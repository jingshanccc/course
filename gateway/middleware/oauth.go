package middleware

import (
	"context"
	"course/config"
	"course/middleware/redis"
	"course/public"
	"course/user-srv/proto/dto"
	"course/user-srv/proto/user"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/server"
	oredis "github.com/go-oauth2/redis/v4"
	"github.com/micro/go-micro/v2/errors"
	"net/http"
	"strings"
	"time"
)

var (
	Services   map[string]interface{}
	AuthServer *server.Server
)

type ClientStore struct {
}

func (c *ClientStore) GetByID(ctx context.Context, id string) (oauth2.ClientInfo, error) {
	if id == "" {
		return nil, nil
	}
	res := &models.Client{}
	err := public.DB.Raw("select * from oauth_client where id = ?", id).Scan(res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}

func init() {
	// 初始化services map
	Services = make(map[string]interface{})

	// 初始化authServer
	manager := manage.NewDefaultManager()
	// use redis token store
	manager.MapTokenStorage(oredis.NewRedisStoreWithCli(redis.RedisClient))
	// use mysql client store
	manager.MapClientStorage(&ClientStore{})
	// use jwt token
	manager.MapAccessGenerate(
		generates.NewJWTAccessGenerate("", []byte(config.JwtKey), jwt.SigningMethodHS512))
	AuthServer = server.NewDefaultServer(manager)
	AuthServer.SetAccessTokenExpHandler(TokenExpHandler)
	AuthServer.SetAllowGetAccessRequest(true)
	AuthServer.SetClientInfoHandler(server.ClientFormHandler)
	// 授权码式授权
	AuthServer.SetUserAuthorizationHandler(UserAuthorizeHandler)
	// 密码式授权
	AuthServer.SetPasswordAuthorizationHandler(ValidPasswordHandler)
}

//SaveServices : 将服务实例存放到gin中
func SaveServices(service []interface{}) gin.HandlerFunc {
	Services[config.UserServiceName] = service[0]
	Services[config.CourseServiceName] = service[1]
	Services[config.FileServiceName] = service[2]
	return func(context *gin.Context) {
		//将实例存到gin.Keys里
		context.Keys = make(map[string]interface{})
		context.Keys[config.UserServiceName] = service[0]
		context.Keys[config.CourseServiceName] = service[1]
		context.Keys[config.FileServiceName] = service[2]
		context.Next()
	}
}

//GetCurrentUser: 获取当前请求登陆的用户
func GetCurrentUser(ctx *gin.Context) (string, *dto.UserDto) {
	token := strings.Split(ctx.Request.Header.Get("Authorization"), " ")[1]
	info, err := AuthServer.Manager.LoadAccessToken(context.Background(), token)
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return "", nil
	} else if usrStr, e := redis.RedisClient.Get(context.Background(), config.UserInfoKey+info.GetUserID()).Result(); e == nil {
		var usr dto.UserDto
		_ = json.Unmarshal([]byte(usrStr), &usr)
		return info.GetUserID(), &usr
	} else {
		return info.GetUserID(), nil
	}
}

//ValidPasswordHandler: 密码式授权-校验密码
func ValidPasswordHandler(username, password string) (userID string, err error) {
	userService := Services[config.UserServiceName].(user.UserService)
	loginUserDto, err := userService.Login(context.Background(), &dto.LoginUserDto{
		LoginName: username,
		Password:  password,
	})
	if err != nil {
		return "", err
	}
	return loginUserDto.Id, nil
}

//UserAuthorizeHandler: 授权码式授权-用户授权环节-从请求中解析出userID
func UserAuthorizeHandler(w http.ResponseWriter, r *http.Request) (userID string, err error) {
	defer func() {
		if err != nil {
			errs := err.(*errors.Error)
			w.WriteHeader(http.StatusOK)
			resp := &public.Response{Success: false, Code: errs.GetCode(), Message: errs.GetDetail(), Content: nil}
			bytes, _ := json.Marshal(resp)
			w.Write(bytes)
		}
	}()
	// 获取参数
	userDto := &dto.LoginUserDto{
		Id:        r.FormValue("id"),
		Name:      r.FormValue("name"),
		LoginName: r.FormValue("login_name"),
		Password:  r.FormValue("password"),
	}
	exception := public.VerifyCaptcha(userDto.Id, userDto.Name)
	if exception != nil {
		err = errors.New(config.UserServiceName, exception.Error(), exception.Code()).(*errors.Error)
		return "", err
	}
	// 登录接口
	userService := Services[config.UserServiceName].(user.UserService)
	loginUserDto, err := userService.Login(r.Context(), userDto)
	if loginUserDto != nil {
		userID = loginUserDto.Id
	}
	// 返回userID
	return userID, err
}

//TokenExpHandler: 令牌过期时间 默认两天
func TokenExpHandler(w http.ResponseWriter, r *http.Request) (exp time.Duration, err error) {
	return config.TokenExpire * time.Hour, nil
}

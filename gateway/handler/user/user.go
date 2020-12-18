package user

import (
	"context"
	"course/config"
	"course/gateway/middleware"
	"course/middleware/redis"
	"course/proto/basic"
	"course/public"
	"course/public/util"
	"course/user-srv/proto/dto"
	"course/user-srv/proto/user"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func GetUserList(ctx *gin.Context) {
	var req dto.PageDto
	if err := ctx.Bind(&req); err == nil {
		userService := ctx.Keys[config.UserServiceName].(user.UserService)
		list, err := userService.List(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}

}

//UserInfo: 获取用户信息
func UserInfo(ctx *gin.Context) {
	currentUser := middleware.GetCurrentUser(ctx)
	userService := ctx.Keys[config.UserServiceName].(user.UserService)
	userDto, err := userService.UserInfo(context.Background(), &basic.String{Str: currentUser})
	public.ResponseAny(ctx, err, userDto)
}

//SavePassword : reset password
func SavePassword(ctx *gin.Context) {
	var req dto.UpdatePass
	if err := ctx.Bind(&req); err == nil {
		req.UserId = middleware.GetCurrentUser(ctx)
		userService := ctx.Keys[config.UserServiceName].(user.UserService)
		result, err := userService.SavePassword(context.Background(), &req)
		public.ResponseAny(ctx, err, result)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//Save : insert or update user
func Save(ctx *gin.Context) {
	var req dto.UserDto
	if err := ctx.Bind(&req); err == nil {
		userService := ctx.Keys[config.UserServiceName].(user.UserService)
		result, err := userService.Save(context.Background(), &req)
		public.ResponseAny(ctx, err, result)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//DeleteUser : 删除用户
func DeleteUser(ctx *gin.Context) {
	var req basic.String
	if err := ctx.Bind(&req); err == nil {
		userService := ctx.Keys[config.UserServiceName].(user.UserService)
		result, err := userService.Delete(context.Background(), &req)
		public.ResponseAny(ctx, err, result)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//Logout : 退出
func Logout(ctx *gin.Context) {
	var req basic.String
	if err := ctx.Bind(&req); err == nil {
		strs := strings.Split(req.Str, "$")
		err := middleware.AuthServer.Manager.RemoveAccessToken(ctx, strs[0])
		err = middleware.AuthServer.Manager.RemoveRefreshToken(ctx, strs[1])
		//userService := ctx.Keys[public.UserServiceName].(user.UserService)
		//result, err := userService.Logout(context.Background(), &req)
		public.ResponseAny(ctx, err, nil)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//UpdateEmail: 更新邮箱
func UpdateEmail(ctx *gin.Context) {
	var req dto.UpdateEmail
	if err := ctx.Bind(&req); err == nil {
		req.UserId = middleware.GetCurrentUser(ctx)
		userService := ctx.Keys[config.UserServiceName].(user.UserService)
		result, err := userService.SaveEmail(context.Background(), &req)
		public.ResponseAny(ctx, err, result)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//SendEmailCode: 发送邮箱验证码
func SendEmailCode(ctx *gin.Context) {
	email := ctx.Query("email")
	redisKey := config.EmailResetEmailCode + email
	var code string
	if c, err := redis.RedisClient.Get(context.Background(), redisKey).Result(); err == nil {
		code = c
	} else {
		code = util.GetVerifyCode()
		redis.RedisClient.Set(context.Background(), redisKey, code, 5*time.Minute)
	}
	err := public.SendHTMLEmail(email, config.TemplatePath+"/email.html", code)
	if err != nil {
		public.ResponseError(ctx, public.NewBusinessException(public.SEND_EMAIL_CODE_ERROR))
		return
	}
	public.ResponseSuccess(ctx, nil)
}

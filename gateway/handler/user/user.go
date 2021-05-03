package user

import (
	"context"
	"gitee.com/jingshanccc/course/gateway/middleware"
	"gitee.com/jingshanccc/course/public"
	"gitee.com/jingshanccc/course/public/config"
	"gitee.com/jingshanccc/course/public/middleware/redis"
	"gitee.com/jingshanccc/course/public/proto/basic"
	"gitee.com/jingshanccc/course/user/proto/dto"
	"gitee.com/jingshanccc/course/user/proto/user"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

func GetUserList(ctx *gin.Context) {
	var req dto.PageDto
	if err := ctx.Bind(&req); err == nil {
		userService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["user"].Name].(user.UserService)
		list, err := userService.List(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}

}

//UserInfo: 获取用户信息
func UserInfo(ctx *gin.Context) {
	var err error
	userId, userDto := middleware.GetCurrentUser(ctx)
	if userDto == nil {
		userService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["user"].Name].(user.UserService)
		userDto, err = userService.UserInfo(context.Background(), &basic.String{Str: userId})
	}
	public.ResponseAny(ctx, err, userDto)
}

//SavePassword : reset password
func SavePassword(ctx *gin.Context) {
	var req dto.UpdatePass
	if err := ctx.Bind(&req); err == nil {
		req.UserId, _ = middleware.GetCurrentUser(ctx)
		userService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["user"].Name].(user.UserService)
		result, err := userService.SavePassword(context.Background(), &req)
		public.ResponseAny(ctx, err, result)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//SaveUserInfo: 更新当前登录用户信息
func SaveUserInfo(ctx *gin.Context) {
	var req dto.UserDto
	if err := ctx.Bind(&req); err == nil {
		_, usr := middleware.GetCurrentUser(ctx)
		req.UpdateBy = usr.LoginName
		userService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["user"].Name].(user.UserService)
		result, err := userService.SaveUserInfo(context.Background(), &req)
		public.ResponseAny(ctx, err, result)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//Save : insert or update user
func Save(ctx *gin.Context) {
	var req dto.UserDto
	if err := ctx.Bind(&req); err == nil {
		_, usr := middleware.GetCurrentUser(ctx)
		req.UpdateBy = usr.LoginName
		userService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["user"].Name].(user.UserService)
		result, err := userService.Save(context.Background(), &req)
		public.ResponseAny(ctx, err, result)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//DeleteUser : 删除用户
func DeleteUser(ctx *gin.Context) {
	var req basic.StringList
	if err := ctx.Bind(&req); err == nil {
		userService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["user"].Name].(user.UserService)
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
		userId, _ := middleware.GetCurrentUser(ctx)
		redis.RedisClient.Del(ctx, config.Conf.Services["user"].Others["userInfoKey"].(string)+userId)
		strs := strings.Split(req.Str, "$")
		err = middleware.AuthServer.Manager.RemoveAccessToken(ctx, strs[0])
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
		req.UserId, _ = middleware.GetCurrentUser(ctx)
		userService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["user"].Name].(user.UserService)
		result, err := userService.SaveEmail(context.Background(), &req)
		public.ResponseAny(ctx, err, result)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//SendEmailCode: 发送邮箱验证码
func SendEmailCode(ctx *gin.Context) {
	email := ctx.Query("email")
	others := config.Conf.Services["user"].Others
	err := public.SendEmailCode(time.Duration(others["emailCodeExpire"].(int))*time.Minute, email, others["emailResetKey"].(string)+email, others["emailTemplatePath"].(string)+"/email.html")
	public.ResponseAny(ctx, err, nil)
}

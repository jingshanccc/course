package user

import (
	"context"
	"course/config"
	"course/gateway/middleware"
	"course/proto/basic"
	"course/public"
	"course/user-srv/proto/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func GetUserList(ctx *gin.Context) {
	var req user.PageDto
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
	currentUser, err := middleware.GetCurrentUser(ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusUnauthorized, err)
		return
	}
	userService := ctx.Keys[config.UserServiceName].(user.UserService)
	userDto, err := userService.UserInfo(context.Background(), &basic.String{Str: currentUser})
	public.ResponseAny(ctx, err, userDto)
}

//SavePassword : reset password
func SavePassword(ctx *gin.Context) {
	var req user.UserDto
	if err := ctx.Bind(&req); err == nil {
		userService := ctx.Keys[config.UserServiceName].(user.UserService)
		result, err := userService.SavePassword(context.Background(), &req)
		public.ResponseAny(ctx, err, result)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//Save : insert or update user
func Save(ctx *gin.Context) {
	var req user.UserDto
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

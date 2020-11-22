package handler

import (
	"context"
	"course/public"
	"course/user-srv/proto/user"
	"github.com/gin-gonic/gin"
)

func GetUserList(ctx *gin.Context) {
	var req user.PageDto
	if err := ctx.Bind(&req); err == nil {
		userService := ctx.Keys[public.UserServiceName].(user.UserService)
		list, err := userService.List(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}

}

func Login(ctx *gin.Context) {
	var req user.User
	if err := ctx.Bind(&req); err == nil {
		//图形验证码校验-user的id作为验证码id name作为验证码值
		exception := VerifyCaptcha(req.Id, req.Name)
		if exception.Code() == int32(public.OK) {
			userService := ctx.Keys[public.UserServiceName].(user.UserService)
			loginUserDto, err := userService.Login(context.Background(), &req)
			public.ResponseAny(ctx, err, loginUserDto)
		} else {
			public.ResponseError(ctx, exception)
		}
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//SavePassword : reset password
func SavePassword(ctx *gin.Context) {
	var req user.User
	if err := ctx.Bind(&req); err == nil {
		userService := ctx.Keys[public.UserServiceName].(user.UserService)
		result, err := userService.SavePassword(context.Background(), &req)
		public.ResponseAny(ctx, err, result)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//Save : insert or update user
func Save(ctx *gin.Context) {
	var req user.User
	if err := ctx.Bind(&req); err == nil {
		userService := ctx.Keys[public.UserServiceName].(user.UserService)
		result, err := userService.Save(context.Background(), &req)
		public.ResponseAny(ctx, err, result)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//DeleteUser : 删除用户
func DeleteUser(ctx *gin.Context) {
	var req user.User
	if err := ctx.Bind(&req); err == nil {
		userService := ctx.Keys[public.UserServiceName].(user.UserService)
		result, err := userService.Delete(context.Background(), &req)
		public.ResponseAny(ctx, err, result)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//Logout : 退出
func Logout(ctx *gin.Context) {
	var req user.LoginUserDto
	if err := ctx.Bind(&req); err == nil {
		userService := ctx.Keys[public.UserServiceName].(user.UserService)
		result, err := userService.Logout(context.Background(), &req)
		public.ResponseAny(ctx, err, result)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

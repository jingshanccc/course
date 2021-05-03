package user

import (
	"context"
	"gitee.com/jingshanccc/course/gateway/middleware"
	"gitee.com/jingshanccc/course/public"
	"gitee.com/jingshanccc/course/public/config"
	"gitee.com/jingshanccc/course/public/proto/basic"
	"gitee.com/jingshanccc/course/user/proto/dto"
	"gitee.com/jingshanccc/course/user/proto/user"
	"github.com/gin-gonic/gin"
)

func EmailRegisterCode(ctx *gin.Context) {
	var req basic.String
	if err := ctx.Bind(&req); err == nil {
		userService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["user"].Name].(user.UserService)
		data, err := userService.SendEmailCode(context.Background(), &req)
		public.ResponseAny(ctx, err, data)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}
func MemberRegister(ctx *gin.Context) {
	var req dto.MemberRegisterDto
	if err := ctx.Bind(&req); err == nil {
		userService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["user"].Name].(user.UserService)
		data, err := userService.MemberRegister(context.Background(), &req)
		public.ResponseAny(ctx, err, data)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

func MemberInfo(ctx *gin.Context) {
	userId, userDto := middleware.GetCurrentMember(ctx)
	var err error
	if userDto == nil {
		userService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["user"].Name].(user.UserService)
		userDto, err = userService.MemberInfo(context.Background(), &basic.String{Str: userId})
	}
	public.ResponseAny(ctx, err, userDto)
}

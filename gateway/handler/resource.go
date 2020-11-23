package handler

import (
	"context"
	"course/public"
	"course/user-srv/proto/dto"
	"course/user-srv/proto/user"
	"github.com/gin-gonic/gin"
)

func LoadTree(ctx *gin.Context) {
	var req dto.String
	if err := ctx.Bind(&req); err == nil {
		resourceService := ctx.Keys[public.UserServiceName].(user.UserService)
		list, err := resourceService.LoadTree(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

func SaveJson(ctx *gin.Context) {
	var req dto.String
	if err := ctx.Bind(&req); err == nil {
		resourceService := ctx.Keys[public.UserServiceName].(user.UserService)
		str, err := resourceService.SaveJson(context.Background(), &req)
		public.ResponseAny(ctx, err, str)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

func Delete(ctx *gin.Context) {
	var req dto.String
	if err := ctx.Bind(&req); err == nil {
		resourceService := ctx.Keys[public.UserServiceName].(user.UserService)
		str, err := resourceService.DeleteResource(context.Background(), &req)
		public.ResponseAny(ctx, err, str)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

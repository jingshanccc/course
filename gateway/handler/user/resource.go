package user

import (
	"context"
	"course/config"
	"course/gateway/middleware"
	"course/proto/basic"
	"course/public"
	"course/user-srv/proto/user"
	"github.com/gin-gonic/gin"
)

//LoadMenus: 加载前端侧边栏菜单
func LoadMenus(ctx *gin.Context) {
	currentUser, _ := middleware.GetCurrentUser(ctx)

	resourceService := ctx.Keys[config.UserServiceName].(user.UserService)
	list, err := resourceService.LoadMenus(context.Background(), &basic.String{Str: currentUser})
	public.ResponseAny(ctx, err, list)
}

//LoadTree: 获取权限树的一个节点的所有子节点
func LoadTree(ctx *gin.Context) {
	var req basic.Integer
	if err := ctx.Bind(&req); err == nil {
		resourceService := ctx.Keys[config.UserServiceName].(user.UserService)
		list, err := resourceService.LoadTree(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//MenuChild: 获取传入权限的所有子权限 ID 包含传入权限自身
func MenuChild(ctx *gin.Context) {
	var req basic.Integer
	if err := ctx.Bind(&req); err == nil {
		resourceService := ctx.Keys[config.UserServiceName].(user.UserService)
		list, err := resourceService.MenuChild(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

func SaveJson(ctx *gin.Context) {
	var req basic.String
	if err := ctx.Bind(&req); err == nil {
		resourceService := ctx.Keys[config.UserServiceName].(user.UserService)
		str, err := resourceService.SaveJson(context.Background(), &req)
		public.ResponseAny(ctx, err, str)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

func Delete(ctx *gin.Context) {
	var req basic.Integer
	if err := ctx.Bind(&req); err == nil {
		resourceService := ctx.Keys[config.UserServiceName].(user.UserService)
		str, err := resourceService.DeleteResource(context.Background(), &req)
		public.ResponseAny(ctx, err, str)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

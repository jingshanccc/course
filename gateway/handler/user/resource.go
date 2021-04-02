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

//LoadMenus: 加载前端侧边栏菜单
func LoadMenus(ctx *gin.Context) {
	currentUser, _ := middleware.GetCurrentUser(ctx)

	resourceService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["user"].Name].(user.UserService)
	list, err := resourceService.LoadMenus(context.Background(), &basic.String{Str: currentUser})
	public.ResponseAny(ctx, err, list)
}

//LoadTree: 获取权限树的一个节点的所有子节点
func LoadTree(ctx *gin.Context) {
	var req basic.Integer
	if err := ctx.Bind(&req); err == nil {
		resourceService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["user"].Name].(user.UserService)
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
		resourceService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["user"].Name].(user.UserService)
		list, err := resourceService.MenuChild(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//MenuList: 权限管理 权限表格数据
func MenuList(ctx *gin.Context) {
	var req dto.ResourcePageDto
	if err := ctx.Bind(&req); err == nil {
		resourceService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["user"].Name].(user.UserService)
		list, err := resourceService.MenuList(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//MenuParent: 获取传入权限的所有父权限和同级权限
func MenuParent(ctx *gin.Context) {
	var req basic.IntegerList
	if err := ctx.Bind(&req); err == nil {
		resourceService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["user"].Name].(user.UserService)
		list, err := resourceService.MenuParent(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//SaveResource: 新增/更新资源
func SaveResource(ctx *gin.Context) {
	var req dto.ResourceDto
	if err := ctx.Bind(&req); err == nil {
		req.Label = req.Title
		_, u := middleware.GetCurrentUser(ctx)
		req.UpdateBy = u.LoginName
		resourceService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["user"].Name].(user.UserService)
		str, err := resourceService.SaveResource(context.Background(), &req)
		public.ResponseAny(ctx, err, str)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

func Delete(ctx *gin.Context) {
	var req basic.IntegerList
	if err := ctx.Bind(&req); err == nil {
		resourceService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["user"].Name].(user.UserService)
		str, err := resourceService.DeleteResource(context.Background(), &req)
		public.ResponseAny(ctx, err, str)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

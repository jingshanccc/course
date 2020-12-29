package user

import (
	"context"
	"course/config"
	"course/gateway/middleware"
	"course/proto/basic"
	"course/public"
	"course/user-srv/dao"
	"course/user-srv/proto/dto"
	"course/user-srv/proto/user"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

//SaveRole: gateway handler 保存角色
func SaveRole(ctx *gin.Context) {
	var req dto.RoleDto
	if err := ctx.Bind(&req); err == nil {
		roleService := ctx.Keys[config.UserServiceName].(user.UserService)
		list, err := roleService.SaveRole(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//AllRole: gateway handler 获取所有角色
func AllRole(ctx *gin.Context) {
	var req basic.String
	if err := ctx.Bind(&req); err == nil {
		roleService := ctx.Keys[config.UserServiceName].(user.UserService)
		list, err := roleService.AllRole(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//GetRole: gateway handler 获取传入 ID 角色
func GetRole(ctx *gin.Context) {
	var req basic.String
	if err := ctx.Bind(&req); err == nil {
		roleService := ctx.Keys[config.UserServiceName].(user.UserService)
		res, err := roleService.GetRole(context.Background(), &req)
		public.ResponseAny(ctx, err, res)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//RoleList: gateway handler 获取角色列表
func RoleList(ctx *gin.Context) {
	var req dto.RolePageDto
	if err := ctx.Bind(&req); err == nil {
		roleService := ctx.Keys[config.UserServiceName].(user.UserService)
		list, err := roleService.RoleList(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//RoleLevel: gateway handler 获取用户角色级别
func RoleLevel(ctx *gin.Context) {
	userId, curUser := middleware.GetCurrentUser(ctx)
	roleService := ctx.Keys[config.UserServiceName].(user.UserService)
	if curUser == nil {
		level, err := roleService.RoleLevel(ctx, &basic.String{Str: userId})
		public.ResponseAny(ctx, err, level)
	} else {
		var roles []dao.Role
		_ = json.Unmarshal([]byte(curUser.Roles), &roles)
		public.ResponseSuccess(ctx, roles[0].Level)
	}

}

//DeleteRole: gateway handler 删除角色
func DeleteRole(ctx *gin.Context) {
	var req basic.StringList
	if err := ctx.Bind(&req); err == nil {
		roleService := ctx.Keys[config.UserServiceName].(user.UserService)
		list, err := roleService.DeleteRole(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//SaveRoleResource: gateway handler  保存角色-资源关联记录
func SaveRoleResource(ctx *gin.Context) {
	var req dto.RoleDto
	if err := ctx.Bind(&req); err == nil {
		roleService := ctx.Keys[config.UserServiceName].(user.UserService)
		list, err := roleService.SaveRoleResource(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//ListRoleResource: gateway handler 按角色加载权限
func ListRoleResource(ctx *gin.Context) {
	var req basic.String
	if err := ctx.Bind(&req); err == nil {
		roleService := ctx.Keys[config.UserServiceName].(user.UserService)
		list, err := roleService.ListRoleResource(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//SaveRoleUser: gateway handler 保存角色的所有用户
func SaveRoleUser(ctx *gin.Context) {
	var req dto.RoleDto
	if err := ctx.Bind(&req); err == nil {
		roleService := ctx.Keys[config.UserServiceName].(user.UserService)
		list, err := roleService.SaveRoleUser(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//ListRoleUser: gateway handler 按角色加载用户
func ListRoleUser(ctx *gin.Context) {
	var req basic.String
	if err := ctx.Bind(&req); err == nil {
		roleService := ctx.Keys[config.UserServiceName].(user.UserService)
		list, err := roleService.ListRoleUser(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

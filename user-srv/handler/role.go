package handler

import (
	"context"
	"course/config"
	"course/proto/basic"
	"course/public"
	"course/public/util"
	"course/user-srv/proto/dto"
	"github.com/micro/go-micro/v2/errors"
)

//RoleList : 获取所有角色
func (u *UserServiceHandler) RoleList(ctx context.Context, in *dto.RolePageDto, out *dto.RolePageDto) error {
	list, exception := roleDao.List(ctx, in)
	if exception != nil {
		return errors.New(config.UserServiceName, exception.Error(), exception.Code())
	}
	_ = util.CopyProperties(out, in)
	out.Rows = list
	return nil
}

//SaveRole: 保存一种角色
func (u *UserServiceHandler) SaveRole(ctx context.Context, in *dto.RoleDto, out *dto.RoleDto) error {
	roleDto, err := roleDao.Save(ctx, in)
	if err.Code() != int32(public.OK) {
		return errors.New(config.UserServiceName, err.Error(), err.Code())
	}
	_ = util.CopyProperties(out, roleDto)
	return nil
}

//DeleteRole: 删除一种角色
func (u *UserServiceHandler) DeleteRole(ctx context.Context, in *basic.String, out *basic.String) error {
	exception := roleDao.Delete(ctx, in.Str)
	if exception != nil {
		return errors.New(config.UserServiceName, exception.Error(), exception.Code())
	}
	return nil
}

//SaveRoleResource: 保存角色权限
func (u *UserServiceHandler) SaveRoleResource(ctx context.Context, in *dto.RoleDto, out *dto.RoleDto) error {
	exception := roleDao.SaveRoleResource(ctx, in)
	if exception != nil {
		return errors.New(config.UserServiceName, exception.Error(), exception.Code())
	}
	return nil
}

//ListRoleResource: 获取角色权限
func (u *UserServiceHandler) ListRoleResource(ctx context.Context, in *basic.String, out *basic.StringList) error {
	resources, exception := roleDao.ListRoleResource(ctx, in.Str)
	if exception != nil {
		return errors.New(config.UserServiceName, exception.Error(), exception.Code())
	}
	out.Rows = resources
	return nil
}

//SaveRoleUser: 保存某个角色的所有用户
func (u *UserServiceHandler) SaveRoleUser(ctx context.Context, in *dto.RoleDto, out *dto.RoleDto) error {
	exception := roleDao.SaveRoleUser(ctx, in)
	if exception != nil {
		return errors.New(config.UserServiceName, exception.Error(), exception.Code())
	}
	return nil
}

//ListRoleUser: 获取某个角色的所有用户
func (u *UserServiceHandler) ListRoleUser(ctx context.Context, in *basic.String, out *basic.StringList) error {
	users, exception := roleDao.ListRoleUser(ctx, in.Str)
	if exception != nil {
		return errors.New(config.UserServiceName, exception.Error(), exception.Code())
	}
	out.Rows = users
	return nil
}

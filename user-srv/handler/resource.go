package handler

import (
	"context"
	"course/config"
	"course/proto/basic"
	"course/user-srv/proto/dto"
	"github.com/micro/go-micro/v2/errors"
)

//LoadMenus : 加载后台管理系统用户权限树
func (u *UserServiceHandler) LoadMenus(ctx context.Context, in *basic.String, out *dto.ResourceDtoList) error {
	resourceDtos, exception := resourceDao.FindUserResources(ctx, in.Str)
	if exception != nil {
		return errors.New(config.UserServiceName, exception.Error(), exception.Code())
	}
	out.Rows = resourceDtos
	return nil
}

//LoadTree : 加载权限树
func (u *UserServiceHandler) LoadTree(ctx context.Context, in *basic.String, out *dto.ResourceDtoList) error {
	resourceDtos, exception := resourceDao.LoadTree(ctx)
	if exception != nil {
		return errors.New(config.UserServiceName, exception.Error(), exception.Code())
	}
	out.Rows = resourceDtos
	return nil
}

// SaveJson : 保存权限树
func (u *UserServiceHandler) SaveJson(ctx context.Context, in *basic.String, out *basic.String) error {
	exception := resourceDao.SaveJson(ctx, in.Str)
	if exception != nil {
		return errors.New(config.UserServiceName, exception.Error(), exception.Code())
	}
	return nil
}

// DeleteResource: 删除权限
func (u *UserServiceHandler) DeleteResource(ctx context.Context, in *basic.Integer, out *basic.String) error {
	exception := resourceDao.Delete(ctx, in.Id)
	if exception != nil {
		return errors.New(config.UserServiceName, exception.Error(), exception.Code())
	}
	return nil
}

package handler

import (
	"context"
	"course/proto/basic"
	"course/public"
	"course/user-srv/proto/dto"
	"github.com/micro/go-micro/v2/errors"
)

//LoadTree : 加载权限树
func (u *UserServiceHandler) LoadTree(ctx context.Context, in *basic.String, out *dto.ResourceDtoList) error {
	resourceDtos, exception := resourceDao.LoadTree(ctx)
	if exception.Code() != int32(public.OK) {
		return errors.New(public.UserServiceName, exception.Error(), exception.Code())
	}
	out.Rows = resourceDtos
	return nil
}

// SaveJson : 保存权限树
func (u *UserServiceHandler) SaveJson(ctx context.Context, in *basic.String, out *basic.String) error {
	exception := resourceDao.SaveJson(ctx, in.Str)
	if exception.Code() != int32(public.OK) {
		return errors.New(public.UserServiceName, exception.Error(), exception.Code())
	}
	return nil
}

// DeleteResource: 删除权限
func (u *UserServiceHandler) DeleteResource(ctx context.Context, in *basic.String, out *basic.String) error {
	exception := resourceDao.Delete(ctx, in.Str)
	if exception.Code() != int32(public.OK) {
		return errors.New(public.UserServiceName, exception.Error(), exception.Code())
	}
	return nil
}

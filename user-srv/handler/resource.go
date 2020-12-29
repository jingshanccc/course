package handler

import (
	"context"
	"course/config"
	"course/proto/basic"
	"course/public"
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

//LoadTree: 加载权限树中的节点
func (u *UserServiceHandler) LoadTree(ctx context.Context, in *basic.Integer, out *dto.ResourceDtoList) error {
	resourceDtos, exception := resourceDao.GetByParent(ctx, in.Id)
	if exception != nil {
		return errors.New(config.UserServiceName, exception.Error(), exception.Code())
	}
	out.Rows = resourceDtos
	return nil
}

//MenuChild: 获取传入权限的子权限 id
func (u *UserServiceHandler) MenuChild(ctx context.Context, in *basic.Integer, out *basic.IntegerList) error {
	// 通过一个父 id 获取其所有的子权限
	resource, exception := resourceDao.SelectById(ctx, in.Id)
	if exception != nil {
		return errors.New(config.UserServiceName, exception.Error(), exception.Code())
	}
	resourceSet := public.NewHashSet(resource)
	resourceDtos, exception := resourceDao.GetByParent(ctx, in.Id)
	if exception != nil {
		return errors.New(config.UserServiceName, exception.Error(), exception.Code())
	}
	resourceDao.GetChildTree(ctx, resourceDtos, resourceSet)
	var res []int32
	for _, r := range resourceSet.Values() {
		re := r.(*dto.ResourceDto)
		res = append(res, re.Id)
	}
	out.Ids = res
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

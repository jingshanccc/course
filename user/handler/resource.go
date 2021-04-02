package handler

import (
	"context"
	"gitee.com/jingshanccc/course/public"
	"gitee.com/jingshanccc/course/public/config"
	"gitee.com/jingshanccc/course/public/proto/basic"
	"gitee.com/jingshanccc/course/user/proto/dto"
	"github.com/micro/go-micro/v2/errors"
)

//LoadMenus : 加载后台管理系统用户权限树
func (u *UserServiceHandler) LoadMenus(ctx context.Context, in *basic.String, out *dto.ResourceDtoList) error {
	resourceDtos, exception := resourceDao.FindUserResources(ctx, in.Str)
	if exception != nil {
		return errors.New(config.Conf.Services["user"].Name, exception.Error(), exception.Code())
	}
	out.Rows = resourceDtos
	return nil
}

//LoadTree: 加载权限树中的节点
func (u *UserServiceHandler) LoadTree(ctx context.Context, in *basic.Integer, out *dto.ResourceDtoList) error {
	resourceDtos, exception := resourceDao.GetByParent(ctx, in.Id)
	if exception != nil {
		return errors.New(config.Conf.Services["user"].Name, exception.Error(), exception.Code())
	}
	out.Rows = resourceDtos
	return nil
}

//MenuChild: 获取传入权限的子权限 id
func (u *UserServiceHandler) MenuChild(ctx context.Context, in *basic.Integer, out *basic.IntegerList) error {
	// 通过一个父 id 获取其所有的子权限
	resourceSet := public.NewHashSet()
	exception := resourceDao.GetChildMenus(ctx, []int32{in.Id}, resourceSet)
	if exception != nil {
		return errors.New(config.Conf.Services["user"].Name, exception.Error(), exception.Code())
	}
	var res []int32
	for _, r := range resourceSet.Values() {
		re := r.(*dto.ResourceDto)
		res = append(res, re.Id)
	}
	out.Ids = res
	return nil
}

//MenuList: 权限列表
func (u *UserServiceHandler) MenuList(ctx context.Context, in *dto.ResourcePageDto, out *dto.ResourcePageDto) error {
	var list []*dto.ResourceDto
	var exception *public.BusinessException
	var count int64
	if in.Blurry == "" && in.CreateTime == nil {
		list, exception = resourceDao.GetByParent(ctx, in.Parent)
	} else {
		count, list, exception = resourceDao.List(ctx, in)
	}
	if exception != nil {
		return errors.New(config.Conf.Services["user"].Name, exception.Error(), exception.Code())
	}
	for _, resourceDto := range list {
		if resourceDto.SubCount > 0 {
			resourceDto.HasChildren = true
		}
	}
	out.Rows = list
	out.Total = count
	return nil
}

//MenuParent: 获取传入权限ID的所有同级和父级权限
func (u *UserServiceHandler) MenuParent(ctx context.Context, in *basic.IntegerList, out *dto.ResourceDtoList) error {
	var exception *public.BusinessException
	var list []*dto.ResourceDto
	if in.Ids == nil {
		list, exception = resourceDao.GetByParent(ctx, 0)
	} else {
		list, exception = resourceDao.GetSuperior(ctx, in.Ids)
	}
	if exception != nil {
		return errors.New(config.Conf.Services["user"].Name, exception.Error(), exception.Code())
	}
	out.Rows = list
	return nil
}

// SaveJson : 保存权限树
func (u *UserServiceHandler) SaveResource(ctx context.Context, in *dto.ResourceDto, out *basic.String) error {
	exception := resourceDao.Save(ctx, in)
	if exception != nil {
		return errors.New(config.Conf.Services["user"].Name, exception.Error(), exception.Code())
	}
	return nil
}

// DeleteResource: 删除权限
func (u *UserServiceHandler) DeleteResource(ctx context.Context, in *basic.IntegerList, out *basic.String) error {
	exception := resourceDao.Delete(ctx, in.Ids)
	if exception != nil {
		return errors.New(config.Conf.Services["user"].Name, exception.Error(), exception.Code())
	}
	return nil
}

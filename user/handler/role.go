package handler

import (
	"context"
	"gitee.com/jingshanccc/course/public"
	"gitee.com/jingshanccc/course/public/config"
	"gitee.com/jingshanccc/course/public/proto/basic"
	"gitee.com/jingshanccc/course/public/util"
	"gitee.com/jingshanccc/course/user/proto/dto"
	"github.com/micro/go-micro/v2/errors"
)

//AllRole : 获取所有角色
func (u *UserServiceHandler) AllRole(ctx context.Context, in *basic.String, out *dto.RoleDtoList) error {
	list, exception := roleDao.All(ctx)
	if exception != nil {
		return errors.New(config.Conf.Services["user"].Name, exception.Error(), exception.Code())
	}
	out.Rows = list
	return nil
}

//GetRole : 获取传入 ID 角色
func (u *UserServiceHandler) GetRole(ctx context.Context, in *basic.String, out *dto.RoleDto) error {
	r, exception := roleDao.SelectById(ctx, in.Str)
	if exception != nil {
		return errors.New(config.Conf.Services["user"].Name, exception.Error(), exception.Code())
	}
	_ = util.CopyProperties(out, r)
	return nil
}

//RoleList : 获取所有角色
func (u *UserServiceHandler) RoleList(ctx context.Context, in *dto.RolePageDto, out *dto.RolePageDto) error {
	total, list, exception := roleDao.List(ctx, in)
	if exception != nil {
		return errors.New(config.Conf.Services["user"].Name, exception.Error(), exception.Code())
	}
	_ = util.CopyProperties(out, in)
	out.Total = total
	out.Rows = list
	return nil
}

//RoleLevel : 获取所有角色
func (u *UserServiceHandler) RoleLevel(ctx context.Context, in *basic.String, out *basic.Integer) error {
	roles, exception := roleDao.SelectByUserId(in.Str)
	if exception != nil {
		return errors.New(config.Conf.Services["user"].Name, exception.Error(), exception.Code())
	}
	_ = util.CopyProperties(out, in)
	out.Id = roles[0].Level
	return nil
}

//SaveRole: 保存一种角色
func (u *UserServiceHandler) SaveRole(ctx context.Context, in *dto.RoleDto, out *dto.RoleDto) error {
	roleDto, err := roleDao.Save(ctx, in)
	if err != nil {
		return errors.New(config.Conf.Services["user"].Name, err.Error(), err.Code())
	}
	_ = util.CopyProperties(out, roleDto)
	return nil
}

//DeleteRole: 删除一种角色
func (u *UserServiceHandler) DeleteRole(ctx context.Context, in *basic.StringList, out *basic.String) error {
	var exception *public.BusinessException
	if count := userDao.CountByRoles(ctx, in.Rows); count > 0 {
		exception = public.NewBusinessException(public.ROLE_USER_EXIST)
	} else {
		exception = roleDao.Delete(ctx, in.Rows)
	}
	if exception != nil {
		return errors.New(config.Conf.Services["user"].Name, exception.Error(), exception.Code())
	}
	return nil
}

//SaveRoleResource: 保存角色权限
func (u *UserServiceHandler) SaveRoleResource(ctx context.Context, in *dto.RoleDto, out *dto.RoleDto) error {
	exception := roleDao.SaveRoleResource(ctx, in)
	if exception != nil {
		return errors.New(config.Conf.Services["user"].Name, exception.Error(), exception.Code())
	}
	return nil
}

//ListRoleResource: 获取角色权限
func (u *UserServiceHandler) ListRoleResource(ctx context.Context, in *basic.String, out *basic.IntegerList) error {
	resources, exception := roleDao.ListRoleResource(ctx, in.Str)
	if exception != nil {
		return errors.New(config.Conf.Services["user"].Name, exception.Error(), exception.Code())
	}
	out.Ids = resources
	return nil
}

//SaveRoleUser: 保存某个角色的所有用户
func (u *UserServiceHandler) SaveRoleUser(ctx context.Context, in *dto.RoleDto, out *dto.RoleDto) error {
	exception := roleDao.SaveRoleUser(ctx, in)
	if exception != nil {
		return errors.New(config.Conf.Services["user"].Name, exception.Error(), exception.Code())
	}
	return nil
}

//ListRoleUser: 获取某个角色的所有用户
func (u *UserServiceHandler) ListRoleUser(ctx context.Context, in *basic.String, out *basic.StringList) error {
	users, exception := roleDao.ListRoleUser(ctx, in.Str)
	if exception != nil {
		return errors.New(config.Conf.Services["user"].Name, exception.Error(), exception.Code())
	}
	out.Rows = users
	return nil
}

package handler

import (
	"context"
	"course/middleware/redis"
	"course/proto/basic"
	"course/public"
	"course/public/util"
	"course/user-srv/dao"
	"course/user-srv/proto/dto"
	"course/user-srv/proto/user"
	"crypto/md5"
	"fmt"
	"github.com/micro/go-micro/v2/errors"
	"log"
	"time"
)

/*
 when service handler return not nil, the default err code will be 500.
 it leads to the gateway response internal server error.
 so when service handler throws error, we should generate a *errors.Error(go-micro) with a code not 500,
 so that we can let the gateway response in our way.
*/

var (
	userDao     = &dao.UserDao{}
	resourceDao = &dao.ResourceDao{}
	roleDao     = &dao.RoleDao{}
)

type UserServiceHandler struct {
}

func (u *UserServiceHandler) List(ctx context.Context, in *user.PageDto, out *user.PageDto) error {
	users, err := userDao.List(ctx, in)
	if err.Code() != int32(public.OK) {
		return errors.New(public.UserServiceName, err.Error(), err.Code())
	}
	_ = util.CopyProperties(out, in)
	out.Users = users

	return nil
}

func (u *UserServiceHandler) Save(ctx context.Context, in *user.UserDto, out *user.UserDto) error {
	outUser, err := userDao.Save(ctx, in)
	if err.Code() != int32(public.OK) {
		return errors.New(public.UserServiceName, err.Error(), err.Code())
	}
	_ = util.CopyProperties(out, outUser)
	return nil
}

func (u *UserServiceHandler) Delete(ctx context.Context, in *basic.String, out *basic.String) error {
	exception := userDao.Delete(ctx, in.Str)
	if exception.Code() != int32(public.OK) {
		return errors.New(public.UserServiceName, exception.Error(), exception.Code())
	}
	return nil
}

//SavePassword : reset password
func (u *UserServiceHandler) SavePassword(ctx context.Context, in *user.UserDto, out *user.UserDto) error {
	str := fmt.Sprintf("%x", md5.Sum([]byte(in.Password)))
	in.Password = str
	password, err := userDao.SavePassword(ctx, in)
	if err.Code() != int32(public.OK) {
		return errors.New(public.UserServiceName, err.Error(), err.Code())
	}
	_ = util.CopyProperties(out, in)
	out.Password = password
	return nil
}

func (u *UserServiceHandler) Login(ctx context.Context, in *user.UserDto, out *dto.LoginUserDto) error {
	//better than in.Password = fmt.Sprintf("%x", sum)
	//in.Password = *(*string)(unsafe.Pointer(&sum))
	str := fmt.Sprintf("%x", md5.Sum([]byte(in.Password)))
	in.Password = str
	loginUserDto, err := userDao.Login(ctx, in)
	if err.Code() != int32(public.OK) {
		return errors.New(public.UserServiceName, err.Error(), err.Code())
	}
	_ = util.CopyProperties(out, loginUserDto)
	// 存储token
	jsonString, _ := util.ToJsonString(loginUserDto)
	redis.RedisClient.Set(loginUserDto.Token, jsonString, time.Second*3600)
	return nil
}

func (u *UserServiceHandler) Logout(ctx context.Context, in *basic.String, out *basic.String) error {
	redis.RedisClient.Del(in.Str)
	log.Println("从redis中删除token: ", in.Str)
	return nil
}

//---------- 权限管理 -------------

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

//RoleList : 获取所有角色
func (u *UserServiceHandler) RoleList(ctx context.Context, in *dto.RolePageDto, out *dto.RolePageDto) error {
	list, exception := roleDao.List(ctx, in)
	if exception.Code() != int32(public.OK) {
		return errors.New(public.UserServiceName, exception.Error(), exception.Code())
	}
	_ = util.CopyProperties(out, in)
	out.Rows = list
	return nil
}

//SaveRole: 保存一种角色
func (u *UserServiceHandler) SaveRole(ctx context.Context, in *dto.RoleDto, out *dto.RoleDto) error {
	roleDto, err := roleDao.Save(ctx, in)
	if err.Code() != int32(public.OK) {
		return errors.New(public.UserServiceName, err.Error(), err.Code())
	}
	_ = util.CopyProperties(out, roleDto)
	return nil
}

//DeleteRole: 删除一种角色
func (u *UserServiceHandler) DeleteRole(ctx context.Context, in *basic.String, out *basic.String) error {
	exception := roleDao.Delete(ctx, in.Str)
	if exception.Code() != int32(public.OK) {
		return errors.New(public.UserServiceName, exception.Error(), exception.Code())
	}
	return nil
}

//SaveRoleResource: 保存角色权限
func (u *UserServiceHandler) SaveRoleResource(ctx context.Context, in *dto.RoleDto, out *dto.RoleDto) error {
	exception := roleDao.SaveRoleResource(ctx, in)
	if exception.Code() != int32(public.OK) {
		return errors.New(public.UserServiceName, exception.Error(), exception.Code())
	}
	return nil
}

//ListRoleResource: 获取角色权限
func (u *UserServiceHandler) ListRoleResource(ctx context.Context, in *basic.String, out *basic.StringList) error {
	resources, exception := roleDao.ListRoleResource(ctx, in.Str)
	if exception.Code() != int32(public.OK) {
		return errors.New(public.UserServiceName, exception.Error(), exception.Code())
	}
	out.Rows = resources
	return nil
}

//SaveRoleUser: 保存某个角色的所有用户
func (u *UserServiceHandler) SaveRoleUser(ctx context.Context, in *dto.RoleDto, out *dto.RoleDto) error {
	exception := roleDao.SaveRoleUser(ctx, in)
	if exception.Code() != int32(public.OK) {
		return errors.New(public.UserServiceName, exception.Error(), exception.Code())
	}
	return nil
}

//ListRoleUser: 获取某个角色的所有用户
func (u *UserServiceHandler) ListRoleUser(ctx context.Context, in *basic.String, out *basic.StringList) error {
	users, exception := roleDao.ListRoleUser(ctx, in.Str)
	if exception.Code() != int32(public.OK) {
		return errors.New(public.UserServiceName, exception.Error(), exception.Code())
	}
	out.Rows = users
	return nil
}

package handler

import (
	"context"
	"course/middleware/redis"
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
)

type UserServiceHandler struct {
}

func (u *UserServiceHandler) List(ctx context.Context, in *user.PageDto, out *user.PageDto) error {
	users, err := userDao.List(ctx, in)
	if err.Code() != int32(public.OK) {
		return errors.New(public.UserServiceName, err.Error(), err.Code())
	}
	out.PageSize = in.PageSize
	out.PageNum = in.PageNum
	out.SortBy = in.SortBy
	out.Asc = in.Asc
	out.Users = users

	return nil
}

func (h *UserServiceHandler) Save(ctx context.Context, in *user.UserDto, out *user.UserDto) error {
	outUser, err := userDao.Save(ctx, in)
	if err.Code() != int32(public.OK) {
		return errors.New(public.UserServiceName, err.Error(), err.Code())
	}
	out.Id = outUser.Id
	out.LoginName = outUser.LoginName
	out.Name = outUser.Name
	out.Password = outUser.Password
	return nil
}

func (h *UserServiceHandler) Delete(ctx context.Context, in *dto.String, out *dto.String) error {
	exception := userDao.Delete(ctx, in.Str)
	if exception.Code() != int32(public.OK) {
		return errors.New(public.UserServiceName, exception.Error(), exception.Code())
	}
	return nil
}

//SavePassword : reset password
func (h *UserServiceHandler) SavePassword(ctx context.Context, in *user.UserDto, out *user.UserDto) error {
	str := fmt.Sprintf("%x", md5.Sum([]byte(in.Password)))
	in.Password = str
	password, err := userDao.SavePassword(ctx, in)
	if err.Code() != int32(public.OK) {
		return errors.New(public.UserServiceName, err.Error(), err.Code())
	}
	out.Id = in.Id
	out.LoginName = in.LoginName
	out.Name = in.Name
	out.Password = password
	return nil
}

func (h *UserServiceHandler) Login(ctx context.Context, in *user.UserDto, out *dto.LoginUserDto) error {
	//better than in.Password = fmt.Sprintf("%x", sum)
	//in.Password = *(*string)(unsafe.Pointer(&sum))
	str := fmt.Sprintf("%x", md5.Sum([]byte(in.Password)))
	in.Password = str
	loginUserDto, err := userDao.Login(ctx, in)
	if err.Code() != int32(public.OK) {
		return errors.New(public.UserServiceName, err.Error(), err.Code())
	}
	out.Id = loginUserDto.Id
	out.Name = loginUserDto.Name
	out.LoginName = loginUserDto.LoginName
	out.Token = loginUserDto.Token
	out.Requests = loginUserDto.Requests
	out.Resources = loginUserDto.Resources
	// 存储token
	jsonString, _ := util.ToJsonString(loginUserDto)
	redis.RedisClient.Set(loginUserDto.Token, jsonString, time.Second*3600)
	return nil
}

func (h *UserServiceHandler) Logout(ctx context.Context, in *dto.String, out *dto.String) error {
	redis.RedisClient.Del(in.Str)
	log.Println("从redis中删除token: ", in.Str)
	return nil
}

//---------- 权限管理 -------------

//LoadTree : 加载权限树
func (r *UserServiceHandler) LoadTree(ctx context.Context, in *dto.String, out *dto.ResourceDtoList) error {
	resourceDtos, exception := resourceDao.LoadTree(ctx)
	if exception.Code() != int32(public.OK) {
		return errors.New(public.UserServiceName, exception.Error(), exception.Code())
	}
	out.Rows = resourceDtos
	return nil
}

// SaveJson : 保存权限树
func (r *UserServiceHandler) SaveJson(ctx context.Context, in *dto.String, out *dto.String) error {
	exception := resourceDao.SaveJson(ctx, in.Str)
	if exception.Code() != int32(public.OK) {
		return errors.New(public.UserServiceName, exception.Error(), exception.Code())
	}
	return nil
}

// DeleteResource: 删除权限
func (r *UserServiceHandler) DeleteResource(ctx context.Context, in *dto.String, out *dto.String) error {
	exception := resourceDao.Delete(ctx, in.Str)
	if exception.Code() != int32(public.OK) {
		return errors.New(public.UserServiceName, exception.Error(), exception.Code())
	}
	return nil
}

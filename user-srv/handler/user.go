package handler

import (
	"context"
	"course/public"
	"course/user-srv/dao"
	"course/user-srv/proto/user"
	"crypto/md5"
	"fmt"
	"github.com/micro/go-micro/v2/errors"
)

/*
 when service handler return not nil, the default err code will be 500.
 it leads to the gateway response internal server error.
 so when service handler throws error, we should generate a *errors.Error(go-micro) with a code not 500,
 so that we can let the gateway response in our way.
*/

var userDao = &dao.UserDao{}

type UserHandler struct {
}

func (u *UserHandler) List(ctx context.Context, in *user.PageDto, out *user.PageDto) error {
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

func (h *UserHandler) Save(ctx context.Context, in *user.User, out *user.User) error {
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

func (h *UserHandler) Delete(ctx context.Context, in *user.User, out *user.User) error {
	exception := userDao.Delete(ctx, in)
	if exception.Code() != int32(public.OK) {
		return errors.New(public.UserServiceName, exception.Error(), exception.Code())
	}
	return nil
}

//SavePassword : reset password
func (h *UserHandler) SavePassword(ctx context.Context, in *user.User, out *user.User) error {
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

func (h *UserHandler) Login(ctx context.Context, in *user.User, out *user.LoginUserDto) error {
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
	return nil
}

func (h *UserHandler) Logout(ctx context.Context, in *user.LoginUserDto, out *user.User) error {
	return nil
}

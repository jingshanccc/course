package handler

import (
	"context"
	"course/config"
	"course/middleware/redis"
	"course/proto/basic"
	"course/public"
	"course/public/util"
	"course/user-srv/proto/dto"
	"github.com/micro/go-micro/v2/errors"
	"log"
)

func (u *UserServiceHandler) List(ctx context.Context, in *dto.PageDto, out *dto.PageDto) error {
	users, err := userDao.List(ctx, in)
	if err.Code() != int32(public.OK) {
		return errors.New(config.UserServiceName, err.Error(), err.Code())
	}
	_ = util.CopyProperties(out, in)
	out.Users = users

	return nil
}

func (u *UserServiceHandler) Save(ctx context.Context, in *dto.UserDto, out *dto.UserDto) error {
	outUser, err := userDao.Save(ctx, in)
	if err.Code() != int32(public.OK) {
		return errors.New(config.UserServiceName, err.Error(), err.Code())
	}
	_ = util.CopyProperties(out, outUser)
	return nil
}

func (u *UserServiceHandler) Delete(ctx context.Context, in *basic.String, out *basic.String) error {
	exception := userDao.Delete(ctx, in.Str)
	if exception.Code() != int32(public.OK) {
		return errors.New(config.UserServiceName, exception.Error(), exception.Code())
	}
	return nil
}

//UserInfo: 获取用户信息
func (u *UserServiceHandler) UserInfo(ctx context.Context, in *basic.String, out *dto.UserDto) error {
	userDto, exception := userDao.SelectById(ctx, in.Str)

	if userDto == nil || exception.Code() != int32(public.OK) {
		return errors.New(config.UserServiceName, exception.Error(), exception.Code())
	}
	_ = util.CopyProperties(out, userDto)
	return nil
}

//SavePassword : reset password
func (u *UserServiceHandler) SavePassword(ctx context.Context, in *dto.UpdatePass, out *basic.String) error {
	err := userDao.SavePassword(ctx, in)
	if err.Code() != int32(public.OK) {
		return errors.New(config.UserServiceName, err.Error(), err.Code())
	}
	return nil
}

func (u *UserServiceHandler) Login(ctx context.Context, in *dto.UserDto, out *dto.LoginUserDto) error {
	loginUserDto, err := userDao.Login(ctx, in)
	if err.Code() != int32(public.OK) {
		return errors.New(config.UserServiceName, err.Error(), err.Code())
	}
	_ = util.CopyProperties(out, loginUserDto)
	return nil
}

func (u *UserServiceHandler) Logout(ctx context.Context, in *basic.String, out *basic.String) error {
	redis.RedisClient.Del(ctx, in.Str)
	log.Println("从redis中删除token: ", in.Str)
	return nil
}

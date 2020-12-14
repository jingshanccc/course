package handler

import (
	"context"
	"course/config"
	"course/middleware/redis"
	"course/proto/basic"
	"course/public"
	"course/public/util"
	"course/user-srv/proto/dto"
	"course/user-srv/proto/user"
	"crypto/md5"
	"fmt"
	"github.com/micro/go-micro/v2/errors"
	"log"
)

func (u *UserServiceHandler) List(ctx context.Context, in *user.PageDto, out *user.PageDto) error {
	users, err := userDao.List(ctx, in)
	if err.Code() != int32(public.OK) {
		return errors.New(config.UserServiceName, err.Error(), err.Code())
	}
	_ = util.CopyProperties(out, in)
	out.Users = users

	return nil
}

func (u *UserServiceHandler) Save(ctx context.Context, in *user.UserDto, out *user.UserDto) error {
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
func (u *UserServiceHandler) UserInfo(ctx context.Context, in *basic.String, out *user.UserDto) error {
	userDto, exception := userDao.SelectById(ctx, in.Str)

	if userDto == nil || exception.Code() != int32(public.OK) {
		return errors.New(config.UserServiceName, exception.Error(), exception.Code())
	}
	_ = util.CopyProperties(out, userDto)
	return nil
}

//SavePassword : reset password
func (u *UserServiceHandler) SavePassword(ctx context.Context, in *user.UserDto, out *user.UserDto) error {
	str := fmt.Sprintf("%x", md5.Sum([]byte(in.Password)))
	in.Password = str
	password, err := userDao.SavePassword(ctx, in)
	if err.Code() != int32(public.OK) {
		return errors.New(config.UserServiceName, err.Error(), err.Code())
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
		return errors.New(config.UserServiceName, err.Error(), err.Code())
	}
	_ = util.CopyProperties(out, loginUserDto)
	// 存储token
	//jsonString, _ := util.ToJsonString(loginUserDto)
	//redis.RedisClient.Set(ctx, loginUserDto.Token, jsonString, time.Second*3600)
	return nil
}

func (u *UserServiceHandler) Logout(ctx context.Context, in *basic.String, out *basic.String) error {
	redis.RedisClient.Del(ctx, in.Str)
	log.Println("从redis中删除token: ", in.Str)
	return nil
}

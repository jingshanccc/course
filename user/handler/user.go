package handler

import (
	"context"
	"encoding/json"
	"gitee.com/jingshanccc/course/public"
	"gitee.com/jingshanccc/course/public/config"
	"gitee.com/jingshanccc/course/public/middleware/redis"
	"gitee.com/jingshanccc/course/public/proto/basic"
	"gitee.com/jingshanccc/course/public/util"
	"gitee.com/jingshanccc/course/user/dao"
	"gitee.com/jingshanccc/course/user/proto/dto"
	"github.com/micro/go-micro/v2/errors"
	"log"
	"time"
)

func (u *UserServiceHandler) List(ctx context.Context, in *dto.PageDto, out *dto.PageDto) error {
	total, users, err := userDao.List(ctx, in)
	if err != nil {
		return errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["user"].Name, err.Error(), err.Code())
	}
	_ = util.CopyProperties(out, in)
	out.Rows = users
	out.Total = total
	return nil
}

//SaveUserInfo: 更新当前登录用户信息
func (u *UserServiceHandler) SaveUserInfo(ctx context.Context, in *dto.UserDto, out *dto.UserDto) error {
	var usr dao.User
	_ = util.CopyProperties(&usr, in)
	usr.UpdateTime = time.Now()
	_, exception := userDao.Update(ctx, &usr)
	if exception != nil {
		return errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["user"].Name, exception.Error(), exception.Code())
	}
	_ = util.CopyProperties(out, in)
	redis.RedisClient.Del(ctx, config.Conf.Services["user"].Others["userInfoKey"].(string)+in.Id)
	return nil
}

//Save: 新增/更新用户信息
func (u *UserServiceHandler) Save(ctx context.Context, in *dto.UserDto, out *dto.UserDto) error {
	var usr dao.User
	_ = util.CopyProperties(&usr, in)
	usr.UpdateTime = time.Now()
	if usr.Id != "" {
		_, exception := userDao.Update(ctx, &usr)
		if exception != nil {
			return errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["user"].Name, exception.Error(), exception.Code())
		}
		exception = roleUserDao.DeleteByUserId(ctx, in.Id)
		if exception != nil {
			return errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["user"].Name, exception.Error(), exception.Code())
		}
		usr.Id = in.Id
	} else {
		usr.Id = util.GetShortUuid()
		_, exception := userDao.Create(ctx, &usr)
		if exception != nil {
			return errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["user"].Name, exception.Error(), exception.Code())
		}
	}
	// 保存用户角色
	var roles []string
	_ = json.Unmarshal([]byte(in.Roles), &roles)
	rts := make([]dao.RoleUser, len(roles))
	for i, roleId := range roles {
		rts[i] = dao.RoleUser{
			Id:     util.GetShortUuid(),
			RoleId: roleId,
			UserId: usr.Id,
		}
	}
	exception := roleUserDao.BatchInsert(ctx, rts)
	if exception != nil {
		return errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["user"].Name, exception.Error(), exception.Code())
	}
	_ = util.CopyProperties(out, in)
	return nil
}

func (u *UserServiceHandler) Delete(ctx context.Context, in *basic.StringList, out *basic.String) error {
	exception := userDao.Delete(ctx, in.Rows)
	if exception != nil {
		return errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["user"].Name, exception.Error(), exception.Code())
	}
	return nil
}

//UserInfo: 获取用户信息
func (u *UserServiceHandler) UserInfo(ctx context.Context, in *basic.String, out *dto.UserDto) error {
	redis.RedisClient.Del(ctx, config.Conf.Services["user"].Others["userInfoKey"].(string)+in.Str)
	var exception *public.BusinessException
	var err error
	defer func() {
		if exception != nil {
			err = errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["user"].Name, exception.Error(), exception.Code())
		} else {
			info, e := json.Marshal(out)
			if e == nil {
				redis.RedisClient.Set(ctx, config.Conf.Services["user"].Others["userInfoKey"].(string)+in.Str, info, time.Duration(config.Conf.Services["user"].Others["tokenExpire"].(int))*time.Hour)
			}
		}
	}()
	usr := userDao.SelectById(ctx, in.Str)
	_ = util.CopyProperties(out, usr)
	if out.Id == "" {
		exception := public.NewBusinessException(public.USER_NOT_EXIST)
		return errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["user"].Name, exception.Error(), exception.Code())
	}
	roles, exception := roleDao.SelectByUserId(out.Id)
	bytes, _ := json.Marshal(roles)
	out.Roles = string(bytes)
	if out.IsAdmin {
		out.Permissions = []string{"admin"}
	} else {
		out.Permissions, exception = resourceDao.SelectPermissionByUserId(ctx, in.Str)
	}
	return err
}

//SavePassword : reset password
func (u *UserServiceHandler) SavePassword(ctx context.Context, in *dto.UpdatePass, out *basic.String) error {
	err := userDao.SavePassword(ctx, in)
	if err != nil {
		return errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["user"].Name, err.Error(), err.Code())
	}
	return nil
}

func (u *UserServiceHandler) Login(ctx context.Context, in *dto.LoginUserDto, out *dto.LoginUserDto) error {
	loginUserDto, err := userDao.Login(ctx, in)
	if err != nil {
		return errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["user"].Name, err.Error(), err.Code())
	}
	_ = util.CopyProperties(out, loginUserDto)
	return nil
}

func (u *UserServiceHandler) Logout(ctx context.Context, in *basic.String, out *basic.String) error {
	redis.RedisClient.Del(ctx, in.Str)
	log.Println("从redis中删除token: ", in.Str)
	return nil
}

//SaveEmail: 用户修改邮箱
func (u *UserServiceHandler) SaveEmail(ctx context.Context, in *dto.UpdateEmail, out *basic.String) error {
	err := userDao.UpdateEmail(ctx, in)
	if err != nil {
		return errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["user"].Name, err.Error(), err.Code())
	}
	redis.RedisClient.Del(ctx, config.Conf.Services["user"].Others["userInfoKey"].(string)+in.UserId)
	return nil
}

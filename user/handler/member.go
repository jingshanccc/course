package handler

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"gitee.com/jingshanccc/course/public"
	"gitee.com/jingshanccc/course/public/config"
	"gitee.com/jingshanccc/course/public/middleware/redis"
	"gitee.com/jingshanccc/course/public/proto/basic"
	"gitee.com/jingshanccc/course/public/util"
	"gitee.com/jingshanccc/course/user/dao"
	"gitee.com/jingshanccc/course/user/proto/dto"
	"github.com/micro/go-micro/v2/errors"
	"strings"
	"time"
)

func (u *UserServiceHandler) SendEmailCode(ctx context.Context, req *basic.String, rsp *basic.String) (err error) {
	var (
		existEmails []string
		exception   *public.BusinessException
	)
	defer func() {
		if exception != nil {
			err = errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["user"].Name, exception.Error(), exception.Code())
		}
	}()
	// 校验邮箱是否唯一
	public.DB.Table("member").Select("email").Find(&existEmails)
	if len(existEmails) > 0 {
		if index := util.Contains(existEmails, req.Str); index != -1 {
			exception = public.BadRequestException("当前邮箱已绑定用户，请更换邮箱注册或直接登陆")
			return
		}
	}
	conf := config.Conf.Services["user"]
	exception = public.SendEmailCode(time.Duration(conf.Others["emailCodeExpire"].(int))*time.Minute, req.Str, conf.Others["emailRegisterKey"].(string)+req.Str, strings.Replace(conf.Others["emailTemplatePath"].(string), "user", "public", 1)+"/email.html")
	return
}
func (u *UserServiceHandler) MemberRegister(ctx context.Context, req *dto.MemberRegisterDto, rsp *basic.String) (err error) {
	var (
		existNames []string
		exception  *public.BusinessException
	)
	defer func() {
		if exception != nil {
			err = errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["user"].Name, exception.Error(), exception.Code())
		}
	}()
	// 校验用户名是否唯一
	public.DB.Table("member").Select("login_name").Find(&existNames)
	if len(existNames) > 0 {
		if index := util.Contains(existNames, req.LoginName); index != -1 {
			exception = public.BadRequestException("用户名已存在，请更换一个用户名")
			return
		}
	}
	// 校验验证码
	code, _ := redis.RedisClient.Get(ctx, config.Conf.Services["user"].Others["emailRegisterKey"].(string)+req.Email).Result()
	if code == "" {
		exception = public.NewBusinessException(public.VERIFY_CODE_EXPIRED)
		return
	}
	if code != req.Code {
		exception = public.NewBusinessException(public.VERIFY_CODE_ERROR)
		return
	}
	// 密码校验
	bytes, e := base64.StdEncoding.DecodeString(req.Password)
	req.Password, e = util.RsaDecrypt(bytes)
	if e != nil {
		exception = public.NewBusinessException(public.VALID_PARM_ERROR)
		return
	}
	req.Password = fmt.Sprintf("%x", md5.Sum([]byte(req.Password)))
	// 完成注册
	exception = memberDao.Create(&dao.Member{
		Id:           util.GetShortUuid(),
		Email:        req.Email,
		LoginName:    req.LoginName,
		Name:         req.Name,
		Password:     req.Password,
		RegisterTime: time.Now(),
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
	})
	return
}
func (u *UserServiceHandler) MemberLogin(ctx context.Context, req *dto.LoginUserDto, rsp *basic.String) (err error) {
	var (
		member    []*dao.Member
		exception *public.BusinessException
	)
	defer func() {
		if exception != nil {
			err = errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["user"].Name, exception.Error(), exception.Code())
		}
	}()
	exception = memberDao.SelectByProperty(&member, "login_name", req.LoginName)
	if len(member) == 0 {
		exception = public.NewBusinessException(public.USER_LOGIN_NAME_EXIST)
	} else {
		bytes, _ := base64.StdEncoding.DecodeString(req.Password)
		req.Password, _ = util.RsaDecrypt(bytes)
		req.Password = fmt.Sprintf("%x", md5.Sum([]byte(req.Password)))
		if member[0].Password == req.Password {
			rsp.Str = member[0].Id
		} else {
			exception = public.NewBusinessException(public.ERROR_PASSWORD)
		}
	}
	return
}

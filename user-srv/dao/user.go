package dao

import (
	"context"
	"course/config"
	"course/middleware/redis"
	"course/public"
	"course/public/util"
	"course/user-srv/proto/dto"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"log"
)

type UserDao struct {
}
type User struct {
	Id         string
	Name       string
	LoginName  string
	Password   string
	Gender     string
	Phone      string
	Email      string
	AvatarName string
	AvatarPath string
}

func (User) TableName() string {
	return "user"
}

//List : get user page
func (u *UserDao) List(ctx context.Context, in *dto.PageDto) ([]*dto.UserDto, public.BusinessException) {
	orderby := "desc"
	if in.Asc {
		orderby = "asc"
	}
	var res []*dto.UserDto
	err := public.DB.Model(&User{}).Order(in.SortBy + " " + orderby).Limit(int(in.PageSize)).Offset(int((in.PageNum - 1) * in.PageSize)).Find(&res).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return res, public.NewBusinessException(public.OK)
}

//Login : login
func (u *UserDao) Login(ctx context.Context, user *dto.UserDto) (*dto.LoginUserDto, public.BusinessException) {
	usr := u.SelectByLoginName(ctx, user.LoginName)
	if usr == nil {
		err := public.NewBusinessException(public.USER_NOT_EXIST)
		log.Println(err.Error() + user.LoginName)
		return nil, err
	} else {
		bytes, _ := base64.StdEncoding.DecodeString(user.Password)
		user.Password, _ = util.RsaDecrypt(bytes)
		user.Password = fmt.Sprintf("%x", md5.Sum([]byte(user.Password)))
		if usr.Password == user.Password {
			res := &dto.LoginUserDto{
				Id:        usr.Id,
				LoginName: usr.LoginName,
				Name:      usr.Name,
			}
			return res, public.NoException("")
		} else {
			err := public.NewBusinessException(public.ERROR_PASSWORD)
			log.Println(err.Error() + user.LoginName)
			return nil, err
		}
	}
}

//SavePassword : reset password
func (u *UserDao) SavePassword(ctx context.Context, updatePass *dto.UpdatePass) public.BusinessException {
	byId := u.SelectById(ctx, updatePass.UserId)
	if byId.Id == "" {
		return public.NewBusinessException(public.USER_NOT_EXIST)
	}
	oldBytes, _ := base64.StdEncoding.DecodeString(updatePass.OldPass)
	newBytes, _ := base64.StdEncoding.DecodeString(updatePass.NewPass)
	oldP, err := util.RsaDecrypt(oldBytes)
	newP, err := util.RsaDecrypt(newBytes)
	if err != nil {
		return public.NewBusinessException(public.VALID_PARM_ERROR)
	}
	updatePass.OldPass = fmt.Sprintf("%x", md5.Sum([]byte(oldP)))
	updatePass.NewPass = fmt.Sprintf("%x", md5.Sum([]byte(newP)))
	if byId.Password != updatePass.OldPass {
		return public.NewBusinessException(public.ERROR_PASSWORD)
	}
	if byId.Password == updatePass.NewPass {
		return public.NewBusinessException(public.SAME_PASSWORD)
	}
	err = public.DB.Model(&User{}).Where("id=?", updatePass.UserId).Update("password", updatePass.NewPass).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return public.NoException("")
}

//Save : update when dto.id exists, insert otherwise
func (u *UserDao) Save(ctx context.Context, userDto *dto.UserDto) (*dto.UserDto, public.BusinessException) {
	var usr User
	_ = util.CopyProperties(&usr, userDto)
	if userDto.Id != "" { //update
		usr.Id = ""
		usr.Password = ""
		err := public.DB.Model(&User{Id: userDto.Id}).Updates(usr).Error
		if err != nil {
			return &dto.UserDto{}, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
		}
	} else { //insert
		// loginName phone email exists
		us := u.SelectByLoginName(ctx, userDto.LoginName)
		if us != nil {
			return nil, public.NewBusinessException(public.USER_LOGIN_NAME_EXIST)
		}
		us = u.SelectByEmail(ctx, userDto.Email)
		if us != nil {
			return nil, public.NewBusinessException(public.USER_EMAIL_EXIST)
		}
		us = u.SelectByPhone(ctx, userDto.Phone)
		if us != nil {
			return nil, public.NewBusinessException(public.USER_PHONE_EXIST)
		}
		usr.Id = util.GetShortUuid()
		usr.Password = fmt.Sprintf("%x", md5.Sum([]byte(userDto.Password)))
		err1 := public.DB.Create(usr).Error
		if err1 != nil {
			return &dto.UserDto{}, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
		}
	}
	return userDto, public.NoException("")
}

// Delete 删除用户
func (u *UserDao) Delete(ctx context.Context, id string) public.BusinessException {
	err := public.DB.Delete(&User{Id: id}).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return public.NoException("")
}

//UpdateEmail: 更新邮箱
func (u *UserDao) UpdateEmail(ctx context.Context, in *dto.UpdateEmail) public.BusinessException {
	//校验验证码
	code, _ := redis.RedisClient.Get(ctx, config.EmailResetEmailCode+in.Email).Result()
	if code == "" {
		return public.NewBusinessException(public.VERIFY_CODE_EXPIRED)
	}
	if code != in.Code {
		return public.NewBusinessException(public.VERIFY_CODE_ERROR)
	}
	//校验密码
	newBytes, _ := base64.StdEncoding.DecodeString(in.Pass)
	pas, err := util.RsaDecrypt(newBytes)
	if err != nil {
		return public.NewBusinessException(public.VALID_PARM_ERROR)
	}
	in.Pass = fmt.Sprintf("%x", md5.Sum([]byte(pas)))
	byId := u.SelectById(ctx, in.UserId)
	if byId.Id == "" {
		return public.NewBusinessException(public.USER_NOT_EXIST)
	}
	if byId.Password != in.Pass {
		return public.NewBusinessException(public.ERROR_PASSWORD)
	}
	err = public.DB.Model(&User{}).Where("id=?", in.UserId).Update("email", in.Email).Error
	if err != nil {
		return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return public.NoException("")
}

//SelectByLoginName : get user by login name
func (u *UserDao) SelectByLoginName(ctx context.Context, loginName string) *dto.UserDto {
	us := &dto.UserDto{}
	public.DB.Model(&User{}).Where("login_name = ?", loginName).First(&us)
	return us
}

//SelectById : get user by id
func (u *UserDao) SelectById(ctx context.Context, id string) *dto.UserDto {
	us := &dto.UserDto{}
	public.DB.Model(&User{}).Where("id = ?", id).First(&us)
	return us
}

func (u *UserDao) SelectByEmail(ctx context.Context, email string) *dto.UserDto {
	us := &dto.UserDto{}
	public.DB.Model(&User{}).Where("email = ?", email).First(&us)
	return us
}

func (u *UserDao) SelectByPhone(ctx context.Context, phone string) *dto.UserDto {
	us := &dto.UserDto{}
	public.DB.Model(&User{}).Where("phone = ?", phone).First(&us)
	return us
}

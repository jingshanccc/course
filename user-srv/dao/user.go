package dao

import (
	"context"
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
	Id        string
	Name      string
	LoginName string
	Password  string
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

//SelectByLoginName : get user by login name
func (u *UserDao) SelectByLoginName(ctx context.Context, loginName string) (*dto.UserDto, public.BusinessException) {
	us := &dto.UserDto{}
	err := public.DB.Model(&User{}).Where("login_name = ?", loginName).First(&us).Error
	if err != nil {
		return nil, public.NoException("")
	}
	return us, public.NoException("")
}

//SelectById : get user by id
func (u *UserDao) SelectById(ctx context.Context, id string) (*dto.UserDto, public.BusinessException) {
	us := &dto.UserDto{}
	err := public.DB.Model(&User{}).Where("id = ?", id).First(&us).Error
	if err != nil || us == nil {
		return nil, public.NewBusinessException(public.LOGIN_USER_ERROR)
	}
	return us, public.NoException("")
}

//Login : login
func (u *UserDao) Login(ctx context.Context, user *dto.UserDto) (*dto.LoginUserDto, public.BusinessException) {
	usr, err := u.SelectByLoginName(ctx, user.LoginName)
	if usr == nil {
		err = public.NewBusinessException(public.LOGIN_USER_ERROR)
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
			err = public.NewBusinessException(public.ERROR_PASSWORD)
			log.Println(err.Error() + user.LoginName)
			return nil, err
		}
	}
}

//SavePassword : reset password
func (u *UserDao) SavePassword(ctx context.Context, updatePass *dto.UpdatePass) public.BusinessException {
	byId, exception := u.SelectById(ctx, updatePass.UserId)
	if exception.Code() != int32(public.OK) {
		return exception
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
	if byId.Password != updatePass.NewPass {
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
	if userDto.Id != "" { //update
		err := public.DB.Model(&User{Id: userDto.Id}).Updates(&User{Name: userDto.Name, LoginName: userDto.LoginName}).Error
		if err != nil {
			return &dto.UserDto{}, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
		}
	} else { //insert
		//if login_name exists, throw error
		us, _ := u.SelectByLoginName(ctx, userDto.LoginName)
		if us != nil {
			return &dto.UserDto{}, public.NewBusinessException(public.USER_LOGIN_NAME_EXIST)
		}
		userDto.Id = util.GetShortUuid()
		userDto.Password = fmt.Sprintf("%x", md5.Sum([]byte(userDto.Password)))
		err1 := public.DB.Create(&User{
			Id:        userDto.Id,
			Name:      userDto.Name,
			LoginName: userDto.LoginName,
			Password:  userDto.Password,
		}).Error
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

package dao

import (
	"context"
	"course/public"
	"course/public/util"
	"course/user-srv/proto/dto"
	"course/user-srv/proto/user"
	"crypto/md5"
	"encoding/json"
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
func (u *UserDao) List(ctx context.Context, dto *user.PageDto) ([]*user.UserDto, public.BusinessException) {
	orderby := "desc"
	if dto.Asc {
		orderby = "asc"
	}
	var res []*user.UserDto
	err := public.DB.Model(&User{}).Order(dto.SortBy + " " + orderby).Limit(int(dto.PageSize)).Offset(int((dto.PageNum - 1) * dto.PageSize)).Find(&res).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return res, public.NewBusinessException(public.OK)
}

//SelectByLoginName : get user by login name
func (u *UserDao) SelectByLoginName(ctx context.Context, loginName string) (*user.UserDto, public.BusinessException) {
	us := &user.UserDto{}
	err := public.DB.Model(&User{}).Where("login_name = ?", loginName).First(&us).Error
	if err != nil {
		return nil, public.NoException("")
	}
	return us, public.NoException("")
}

//Login : login
func (u *UserDao) Login(ctx context.Context, user *user.UserDto) (*dto.LoginUserDto, public.BusinessException) {
	usr, err := u.SelectByLoginName(ctx, user.LoginName)
	if usr == nil {
		err = public.NewBusinessException(public.MEMBER_NOT_EXIST)
		log.Println(err.Error() + user.LoginName)
		return nil, err
	} else {
		if usr.Password == user.Password {
			res := &dto.LoginUserDto{
				Id:        usr.Id,
				LoginName: usr.LoginName,
				Name:      usr.Name,
			}
			res.Token = util.GetUuid()
			exception := setAuth(ctx, res)
			return res, exception
		} else {
			err = public.NewBusinessException(public.LOGIN_USER_ERROR)
			log.Println(err.Error() + user.LoginName)
			return nil, err
		}
	}
}

//setAuth : set user's resources (access control)
func setAuth(ctx context.Context, loginUser *dto.LoginUserDto) public.BusinessException {
	resources, exception := (&ResourceDao{}).FindUserResources(ctx, loginUser.Id)
	if exception.Code() != int32(public.OK) {
		return exception
	}
	requestSet := public.NewHashSet()
	if len(resources) > 0 {
		for _, resource := range resources {
			var requests []string
			request := resource.Request
			json.Unmarshal([]byte(request), &requests)
			if len(requests) > 0 {
				for _, v := range requests {
					requestSet.Add(v)
				}
			}
		}
	}
	var reqs []string
	requestJson, _ := requestSet.ToJSON()
	json.Unmarshal(requestJson, &reqs)
	loginUser.Resources = resources
	loginUser.Requests = reqs
	return exception
}

//SavePassword : reset password
func (u *UserDao) SavePassword(ctx context.Context, dto *user.UserDto) (string, public.BusinessException) {
	err := public.DB.Model(&User{LoginName: dto.LoginName}).Update("password", dto.Password).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return "", public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return dto.Password, public.NoException("")
}

//Save : update when dto.id exists, insert otherwise
func (u *UserDao) Save(ctx context.Context, dto *user.UserDto) (*user.UserDto, public.BusinessException) {
	if dto.Id != "" { //update
		err := public.DB.Model(&User{Id: dto.Id}).Updates(&User{Name: dto.Name, LoginName: dto.LoginName}).Error
		if err != nil {
			return &user.UserDto{}, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
		}
	} else { //insert
		//if login_name exists, throw error
		us, _ := u.SelectByLoginName(ctx, dto.LoginName)
		if us != nil {
			return &user.UserDto{}, public.NewBusinessException(public.USER_LOGIN_NAME_EXIST)
		}
		dto.Id = util.GetShortUuid()
		dto.Password = fmt.Sprintf("%x", md5.Sum([]byte(dto.Password)))
		err1 := public.DB.Create(&User{
			Id:        dto.Id,
			Name:      dto.Name,
			LoginName: dto.LoginName,
			Password:  dto.Password,
		}).Error
		if err1 != nil {
			return &user.UserDto{}, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
		}
	}
	return dto, public.NoException("")
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

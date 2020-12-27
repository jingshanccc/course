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
	"strings"
	"time"
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
	IsAdmin    bool
	CreateBy   string
	CreateTime time.Time
	UpdateBy   string
	UpdateTime time.Time
}

func (User) TableName() string {
	return "user"
}

//List : get user page
func (u *UserDao) List(ctx context.Context, in *dto.PageDto) (int64, []*dto.UserDto, *public.BusinessException) {
	conditions := ""
	join := "join role_user ru on ru.user_id = u.id join role r on r.id = ru.role_id "
	// blurry 模糊搜索字段 sort 排序数组 ["id,desc","name,asc"]
	and := true
	if in.CreateTime != nil {
		and = false
		conditions += "where u.create_time between '" + in.CreateTime[0] + "' and '" + in.CreateTime[1] + "'"
	}
	if in.Blurry != "" {
		if and {
			conditions += "where "
		} else {
			conditions += " and "
		}
		conditions += "(u.name like '%" + in.Blurry + "%' or "
		conditions += "u.login_name like '%" + in.Blurry + "%' or "
		conditions += "u.email like '%" + in.Blurry + "%')"
	}
	var count int64
	err := public.DB.Raw("select count(1) from user u " + conditions).Find(&count).Error
	// 当使用 group_concat 若无 group by 只能选出一条记录
	conditions += " group by u.id"
	if in.Sort != nil {
		conditions += " order by "
		for i, s := range in.Sort {
			sort := strings.Split(s, ",")
			conditions += "u." + sort[0] + " " + sort[1]
			if i != len(in.Sort)-1 {
				conditions += ","
			}
		}
	}

	conditions += " limit ?,?"
	var res []*dto.UserDto
	err = public.DB.Raw("select u.*, concat('[', group_concat(json_object('id',r.id, 'name',r.name, 'desc',r.`desc`)),']') as 'roles' from user u "+join+conditions, (in.Page-1)*in.Size, in.Size).Find(&res).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return 0, nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return count, res, nil
}

//Login : login
func (u *UserDao) Login(ctx context.Context, user *dto.LoginUserDto) (*dto.LoginUserDto, *public.BusinessException) {
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
			return res, nil
		} else {
			err := public.NewBusinessException(public.ERROR_PASSWORD)
			log.Println(err.Error() + user.LoginName)
			return nil, err
		}
	}
}

//SavePassword : reset password
func (u *UserDao) SavePassword(ctx context.Context, updatePass *dto.UpdatePass) *public.BusinessException {
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
	return nil
}

//Create: 创建新用户
func (u *UserDao) Create(ctx context.Context, userDto *User) (*User, *public.BusinessException) {
	//insert
	// loginName phone email exists
	us := u.SelectByLoginName(ctx, userDto.LoginName)
	if us.Id != "" {
		return nil, public.NewBusinessException(public.USER_LOGIN_NAME_EXIST)
	}
	usd := u.SelectByEmail(ctx, userDto.Email)
	if usd.Id != "" {
		return nil, public.NewBusinessException(public.USER_EMAIL_EXIST)
	}
	usd = u.SelectByPhone(ctx, userDto.Phone)
	if usd.Id != "" {
		return nil, public.NewBusinessException(public.USER_PHONE_EXIST)
	}
	userDto.Password = fmt.Sprintf("%x", md5.Sum([]byte(config.DefaultPassword)))
	userDto.CreateBy = userDto.UpdateBy
	userDto.CreateTime = userDto.UpdateTime
	err := public.DB.Create(userDto).Error
	if err != nil {
		return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return userDto, nil
}

//Update : 更新用户
func (u *UserDao) Update(ctx context.Context, userDto *User) (*User, *public.BusinessException) {
	// loginName phone email exists
	us := u.SelectByLoginName(ctx, userDto.LoginName)
	if us.Id != "" && us.Id != userDto.Id {
		return nil, public.NewBusinessException(public.USER_LOGIN_NAME_EXIST)
	}
	usd := u.SelectByEmail(ctx, userDto.Email)
	if usd.Id != "" && us.Id != userDto.Id {
		return nil, public.NewBusinessException(public.USER_EMAIL_EXIST)
	}
	usd = u.SelectByPhone(ctx, userDto.Phone)
	if usd.Id != "" && us.Id != userDto.Id {
		return nil, public.NewBusinessException(public.USER_PHONE_EXIST)
	}

	userDto.Id = ""
	userDto.Password = ""
	err := public.DB.Model(&User{Id: userDto.Id}).Updates(userDto).Error
	if err != nil {
		return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return userDto, nil
}

// Delete 删除用户
func (u *UserDao) Delete(ctx context.Context, ids []string) *public.BusinessException {
	err := public.DB.Where("id in ?", ids).Delete(&User{}).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return nil
}

//UpdateEmail: 更新邮箱
func (u *UserDao) UpdateEmail(ctx context.Context, in *dto.UpdateEmail) *public.BusinessException {
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
	return nil
}

//SelectByLoginName : get user by login name
func (u *UserDao) SelectByLoginName(ctx context.Context, loginName string) *User {
	var us User
	public.DB.Model(&User{}).Where("login_name = ?", loginName).First(&us)
	return &us
}

//SelectById : get user by id
func (u *UserDao) SelectById(ctx context.Context, id string) *User {
	var us User
	public.DB.Model(&User{}).Where("id = ?", id).First(&us)
	return &us
}

func (u *UserDao) SelectByEmail(ctx context.Context, email string) *dto.UserDto {
	var us dto.UserDto
	public.DB.Model(&User{}).Where("email = ?", email).First(&us)
	return &us
}

func (u *UserDao) SelectByPhone(ctx context.Context, phone string) *dto.UserDto {
	var us dto.UserDto
	public.DB.Model(&User{}).Where("phone = ?", phone).First(&us)
	return &us
}

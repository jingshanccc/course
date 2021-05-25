package dao

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"gitee.com/jingshanccc/course/public"
	"gitee.com/jingshanccc/course/public/config"
	"gitee.com/jingshanccc/course/public/util"
	"gitee.com/jingshanccc/course/user/proto/dto"
	"log"
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
	join := "join role_user ru on ru.user_id = x.id join role r on r.id = ru.role_id "
	// 当使用 group_concat 若无 group by 只能选出一条记录
	forCount, forPage := util.GeneratePageSql(in.CreateTime, in.Blurry, in.Sort, []string{"login_name", "name", "email"}, " group by x.id")
	var count int64
	err := public.DB.Raw("select count(1) from user x " + forCount).Find(&count).Error
	var res []*dto.UserDto
	err = public.DB.Raw("select x.*, concat('[', group_concat(json_object('id',r.id, 'name',r.name, 'desc',r.`desc`)),']') as 'roles' from user x "+join+forPage, (in.Page-1)*in.Size, in.Size).Find(&res).Error
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
	if byId.Password != updatePass.OldPass {
		return public.NewBusinessException(public.ERROR_PASSWORD)
	}
	if byId.Password == updatePass.NewPass {
		return public.NewBusinessException(public.SAME_PASSWORD)
	}
	err := public.DB.Model(&User{}).Where("id=?", updatePass.UserId).Update("password", updatePass.NewPass).Error
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
	userDto.Password = fmt.Sprintf("%x", md5.Sum([]byte(config.Conf.Services["user"].Others["defaultPassword"].(string))))
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
	if usd.Id != "" && usd.Id != userDto.Id {
		return nil, public.NewBusinessException(public.USER_EMAIL_EXIST)
	}
	usd = u.SelectByPhone(ctx, userDto.Phone)
	if usd.Id != "" && usd.Id != userDto.Id {
		return nil, public.NewBusinessException(public.USER_PHONE_EXIST)
	}

	id := userDto.Id
	userDto.Id = ""
	userDto.Password = ""
	err := public.DB.Model(&User{Id: id}).Updates(userDto).Error
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
	byId := u.SelectById(ctx, in.UserId)
	if byId.Id == "" {
		return public.NewBusinessException(public.USER_NOT_EXIST)
	}
	if byId.Password != in.Pass {
		return public.NewBusinessException(public.ERROR_PASSWORD)
	}
	err := public.DB.Model(&User{}).Where("id=?", in.UserId).Update("email", in.Email).Error
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

//CountByRoles: 查询角色是否关联了用户
func (u *UserDao) CountByRoles(ctx context.Context, roles []string) int64 {
	var count int64
	public.DB.Raw("SELECT count(1) FROM user u, role_user r WHERE u.id = r.user_id AND r.role_id in ?", roles).Find(&count)
	return count
}

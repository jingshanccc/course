package dao

import (
	"context"
	"gitee.com/jingshanccc/course/public"
	"gitee.com/jingshanccc/course/public/util"
	"gitee.com/jingshanccc/course/user/proto/dto"
	"log"
	"time"
)

type Member struct {
	Id           string    // id
	Email        string    // 邮箱
	Name         string    // 昵称
	LoginName    string    // 用户名
	Password     string    // 密码
	Sex          string    // 性别
	Photo        string    // 用户头像
	Integral     uint32    // 积分
	LearningTime uint32    // 学习时长
	RegisterTime time.Time // 注册时间
	CreateTime   time.Time // 创建时间
	UpdateTime   time.Time // 更新时间
	Creator      string    // 创建者
	Updater      string    // 更新者
}

func (Member) TableName() string {
	return "member"
}

type MemberDao struct {
}

func (m *MemberDao) Create(member *Member) *public.BusinessException {
	member.Id = util.GetShortUuid()
	member.Sex = "男"
	member.RegisterTime = time.Now()
	member.CreateTime = time.Now()
	member.UpdateTime = time.Now()
	err := public.DB.Create(&member).Error
	if err != nil {
		return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return nil
}

func (m *MemberDao) SelectByProperty(member *[]*Member, property string, val interface{}) *public.BusinessException {
	err := public.DB.Model(member).Where(property+" = ?", val).Find(&member).Error
	if err != nil {
		return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return nil
}

//SavePassword : reset password
func (m *MemberDao) SavePassword(ctx context.Context, updatePass *dto.UpdatePass) *public.BusinessException {
	var (
		members []*Member
		byId    *Member
	)
	exception := m.SelectByProperty(&members, "id", updatePass.UserId)
	if exception != nil {
		return exception
	}
	if len(members) == 0 {
		return public.NewBusinessException(public.USER_NOT_EXIST)
	}
	byId = members[0]
	if byId.Password != updatePass.OldPass {
		return public.NewBusinessException(public.ERROR_PASSWORD)
	}
	if byId.Password == updatePass.NewPass {
		return public.NewBusinessException(public.SAME_PASSWORD)
	}
	err := public.DB.Model(&Member{}).Where("id=?", updatePass.UserId).Update("password", updatePass.NewPass).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return nil
}

//UpdateEmail: 更新邮箱
func (m *MemberDao) UpdateEmail(ctx context.Context, in *dto.UpdateEmail) *public.BusinessException {
	var (
		members []*Member
		byId    *Member
	)
	exception := m.SelectByProperty(&members, "id", in.UserId)
	if exception != nil {
		return exception
	}
	if len(members) == 0 {
		return public.NewBusinessException(public.USER_NOT_EXIST)
	}
	byId = members[0]
	if byId.Id == "" {
		return public.NewBusinessException(public.USER_NOT_EXIST)
	}
	if byId.Password != in.Pass {
		return public.NewBusinessException(public.ERROR_PASSWORD)
	}
	err := public.DB.Model(&Member{}).Where("id=?", in.UserId).Update("email", in.Email).Error
	if err != nil {
		return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return nil
}

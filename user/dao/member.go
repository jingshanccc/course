package dao

import (
	"gitee.com/jingshanccc/course/public"
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

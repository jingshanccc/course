package dao

import (
	"context"
	"course/public"
)

type RoleUserDao struct {
}
type RoleUser struct {
	Id     string
	RoleId string
	UserId string
}

func (RoleUser) TableName() string {
	return "role_user"
}

//DeleteByRoleId: 删除角色关联的所有记录
func (r *RoleUserDao) DeleteByRoleId(ctx context.Context, roleId string) public.BusinessException {
	err := public.DB.Where("role_id= ?", roleId).Delete(&RoleUser{}).Error
	if err != nil {
		return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return public.NoException("")
}

//Save: 创建角色-用户关联记录
func (r *RoleUserDao) Save(ctx context.Context, rt RoleUser) public.BusinessException {
	err := public.DB.Create(&rt).Error
	if err != nil {
		return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return public.NoException("")
}

//SelectByRoleId: 查询角色关联的所有记录
func (r *RoleUserDao) SelectByRoleId(ctx context.Context, roleId string) ([]string, public.BusinessException) {
	var res []string
	err := public.DB.Model(&RoleUser{}).Select("user_id").Where("role_id = ?", roleId).Find(&res).Error
	if err != nil {
		return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return res, public.NoException("")
}

package dao

import (
	"context"
	"course/public"
)

type RoleResourceDao struct {
}
type RoleResource struct {
	Id         string
	RoleId     string
	ResourceId int32
}

func (RoleResource) TableName() string {
	return "role_resource"
}

//DeleteByRoleId: 删除角色关联的所有记录
func (r *RoleResourceDao) DeleteByRoleId(ctx context.Context, roleId string) *public.BusinessException {
	err := public.DB.Where("role_id= ?", roleId).Delete(&RoleResource{}).Error
	if err != nil {
		return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return nil
}

//Save: 创建角色-资源关联记录
func (r *RoleResourceDao) Save(ctx context.Context, rt RoleResource) *public.BusinessException {
	err := public.DB.Create(&rt).Error
	if err != nil {
		return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return nil
}

//SelectByRoleId: 查询角色关联的所有记录
func (r *RoleResourceDao) SelectByRoleId(ctx context.Context, roleId string) ([]string, *public.BusinessException) {
	var res []string
	err := public.DB.Model(&RoleResource{}).Select("resource_id").Where("role_id = ?", roleId).Find(&res).Error
	if err != nil {
		return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return res, nil
}

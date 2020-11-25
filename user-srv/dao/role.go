package dao

import (
	"context"
	"course/public"
	"course/public/util"
	"course/user-srv/proto/dto"
	"log"
)

type RoleDao struct {
}

type Role struct {
	Id   string
	Name string
	Desc string
}

func (Role) TableName() string {
	return "role"
}

var (
	roleResourceDao = &RoleResourceDao{}
	roleUserDao     = &RoleUserDao{}
)

//List: 获取角色列表
func (r *RoleDao) List(ctx context.Context, page *dto.RolePageDto) ([]*dto.RoleDto, public.BusinessException) {
	orderby := "desc"
	if page.Asc {
		orderby = "asc"
	}
	var roles []*Role
	err := public.DB.Model(&Role{}).Order(page.SortBy + " " + orderby).Limit(int(page.PageSize)).Offset(int((page.PageNum - 1) * page.PageSize)).Find(&roles).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	res := make([]*dto.RoleDto, len(roles))
	for i, role := range roles {
		a := &dto.RoleDto{}
		_ = util.CopyProperties(a, role)
		res[i] = a
	}
	return res, public.NewBusinessException(public.OK)
}

//Save : update when rr.id exists, insert otherwise
func (r *RoleDao) Save(ctx context.Context, rr *dto.RoleDto) (*dto.RoleDto, public.BusinessException) {
	if rr.Id != "" { //update
		err := public.DB.Model(&Role{Id: rr.Id}).Updates(&Role{Name: rr.Name, Desc: rr.Desc}).Error
		if err != nil {
			return &dto.RoleDto{}, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
		}
	} else { //insert
		rr.Id = util.GetShortUuid()
		err1 := public.DB.Create(&Role{
			Id:   rr.Id,
			Name: rr.Name,
			Desc: rr.Desc,
		}).Error
		if err1 != nil {
			return &dto.RoleDto{}, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
		}
	}
	return rr, public.NoException("")
}

// Delete 删除角色
func (r *RoleDao) Delete(ctx context.Context, id string) public.BusinessException {
	err := public.DB.Delete(&Role{Id: id}).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return public.NoException("")
}

// SaveRoleResource: 保存角色资源关联记录
func (r *RoleDao) SaveRoleResource(ctx context.Context, rt *dto.RoleDto) public.BusinessException {
	exception := roleResourceDao.DeleteByRoleId(ctx, rt.Id)
	for _, resourceId := range rt.ResourceIds {
		exception = roleResourceDao.Save(ctx, RoleResource{
			Id:         util.GetShortUuid(),
			RoleId:     rt.Id,
			ResourceId: resourceId,
		})
	}
	return exception
}

//ListRoleResource: 获取角色所有权限
func (r *RoleDao) ListRoleResource(ctx context.Context, roleId string) ([]string, public.BusinessException) {
	return roleResourceDao.SelectByRoleId(ctx, roleId)
}

//SaveRoleUser: 保存角色的所有用户
func (r *RoleDao) SaveRoleUser(ctx context.Context, rt *dto.RoleDto) public.BusinessException {
	exception := roleUserDao.DeleteByRoleId(ctx, rt.Id)
	for _, userId := range rt.UserIds {
		exception = roleUserDao.Save(ctx, RoleUser{
			Id:     util.GetShortUuid(),
			RoleId: rt.Id,
			UserId: userId,
		})
	}
	return exception
}

//ListRoleResource: 获取角色所有权限
func (r *RoleDao) ListRoleUser(ctx context.Context, roleId string) ([]string, public.BusinessException) {
	return roleUserDao.SelectByRoleId(ctx, roleId)
}

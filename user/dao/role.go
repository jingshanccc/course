package dao

import (
	"context"
	"gitee.com/jingshanccc/course/public"
	"gitee.com/jingshanccc/course/public/util"
	"gitee.com/jingshanccc/course/user/proto/dto"
	"log"
	"time"
)

type RoleDao struct {
}

type Role struct {
	Id         string
	Name       string
	Desc       string
	Level      int32
	CreateTime time.Time
}

func (Role) TableName() string {
	return "role"
}

var (
	roleResourceDao = &RoleResourceDao{}
	roleUserDao     = &RoleUserDao{}
)

//All: 获取所有角色
func (r *RoleDao) All(ctx context.Context) ([]*dto.RoleDto, *public.BusinessException) {
	var res []*dto.RoleDto
	err := public.DB.Model(&Role{}).Order("level").Find(&res).Error
	if err != nil {
		return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return res, nil
}

//List: 获取角色列表
func (r *RoleDao) List(ctx context.Context, in *dto.RolePageDto) (int64, []*dto.RoleDto, *public.BusinessException) {
	forCount, forPage := util.GeneratePageSql(in.CreateTime, in.Blurry, in.Sort, []string{"name", "desc"}, "")
	var count int64
	err := public.DB.Raw("select count(1) from role x " + forCount).Find(&count).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return 0, nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	var res []*dto.RoleDto
	err = public.DB.Raw("select x.* from role x "+forPage, (in.Page-1)*in.Size, in.Size).Find(&res).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return 0, nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	for _, role := range res {
		role.ResourceIds, _ = roleResourceDao.SelectByRoleId(ctx, role.Id)
	}
	return count, res, nil
}

//Save : update when rr.id exists, insert otherwise
func (r *RoleDao) Save(ctx context.Context, rr *dto.RoleDto) (*dto.RoleDto, *public.BusinessException) {
	if rr.Id != "" { //update
		err := public.DB.Model(&Role{Id: rr.Id}).Updates(&Role{Name: rr.Name, Desc: rr.Desc, Level: rr.Level}).Error
		if err != nil {
			return &dto.RoleDto{}, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
		}
	} else { //insert
		rr.Id = util.GetShortUuid()
		err := public.DB.Create(&Role{
			Id:         rr.Id,
			Name:       rr.Name,
			Desc:       rr.Desc,
			Level:      rr.Level,
			CreateTime: time.Now(),
		}).Error
		if err != nil {
			return &dto.RoleDto{}, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
		}
	}
	return rr, nil
}

// Delete 删除角色
func (r *RoleDao) Delete(ctx context.Context, ids []string) *public.BusinessException {
	err := public.DB.Where("id in ?", ids).Delete(&Role{}).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	exception := roleResourceDao.DeleteByRoleId(ctx, ids)
	if exception != nil {
		log.Println("exec sql failed, err is " + exception.Error())
		return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return nil
}

// SaveRoleResource: 保存角色资源关联记录
func (r *RoleDao) SaveRoleResource(ctx context.Context, rt *dto.RoleDto) *public.BusinessException {
	exception := roleResourceDao.DeleteByRoleId(ctx, []string{rt.Id})
	for _, resourceId := range rt.ResourceIds {
		exception = roleResourceDao.Save(ctx, RoleResource{
			RoleId:     rt.Id,
			ResourceId: resourceId,
		})
	}
	return exception
}

//ListRoleResource: 获取角色所有权限
func (r *RoleDao) ListRoleResource(ctx context.Context, roleId string) ([]int32, *public.BusinessException) {
	return roleResourceDao.SelectByRoleId(ctx, roleId)
}

//SaveRoleUser: 保存角色的所有用户
func (r *RoleDao) SaveRoleUser(ctx context.Context, rt *dto.RoleDto) *public.BusinessException {
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

//ListRoleUser: 获取角色所有用户
func (r *RoleDao) ListRoleUser(ctx context.Context, roleId string) ([]string, *public.BusinessException) {
	return roleUserDao.SelectByRoleId(ctx, roleId)
}

//SelectByUserId: 获取用户角色
func (r *RoleDao) SelectByUserId(userId string) ([]Role, *public.BusinessException) {
	var roles []Role
	err := public.DB.Raw("select r.* from role r, role_user ru where ru.role_id = r.id and ru.user_id = ? order by r.level", userId).Find(&roles).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return roles, nil
}

//SelectById: 获取传入 ID 角色
func (r *RoleDao) SelectById(ctx context.Context, id string) (*dto.RoleDto, *public.BusinessException) {
	var res dto.RoleDto
	err := public.DB.Raw("select * from role where id = ?", id).Find(&res).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	res.ResourceIds, _ = roleResourceDao.SelectByRoleId(ctx, res.Id)
	return &res, nil
}

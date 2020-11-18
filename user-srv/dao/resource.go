package dao

import (
	"context"
	"course/public"
	"course/user-srv/proto/user"
	"log"
)

type ResourceDao struct {
}
type Resource struct {
	//可能为空的字段使用指针类型
	Id      string
	Name    string
	Page    *string
	Parent  *string
	Request *string
}

// findUserResources 获取用户的权限
func (r *ResourceDao) findUserResources(ctx context.Context, userId string) ([]*user.ResourceDto, public.BusinessException) {
	stmt, err := public.DB.Prepare("select distinct r.id, r.name, r.page, r.request, r.parent from role_user ru, role_resource rr, resource r where ru.user_id = ? and ru.role_id = rr.role_id and rr.resource_id = r.id order by r.id")
	if err != nil {
		return nil, public.NewBusinessException(public.PREPARE_SQL_ERROR)
	}
	defer stmt.Close()
	rows, err := stmt.Query(userId)
	if err != nil {
		return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}

	var res []*user.ResourceDto

	for rows.Next() {
		r := &Resource{}
		err = rows.Scan(&r.Id, &r.Name, &r.Page, &r.Request, &r.Parent)
		if err != nil {
			log.Println("row scan failed, err is " + err.Error())
			return nil, public.NewBusinessException(public.ROW_SCAN_ERROR)
		}
		dto := &user.ResourceDto{
			Id:   r.Id,
			Name: r.Name,
		}
		if r.Page == nil {
			dto.Page = ""
		} else {
			dto.Page = *r.Page
		}
		if r.Parent == nil {
			dto.Parent = ""
		} else {
			dto.Parent = *r.Parent
		}
		if r.Request == nil {
			dto.Request = ""
		} else {
			dto.Request = *r.Request
		}
		res = append(res, dto)
	}
	return res, public.NoException("")
}

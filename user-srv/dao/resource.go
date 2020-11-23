package dao

import (
	"context"
	"course/public"
	"course/user-srv/proto/dto"
	"database/sql"
	"encoding/json"
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

// Delete : 删除权限
func (r *ResourceDao) Delete(ctx context.Context, id string) public.BusinessException {
	stmt, err := public.DB.PrepareContext(ctx, "delete from resource where id = ?")
	if err != nil {
		return public.NewBusinessException(public.PREPARE_SQL_ERROR)
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return public.NoException("")
}

//SaveJson : 保存权限树
func (r *ResourceDao) SaveJson(ctx context.Context, jsonStr string) public.BusinessException {
	// rollback 逻辑
	var exception public.BusinessException
	var tx *sql.Tx
	defer func() {
		if exception.Code() != int32(public.OK) {
			tx.Rollback()
		}
	}()
	// 从传入的json串中获取资源列表
	var inputList []*dto.ResourceDto
	err := json.Unmarshal([]byte(jsonStr), &inputList)
	var list []*dto.ResourceDto
	if err == nil && len(inputList) > 0 {
		for _, resourceDto := range inputList {
			resourceDto.Parent = ""
			add(&list, resourceDto)
		}
	}
	//删除数据库中的所有resource, 然后保存新的resource
	tx, err = public.DB.Begin()
	if err != nil {
		return public.NewBusinessException(public.BEGIN_TRANSACTION_ERROR)
	}
	_, err = tx.Exec("delete from resource")
	if err != nil {
		return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}

	for _, resourceDto := range list {
		_, err = tx.Exec("insert into resource (id, name, page, request, parent) VALUES (?, ?, ?, ?, ?)", resourceDto.Id, resourceDto.Name, resourceDto.Page, resourceDto.Request, resourceDto.Parent)
		if err != nil {
			return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Println("transaction commit error, need roll back")
		return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return public.NoException("")
}

//add: 将树中的resource都取出来放到list
func add(list *[]*dto.ResourceDto, resourceDto *dto.ResourceDto) {
	*list = append(*list, resourceDto)
	if resourceDto.Children != nil && len(resourceDto.Children) > 0 {
		for _, child := range resourceDto.Children {
			child.Parent = resourceDto.Id
			add(list, child)
		}
	}
}

// LoadTree : 按约定将列表转成树, ID要正序排列
func (r *ResourceDao) LoadTree(ctx context.Context) ([]*dto.ResourceDto, public.BusinessException) {
	stmt, err := public.DB.Prepare("select distinct r.id, r.name, r.page, r.request, r.parent from resource r order by r.id")
	if err != nil {
		return nil, public.NewBusinessException(public.PREPARE_SQL_ERROR)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	res, exception := scanResources(rows)
	if exception.Code() != int32(public.OK) {
		return nil, exception
	}
	for i := len(res) - 1; i >= 0; i-- {
		child := res[i]
		if child.Parent == "" {
			continue
		}
		for j := i - 1; j >= 0; j-- {
			parent := res[j]
			if child.Parent == parent.Id {
				if parent.Children == nil {
					parent.Children = []*dto.ResourceDto{}
				}
				// 此时为倒序 因为循环是从后往前的
				parent.Children = append(parent.Children, child)
				// 子节点找到父节点后，删除列表中的子节点
				for index, val := range res {
					if val.Id == child.Id {
						res = append(res[:index], res[index+1:]...)
						break
					}
				}
			}
		}
	}
	reverseChild(res)
	return res, public.NoException("")
}

//reverseChild : 将子节点list反转为正序
func reverseChild(res []*dto.ResourceDto) {
	for _, val := range res {
		if val.Children != nil {
			left := 0
			right := len(val.Children) - 1
			for left < right {
				temp := val.Children[left]
				val.Children[left] = val.Children[right]
				val.Children[right] = temp
				left++
				right--
			}
			reverseChild(val.Children)
		}
	}
}

// FindUserResources 获取用户的权限
func (r *ResourceDao) FindUserResources(ctx context.Context, userId string) ([]*dto.ResourceDto, public.BusinessException) {
	stmt, err := public.DB.Prepare("select distinct r.id, r.name, r.page, r.request, r.parent from role_user ru, role_resource rr, resource r where ru.user_id = ? and ru.role_id = rr.role_id and rr.resource_id = r.id order by r.id")
	if err != nil {
		return nil, public.NewBusinessException(public.PREPARE_SQL_ERROR)
	}
	defer stmt.Close()
	rows, err := stmt.Query(userId)
	if err != nil {
		return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}

	return scanResources(rows)
}

// scanResources : 从查询得到的rows中将数据scan到返回值中
func scanResources(rows *sql.Rows) ([]*dto.ResourceDto, public.BusinessException) {
	var res []*dto.ResourceDto

	for rows.Next() {
		r := &Resource{}
		err := rows.Scan(&r.Id, &r.Name, &r.Page, &r.Request, &r.Parent)
		if err != nil {
			log.Println("row scan failed, err is " + err.Error())
			return nil, public.NewBusinessException(public.ROW_SCAN_ERROR)
		}
		dto := &dto.ResourceDto{
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

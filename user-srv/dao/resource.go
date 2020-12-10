package dao

import (
	"context"
	"course/public"
	"course/public/util"
	"course/user-srv/proto/dto"
	"encoding/json"
	"gorm.io/gorm"
	"log"
	"time"
)

type ResourceDao struct {
}
type Resource struct {
	//可能为空的字段使用指针类型
	Id         string
	Title      string
	Component  string
	Name       string
	Icon       string
	Path       string
	Parent     string
	Request    string
	CreateBy   string
	UpdateBy   string
	CreateTime time.Time
	UpdateTime time.Time
}

func (Resource) TableName() string {
	return "menu"
}

// Delete : 删除权限
func (r *ResourceDao) Delete(ctx context.Context, id string) public.BusinessException {
	public.DB.Delete(&Resource{Id: id})
	return public.NoException("")
}

//SaveJson : 保存权限树
func (r *ResourceDao) SaveJson(ctx context.Context, jsonStr string) public.BusinessException {
	// rollback 逻辑
	var exception public.BusinessException
	var tx *gorm.DB
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
	tx = public.DB.Begin()
	if err = tx.Error; err != nil {
		return public.NewBusinessException(public.BEGIN_TRANSACTION_ERROR)
	}
	if err = tx.Exec("delete from resource").Error; err != nil {
		return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	for _, resourceDto := range list {
		err = tx.Exec("insert into resource (id, name, page, request, parent) VALUES (?, ?, ?, ?, ?)", resourceDto.Id, resourceDto.Name, resourceDto.Path, resourceDto.Request, resourceDto.Parent).Error
		if err != nil {
			return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
		}
	}

	if err = tx.Commit().Error; err != nil {
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
	var res []*dto.ResourceDto
	err := public.DB.Raw("select * from menu r").Find(&res).Error
	if err != nil {
		return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	//res, exception := scanResources(rows)
	//if exception.Code() != int32(public.OK) {
	//	return nil, exception
	//}
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
	var resources []*Resource
	err := public.DB.Raw("select distinct r.id, r.name, r.title, r.request, r.parent, r.path, r.component, r.icon from role_user ru, role_menu rr, menu r where ru.user_id = ? and ru.role_id = rr.role_id and rr.resource_id = r.id order by r.id", userId).Find(&resources).Error

	if err != nil {
		return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}

	res := make([]*dto.ResourceDto, len(resources))
	for i, resource := range resources {
		r := dto.ResourceDto{}
		_ = util.CopyProperties(&r, resource)
		res[i] = &r
	}
	return res, public.NoException("")
}

// scanResources : 从查询得到的rows中将数据scan到返回值中
//func scanResources(rows *sql.Rows) ([]*dto.ResourceDto, public.BusinessException) {
//	var res []*dto.ResourceDto
//
//	for rows.Next() {
//		r := &Resource{}
//		err := rows.Scan(&r.Id, &r.Name, &r.Page, &r.Request, &r.Parent)
//		if err != nil {
//			log.Println("row scan failed, err is " + err.Error())
//			return nil, public.NewBusinessException(public.ROW_SCAN_ERROR)
//		}
//		d := &dto.ResourceDto{
//			Id:   r.Id,
//			Name: r.Name,
//		}
//		if r.Page == nil {
//			d.Page = ""
//		} else {
//			d.Page = *r.Page
//		}
//		if r.Parent == nil {
//			d.Parent = ""
//		} else {
//			d.Parent = *r.Parent
//		}
//		if r.Request == nil {
//			d.Request = ""
//		} else {
//			d.Request = *r.Request
//		}
//		res = append(res, d)
//	}
//	return res, public.NoException("")
//}

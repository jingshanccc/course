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
	Id         int32
	Parent     int32
	Type       int32
	Title      string
	Name       string
	Component  string
	Sort       int32
	Icon       string
	Path       string
	IFrame     bool
	Cache      bool
	Hidden     bool
	Permission string
	CreateBy   string
	UpdateBy   string
	CreateTime time.Time
	UpdateTime time.Time
}

func (Resource) TableName() string {
	return "menu"
}

// Delete : 删除权限
func (r *ResourceDao) Delete(ctx context.Context, id int32) public.BusinessException {
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
			//resourceDto.Parent = ""
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
		err = tx.Exec("insert into resource (id, name, page, request, parent) VALUES (?, ?, ?, ?, ?)", resourceDto.Id, resourceDto.Name, resourceDto.Path, resourceDto.Parent, resourceDto.Parent).Error
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
	var resources []*Resource
	err := public.DB.Raw("select * from menu r order by id").Find(&resources).Error
	if err != nil {
		return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	res := make([]*dto.ResourceDto, len(resources))
	for i, resource := range resources {
		r := dto.ResourceDto{}
		_ = util.CopyProperties(&r, resource)
		res[i] = &r
	}
	buildTree(res)
	return res, public.NoException("")
}

// FindUserResources 获取用户的权限
func (r *ResourceDao) FindUserResources(ctx context.Context, userId string) ([]*dto.ResourceDto, public.BusinessException) {
	var resources []*Resource
	err := public.DB.Raw("select r.* from role_user ru, role_menu rr, menu r where ru.user_id = ? and type != ? and ru.role_id = rr.role_id and rr.resource_id = r.id order by r.sort asc", userId, 2).Find(&resources).Error
	if err != nil {
		return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}

	res := make([]*dto.ResourceDto, len(resources))
	for i, resource := range resources {
		r := dto.ResourceDto{}
		_ = util.CopyProperties(&r, resource)
		res[i] = &r
	}
	res = buildTree(res)
	return buildMenus(res), public.NoException("")
}

// 将权限转化为前端侧边栏要求的格式
func buildMenus(resources []*dto.ResourceDto) []*dto.ResourceDto {
	var res []*dto.ResourceDto
	for _, resource := range resources {
		children := resource.Children
		r := &dto.ResourceDto{}
		if resource.Name == "" {
			r.Name = resource.Title
		} else {
			r.Name = resource.Name
		}
		if resource.Parent == 0 { //一级菜单需要添加"/"消除警告
			r.Path = "/" + resource.Path
		} else {
			r.Path = resource.Path
		}
		r.Hidden = resource.Hidden
		if !resource.IFrame {
			if resource.Parent == 0 {
				if resource.Component == "" {
					r.Component = "Layout"
				} else {
					r.Component = resource.Component
				}
			} else if resource.Component != "" {
				r.Component = resource.Component
			}
		}
		r.Title = resource.Title
		r.Icon = resource.Icon
		r.Cache = resource.Cache
		if children != nil && len(children) > 0 {
			r.Redirect = "noRedirect"
			r.AlwaysShow = true
			r.Children = buildMenus(children)
		} else if resource.Parent == 0 { //一级菜单并且没有子菜单
			rr := &dto.ResourceDto{}
			rr.Title = r.Title
			rr.Icon = r.Icon
			rr.Cache = r.Cache
			if !resource.IFrame {
				rr.Path = "index"
				rr.Name = r.Name
				rr.Component = r.Component
			} else {
				rr.Path = resource.Path
			}
			r.Name = ""
			r.Component = "Layout"
			children1 := []*dto.ResourceDto{rr}
			r.Children = children1
		}
		res = append(res, r)
	}
	return res
}

// buildTree: 将资源转换为树结构
func buildTree(res []*dto.ResourceDto) []*dto.ResourceDto {
	var resp []*dto.ResourceDto
	for i := len(res) - 1; i >= 0; i-- {
		child := res[i]
		if child.Parent == 0 {
			resp = append(resp, child)
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
	reverseChild(resp)
	return resp
}

//reverseChild : 将子节点list反转为正序
func reverseChild(res []*dto.ResourceDto) {
	if res != nil {
		left := 0
		right := len(res) - 1
		for left < right {
			temp := res[left]
			res[left] = res[right]
			res[right] = temp
			left++
			right--
		}
		for _, re := range res {
			reverseChild(re.Children)
		}
	}
}

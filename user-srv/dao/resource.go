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
	Label      string
	Name       string
	Component  string
	Sort       int32
	Icon       string
	Path       string
	SubCount   int32
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

//SelectById: 通过 ID 获取权限
func (r *ResourceDao) SelectById(ctx context.Context, id int32) (*dto.ResourceDto, *public.BusinessException) {
	var resource Resource
	err := public.DB.Model(&Resource{}).Where("id = ?", id).Find(&resource).Error
	if err != nil {
		return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	var res dto.ResourceDto
	_ = util.CopyProperties(&res, resource)
	return &res, nil
}

// SelectPermissionByUserId: 获取permission字段
func (r *ResourceDao) SelectPermissionByUserId(ctx context.Context, userId string) ([]string, *public.BusinessException) {
	var res []string
	err := public.DB.Raw("select distinct r.permission from role_user ru, role_menu rr, menu r where ru.user_id = ? and type != ? and ru.role_id = rr.role_id and rr.resource_id = r.id and r.permission IS NOT NULL order by r.sort asc", userId, 2).Find(&res).Error
	if err != nil {
		return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return res, nil
}

// Delete : 删除权限
func (r *ResourceDao) Delete(ctx context.Context, id int32) *public.BusinessException {
	public.DB.Delete(&Resource{Id: id})
	return nil
}

//SaveJson : 保存权限树
func (r *ResourceDao) SaveJson(ctx context.Context, jsonStr string) *public.BusinessException {
	// rollback 逻辑
	var exception *public.BusinessException
	var tx *gorm.DB
	defer func() {
		if exception != nil {
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
	return nil
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

// GetByParent : 通过父ID parent查询权限
func (r *ResourceDao) GetByParent(ctx context.Context, pid int32) ([]*dto.ResourceDto, *public.BusinessException) {
	var resources []*dto.ResourceDto
	db := public.DB.Model(&Resource{})
	if pid != 0 {
		db = db.Where("parent = ?", pid)
	} else {
		db = db.Where("parent is null")
	}
	err := db.Order("sort").Find(&resources).Error
	if err != nil {
		return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return resources, nil
}

// FindUserResources 获取用户的权限
func (r *ResourceDao) FindUserResources(ctx context.Context, userId string) ([]*dto.ResourceDto, *public.BusinessException) {
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
	return buildMenus(res), nil
}

//GetChildTree: 获取传入父 ID 的所有子权限
func (r *ResourceDao) GetChildTree(ctx context.Context, resourceList []*dto.ResourceDto, resourceSet *public.HashSet) {
	for _, resource := range resourceList {
		resourceSet.Add(resource)
		resourceDtos, _ := r.GetByParent(ctx, resource.Id)
		if resourceDtos != nil && len(resourceDtos) > 0 {
			r.GetChildTree(ctx, resourceDtos, resourceSet)
		}
	}
}

//List: 分页数据
func (r *ResourceDao) List(ctx context.Context, in *dto.ResourcePageDto) (int64, []*dto.ResourceDto, *public.BusinessException) {
	db := public.DB.Model(&Resource{})
	forCount, forPage := util.GeneratePageSql(in.CreateTime, in.Blurry, in.Sort, []string{"title", "component", "permission"}, "")
	var count int64
	err := db.Raw("select count(1) from menu x " + forCount).Find(&count).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return 0, nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	var res []*dto.ResourceDto
	err = db.Raw("select x.* from menu x "+forPage, (in.Page-1)*in.Size, in.Size).Find(&res).Error

	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return 0, nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return count, res, nil
}

//GetSuperior: 获取同级和父级的权限
func (r *ResourceDao) GetSuperior(ctx context.Context, ids []int32) ([]*dto.ResourceDto, *public.BusinessException) {
	var res []*dto.ResourceDto
	for _, id := range ids {
		resource, _ := r.SelectById(ctx, id)
		var list []*dto.ResourceDto
		res = append(res, r._getSuperior(ctx, resource, list)...)
	}
	reverseChild(res)
	res = buildTree(res)
	reverseChild(res)
	return res, nil
}

func (r *ResourceDao) _getSuperior(ctx context.Context, resourceDto *dto.ResourceDto, list []*dto.ResourceDto) []*dto.ResourceDto {
	l, _ := r.GetByParent(ctx, resourceDto.Parent)
	if l != nil {
		list = append(list, l...)
	}
	if resourceDto.Parent == 0 {
		return list
	} else {
		rr, _ := r.SelectById(ctx, resourceDto.Parent)
		return r._getSuperior(ctx, rr, list)
	}
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

// buildTree: 将资源转换为树结构 此时 res 结构需按照 sort 升序排列
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

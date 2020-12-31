package dao

import (
	"context"
	"course/public"
	"course/public/util"
	"course/user-srv/proto/dto"
	"log"
	"strings"
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

const EntityName = "权限"

func (Resource) TableName() string {
	return "menu"
}

//SelectByProperty: 通过 一个属性 获取权限
func (r *ResourceDao) SelectByProperty(ctx context.Context, property string, value interface{}) (*dto.ResourceDto, *public.BusinessException) {
	var resource Resource
	err := public.DB.Model(&Resource{}).Where(property+" = ?", value).Find(&resource).Error
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
func (r *ResourceDao) Delete(ctx context.Context, ids []int32) *public.BusinessException {
	// 获取传入 id 的权限及所有子权限
	resourceSet := public.NewHashSet()
	exception := r.GetChildMenus(ctx, ids, resourceSet)
	if exception != nil {
		return exception
	}
	var resourceIds []int32
	var parents []int32
	for _, res := range resourceSet.Values() {
		re := res.(*dto.ResourceDto)
		resourceIds = append(resourceIds, re.Id)
		parents = append(parents, re.Parent)
	}
	// 解绑角色-权限
	exception = roleResourceDao.DeleteByResources(ctx, resourceIds)
	if exception != nil {
		return exception
	}
	// 删除权限
	public.DB.Delete(Resource{}, "id in ?", resourceIds)
	// 更新子权限数
	for _, id := range parents {
		r.updateSubCnt(id)
	}
	return nil
}

//Save: 保存/更新资源
func (r *ResourceDao) Save(ctx context.Context, in *dto.ResourceDto) *public.BusinessException {
	if in.IFrame {
		lower := strings.ToLower(in.Path)
		if !(strings.HasPrefix(lower, "http://") || strings.HasPrefix(lower, "https://")) {
			return public.BadRequestException("外链必须以http://或者https://开头")
		}
	}
	byTitle, _ := r.SelectByProperty(ctx, "title", in.Title)
	if byTitle != nil && byTitle.Id != 0 && byTitle.Id != in.Id {
		return public.EntityExistException(EntityName, "名称", in.Title)
	}
	if in.Component != "" {
		byComponent, _ := r.SelectByProperty(ctx, "component", in.Component)
		if byComponent != nil && byComponent.Id != 0 && byComponent.Id != in.Id {
			return public.EntityExistException(EntityName, "组件名称", in.Component)
		}
	}
	var resource Resource
	_ = util.CopyProperties(&resource, in)
	db := public.DB.Model(&Resource{})
	if in.Id == 0 { // create
		resource.CreateTime = time.Now()
		resource.UpdateTime = time.Now()
		resource.CreateBy = resource.UpdateBy
		err := db.Create(&resource).Error
		if err != nil {
			return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
		}
		r.updateSubCnt(resource.Parent)
	} else { // update
		if in.Id == in.Parent {
			return public.BadRequestException("父级权限不能为自己！")
		}
		byId, _ := r.SelectByProperty(ctx, "id", in.Id)
		if byId == nil || byId.Id == 0 {
			return public.BadRequestException("不存在当前记录！")
		}
		resource.Id = 0
		resource.UpdateTime = time.Now()
		err := db.Where("id = ?", in.Id).Updates(&resource).Error
		if err != nil {
			return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
		}
		if in.Parent != byId.Parent {
			r.updateSubCnt(in.Parent)
			r.updateSubCnt(byId.Parent)
		}
	}
	return nil
}

// GetByParent : 通过父ID parent查询权限
func (r *ResourceDao) GetByParent(ctx context.Context, pid int32) ([]*dto.ResourceDto, *public.BusinessException) {
	var resources []*dto.ResourceDto
	err := public.DB.Model(&Resource{}).Where("parent = ?", pid).Order("sort").Find(&resources).Error
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

//GetChildMenus: 获取传入父 ID 的所有子权限
func (r *ResourceDao) GetChildMenus(ctx context.Context, ids []int32, resourceSet *public.HashSet) *public.BusinessException {
	for _, id := range ids {
		parent, _ := r.SelectByProperty(ctx, "id", id)
		resourceSet.Add(parent)
		resourceDtos, exception := r.GetByParent(ctx, parent.Id)
		if exception != nil {
			return exception
		}
		r.getChildTree(ctx, resourceDtos, resourceSet)
	}
	return nil
}

//getChildTree: 获取传入父 ID 的所有子权限
func (r *ResourceDao) getChildTree(ctx context.Context, resourceList []*dto.ResourceDto, resourceSet *public.HashSet) {
	for _, resource := range resourceList {
		resourceSet.Add(resource)
		resourceDtos, _ := r.GetByParent(ctx, resource.Id)
		if resourceDtos != nil && len(resourceDtos) > 0 {
			r.getChildTree(ctx, resourceDtos, resourceSet)
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
		resource, _ := r.SelectByProperty(ctx, "id", id)
		var list []*dto.ResourceDto
		res = append(res, r._getSuperior(ctx, resource, list)...)
	}
	reverseChild(res)
	res = buildTree(res)
	reverseChild(res)
	return res, nil
}

//_getSuperior: 获取同级和父级权限 递归过程
func (r *ResourceDao) _getSuperior(ctx context.Context, resourceDto *dto.ResourceDto, list []*dto.ResourceDto) []*dto.ResourceDto {
	l, _ := r.GetByParent(ctx, resourceDto.Parent)
	if l != nil {
		list = append(list, l...)
	}
	if resourceDto.Parent == 0 {
		return list
	} else {
		rr, _ := r.SelectByProperty(ctx, "id", resourceDto.Parent)
		return r._getSuperior(ctx, rr, list)
	}
}

//updateSubCnt: 更新传入 id 的子菜单数
func (r *ResourceDao) updateSubCnt(parent int32) {
	if parent != 0 {
		var count int64
		public.DB.Model(&Resource{}).Where("parent = ?", parent).Count(&count)
		public.DB.Model(&Resource{}).Where("id = ?", parent).Update("sub_count", count)
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

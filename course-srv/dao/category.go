package dao

import (
	"context"
	"course/course-srv/proto/dto"
	"course/public"
	"course/public/util"
	"gorm.io/gorm"
	"log"
)

type CategoryDao struct {
}
type Category struct {
	Id     string
	Name   string
	Parent string
	Sort   int32
}

func (Category) TableName() string {
	return "category"
}

//PrimaryCategory: 获取所有一级分类
func (c *CategoryDao) PrimaryCategory() ([]*dto.CategoryDto, *public.BusinessException) {
	var res []*dto.CategoryDto
	err := public.DB.Model(&Category{}).Where("parent = '00000000'").Order("sort asc").Find(&res).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return res, nil
}

//All: 获取所有分类
func (c *CategoryDao) All() ([]*dto.CategoryDto, *public.BusinessException) {
	var res []*dto.CategoryDto
	err := public.DB.Model(&Category{}).Find(&res).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return res, nil
}

func (c *CategoryDao) Save(cd *dto.CategoryDto) (*dto.CategoryDto, *public.BusinessException) {
	if cd.Id != "" { //update
		err := public.DB.Model(&Category{Id: cd.Id}).Updates(&Category{Name: cd.Name, Parent: cd.Parent, Sort: cd.Sort}).Error
		if err != nil {
			return &dto.CategoryDto{}, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
		}
	} else { //insert
		cd.Id = util.GetShortUuid()
		err := public.DB.Create(&Category{
			Id:     cd.Id,
			Name:   cd.Name,
			Parent: cd.Parent,
			Sort:   cd.Sort,
		}).Error
		if err != nil {
			return &dto.CategoryDto{}, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
		}
	}
	return cd, nil
}

// Delete 删除分类
func (c *CategoryDao) Delete(ids []string) *public.BusinessException {
	var db *gorm.DB
	var isErr bool
	var exception *public.BusinessException
	defer func() {
		if isErr {
			exception = public.NewBusinessException(public.EXECUTE_SQL_ERROR)
			db.Rollback()
		}
	}()
	db = public.DB.Begin()
	err := db.Delete(&Category{}, "id in ?", ids).Error
	if err != nil {
		isErr = true
	}
	err = db.Delete(&Category{}, "parent in ?", ids).Error
	if err != nil {
		isErr = true
	}
	db.Commit()
	return exception
}

func (c *CategoryDao) List(ctx context.Context, in *dto.CategoryPageDto) (int64, []*dto.CategoryDto, *public.BusinessException) {
	if in.Parent == "" {
		in.Parent = "00000000"
	}
	var beforeOrder string
	if in.Blurry != "" {
		beforeOrder = " and x.parent = ? "
	} else {
		beforeOrder = " where x.parent = ? "
	}
	forCount, forPage := util.GeneratePageSql(nil, in.Blurry, in.Sort, []string{"name"}, beforeOrder)

	var count int64
	err := public.DB.Model(&Category{}).Raw("select count(1) from category x "+forCount+beforeOrder, in.Parent).Find(&count).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return 0, nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	var res []*dto.CategoryDto
	err = public.DB.Model(&Category{}).Raw("select x.* from category x "+forPage, in.Parent, (in.Page-1)*in.Size, in.Size).Find(&res).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return 0, nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return count, res, nil
}

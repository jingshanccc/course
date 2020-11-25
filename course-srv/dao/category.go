package dao

import (
	"course/course-srv/proto/dto"
	"course/public"
	"course/public/util"
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

func (c *CategoryDao) All() ([]*dto.CategoryDto, public.BusinessException) {
	var res []*dto.CategoryDto
	err := public.DB.Model(&Category{}).Order("sort asc").Find(&res).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return res, public.NewBusinessException(public.OK)
}

func (c *CategoryDao) Save(cd *dto.CategoryDto) (*dto.CategoryDto, public.BusinessException) {
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
	return cd, public.NoException("")
}

// Delete 删除分类
func (c *CategoryDao) Delete(id string) public.BusinessException {
	err := public.DB.Delete(&Category{Id: id}).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return public.NoException("")
}

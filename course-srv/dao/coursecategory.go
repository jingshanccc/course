package dao

import (
	"course/course-srv/proto/dto"
	"course/public"
	"course/public/util"
	"gorm.io/gorm"
)

type CourseCategoryDao struct {
}

type CourseCategory struct {
	Id         string
	CourseId   string
	CategoryId string
}

func (CourseCategory) TableName() string {
	return "course_category"
}

//SelectByCourseId: 查询课程关联的所有分类记录
func (r *CourseCategoryDao) SelectByCourseId(courseId string) ([]*dto.CourseCategoryDto, *public.BusinessException) {
	var res []*dto.CourseCategoryDto
	err := public.DB.Model(&CourseCategory{}).Where("course_id = ?", courseId).Find(&res).Error
	if err != nil {
		return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return res, nil
}

//BatchInsert: 批量插入
func (r *CourseCategoryDao) BatchInsert(courseId string, categoryDtos []*dto.CategoryDto) *public.BusinessException {
	var err error
	var tx *gorm.DB
	res := public.NoException("")
	defer func() {
		if err != nil {
			tx.Rollback()
			res = public.NewBusinessException(public.EXECUTE_SQL_ERROR)
		}
	}()
	tx = public.DB.Begin()
	err = tx.Model(&CourseCategory{}).Where("course_id = ?", courseId).Delete(&CourseCategory{}).Error

	for _, categoryDto := range categoryDtos {
		err = tx.Create(&CourseCategory{
			Id:         util.GetShortUuid(),
			CourseId:   courseId,
			CategoryId: categoryDto.Id,
		}).Error
	}
	err = tx.Commit().Error
	return res
}

//CountByCategories: 分类是否被课程关联
func (r *CourseCategoryDao) CountByCategories(categories []string) int64 {
	var count int64
	public.DB.Model(&CourseCategory{}).Where("category_id in ?", categories).Count(&count)
	return count
}

package dao

import (
	"gitee.com/jingshanccc/course/course/proto/dto"
	"gitee.com/jingshanccc/course/public"
	"gitee.com/jingshanccc/course/public/util"
	"log"
)

type CourseFileDao struct {
}

type CourseFile struct {
	Id       string
	CourseId string
	Name     string
	Url      string
	Size     int32
}

func (CourseFile) TableName() string {
	return "course_file"
}

//List : get CourseFile list
func (c *CourseFileDao) List(courseId string) ([]*dto.CourseFileDto, *public.BusinessException) {
	var res []*dto.CourseFileDto
	err := public.DB.Model(&CourseFile{}).Where(&CourseFile{CourseId: courseId}).Find(&res).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}

	return res, nil
}

//Save: 保存/更新小节
func (c *CourseFileDao) Save(cd *dto.CourseFileDto) (*dto.CourseFileDto, *public.BusinessException) {
	courseFileEntity := &CourseFile{}
	_ = util.CopyProperties(courseFileEntity, cd)
	if cd.Id != "" { //update
		courseFileEntity.Id = ""
		err := public.DB.Model(&CourseFile{Id: cd.Id}).Updates(courseFileEntity).Error
		if err != nil {
			return &dto.CourseFileDto{}, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
		}
	} else { //insert
		cd.Id = util.GetShortUuid()
		courseFileEntity.Id = cd.Id
		err := public.DB.Create(courseFileEntity).Error
		if err != nil {
			return &dto.CourseFileDto{}, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
		}
	}
	return cd, nil
}

// Delete 删除小节
func (c *CourseFileDao) Delete(id string) *public.BusinessException {
	err := public.DB.Delete(&CourseFile{Id: id}).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return nil
}

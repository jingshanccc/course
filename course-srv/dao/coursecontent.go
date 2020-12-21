package dao

import (
	"course/course-srv/proto/dto"
	"course/public"
	"log"
)

type CourseContentDao struct {
}

type CourseContent struct {
	Id      string
	Content string
}

func (CourseContent) TableName() string {
	return "course_content"
}

func (c *CourseContentDao) SelectById(id string) (*dto.CourseContentDto, *public.BusinessException) {
	var res dto.CourseContentDto
	err := public.DB.Model(&CourseContent{}).Where("id = ?", id).Find(&res).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return &res, nil
}

//SaveContent: 插入/更新
func (c *CourseContentDao) SaveContent(ccd *dto.CourseContentDto) *public.BusinessException {
	if affected := public.DB.Model(&CourseContent{}).Where("id = ? ", ccd.Id).Update("content", ccd.Content).RowsAffected; affected <= 0 {
		err := public.DB.Create(&CourseContent{
			Id:      ccd.Id,
			Content: ccd.Content,
		}).Error
		if err != nil {
			log.Println("exec sql failed, err is " + err.Error())
			return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
		}
	}
	return nil
}

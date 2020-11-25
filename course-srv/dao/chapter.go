package dao

import (
	"course/course-srv/proto/dto"
	"course/public"
	"course/public/util"
	"log"
)

type ChapterDao struct {
}
type Chapter struct {
	Id       string
	CourseId string
	Name     string
}

func (Chapter) TableName() string {
	return "chapter"
}

//List : get Chapter page
func (c *ChapterDao) List(cd *dto.ChapterPageDto) ([]*dto.ChapterDto, public.BusinessException) {
	orderby := "desc"
	if cd.Asc {
		orderby = "asc"
	}
	var res []*dto.ChapterDto
	err := public.DB.Model(&Chapter{}).Where(&Chapter{CourseId: cd.CourseId}).Order(cd.SortBy + " " + orderby).Limit(int(cd.PageSize)).Offset(int((cd.PageNum - 1) * cd.PageSize)).Find(&res).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}

	return res, public.NewBusinessException(public.OK)
}

//Save: 保存/更新大章
func (c *ChapterDao) Save(cd *dto.ChapterDto) (*dto.ChapterDto, public.BusinessException) {
	if cd.Id != "" { //update
		err := public.DB.Model(&Chapter{Id: cd.Id}).Updates(&Chapter{Name: cd.Name, CourseId: cd.CourseId}).Error
		if err != nil {
			return &dto.ChapterDto{}, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
		}
	} else { //insert
		cd.Id = util.GetShortUuid()
		err := public.DB.Create(&Chapter{
			Id:       cd.Id,
			Name:     cd.Name,
			CourseId: cd.CourseId,
		}).Error
		if err != nil {
			return &dto.ChapterDto{}, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
		}
	}
	return cd, public.NoException("")
}

// Delete 删除大章
func (c *ChapterDao) Delete(id string) public.BusinessException {
	err := public.DB.Delete(&Chapter{Id: id}).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return public.NoException("")
}

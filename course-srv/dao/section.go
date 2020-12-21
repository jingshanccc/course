package dao

import (
	"course/course-srv/proto/dto"
	"course/public"
	"course/public/util"
	"log"
	"time"
)

type SectionDao struct {
}
type Section struct {
	// id
	Id string
	// 标题
	Title string
	// 课程|course.id
	CourseId string
	// 大章|chapter.id
	ChapterId string
	// 视频
	Video string
	// 时长|单位秒
	Time int32
	// 收费|枚举[SectionChargeEnum]：CHARGE("C", "收费"),FREE("F", "免费")
	Charge string
	// 顺序
	Sort int32
	// 创建时间
	CreatedAt time.Time
	// 修改时间
	UpdatedAt time.Time
}

func (Section) TableName() string {
	return "section"
}

//List : get Section page
func (c *SectionDao) List(cd *dto.SectionPageDto) ([]*dto.SectionDto, *public.BusinessException) {
	orderby := "desc"
	if cd.Asc {
		orderby = "asc"
	}
	var res []*dto.SectionDto
	err := public.DB.Model(&Section{}).Where(&Section{ChapterId: cd.ChapterId, CourseId: cd.CourseId}).Order(cd.SortBy + " " + orderby).Limit(int(cd.PageSize)).Offset(int((cd.PageNum - 1) * cd.PageSize)).Find(&res).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}

	return res, nil
}

//Save: 保存/更新小节
func (c *SectionDao) Save(cd *dto.SectionDto) (*dto.SectionDto, *public.BusinessException) {
	sectionEntity := &Section{}
	_ = util.CopyProperties(sectionEntity, cd)
	sectionEntity.UpdatedAt = time.Time{}
	if cd.Id != "" { //update
		sectionEntity.Id = ""
		err := public.DB.Model(&Section{Id: cd.Id}).Updates(sectionEntity).Error
		if err != nil {
			return &dto.SectionDto{}, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
		}
	} else { //insert
		cd.Id = util.GetShortUuid()
		sectionEntity.Id = cd.Id
		sectionEntity.CreatedAt = time.Time{}
		err := public.DB.Create(sectionEntity).Error
		if err != nil {
			return &dto.SectionDto{}, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
		}
	}
	return cd, nil
}

// Delete 删除小节
func (c *SectionDao) Delete(id string) *public.BusinessException {
	err := public.DB.Delete(&Section{Id: id}).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return nil
}

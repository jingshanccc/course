package dao

import (
	"gitee.com/jingshanccc/course/course/proto/dto"
	"gitee.com/jingshanccc/course/public"
	"gitee.com/jingshanccc/course/public/util"
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
	CreatedTime time.Time
	// 修改时间
	UpdatedTime time.Time
}

func (Section) TableName() string {
	return "section"
}

//List : get Section page
func (c *SectionDao) List(cd *dto.SectionPageDto) (int64, []*dto.SectionDto, *public.BusinessException) {
	var beforeOrder string
	if cd.Blurry != "" {
		beforeOrder = " and x.course_id = ? and x.chapter_id = ? "
	} else {
		beforeOrder = " where x.course_id = ? and x.chapter_id = ? "
	}
	forCount, forPage := util.GeneratePageSql(cd.CreateTime, cd.Blurry, cd.Sort, []string{"title"}, beforeOrder)

	var count int64
	err := public.DB.Model(&Section{}).Raw("select count(1) from section x "+forCount+beforeOrder, cd.CourseId, cd.ChapterId).Find(&count).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return 0, nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	var res []*dto.SectionDto
	err = public.DB.Model(&Section{}).Raw("select x.* from section x "+forPage, cd.CourseId, cd.ChapterId, (cd.Page-1)*cd.Size, cd.Size).Find(&res).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return 0, nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return count, res, nil
}

//Save: 保存/更新小节
func (c *SectionDao) Save(cd *dto.SectionDto) (*dto.SectionDto, *public.BusinessException) {
	var sectionEntity Section
	_ = util.CopyProperties(&sectionEntity, cd)
	sectionEntity.UpdatedTime = time.Now()
	if isSortAllow(sectionEntity) {
		if cd.Id != "" { //update
			sectionEntity.Id = ""
			err := public.DB.Model(&Section{}).Where("id = ?", cd.Id).Updates(&sectionEntity).Error
			if err != nil {
				return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
			}
		} else { //insert
			sectionEntity.Id = util.GetShortUuid()
			sectionEntity.CreatedTime = sectionEntity.UpdatedTime
			err := public.DB.Create(&sectionEntity).Error
			if err != nil {
				return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
			}
		}
	} else {
		return nil, public.BadRequestException("已存在该序号的小节，请重新输入！")
	}
	return cd, nil
}

//isSortAllow: 判断序号是否已存在
func isSortAllow(section Section) bool {
	var fromDb Section
	public.DB.Model(&Section{}).Where("chapter_id = ? and course_id = ? and sort = ?", section.ChapterId, section.CourseId, section.Sort).Find(&fromDb)
	if fromDb.Id == "" {
		return true
	} else {
		if section.Id != "" {
			return section.Id == fromDb.Id
		} else {
			return false
		}
	}
}

// Delete 删除小节
func (c *SectionDao) Delete(ids []string) *public.BusinessException {
	err := public.DB.Delete(&Section{}, "id in ?", ids).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return nil
}

func (c *SectionDao) SelectByProperty(property, value string) ([]*dto.SectionDto, *public.BusinessException) {
	var res []*dto.SectionDto
	err := public.DB.Model(&Section{}).Where(property+" = ?", value).Find(&res).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return res, nil
}

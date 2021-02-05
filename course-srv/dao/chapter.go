package dao

import (
	"course/course-srv/proto/dto"
	"course/public"
	"course/public/util"
	"log"
	"time"
)

type ChapterDao struct {
}
type Chapter struct {
	Id       string
	CourseId string
	Name     string
	Sort     int32
	// 创建时间
	CreateTime time.Time
	// 修改时间
	UpdateTime time.Time
	Creator    string
	Updater    string
}

func (Chapter) TableName() string {
	return "chapter"
}

//All: 获取一门课程的所有大章
func (c *ChapterDao) All(courseId string) ([]*dto.ChapterDto, *public.BusinessException) {
	var res []*dto.ChapterDto
	err := public.DB.Model(&Chapter{}).Where("course_id = ?", courseId).Order("sort").Find(&res).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return res, nil
}

//List : get Chapter page
func (c *ChapterDao) List(cd *dto.ChapterPageDto) (int64, []*dto.ChapterDto, *public.BusinessException) {
	var beforeOrder string
	if cd.Blurry != "" {
		beforeOrder = " and x.course_id = ? "
	} else {
		beforeOrder = " where x.course_id = ? "
	}
	forCount, forPage := util.GeneratePageSql(cd.CreateTime, cd.Blurry, cd.Sort, []string{"name"}, beforeOrder)

	var count int64
	err := public.DB.Model(&Chapter{}).Raw("select count(1) from chapter x "+forCount+beforeOrder, cd.CourseId).Find(&count).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return 0, nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	var res []*dto.ChapterDto
	err = public.DB.Model(&Chapter{}).Raw("select x.* from chapter x "+forPage, cd.CourseId, (cd.Page-1)*cd.Size, cd.Size).Find(&res).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return 0, nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return count, res, nil
}

//Save: 保存/更新大章
func (c *ChapterDao) Save(cd *dto.ChapterDto) (*dto.ChapterDto, *public.BusinessException) {
	var chapter Chapter
	_ = util.CopyProperties(&chapter, cd)
	chapter.UpdateTime = time.Now()
	if isAllowedSort(chapter) {
		if chapter.Id != "" { //update
			chapter.Id = ""
			err := public.DB.Model(&Chapter{}).Where("id = ?", cd.Id).Updates(&chapter).Error
			if err != nil {
				return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
			}
		} else { //insert
			chapter.CreateTime = chapter.UpdateTime
			chapter.Id = util.GetShortUuid()
			err := public.DB.Create(&chapter).Error
			if err != nil {
				return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
			}
		}
	} else {
		return nil, public.BadRequestException("已存在该序号的大章，请重新输入！")
	}
	return cd, nil
}

//isAllowedSort: 判断序号是否已存在
func isAllowedSort(chapter Chapter) bool {
	var fromDb Chapter
	public.DB.Model(&Chapter{}).Where("course_id = ? and sort = ?", chapter.CourseId, chapter.Sort).Find(&fromDb)
	if fromDb.Id == "" {
		return true
	} else {
		if chapter.Id != "" {
			return chapter.Id == fromDb.Id
		} else {
			return false
		}
	}
}

// Delete 删除大章
func (c *ChapterDao) Delete(ids []string) *public.BusinessException {
	err := public.DB.Delete(&Chapter{}, "id in ?", ids).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return nil
}

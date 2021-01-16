package dao

import (
	"course/course-srv/proto/dto"
	"course/public"
	"course/public/util"
	"gorm.io/gorm"
	"log"
	"time"
)

var courseCategoryDao = &CourseCategoryDao{}

type CourseDao struct {
}
type Course struct {
	Id        string
	Name      string
	Summary   string
	Time      int32
	Price     float64
	Image     string
	Level     string //级别
	Charge    string //收费
	Status    string
	Enroll    int32 //报名人数
	Sort      int32 //排序
	TeacherId string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Course) TableName() string {
	return "course"
}

//List : get course page
func (c *CourseDao) List(in *dto.CoursePageDto) (int64, []*dto.CourseDto, *public.BusinessException) {
	forCount, forPage := util.GeneratePageSql(in.CreateTime, in.Blurry, in.Sort, []string{"name", "summary"}, "")
	var count int64
	err := public.DB.Model(&Course{}).Raw("select count(1) from course x" + forCount).Find(&count).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return 0, nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	var res []*dto.CourseDto
	err = public.DB.Model(&Course{}).Raw("select x.* from course x"+forPage, (in.Page-1)*in.Size, in.Size).Find(&res).Error
	return count, res, nil
}

//Save: 保存/更新
func (c *CourseDao) Save(cd *dto.CourseDto) (*dto.CourseDto, *public.BusinessException) {
	courseEntity := &Course{}
	_ = util.CopyProperties(courseEntity, cd)
	courseEntity.UpdatedAt = time.Time{}
	if cd.Id != "" { //update
		courseEntity.Id = ""
		err := public.DB.Model(&Course{Id: cd.Id}).Updates(courseEntity).Error
		if err != nil {
			return &dto.CourseDto{}, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
		}
	} else { //insert
		cd.Id = util.GetShortUuid()
		courseEntity.Id = cd.Id
		courseEntity.CreatedAt = time.Time{}
		err := public.DB.Create(courseEntity).Error
		if err != nil {
			return &dto.CourseDto{}, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
		}
	}
	//保存分类
	err := courseCategoryDao.BatchInsert(cd.Id, cd.Categories)
	return cd, err
}

// Delete 删除课程
func (c *CourseDao) Delete(id string) *public.BusinessException {
	err := public.DB.Delete(&Course{Id: id}).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return nil
}

//UpdateSort: 更新课程排序
func (c *CourseDao) UpdateSort(tx *gorm.DB, sortDto *dto.SortDto) *public.BusinessException {
	err := tx.Model(&Course{}).Where("id = ? ", sortDto.Id).Update("sort", sortDto.NewSort).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return nil
}

//MoveSortsForward: 将顺序在oldSort~newSort之间的记录往前一个位置
func (c *CourseDao) MoveSortsForward(tx *gorm.DB, sortDto *dto.SortDto) *public.BusinessException {
	err := tx.Exec("update course set `sort`= (`sort`-1) where `sort` <= ? and `sort` >= ? and id != ?", sortDto.NewSort, sortDto.OldSort, sortDto.Id).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return nil
}

//MoveSortsForward: 将顺序在oldSort之前newSort之后的记录往后一个位置
func (c *CourseDao) MoveSortsBackward(tx *gorm.DB, sortDto *dto.SortDto) *public.BusinessException {
	err := tx.Exec("update course set `sort`= (`sort`+1) where `sort` >= ? and `sort` <= ? and id != ?", sortDto.NewSort, sortDto.OldSort, sortDto.Id).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return nil
}

//FindContent: 获取课程内容
func (c *CourseDao) FindContent(id string) (*dto.CourseContentDto, *public.BusinessException) {
	return (&CourseContentDao{}).SelectById(id)
}

//SaveContent: 保存课程内容
func (c *CourseDao) SaveContent(ccd *dto.CourseContentDto) *public.BusinessException {
	return (&CourseContentDao{}).SaveContent(ccd)
}

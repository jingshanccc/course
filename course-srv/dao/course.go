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
func (c *CourseDao) List(cd *dto.CoursePageDto) ([]*dto.CourseDto, public.BusinessException) {
	orderby := "desc"
	if cd.Asc {
		orderby = "asc"
	}
	var courses []Course
	err := public.DB.Model(&Course{}).Order(cd.SortBy + " " + orderby).Limit(int(cd.PageSize)).Offset(int((cd.PageNum - 1) * cd.PageSize)).Find(&courses).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	//解决CourseDto直接被sql赋值时列表字段未被sql忽略而产生的报错
	res := make([]*dto.CourseDto, len(courses))
	for index, val := range courses {
		r := dto.CourseDto{}
		_ = util.CopyProperties(&r, val)
		res[index] = &r
	}
	return res, public.NewBusinessException(public.OK)
}

//Save: 保存/更新
func (c *CourseDao) Save(cd *dto.CourseDto) (*dto.CourseDto, public.BusinessException) {
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
func (c *CourseDao) Delete(id string) public.BusinessException {
	err := public.DB.Delete(&Course{Id: id}).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return public.NoException("")
}

//UpdateSort: 更新课程排序
func (c *CourseDao) UpdateSort(tx *gorm.DB, sortDto *dto.SortDto) public.BusinessException {
	err := tx.Model(&Course{}).Where("id = ? ", sortDto.Id).Update("sort", sortDto.NewSort).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return public.NoException("")
}

//MoveSortsForward: 将顺序在oldSort~newSort之间的记录往前一个位置
func (c *CourseDao) MoveSortsForward(tx *gorm.DB, sortDto *dto.SortDto) public.BusinessException {
	err := tx.Exec("update course set `sort`= (`sort`-1) where `sort` <= ? and `sort` >= ? and id != ?", sortDto.NewSort, sortDto.OldSort, sortDto.Id).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return public.NoException("")
}

//MoveSortsForward: 将顺序在oldSort之前newSort之后的记录往后一个位置
func (c *CourseDao) MoveSortsBackward(tx *gorm.DB, sortDto *dto.SortDto) public.BusinessException {
	err := tx.Exec("update course set `sort`= (`sort`+1) where `sort` >= ? and `sort` <= ? and id != ?", sortDto.NewSort, sortDto.OldSort, sortDto.Id).Error
	if err != nil {
		log.Println("exec sql failed, err is " + err.Error())
		return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return public.NoException("")
}

//FindContent: 获取课程内容
func (c *CourseDao) FindContent(id string) (*dto.CourseContentDto, public.BusinessException) {
	return (&CourseContentDao{}).SelectById(id)
}

//SaveContent: 保存课程内容
func (c *CourseDao) SaveContent(ccd *dto.CourseContentDto) public.BusinessException {
	return (&CourseContentDao{}).SaveContent(ccd)
}

package dao

import (
	"gitee.com/jingshanccc/course/course/proto/dto"
	"gitee.com/jingshanccc/course/public"
	"gitee.com/jingshanccc/course/public/util"
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
	for _, cou := range res {
		var categories []string
		public.DB.Model(&CourseCategory{}).Select("category_id").Where(" course_id = ?", cou.Id).Find(&categories)
		cou.Categories = categories
	}
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
func (c *CourseDao) Delete(ids []string) *public.BusinessException {
	err := public.DB.Model(&Teacher{}).Where("id in ?", ids).Delete(&Course{}).Error
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

//UpdateCourseDuration: 更新课程时长
func (c *CourseDao) UpdateCourseDuration(id string) *public.BusinessException {
	err := public.DB.Model(&Course{}).Raw("update course c set c.time = (select sum(time) from section where course_id = c.id) where c.id = ?", id).Find(nil).Error
	if err != nil {
		return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return nil
}

//CarouselCourse: 轮播图课程
func (c *CourseDao) CarouselCourse() ([]*dto.CourseDto, *public.BusinessException) {
	var res []*dto.CourseDto
	err := public.DB.Model(&Course{}).Select("id", "image").Where("status = '发布'").Order("sort").Limit(4).Find(&res).Error
	if err != nil {
		return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return res, nil
}

//NewPublish: 新上好课
func (c *CourseDao) NewPublish() ([]*dto.CourseDto, *public.BusinessException) {
	var res []*dto.CourseDto
	err := public.DB.Model(&Course{}).Where("status = '发布'").Order("updated_at desc, sort").Limit(8).Find(&res).Error
	if err != nil {
		return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return res, nil
}

//SelectCourseByIds: 批量获取课程 all-当ids为空时是否返回所有course
func (c *CourseDao) SelectCourseByIds(ids []string, all bool) ([]*dto.CourseDto, *public.BusinessException) {
	var res []*dto.CourseDto
	db := public.DB.Model(&Course{})
	if !all && len(ids) == 0 {
		return nil, nil
	}
	if len(ids) > 0 {
		db = db.Where("id in ?", ids)
	}
	err := db.Where("status = '发布'").Order("sort").Find(&res).Error
	if err != nil {
		return nil, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return res, nil
}

//SelectByProperty: 查询课程
func (c *CourseDao) SelectByProperty(property, value string) *Course {
	var res Course
	public.DB.Model(&Course{}).Where(property+" = ?", value).Find(&res)
	return &res
}

package dao

import (
	"gitee.com/jingshanccc/course/course/proto/dto"
	"gitee.com/jingshanccc/course/public"
	"time"
)

type MemberCourse struct {
	Id        uint32
	MemberId  string
	CourseId  string
	LastLearn string // 上次学习到 格式为 1 1 1280 第一章第一节第1280秒
	CreateAt  time.Time
	UpdateAt  time.Time
}

func (m MemberCourse) TableName() string {
	return "member_course"
}

type MemberCourseDao struct {
}

func (m *MemberCourseDao) MyCourse(str string) ([]*dto.CourseDto, *public.BusinessException) {
	var res []*dto.CourseDto
	err := public.DB.Model(&MemberCourse{}).Select("course.*").Where("member_course.member_id = ?", str).Joins("join course on course.id = member_course.course_id").Scan(&res).Error
	if err != nil {
		return res, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return res, nil
}

func (m *MemberCourseDao) AddToMyCourse(data *MemberCourse) *public.BusinessException {
	err := public.DB.Create(&data).Error
	if err != nil {
		return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return nil
}

func (m *MemberCourseDao) CourseInfo(memberId, courseId string) (string, *public.BusinessException) {
	var info string
	err := public.DB.Model(&MemberCourse{}).Select("last_learn").Where("member_id = ?", memberId).Where("course_id", courseId).Find(&info).Error
	if err != nil {
		return info, public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return info, nil
}

func (m *MemberCourseDao) SaveLearnInfo(rows []string) *public.BusinessException {
	err := public.DB.Model(&MemberCourse{}).Where("member_id = ? and course_id = ?", rows[2], rows[0]).Update("last_learn", rows[1]).Error
	if err != nil {
		return public.NewBusinessException(public.EXECUTE_SQL_ERROR)
	}
	return nil
}

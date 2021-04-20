package handler

import "gitee.com/jingshanccc/course/course/dao"

type CourseServiceHandler struct {
}

var (
	categoryDao       = &dao.CategoryDao{}
	courseDao         = &dao.CourseDao{}
	courseCategoryDao = &dao.CourseCategoryDao{}
	chapterDao        = &dao.ChapterDao{}
	sectionDao        = &dao.SectionDao{}
	courseFileDao     = &dao.CourseFileDao{}
	teacherDao        = &dao.TeacherDao{}
	memberCourseDao   = &dao.MemberCourseDao{}
)

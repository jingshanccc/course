package handler

import "course/course-srv/dao"

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
)

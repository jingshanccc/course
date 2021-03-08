package course

import (
	"context"
	"course/config"
	"course/course-srv/proto/course"
	"course/course-srv/proto/dto"
	"course/proto/basic"
	"course/public"
	"github.com/gin-gonic/gin"
)

//ListCourse: 获取课程列表页
func ListCourse(ctx *gin.Context) {
	var req dto.CoursePageDto
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[config.CourseServiceName].(course.CourseService)
		list, err := courseService.CourseList(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//CarouselCourse: 轮播图课程
func CarouselCourse(ctx *gin.Context) {
	courseService := ctx.Keys[config.CourseServiceName].(course.CourseService)
	list, err := courseService.CarouselCourse(context.Background(), &basic.String{})
	public.ResponseAny(ctx, err, list)
}

//NewPublishCourse: 新上好课
func NewPublishCourse(ctx *gin.Context) {
	courseService := ctx.Keys[config.CourseServiceName].(course.CourseService)
	list, err := courseService.NewPublishCourse(context.Background(), &basic.String{})
	public.ResponseAny(ctx, err, list)
}

//SaveCourse: 保存课程
func SaveCourse(ctx *gin.Context) {
	var req dto.CourseDto
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[config.CourseServiceName].(course.CourseService)
		list, err := courseService.SaveCourse(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//DelCourse: 删除课程
func DelCourse(ctx *gin.Context) {
	var req basic.StringList
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[config.CourseServiceName].(course.CourseService)
		list, err := courseService.DeleteCourse(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//FindCourseContent: 获取课程内容
func FindCourseContent(ctx *gin.Context) {
	var req basic.String
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[config.CourseServiceName].(course.CourseService)
		list, err := courseService.FindCourseContent(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//SaveCourseContent: 保存课程内容
func SaveCourseContent(ctx *gin.Context) {
	var req dto.CourseContentDto
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[config.CourseServiceName].(course.CourseService)
		list, err := courseService.SaveCourseContent(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//ListCourseCategory: 获取课程所属的所有分类
func ListCourseCategory(ctx *gin.Context) {
	var req basic.String
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[config.CourseServiceName].(course.CourseService)
		list, err := courseService.ListCourseCategory(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//SortCourse: 课程排序
func SortCourse(ctx *gin.Context) {
	var req dto.SortDto
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[config.CourseServiceName].(course.CourseService)
		list, err := courseService.SortCourse(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//ListCourseFile: 获取课程文件
func ListCourseFile(ctx *gin.Context) {
	var req basic.String
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[config.CourseServiceName].(course.CourseService)
		list, err := courseService.ListCourseFile(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//SaveCourseFile: 保存课程
func SaveCourseFile(ctx *gin.Context) {
	var req dto.CourseFileDto
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[config.CourseServiceName].(course.CourseService)
		list, err := courseService.SaveCourseFile(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//DelCourseFile: 删除课程
func DelCourseFile(ctx *gin.Context) {
	var req basic.String
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[config.CourseServiceName].(course.CourseService)
		list, err := courseService.DeleteCourseFile(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

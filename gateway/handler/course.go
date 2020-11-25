package handler

import (
	"context"
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
		courseService := ctx.Keys[public.CourseServiceName].(course.CourseService)
		list, err := courseService.CourseList(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//SaveCourse: 保存课程
func SaveCourse(ctx *gin.Context) {
	var req dto.CourseDto
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[public.CourseServiceName].(course.CourseService)
		list, err := courseService.SaveCourse(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//DelCourse: 删除课程
func DelCourse(ctx *gin.Context) {
	var req basic.String
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[public.CourseServiceName].(course.CourseService)
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
		courseService := ctx.Keys[public.CourseServiceName].(course.CourseService)
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
		courseService := ctx.Keys[public.CourseServiceName].(course.CourseService)
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
		courseService := ctx.Keys[public.CourseServiceName].(course.CourseService)
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
		courseService := ctx.Keys[public.CourseServiceName].(course.CourseService)
		list, err := courseService.SortCourse(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

/*------------- Category ----------------*/

//AllCategory: 获取所有分类
func AllCategory(ctx *gin.Context) {
	var req basic.String
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[public.CourseServiceName].(course.CourseService)
		list, err := courseService.AllCategory(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//SaveCategory: 保存分类
func SaveCategory(ctx *gin.Context) {
	var req dto.CategoryDto
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[public.CourseServiceName].(course.CourseService)
		list, err := courseService.SaveCategory(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//DeleteCategory: 删除分类
func DeleteCategory(ctx *gin.Context) {
	var req basic.String
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[public.CourseServiceName].(course.CourseService)
		list, err := courseService.DeleteCategory(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

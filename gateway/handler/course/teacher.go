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

//AllTeacher: 获取所有讲师
func AllTeacher(ctx *gin.Context) {
	var req basic.String
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[config.CourseServiceName].(course.CourseService)
		list, err := courseService.AllTeacher(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//ListTeacher: get Teacher page
func ListTeacher(ctx *gin.Context) {
	var req dto.TeacherPageDto
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[config.CourseServiceName].(course.CourseService)
		list, err := courseService.ListTeacher(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//SaveTeacher: 保存讲师
func SaveTeacher(ctx *gin.Context) {
	var req dto.TeacherDto
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[config.CourseServiceName].(course.CourseService)
		list, err := courseService.SaveTeacher(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//DeleteTeacher: 删除讲师
func DeleteTeacher(ctx *gin.Context) {
	var req basic.StringList
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[config.CourseServiceName].(course.CourseService)
		list, err := courseService.DeleteTeacher(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

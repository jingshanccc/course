package course

import (
	"context"
	"course/course-srv/proto/course"
	"course/course-srv/proto/dto"
	"course/proto/basic"
	"course/public"
	"github.com/gin-gonic/gin"
)

func ListSection(ctx *gin.Context) {
	var req dto.SectionPageDto
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[public.CourseServiceName].(course.CourseService)
		list, err := courseService.ListSection(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//SaveSection: 保存课程
func SaveSection(ctx *gin.Context) {
	var req dto.SectionDto
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[public.CourseServiceName].(course.CourseService)
		list, err := courseService.SaveSection(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//DelSection: 删除课程
func DelSection(ctx *gin.Context) {
	var req basic.String
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[public.CourseServiceName].(course.CourseService)
		list, err := courseService.DeleteSection(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

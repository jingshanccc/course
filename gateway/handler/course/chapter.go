package course

import (
	"context"
	"course/course-srv/proto/course"
	"course/course-srv/proto/dto"
	"course/proto/basic"
	"course/public"
	"github.com/gin-gonic/gin"
)

func ListChapter(ctx *gin.Context) {
	var req dto.ChapterPageDto
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[public.CourseServiceName].(course.CourseService)
		list, err := courseService.ListChapter(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//SaveChapter: 保存课程
func SaveChapter(ctx *gin.Context) {
	var req dto.ChapterDto
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[public.CourseServiceName].(course.CourseService)
		list, err := courseService.SaveChapter(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//DelChapter: 删除课程
func DelChapter(ctx *gin.Context) {
	var req basic.String
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[public.CourseServiceName].(course.CourseService)
		list, err := courseService.DeleteChapter(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

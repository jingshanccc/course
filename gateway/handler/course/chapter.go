package course

import (
	"context"
	"gitee.com/jingshanccc/course/course/proto/course"
	"gitee.com/jingshanccc/course/course/proto/dto"
	"gitee.com/jingshanccc/course/public"
	"gitee.com/jingshanccc/course/public/config"
	"gitee.com/jingshanccc/course/public/proto/basic"
	"github.com/gin-gonic/gin"
)

func ListChapter(ctx *gin.Context) {
	var req dto.ChapterPageDto
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name].(course.CourseService)
		list, err := courseService.ListChapter(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}
func AllChapter(ctx *gin.Context) {
	var req basic.String
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name].(course.CourseService)
		list, err := courseService.AllChapter(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//SaveChapter: 保存课程
func SaveChapter(ctx *gin.Context) {
	var req dto.ChapterDto
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name].(course.CourseService)
		list, err := courseService.SaveChapter(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//DelChapter: 删除课程
func DelChapter(ctx *gin.Context) {
	var req basic.StringList
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name].(course.CourseService)
		list, err := courseService.DeleteChapter(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

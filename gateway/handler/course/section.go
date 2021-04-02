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

func ListSection(ctx *gin.Context) {
	var req dto.SectionPageDto
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name].(course.CourseService)
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
		courseService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name].(course.CourseService)
		list, err := courseService.SaveSection(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//DelSection: 删除课程
func DelSection(ctx *gin.Context) {
	var req basic.StringList
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name].(course.CourseService)
		list, err := courseService.DeleteSection(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

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

//ListCategory: 获取分类列表
func ListCategory(ctx *gin.Context) {
	var req dto.CategoryPageDto
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name].(course.CourseService)
		list, err := courseService.ListCategory(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//PrimaryCategory: 获取所有一级分类
func PrimaryCategory(ctx *gin.Context) {
	var req basic.String
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name].(course.CourseService)
		list, err := courseService.PrimaryCategory(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//AllCategory: 获取所有分类
func AllCategory(ctx *gin.Context) {
	var req basic.String
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name].(course.CourseService)
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
		courseService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name].(course.CourseService)
		list, err := courseService.SaveCategory(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//DeleteCategory: 删除分类
func DeleteCategory(ctx *gin.Context) {
	var req basic.StringList
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name].(course.CourseService)
		list, err := courseService.DeleteCategory(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

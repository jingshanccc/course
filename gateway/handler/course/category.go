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

//ListCategory: 获取分类列表
func ListCategory(ctx *gin.Context) {
	var req dto.CategoryPageDto
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[config.CourseServiceName].(course.CourseService)
		list, err := courseService.ListCategory(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//AllCategory: 获取所有分类
func AllCategory(ctx *gin.Context) {
	var req basic.String
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[config.CourseServiceName].(course.CourseService)
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
		courseService := ctx.Keys[config.CourseServiceName].(course.CourseService)
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
		courseService := ctx.Keys[config.CourseServiceName].(course.CourseService)
		list, err := courseService.DeleteCategory(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

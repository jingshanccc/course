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

//AllTeacher: 获取所有讲师
func AllTeacher(ctx *gin.Context) {
	var req basic.String
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name].(course.CourseService)
		list, err := courseService.AllTeacher(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//SearchTeacher: 获取所有讲师
func SearchTeacher(ctx *gin.Context) {
	var req basic.String
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name].(course.CourseService)
		list, err := courseService.SearchTeacher(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//ListTeacher: get Teacher page
func ListTeacher(ctx *gin.Context) {
	var req dto.TeacherPageDto
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name].(course.CourseService)
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
		courseService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name].(course.CourseService)
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
		courseService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name].(course.CourseService)
		list, err := courseService.DeleteTeacher(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

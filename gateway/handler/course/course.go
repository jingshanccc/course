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

//ListCourse: 获取课程列表页
func ListCourse(ctx *gin.Context) {
	var req dto.CoursePageDto
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name].(course.CourseService)
		list, err := courseService.CourseList(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//CarouselCourse: 轮播图课程
func CarouselCourse(ctx *gin.Context) {
	courseService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name].(course.CourseService)
	list, err := courseService.CarouselCourse(context.Background(), &basic.String{})
	public.ResponseAny(ctx, err, list)
}

//NewPublishCourse: 新上好课
func NewPublishCourse(ctx *gin.Context) {
	courseService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name].(course.CourseService)
	list, err := courseService.NewPublishCourse(context.Background(), &basic.String{})
	public.ResponseAny(ctx, err, list)
}

//CategoryCourse: 分类课程
func CategoryCourse(ctx *gin.Context) {
	var req basic.String
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name].(course.CourseService)
		list, err := courseService.CategoryCourse(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//RelatedCourse: 相关课程
func RelatedCourse(ctx *gin.Context) {
	var req basic.String
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name].(course.CourseService)
		list, err := courseService.RelatedCourse(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//MyCourse: 我的课程
func MyCourse(ctx *gin.Context) {
	var req basic.String
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name].(course.CourseService)
		list, err := courseService.MyCourse(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//AddToMyCourse: 添加到我的课程
func AddToMyCourse(ctx *gin.Context) {
	var req basic.String
	//user, _ := middleware.GetCurrentUser(ctx)
	user := "IX7UQht2"
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name].(course.CourseService)
		list, err := courseService.AddToMyCourse(context.Background(), &dto.MemberCourseDto{
			MemberId: user,
			CourseId: req.Str,
		})
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//CourseInfo: 获取课程学习进度
func CourseInfo(ctx *gin.Context) {
	var req basic.String
	//user, _ := middleware.GetCurrentUser(ctx)
	user := "IX7UQht2"
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name].(course.CourseService)
		list, err := courseService.CourseInfo(context.Background(), &basic.StringList{Rows: []string{user, req.Str}})
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//DownloadCourseContent: 下载课程讲义
func DownloadCourseContent(ctx *gin.Context) {
	var req basic.String
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name].(course.CourseService)
		bytes, err := courseService.DownloadCourseContent(context.Background(), &req)
		public.ResponseAny(ctx, err, bytes)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//GetCourse: 课程详情
func GetCourse(ctx *gin.Context) {
	var req basic.String
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name].(course.CourseService)
		list, err := courseService.CourseDetail(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//SaveCourse: 保存课程
func SaveCourse(ctx *gin.Context) {
	var req dto.CourseDto
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name].(course.CourseService)
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
		courseService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name].(course.CourseService)
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
		courseService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name].(course.CourseService)
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
		courseService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name].(course.CourseService)
		list, err := courseService.SaveCourseContent(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

//SortCourse: 课程排序
func SortCourse(ctx *gin.Context) {
	var req dto.SortDto
	if err := ctx.Bind(&req); err == nil {
		courseService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name].(course.CourseService)
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
		courseService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name].(course.CourseService)
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
		courseService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name].(course.CourseService)
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
		courseService := ctx.Keys[config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name].(course.CourseService)
		list, err := courseService.DeleteCourseFile(context.Background(), &req)
		public.ResponseAny(ctx, err, list)
	} else {
		public.ResponseError(ctx, public.NewBusinessException(public.VALID_PARM_ERROR))
	}
}

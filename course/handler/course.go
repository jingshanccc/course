package handler

import (
	"context"
	"gitee.com/jingshanccc/course/course/proto/dto"
	"gitee.com/jingshanccc/course/public"
	"gitee.com/jingshanccc/course/public/config"
	"gitee.com/jingshanccc/course/public/proto/basic"
	"gitee.com/jingshanccc/course/public/util"
	"github.com/micro/go-micro/v2/errors"
	"gorm.io/gorm"
)

//CourseList: get course page
func (c *CourseServiceHandler) CourseList(ctx context.Context, in *dto.CoursePageDto, out *dto.CoursePageDto) error {
	count, courseDtos, err := courseDao.List(in)
	if err != nil {
		return errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name, err.Error(), err.Code())
	}
	_ = util.CopyProperties(out, in)
	out.Rows = courseDtos
	out.Total = count
	return nil
}

//CarouselCourse: 轮播图课程
func (c *CourseServiceHandler) CarouselCourse(ctx context.Context, in *basic.String, out *dto.CourseDtoList) error {
	courseDtos, err := courseDao.CarouselCourse()
	if err != nil {
		return errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name, err.Error(), err.Code())
	}
	out.Rows = courseDtos
	return nil
}

//NewPublishCourse: 新上好课
func (c *CourseServiceHandler) NewPublishCourse(ctx context.Context, in *basic.String, out *dto.CourseDtoList) error {
	courseDtos, err := courseDao.NewPublish()
	if err != nil {
		return errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name, err.Error(), err.Code())
	}
	out.Rows = courseDtos
	return nil
}

//RelatedCourse: 新上好课
func (c *CourseServiceHandler) RelatedCourse(ctx context.Context, in *basic.String, out *dto.CourseDtoList) error {
	ccds, _ := courseCategoryDao.SelectByCourseId(in.Str)
	courseDtos, err := courseDao.SelectCourseByIds(courseCategoryDao.SelectCourseIds(ccds...), false)
	if err != nil {
		return errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name, err.Error(), err.Code())
	}
	if len(courseDtos) == 1 {
		courseDtos, _ = courseDao.NewPublish()
	}
	out.Rows = courseDtos
	return nil
}

//CategoryCourse: 分类搜索课程
func (c *CourseServiceHandler) CategoryCourse(ctx context.Context, in *basic.String, out *dto.CourseDtoList) error {
	// 通过 分类id 获取分类id和其父分类id 获取两个id 对应的所有课程id
	var ids []string
	var err *public.BusinessException
	var courseDtos []*dto.CourseDto
	if in.Str != "" {
		ids = courseCategoryDao.SelectCourseIds(in.Str)
		courseDtos, err = courseDao.SelectCourseByIds(ids, false)
	} else {
		courseDtos, err = courseDao.SelectCourseByIds(ids, true)
	}
	if err != nil {
		return errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name, err.Error(), err.Code())
	}
	out.Rows = courseDtos
	return nil
}

//DownloadCourseContent: 下载课程讲义
func (c *CourseServiceHandler) DownloadCourseContent(ctx context.Context, in *basic.String, out *basic.String) error {
	content, exception := courseDao.FindContent(in.Str)
	if exception != nil {
		return errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name, exception.Error(), exception.Code())
	}
	out.Str = util.TransToHtml(content.Content)
	return nil
}

//CourseDetail: 课程详情
func (c *CourseServiceHandler) CourseDetail(ctx context.Context, in *basic.String, out *dto.CourseDto) error {
	courseDb := courseDao.SelectByProperty("id", in.Str)
	if courseDb.Id != "" {
		_ = util.CopyProperties(out, courseDb)
		teacherDb, _ := teacherDao.SelectByProperty("id", courseDb.TeacherId)
		chapters, _ := chapterDao.SelectByProperty("course_id", courseDb.Id)
		sections, _ := sectionDao.SelectByProperty("course_id", courseDb.Id)
		out.Chapters = chapters
		out.Teacher = teacherDb
		out.Sections = sections
	}

	return nil
}

//SaveCourse: 保存/更新课程
func (c *CourseServiceHandler) SaveCourse(ctx context.Context, in *dto.CourseDto, out *dto.CourseDto) error {
	cd, err := courseDao.Save(in)
	if err != nil {
		return errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name, err.Error(), err.Code())
	}
	_ = util.CopyProperties(out, cd)
	return nil
}

//DeleteCourse: 删除课程
func (c *CourseServiceHandler) DeleteCourse(ctx context.Context, in *basic.StringList, out *basic.String) error {
	exception := courseDao.Delete(in.Rows)
	if exception != nil {
		return errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name, exception.Error(), exception.Code())
	}
	return nil
}

//SortCourse: 更新课程排序
func (c *CourseServiceHandler) SortCourse(ctx context.Context, in *dto.SortDto, out *basic.String) error {
	var exception *public.BusinessException
	var tx *gorm.DB
	var err error
	defer func() {
		if exception != nil {
			tx.Rollback()
			err = errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name, exception.Error(), exception.Code())
		} else {
			err = nil
		}
	}()
	tx = public.DB.Begin()
	exception = courseDao.UpdateSort(tx, in)
	if in.NewSort > in.OldSort {
		exception = courseDao.MoveSortsForward(tx, in)
	}
	if in.OldSort > in.NewSort {
		exception = courseDao.MoveSortsBackward(tx, in)
	}
	tx.Commit()
	return err
}

//FindCourseContent: 查找课程内容
func (c *CourseServiceHandler) FindCourseContent(ctx context.Context, in *basic.String, out *dto.CourseContentDto) error {
	content, exception := courseDao.FindContent(in.Str)
	if exception != nil {
		return errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name, exception.Error(), exception.Code())
	}
	out.Id = content.Id
	out.Content = content.Content
	return nil
}

//SaveCourseContent: 保存课程内容
func (c *CourseServiceHandler) SaveCourseContent(ctx context.Context, in *dto.CourseContentDto, out *basic.String) error {
	exception := courseDao.SaveContent(in)
	if exception != nil {
		return errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name, exception.Error(), exception.Code())
	}
	return nil
}

//ListCourseFile: 获取课程文件
func (c *CourseServiceHandler) ListCourseFile(ctx context.Context, in *basic.String, out *dto.CourseFileDtoList) error {
	list, exception := courseFileDao.List(in.Str)
	if exception != nil {
		return errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name, exception.Error(), exception.Code())
	}
	out.Rows = list
	return nil
}

//SaveCourseFile: 保存课程文件
func (c *CourseServiceHandler) SaveCourseFile(ctx context.Context, in *dto.CourseFileDto, out *dto.CourseFileDto) error {
	fileDto, exception := courseFileDao.Save(in)
	if exception != nil {
		return errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name, exception.Error(), exception.Code())
	}
	_ = util.CopyProperties(out, fileDto)
	return nil
}

//DeleteCourseFile: 删除课程文件
func (c *CourseServiceHandler) DeleteCourseFile(ctx context.Context, in *basic.String, out *basic.String) error {
	exception := courseFileDao.Delete(in.Str)
	if exception != nil {
		return errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name, exception.Error(), exception.Code())
	}
	return nil
}

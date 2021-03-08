package handler

import (
	"context"
	"course/config"
	"course/course-srv/proto/dto"
	"course/proto/basic"
	"course/public"
	"course/public/util"
	"github.com/micro/go-micro/v2/errors"
	"gorm.io/gorm"
)

//CourseList: get course page
func (c *CourseServiceHandler) CourseList(ctx context.Context, in *dto.CoursePageDto, out *dto.CoursePageDto) error {
	count, courseDtos, err := courseDao.List(in)
	if err != nil {
		return errors.New(config.CourseServiceName, err.Error(), err.Code())
	}
	_ = util.CopyProperties(out, in)
	out.Rows = courseDtos
	out.Total = count
	return nil
}

//CarouselCourse: get course page
func (c *CourseServiceHandler) CarouselCourse(ctx context.Context, in *basic.String, out *dto.CourseDtoList) error {
	courseDtos, err := courseDao.CarouselCourse()
	if err != nil {
		return errors.New(config.CourseServiceName, err.Error(), err.Code())
	}
	out.Rows = courseDtos
	return nil
}

//NewPublishCourse: get course page
func (c *CourseServiceHandler) NewPublishCourse(ctx context.Context, in *basic.String, out *dto.CourseDtoList) error {
	courseDtos, err := courseDao.NewPublish()
	if err != nil {
		return errors.New(config.CourseServiceName, err.Error(), err.Code())
	}
	out.Rows = courseDtos
	return nil
}

//SaveCourse: 保存/更新课程
func (c *CourseServiceHandler) SaveCourse(ctx context.Context, in *dto.CourseDto, out *dto.CourseDto) error {
	cd, err := courseDao.Save(in)
	if err != nil {
		return errors.New(config.CourseServiceName, err.Error(), err.Code())
	}
	_ = util.CopyProperties(out, cd)
	return nil
}

//DeleteCourse: 删除课程
func (c *CourseServiceHandler) DeleteCourse(ctx context.Context, in *basic.StringList, out *basic.String) error {
	exception := courseDao.Delete(in.Rows)
	if exception != nil {
		return errors.New(config.CourseServiceName, exception.Error(), exception.Code())
	}
	return nil
}

//ListCourseCategory: 获取课程所属的所有分类
func (c *CourseServiceHandler) ListCourseCategory(ctx context.Context, in *basic.String, out *dto.CourseCategoryDtoList) error {
	courseCategoryDtos, exception := courseCategoryDao.SelectByCourseId(in.Str)
	if exception != nil {
		return errors.New(config.CourseServiceName, exception.Error(), exception.Code())
	}
	out.Rows = courseCategoryDtos
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
			err = errors.New(config.CourseServiceName, exception.Error(), exception.Code())
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
		return errors.New(config.CourseServiceName, exception.Error(), exception.Code())
	}
	out.Id = content.Id
	out.Content = content.Content
	return nil
}

//SaveCourseContent: 保存课程内容
func (c *CourseServiceHandler) SaveCourseContent(ctx context.Context, in *dto.CourseContentDto, out *basic.String) error {
	exception := courseDao.SaveContent(in)
	if exception != nil {
		return errors.New(config.CourseServiceName, exception.Error(), exception.Code())
	}
	return nil
}

//ListCourseFile: 获取课程文件
func (c *CourseServiceHandler) ListCourseFile(ctx context.Context, in *basic.String, out *dto.CourseFileDtoList) error {
	list, exception := courseFileDao.List(in.Str)
	if exception != nil {
		return errors.New(config.CourseServiceName, exception.Error(), exception.Code())
	}
	out.Rows = list
	return nil
}

//SaveCourseFile: 保存课程文件
func (c *CourseServiceHandler) SaveCourseFile(ctx context.Context, in *dto.CourseFileDto, out *dto.CourseFileDto) error {
	fileDto, exception := courseFileDao.Save(in)
	if exception != nil {
		return errors.New(config.CourseServiceName, exception.Error(), exception.Code())
	}
	_ = util.CopyProperties(out, fileDto)
	return nil
}

//DeleteCourseFile: 删除课程文件
func (c *CourseServiceHandler) DeleteCourseFile(ctx context.Context, in *basic.String, out *basic.String) error {
	exception := courseFileDao.Delete(in.Str)
	if exception != nil {
		return errors.New(config.CourseServiceName, exception.Error(), exception.Code())
	}
	return nil
}

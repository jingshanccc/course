package handler

import (
	"context"
	"gitee.com/jingshanccc/course/course/proto/dto"
	"gitee.com/jingshanccc/course/public/config"
	"gitee.com/jingshanccc/course/public/proto/basic"
	"gitee.com/jingshanccc/course/public/util"
	"github.com/micro/go-micro/v2/errors"
)

func (c *CourseServiceHandler) ListSection(ctx context.Context, in *dto.SectionPageDto, out *dto.SectionPageDto) error {
	total, list, exception := sectionDao.List(in)
	if exception != nil {
		return errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name, exception.Error(), exception.Code())
	}
	_ = util.CopyProperties(out, in)
	out.Rows = list
	out.Total = total
	return nil
}

func (c *CourseServiceHandler) SaveSection(ctx context.Context, in *dto.SectionDto, out *dto.SectionDto) error {
	sectionDto, exception := sectionDao.Save(in)
	if exception != nil {
		return errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name, exception.Error(), exception.Code())
	}
	exception = courseDao.UpdateCourseDuration(in.CourseId)
	if exception != nil {
		return errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name, exception.Error(), exception.Code())
	}
	_ = util.CopyProperties(out, sectionDto)
	return nil
}

func (c *CourseServiceHandler) DeleteSection(ctx context.Context, in *basic.StringList, out *basic.String) error {
	exception := sectionDao.Delete(in.Rows)
	if exception != nil {
		return errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name, exception.Error(), exception.Code())
	}
	return nil
}

package handler

import (
	"context"
	"gitee.com/jingshanccc/course/course/proto/dto"
	"gitee.com/jingshanccc/course/public/config"
	"gitee.com/jingshanccc/course/public/proto/basic"
	"gitee.com/jingshanccc/course/public/util"
	"github.com/micro/go-micro/v2/errors"
)

func (c *CourseServiceHandler) ListChapter(ctx context.Context, in *dto.ChapterPageDto, out *dto.ChapterPageDto) error {
	total, list, exception := chapterDao.List(in)
	if exception != nil {
		return errors.New(config.Conf.Services["course"].Name, exception.Error(), exception.Code())
	}
	_ = util.CopyProperties(out, in)
	out.Rows = list
	out.Total = total
	return nil
}

func (c *CourseServiceHandler) AllChapter(ctx context.Context, in *basic.String, out *dto.ChapterDtoList) error {
	list, exception := chapterDao.All(in.Str)
	if exception != nil {
		return errors.New(config.Conf.Services["course"].Name, exception.Error(), exception.Code())
	}
	_ = util.CopyProperties(out, in)
	out.Rows = list
	return nil
}

func (c *CourseServiceHandler) SaveChapter(ctx context.Context, in *dto.ChapterDto, out *dto.ChapterDto) error {
	chapterDto, exception := chapterDao.Save(in)
	if exception != nil {
		return errors.New(config.Conf.Services["course"].Name, exception.Error(), exception.Code())
	}
	_ = util.CopyProperties(out, chapterDto)
	return nil
}

func (c *CourseServiceHandler) DeleteChapter(ctx context.Context, in *basic.StringList, out *basic.String) error {
	exception := chapterDao.Delete(in.Rows)
	if exception != nil {
		return errors.New(config.Conf.Services["course"].Name, exception.Error(), exception.Code())
	}
	return nil
}

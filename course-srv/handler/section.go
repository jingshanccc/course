package handler

import (
	"context"
	"course/course-srv/proto/dto"
	"course/proto/basic"
	"course/public"
	"course/public/util"
	"github.com/micro/go-micro/v2/errors"
)

func (c *CourseServiceHandler) ListSection(ctx context.Context, in *dto.SectionPageDto, out *dto.SectionPageDto) error {
	list, exception := sectionDao.List(in)
	if exception.Code() != int32(public.OK) {
		return errors.New(public.CourseServiceName, exception.Error(), exception.Code())
	}
	_ = util.CopyProperties(out, in)
	out.Rows = list
	return nil
}

func (c *CourseServiceHandler) SaveSection(ctx context.Context, in *dto.SectionDto, out *dto.SectionDto) error {
	sectionDto, exception := sectionDao.Save(in)
	if exception.Code() != int32(public.OK) {
		return errors.New(public.CourseServiceName, exception.Error(), exception.Code())
	}
	_ = util.CopyProperties(out, sectionDto)
	return nil
}

func (c *CourseServiceHandler) DeleteSection(ctx context.Context, in *basic.String, out *basic.String) error {
	exception := sectionDao.Delete(in.Str)
	if exception.Code() != int32(public.OK) {
		return errors.New(public.CourseServiceName, exception.Error(), exception.Code())
	}
	return nil
}

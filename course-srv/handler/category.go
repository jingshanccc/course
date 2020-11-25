package handler

import (
	"context"
	"course/course-srv/proto/dto"
	"course/proto/basic"
	"course/public"
	"course/public/util"
	"github.com/micro/go-micro/v2/errors"
)

//AllCategory: 获取所有分类
func (c *CourseServiceHandler) AllCategory(ctx context.Context, in *basic.String, out *dto.CategoryDtoList) error {
	dtos, err := categoryDao.All()
	if err.Code() != int32(public.OK) {
		return errors.New(public.CourseServiceName, err.Error(), err.Code())
	}
	out.Rows = dtos
	return nil
}

//SaveCategory: 保存分类
func (c *CourseServiceHandler) SaveCategory(ctx context.Context, in *dto.CategoryDto, out *dto.CategoryDto) error {
	cd, err := categoryDao.Save(in)
	if err.Code() != int32(public.OK) {
		return errors.New(public.CourseServiceName, err.Error(), err.Code())
	}
	_ = util.CopyProperties(out, cd)
	return nil
}

//DeleteCategory: 删除分类
func (c *CourseServiceHandler) DeleteCategory(ctx context.Context, in *basic.String, out *basic.String) error {
	exception := categoryDao.Delete(in.Str)
	if exception.Code() != int32(public.OK) {
		return errors.New(public.CourseServiceName, exception.Error(), exception.Code())
	}
	return nil
}

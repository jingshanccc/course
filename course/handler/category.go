package handler

import (
	"context"
	"gitee.com/jingshanccc/course/course/proto/dto"
	"gitee.com/jingshanccc/course/public"
	"gitee.com/jingshanccc/course/public/config"
	"gitee.com/jingshanccc/course/public/proto/basic"
	"gitee.com/jingshanccc/course/public/util"
	"github.com/micro/go-micro/v2/errors"
)

//ListCategory: 获取分类列表
func (c *CourseServiceHandler) ListCategory(ctx context.Context, in *dto.CategoryPageDto, out *dto.CategoryPageDto) error {
	count, list, exception := categoryDao.List(ctx, in)
	if exception != nil {
		return errors.New(config.Conf.Services["course"].Name, exception.Error(), exception.Code())
	}
	_ = util.CopyProperties(out, in)
	out.Total = count
	out.Rows = list
	return nil
}

//PrimaryCategory: 获取所有一级分类
func (c *CourseServiceHandler) PrimaryCategory(ctx context.Context, in *basic.String, out *dto.CategoryDtoList) error {
	dtos, err := categoryDao.PrimaryCategory()
	if err != nil {
		return errors.New(config.Conf.Services["course"].Name, err.Error(), err.Code())
	}
	out.Rows = dtos
	return nil
}

//AllCategory: 获取所有分类
func (c *CourseServiceHandler) AllCategory(ctx context.Context, in *basic.String, out *dto.CategoryDtoList) error {
	dtos, err := categoryDao.All()
	if err != nil {
		return errors.New(config.Conf.Services["course"].Name, err.Error(), err.Code())
	}
	out.Rows = dtos
	return nil
}

//SaveCategory: 保存分类
func (c *CourseServiceHandler) SaveCategory(ctx context.Context, in *dto.CategoryDto, out *dto.CategoryDto) error {
	cd, err := categoryDao.Save(in)
	if err != nil {
		return errors.New(config.Conf.Services["course"].Name, err.Error(), err.Code())
	}
	_ = util.CopyProperties(out, cd)
	return nil
}

//DeleteCategory: 删除分类
func (c *CourseServiceHandler) DeleteCategory(ctx context.Context, in *basic.StringList, out *basic.String) error {
	var exception *public.BusinessException
	if count := courseCategoryDao.CountByCategories(in.Rows); count > 0 {
		exception = public.BadRequestException("所选分类存在课程关联，请解除关联再试！")
	} else {
		exception = categoryDao.Delete(in.Rows)
	}
	if exception != nil {
		return errors.New(config.Conf.Services["course"].Name, exception.Error(), exception.Code())
	}
	return nil
}

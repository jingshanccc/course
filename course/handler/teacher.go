package handler

import (
	"context"
	"gitee.com/jingshanccc/course/course/proto/dto"
	"gitee.com/jingshanccc/course/public/config"
	"gitee.com/jingshanccc/course/public/proto/basic"
	"gitee.com/jingshanccc/course/public/util"
	"github.com/micro/go-micro/v2/errors"
)

//AllTeacher: 获取所有讲师
func (c *CourseServiceHandler) AllTeacher(ctx context.Context, in *basic.String, out *dto.TeacherDtoList) error {
	dtos, err := teacherDao.All()
	if err != nil {
		return errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name, err.Error(), err.Code())
	}
	out.Rows = dtos
	return nil
}

//SearchTeacher: 搜索讲师
func (c *CourseServiceHandler) SearchTeacher(ctx context.Context, in *basic.String, out *dto.TeacherDtoList) error {
	dtos, err := teacherDao.SearchByProperty("name", in.Str)
	if err != nil {
		return errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name, err.Error(), err.Code())
	}
	out.Rows = dtos
	return nil
}

//ListTeacher: get Teacher page
func (c *CourseServiceHandler) ListTeacher(ctx context.Context, in *dto.TeacherPageDto, out *dto.TeacherPageDto) error {
	count, list, exception := teacherDao.List(in)
	if exception != nil {
		return errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name, exception.Error(), exception.Code())
	}
	_ = util.CopyProperties(out, in)
	out.Rows = list
	out.Total = count
	return nil
}

//Save: 保存/更新讲师
func (c *CourseServiceHandler) SaveTeacher(ctx context.Context, in *dto.TeacherDto, out *dto.TeacherDto) error {
	TeacherDto, exception := teacherDao.Save(in)
	if exception != nil {
		return errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name, exception.Error(), exception.Code())
	}
	_ = util.CopyProperties(out, TeacherDto)
	return nil
}

// Delete: 删除讲师
func (c *CourseServiceHandler) DeleteTeacher(ctx context.Context, in *basic.StringList, out *basic.String) error {
	exception := teacherDao.Delete(in.Rows)
	if exception != nil {
		return errors.New(config.Conf.BasicConfig.BasicName+config.Conf.Services["course"].Name, exception.Error(), exception.Code())
	}
	return nil
}

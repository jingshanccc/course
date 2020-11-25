package handler

import (
	"context"
	"course/course-srv/proto/dto"
	"course/proto/basic"
	"course/public"
	"course/public/util"
	"github.com/micro/go-micro/v2/errors"
)

//AllTeacher: 获取所有讲师
func (c *CourseServiceHandler) AllTeacher(ctx context.Context, in *basic.String, out *dto.TeacherDtoList) error {
	dtos, err := teacherDao.All()
	if err.Code() != int32(public.OK) {
		return errors.New(public.CourseServiceName, err.Error(), err.Code())
	}
	out.Rows = dtos
	return nil
}

//ListTeacher: get Teacher page
func (c *CourseServiceHandler) ListTeacher(ctx context.Context, in *dto.TeacherPageDto, out *dto.TeacherPageDto) error {
	list, exception := teacherDao.List(in)
	if exception.Code() != int32(public.OK) {
		return errors.New(public.CourseServiceName, exception.Error(), exception.Code())
	}
	_ = util.CopyProperties(out, in)
	out.Rows = list
	return nil
}

//Save: 保存/更新讲师
func (c *CourseServiceHandler) SaveTeacher(ctx context.Context, in *dto.TeacherDto, out *dto.TeacherDto) error {
	TeacherDto, exception := teacherDao.Save(in)
	if exception.Code() != int32(public.OK) {
		return errors.New(public.CourseServiceName, exception.Error(), exception.Code())
	}
	_ = util.CopyProperties(out, TeacherDto)
	return nil
}

// Delete: 删除讲师
func (c *CourseServiceHandler) DeleteTeacher(ctx context.Context, in *basic.String, out *basic.String) error {
	exception := teacherDao.Delete(in.Str)
	if exception.Code() != int32(public.OK) {
		return errors.New(public.CourseServiceName, exception.Error(), exception.Code())
	}
	return nil
}

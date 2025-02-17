// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: gitee.com/jingshanccc/course/course/proto/course/course.proto

package course

import (
	fmt "fmt"
	dto "gitee.com/jingshanccc/course/course/proto/dto"
	basic "gitee.com/jingshanccc/course/public/proto/basic"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for CourseService service

func NewCourseServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for CourseService service

type CourseService interface {
	//  ------------  管理端 --------------------
	//Course
	CourseList(ctx context.Context, in *dto.CoursePageDto, opts ...client.CallOption) (*dto.CoursePageDto, error)
	SaveCourse(ctx context.Context, in *dto.CourseDto, opts ...client.CallOption) (*dto.CourseDto, error)
	DeleteCourse(ctx context.Context, in *basic.StringList, opts ...client.CallOption) (*basic.String, error)
	SortCourse(ctx context.Context, in *dto.SortDto, opts ...client.CallOption) (*basic.String, error)
	FindCourseContent(ctx context.Context, in *basic.String, opts ...client.CallOption) (*dto.CourseContentDto, error)
	SaveCourseContent(ctx context.Context, in *dto.CourseContentDto, opts ...client.CallOption) (*basic.String, error)
	//Category
	ListCategory(ctx context.Context, in *dto.CategoryPageDto, opts ...client.CallOption) (*dto.CategoryPageDto, error)
	PrimaryCategory(ctx context.Context, in *basic.String, opts ...client.CallOption) (*dto.CategoryDtoList, error)
	SaveCategory(ctx context.Context, in *dto.CategoryDto, opts ...client.CallOption) (*dto.CategoryDto, error)
	DeleteCategory(ctx context.Context, in *basic.StringList, opts ...client.CallOption) (*basic.String, error)
	//Chapter
	ListChapter(ctx context.Context, in *dto.ChapterPageDto, opts ...client.CallOption) (*dto.ChapterPageDto, error)
	AllChapter(ctx context.Context, in *basic.String, opts ...client.CallOption) (*dto.ChapterDtoList, error)
	SaveChapter(ctx context.Context, in *dto.ChapterDto, opts ...client.CallOption) (*dto.ChapterDto, error)
	DeleteChapter(ctx context.Context, in *basic.StringList, opts ...client.CallOption) (*basic.String, error)
	//Section
	ListSection(ctx context.Context, in *dto.SectionPageDto, opts ...client.CallOption) (*dto.SectionPageDto, error)
	SaveSection(ctx context.Context, in *dto.SectionDto, opts ...client.CallOption) (*dto.SectionDto, error)
	DeleteSection(ctx context.Context, in *basic.StringList, opts ...client.CallOption) (*basic.String, error)
	//CourseFile
	ListCourseFile(ctx context.Context, in *basic.String, opts ...client.CallOption) (*dto.CourseFileDtoList, error)
	SaveCourseFile(ctx context.Context, in *dto.CourseFileDto, opts ...client.CallOption) (*dto.CourseFileDto, error)
	DeleteCourseFile(ctx context.Context, in *basic.String, opts ...client.CallOption) (*basic.String, error)
	//Teacher
	ListTeacher(ctx context.Context, in *dto.TeacherPageDto, opts ...client.CallOption) (*dto.TeacherPageDto, error)
	AllTeacher(ctx context.Context, in *basic.String, opts ...client.CallOption) (*dto.TeacherDtoList, error)
	SearchTeacher(ctx context.Context, in *basic.String, opts ...client.CallOption) (*dto.TeacherDtoList, error)
	SaveTeacher(ctx context.Context, in *dto.TeacherDto, opts ...client.CallOption) (*dto.TeacherDto, error)
	DeleteTeacher(ctx context.Context, in *basic.StringList, opts ...client.CallOption) (*basic.String, error)
	//  ------------  平台端 --------------------
	// 分类标签云
	AllCategory(ctx context.Context, in *basic.String, opts ...client.CallOption) (*dto.CategoryDtoList, error)
	// 轮播图课程
	CarouselCourse(ctx context.Context, in *basic.String, opts ...client.CallOption) (*dto.CourseDtoList, error)
	// 新上好课
	NewPublishCourse(ctx context.Context, in *basic.String, opts ...client.CallOption) (*dto.CourseDtoList, error)
	// 分类搜索课程
	CategoryCourse(ctx context.Context, in *basic.String, opts ...client.CallOption) (*dto.CourseDtoList, error)
	// 课程详情
	CourseDetail(ctx context.Context, in *basic.String, opts ...client.CallOption) (*dto.CourseDto, error)
	// 相关课程
	RelatedCourse(ctx context.Context, in *basic.String, opts ...client.CallOption) (*dto.CourseDtoList, error)
	// 下载课程讲义
	DownloadCourseContent(ctx context.Context, in *basic.String, opts ...client.CallOption) (*basic.String, error)
	// 我的课程
	MyCourse(ctx context.Context, in *basic.String, opts ...client.CallOption) (*dto.CourseDtoList, error)
	// 添加到我的课程
	AddToMyCourse(ctx context.Context, in *dto.MemberCourseDto, opts ...client.CallOption) (*basic.String, error)
	// 获取课程学习进度
	CourseInfo(ctx context.Context, in *basic.StringList, opts ...client.CallOption) (*basic.String, error)
	// 保存课程学习进度
	SaveLearnInfo(ctx context.Context, in *basic.StringList, opts ...client.CallOption) (*basic.String, error)
}

type courseService struct {
	c    client.Client
	name string
}

func NewCourseService(name string, c client.Client) CourseService {
	return &courseService{
		c:    c,
		name: name,
	}
}

func (c *courseService) CourseList(ctx context.Context, in *dto.CoursePageDto, opts ...client.CallOption) (*dto.CoursePageDto, error) {
	req := c.c.NewRequest(c.name, "CourseService.CourseList", in)
	out := new(dto.CoursePageDto)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseService) SaveCourse(ctx context.Context, in *dto.CourseDto, opts ...client.CallOption) (*dto.CourseDto, error) {
	req := c.c.NewRequest(c.name, "CourseService.SaveCourse", in)
	out := new(dto.CourseDto)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseService) DeleteCourse(ctx context.Context, in *basic.StringList, opts ...client.CallOption) (*basic.String, error) {
	req := c.c.NewRequest(c.name, "CourseService.DeleteCourse", in)
	out := new(basic.String)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseService) SortCourse(ctx context.Context, in *dto.SortDto, opts ...client.CallOption) (*basic.String, error) {
	req := c.c.NewRequest(c.name, "CourseService.SortCourse", in)
	out := new(basic.String)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseService) FindCourseContent(ctx context.Context, in *basic.String, opts ...client.CallOption) (*dto.CourseContentDto, error) {
	req := c.c.NewRequest(c.name, "CourseService.FindCourseContent", in)
	out := new(dto.CourseContentDto)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseService) SaveCourseContent(ctx context.Context, in *dto.CourseContentDto, opts ...client.CallOption) (*basic.String, error) {
	req := c.c.NewRequest(c.name, "CourseService.SaveCourseContent", in)
	out := new(basic.String)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseService) ListCategory(ctx context.Context, in *dto.CategoryPageDto, opts ...client.CallOption) (*dto.CategoryPageDto, error) {
	req := c.c.NewRequest(c.name, "CourseService.ListCategory", in)
	out := new(dto.CategoryPageDto)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseService) PrimaryCategory(ctx context.Context, in *basic.String, opts ...client.CallOption) (*dto.CategoryDtoList, error) {
	req := c.c.NewRequest(c.name, "CourseService.PrimaryCategory", in)
	out := new(dto.CategoryDtoList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseService) SaveCategory(ctx context.Context, in *dto.CategoryDto, opts ...client.CallOption) (*dto.CategoryDto, error) {
	req := c.c.NewRequest(c.name, "CourseService.SaveCategory", in)
	out := new(dto.CategoryDto)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseService) DeleteCategory(ctx context.Context, in *basic.StringList, opts ...client.CallOption) (*basic.String, error) {
	req := c.c.NewRequest(c.name, "CourseService.DeleteCategory", in)
	out := new(basic.String)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseService) ListChapter(ctx context.Context, in *dto.ChapterPageDto, opts ...client.CallOption) (*dto.ChapterPageDto, error) {
	req := c.c.NewRequest(c.name, "CourseService.ListChapter", in)
	out := new(dto.ChapterPageDto)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseService) AllChapter(ctx context.Context, in *basic.String, opts ...client.CallOption) (*dto.ChapterDtoList, error) {
	req := c.c.NewRequest(c.name, "CourseService.AllChapter", in)
	out := new(dto.ChapterDtoList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseService) SaveChapter(ctx context.Context, in *dto.ChapterDto, opts ...client.CallOption) (*dto.ChapterDto, error) {
	req := c.c.NewRequest(c.name, "CourseService.SaveChapter", in)
	out := new(dto.ChapterDto)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseService) DeleteChapter(ctx context.Context, in *basic.StringList, opts ...client.CallOption) (*basic.String, error) {
	req := c.c.NewRequest(c.name, "CourseService.DeleteChapter", in)
	out := new(basic.String)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseService) ListSection(ctx context.Context, in *dto.SectionPageDto, opts ...client.CallOption) (*dto.SectionPageDto, error) {
	req := c.c.NewRequest(c.name, "CourseService.ListSection", in)
	out := new(dto.SectionPageDto)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseService) SaveSection(ctx context.Context, in *dto.SectionDto, opts ...client.CallOption) (*dto.SectionDto, error) {
	req := c.c.NewRequest(c.name, "CourseService.SaveSection", in)
	out := new(dto.SectionDto)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseService) DeleteSection(ctx context.Context, in *basic.StringList, opts ...client.CallOption) (*basic.String, error) {
	req := c.c.NewRequest(c.name, "CourseService.DeleteSection", in)
	out := new(basic.String)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseService) ListCourseFile(ctx context.Context, in *basic.String, opts ...client.CallOption) (*dto.CourseFileDtoList, error) {
	req := c.c.NewRequest(c.name, "CourseService.ListCourseFile", in)
	out := new(dto.CourseFileDtoList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseService) SaveCourseFile(ctx context.Context, in *dto.CourseFileDto, opts ...client.CallOption) (*dto.CourseFileDto, error) {
	req := c.c.NewRequest(c.name, "CourseService.SaveCourseFile", in)
	out := new(dto.CourseFileDto)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseService) DeleteCourseFile(ctx context.Context, in *basic.String, opts ...client.CallOption) (*basic.String, error) {
	req := c.c.NewRequest(c.name, "CourseService.DeleteCourseFile", in)
	out := new(basic.String)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseService) ListTeacher(ctx context.Context, in *dto.TeacherPageDto, opts ...client.CallOption) (*dto.TeacherPageDto, error) {
	req := c.c.NewRequest(c.name, "CourseService.ListTeacher", in)
	out := new(dto.TeacherPageDto)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseService) AllTeacher(ctx context.Context, in *basic.String, opts ...client.CallOption) (*dto.TeacherDtoList, error) {
	req := c.c.NewRequest(c.name, "CourseService.AllTeacher", in)
	out := new(dto.TeacherDtoList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseService) SearchTeacher(ctx context.Context, in *basic.String, opts ...client.CallOption) (*dto.TeacherDtoList, error) {
	req := c.c.NewRequest(c.name, "CourseService.SearchTeacher", in)
	out := new(dto.TeacherDtoList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseService) SaveTeacher(ctx context.Context, in *dto.TeacherDto, opts ...client.CallOption) (*dto.TeacherDto, error) {
	req := c.c.NewRequest(c.name, "CourseService.SaveTeacher", in)
	out := new(dto.TeacherDto)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseService) DeleteTeacher(ctx context.Context, in *basic.StringList, opts ...client.CallOption) (*basic.String, error) {
	req := c.c.NewRequest(c.name, "CourseService.DeleteTeacher", in)
	out := new(basic.String)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseService) AllCategory(ctx context.Context, in *basic.String, opts ...client.CallOption) (*dto.CategoryDtoList, error) {
	req := c.c.NewRequest(c.name, "CourseService.AllCategory", in)
	out := new(dto.CategoryDtoList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseService) CarouselCourse(ctx context.Context, in *basic.String, opts ...client.CallOption) (*dto.CourseDtoList, error) {
	req := c.c.NewRequest(c.name, "CourseService.CarouselCourse", in)
	out := new(dto.CourseDtoList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseService) NewPublishCourse(ctx context.Context, in *basic.String, opts ...client.CallOption) (*dto.CourseDtoList, error) {
	req := c.c.NewRequest(c.name, "CourseService.NewPublishCourse", in)
	out := new(dto.CourseDtoList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseService) CategoryCourse(ctx context.Context, in *basic.String, opts ...client.CallOption) (*dto.CourseDtoList, error) {
	req := c.c.NewRequest(c.name, "CourseService.CategoryCourse", in)
	out := new(dto.CourseDtoList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseService) CourseDetail(ctx context.Context, in *basic.String, opts ...client.CallOption) (*dto.CourseDto, error) {
	req := c.c.NewRequest(c.name, "CourseService.CourseDetail", in)
	out := new(dto.CourseDto)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseService) RelatedCourse(ctx context.Context, in *basic.String, opts ...client.CallOption) (*dto.CourseDtoList, error) {
	req := c.c.NewRequest(c.name, "CourseService.RelatedCourse", in)
	out := new(dto.CourseDtoList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseService) DownloadCourseContent(ctx context.Context, in *basic.String, opts ...client.CallOption) (*basic.String, error) {
	req := c.c.NewRequest(c.name, "CourseService.DownloadCourseContent", in)
	out := new(basic.String)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseService) MyCourse(ctx context.Context, in *basic.String, opts ...client.CallOption) (*dto.CourseDtoList, error) {
	req := c.c.NewRequest(c.name, "CourseService.MyCourse", in)
	out := new(dto.CourseDtoList)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseService) AddToMyCourse(ctx context.Context, in *dto.MemberCourseDto, opts ...client.CallOption) (*basic.String, error) {
	req := c.c.NewRequest(c.name, "CourseService.AddToMyCourse", in)
	out := new(basic.String)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseService) CourseInfo(ctx context.Context, in *basic.StringList, opts ...client.CallOption) (*basic.String, error) {
	req := c.c.NewRequest(c.name, "CourseService.CourseInfo", in)
	out := new(basic.String)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *courseService) SaveLearnInfo(ctx context.Context, in *basic.StringList, opts ...client.CallOption) (*basic.String, error) {
	req := c.c.NewRequest(c.name, "CourseService.SaveLearnInfo", in)
	out := new(basic.String)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CourseService service

type CourseServiceHandler interface {
	//  ------------  管理端 --------------------
	//Course
	CourseList(context.Context, *dto.CoursePageDto, *dto.CoursePageDto) error
	SaveCourse(context.Context, *dto.CourseDto, *dto.CourseDto) error
	DeleteCourse(context.Context, *basic.StringList, *basic.String) error
	SortCourse(context.Context, *dto.SortDto, *basic.String) error
	FindCourseContent(context.Context, *basic.String, *dto.CourseContentDto) error
	SaveCourseContent(context.Context, *dto.CourseContentDto, *basic.String) error
	//Category
	ListCategory(context.Context, *dto.CategoryPageDto, *dto.CategoryPageDto) error
	PrimaryCategory(context.Context, *basic.String, *dto.CategoryDtoList) error
	SaveCategory(context.Context, *dto.CategoryDto, *dto.CategoryDto) error
	DeleteCategory(context.Context, *basic.StringList, *basic.String) error
	//Chapter
	ListChapter(context.Context, *dto.ChapterPageDto, *dto.ChapterPageDto) error
	AllChapter(context.Context, *basic.String, *dto.ChapterDtoList) error
	SaveChapter(context.Context, *dto.ChapterDto, *dto.ChapterDto) error
	DeleteChapter(context.Context, *basic.StringList, *basic.String) error
	//Section
	ListSection(context.Context, *dto.SectionPageDto, *dto.SectionPageDto) error
	SaveSection(context.Context, *dto.SectionDto, *dto.SectionDto) error
	DeleteSection(context.Context, *basic.StringList, *basic.String) error
	//CourseFile
	ListCourseFile(context.Context, *basic.String, *dto.CourseFileDtoList) error
	SaveCourseFile(context.Context, *dto.CourseFileDto, *dto.CourseFileDto) error
	DeleteCourseFile(context.Context, *basic.String, *basic.String) error
	//Teacher
	ListTeacher(context.Context, *dto.TeacherPageDto, *dto.TeacherPageDto) error
	AllTeacher(context.Context, *basic.String, *dto.TeacherDtoList) error
	SearchTeacher(context.Context, *basic.String, *dto.TeacherDtoList) error
	SaveTeacher(context.Context, *dto.TeacherDto, *dto.TeacherDto) error
	DeleteTeacher(context.Context, *basic.StringList, *basic.String) error
	//  ------------  平台端 --------------------
	// 分类标签云
	AllCategory(context.Context, *basic.String, *dto.CategoryDtoList) error
	// 轮播图课程
	CarouselCourse(context.Context, *basic.String, *dto.CourseDtoList) error
	// 新上好课
	NewPublishCourse(context.Context, *basic.String, *dto.CourseDtoList) error
	// 分类搜索课程
	CategoryCourse(context.Context, *basic.String, *dto.CourseDtoList) error
	// 课程详情
	CourseDetail(context.Context, *basic.String, *dto.CourseDto) error
	// 相关课程
	RelatedCourse(context.Context, *basic.String, *dto.CourseDtoList) error
	// 下载课程讲义
	DownloadCourseContent(context.Context, *basic.String, *basic.String) error
	// 我的课程
	MyCourse(context.Context, *basic.String, *dto.CourseDtoList) error
	// 添加到我的课程
	AddToMyCourse(context.Context, *dto.MemberCourseDto, *basic.String) error
	// 获取课程学习进度
	CourseInfo(context.Context, *basic.StringList, *basic.String) error
	// 保存课程学习进度
	SaveLearnInfo(context.Context, *basic.StringList, *basic.String) error
}

func RegisterCourseServiceHandler(s server.Server, hdlr CourseServiceHandler, opts ...server.HandlerOption) error {
	type courseService interface {
		CourseList(ctx context.Context, in *dto.CoursePageDto, out *dto.CoursePageDto) error
		SaveCourse(ctx context.Context, in *dto.CourseDto, out *dto.CourseDto) error
		DeleteCourse(ctx context.Context, in *basic.StringList, out *basic.String) error
		SortCourse(ctx context.Context, in *dto.SortDto, out *basic.String) error
		FindCourseContent(ctx context.Context, in *basic.String, out *dto.CourseContentDto) error
		SaveCourseContent(ctx context.Context, in *dto.CourseContentDto, out *basic.String) error
		ListCategory(ctx context.Context, in *dto.CategoryPageDto, out *dto.CategoryPageDto) error
		PrimaryCategory(ctx context.Context, in *basic.String, out *dto.CategoryDtoList) error
		SaveCategory(ctx context.Context, in *dto.CategoryDto, out *dto.CategoryDto) error
		DeleteCategory(ctx context.Context, in *basic.StringList, out *basic.String) error
		ListChapter(ctx context.Context, in *dto.ChapterPageDto, out *dto.ChapterPageDto) error
		AllChapter(ctx context.Context, in *basic.String, out *dto.ChapterDtoList) error
		SaveChapter(ctx context.Context, in *dto.ChapterDto, out *dto.ChapterDto) error
		DeleteChapter(ctx context.Context, in *basic.StringList, out *basic.String) error
		ListSection(ctx context.Context, in *dto.SectionPageDto, out *dto.SectionPageDto) error
		SaveSection(ctx context.Context, in *dto.SectionDto, out *dto.SectionDto) error
		DeleteSection(ctx context.Context, in *basic.StringList, out *basic.String) error
		ListCourseFile(ctx context.Context, in *basic.String, out *dto.CourseFileDtoList) error
		SaveCourseFile(ctx context.Context, in *dto.CourseFileDto, out *dto.CourseFileDto) error
		DeleteCourseFile(ctx context.Context, in *basic.String, out *basic.String) error
		ListTeacher(ctx context.Context, in *dto.TeacherPageDto, out *dto.TeacherPageDto) error
		AllTeacher(ctx context.Context, in *basic.String, out *dto.TeacherDtoList) error
		SearchTeacher(ctx context.Context, in *basic.String, out *dto.TeacherDtoList) error
		SaveTeacher(ctx context.Context, in *dto.TeacherDto, out *dto.TeacherDto) error
		DeleteTeacher(ctx context.Context, in *basic.StringList, out *basic.String) error
		AllCategory(ctx context.Context, in *basic.String, out *dto.CategoryDtoList) error
		CarouselCourse(ctx context.Context, in *basic.String, out *dto.CourseDtoList) error
		NewPublishCourse(ctx context.Context, in *basic.String, out *dto.CourseDtoList) error
		CategoryCourse(ctx context.Context, in *basic.String, out *dto.CourseDtoList) error
		CourseDetail(ctx context.Context, in *basic.String, out *dto.CourseDto) error
		RelatedCourse(ctx context.Context, in *basic.String, out *dto.CourseDtoList) error
		DownloadCourseContent(ctx context.Context, in *basic.String, out *basic.String) error
		MyCourse(ctx context.Context, in *basic.String, out *dto.CourseDtoList) error
		AddToMyCourse(ctx context.Context, in *dto.MemberCourseDto, out *basic.String) error
		CourseInfo(ctx context.Context, in *basic.StringList, out *basic.String) error
		SaveLearnInfo(ctx context.Context, in *basic.StringList, out *basic.String) error
	}
	type CourseService struct {
		courseService
	}
	h := &courseServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&CourseService{h}, opts...))
}

type courseServiceHandler struct {
	CourseServiceHandler
}

func (h *courseServiceHandler) CourseList(ctx context.Context, in *dto.CoursePageDto, out *dto.CoursePageDto) error {
	return h.CourseServiceHandler.CourseList(ctx, in, out)
}

func (h *courseServiceHandler) SaveCourse(ctx context.Context, in *dto.CourseDto, out *dto.CourseDto) error {
	return h.CourseServiceHandler.SaveCourse(ctx, in, out)
}

func (h *courseServiceHandler) DeleteCourse(ctx context.Context, in *basic.StringList, out *basic.String) error {
	return h.CourseServiceHandler.DeleteCourse(ctx, in, out)
}

func (h *courseServiceHandler) SortCourse(ctx context.Context, in *dto.SortDto, out *basic.String) error {
	return h.CourseServiceHandler.SortCourse(ctx, in, out)
}

func (h *courseServiceHandler) FindCourseContent(ctx context.Context, in *basic.String, out *dto.CourseContentDto) error {
	return h.CourseServiceHandler.FindCourseContent(ctx, in, out)
}

func (h *courseServiceHandler) SaveCourseContent(ctx context.Context, in *dto.CourseContentDto, out *basic.String) error {
	return h.CourseServiceHandler.SaveCourseContent(ctx, in, out)
}

func (h *courseServiceHandler) ListCategory(ctx context.Context, in *dto.CategoryPageDto, out *dto.CategoryPageDto) error {
	return h.CourseServiceHandler.ListCategory(ctx, in, out)
}

func (h *courseServiceHandler) PrimaryCategory(ctx context.Context, in *basic.String, out *dto.CategoryDtoList) error {
	return h.CourseServiceHandler.PrimaryCategory(ctx, in, out)
}

func (h *courseServiceHandler) SaveCategory(ctx context.Context, in *dto.CategoryDto, out *dto.CategoryDto) error {
	return h.CourseServiceHandler.SaveCategory(ctx, in, out)
}

func (h *courseServiceHandler) DeleteCategory(ctx context.Context, in *basic.StringList, out *basic.String) error {
	return h.CourseServiceHandler.DeleteCategory(ctx, in, out)
}

func (h *courseServiceHandler) ListChapter(ctx context.Context, in *dto.ChapterPageDto, out *dto.ChapterPageDto) error {
	return h.CourseServiceHandler.ListChapter(ctx, in, out)
}

func (h *courseServiceHandler) AllChapter(ctx context.Context, in *basic.String, out *dto.ChapterDtoList) error {
	return h.CourseServiceHandler.AllChapter(ctx, in, out)
}

func (h *courseServiceHandler) SaveChapter(ctx context.Context, in *dto.ChapterDto, out *dto.ChapterDto) error {
	return h.CourseServiceHandler.SaveChapter(ctx, in, out)
}

func (h *courseServiceHandler) DeleteChapter(ctx context.Context, in *basic.StringList, out *basic.String) error {
	return h.CourseServiceHandler.DeleteChapter(ctx, in, out)
}

func (h *courseServiceHandler) ListSection(ctx context.Context, in *dto.SectionPageDto, out *dto.SectionPageDto) error {
	return h.CourseServiceHandler.ListSection(ctx, in, out)
}

func (h *courseServiceHandler) SaveSection(ctx context.Context, in *dto.SectionDto, out *dto.SectionDto) error {
	return h.CourseServiceHandler.SaveSection(ctx, in, out)
}

func (h *courseServiceHandler) DeleteSection(ctx context.Context, in *basic.StringList, out *basic.String) error {
	return h.CourseServiceHandler.DeleteSection(ctx, in, out)
}

func (h *courseServiceHandler) ListCourseFile(ctx context.Context, in *basic.String, out *dto.CourseFileDtoList) error {
	return h.CourseServiceHandler.ListCourseFile(ctx, in, out)
}

func (h *courseServiceHandler) SaveCourseFile(ctx context.Context, in *dto.CourseFileDto, out *dto.CourseFileDto) error {
	return h.CourseServiceHandler.SaveCourseFile(ctx, in, out)
}

func (h *courseServiceHandler) DeleteCourseFile(ctx context.Context, in *basic.String, out *basic.String) error {
	return h.CourseServiceHandler.DeleteCourseFile(ctx, in, out)
}

func (h *courseServiceHandler) ListTeacher(ctx context.Context, in *dto.TeacherPageDto, out *dto.TeacherPageDto) error {
	return h.CourseServiceHandler.ListTeacher(ctx, in, out)
}

func (h *courseServiceHandler) AllTeacher(ctx context.Context, in *basic.String, out *dto.TeacherDtoList) error {
	return h.CourseServiceHandler.AllTeacher(ctx, in, out)
}

func (h *courseServiceHandler) SearchTeacher(ctx context.Context, in *basic.String, out *dto.TeacherDtoList) error {
	return h.CourseServiceHandler.SearchTeacher(ctx, in, out)
}

func (h *courseServiceHandler) SaveTeacher(ctx context.Context, in *dto.TeacherDto, out *dto.TeacherDto) error {
	return h.CourseServiceHandler.SaveTeacher(ctx, in, out)
}

func (h *courseServiceHandler) DeleteTeacher(ctx context.Context, in *basic.StringList, out *basic.String) error {
	return h.CourseServiceHandler.DeleteTeacher(ctx, in, out)
}

func (h *courseServiceHandler) AllCategory(ctx context.Context, in *basic.String, out *dto.CategoryDtoList) error {
	return h.CourseServiceHandler.AllCategory(ctx, in, out)
}

func (h *courseServiceHandler) CarouselCourse(ctx context.Context, in *basic.String, out *dto.CourseDtoList) error {
	return h.CourseServiceHandler.CarouselCourse(ctx, in, out)
}

func (h *courseServiceHandler) NewPublishCourse(ctx context.Context, in *basic.String, out *dto.CourseDtoList) error {
	return h.CourseServiceHandler.NewPublishCourse(ctx, in, out)
}

func (h *courseServiceHandler) CategoryCourse(ctx context.Context, in *basic.String, out *dto.CourseDtoList) error {
	return h.CourseServiceHandler.CategoryCourse(ctx, in, out)
}

func (h *courseServiceHandler) CourseDetail(ctx context.Context, in *basic.String, out *dto.CourseDto) error {
	return h.CourseServiceHandler.CourseDetail(ctx, in, out)
}

func (h *courseServiceHandler) RelatedCourse(ctx context.Context, in *basic.String, out *dto.CourseDtoList) error {
	return h.CourseServiceHandler.RelatedCourse(ctx, in, out)
}

func (h *courseServiceHandler) DownloadCourseContent(ctx context.Context, in *basic.String, out *basic.String) error {
	return h.CourseServiceHandler.DownloadCourseContent(ctx, in, out)
}

func (h *courseServiceHandler) MyCourse(ctx context.Context, in *basic.String, out *dto.CourseDtoList) error {
	return h.CourseServiceHandler.MyCourse(ctx, in, out)
}

func (h *courseServiceHandler) AddToMyCourse(ctx context.Context, in *dto.MemberCourseDto, out *basic.String) error {
	return h.CourseServiceHandler.AddToMyCourse(ctx, in, out)
}

func (h *courseServiceHandler) CourseInfo(ctx context.Context, in *basic.StringList, out *basic.String) error {
	return h.CourseServiceHandler.CourseInfo(ctx, in, out)
}

func (h *courseServiceHandler) SaveLearnInfo(ctx context.Context, in *basic.StringList, out *basic.String) error {
	return h.CourseServiceHandler.SaveLearnInfo(ctx, in, out)
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: course/course-srv/proto/course/course.proto

package course

import (
	dto "course/course-srv/proto/dto"
	basic "course/proto/basic"
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

var File_course_course_srv_proto_course_course_proto protoreflect.FileDescriptor

var file_course_course_srv_proto_course_course_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2d,
	0x73, 0x72, 0x76, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x1a, 0x1e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x2d, 0x73, 0x72, 0x76, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64,
	0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32,
	0xbc, 0x04, 0x0a, 0x0d, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x15, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x50,
	0x61, 0x67, 0x65, 0x44, 0x74, 0x6f, 0x1a, 0x15, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x50, 0x61, 0x67, 0x65, 0x44, 0x74, 0x6f, 0x12, 0x32, 0x0a,
	0x0a, 0x53, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x11, 0x2e, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x44, 0x74, 0x6f, 0x1a, 0x11,
	0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x44, 0x74,
	0x6f, 0x12, 0x2c, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x12, 0x0d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x1a, 0x0d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12,
	0x42, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x43, 0x61, 0x74,
	0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x1a, 0x1d, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x43, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x44, 0x74, 0x6f, 0x4c,
	0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x0a, 0x53, 0x6f, 0x72, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x12, 0x0f, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x53, 0x6f, 0x72, 0x74, 0x44,
	0x74, 0x6f, 0x1a, 0x0d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x12, 0x3c, 0x0a, 0x11, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x0d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x43,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x44, 0x74, 0x6f, 0x12,
	0x3c, 0x0a, 0x11, 0x53, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x43, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x44, 0x74, 0x6f, 0x1a, 0x0d,
	0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x35, 0x0a,
	0x0b, 0x41, 0x6c, 0x6c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0d, 0x2e, 0x62,
	0x61, 0x73, 0x69, 0x63, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x17, 0x2e, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x44, 0x74, 0x6f,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0c, 0x53, 0x61, 0x76, 0x65, 0x43, 0x61, 0x74, 0x65,
	0x67, 0x6f, 0x72, 0x79, 0x12, 0x13, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x44, 0x74, 0x6f, 0x1a, 0x13, 0x2e, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x44, 0x74, 0x6f, 0x12, 0x2e,
	0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x12, 0x0d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a,
	0x0d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x42, 0x20,
	0x5a, 0x1e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2d,
	0x73, 0x72, 0x76, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_course_course_srv_proto_course_course_proto_goTypes = []interface{}{
	(*dto.CoursePageDto)(nil),         // 0: course.CoursePageDto
	(*dto.CourseDto)(nil),             // 1: course.CourseDto
	(*basic.String)(nil),              // 2: basic.String
	(*dto.SortDto)(nil),               // 3: course.SortDto
	(*dto.CourseContentDto)(nil),      // 4: course.CourseContentDto
	(*dto.CategoryDto)(nil),           // 5: course.CategoryDto
	(*dto.CourseCategoryDtoList)(nil), // 6: course.CourseCategoryDtoList
	(*dto.CategoryDtoList)(nil),       // 7: course.CategoryDtoList
}
var file_course_course_srv_proto_course_course_proto_depIdxs = []int32{
	0,  // 0: course.CourseService.CourseList:input_type -> course.CoursePageDto
	1,  // 1: course.CourseService.SaveCourse:input_type -> course.CourseDto
	2,  // 2: course.CourseService.DeleteCourse:input_type -> basic.String
	2,  // 3: course.CourseService.ListCourseCategory:input_type -> basic.String
	3,  // 4: course.CourseService.SortCourse:input_type -> course.SortDto
	2,  // 5: course.CourseService.FindCourseContent:input_type -> basic.String
	4,  // 6: course.CourseService.SaveCourseContent:input_type -> course.CourseContentDto
	2,  // 7: course.CourseService.AllCategory:input_type -> basic.String
	5,  // 8: course.CourseService.SaveCategory:input_type -> course.CategoryDto
	2,  // 9: course.CourseService.DeleteCategory:input_type -> basic.String
	0,  // 10: course.CourseService.CourseList:output_type -> course.CoursePageDto
	1,  // 11: course.CourseService.SaveCourse:output_type -> course.CourseDto
	2,  // 12: course.CourseService.DeleteCourse:output_type -> basic.String
	6,  // 13: course.CourseService.ListCourseCategory:output_type -> course.CourseCategoryDtoList
	2,  // 14: course.CourseService.SortCourse:output_type -> basic.String
	4,  // 15: course.CourseService.FindCourseContent:output_type -> course.CourseContentDto
	2,  // 16: course.CourseService.SaveCourseContent:output_type -> basic.String
	7,  // 17: course.CourseService.AllCategory:output_type -> course.CategoryDtoList
	5,  // 18: course.CourseService.SaveCategory:output_type -> course.CategoryDto
	2,  // 19: course.CourseService.DeleteCategory:output_type -> basic.String
	10, // [10:20] is the sub-list for method output_type
	0,  // [0:10] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_course_course_srv_proto_course_course_proto_init() }
func file_course_course_srv_proto_course_course_proto_init() {
	if File_course_course_srv_proto_course_course_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_course_course_srv_proto_course_course_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_course_course_srv_proto_course_course_proto_goTypes,
		DependencyIndexes: file_course_course_srv_proto_course_course_proto_depIdxs,
	}.Build()
	File_course_course_srv_proto_course_course_proto = out.File
	file_course_course_srv_proto_course_course_proto_rawDesc = nil
	file_course_course_srv_proto_course_course_proto_goTypes = nil
	file_course_course_srv_proto_course_course_proto_depIdxs = nil
}

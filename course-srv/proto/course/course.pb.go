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
	0x99, 0x0b, 0x0a, 0x0d, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x15, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x50,
	0x61, 0x67, 0x65, 0x44, 0x74, 0x6f, 0x1a, 0x15, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e,
	0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x50, 0x61, 0x67, 0x65, 0x44, 0x74, 0x6f, 0x12, 0x32, 0x0a,
	0x0a, 0x53, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x11, 0x2e, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x44, 0x74, 0x6f, 0x1a, 0x11,
	0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x44, 0x74,
	0x6f, 0x12, 0x30, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x12, 0x11, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x4c, 0x69, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x12, 0x42, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x0d, 0x2e, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x1d, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79,
	0x44, 0x74, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x0a, 0x53, 0x6f, 0x72, 0x74, 0x43,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x0f, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x53,
	0x6f, 0x72, 0x74, 0x44, 0x74, 0x6f, 0x1a, 0x0d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x3c, 0x0a, 0x11, 0x46, 0x69, 0x6e, 0x64, 0x43, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x0d, 0x2e, 0x62, 0x61, 0x73,
	0x69, 0x63, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x18, 0x2e, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x44, 0x74, 0x6f, 0x12, 0x3c, 0x0a, 0x11, 0x53, 0x61, 0x76, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x18, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x44,
	0x74, 0x6f, 0x1a, 0x0d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x12, 0x40, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x12, 0x17, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x50, 0x61, 0x67, 0x65, 0x44, 0x74, 0x6f, 0x1a, 0x17, 0x2e, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x50, 0x61, 0x67, 0x65,
	0x44, 0x74, 0x6f, 0x12, 0x35, 0x0a, 0x0b, 0x41, 0x6c, 0x6c, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f,
	0x72, 0x79, 0x12, 0x0d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e,
	0x67, 0x1a, 0x17, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67,
	0x6f, 0x72, 0x79, 0x44, 0x74, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x0c, 0x53, 0x61,
	0x76, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x13, 0x2e, 0x63, 0x6f, 0x75,
	0x72, 0x73, 0x65, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x44, 0x74, 0x6f, 0x1a,
	0x13, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72,
	0x79, 0x44, 0x74, 0x6f, 0x12, 0x32, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61,
	0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x11, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x3d, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74,
	0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x2e, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x44, 0x74, 0x6f, 0x1a,
	0x16, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72,
	0x50, 0x61, 0x67, 0x65, 0x44, 0x74, 0x6f, 0x12, 0x33, 0x0a, 0x0a, 0x41, 0x6c, 0x6c, 0x43, 0x68,
	0x61, 0x70, 0x74, 0x65, 0x72, 0x12, 0x0d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x1a, 0x16, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x43, 0x68,
	0x61, 0x70, 0x74, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x35, 0x0a, 0x0b,
	0x53, 0x61, 0x76, 0x65, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x12, 0x12, 0x2e, 0x63, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x1a,
	0x12, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x43, 0x68, 0x61, 0x70, 0x74, 0x65, 0x72,
	0x44, 0x74, 0x6f, 0x12, 0x31, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x68, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x3d, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x53,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x67, 0x65, 0x44, 0x74, 0x6f, 0x1a, 0x16, 0x2e,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61,
	0x67, 0x65, 0x44, 0x74, 0x6f, 0x12, 0x35, 0x0a, 0x0b, 0x53, 0x61, 0x76, 0x65, 0x53, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x53, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x74, 0x6f, 0x1a, 0x12, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73,
	0x65, 0x2e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x74, 0x6f, 0x12, 0x31, 0x0a, 0x0d,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x11, 0x2e,
	0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74,
	0x1a, 0x0d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12,
	0x3a, 0x0a, 0x0e, 0x4c, 0x69, 0x73, 0x74, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x46, 0x69, 0x6c,
	0x65, 0x12, 0x0d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67,
	0x1a, 0x19, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x46, 0x69, 0x6c, 0x65, 0x44, 0x74, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x0e, 0x53,
	0x61, 0x76, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x15, 0x2e,
	0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x46, 0x69, 0x6c,
	0x65, 0x44, 0x74, 0x6f, 0x1a, 0x15, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x43, 0x6f,
	0x75, 0x72, 0x73, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x74, 0x6f, 0x12, 0x30, 0x0a, 0x10, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12,
	0x0d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x0d,
	0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x3d, 0x0a,
	0x0b, 0x4c, 0x69, 0x73, 0x74, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x16, 0x2e, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x50, 0x61, 0x67,
	0x65, 0x44, 0x74, 0x6f, 0x1a, 0x16, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x54, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x50, 0x61, 0x67, 0x65, 0x44, 0x74, 0x6f, 0x12, 0x33, 0x0a, 0x0a,
	0x41, 0x6c, 0x6c, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x0d, 0x2e, 0x62, 0x61, 0x73,
	0x69, 0x63, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x16, 0x2e, 0x63, 0x6f, 0x75, 0x72,
	0x73, 0x65, 0x2e, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x35, 0x0a, 0x0b, 0x53, 0x61, 0x76, 0x65, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72,
	0x12, 0x12, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65,
	0x72, 0x44, 0x74, 0x6f, 0x1a, 0x12, 0x2e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2e, 0x54, 0x65,
	0x61, 0x63, 0x68, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x12, 0x31, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x54, 0x65, 0x61, 0x63, 0x68, 0x65, 0x72, 0x12, 0x11, 0x2e, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x62,
	0x61, 0x73, 0x69, 0x63, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x42, 0x20, 0x5a, 0x1e, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2d, 0x73, 0x72, 0x76,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_course_course_srv_proto_course_course_proto_goTypes = []interface{}{
	(*dto.CoursePageDto)(nil),         // 0: course.CoursePageDto
	(*dto.CourseDto)(nil),             // 1: course.CourseDto
	(*basic.StringList)(nil),          // 2: basic.StringList
	(*basic.String)(nil),              // 3: basic.String
	(*dto.SortDto)(nil),               // 4: course.SortDto
	(*dto.CourseContentDto)(nil),      // 5: course.CourseContentDto
	(*dto.CategoryPageDto)(nil),       // 6: course.CategoryPageDto
	(*dto.CategoryDto)(nil),           // 7: course.CategoryDto
	(*dto.ChapterPageDto)(nil),        // 8: course.ChapterPageDto
	(*dto.ChapterDto)(nil),            // 9: course.ChapterDto
	(*dto.SectionPageDto)(nil),        // 10: course.SectionPageDto
	(*dto.SectionDto)(nil),            // 11: course.SectionDto
	(*dto.CourseFileDto)(nil),         // 12: course.CourseFileDto
	(*dto.TeacherPageDto)(nil),        // 13: course.TeacherPageDto
	(*dto.TeacherDto)(nil),            // 14: course.TeacherDto
	(*dto.CourseCategoryDtoList)(nil), // 15: course.CourseCategoryDtoList
	(*dto.CategoryDtoList)(nil),       // 16: course.CategoryDtoList
	(*dto.ChapterDtoList)(nil),        // 17: course.ChapterDtoList
	(*dto.CourseFileDtoList)(nil),     // 18: course.CourseFileDtoList
	(*dto.TeacherDtoList)(nil),        // 19: course.TeacherDtoList
}
var file_course_course_srv_proto_course_course_proto_depIdxs = []int32{
	0,  // 0: course.CourseService.CourseList:input_type -> course.CoursePageDto
	1,  // 1: course.CourseService.SaveCourse:input_type -> course.CourseDto
	2,  // 2: course.CourseService.DeleteCourse:input_type -> basic.StringList
	3,  // 3: course.CourseService.ListCourseCategory:input_type -> basic.String
	4,  // 4: course.CourseService.SortCourse:input_type -> course.SortDto
	3,  // 5: course.CourseService.FindCourseContent:input_type -> basic.String
	5,  // 6: course.CourseService.SaveCourseContent:input_type -> course.CourseContentDto
	6,  // 7: course.CourseService.ListCategory:input_type -> course.CategoryPageDto
	3,  // 8: course.CourseService.AllCategory:input_type -> basic.String
	7,  // 9: course.CourseService.SaveCategory:input_type -> course.CategoryDto
	2,  // 10: course.CourseService.DeleteCategory:input_type -> basic.StringList
	8,  // 11: course.CourseService.ListChapter:input_type -> course.ChapterPageDto
	3,  // 12: course.CourseService.AllChapter:input_type -> basic.String
	9,  // 13: course.CourseService.SaveChapter:input_type -> course.ChapterDto
	2,  // 14: course.CourseService.DeleteChapter:input_type -> basic.StringList
	10, // 15: course.CourseService.ListSection:input_type -> course.SectionPageDto
	11, // 16: course.CourseService.SaveSection:input_type -> course.SectionDto
	2,  // 17: course.CourseService.DeleteSection:input_type -> basic.StringList
	3,  // 18: course.CourseService.ListCourseFile:input_type -> basic.String
	12, // 19: course.CourseService.SaveCourseFile:input_type -> course.CourseFileDto
	3,  // 20: course.CourseService.DeleteCourseFile:input_type -> basic.String
	13, // 21: course.CourseService.ListTeacher:input_type -> course.TeacherPageDto
	3,  // 22: course.CourseService.AllTeacher:input_type -> basic.String
	14, // 23: course.CourseService.SaveTeacher:input_type -> course.TeacherDto
	2,  // 24: course.CourseService.DeleteTeacher:input_type -> basic.StringList
	0,  // 25: course.CourseService.CourseList:output_type -> course.CoursePageDto
	1,  // 26: course.CourseService.SaveCourse:output_type -> course.CourseDto
	3,  // 27: course.CourseService.DeleteCourse:output_type -> basic.String
	15, // 28: course.CourseService.ListCourseCategory:output_type -> course.CourseCategoryDtoList
	3,  // 29: course.CourseService.SortCourse:output_type -> basic.String
	5,  // 30: course.CourseService.FindCourseContent:output_type -> course.CourseContentDto
	3,  // 31: course.CourseService.SaveCourseContent:output_type -> basic.String
	6,  // 32: course.CourseService.ListCategory:output_type -> course.CategoryPageDto
	16, // 33: course.CourseService.AllCategory:output_type -> course.CategoryDtoList
	7,  // 34: course.CourseService.SaveCategory:output_type -> course.CategoryDto
	3,  // 35: course.CourseService.DeleteCategory:output_type -> basic.String
	8,  // 36: course.CourseService.ListChapter:output_type -> course.ChapterPageDto
	17, // 37: course.CourseService.AllChapter:output_type -> course.ChapterDtoList
	9,  // 38: course.CourseService.SaveChapter:output_type -> course.ChapterDto
	3,  // 39: course.CourseService.DeleteChapter:output_type -> basic.String
	10, // 40: course.CourseService.ListSection:output_type -> course.SectionPageDto
	11, // 41: course.CourseService.SaveSection:output_type -> course.SectionDto
	3,  // 42: course.CourseService.DeleteSection:output_type -> basic.String
	18, // 43: course.CourseService.ListCourseFile:output_type -> course.CourseFileDtoList
	12, // 44: course.CourseService.SaveCourseFile:output_type -> course.CourseFileDto
	3,  // 45: course.CourseService.DeleteCourseFile:output_type -> basic.String
	13, // 46: course.CourseService.ListTeacher:output_type -> course.TeacherPageDto
	19, // 47: course.CourseService.AllTeacher:output_type -> course.TeacherDtoList
	14, // 48: course.CourseService.SaveTeacher:output_type -> course.TeacherDto
	3,  // 49: course.CourseService.DeleteTeacher:output_type -> basic.String
	25, // [25:50] is the sub-list for method output_type
	0,  // [0:25] is the sub-list for method input_type
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

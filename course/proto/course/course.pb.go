// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gitee.com/jingshanccc/course/course/proto/course/course.proto

package course

import (
	fmt "fmt"
	_ "gitee.com/jingshanccc/course/course/proto/dto"
	_ "gitee.com/jingshanccc/course/public/proto/basic"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

func init() {
	proto.RegisterFile("gitee.com/jingshanccc/course/course/proto/course/course.proto", fileDescriptor_6f6fb42b0d6a1ba7)
}

var fileDescriptor_6f6fb42b0d6a1ba7 = []byte{
	// 604 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x96, 0x6d, 0x6f, 0xda, 0x30,
	0x10, 0xc7, 0xdf, 0x55, 0xd3, 0x95, 0x40, 0xf1, 0xc4, 0x1e, 0xf2, 0x19, 0x26, 0x4a, 0xa9, 0xe8,
	0x36, 0xd6, 0x4e, 0xeb, 0x40, 0x95, 0x26, 0xb5, 0x13, 0x5a, 0xfa, 0x6a, 0xef, 0x8c, 0xb9, 0x81,
	0xa7, 0x10, 0x23, 0xc7, 0xb4, 0xea, 0xb7, 0xdd, 0x47, 0x99, 0xe2, 0x87, 0xa4, 0x71, 0x42, 0x21,
	0x6f, 0x5a, 0xe5, 0x9f, 0xff, 0xcf, 0x77, 0xe7, 0xbb, 0xd8, 0xc0, 0xd5, 0x92, 0x2b, 0xc4, 0x3e,
	0x13, 0xeb, 0xd3, 0xbf, 0x3c, 0x59, 0xa6, 0x2b, 0x9a, 0x30, 0xc6, 0x4e, 0x99, 0xd8, 0xca, 0x14,
	0xdd, 0xbf, 0x8d, 0x14, 0x4a, 0x94, 0xb5, 0xbe, 0xd6, 0xc8, 0x91, 0x79, 0x0a, 0xc7, 0x87, 0x2f,
	0xb3, 0xc8, 0x97, 0x32, 0x6b, 0x84, 0x5f, 0x5e, 0x64, 0x37, 0xdb, 0x79, 0xcc, 0x99, 0x65, 0xe7,
	0x34, 0xe5, 0xcc, 0xfc, 0x35, 0xf0, 0xf0, 0x5f, 0x07, 0x82, 0x89, 0x76, 0x46, 0x28, 0x1f, 0x38,
	0x43, 0x32, 0x06, 0x30, 0xc2, 0x2d, 0x4f, 0x15, 0xe9, 0xf5, 0x6d, 0x2c, 0xa3, 0xcd, 0xe8, 0x12,
	0xa7, 0x4a, 0x84, 0xf5, 0x32, 0x19, 0x02, 0x44, 0xf4, 0x01, 0x8d, 0x48, 0xba, 0x65, 0x53, 0xc6,
	0x55, 0x25, 0x32, 0x80, 0xd6, 0x14, 0x63, 0x54, 0x05, 0x65, 0xf2, 0x8b, 0x94, 0xe4, 0xc9, 0x32,
	0x4b, 0x22, 0x0c, 0x4a, 0x12, 0xf9, 0x00, 0x10, 0x09, 0xa9, 0xac, 0xbf, 0xe3, 0x96, 0xcc, 0xb4,
	0x2c, 0x86, 0xe7, 0xbe, 0x84, 0xee, 0x0d, 0x4f, 0x16, 0xc6, 0x3d, 0x11, 0x89, 0xc2, 0x44, 0x91,
	0xb2, 0x27, 0x7c, 0x57, 0x4e, 0xcb, 0xba, 0xb2, 0xec, 0x2e, 0xa1, 0x5b, 0x54, 0xe4, 0xe8, 0x9d,
	0x76, 0x3f, 0xf6, 0x37, 0x68, 0x65, 0x05, 0x4c, 0xa8, 0xc2, 0xa5, 0x90, 0x4f, 0xe4, 0x6d, 0x0e,
	0x5a, 0xc5, 0xed, 0xe7, 0xae, 0x17, 0xe4, 0x33, 0x74, 0x66, 0x92, 0xaf, 0xa9, 0x7c, 0xca, 0x17,
	0xf1, 0x72, 0xaf, 0xa0, 0x53, 0x25, 0x74, 0xeb, 0x3e, 0x41, 0x4b, 0xa7, 0xee, 0xb8, 0xd7, 0x35,
	0xc6, 0xb0, 0x4e, 0x24, 0x43, 0x68, 0xdb, 0x96, 0x38, 0x76, 0x7f, 0x53, 0xae, 0xe0, 0x58, 0x97,
	0xba, 0xa2, 0x1b, 0x85, 0x92, 0xbc, 0xc9, 0xd7, 0x35, 0x82, 0x2b, 0x74, 0x87, 0x4e, 0xce, 0x01,
	0xae, 0xe3, 0xd8, 0xd1, 0x5e, 0x89, 0x3e, 0xe4, 0x2a, 0x1c, 0xc1, 0xb1, 0xae, 0xd0, 0x52, 0xa4,
	0x6a, 0x0b, 0x6b, 0x34, 0x72, 0x06, 0x81, 0x2d, 0xcf, 0x82, 0x07, 0x57, 0x17, 0x21, 0x53, 0x5c,
	0x24, 0x45, 0x75, 0x56, 0xa8, 0x54, 0x57, 0xd6, 0x5d, 0xa2, 0x0e, 0x27, 0x9e, 0xad, 0x94, 0x68,
	0xa1, 0x15, 0x89, 0x3a, 0x70, 0x7f, 0xa2, 0x63, 0x68, 0xeb, 0x36, 0xe8, 0xb5, 0x6e, 0x78, 0x8c,
	0xfe, 0x5e, 0xbe, 0x2f, 0xcf, 0x6e, 0x66, 0x71, 0xdb, 0xf9, 0x15, 0xda, 0xc5, 0xac, 0x6b, 0xb6,
	0x57, 0x6b, 0x0e, 0xeb, 0x65, 0x32, 0x80, 0x93, 0xe7, 0x5f, 0x72, 0x5d, 0xf4, 0xfa, 0x6d, 0xbd,
	0x47, 0xca, 0x56, 0xcf, 0x87, 0xc6, 0x0a, 0x95, 0x6d, 0x2d, 0xeb, 0x76, 0x68, 0x1c, 0xbd, 0x6b,
	0x68, 0xec, 0x7b, 0x57, 0xe5, 0x05, 0x04, 0x11, 0x52, 0xc9, 0x56, 0x0d, 0x39, 0xdb, 0x43, 0x47,
	0x91, 0xaa, 0x2d, 0xac, 0xd1, 0x8a, 0x1e, 0x3a, 0x70, 0x7f, 0x0f, 0x47, 0x70, 0x9c, 0x7d, 0x0b,
	0x4d, 0xbf, 0xf7, 0x0b, 0x68, 0x4f, 0xa8, 0x14, 0xdb, 0x14, 0x63, 0x7b, 0x34, 0x7a, 0x64, 0xaf,
	0x72, 0xf8, 0xda, 0x73, 0xe2, 0xe4, 0x27, 0x3e, 0xce, 0xb2, 0x7b, 0x22, 0x5d, 0x35, 0x22, 0x75,
	0x44, 0x93, 0x44, 0x23, 0x6e, 0x00, 0x2d, 0x2b, 0xa0, 0xa2, 0x3c, 0xf6, 0xa9, 0x9a, 0x4b, 0x62,
	0x04, 0xc1, 0x2f, 0x8c, 0xa9, 0xc2, 0x45, 0xa3, 0x40, 0x23, 0xe8, 0x4d, 0xc5, 0x63, 0x12, 0x0b,
	0xfa, 0xf2, 0xf9, 0xef, 0x35, 0x60, 0x00, 0xaf, 0xee, 0x9a, 0x55, 0xf4, 0x11, 0x82, 0xeb, 0xc5,
	0xe2, 0x5e, 0xe4, 0x58, 0xde, 0xa5, 0x3b, 0x5c, 0xcf, 0x51, 0x16, 0x37, 0xa0, 0x17, 0xaa, 0xef,
	0x6e, 0xdb, 0x1f, 0xc9, 0x1f, 0x71, 0xc0, 0x6c, 0x9c, 0x41, 0x90, 0x4d, 0xe1, 0x2d, 0x52, 0x99,
	0x1c, 0x86, 0x7c, 0x1f, 0xfe, 0x1e, 0x34, 0xfd, 0x91, 0x32, 0x3f, 0xd2, 0x4f, 0xe7, 0xff, 0x03,
	0x00, 0x00, 0xff, 0xff, 0xd3, 0xa5, 0xf1, 0xd1, 0xdf, 0x08, 0x00, 0x00,
}

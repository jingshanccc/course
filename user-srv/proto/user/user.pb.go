// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: course/user-srv/proto/user/user.proto

package user

import (
	basic "course/proto/basic"
	dto "course/user-srv/proto/dto"
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

var File_course_user_srv_proto_user_user_proto protoreflect.FileDescriptor

var file_course_user_srv_proto_user_user_proto_rawDesc = []byte{
	0x0a, 0x25, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x72,
	0x76, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x24, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2d, 0x73, 0x72, 0x76, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2d, 0x73, 0x72, 0x76, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x74, 0x6f, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xe1, 0x07,
	0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x26, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x67,
	0x65, 0x44, 0x74, 0x6f, 0x1a, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x67, 0x65,
	0x44, 0x74, 0x6f, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x04, 0x53, 0x61, 0x76, 0x65, 0x12, 0x0d, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x1a, 0x0d, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x22, 0x00, 0x12, 0x2c, 0x0a,
	0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x11, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x62, 0x61, 0x73,
	0x69, 0x63, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x0c, 0x53,
	0x61, 0x76, 0x65, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x10, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x1a, 0x0d, 0x2e,
	0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12, 0x31,
	0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x1a, 0x12, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x22,
	0x00, 0x12, 0x28, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x0d, 0x2e, 0x62, 0x61,
	0x73, 0x69, 0x63, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x0d, 0x2e, 0x62, 0x61, 0x73,
	0x69, 0x63, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12, 0x2a, 0x0a, 0x08, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x44, 0x74, 0x6f, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x09, 0x53, 0x61, 0x76, 0x65, 0x45,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x1a, 0x0d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x09, 0x4c, 0x6f, 0x61, 0x64,
	0x4d, 0x65, 0x6e, 0x75, 0x73, 0x12, 0x0d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x1a, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x44, 0x74, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x08, 0x4c,
	0x6f, 0x61, 0x64, 0x54, 0x72, 0x65, 0x65, 0x12, 0x0d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x15, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x74, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x28, 0x0a,
	0x08, 0x53, 0x61, 0x76, 0x65, 0x4a, 0x73, 0x6f, 0x6e, 0x12, 0x0d, 0x2e, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x0d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63,
	0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x2f, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x0e, 0x2e, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x1a, 0x0d, 0x2e, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x30, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6c, 0x65,
	0x50, 0x61, 0x67, 0x65, 0x44, 0x74, 0x6f, 0x1a, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52,
	0x6f, 0x6c, 0x65, 0x50, 0x61, 0x67, 0x65, 0x44, 0x74, 0x6f, 0x12, 0x2b, 0x0a, 0x07, 0x41, 0x6c,
	0x6c, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x1a, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6c, 0x65,
	0x44, 0x74, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2a, 0x0a, 0x09, 0x52, 0x6f, 0x6c, 0x65, 0x4c,
	0x65, 0x76, 0x65, 0x6c, 0x12, 0x0d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x53, 0x74, 0x72,
	0x69, 0x6e, 0x67, 0x1a, 0x0e, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x49, 0x6e, 0x74, 0x65,
	0x67, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x08, 0x53, 0x61, 0x76, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12,
	0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x44, 0x74, 0x6f, 0x1a, 0x0d,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x44, 0x74, 0x6f, 0x12, 0x2a, 0x0a,
	0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x0d, 0x2e, 0x62, 0x61,
	0x73, 0x69, 0x63, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x0d, 0x2e, 0x62, 0x61, 0x73,
	0x69, 0x63, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x30, 0x0a, 0x10, 0x53, 0x61, 0x76,
	0x65, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x0d, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x44, 0x74, 0x6f, 0x1a, 0x0d, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x44, 0x74, 0x6f, 0x12, 0x34, 0x0a, 0x10, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12,
	0x0d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x11,
	0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x2c, 0x0a, 0x0c, 0x53, 0x61, 0x76, 0x65, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x44, 0x74, 0x6f,
	0x1a, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x52, 0x6f, 0x6c, 0x65, 0x44, 0x74, 0x6f, 0x12,
	0x30, 0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x6f, 0x6c, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x0d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x11,
	0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73,
	0x74, 0x42, 0x1c, 0x5a, 0x1a, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2d, 0x73, 0x72, 0x76, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_course_user_srv_proto_user_user_proto_goTypes = []interface{}{
	(*dto.PageDto)(nil),         // 0: user.PageDto
	(*dto.UserDto)(nil),         // 1: user.UserDto
	(*basic.StringList)(nil),    // 2: basic.StringList
	(*dto.UpdatePass)(nil),      // 3: user.UpdatePass
	(*dto.LoginUserDto)(nil),    // 4: user.LoginUserDto
	(*basic.String)(nil),        // 5: basic.String
	(*dto.UpdateEmail)(nil),     // 6: user.UpdateEmail
	(*basic.Integer)(nil),       // 7: basic.Integer
	(*dto.RolePageDto)(nil),     // 8: user.RolePageDto
	(*dto.RoleDto)(nil),         // 9: user.RoleDto
	(*dto.ResourceDtoList)(nil), // 10: user.ResourceDtoList
	(*dto.RoleDtoList)(nil),     // 11: user.RoleDtoList
}
var file_course_user_srv_proto_user_user_proto_depIdxs = []int32{
	0,  // 0: user.UserService.List:input_type -> user.PageDto
	1,  // 1: user.UserService.Save:input_type -> user.UserDto
	2,  // 2: user.UserService.Delete:input_type -> basic.StringList
	3,  // 3: user.UserService.SavePassword:input_type -> user.UpdatePass
	4,  // 4: user.UserService.Login:input_type -> user.LoginUserDto
	5,  // 5: user.UserService.Logout:input_type -> basic.String
	5,  // 6: user.UserService.UserInfo:input_type -> basic.String
	6,  // 7: user.UserService.SaveEmail:input_type -> user.UpdateEmail
	5,  // 8: user.UserService.LoadMenus:input_type -> basic.String
	5,  // 9: user.UserService.LoadTree:input_type -> basic.String
	5,  // 10: user.UserService.SaveJson:input_type -> basic.String
	7,  // 11: user.UserService.DeleteResource:input_type -> basic.Integer
	8,  // 12: user.UserService.RoleList:input_type -> user.RolePageDto
	5,  // 13: user.UserService.AllRole:input_type -> basic.String
	5,  // 14: user.UserService.RoleLevel:input_type -> basic.String
	9,  // 15: user.UserService.SaveRole:input_type -> user.RoleDto
	5,  // 16: user.UserService.DeleteRole:input_type -> basic.String
	9,  // 17: user.UserService.SaveRoleResource:input_type -> user.RoleDto
	5,  // 18: user.UserService.ListRoleResource:input_type -> basic.String
	9,  // 19: user.UserService.SaveRoleUser:input_type -> user.RoleDto
	5,  // 20: user.UserService.ListRoleUser:input_type -> basic.String
	0,  // 21: user.UserService.List:output_type -> user.PageDto
	1,  // 22: user.UserService.Save:output_type -> user.UserDto
	5,  // 23: user.UserService.Delete:output_type -> basic.String
	5,  // 24: user.UserService.SavePassword:output_type -> basic.String
	4,  // 25: user.UserService.Login:output_type -> user.LoginUserDto
	5,  // 26: user.UserService.Logout:output_type -> basic.String
	1,  // 27: user.UserService.UserInfo:output_type -> user.UserDto
	5,  // 28: user.UserService.SaveEmail:output_type -> basic.String
	10, // 29: user.UserService.LoadMenus:output_type -> user.ResourceDtoList
	10, // 30: user.UserService.LoadTree:output_type -> user.ResourceDtoList
	5,  // 31: user.UserService.SaveJson:output_type -> basic.String
	5,  // 32: user.UserService.DeleteResource:output_type -> basic.String
	8,  // 33: user.UserService.RoleList:output_type -> user.RolePageDto
	11, // 34: user.UserService.AllRole:output_type -> user.RoleDtoList
	7,  // 35: user.UserService.RoleLevel:output_type -> basic.Integer
	9,  // 36: user.UserService.SaveRole:output_type -> user.RoleDto
	5,  // 37: user.UserService.DeleteRole:output_type -> basic.String
	9,  // 38: user.UserService.SaveRoleResource:output_type -> user.RoleDto
	2,  // 39: user.UserService.ListRoleResource:output_type -> basic.StringList
	9,  // 40: user.UserService.SaveRoleUser:output_type -> user.RoleDto
	2,  // 41: user.UserService.ListRoleUser:output_type -> basic.StringList
	21, // [21:42] is the sub-list for method output_type
	0,  // [0:21] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func init() { file_course_user_srv_proto_user_user_proto_init() }
func file_course_user_srv_proto_user_user_proto_init() {
	if File_course_user_srv_proto_user_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_course_user_srv_proto_user_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_course_user_srv_proto_user_user_proto_goTypes,
		DependencyIndexes: file_course_user_srv_proto_user_user_proto_depIdxs,
	}.Build()
	File_course_user_srv_proto_user_user_proto = out.File
	file_course_user_srv_proto_user_user_proto_rawDesc = nil
	file_course_user_srv_proto_user_user_proto_goTypes = nil
	file_course_user_srv_proto_user_user_proto_depIdxs = nil
}

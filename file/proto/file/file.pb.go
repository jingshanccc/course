// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: gitee.com/jingshanccc/course/file/proto/file/file.proto

package file

import (
	dto "gitee.com/jingshanccc/course/file/proto/dto"
	basic "gitee.com/jingshanccc/course/public/proto/basic"
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

var File_gitee_com_jingshanccc_course_file_proto_file_file_proto protoreflect.FileDescriptor

var file_gitee_com_jingshanccc_course_file_proto_file_file_proto_rawDesc = []byte{
	0x0a, 0x37, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x69, 0x6e, 0x67,
	0x73, 0x68, 0x61, 0x6e, 0x63, 0x63, 0x63, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x66,
	0x69, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x66,
	0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x1a,
	0x36, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x69, 0x6e, 0x67, 0x73,
	0x68, 0x61, 0x6e, 0x63, 0x63, 0x63, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x66, 0x69,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x74, 0x6f, 0x2f, 0x66, 0x69, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x3b, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6a, 0x69, 0x6e, 0x67, 0x73, 0x68, 0x61, 0x6e, 0x63, 0x63, 0x63, 0x2f, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0x8e, 0x02, 0x0a, 0x0b, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x26, 0x0a, 0x06, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x0d,
	0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x74, 0x6f, 0x1a, 0x0d, 0x2e,
	0x66, 0x69, 0x6c, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x74, 0x6f, 0x12, 0x25, 0x0a, 0x05,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x0d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x53, 0x74,
	0x72, 0x69, 0x6e, 0x67, 0x1a, 0x0d, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x44, 0x74, 0x6f, 0x12, 0x2e, 0x0a, 0x0c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x55, 0x70, 0x6c,
	0x6f, 0x61, 0x64, 0x12, 0x0d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x53, 0x74, 0x72, 0x69,
	0x6e, 0x67, 0x1a, 0x0f, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79,
	0x52, 0x65, 0x73, 0x12, 0x31, 0x0a, 0x0b, 0x55, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x53, 0x68, 0x61,
	0x72, 0x64, 0x12, 0x12, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x53, 0x68,
	0x61, 0x72, 0x64, 0x44, 0x74, 0x6f, 0x1a, 0x0e, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x42,
	0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x12, 0x25, 0x0a, 0x05, 0x4d, 0x65, 0x72, 0x67, 0x65, 0x12,
	0x0d, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x74, 0x6f, 0x1a, 0x0d,
	0x2e, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x74, 0x6f, 0x12, 0x26, 0x0a,
	0x06, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x12, 0x0d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x1a, 0x0d, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6a, 0x69, 0x6e, 0x67, 0x73, 0x68, 0x61, 0x6e, 0x63, 0x63, 0x63, 0x2f, 0x63,
	0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x66, 0x69, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_gitee_com_jingshanccc_course_file_proto_file_file_proto_goTypes = []interface{}{
	(*dto.FileDto)(nil),      // 0: file.FileDto
	(*basic.String)(nil),     // 1: basic.String
	(*dto.FileShardDto)(nil), // 2: file.FileShardDto
	(*dto.VerifyRes)(nil),    // 3: file.VerifyRes
	(*basic.Boolean)(nil),    // 4: basic.Boolean
}
var file_gitee_com_jingshanccc_course_file_proto_file_file_proto_depIdxs = []int32{
	0, // 0: file.FileService.Upload:input_type -> file.FileDto
	1, // 1: file.FileService.Check:input_type -> basic.String
	1, // 2: file.FileService.VerifyUpload:input_type -> basic.String
	2, // 3: file.FileService.UploadShard:input_type -> file.FileShardDto
	0, // 4: file.FileService.Merge:input_type -> file.FileDto
	1, // 5: file.FileService.Cancel:input_type -> basic.String
	0, // 6: file.FileService.Upload:output_type -> file.FileDto
	0, // 7: file.FileService.Check:output_type -> file.FileDto
	3, // 8: file.FileService.VerifyUpload:output_type -> file.VerifyRes
	4, // 9: file.FileService.UploadShard:output_type -> basic.Boolean
	0, // 10: file.FileService.Merge:output_type -> file.FileDto
	1, // 11: file.FileService.Cancel:output_type -> basic.String
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gitee_com_jingshanccc_course_file_proto_file_file_proto_init() }
func file_gitee_com_jingshanccc_course_file_proto_file_file_proto_init() {
	if File_gitee_com_jingshanccc_course_file_proto_file_file_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gitee_com_jingshanccc_course_file_proto_file_file_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gitee_com_jingshanccc_course_file_proto_file_file_proto_goTypes,
		DependencyIndexes: file_gitee_com_jingshanccc_course_file_proto_file_file_proto_depIdxs,
	}.Build()
	File_gitee_com_jingshanccc_course_file_proto_file_file_proto = out.File
	file_gitee_com_jingshanccc_course_file_proto_file_file_proto_rawDesc = nil
	file_gitee_com_jingshanccc_course_file_proto_file_file_proto_goTypes = nil
	file_gitee_com_jingshanccc_course_file_proto_file_file_proto_depIdxs = nil
}
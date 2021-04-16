// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.7
// source: gitee.com/jingshanccc/course/user/proto/dto/user.proto

package dto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	LoginName   string   `protobuf:"bytes,3,opt,name=login_name,json=loginName,proto3" json:"login_name,omitempty"`
	Phone       string   `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"`
	Email       string   `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	AvatarName  string   `protobuf:"bytes,6,opt,name=avatar_name,json=avatarName,proto3" json:"avatar_name,omitempty"`
	AvatarPath  string   `protobuf:"bytes,7,opt,name=avatar_path,json=avatarPath,proto3" json:"avatar_path,omitempty"`
	Gender      string   `protobuf:"bytes,8,opt,name=gender,proto3" json:"gender,omitempty"`
	IsAdmin     bool     `protobuf:"varint,9,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
	CreateBy    string   `protobuf:"bytes,10,opt,name=create_by,json=createBy,proto3" json:"create_by,omitempty"`
	CreateTime  string   `protobuf:"bytes,11,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	UpdateBy    string   `protobuf:"bytes,12,opt,name=update_by,json=updateBy,proto3" json:"update_by,omitempty"`
	UpdateTime  string   `protobuf:"bytes,13,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
	Roles       string   `protobuf:"bytes,14,opt,name=roles,proto3" json:"roles,omitempty"`
	Permissions []string `protobuf:"bytes,15,rep,name=permissions,proto3" json:"permissions,omitempty"`
}

func (x *UserDto) Reset() {
	*x = UserDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDto) ProtoMessage() {}

func (x *UserDto) ProtoReflect() protoreflect.Message {
	mi := &file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDto.ProtoReflect.Descriptor instead.
func (*UserDto) Descriptor() ([]byte, []int) {
	return file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserDto) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserDto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserDto) GetLoginName() string {
	if x != nil {
		return x.LoginName
	}
	return ""
}

func (x *UserDto) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *UserDto) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserDto) GetAvatarName() string {
	if x != nil {
		return x.AvatarName
	}
	return ""
}

func (x *UserDto) GetAvatarPath() string {
	if x != nil {
		return x.AvatarPath
	}
	return ""
}

func (x *UserDto) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *UserDto) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

func (x *UserDto) GetCreateBy() string {
	if x != nil {
		return x.CreateBy
	}
	return ""
}

func (x *UserDto) GetCreateTime() string {
	if x != nil {
		return x.CreateTime
	}
	return ""
}

func (x *UserDto) GetUpdateBy() string {
	if x != nil {
		return x.UpdateBy
	}
	return ""
}

func (x *UserDto) GetUpdateTime() string {
	if x != nil {
		return x.UpdateTime
	}
	return ""
}

func (x *UserDto) GetRoles() string {
	if x != nil {
		return x.Roles
	}
	return ""
}

func (x *UserDto) GetPermissions() []string {
	if x != nil {
		return x.Permissions
	}
	return nil
}

type PageDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page       int64      `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Size       int64      `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
	Sort       []string   `protobuf:"bytes,3,rep,name=sort,proto3" json:"sort,omitempty"`
	Blurry     string     `protobuf:"bytes,4,opt,name=blurry,proto3" json:"blurry,omitempty"`
	CreateTime []string   `protobuf:"bytes,5,rep,name=createTime,proto3" json:"createTime,omitempty"`
	Total      int64      `protobuf:"varint,6,opt,name=total,proto3" json:"total,omitempty"`
	Rows       []*UserDto `protobuf:"bytes,7,rep,name=rows,proto3" json:"rows,omitempty"`
}

func (x *PageDto) Reset() {
	*x = PageDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PageDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PageDto) ProtoMessage() {}

func (x *PageDto) ProtoReflect() protoreflect.Message {
	mi := &file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PageDto.ProtoReflect.Descriptor instead.
func (*PageDto) Descriptor() ([]byte, []int) {
	return file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_rawDescGZIP(), []int{1}
}

func (x *PageDto) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PageDto) GetSize() int64 {
	if x != nil {
		return x.Size
	}
	return 0
}

func (x *PageDto) GetSort() []string {
	if x != nil {
		return x.Sort
	}
	return nil
}

func (x *PageDto) GetBlurry() string {
	if x != nil {
		return x.Blurry
	}
	return ""
}

func (x *PageDto) GetCreateTime() []string {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *PageDto) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PageDto) GetRows() []*UserDto {
	if x != nil {
		return x.Rows
	}
	return nil
}

type LoginUserDto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	LoginName string `protobuf:"bytes,2,opt,name=login_name,json=loginName,proto3" json:"login_name,omitempty"`
	Name      string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"` // 登陆请求时用作图形验证码 响应时作为
	Token     string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
	Password  string `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *LoginUserDto) Reset() {
	*x = LoginUserDto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginUserDto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginUserDto) ProtoMessage() {}

func (x *LoginUserDto) ProtoReflect() protoreflect.Message {
	mi := &file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginUserDto.ProtoReflect.Descriptor instead.
func (*LoginUserDto) Descriptor() ([]byte, []int) {
	return file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_rawDescGZIP(), []int{2}
}

func (x *LoginUserDto) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *LoginUserDto) GetLoginName() string {
	if x != nil {
		return x.LoginName
	}
	return ""
}

func (x *LoginUserDto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LoginUserDto) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *LoginUserDto) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UpdatePass struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OldPass string `protobuf:"bytes,1,opt,name=oldPass,proto3" json:"oldPass,omitempty"`
	NewPass string `protobuf:"bytes,2,opt,name=newPass,proto3" json:"newPass,omitempty"`
	UserId  string `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *UpdatePass) Reset() {
	*x = UpdatePass{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePass) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePass) ProtoMessage() {}

func (x *UpdatePass) ProtoReflect() protoreflect.Message {
	mi := &file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePass.ProtoReflect.Descriptor instead.
func (*UpdatePass) Descriptor() ([]byte, []int) {
	return file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_rawDescGZIP(), []int{3}
}

func (x *UpdatePass) GetOldPass() string {
	if x != nil {
		return x.OldPass
	}
	return ""
}

func (x *UpdatePass) GetNewPass() string {
	if x != nil {
		return x.NewPass
	}
	return ""
}

func (x *UpdatePass) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type UpdateEmail struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pass   string `protobuf:"bytes,1,opt,name=pass,proto3" json:"pass,omitempty"`
	Email  string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Code   string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	UserId string `protobuf:"bytes,4,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *UpdateEmail) Reset() {
	*x = UpdateEmail{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateEmail) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateEmail) ProtoMessage() {}

func (x *UpdateEmail) ProtoReflect() protoreflect.Message {
	mi := &file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateEmail.ProtoReflect.Descriptor instead.
func (*UpdateEmail) Descriptor() ([]byte, []int) {
	return file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateEmail) GetPass() string {
	if x != nil {
		return x.Pass
	}
	return ""
}

func (x *UpdateEmail) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UpdateEmail) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *UpdateEmail) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_gitee_com_jingshanccc_course_user_proto_dto_user_proto protoreflect.FileDescriptor

var file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_rawDesc = []byte{
	0x0a, 0x36, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x69, 0x6e, 0x67,
	0x73, 0x68, 0x61, 0x6e, 0x63, 0x63, 0x63, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x74, 0x6f, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0xa1,
	0x03, 0x0a, 0x07, 0x55, 0x73, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d,
	0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x5f, 0x70, 0x61, 0x74, 0x68, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x50, 0x61, 0x74, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x67,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1b, 0x0a, 0x09,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x62, 0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x42, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f,
	0x6c, 0x65, 0x73, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73,
	0x12, 0x20, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x0f, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x73, 0x22, 0xb6, 0x01, 0x0a, 0x07, 0x50, 0x61, 0x67, 0x65, 0x44, 0x74, 0x6f, 0x12, 0x12,
	0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x6c,
	0x75, 0x72, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x62, 0x6c, 0x75, 0x72,
	0x72, 0x79, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x21, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x44, 0x74, 0x6f, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x22, 0x83, 0x01, 0x0a, 0x0c,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x44, 0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a,
	0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x22, 0x58, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x73, 0x73, 0x12,
	0x18, 0x0a, 0x07, 0x6f, 0x6c, 0x64, 0x50, 0x61, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x6f, 0x6c, 0x64, 0x50, 0x61, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x77,
	0x50, 0x61, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x77, 0x50,
	0x61, 0x73, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x63, 0x0a, 0x0b, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x73, 0x73, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x69,
	0x6e, 0x67, 0x73, 0x68, 0x61, 0x6e, 0x63, 0x63, 0x63, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x74, 0x6f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_rawDescOnce sync.Once
	file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_rawDescData = file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_rawDesc
)

func file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_rawDescGZIP() []byte {
	file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_rawDescOnce.Do(func() {
		file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_rawDescData)
	})
	return file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_rawDescData
}

var file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_goTypes = []interface{}{
	(*UserDto)(nil),      // 0: user.UserDto
	(*PageDto)(nil),      // 1: user.PageDto
	(*LoginUserDto)(nil), // 2: user.LoginUserDto
	(*UpdatePass)(nil),   // 3: user.UpdatePass
	(*UpdateEmail)(nil),  // 4: user.UpdateEmail
}
var file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_depIdxs = []int32{
	0, // 0: user.PageDto.rows:type_name -> user.UserDto
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_init() }
func file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_init() {
	if File_gitee_com_jingshanccc_course_user_proto_dto_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserDto); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PageDto); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginUserDto); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePass); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateEmail); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_goTypes,
		DependencyIndexes: file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_depIdxs,
		MessageInfos:      file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_msgTypes,
	}.Build()
	File_gitee_com_jingshanccc_course_user_proto_dto_user_proto = out.File
	file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_rawDesc = nil
	file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_goTypes = nil
	file_gitee_com_jingshanccc_course_user_proto_dto_user_proto_depIdxs = nil
}

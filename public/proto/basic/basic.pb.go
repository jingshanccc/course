// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: gitee.com/jingshanccc/course/public/proto/basic/basic.proto

package basic

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

type Integer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Integer) Reset() {
	*x = Integer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Integer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Integer) ProtoMessage() {}

func (x *Integer) ProtoReflect() protoreflect.Message {
	mi := &file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Integer.ProtoReflect.Descriptor instead.
func (*Integer) Descriptor() ([]byte, []int) {
	return file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_rawDescGZIP(), []int{0}
}

func (x *Integer) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type IntegerList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []int32 `protobuf:"varint,1,rep,packed,name=ids,proto3" json:"ids,omitempty"`
}

func (x *IntegerList) Reset() {
	*x = IntegerList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IntegerList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IntegerList) ProtoMessage() {}

func (x *IntegerList) ProtoReflect() protoreflect.Message {
	mi := &file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IntegerList.ProtoReflect.Descriptor instead.
func (*IntegerList) Descriptor() ([]byte, []int) {
	return file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_rawDescGZIP(), []int{1}
}

func (x *IntegerList) GetIds() []int32 {
	if x != nil {
		return x.Ids
	}
	return nil
}

type String struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Str string `protobuf:"bytes,1,opt,name=str,proto3" json:"str,omitempty"`
}

func (x *String) Reset() {
	*x = String{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *String) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*String) ProtoMessage() {}

func (x *String) ProtoReflect() protoreflect.Message {
	mi := &file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use String.ProtoReflect.Descriptor instead.
func (*String) Descriptor() ([]byte, []int) {
	return file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_rawDescGZIP(), []int{2}
}

func (x *String) GetStr() string {
	if x != nil {
		return x.Str
	}
	return ""
}

type StringList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rows []string `protobuf:"bytes,1,rep,name=rows,proto3" json:"rows,omitempty"`
}

func (x *StringList) Reset() {
	*x = StringList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StringList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StringList) ProtoMessage() {}

func (x *StringList) ProtoReflect() protoreflect.Message {
	mi := &file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StringList.ProtoReflect.Descriptor instead.
func (*StringList) Descriptor() ([]byte, []int) {
	return file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_rawDescGZIP(), []int{3}
}

func (x *StringList) GetRows() []string {
	if x != nil {
		return x.Rows
	}
	return nil
}

type Boolean struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Is bool `protobuf:"varint,1,opt,name=is,proto3" json:"is,omitempty"`
}

func (x *Boolean) Reset() {
	*x = Boolean{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Boolean) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Boolean) ProtoMessage() {}

func (x *Boolean) ProtoReflect() protoreflect.Message {
	mi := &file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Boolean.ProtoReflect.Descriptor instead.
func (*Boolean) Descriptor() ([]byte, []int) {
	return file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_rawDescGZIP(), []int{4}
}

func (x *Boolean) GetIs() bool {
	if x != nil {
		return x.Is
	}
	return false
}

type Pair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key    string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Values []string `protobuf:"bytes,2,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *Pair) Reset() {
	*x = Pair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pair) ProtoMessage() {}

func (x *Pair) ProtoReflect() protoreflect.Message {
	mi := &file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pair.ProtoReflect.Descriptor instead.
func (*Pair) Descriptor() ([]byte, []int) {
	return file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_rawDescGZIP(), []int{5}
}

func (x *Pair) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *Pair) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

// A HTTP request as RPC
// Forward by the api handler
type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method string           `protobuf:"bytes,1,opt,name=method,proto3" json:"method,omitempty"`
	Path   string           `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Header map[string]*Pair `protobuf:"bytes,3,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Get    map[string]*Pair `protobuf:"bytes,4,rep,name=get,proto3" json:"get,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Post   map[string]*Pair `protobuf:"bytes,5,rep,name=post,proto3" json:"post,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body   string           `protobuf:"bytes,6,opt,name=body,proto3" json:"body,omitempty"` // raw request body; if not application/x-www-form-urlencoded
	Url    string           `protobuf:"bytes,7,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_rawDescGZIP(), []int{6}
}

func (x *Request) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Request) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Request) GetHeader() map[string]*Pair {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Request) GetGet() map[string]*Pair {
	if x != nil {
		return x.Get
	}
	return nil
}

func (x *Request) GetPost() map[string]*Pair {
	if x != nil {
		return x.Post
	}
	return nil
}

func (x *Request) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Request) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

// A HTTP response as RPC
// Expected response for the api handler
type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32            `protobuf:"varint,1,opt,name=statusCode,proto3" json:"statusCode,omitempty"`
	Header     map[string]*Pair `protobuf:"bytes,2,rep,name=header,proto3" json:"header,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Body       string           `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_rawDescGZIP(), []int{7}
}

func (x *Response) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *Response) GetHeader() map[string]*Pair {
	if x != nil {
		return x.Header
	}
	return nil
}

func (x *Response) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

var File_gitee_com_jingshanccc_course_public_proto_basic_basic_proto protoreflect.FileDescriptor

var file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x67, 0x69, 0x74, 0x65, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x69, 0x6e, 0x67,
	0x73, 0x68, 0x61, 0x6e, 0x63, 0x63, 0x63, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x73, 0x69,
	0x63, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x62,
	0x61, 0x73, 0x69, 0x63, 0x22, 0x19, 0x0a, 0x07, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x1f, 0x0a, 0x0b, 0x49, 0x6e, 0x74, 0x65, 0x67, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x10,
	0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x03, 0x69, 0x64, 0x73,
	0x22, 0x1a, 0x0a, 0x06, 0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x74,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x74, 0x72, 0x22, 0x20, 0x0a, 0x0a,
	0x53, 0x74, 0x72, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x6f,
	0x77, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x22, 0x19,
	0x0a, 0x07, 0x42, 0x6f, 0x6f, 0x6c, 0x65, 0x61, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x69, 0x73, 0x22, 0x30, 0x0a, 0x04, 0x50, 0x61, 0x69,
	0x72, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6b, 0x65, 0x79, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0xbb, 0x03, 0x0a, 0x07,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x12, 0x32, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x03, 0x67, 0x65, 0x74, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x67,
	0x65, 0x74, 0x12, 0x2c, 0x0a, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x18, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x2e, 0x50, 0x6f, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x62, 0x6f, 0x64, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x1a, 0x46, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x50,
	0x61, 0x69, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x43,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x62, 0x61,
	0x73, 0x69, 0x63, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x44, 0x0a, 0x09, 0x50, 0x6f, 0x73, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x21, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x52, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xbb, 0x01, 0x0a, 0x08, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x33, 0x0a, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x62,
	0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x1a,
	0x46, 0x0a, 0x0b, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x21, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x62, 0x61, 0x73, 0x69, 0x63, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x65, 0x65,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x69, 0x6e, 0x67, 0x73, 0x68, 0x61, 0x6e, 0x63, 0x63, 0x63,
	0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62, 0x61, 0x73, 0x69, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_rawDescOnce sync.Once
	file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_rawDescData = file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_rawDesc
)

func file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_rawDescGZIP() []byte {
	file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_rawDescOnce.Do(func() {
		file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_rawDescData = protoimpl.X.CompressGZIP(file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_rawDescData)
	})
	return file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_rawDescData
}

var file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_goTypes = []interface{}{
	(*Integer)(nil),     // 0: basic.Integer
	(*IntegerList)(nil), // 1: basic.IntegerList
	(*String)(nil),      // 2: basic.String
	(*StringList)(nil),  // 3: basic.StringList
	(*Boolean)(nil),     // 4: basic.Boolean
	(*Pair)(nil),        // 5: basic.Pair
	(*Request)(nil),     // 6: basic.Request
	(*Response)(nil),    // 7: basic.Response
	nil,                 // 8: basic.Request.HeaderEntry
	nil,                 // 9: basic.Request.GetEntry
	nil,                 // 10: basic.Request.PostEntry
	nil,                 // 11: basic.Response.HeaderEntry
}
var file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_depIdxs = []int32{
	8,  // 0: basic.Request.header:type_name -> basic.Request.HeaderEntry
	9,  // 1: basic.Request.get:type_name -> basic.Request.GetEntry
	10, // 2: basic.Request.post:type_name -> basic.Request.PostEntry
	11, // 3: basic.Response.header:type_name -> basic.Response.HeaderEntry
	5,  // 4: basic.Request.HeaderEntry.value:type_name -> basic.Pair
	5,  // 5: basic.Request.GetEntry.value:type_name -> basic.Pair
	5,  // 6: basic.Request.PostEntry.value:type_name -> basic.Pair
	5,  // 7: basic.Response.HeaderEntry.value:type_name -> basic.Pair
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_init() }
func file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_init() {
	if File_gitee_com_jingshanccc_course_public_proto_basic_basic_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Integer); i {
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
		file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IntegerList); i {
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
		file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*String); i {
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
		file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StringList); i {
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
		file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Boolean); i {
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
		file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pair); i {
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
		file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_goTypes,
		DependencyIndexes: file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_depIdxs,
		MessageInfos:      file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_msgTypes,
	}.Build()
	File_gitee_com_jingshanccc_course_public_proto_basic_basic_proto = out.File
	file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_rawDesc = nil
	file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_goTypes = nil
	file_gitee_com_jingshanccc_course_public_proto_basic_basic_proto_depIdxs = nil
}

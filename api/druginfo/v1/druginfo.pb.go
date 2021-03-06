// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: api/druginfo/v1/druginfo.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/go-kratos/kratos/v2/errors"
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

type ErrorReason int32

const (
	// 为某个枚举单独设置错误码
	ErrorReason_DRUG_NOT_FOUND ErrorReason = 0
)

// Enum value maps for ErrorReason.
var (
	ErrorReason_name = map[int32]string{
		0: "DRUG_NOT_FOUND",
	}
	ErrorReason_value = map[string]int32{
		"DRUG_NOT_FOUND": 0,
	}
)

func (x ErrorReason) Enum() *ErrorReason {
	p := new(ErrorReason)
	*p = x
	return p
}

func (x ErrorReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ErrorReason) Descriptor() protoreflect.EnumDescriptor {
	return file_api_druginfo_v1_druginfo_proto_enumTypes[0].Descriptor()
}

func (ErrorReason) Type() protoreflect.EnumType {
	return &file_api_druginfo_v1_druginfo_proto_enumTypes[0]
}

func (x ErrorReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ErrorReason.Descriptor instead.
func (ErrorReason) EnumDescriptor() ([]byte, []int) {
	return file_api_druginfo_v1_druginfo_proto_rawDescGZIP(), []int{0}
}

type Druginfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	CnName string `protobuf:"bytes,2,opt,name=cn_name,json=cnName,proto3" json:"cn_name,omitempty"`
}

func (x *Druginfo) Reset() {
	*x = Druginfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_druginfo_v1_druginfo_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Druginfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Druginfo) ProtoMessage() {}

func (x *Druginfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_druginfo_v1_druginfo_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Druginfo.ProtoReflect.Descriptor instead.
func (*Druginfo) Descriptor() ([]byte, []int) {
	return file_api_druginfo_v1_druginfo_proto_rawDescGZIP(), []int{0}
}

func (x *Druginfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Druginfo) GetCnName() string {
	if x != nil {
		return x.CnName
	}
	return ""
}

type CreateDrugRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CnName string `protobuf:"bytes,1,opt,name=cn_name,json=cnName,proto3" json:"cn_name,omitempty"`
}

func (x *CreateDrugRequest) Reset() {
	*x = CreateDrugRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_druginfo_v1_druginfo_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDrugRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDrugRequest) ProtoMessage() {}

func (x *CreateDrugRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_druginfo_v1_druginfo_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDrugRequest.ProtoReflect.Descriptor instead.
func (*CreateDrugRequest) Descriptor() ([]byte, []int) {
	return file_api_druginfo_v1_druginfo_proto_rawDescGZIP(), []int{1}
}

func (x *CreateDrugRequest) GetCnName() string {
	if x != nil {
		return x.CnName
	}
	return ""
}

type UpdatedruginfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	DrugName string `protobuf:"bytes,2,opt,name=drug_name,json=drugName,proto3" json:"drug_name,omitempty"` // the drug_name of string must be between 1 and 63 character;
}

func (x *UpdatedruginfoRequest) Reset() {
	*x = UpdatedruginfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_druginfo_v1_druginfo_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatedruginfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatedruginfoRequest) ProtoMessage() {}

func (x *UpdatedruginfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_druginfo_v1_druginfo_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatedruginfoRequest.ProtoReflect.Descriptor instead.
func (*UpdatedruginfoRequest) Descriptor() ([]byte, []int) {
	return file_api_druginfo_v1_druginfo_proto_rawDescGZIP(), []int{2}
}

func (x *UpdatedruginfoRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdatedruginfoRequest) GetDrugName() string {
	if x != nil {
		return x.DrugName
	}
	return ""
}

type UpdatedruginfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Druginfo *Druginfo `protobuf:"bytes,1,opt,name=druginfo,proto3" json:"druginfo,omitempty"`
}

func (x *UpdatedruginfoReply) Reset() {
	*x = UpdatedruginfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_druginfo_v1_druginfo_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatedruginfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatedruginfoReply) ProtoMessage() {}

func (x *UpdatedruginfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_druginfo_v1_druginfo_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatedruginfoReply.ProtoReflect.Descriptor instead.
func (*UpdatedruginfoReply) Descriptor() ([]byte, []int) {
	return file_api_druginfo_v1_druginfo_proto_rawDescGZIP(), []int{3}
}

func (x *UpdatedruginfoReply) GetDruginfo() *Druginfo {
	if x != nil {
		return x.Druginfo
	}
	return nil
}

type GetDrugRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetDrugRequest) Reset() {
	*x = GetDrugRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_druginfo_v1_druginfo_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDrugRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDrugRequest) ProtoMessage() {}

func (x *GetDrugRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_druginfo_v1_druginfo_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDrugRequest.ProtoReflect.Descriptor instead.
func (*GetDrugRequest) Descriptor() ([]byte, []int) {
	return file_api_druginfo_v1_druginfo_proto_rawDescGZIP(), []int{4}
}

func (x *GetDrugRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetDrugReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Druginfo *Druginfo `protobuf:"bytes,1,opt,name=druginfo,proto3" json:"druginfo,omitempty"`
	Code     int32     `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Msg      string    `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *GetDrugReply) Reset() {
	*x = GetDrugReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_druginfo_v1_druginfo_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDrugReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDrugReply) ProtoMessage() {}

func (x *GetDrugReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_druginfo_v1_druginfo_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDrugReply.ProtoReflect.Descriptor instead.
func (*GetDrugReply) Descriptor() ([]byte, []int) {
	return file_api_druginfo_v1_druginfo_proto_rawDescGZIP(), []int{5}
}

func (x *GetDrugReply) GetDruginfo() *Druginfo {
	if x != nil {
		return x.Druginfo
	}
	return nil
}

func (x *GetDrugReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetDrugReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type ListdruginfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListdruginfoRequest) Reset() {
	*x = ListdruginfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_druginfo_v1_druginfo_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListdruginfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListdruginfoRequest) ProtoMessage() {}

func (x *ListdruginfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_druginfo_v1_druginfo_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListdruginfoRequest.ProtoReflect.Descriptor instead.
func (*ListdruginfoRequest) Descriptor() ([]byte, []int) {
	return file_api_druginfo_v1_druginfo_proto_rawDescGZIP(), []int{6}
}

type ListdruginfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Results []*Druginfo `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	Code    int32       `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
	Msg     string      `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ListdruginfoReply) Reset() {
	*x = ListdruginfoReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_druginfo_v1_druginfo_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListdruginfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListdruginfoReply) ProtoMessage() {}

func (x *ListdruginfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_druginfo_v1_druginfo_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListdruginfoReply.ProtoReflect.Descriptor instead.
func (*ListdruginfoReply) Descriptor() ([]byte, []int) {
	return file_api_druginfo_v1_druginfo_proto_rawDescGZIP(), []int{7}
}

func (x *ListdruginfoReply) GetResults() []*Druginfo {
	if x != nil {
		return x.Results
	}
	return nil
}

func (x *ListdruginfoReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ListdruginfoReply) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_api_druginfo_v1_druginfo_proto protoreflect.FileDescriptor

var file_api_druginfo_v1_druginfo_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x72, 0x75, 0x67, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x76,
	0x31, 0x2f, 0x64, 0x72, 0x75, 0x67, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x0b, 0x64, 0x72, 0x75, 0x67, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31, 0x1a, 0x13, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x33, 0x0a, 0x08, 0x44,
	0x72, 0x75, 0x67, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x6e, 0x5f, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6e, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0x2c, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x72, 0x75, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x58,
	0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x72, 0x75, 0x67, 0x69, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x26, 0x0a, 0x09, 0x64, 0x72, 0x75, 0x67, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x09, 0xfa, 0x42, 0x06, 0x72, 0x04, 0x10, 0x01, 0x18, 0x3f, 0x52, 0x08,
	0x64, 0x72, 0x75, 0x67, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x48, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x64, 0x72, 0x75, 0x67, 0x69, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x31, 0x0a, 0x08, 0x64, 0x72, 0x75, 0x67, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x15, 0x2e, 0x64, 0x72, 0x75, 0x67, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x72, 0x75, 0x67, 0x69, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x64, 0x72, 0x75, 0x67, 0x69, 0x6e,
	0x66, 0x6f, 0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x44, 0x72, 0x75, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x02, 0x69, 0x64, 0x22, 0x67, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x44, 0x72, 0x75, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x31, 0x0a,
	0x08, 0x64, 0x72, 0x75, 0x67, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x64, 0x72, 0x75, 0x67, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x72,
	0x75, 0x67, 0x69, 0x6e, 0x66, 0x6f, 0x52, 0x08, 0x64, 0x72, 0x75, 0x67, 0x69, 0x6e, 0x66, 0x6f,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x64, 0x72,
	0x75, 0x67, 0x69, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x6a, 0x0a,
	0x11, 0x4c, 0x69, 0x73, 0x74, 0x64, 0x72, 0x75, 0x67, 0x69, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x2f, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x64, 0x72, 0x75, 0x67, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x44, 0x72, 0x75, 0x67, 0x69, 0x6e, 0x66, 0x6f, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x2a, 0x2d, 0x0a, 0x0b, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x0e, 0x44, 0x52, 0x55, 0x47,
	0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x46, 0x4f, 0x55, 0x4e, 0x44, 0x10, 0x00, 0x1a, 0x04, 0xa8, 0x45,
	0x94, 0x03, 0x1a, 0x04, 0xa0, 0x45, 0xf4, 0x03, 0x32, 0xc5, 0x02, 0x0a, 0x09, 0x44, 0x72, 0x75,
	0x67, 0x69, 0x6e, 0x66, 0x6f, 0x73, 0x12, 0x49, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x44, 0x72, 0x75, 0x67, 0x12, 0x1e, 0x2e, 0x64, 0x72, 0x75, 0x67, 0x69, 0x6e, 0x66, 0x6f, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x72, 0x75, 0x67, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x64, 0x72, 0x75, 0x67, 0x69, 0x6e, 0x66, 0x6f, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x72, 0x75, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x12, 0x54, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x72, 0x75, 0x67, 0x12,
	0x22, 0x2e, 0x64, 0x72, 0x75, 0x67, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x64, 0x72, 0x75, 0x67, 0x69, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x64, 0x72, 0x75, 0x67, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x72, 0x75, 0x67, 0x69, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x44, 0x72,
	0x75, 0x67, 0x12, 0x1b, 0x2e, 0x64, 0x72, 0x75, 0x67, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x44, 0x72, 0x75, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x64, 0x72, 0x75, 0x67, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x44, 0x72, 0x75, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x52, 0x0a, 0x0c,
	0x4c, 0x69, 0x73, 0x74, 0x64, 0x72, 0x75, 0x67, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x20, 0x2e, 0x64,
	0x72, 0x75, 0x67, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x64,
	0x72, 0x75, 0x67, 0x69, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x64, 0x72, 0x75, 0x67, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x64, 0x72, 0x75, 0x67, 0x69, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x42, 0x27, 0x0a, 0x0f, 0x64, 0x72, 0x75, 0x67, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x72, 0x75, 0x67, 0x69,
	0x6e, 0x66, 0x6f, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_druginfo_v1_druginfo_proto_rawDescOnce sync.Once
	file_api_druginfo_v1_druginfo_proto_rawDescData = file_api_druginfo_v1_druginfo_proto_rawDesc
)

func file_api_druginfo_v1_druginfo_proto_rawDescGZIP() []byte {
	file_api_druginfo_v1_druginfo_proto_rawDescOnce.Do(func() {
		file_api_druginfo_v1_druginfo_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_druginfo_v1_druginfo_proto_rawDescData)
	})
	return file_api_druginfo_v1_druginfo_proto_rawDescData
}

var file_api_druginfo_v1_druginfo_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_api_druginfo_v1_druginfo_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_druginfo_v1_druginfo_proto_goTypes = []interface{}{
	(ErrorReason)(0),              // 0: druginfo.v1.ErrorReason
	(*Druginfo)(nil),              // 1: druginfo.v1.Druginfo
	(*CreateDrugRequest)(nil),     // 2: druginfo.v1.CreateDrugRequest
	(*UpdatedruginfoRequest)(nil), // 3: druginfo.v1.UpdatedruginfoRequest
	(*UpdatedruginfoReply)(nil),   // 4: druginfo.v1.UpdatedruginfoReply
	(*GetDrugRequest)(nil),        // 5: druginfo.v1.GetDrugRequest
	(*GetDrugReply)(nil),          // 6: druginfo.v1.GetDrugReply
	(*ListdruginfoRequest)(nil),   // 7: druginfo.v1.ListdruginfoRequest
	(*ListdruginfoReply)(nil),     // 8: druginfo.v1.ListdruginfoReply
}
var file_api_druginfo_v1_druginfo_proto_depIdxs = []int32{
	1, // 0: druginfo.v1.UpdatedruginfoReply.druginfo:type_name -> druginfo.v1.Druginfo
	1, // 1: druginfo.v1.GetDrugReply.druginfo:type_name -> druginfo.v1.Druginfo
	1, // 2: druginfo.v1.ListdruginfoReply.results:type_name -> druginfo.v1.Druginfo
	2, // 3: druginfo.v1.Druginfos.CreateDrug:input_type -> druginfo.v1.CreateDrugRequest
	3, // 4: druginfo.v1.Druginfos.UpdateDrug:input_type -> druginfo.v1.UpdatedruginfoRequest
	5, // 5: druginfo.v1.Druginfos.GetDrug:input_type -> druginfo.v1.GetDrugRequest
	7, // 6: druginfo.v1.Druginfos.Listdruginfo:input_type -> druginfo.v1.ListdruginfoRequest
	6, // 7: druginfo.v1.Druginfos.CreateDrug:output_type -> druginfo.v1.GetDrugReply
	4, // 8: druginfo.v1.Druginfos.UpdateDrug:output_type -> druginfo.v1.UpdatedruginfoReply
	6, // 9: druginfo.v1.Druginfos.GetDrug:output_type -> druginfo.v1.GetDrugReply
	8, // 10: druginfo.v1.Druginfos.Listdruginfo:output_type -> druginfo.v1.ListdruginfoReply
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_api_druginfo_v1_druginfo_proto_init() }
func file_api_druginfo_v1_druginfo_proto_init() {
	if File_api_druginfo_v1_druginfo_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_druginfo_v1_druginfo_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Druginfo); i {
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
		file_api_druginfo_v1_druginfo_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDrugRequest); i {
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
		file_api_druginfo_v1_druginfo_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatedruginfoRequest); i {
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
		file_api_druginfo_v1_druginfo_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatedruginfoReply); i {
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
		file_api_druginfo_v1_druginfo_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDrugRequest); i {
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
		file_api_druginfo_v1_druginfo_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDrugReply); i {
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
		file_api_druginfo_v1_druginfo_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListdruginfoRequest); i {
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
		file_api_druginfo_v1_druginfo_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListdruginfoReply); i {
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
			RawDescriptor: file_api_druginfo_v1_druginfo_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_druginfo_v1_druginfo_proto_goTypes,
		DependencyIndexes: file_api_druginfo_v1_druginfo_proto_depIdxs,
		EnumInfos:         file_api_druginfo_v1_druginfo_proto_enumTypes,
		MessageInfos:      file_api_druginfo_v1_druginfo_proto_msgTypes,
	}.Build()
	File_api_druginfo_v1_druginfo_proto = out.File
	file_api_druginfo_v1_druginfo_proto_rawDesc = nil
	file_api_druginfo_v1_druginfo_proto_goTypes = nil
	file_api_druginfo_v1_druginfo_proto_depIdxs = nil
}

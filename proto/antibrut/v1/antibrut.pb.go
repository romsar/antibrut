// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: antibrut/v1/antibrut.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login    string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Ip       string `protobuf:"bytes,3,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *CheckRequest) Reset() {
	*x = CheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_antibrut_v1_antibrut_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckRequest) ProtoMessage() {}

func (x *CheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_antibrut_v1_antibrut_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckRequest.ProtoReflect.Descriptor instead.
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return file_antibrut_v1_antibrut_proto_rawDescGZIP(), []int{0}
}

func (x *CheckRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *CheckRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *CheckRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type CheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *CheckResponse) Reset() {
	*x = CheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_antibrut_v1_antibrut_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckResponse) ProtoMessage() {}

func (x *CheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_antibrut_v1_antibrut_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckResponse.ProtoReflect.Descriptor instead.
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return file_antibrut_v1_antibrut_proto_rawDescGZIP(), []int{1}
}

func (x *CheckResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type ResetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Login string `protobuf:"bytes,1,opt,name=login,proto3" json:"login,omitempty"`
	Ip    string `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *ResetRequest) Reset() {
	*x = ResetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_antibrut_v1_antibrut_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResetRequest) ProtoMessage() {}

func (x *ResetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_antibrut_v1_antibrut_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResetRequest.ProtoReflect.Descriptor instead.
func (*ResetRequest) Descriptor() ([]byte, []int) {
	return file_antibrut_v1_antibrut_proto_rawDescGZIP(), []int{2}
}

func (x *ResetRequest) GetLogin() string {
	if x != nil {
		return x.Login
	}
	return ""
}

func (x *ResetRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type AddIPToWhiteListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subnet string `protobuf:"bytes,1,opt,name=subnet,proto3" json:"subnet,omitempty"`
}

func (x *AddIPToWhiteListRequest) Reset() {
	*x = AddIPToWhiteListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_antibrut_v1_antibrut_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddIPToWhiteListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddIPToWhiteListRequest) ProtoMessage() {}

func (x *AddIPToWhiteListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_antibrut_v1_antibrut_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddIPToWhiteListRequest.ProtoReflect.Descriptor instead.
func (*AddIPToWhiteListRequest) Descriptor() ([]byte, []int) {
	return file_antibrut_v1_antibrut_proto_rawDescGZIP(), []int{3}
}

func (x *AddIPToWhiteListRequest) GetSubnet() string {
	if x != nil {
		return x.Subnet
	}
	return ""
}

type DeleteIPFromWhiteListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subnet string `protobuf:"bytes,1,opt,name=subnet,proto3" json:"subnet,omitempty"`
}

func (x *DeleteIPFromWhiteListRequest) Reset() {
	*x = DeleteIPFromWhiteListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_antibrut_v1_antibrut_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteIPFromWhiteListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteIPFromWhiteListRequest) ProtoMessage() {}

func (x *DeleteIPFromWhiteListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_antibrut_v1_antibrut_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteIPFromWhiteListRequest.ProtoReflect.Descriptor instead.
func (*DeleteIPFromWhiteListRequest) Descriptor() ([]byte, []int) {
	return file_antibrut_v1_antibrut_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteIPFromWhiteListRequest) GetSubnet() string {
	if x != nil {
		return x.Subnet
	}
	return ""
}

type AddIPToBlackListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subnet string `protobuf:"bytes,1,opt,name=subnet,proto3" json:"subnet,omitempty"`
}

func (x *AddIPToBlackListRequest) Reset() {
	*x = AddIPToBlackListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_antibrut_v1_antibrut_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddIPToBlackListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddIPToBlackListRequest) ProtoMessage() {}

func (x *AddIPToBlackListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_antibrut_v1_antibrut_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddIPToBlackListRequest.ProtoReflect.Descriptor instead.
func (*AddIPToBlackListRequest) Descriptor() ([]byte, []int) {
	return file_antibrut_v1_antibrut_proto_rawDescGZIP(), []int{5}
}

func (x *AddIPToBlackListRequest) GetSubnet() string {
	if x != nil {
		return x.Subnet
	}
	return ""
}

type DeleteIPFromBlackListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Subnet string `protobuf:"bytes,1,opt,name=subnet,proto3" json:"subnet,omitempty"`
}

func (x *DeleteIPFromBlackListRequest) Reset() {
	*x = DeleteIPFromBlackListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_antibrut_v1_antibrut_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteIPFromBlackListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteIPFromBlackListRequest) ProtoMessage() {}

func (x *DeleteIPFromBlackListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_antibrut_v1_antibrut_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteIPFromBlackListRequest.ProtoReflect.Descriptor instead.
func (*DeleteIPFromBlackListRequest) Descriptor() ([]byte, []int) {
	return file_antibrut_v1_antibrut_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteIPFromBlackListRequest) GetSubnet() string {
	if x != nil {
		return x.Subnet
	}
	return ""
}

var File_antibrut_v1_antibrut_proto protoreflect.FileDescriptor

var file_antibrut_v1_antibrut_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x61, 0x6e, 0x74, 0x69, 0x62, 0x72, 0x75, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x6e,
	0x74, 0x69, 0x62, 0x72, 0x75, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x61, 0x6e,
	0x74, 0x69, 0x62, 0x72, 0x75, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x50, 0x0a, 0x0c, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x1a, 0x0a, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x22, 0x1f, 0x0a, 0x0d, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x34, 0x0a, 0x0c, 0x52, 0x65, 0x73,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x6f, 0x67,
	0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x22,
	0x31, 0x0a, 0x17, 0x41, 0x64, 0x64, 0x49, 0x50, 0x54, 0x6f, 0x57, 0x68, 0x69, 0x74, 0x65, 0x4c,
	0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75,
	0x62, 0x6e, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x75, 0x62, 0x6e,
	0x65, 0x74, 0x22, 0x36, 0x0a, 0x1c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x50, 0x46, 0x72,
	0x6f, 0x6d, 0x57, 0x68, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x22, 0x31, 0x0a, 0x17, 0x41, 0x64,
	0x64, 0x49, 0x50, 0x54, 0x6f, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x22, 0x36, 0x0a,
	0x1c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x50, 0x46, 0x72, 0x6f, 0x6d, 0x42, 0x6c, 0x61,
	0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x75, 0x62, 0x6e, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x75, 0x62, 0x6e, 0x65, 0x74, 0x32, 0xe9, 0x03, 0x0a, 0x0f, 0x41, 0x6e, 0x74, 0x69, 0x42, 0x72,
	0x75, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3e, 0x0a, 0x05, 0x43, 0x68, 0x65,
	0x63, 0x6b, 0x12, 0x19, 0x2e, 0x61, 0x6e, 0x74, 0x69, 0x62, 0x72, 0x75, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e,
	0x61, 0x6e, 0x74, 0x69, 0x62, 0x72, 0x75, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3a, 0x0a, 0x05, 0x52, 0x65, 0x73,
	0x65, 0x74, 0x12, 0x19, 0x2e, 0x61, 0x6e, 0x74, 0x69, 0x62, 0x72, 0x75, 0x74, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x73, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x50, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x49, 0x50, 0x54, 0x6f,
	0x57, 0x68, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x2e, 0x61, 0x6e, 0x74, 0x69,
	0x62, 0x72, 0x75, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x50, 0x54, 0x6f, 0x57,
	0x68, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x5a, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x49, 0x50, 0x46, 0x72, 0x6f, 0x6d, 0x57, 0x68, 0x69, 0x74, 0x65, 0x4c, 0x69, 0x73, 0x74,
	0x12, 0x29, 0x2e, 0x61, 0x6e, 0x74, 0x69, 0x62, 0x72, 0x75, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x49, 0x50, 0x46, 0x72, 0x6f, 0x6d, 0x57, 0x68, 0x69, 0x74, 0x65,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x12, 0x50, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x49, 0x50, 0x54, 0x6f, 0x42, 0x6c,
	0x61, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x24, 0x2e, 0x61, 0x6e, 0x74, 0x69, 0x62, 0x72,
	0x75, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x50, 0x54, 0x6f, 0x42, 0x6c, 0x61,
	0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x5a, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x49,
	0x50, 0x46, 0x72, 0x6f, 0x6d, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x29,
	0x2e, 0x61, 0x6e, 0x74, 0x69, 0x62, 0x72, 0x75, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x49, 0x50, 0x46, 0x72, 0x6f, 0x6d, 0x42, 0x6c, 0x61, 0x63, 0x6b, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_antibrut_v1_antibrut_proto_rawDescOnce sync.Once
	file_antibrut_v1_antibrut_proto_rawDescData = file_antibrut_v1_antibrut_proto_rawDesc
)

func file_antibrut_v1_antibrut_proto_rawDescGZIP() []byte {
	file_antibrut_v1_antibrut_proto_rawDescOnce.Do(func() {
		file_antibrut_v1_antibrut_proto_rawDescData = protoimpl.X.CompressGZIP(file_antibrut_v1_antibrut_proto_rawDescData)
	})
	return file_antibrut_v1_antibrut_proto_rawDescData
}

var file_antibrut_v1_antibrut_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_antibrut_v1_antibrut_proto_goTypes = []interface{}{
	(*CheckRequest)(nil),                 // 0: antibrut.v1.CheckRequest
	(*CheckResponse)(nil),                // 1: antibrut.v1.CheckResponse
	(*ResetRequest)(nil),                 // 2: antibrut.v1.ResetRequest
	(*AddIPToWhiteListRequest)(nil),      // 3: antibrut.v1.AddIPToWhiteListRequest
	(*DeleteIPFromWhiteListRequest)(nil), // 4: antibrut.v1.DeleteIPFromWhiteListRequest
	(*AddIPToBlackListRequest)(nil),      // 5: antibrut.v1.AddIPToBlackListRequest
	(*DeleteIPFromBlackListRequest)(nil), // 6: antibrut.v1.DeleteIPFromBlackListRequest
	(*emptypb.Empty)(nil),                // 7: google.protobuf.Empty
}
var file_antibrut_v1_antibrut_proto_depIdxs = []int32{
	0, // 0: antibrut.v1.AntiBrutService.Check:input_type -> antibrut.v1.CheckRequest
	2, // 1: antibrut.v1.AntiBrutService.Reset:input_type -> antibrut.v1.ResetRequest
	3, // 2: antibrut.v1.AntiBrutService.AddIPToWhiteList:input_type -> antibrut.v1.AddIPToWhiteListRequest
	4, // 3: antibrut.v1.AntiBrutService.DeleteIPFromWhiteList:input_type -> antibrut.v1.DeleteIPFromWhiteListRequest
	5, // 4: antibrut.v1.AntiBrutService.AddIPToBlackList:input_type -> antibrut.v1.AddIPToBlackListRequest
	6, // 5: antibrut.v1.AntiBrutService.DeleteIPFromBlackList:input_type -> antibrut.v1.DeleteIPFromBlackListRequest
	1, // 6: antibrut.v1.AntiBrutService.Check:output_type -> antibrut.v1.CheckResponse
	7, // 7: antibrut.v1.AntiBrutService.Reset:output_type -> google.protobuf.Empty
	7, // 8: antibrut.v1.AntiBrutService.AddIPToWhiteList:output_type -> google.protobuf.Empty
	7, // 9: antibrut.v1.AntiBrutService.DeleteIPFromWhiteList:output_type -> google.protobuf.Empty
	7, // 10: antibrut.v1.AntiBrutService.AddIPToBlackList:output_type -> google.protobuf.Empty
	7, // 11: antibrut.v1.AntiBrutService.DeleteIPFromBlackList:output_type -> google.protobuf.Empty
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_antibrut_v1_antibrut_proto_init() }
func file_antibrut_v1_antibrut_proto_init() {
	if File_antibrut_v1_antibrut_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_antibrut_v1_antibrut_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckRequest); i {
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
		file_antibrut_v1_antibrut_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckResponse); i {
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
		file_antibrut_v1_antibrut_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResetRequest); i {
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
		file_antibrut_v1_antibrut_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddIPToWhiteListRequest); i {
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
		file_antibrut_v1_antibrut_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteIPFromWhiteListRequest); i {
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
		file_antibrut_v1_antibrut_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddIPToBlackListRequest); i {
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
		file_antibrut_v1_antibrut_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteIPFromBlackListRequest); i {
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
			RawDescriptor: file_antibrut_v1_antibrut_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_antibrut_v1_antibrut_proto_goTypes,
		DependencyIndexes: file_antibrut_v1_antibrut_proto_depIdxs,
		MessageInfos:      file_antibrut_v1_antibrut_proto_msgTypes,
	}.Build()
	File_antibrut_v1_antibrut_proto = out.File
	file_antibrut_v1_antibrut_proto_rawDesc = nil
	file_antibrut_v1_antibrut_proto_goTypes = nil
	file_antibrut_v1_antibrut_proto_depIdxs = nil
}

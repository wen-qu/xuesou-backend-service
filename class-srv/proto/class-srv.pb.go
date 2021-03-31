// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: proto/class-srv.proto

package classsrv

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Class struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgencyID  string  `protobuf:"bytes,1,opt,name=agencyID,proto3" json:"agencyID,omitempty"`
	ClassID   string  `protobuf:"bytes,2,opt,name=classID,proto3" json:"classID,omitempty"`
	Price     float32 `protobuf:"fixed32,3,opt,name=price,proto3" json:"price,omitempty"`
	Name      string  `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Age       string  `protobuf:"bytes,5,opt,name=age,proto3" json:"age,omitempty"`
	StuNumber int32   `protobuf:"varint,6,opt,name=stuNumber,proto3" json:"stuNumber,omitempty"`
	Level     string  `protobuf:"bytes,7,opt,name=level,proto3" json:"level,omitempty"`
	Sales     int32   `protobuf:"varint,8,opt,name=sales,proto3" json:"sales,omitempty"`
}

func (x *Class) Reset() {
	*x = Class{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_class_srv_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Class) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Class) ProtoMessage() {}

func (x *Class) ProtoReflect() protoreflect.Message {
	mi := &file_proto_class_srv_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Class.ProtoReflect.Descriptor instead.
func (*Class) Descriptor() ([]byte, []int) {
	return file_proto_class_srv_proto_rawDescGZIP(), []int{0}
}

func (x *Class) GetAgencyID() string {
	if x != nil {
		return x.AgencyID
	}
	return ""
}

func (x *Class) GetClassID() string {
	if x != nil {
		return x.ClassID
	}
	return ""
}

func (x *Class) GetPrice() float32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Class) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Class) GetAge() string {
	if x != nil {
		return x.Age
	}
	return ""
}

func (x *Class) GetStuNumber() int32 {
	if x != nil {
		return x.StuNumber
	}
	return 0
}

func (x *Class) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *Class) GetSales() int32 {
	if x != nil {
		return x.Sales
	}
	return 0
}

type ReadClassRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgencyID string `protobuf:"bytes,1,opt,name=agencyID,proto3" json:"agencyID,omitempty"`
}

func (x *ReadClassRequest) Reset() {
	*x = ReadClassRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_class_srv_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadClassRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadClassRequest) ProtoMessage() {}

func (x *ReadClassRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_class_srv_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadClassRequest.ProtoReflect.Descriptor instead.
func (*ReadClassRequest) Descriptor() ([]byte, []int) {
	return file_proto_class_srv_proto_rawDescGZIP(), []int{1}
}

func (x *ReadClassRequest) GetAgencyID() string {
	if x != nil {
		return x.AgencyID
	}
	return ""
}

type ReadClassResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Classes []*Class `protobuf:"bytes,2,rep,name=classes,proto3" json:"classes,omitempty"`
	Msg     string   `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *ReadClassResponse) Reset() {
	*x = ReadClassResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_class_srv_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadClassResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadClassResponse) ProtoMessage() {}

func (x *ReadClassResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_class_srv_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadClassResponse.ProtoReflect.Descriptor instead.
func (*ReadClassResponse) Descriptor() ([]byte, []int) {
	return file_proto_class_srv_proto_rawDescGZIP(), []int{2}
}

func (x *ReadClassResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *ReadClassResponse) GetClasses() []*Class {
	if x != nil {
		return x.Classes
	}
	return nil
}

func (x *ReadClassResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type AddClassRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Class *Class `protobuf:"bytes,1,opt,name=class,proto3" json:"class,omitempty"`
}

func (x *AddClassRequest) Reset() {
	*x = AddClassRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_class_srv_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddClassRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddClassRequest) ProtoMessage() {}

func (x *AddClassRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_class_srv_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddClassRequest.ProtoReflect.Descriptor instead.
func (*AddClassRequest) Descriptor() ([]byte, []int) {
	return file_proto_class_srv_proto_rawDescGZIP(), []int{3}
}

func (x *AddClassRequest) GetClass() *Class {
	if x != nil {
		return x.Class
	}
	return nil
}

type AddClassResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	ClassID string `protobuf:"bytes,2,opt,name=classID,proto3" json:"classID,omitempty"`
	Msg     string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *AddClassResponse) Reset() {
	*x = AddClassResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_class_srv_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddClassResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddClassResponse) ProtoMessage() {}

func (x *AddClassResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_class_srv_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddClassResponse.ProtoReflect.Descriptor instead.
func (*AddClassResponse) Descriptor() ([]byte, []int) {
	return file_proto_class_srv_proto_rawDescGZIP(), []int{4}
}

func (x *AddClassResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *AddClassResponse) GetClassID() string {
	if x != nil {
		return x.ClassID
	}
	return ""
}

func (x *AddClassResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type UpdateClassRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Class *Class `protobuf:"bytes,1,opt,name=class,proto3" json:"class,omitempty"`
}

func (x *UpdateClassRequest) Reset() {
	*x = UpdateClassRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_class_srv_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateClassRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateClassRequest) ProtoMessage() {}

func (x *UpdateClassRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_class_srv_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateClassRequest.ProtoReflect.Descriptor instead.
func (*UpdateClassRequest) Descriptor() ([]byte, []int) {
	return file_proto_class_srv_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateClassRequest) GetClass() *Class {
	if x != nil {
		return x.Class
	}
	return nil
}

type UpdateClassResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *UpdateClassResponse) Reset() {
	*x = UpdateClassResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_class_srv_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateClassResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateClassResponse) ProtoMessage() {}

func (x *UpdateClassResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_class_srv_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateClassResponse.ProtoReflect.Descriptor instead.
func (*UpdateClassResponse) Descriptor() ([]byte, []int) {
	return file_proto_class_srv_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateClassResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *UpdateClassResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type DeleteClassRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AgencyID string `protobuf:"bytes,1,opt,name=agencyID,proto3" json:"agencyID,omitempty"`
	ClassID  string `protobuf:"bytes,2,opt,name=classID,proto3" json:"classID,omitempty"`
}

func (x *DeleteClassRequest) Reset() {
	*x = DeleteClassRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_class_srv_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteClassRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteClassRequest) ProtoMessage() {}

func (x *DeleteClassRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_class_srv_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteClassRequest.ProtoReflect.Descriptor instead.
func (*DeleteClassRequest) Descriptor() ([]byte, []int) {
	return file_proto_class_srv_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteClassRequest) GetAgencyID() string {
	if x != nil {
		return x.AgencyID
	}
	return ""
}

func (x *DeleteClassRequest) GetClassID() string {
	if x != nil {
		return x.ClassID
	}
	return ""
}

type DeleteClassResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status int32  `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *DeleteClassResponse) Reset() {
	*x = DeleteClassResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_class_srv_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteClassResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteClassResponse) ProtoMessage() {}

func (x *DeleteClassResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_class_srv_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteClassResponse.ProtoReflect.Descriptor instead.
func (*DeleteClassResponse) Descriptor() ([]byte, []int) {
	return file_proto_class_srv_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteClassResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *DeleteClassResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_proto_class_srv_proto protoreflect.FileDescriptor

var file_proto_class_srv_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x2d, 0x73, 0x72,
	0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x73, 0x72,
	0x76, 0x22, 0xc3, 0x01, 0x0a, 0x05, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x61,
	0x67, 0x65, 0x6e, 0x63, 0x79, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61,
	0x67, 0x65, 0x6e, 0x63, 0x79, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x49,
	0x44, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61,
	0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x74, 0x75, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x09, 0x73, 0x74, 0x75, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6c,
	0x65, 0x76, 0x65, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x22, 0x2e, 0x0a, 0x10, 0x52, 0x65, 0x61, 0x64, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x61,
	0x67, 0x65, 0x6e, 0x63, 0x79, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61,
	0x67, 0x65, 0x6e, 0x63, 0x79, 0x49, 0x44, 0x22, 0x68, 0x0a, 0x11, 0x52, 0x65, 0x61, 0x64, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x29, 0x0a, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x73, 0x72, 0x76,
	0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x22, 0x38, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x73, 0x72, 0x76, 0x2e, 0x43,
	0x6c, 0x61, 0x73, 0x73, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x22, 0x56, 0x0a, 0x10, 0x41,
	0x64, 0x64, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x49,
	0x44, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x22, 0x3b, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61,
	0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x05, 0x63, 0x6c, 0x61,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x73, 0x72, 0x76, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x05, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x22, 0x3f, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73,
	0x67, 0x22, 0x4a, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x63,
	0x79, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61, 0x67, 0x65, 0x6e, 0x63,
	0x79, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x44, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x44, 0x22, 0x3f, 0x0a,
	0x13, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0xc1,
	0x02, 0x0a, 0x08, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x53, 0x72, 0x76, 0x12, 0x52, 0x0a, 0x15, 0x52,
	0x65, 0x61, 0x64, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x42, 0x79, 0x41, 0x67, 0x65, 0x6e,
	0x63, 0x79, 0x49, 0x44, 0x12, 0x1a, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x73, 0x72, 0x76, 0x2e,
	0x52, 0x65, 0x61, 0x64, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x73, 0x72, 0x76, 0x2e, 0x52, 0x65, 0x61, 0x64,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x45, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x65, 0x73, 0x12, 0x19, 0x2e,
	0x63, 0x6c, 0x61, 0x73, 0x73, 0x73, 0x72, 0x76, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6c, 0x61, 0x73,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73,
	0x73, 0x72, 0x76, 0x2e, 0x41, 0x64, 0x64, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x6c, 0x61, 0x73, 0x73, 0x12, 0x1c, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x73, 0x72, 0x76,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x73, 0x72, 0x76, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c,
	0x61, 0x73, 0x73, 0x12, 0x1c, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x73, 0x72, 0x76, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x73, 0x72, 0x76, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x10, 0x5a, 0x0e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x63, 0x6c, 0x61, 0x73,
	0x73, 0x73, 0x72, 0x76, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_class_srv_proto_rawDescOnce sync.Once
	file_proto_class_srv_proto_rawDescData = file_proto_class_srv_proto_rawDesc
)

func file_proto_class_srv_proto_rawDescGZIP() []byte {
	file_proto_class_srv_proto_rawDescOnce.Do(func() {
		file_proto_class_srv_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_class_srv_proto_rawDescData)
	})
	return file_proto_class_srv_proto_rawDescData
}

var file_proto_class_srv_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_class_srv_proto_goTypes = []interface{}{
	(*Class)(nil),               // 0: classsrv.Class
	(*ReadClassRequest)(nil),    // 1: classsrv.ReadClassRequest
	(*ReadClassResponse)(nil),   // 2: classsrv.ReadClassResponse
	(*AddClassRequest)(nil),     // 3: classsrv.AddClassRequest
	(*AddClassResponse)(nil),    // 4: classsrv.AddClassResponse
	(*UpdateClassRequest)(nil),  // 5: classsrv.UpdateClassRequest
	(*UpdateClassResponse)(nil), // 6: classsrv.UpdateClassResponse
	(*DeleteClassRequest)(nil),  // 7: classsrv.DeleteClassRequest
	(*DeleteClassResponse)(nil), // 8: classsrv.DeleteClassResponse
}
var file_proto_class_srv_proto_depIdxs = []int32{
	0, // 0: classsrv.ReadClassResponse.classes:type_name -> classsrv.Class
	0, // 1: classsrv.AddClassRequest.class:type_name -> classsrv.Class
	0, // 2: classsrv.UpdateClassRequest.class:type_name -> classsrv.Class
	1, // 3: classsrv.ClassSrv.ReadClassesByAgencyID:input_type -> classsrv.ReadClassRequest
	3, // 4: classsrv.ClassSrv.AddClasses:input_type -> classsrv.AddClassRequest
	5, // 5: classsrv.ClassSrv.UpdateClass:input_type -> classsrv.UpdateClassRequest
	7, // 6: classsrv.ClassSrv.DeleteClass:input_type -> classsrv.DeleteClassRequest
	2, // 7: classsrv.ClassSrv.ReadClassesByAgencyID:output_type -> classsrv.ReadClassResponse
	4, // 8: classsrv.ClassSrv.AddClasses:output_type -> classsrv.AddClassResponse
	6, // 9: classsrv.ClassSrv.UpdateClass:output_type -> classsrv.UpdateClassResponse
	8, // 10: classsrv.ClassSrv.DeleteClass:output_type -> classsrv.DeleteClassResponse
	7, // [7:11] is the sub-list for method output_type
	3, // [3:7] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_class_srv_proto_init() }
func file_proto_class_srv_proto_init() {
	if File_proto_class_srv_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_class_srv_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Class); i {
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
		file_proto_class_srv_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadClassRequest); i {
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
		file_proto_class_srv_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadClassResponse); i {
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
		file_proto_class_srv_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddClassRequest); i {
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
		file_proto_class_srv_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddClassResponse); i {
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
		file_proto_class_srv_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateClassRequest); i {
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
		file_proto_class_srv_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateClassResponse); i {
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
		file_proto_class_srv_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteClassRequest); i {
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
		file_proto_class_srv_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteClassResponse); i {
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
			RawDescriptor: file_proto_class_srv_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_class_srv_proto_goTypes,
		DependencyIndexes: file_proto_class_srv_proto_depIdxs,
		MessageInfos:      file_proto_class_srv_proto_msgTypes,
	}.Build()
	File_proto_class_srv_proto = out.File
	file_proto_class_srv_proto_rawDesc = nil
	file_proto_class_srv_proto_goTypes = nil
	file_proto_class_srv_proto_depIdxs = nil
}

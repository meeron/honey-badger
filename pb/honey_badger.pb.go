// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.2
// source: pb/honey_badger.proto

package pb

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

type SetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Db   string `protobuf:"bytes,1,opt,name=db,proto3" json:"db,omitempty"`
	Key  string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Data []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Ttl  *int32 `protobuf:"varint,4,opt,name=ttl,proto3,oneof" json:"ttl,omitempty"`
}

func (x *SetRequest) Reset() {
	*x = SetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_honey_badger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRequest) ProtoMessage() {}

func (x *SetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_honey_badger_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetRequest.ProtoReflect.Descriptor instead.
func (*SetRequest) Descriptor() ([]byte, []int) {
	return file_pb_honey_badger_proto_rawDescGZIP(), []int{0}
}

func (x *SetRequest) GetDb() string {
	if x != nil {
		return x.Db
	}
	return ""
}

func (x *SetRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SetRequest) GetTtl() int32 {
	if x != nil && x.Ttl != nil {
		return *x.Ttl
	}
	return 0
}

type KeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Db  string `protobuf:"bytes,1,opt,name=db,proto3" json:"db,omitempty"`
	Key string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *KeyRequest) Reset() {
	*x = KeyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_honey_badger_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyRequest) ProtoMessage() {}

func (x *KeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_honey_badger_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KeyRequest.ProtoReflect.Descriptor instead.
func (*KeyRequest) Descriptor() ([]byte, []int) {
	return file_pb_honey_badger_proto_rawDescGZIP(), []int{1}
}

func (x *KeyRequest) GetDb() string {
	if x != nil {
		return x.Db
	}
	return ""
}

func (x *KeyRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type GetResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hit  bool   `protobuf:"varint,1,opt,name=hit,proto3" json:"hit,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *GetResult) Reset() {
	*x = GetResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_honey_badger_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResult) ProtoMessage() {}

func (x *GetResult) ProtoReflect() protoreflect.Message {
	mi := &file_pb_honey_badger_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResult.ProtoReflect.Descriptor instead.
func (*GetResult) Descriptor() ([]byte, []int) {
	return file_pb_honey_badger_proto_rawDescGZIP(), []int{2}
}

func (x *GetResult) GetHit() bool {
	if x != nil {
		return x.Hit
	}
	return false
}

func (x *GetResult) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  string  `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Error *string `protobuf:"bytes,2,opt,name=error,proto3,oneof" json:"error,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_honey_badger_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_pb_honey_badger_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_pb_honey_badger_proto_rawDescGZIP(), []int{3}
}

func (x *Result) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Result) GetError() string {
	if x != nil && x.Error != nil {
		return *x.Error
	}
	return ""
}

type PrefixRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Db     string `protobuf:"bytes,1,opt,name=db,proto3" json:"db,omitempty"`
	Prefix string `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
}

func (x *PrefixRequest) Reset() {
	*x = PrefixRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_honey_badger_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrefixRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrefixRequest) ProtoMessage() {}

func (x *PrefixRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_honey_badger_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrefixRequest.ProtoReflect.Descriptor instead.
func (*PrefixRequest) Descriptor() ([]byte, []int) {
	return file_pb_honey_badger_proto_rawDescGZIP(), []int{4}
}

func (x *PrefixRequest) GetDb() string {
	if x != nil {
		return x.Db
	}
	return ""
}

func (x *PrefixRequest) GetPrefix() string {
	if x != nil {
		return x.Prefix
	}
	return ""
}

type PrefixResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data map[string][]byte `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PrefixResult) Reset() {
	*x = PrefixResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_honey_badger_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrefixResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrefixResult) ProtoMessage() {}

func (x *PrefixResult) ProtoReflect() protoreflect.Message {
	mi := &file_pb_honey_badger_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PrefixResult.ProtoReflect.Descriptor instead.
func (*PrefixResult) Descriptor() ([]byte, []int) {
	return file_pb_honey_badger_proto_rawDescGZIP(), []int{5}
}

func (x *PrefixResult) GetData() map[string][]byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type SetBatchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Db   string            `protobuf:"bytes,1,opt,name=db,proto3" json:"db,omitempty"`
	Data map[string][]byte `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SetBatchRequest) Reset() {
	*x = SetBatchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_honey_badger_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetBatchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetBatchRequest) ProtoMessage() {}

func (x *SetBatchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_pb_honey_badger_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetBatchRequest.ProtoReflect.Descriptor instead.
func (*SetBatchRequest) Descriptor() ([]byte, []int) {
	return file_pb_honey_badger_proto_rawDescGZIP(), []int{6}
}

func (x *SetBatchRequest) GetDb() string {
	if x != nil {
		return x.Db
	}
	return ""
}

func (x *SetBatchRequest) GetData() map[string][]byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_pb_honey_badger_proto protoreflect.FileDescriptor

var file_pb_honey_badger_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x62, 0x2f, 0x68, 0x6f, 0x6e, 0x65, 0x79, 0x5f, 0x62, 0x61, 0x64, 0x67, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x61, 0x0a, 0x0a, 0x53,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x62, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x64, 0x62, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12,
	0x15, 0x0a, 0x03, 0x74, 0x74, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x03,
	0x74, 0x74, 0x6c, 0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x74, 0x74, 0x6c, 0x22, 0x2e,
	0x0a, 0x0a, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02,
	0x64, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x64, 0x62, 0x12, 0x10, 0x0a, 0x03,
	0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x31,
	0x0a, 0x09, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x68,
	0x69, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x68, 0x69, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x41, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12,
	0x19, 0x0a, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00,
	0x52, 0x05, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x88, 0x01, 0x01, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x65,
	0x72, 0x72, 0x6f, 0x72, 0x22, 0x37, 0x0a, 0x0d, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x64, 0x62, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x22, 0x77, 0x0a,
	0x0c, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2e, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x62,
	0x2e, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x37, 0x0a,
	0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x8d, 0x01, 0x0a, 0x0f, 0x53, 0x65, 0x74, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x62,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x64, 0x62, 0x12, 0x31, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65,
	0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x37, 0x0a,
	0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x9a, 0x02, 0x0a, 0x0b, 0x48, 0x6f, 0x6e, 0x65, 0x79,
	0x42, 0x61, 0x64, 0x67, 0x65, 0x72, 0x12, 0x23, 0x0a, 0x03, 0x53, 0x65, 0x74, 0x12, 0x0e, 0x2e,
	0x70, 0x62, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e,
	0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x42, 0x79, 0x50, 0x72, 0x65, 0x66,
	0x69, 0x78, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x65, 0x66, 0x69,
	0x78, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x06, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x00, 0x12, 0x31, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x42, 0x79, 0x50, 0x72, 0x65,
	0x66, 0x69, 0x78, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x00, 0x12, 0x2d, 0x0a, 0x08, 0x53, 0x65, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68,
	0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x65, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0a, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x00, 0x42, 0x23, 0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6d, 0x65, 0x65, 0x72, 0x6f, 0x6e, 0x2f, 0x68, 0x6f, 0x6e, 0x65, 0x79, 0x2d, 0x62,
	0x61, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_honey_badger_proto_rawDescOnce sync.Once
	file_pb_honey_badger_proto_rawDescData = file_pb_honey_badger_proto_rawDesc
)

func file_pb_honey_badger_proto_rawDescGZIP() []byte {
	file_pb_honey_badger_proto_rawDescOnce.Do(func() {
		file_pb_honey_badger_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_honey_badger_proto_rawDescData)
	})
	return file_pb_honey_badger_proto_rawDescData
}

var file_pb_honey_badger_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_pb_honey_badger_proto_goTypes = []interface{}{
	(*SetRequest)(nil),      // 0: pb.SetRequest
	(*KeyRequest)(nil),      // 1: pb.KeyRequest
	(*GetResult)(nil),       // 2: pb.GetResult
	(*Result)(nil),          // 3: pb.Result
	(*PrefixRequest)(nil),   // 4: pb.PrefixRequest
	(*PrefixResult)(nil),    // 5: pb.PrefixResult
	(*SetBatchRequest)(nil), // 6: pb.SetBatchRequest
	nil,                     // 7: pb.PrefixResult.DataEntry
	nil,                     // 8: pb.SetBatchRequest.DataEntry
}
var file_pb_honey_badger_proto_depIdxs = []int32{
	7, // 0: pb.PrefixResult.data:type_name -> pb.PrefixResult.DataEntry
	8, // 1: pb.SetBatchRequest.data:type_name -> pb.SetBatchRequest.DataEntry
	0, // 2: pb.HoneyBadger.Set:input_type -> pb.SetRequest
	1, // 3: pb.HoneyBadger.Get:input_type -> pb.KeyRequest
	4, // 4: pb.HoneyBadger.GetByPrefix:input_type -> pb.PrefixRequest
	1, // 5: pb.HoneyBadger.Delete:input_type -> pb.KeyRequest
	4, // 6: pb.HoneyBadger.DeleteByPrefix:input_type -> pb.PrefixRequest
	6, // 7: pb.HoneyBadger.SetBatch:input_type -> pb.SetBatchRequest
	3, // 8: pb.HoneyBadger.Set:output_type -> pb.Result
	2, // 9: pb.HoneyBadger.Get:output_type -> pb.GetResult
	5, // 10: pb.HoneyBadger.GetByPrefix:output_type -> pb.PrefixResult
	3, // 11: pb.HoneyBadger.Delete:output_type -> pb.Result
	3, // 12: pb.HoneyBadger.DeleteByPrefix:output_type -> pb.Result
	3, // 13: pb.HoneyBadger.SetBatch:output_type -> pb.Result
	8, // [8:14] is the sub-list for method output_type
	2, // [2:8] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pb_honey_badger_proto_init() }
func file_pb_honey_badger_proto_init() {
	if File_pb_honey_badger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_honey_badger_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetRequest); i {
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
		file_pb_honey_badger_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KeyRequest); i {
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
		file_pb_honey_badger_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResult); i {
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
		file_pb_honey_badger_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_pb_honey_badger_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrefixRequest); i {
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
		file_pb_honey_badger_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PrefixResult); i {
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
		file_pb_honey_badger_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetBatchRequest); i {
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
	file_pb_honey_badger_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_pb_honey_badger_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_honey_badger_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_honey_badger_proto_goTypes,
		DependencyIndexes: file_pb_honey_badger_proto_depIdxs,
		MessageInfos:      file_pb_honey_badger_proto_msgTypes,
	}.Build()
	File_pb_honey_badger_proto = out.File
	file_pb_honey_badger_proto_rawDesc = nil
	file_pb_honey_badger_proto_goTypes = nil
	file_pb_honey_badger_proto_depIdxs = nil
}
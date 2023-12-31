// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: honey_badger.proto

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
		mi := &file_honey_badger_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetRequest) ProtoMessage() {}

func (x *SetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_honey_badger_proto_msgTypes[0]
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
	return file_honey_badger_proto_rawDescGZIP(), []int{0}
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
		mi := &file_honey_badger_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KeyRequest) ProtoMessage() {}

func (x *KeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_honey_badger_proto_msgTypes[1]
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
	return file_honey_badger_proto_rawDescGZIP(), []int{1}
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
		mi := &file_honey_badger_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResult) ProtoMessage() {}

func (x *GetResult) ProtoReflect() protoreflect.Message {
	mi := &file_honey_badger_proto_msgTypes[2]
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
	return file_honey_badger_proto_rawDescGZIP(), []int{2}
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
		mi := &file_honey_badger_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PrefixRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PrefixRequest) ProtoMessage() {}

func (x *PrefixRequest) ProtoReflect() protoreflect.Message {
	mi := &file_honey_badger_proto_msgTypes[3]
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
	return file_honey_badger_proto_rawDescGZIP(), []int{3}
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

type CreateDbRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	InMemory bool   `protobuf:"varint,2,opt,name=inMemory,proto3" json:"inMemory,omitempty"`
}

func (x *CreateDbRequest) Reset() {
	*x = CreateDbRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_honey_badger_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateDbRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateDbRequest) ProtoMessage() {}

func (x *CreateDbRequest) ProtoReflect() protoreflect.Message {
	mi := &file_honey_badger_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateDbRequest.ProtoReflect.Descriptor instead.
func (*CreateDbRequest) Descriptor() ([]byte, []int) {
	return file_honey_badger_proto_rawDescGZIP(), []int{4}
}

func (x *CreateDbRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateDbRequest) GetInMemory() bool {
	if x != nil {
		return x.InMemory
	}
	return false
}

type DropDbRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DropDbRequest) Reset() {
	*x = DropDbRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_honey_badger_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DropDbRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DropDbRequest) ProtoMessage() {}

func (x *DropDbRequest) ProtoReflect() protoreflect.Message {
	mi := &file_honey_badger_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DropDbRequest.ProtoReflect.Descriptor instead.
func (*DropDbRequest) Descriptor() ([]byte, []int) {
	return file_honey_badger_proto_rawDescGZIP(), []int{5}
}

func (x *DropDbRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_honey_badger_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_honey_badger_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_honey_badger_proto_rawDescGZIP(), []int{6}
}

type PingResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mesage string `protobuf:"bytes,1,opt,name=mesage,proto3" json:"mesage,omitempty"`
}

func (x *PingResult) Reset() {
	*x = PingResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_honey_badger_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingResult) ProtoMessage() {}

func (x *PingResult) ProtoReflect() protoreflect.Message {
	mi := &file_honey_badger_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingResult.ProtoReflect.Descriptor instead.
func (*PingResult) Descriptor() ([]byte, []int) {
	return file_honey_badger_proto_rawDescGZIP(), []int{7}
}

func (x *PingResult) GetMesage() string {
	if x != nil {
		return x.Mesage
	}
	return ""
}

type EmptyResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyResult) Reset() {
	*x = EmptyResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_honey_badger_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyResult) ProtoMessage() {}

func (x *EmptyResult) ProtoReflect() protoreflect.Message {
	mi := &file_honey_badger_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyResult.ProtoReflect.Descriptor instead.
func (*EmptyResult) Descriptor() ([]byte, []int) {
	return file_honey_badger_proto_rawDescGZIP(), []int{8}
}

type ReadStreamReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Db     string  `protobuf:"bytes,1,opt,name=db,proto3" json:"db,omitempty"`
	Prefix *string `protobuf:"bytes,2,opt,name=prefix,proto3,oneof" json:"prefix,omitempty"`
}

func (x *ReadStreamReq) Reset() {
	*x = ReadStreamReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_honey_badger_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadStreamReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadStreamReq) ProtoMessage() {}

func (x *ReadStreamReq) ProtoReflect() protoreflect.Message {
	mi := &file_honey_badger_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadStreamReq.ProtoReflect.Descriptor instead.
func (*ReadStreamReq) Descriptor() ([]byte, []int) {
	return file_honey_badger_proto_rawDescGZIP(), []int{9}
}

func (x *ReadStreamReq) GetDb() string {
	if x != nil {
		return x.Db
	}
	return ""
}

func (x *ReadStreamReq) GetPrefix() string {
	if x != nil && x.Prefix != nil {
		return *x.Prefix
	}
	return ""
}

type SendStreamReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Item *DataItem `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	Db   string    `protobuf:"bytes,2,opt,name=db,proto3" json:"db,omitempty"`
}

func (x *SendStreamReq) Reset() {
	*x = SendStreamReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_honey_badger_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendStreamReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendStreamReq) ProtoMessage() {}

func (x *SendStreamReq) ProtoReflect() protoreflect.Message {
	mi := &file_honey_badger_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendStreamReq.ProtoReflect.Descriptor instead.
func (*SendStreamReq) Descriptor() ([]byte, []int) {
	return file_honey_badger_proto_rawDescGZIP(), []int{10}
}

func (x *SendStreamReq) GetItem() *DataItem {
	if x != nil {
		return x.Item
	}
	return nil
}

func (x *SendStreamReq) GetDb() string {
	if x != nil {
		return x.Db
	}
	return ""
}

type DataItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key  string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Data []byte `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *DataItem) Reset() {
	*x = DataItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_honey_badger_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataItem) ProtoMessage() {}

func (x *DataItem) ProtoReflect() protoreflect.Message {
	mi := &file_honey_badger_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataItem.ProtoReflect.Descriptor instead.
func (*DataItem) Descriptor() ([]byte, []int) {
	return file_honey_badger_proto_rawDescGZIP(), []int{11}
}

func (x *DataItem) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *DataItem) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_honey_badger_proto protoreflect.FileDescriptor

var file_honey_badger_proto_rawDesc = []byte{
	0x0a, 0x12, 0x68, 0x6f, 0x6e, 0x65, 0x79, 0x5f, 0x62, 0x61, 0x64, 0x67, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x68, 0x62, 0x22, 0x61, 0x0a, 0x0a, 0x53, 0x65, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x62, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x64, 0x62, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x15, 0x0a, 0x03,
	0x74, 0x74, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x03, 0x74, 0x74, 0x6c,
	0x88, 0x01, 0x01, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x74, 0x74, 0x6c, 0x22, 0x2e, 0x0a, 0x0a, 0x4b,
	0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x62, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x64, 0x62, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x31, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x68, 0x69, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x68, 0x69, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x37,
	0x0a, 0x0d, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x64, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x64, 0x62, 0x12,
	0x16, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x22, 0x41, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x44, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x69, 0x6e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x08, 0x69, 0x6e, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x22, 0x23, 0x0a, 0x0d, 0x44, 0x72,
	0x6f, 0x70, 0x44, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x0d, 0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x24,
	0x0a, 0x0a, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x0a, 0x06,
	0x6d, 0x65, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x0d, 0x0a, 0x0b, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x22, 0x47, 0x0a, 0x0d, 0x52, 0x65, 0x61, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x64, 0x62, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x64, 0x62, 0x12, 0x1b, 0x0a, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x88, 0x01,
	0x01, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x70, 0x72, 0x65, 0x66, 0x69, 0x78, 0x22, 0x41, 0x0a, 0x0d,
	0x53, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x12, 0x20, 0x0a,
	0x04, 0x69, 0x74, 0x65, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x68, 0x62,
	0x2e, 0x44, 0x61, 0x74, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x04, 0x69, 0x74, 0x65, 0x6d, 0x12,
	0x0e, 0x0a, 0x02, 0x64, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x64, 0x62, 0x22,
	0x30, 0x0a, 0x08, 0x44, 0x61, 0x74, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x32, 0xb2, 0x02, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x28, 0x0a, 0x03, 0x53, 0x65,
	0x74, 0x12, 0x0e, 0x2e, 0x68, 0x62, 0x2e, 0x53, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x0f, 0x2e, 0x68, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x00, 0x12, 0x26, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x0e, 0x2e, 0x68, 0x62,
	0x2e, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x68, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x06,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x0e, 0x2e, 0x68, 0x62, 0x2e, 0x4b, 0x65, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x68, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x36, 0x0a, 0x0e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x42, 0x79, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x12, 0x11, 0x2e, 0x68, 0x62,
	0x2e, 0x50, 0x72, 0x65, 0x66, 0x69, 0x78, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f,
	0x2e, 0x68, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22,
	0x00, 0x12, 0x37, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x61, 0x64, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x11, 0x2e, 0x68, 0x62, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x53,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x71, 0x1a, 0x0c, 0x2e, 0x68, 0x62, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x49, 0x74, 0x65, 0x6d, 0x22, 0x00, 0x30, 0x01, 0x12, 0x3a, 0x0a, 0x10, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x11,
	0x2e, 0x68, 0x62, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x71, 0x1a, 0x0f, 0x2e, 0x68, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x22, 0x00, 0x28, 0x01, 0x32, 0x64, 0x0a, 0x02, 0x44, 0x62, 0x12, 0x30, 0x0a, 0x06,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x13, 0x2e, 0x68, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x44, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x68, 0x62,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x2c,
	0x0a, 0x04, 0x44, 0x72, 0x6f, 0x70, 0x12, 0x11, 0x2e, 0x68, 0x62, 0x2e, 0x44, 0x72, 0x6f, 0x70,
	0x44, 0x62, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x68, 0x62, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x32, 0x30, 0x0a, 0x03,
	0x53, 0x79, 0x73, 0x12, 0x29, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x12, 0x0f, 0x2e, 0x68, 0x62,
	0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x68,
	0x62, 0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x42, 0x3b,
	0x5a, 0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x65, 0x65,
	0x72, 0x6f, 0x6e, 0x2f, 0x68, 0x6f, 0x6e, 0x65, 0x79, 0x2d, 0x62, 0x61, 0x64, 0x67, 0x65, 0x72,
	0x2f, 0x70, 0x62, 0xaa, 0x02, 0x15, 0x48, 0x6f, 0x6e, 0x65, 0x79, 0x42, 0x61, 0x64, 0x67, 0x65,
	0x72, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e, 0x48, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_honey_badger_proto_rawDescOnce sync.Once
	file_honey_badger_proto_rawDescData = file_honey_badger_proto_rawDesc
)

func file_honey_badger_proto_rawDescGZIP() []byte {
	file_honey_badger_proto_rawDescOnce.Do(func() {
		file_honey_badger_proto_rawDescData = protoimpl.X.CompressGZIP(file_honey_badger_proto_rawDescData)
	})
	return file_honey_badger_proto_rawDescData
}

var file_honey_badger_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_honey_badger_proto_goTypes = []interface{}{
	(*SetRequest)(nil),      // 0: hb.SetRequest
	(*KeyRequest)(nil),      // 1: hb.KeyRequest
	(*GetResult)(nil),       // 2: hb.GetResult
	(*PrefixRequest)(nil),   // 3: hb.PrefixRequest
	(*CreateDbRequest)(nil), // 4: hb.CreateDbRequest
	(*DropDbRequest)(nil),   // 5: hb.DropDbRequest
	(*PingRequest)(nil),     // 6: hb.PingRequest
	(*PingResult)(nil),      // 7: hb.PingResult
	(*EmptyResult)(nil),     // 8: hb.EmptyResult
	(*ReadStreamReq)(nil),   // 9: hb.ReadStreamReq
	(*SendStreamReq)(nil),   // 10: hb.SendStreamReq
	(*DataItem)(nil),        // 11: hb.DataItem
}
var file_honey_badger_proto_depIdxs = []int32{
	11, // 0: hb.SendStreamReq.item:type_name -> hb.DataItem
	0,  // 1: hb.Data.Set:input_type -> hb.SetRequest
	1,  // 2: hb.Data.Get:input_type -> hb.KeyRequest
	1,  // 3: hb.Data.Delete:input_type -> hb.KeyRequest
	3,  // 4: hb.Data.DeleteByPrefix:input_type -> hb.PrefixRequest
	9,  // 5: hb.Data.CreateReadStream:input_type -> hb.ReadStreamReq
	10, // 6: hb.Data.CreateSendStream:input_type -> hb.SendStreamReq
	4,  // 7: hb.Db.Create:input_type -> hb.CreateDbRequest
	5,  // 8: hb.Db.Drop:input_type -> hb.DropDbRequest
	6,  // 9: hb.Sys.Ping:input_type -> hb.PingRequest
	8,  // 10: hb.Data.Set:output_type -> hb.EmptyResult
	2,  // 11: hb.Data.Get:output_type -> hb.GetResult
	8,  // 12: hb.Data.Delete:output_type -> hb.EmptyResult
	8,  // 13: hb.Data.DeleteByPrefix:output_type -> hb.EmptyResult
	11, // 14: hb.Data.CreateReadStream:output_type -> hb.DataItem
	8,  // 15: hb.Data.CreateSendStream:output_type -> hb.EmptyResult
	8,  // 16: hb.Db.Create:output_type -> hb.EmptyResult
	8,  // 17: hb.Db.Drop:output_type -> hb.EmptyResult
	7,  // 18: hb.Sys.Ping:output_type -> hb.PingResult
	10, // [10:19] is the sub-list for method output_type
	1,  // [1:10] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_honey_badger_proto_init() }
func file_honey_badger_proto_init() {
	if File_honey_badger_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_honey_badger_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_honey_badger_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_honey_badger_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_honey_badger_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_honey_badger_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateDbRequest); i {
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
		file_honey_badger_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DropDbRequest); i {
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
		file_honey_badger_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_honey_badger_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingResult); i {
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
		file_honey_badger_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyResult); i {
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
		file_honey_badger_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadStreamReq); i {
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
		file_honey_badger_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendStreamReq); i {
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
		file_honey_badger_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataItem); i {
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
	file_honey_badger_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_honey_badger_proto_msgTypes[9].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_honey_badger_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   3,
		},
		GoTypes:           file_honey_badger_proto_goTypes,
		DependencyIndexes: file_honey_badger_proto_depIdxs,
		MessageInfos:      file_honey_badger_proto_msgTypes,
	}.Build()
	File_honey_badger_proto = out.File
	file_honey_badger_proto_rawDesc = nil
	file_honey_badger_proto_goTypes = nil
	file_honey_badger_proto_depIdxs = nil
}

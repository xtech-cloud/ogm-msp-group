// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.7
// source: proto/group/collection.proto

package group

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

// 创建集合的请求
type CollectionMakeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`          // 名称
	Capacity uint64 `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty"` // 成员上限，0为无限制
}

func (x *CollectionMakeRequest) Reset() {
	*x = CollectionMakeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_group_collection_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectionMakeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectionMakeRequest) ProtoMessage() {}

func (x *CollectionMakeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_group_collection_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectionMakeRequest.ProtoReflect.Descriptor instead.
func (*CollectionMakeRequest) Descriptor() ([]byte, []int) {
	return file_proto_group_collection_proto_rawDescGZIP(), []int{0}
}

func (x *CollectionMakeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CollectionMakeRequest) GetCapacity() uint64 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

// 更新集合的请求
type CollectionUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid     string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`          // uuid
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`          // 名称
	Capacity uint64 `protobuf:"varint,3,opt,name=capacity,proto3" json:"capacity,omitempty"` // 成员上限，0为无限制
}

func (x *CollectionUpdateRequest) Reset() {
	*x = CollectionUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_group_collection_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectionUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectionUpdateRequest) ProtoMessage() {}

func (x *CollectionUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_group_collection_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectionUpdateRequest.ProtoReflect.Descriptor instead.
func (*CollectionUpdateRequest) Descriptor() ([]byte, []int) {
	return file_proto_group_collection_proto_rawDescGZIP(), []int{1}
}

func (x *CollectionUpdateRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *CollectionUpdateRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CollectionUpdateRequest) GetCapacity() uint64 {
	if x != nil {
		return x.Capacity
	}
	return 0
}

// 列举集合的请求
type CollectionListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64 `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"` // 偏移值
	Count  int64 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`   // 数量
}

func (x *CollectionListRequest) Reset() {
	*x = CollectionListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_group_collection_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectionListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectionListRequest) ProtoMessage() {}

func (x *CollectionListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_group_collection_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectionListRequest.ProtoReflect.Descriptor instead.
func (*CollectionListRequest) Descriptor() ([]byte, []int) {
	return file_proto_group_collection_proto_rawDescGZIP(), []int{2}
}

func (x *CollectionListRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *CollectionListRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

// 搜索集合的请求
type CollectionSearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"` // 偏移值
	Count  int64  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`   // 数量
	Name   string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`      // 名称
}

func (x *CollectionSearchRequest) Reset() {
	*x = CollectionSearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_group_collection_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectionSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectionSearchRequest) ProtoMessage() {}

func (x *CollectionSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_group_collection_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectionSearchRequest.ProtoReflect.Descriptor instead.
func (*CollectionSearchRequest) Descriptor() ([]byte, []int) {
	return file_proto_group_collection_proto_rawDescGZIP(), []int{3}
}

func (x *CollectionSearchRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *CollectionSearchRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *CollectionSearchRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// 列举集合的回复
type CollectionListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status             `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` // 状态
	Total  uint64              `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`  // 总数
	Entity []*CollectionEntity `protobuf:"bytes,3,rep,name=entity,proto3" json:"entity,omitempty"` // 实体列表
}

func (x *CollectionListResponse) Reset() {
	*x = CollectionListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_group_collection_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectionListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectionListResponse) ProtoMessage() {}

func (x *CollectionListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_group_collection_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectionListResponse.ProtoReflect.Descriptor instead.
func (*CollectionListResponse) Descriptor() ([]byte, []int) {
	return file_proto_group_collection_proto_rawDescGZIP(), []int{4}
}

func (x *CollectionListResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *CollectionListResponse) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *CollectionListResponse) GetEntity() []*CollectionEntity {
	if x != nil {
		return x.Entity
	}
	return nil
}

// 删除集合的请求
type CollectionRemoveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"` // uuid
}

func (x *CollectionRemoveRequest) Reset() {
	*x = CollectionRemoveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_group_collection_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectionRemoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectionRemoveRequest) ProtoMessage() {}

func (x *CollectionRemoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_group_collection_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectionRemoveRequest.ProtoReflect.Descriptor instead.
func (*CollectionRemoveRequest) Descriptor() ([]byte, []int) {
	return file_proto_group_collection_proto_rawDescGZIP(), []int{5}
}

func (x *CollectionRemoveRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

// 获取集合信息的请求
type CollectionGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"` // uuid
}

func (x *CollectionGetRequest) Reset() {
	*x = CollectionGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_group_collection_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectionGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectionGetRequest) ProtoMessage() {}

func (x *CollectionGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_group_collection_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectionGetRequest.ProtoReflect.Descriptor instead.
func (*CollectionGetRequest) Descriptor() ([]byte, []int) {
	return file_proto_group_collection_proto_rawDescGZIP(), []int{6}
}

func (x *CollectionGetRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

// 获取集合信息的回复
type CollectionGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status           `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` // 状态
	Entity *CollectionEntity `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"` // 实体
}

func (x *CollectionGetResponse) Reset() {
	*x = CollectionGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_group_collection_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CollectionGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CollectionGetResponse) ProtoMessage() {}

func (x *CollectionGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_group_collection_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CollectionGetResponse.ProtoReflect.Descriptor instead.
func (*CollectionGetResponse) Descriptor() ([]byte, []int) {
	return file_proto_group_collection_proto_rawDescGZIP(), []int{7}
}

func (x *CollectionGetResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *CollectionGetResponse) GetEntity() *CollectionEntity {
	if x != nil {
		return x.Entity
	}
	return nil
}

var File_proto_group_collection_proto protoreflect.FileDescriptor

var file_proto_group_collection_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x63, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x67, 0x72, 0x6f, 0x75, 0x70, 0x1a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x6f,
	0x75, 0x70, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x47, 0x0a, 0x15, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x6b,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08,
	0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08,
	0x63, 0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x22, 0x5d, 0x0a, 0x17, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63,
	0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x63,
	0x61, 0x70, 0x61, 0x63, 0x69, 0x74, 0x79, 0x22, 0x45, 0x0a, 0x15, 0x43, 0x6f, 0x6c, 0x6c, 0x65,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x5b,
	0x0a, 0x17, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x86, 0x01, 0x0a, 0x16,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x05, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x12, 0x2f, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x22, 0x2d, 0x0a, 0x17, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x22, 0x2a, 0x0a, 0x14, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22,
	0x6f, 0x0a, 0x15, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x2f, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x17, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x32, 0xa1, 0x03, 0x0a, 0x0a, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x3b, 0x0a, 0x04, 0x4d, 0x61, 0x6b, 0x65, 0x12, 0x1c, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4d, 0x61, 0x6b, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x55, 0x75,
	0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x06,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x43,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x55,
	0x75, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x43, 0x6f,
	0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x43, 0x6f, 0x6c, 0x6c,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1e,
	0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3f, 0x0a, 0x06, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x1e, 0x2e, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x2e, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x2e, 0x55, 0x75, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x42, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1b, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e,
	0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x3b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_proto_group_collection_proto_rawDescOnce sync.Once
	file_proto_group_collection_proto_rawDescData = file_proto_group_collection_proto_rawDesc
)

func file_proto_group_collection_proto_rawDescGZIP() []byte {
	file_proto_group_collection_proto_rawDescOnce.Do(func() {
		file_proto_group_collection_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_group_collection_proto_rawDescData)
	})
	return file_proto_group_collection_proto_rawDescData
}

var file_proto_group_collection_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_proto_group_collection_proto_goTypes = []interface{}{
	(*CollectionMakeRequest)(nil),   // 0: group.CollectionMakeRequest
	(*CollectionUpdateRequest)(nil), // 1: group.CollectionUpdateRequest
	(*CollectionListRequest)(nil),   // 2: group.CollectionListRequest
	(*CollectionSearchRequest)(nil), // 3: group.CollectionSearchRequest
	(*CollectionListResponse)(nil),  // 4: group.CollectionListResponse
	(*CollectionRemoveRequest)(nil), // 5: group.CollectionRemoveRequest
	(*CollectionGetRequest)(nil),    // 6: group.CollectionGetRequest
	(*CollectionGetResponse)(nil),   // 7: group.CollectionGetResponse
	(*Status)(nil),                  // 8: group.Status
	(*CollectionEntity)(nil),        // 9: group.CollectionEntity
	(*UuidResponse)(nil),            // 10: group.UuidResponse
}
var file_proto_group_collection_proto_depIdxs = []int32{
	8,  // 0: group.CollectionListResponse.status:type_name -> group.Status
	9,  // 1: group.CollectionListResponse.entity:type_name -> group.CollectionEntity
	8,  // 2: group.CollectionGetResponse.status:type_name -> group.Status
	9,  // 3: group.CollectionGetResponse.entity:type_name -> group.CollectionEntity
	0,  // 4: group.Collection.Make:input_type -> group.CollectionMakeRequest
	1,  // 5: group.Collection.Update:input_type -> group.CollectionUpdateRequest
	2,  // 6: group.Collection.List:input_type -> group.CollectionListRequest
	3,  // 7: group.Collection.Search:input_type -> group.CollectionSearchRequest
	5,  // 8: group.Collection.Remove:input_type -> group.CollectionRemoveRequest
	6,  // 9: group.Collection.Get:input_type -> group.CollectionGetRequest
	10, // 10: group.Collection.Make:output_type -> group.UuidResponse
	10, // 11: group.Collection.Update:output_type -> group.UuidResponse
	4,  // 12: group.Collection.List:output_type -> group.CollectionListResponse
	4,  // 13: group.Collection.Search:output_type -> group.CollectionListResponse
	10, // 14: group.Collection.Remove:output_type -> group.UuidResponse
	7,  // 15: group.Collection.Get:output_type -> group.CollectionGetResponse
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_proto_group_collection_proto_init() }
func file_proto_group_collection_proto_init() {
	if File_proto_group_collection_proto != nil {
		return
	}
	file_proto_group_shared_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_group_collection_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectionMakeRequest); i {
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
		file_proto_group_collection_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectionUpdateRequest); i {
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
		file_proto_group_collection_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectionListRequest); i {
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
		file_proto_group_collection_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectionSearchRequest); i {
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
		file_proto_group_collection_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectionListResponse); i {
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
		file_proto_group_collection_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectionRemoveRequest); i {
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
		file_proto_group_collection_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectionGetRequest); i {
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
		file_proto_group_collection_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CollectionGetResponse); i {
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
			RawDescriptor: file_proto_group_collection_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_group_collection_proto_goTypes,
		DependencyIndexes: file_proto_group_collection_proto_depIdxs,
		MessageInfos:      file_proto_group_collection_proto_msgTypes,
	}.Build()
	File_proto_group_collection_proto = out.File
	file_proto_group_collection_proto_rawDesc = nil
	file_proto_group_collection_proto_goTypes = nil
	file_proto_group_collection_proto_depIdxs = nil
}

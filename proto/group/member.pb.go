// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.7
// source: proto/group/member.proto

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

// 加入成员的请求
type MemberAddRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Collection string `protobuf:"bytes,1,opt,name=collection,proto3" json:"collection,omitempty"` // 集合UUID
	Element    string `protobuf:"bytes,2,opt,name=element,proto3" json:"element,omitempty"`       // 元素
	Alias      string `protobuf:"bytes,3,opt,name=alias,proto3" json:"alias,omitempty"`           // 别名
}

func (x *MemberAddRequest) Reset() {
	*x = MemberAddRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_group_member_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberAddRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberAddRequest) ProtoMessage() {}

func (x *MemberAddRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_group_member_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberAddRequest.ProtoReflect.Descriptor instead.
func (*MemberAddRequest) Descriptor() ([]byte, []int) {
	return file_proto_group_member_proto_rawDescGZIP(), []int{0}
}

func (x *MemberAddRequest) GetCollection() string {
	if x != nil {
		return x.Collection
	}
	return ""
}

func (x *MemberAddRequest) GetElement() string {
	if x != nil {
		return x.Element
	}
	return ""
}

func (x *MemberAddRequest) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

// 更新成员的请求
type MemberUpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid    string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`       // 成员UUID
	Element string `protobuf:"bytes,2,opt,name=element,proto3" json:"element,omitempty"` // 元素
	Alias   string `protobuf:"bytes,3,opt,name=alias,proto3" json:"alias,omitempty"`     // 别名
}

func (x *MemberUpdateRequest) Reset() {
	*x = MemberUpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_group_member_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberUpdateRequest) ProtoMessage() {}

func (x *MemberUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_group_member_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberUpdateRequest.ProtoReflect.Descriptor instead.
func (*MemberUpdateRequest) Descriptor() ([]byte, []int) {
	return file_proto_group_member_proto_rawDescGZIP(), []int{1}
}

func (x *MemberUpdateRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *MemberUpdateRequest) GetElement() string {
	if x != nil {
		return x.Element
	}
	return ""
}

func (x *MemberUpdateRequest) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

// 列举成员的请求
type MemberListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset     int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`        // 偏移值
	Count      int64  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`          // 数量
	Collection string `protobuf:"bytes,3,opt,name=collection,proto3" json:"collection,omitempty"` // 集合UUID, 空值时列举所有集合
}

func (x *MemberListRequest) Reset() {
	*x = MemberListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_group_member_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberListRequest) ProtoMessage() {}

func (x *MemberListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_group_member_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberListRequest.ProtoReflect.Descriptor instead.
func (*MemberListRequest) Descriptor() ([]byte, []int) {
	return file_proto_group_member_proto_rawDescGZIP(), []int{2}
}

func (x *MemberListRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *MemberListRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *MemberListRequest) GetCollection() string {
	if x != nil {
		return x.Collection
	}
	return ""
}

// 搜索成员的请求
type MemberSearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset     int64  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`        // 偏移值
	Count      int64  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`          // 数量
	Collection string `protobuf:"bytes,3,opt,name=collection,proto3" json:"collection,omitempty"` // 集合UUID, 空值时列举所有集合
	Element    string `protobuf:"bytes,4,opt,name=element,proto3" json:"element,omitempty"`       // 元素
	Alias      string `protobuf:"bytes,5,opt,name=alias,proto3" json:"alias,omitempty"`           // 别名
}

func (x *MemberSearchRequest) Reset() {
	*x = MemberSearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_group_member_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberSearchRequest) ProtoMessage() {}

func (x *MemberSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_group_member_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberSearchRequest.ProtoReflect.Descriptor instead.
func (*MemberSearchRequest) Descriptor() ([]byte, []int) {
	return file_proto_group_member_proto_rawDescGZIP(), []int{3}
}

func (x *MemberSearchRequest) GetOffset() int64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *MemberSearchRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *MemberSearchRequest) GetCollection() string {
	if x != nil {
		return x.Collection
	}
	return ""
}

func (x *MemberSearchRequest) GetElement() string {
	if x != nil {
		return x.Element
	}
	return ""
}

func (x *MemberSearchRequest) GetAlias() string {
	if x != nil {
		return x.Alias
	}
	return ""
}

// 列举成员的回复
type MemberListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status         `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` // 状态
	Total  uint64          `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`  // 总数
	Entity []*MemberEntity `protobuf:"bytes,3,rep,name=entity,proto3" json:"entity,omitempty"` // 实体列表
}

func (x *MemberListResponse) Reset() {
	*x = MemberListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_group_member_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberListResponse) ProtoMessage() {}

func (x *MemberListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_group_member_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberListResponse.ProtoReflect.Descriptor instead.
func (*MemberListResponse) Descriptor() ([]byte, []int) {
	return file_proto_group_member_proto_rawDescGZIP(), []int{4}
}

func (x *MemberListResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *MemberListResponse) GetTotal() uint64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *MemberListResponse) GetEntity() []*MemberEntity {
	if x != nil {
		return x.Entity
	}
	return nil
}

// 移除成员的请求
type MemberRemoveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"` // 成员UUID
}

func (x *MemberRemoveRequest) Reset() {
	*x = MemberRemoveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_group_member_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberRemoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberRemoveRequest) ProtoMessage() {}

func (x *MemberRemoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_group_member_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberRemoveRequest.ProtoReflect.Descriptor instead.
func (*MemberRemoveRequest) Descriptor() ([]byte, []int) {
	return file_proto_group_member_proto_rawDescGZIP(), []int{5}
}

func (x *MemberRemoveRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

// 获取成员的请求
type MemberGetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid string `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"` // 成员uuid
}

func (x *MemberGetRequest) Reset() {
	*x = MemberGetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_group_member_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberGetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberGetRequest) ProtoMessage() {}

func (x *MemberGetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_group_member_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberGetRequest.ProtoReflect.Descriptor instead.
func (*MemberGetRequest) Descriptor() ([]byte, []int) {
	return file_proto_group_member_proto_rawDescGZIP(), []int{6}
}

func (x *MemberGetRequest) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

// 获取成员信息的回复
type MemberGetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status *Status       `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"` // 状态
	Entity *MemberEntity `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"` // 实体
}

func (x *MemberGetResponse) Reset() {
	*x = MemberGetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_group_member_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberGetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberGetResponse) ProtoMessage() {}

func (x *MemberGetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_group_member_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberGetResponse.ProtoReflect.Descriptor instead.
func (*MemberGetResponse) Descriptor() ([]byte, []int) {
	return file_proto_group_member_proto_rawDescGZIP(), []int{7}
}

func (x *MemberGetResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *MemberGetResponse) GetEntity() *MemberEntity {
	if x != nil {
		return x.Entity
	}
	return nil
}

// 获取成员所在的集合的请求
type MemberWhereRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Element string `protobuf:"bytes,1,opt,name=element,proto3" json:"element,omitempty"` // 元素
}

func (x *MemberWhereRequest) Reset() {
	*x = MemberWhereRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_group_member_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberWhereRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberWhereRequest) ProtoMessage() {}

func (x *MemberWhereRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_group_member_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberWhereRequest.ProtoReflect.Descriptor instead.
func (*MemberWhereRequest) Descriptor() ([]byte, []int) {
	return file_proto_group_member_proto_rawDescGZIP(), []int{8}
}

func (x *MemberWhereRequest) GetElement() string {
	if x != nil {
		return x.Element
	}
	return ""
}

// 获取成员所在的集合的回复
type MemberWhereResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status     *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`         // 状态
	Collection []string `protobuf:"bytes,2,rep,name=collection,proto3" json:"collection,omitempty"` // 集合列表
}

func (x *MemberWhereResponse) Reset() {
	*x = MemberWhereResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_group_member_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberWhereResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberWhereResponse) ProtoMessage() {}

func (x *MemberWhereResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_group_member_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberWhereResponse.ProtoReflect.Descriptor instead.
func (*MemberWhereResponse) Descriptor() ([]byte, []int) {
	return file_proto_group_member_proto_rawDescGZIP(), []int{9}
}

func (x *MemberWhereResponse) GetStatus() *Status {
	if x != nil {
		return x.Status
	}
	return nil
}

func (x *MemberWhereResponse) GetCollection() []string {
	if x != nil {
		return x.Collection
	}
	return nil
}

var File_proto_group_member_proto protoreflect.FileDescriptor

var file_proto_group_member_proto_rawDesc = []byte{
	0x0a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x6d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x1a, 0x18, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2f, 0x73,
	0x68, 0x61, 0x72, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x62, 0x0a, 0x10, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x64, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x18, 0x0a, 0x07, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69,
	0x61, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x22,
	0x59, 0x0a, 0x13, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6c,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6c, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x22, 0x61, 0x0a, 0x11, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x93, 0x01,
	0x0a, 0x13, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x61, 0x6c, 0x69, 0x61, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x61, 0x6c,
	0x69, 0x61, 0x73, 0x22, 0x7e, 0x0a, 0x12, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2b, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x4d,
	0x65, 0x6d, 0x62, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x22, 0x29, 0x0a, 0x13, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x26,
	0x0a, 0x10, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x67, 0x0a, 0x11, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x2b, 0x0a, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x06, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22,
	0x2e, 0x0a, 0x12, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x57, 0x68, 0x65, 0x72, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x65, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22,
	0x5c, 0x0a, 0x13, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x57, 0x68, 0x65, 0x72, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a,
	0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0xb9, 0x03,
	0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x35, 0x0a, 0x03, 0x41, 0x64, 0x64, 0x12,
	0x17, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x41, 0x64,
	0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2e, 0x55, 0x75, 0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x3b, 0x0a, 0x06, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x67, 0x72, 0x6f, 0x75,
	0x70, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x55, 0x75,
	0x69, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3d, 0x0a, 0x04,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x4d, 0x65, 0x6d,
	0x62, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x06, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x1a, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b,
	0x0a, 0x06, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x12, 0x1a, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70,
	0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x55, 0x75, 0x69,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x03, 0x47,
	0x65, 0x74, 0x12, 0x17, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x40, 0x0a, 0x05, 0x57, 0x68, 0x65, 0x72, 0x65,
	0x12, 0x19, 0x2e, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x57,
	0x68, 0x65, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x72,
	0x6f, 0x75, 0x70, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x57, 0x68, 0x65, 0x72, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x3b, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_group_member_proto_rawDescOnce sync.Once
	file_proto_group_member_proto_rawDescData = file_proto_group_member_proto_rawDesc
)

func file_proto_group_member_proto_rawDescGZIP() []byte {
	file_proto_group_member_proto_rawDescOnce.Do(func() {
		file_proto_group_member_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_group_member_proto_rawDescData)
	})
	return file_proto_group_member_proto_rawDescData
}

var file_proto_group_member_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_proto_group_member_proto_goTypes = []interface{}{
	(*MemberAddRequest)(nil),    // 0: group.MemberAddRequest
	(*MemberUpdateRequest)(nil), // 1: group.MemberUpdateRequest
	(*MemberListRequest)(nil),   // 2: group.MemberListRequest
	(*MemberSearchRequest)(nil), // 3: group.MemberSearchRequest
	(*MemberListResponse)(nil),  // 4: group.MemberListResponse
	(*MemberRemoveRequest)(nil), // 5: group.MemberRemoveRequest
	(*MemberGetRequest)(nil),    // 6: group.MemberGetRequest
	(*MemberGetResponse)(nil),   // 7: group.MemberGetResponse
	(*MemberWhereRequest)(nil),  // 8: group.MemberWhereRequest
	(*MemberWhereResponse)(nil), // 9: group.MemberWhereResponse
	(*Status)(nil),              // 10: group.Status
	(*MemberEntity)(nil),        // 11: group.MemberEntity
	(*UuidResponse)(nil),        // 12: group.UuidResponse
}
var file_proto_group_member_proto_depIdxs = []int32{
	10, // 0: group.MemberListResponse.status:type_name -> group.Status
	11, // 1: group.MemberListResponse.entity:type_name -> group.MemberEntity
	10, // 2: group.MemberGetResponse.status:type_name -> group.Status
	11, // 3: group.MemberGetResponse.entity:type_name -> group.MemberEntity
	10, // 4: group.MemberWhereResponse.status:type_name -> group.Status
	0,  // 5: group.Member.Add:input_type -> group.MemberAddRequest
	1,  // 6: group.Member.Update:input_type -> group.MemberUpdateRequest
	2,  // 7: group.Member.List:input_type -> group.MemberListRequest
	3,  // 8: group.Member.Search:input_type -> group.MemberSearchRequest
	5,  // 9: group.Member.Remove:input_type -> group.MemberRemoveRequest
	6,  // 10: group.Member.Get:input_type -> group.MemberGetRequest
	8,  // 11: group.Member.Where:input_type -> group.MemberWhereRequest
	12, // 12: group.Member.Add:output_type -> group.UuidResponse
	12, // 13: group.Member.Update:output_type -> group.UuidResponse
	4,  // 14: group.Member.List:output_type -> group.MemberListResponse
	4,  // 15: group.Member.Search:output_type -> group.MemberListResponse
	12, // 16: group.Member.Remove:output_type -> group.UuidResponse
	7,  // 17: group.Member.Get:output_type -> group.MemberGetResponse
	9,  // 18: group.Member.Where:output_type -> group.MemberWhereResponse
	12, // [12:19] is the sub-list for method output_type
	5,  // [5:12] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_proto_group_member_proto_init() }
func file_proto_group_member_proto_init() {
	if File_proto_group_member_proto != nil {
		return
	}
	file_proto_group_shared_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_proto_group_member_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberAddRequest); i {
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
		file_proto_group_member_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberUpdateRequest); i {
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
		file_proto_group_member_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberListRequest); i {
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
		file_proto_group_member_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberSearchRequest); i {
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
		file_proto_group_member_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberListResponse); i {
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
		file_proto_group_member_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberRemoveRequest); i {
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
		file_proto_group_member_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberGetRequest); i {
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
		file_proto_group_member_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberGetResponse); i {
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
		file_proto_group_member_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberWhereRequest); i {
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
		file_proto_group_member_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberWhereResponse); i {
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
			RawDescriptor: file_proto_group_member_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_group_member_proto_goTypes,
		DependencyIndexes: file_proto_group_member_proto_depIdxs,
		MessageInfos:      file_proto_group_member_proto_msgTypes,
	}.Build()
	File_proto_group_member_proto = out.File
	file_proto_group_member_proto_rawDesc = nil
	file_proto_group_member_proto_goTypes = nil
	file_proto_group_member_proto_depIdxs = nil
}

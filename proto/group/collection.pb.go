// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/group/collection.proto

package group

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// 创建集合的请求
type CollectionMakeRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Capacity             uint64   `protobuf:"varint,2,opt,name=capacity,proto3" json:"capacity,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollectionMakeRequest) Reset()         { *m = CollectionMakeRequest{} }
func (m *CollectionMakeRequest) String() string { return proto.CompactTextString(m) }
func (*CollectionMakeRequest) ProtoMessage()    {}
func (*CollectionMakeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fbc8a1969e648d9, []int{0}
}

func (m *CollectionMakeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionMakeRequest.Unmarshal(m, b)
}
func (m *CollectionMakeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionMakeRequest.Marshal(b, m, deterministic)
}
func (m *CollectionMakeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionMakeRequest.Merge(m, src)
}
func (m *CollectionMakeRequest) XXX_Size() int {
	return xxx_messageInfo_CollectionMakeRequest.Size(m)
}
func (m *CollectionMakeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionMakeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionMakeRequest proto.InternalMessageInfo

func (m *CollectionMakeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CollectionMakeRequest) GetCapacity() uint64 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

// 列举集合的回复
type CollectionListRequest struct {
	Offset               int64    `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Count                int64    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollectionListRequest) Reset()         { *m = CollectionListRequest{} }
func (m *CollectionListRequest) String() string { return proto.CompactTextString(m) }
func (*CollectionListRequest) ProtoMessage()    {}
func (*CollectionListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fbc8a1969e648d9, []int{1}
}

func (m *CollectionListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionListRequest.Unmarshal(m, b)
}
func (m *CollectionListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionListRequest.Marshal(b, m, deterministic)
}
func (m *CollectionListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionListRequest.Merge(m, src)
}
func (m *CollectionListRequest) XXX_Size() int {
	return xxx_messageInfo_CollectionListRequest.Size(m)
}
func (m *CollectionListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionListRequest proto.InternalMessageInfo

func (m *CollectionListRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *CollectionListRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

// 列举集合的回复
type CollectionListResponse struct {
	Status               *Status             `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Total                uint64              `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Entity               []*CollectionEntity `protobuf:"bytes,3,rep,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *CollectionListResponse) Reset()         { *m = CollectionListResponse{} }
func (m *CollectionListResponse) String() string { return proto.CompactTextString(m) }
func (*CollectionListResponse) ProtoMessage()    {}
func (*CollectionListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fbc8a1969e648d9, []int{2}
}

func (m *CollectionListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionListResponse.Unmarshal(m, b)
}
func (m *CollectionListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionListResponse.Marshal(b, m, deterministic)
}
func (m *CollectionListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionListResponse.Merge(m, src)
}
func (m *CollectionListResponse) XXX_Size() int {
	return xxx_messageInfo_CollectionListResponse.Size(m)
}
func (m *CollectionListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionListResponse proto.InternalMessageInfo

func (m *CollectionListResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CollectionListResponse) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *CollectionListResponse) GetEntity() []*CollectionEntity {
	if m != nil {
		return m.Entity
	}
	return nil
}

// 删除集合的请求
type CollectionRemoveRequest struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollectionRemoveRequest) Reset()         { *m = CollectionRemoveRequest{} }
func (m *CollectionRemoveRequest) String() string { return proto.CompactTextString(m) }
func (*CollectionRemoveRequest) ProtoMessage()    {}
func (*CollectionRemoveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fbc8a1969e648d9, []int{3}
}

func (m *CollectionRemoveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionRemoveRequest.Unmarshal(m, b)
}
func (m *CollectionRemoveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionRemoveRequest.Marshal(b, m, deterministic)
}
func (m *CollectionRemoveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionRemoveRequest.Merge(m, src)
}
func (m *CollectionRemoveRequest) XXX_Size() int {
	return xxx_messageInfo_CollectionRemoveRequest.Size(m)
}
func (m *CollectionRemoveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionRemoveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionRemoveRequest proto.InternalMessageInfo

func (m *CollectionRemoveRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

// 获取集合信息的请求
type CollectionGetRequest struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CollectionGetRequest) Reset()         { *m = CollectionGetRequest{} }
func (m *CollectionGetRequest) String() string { return proto.CompactTextString(m) }
func (*CollectionGetRequest) ProtoMessage()    {}
func (*CollectionGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fbc8a1969e648d9, []int{4}
}

func (m *CollectionGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionGetRequest.Unmarshal(m, b)
}
func (m *CollectionGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionGetRequest.Marshal(b, m, deterministic)
}
func (m *CollectionGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionGetRequest.Merge(m, src)
}
func (m *CollectionGetRequest) XXX_Size() int {
	return xxx_messageInfo_CollectionGetRequest.Size(m)
}
func (m *CollectionGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionGetRequest proto.InternalMessageInfo

func (m *CollectionGetRequest) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

// 获取集合信息的回复
type CollectionGetResponse struct {
	Status               *Status           `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Entity               *CollectionEntity `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CollectionGetResponse) Reset()         { *m = CollectionGetResponse{} }
func (m *CollectionGetResponse) String() string { return proto.CompactTextString(m) }
func (*CollectionGetResponse) ProtoMessage()    {}
func (*CollectionGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7fbc8a1969e648d9, []int{5}
}

func (m *CollectionGetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CollectionGetResponse.Unmarshal(m, b)
}
func (m *CollectionGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CollectionGetResponse.Marshal(b, m, deterministic)
}
func (m *CollectionGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CollectionGetResponse.Merge(m, src)
}
func (m *CollectionGetResponse) XXX_Size() int {
	return xxx_messageInfo_CollectionGetResponse.Size(m)
}
func (m *CollectionGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CollectionGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CollectionGetResponse proto.InternalMessageInfo

func (m *CollectionGetResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *CollectionGetResponse) GetEntity() *CollectionEntity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func init() {
	proto.RegisterType((*CollectionMakeRequest)(nil), "group.CollectionMakeRequest")
	proto.RegisterType((*CollectionListRequest)(nil), "group.CollectionListRequest")
	proto.RegisterType((*CollectionListResponse)(nil), "group.CollectionListResponse")
	proto.RegisterType((*CollectionRemoveRequest)(nil), "group.CollectionRemoveRequest")
	proto.RegisterType((*CollectionGetRequest)(nil), "group.CollectionGetRequest")
	proto.RegisterType((*CollectionGetResponse)(nil), "group.CollectionGetResponse")
}

func init() {
	proto.RegisterFile("proto/group/collection.proto", fileDescriptor_7fbc8a1969e648d9)
}

var fileDescriptor_7fbc8a1969e648d9 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0xcd, 0x4e, 0xc2, 0x40,
	0x10, 0xb6, 0xb4, 0x34, 0x3a, 0xc4, 0xcb, 0x06, 0x81, 0x54, 0x34, 0xa4, 0x89, 0x09, 0x31, 0x11,
	0x12, 0xbc, 0x7a, 0x30, 0x18, 0xc2, 0x45, 0x2f, 0xeb, 0x13, 0xac, 0x65, 0xd1, 0x86, 0xd2, 0xad,
	0xec, 0xd4, 0xc4, 0x17, 0xf0, 0x15, 0x7c, 0x5d, 0xd3, 0x69, 0x4b, 0x57, 0x0b, 0x1a, 0x6f, 0x3b,
	0xf3, 0xfd, 0x30, 0xf3, 0x0d, 0x85, 0x7e, 0xb2, 0x51, 0xa8, 0xc6, 0xcf, 0x1b, 0x95, 0x26, 0xe3,
	0x40, 0x45, 0x91, 0x0c, 0x30, 0x54, 0xf1, 0x88, 0xda, 0xac, 0x49, 0x7d, 0xaf, 0x67, 0x92, 0xf4,
	0x8b, 0xd8, 0xc8, 0x45, 0x4e, 0xf0, 0xe7, 0x70, 0x72, 0xb7, 0x15, 0x3d, 0x88, 0x95, 0xe4, 0xf2,
	0x35, 0x95, 0x1a, 0x19, 0x03, 0x27, 0x16, 0x6b, 0xd9, 0xb3, 0x06, 0xd6, 0xf0, 0x88, 0xd3, 0x9b,
	0x79, 0x70, 0x18, 0x88, 0x44, 0x04, 0x21, 0xbe, 0xf7, 0x1a, 0x03, 0x6b, 0xe8, 0xf0, 0x6d, 0xed,
	0xcf, 0x4c, 0xa3, 0xfb, 0x50, 0x63, 0x69, 0xd4, 0x01, 0x57, 0x2d, 0x97, 0x5a, 0x22, 0x59, 0xd9,
	0xbc, 0xa8, 0x58, 0x1b, 0x9a, 0x81, 0x4a, 0x63, 0x24, 0x27, 0x9b, 0xe7, 0x85, 0xff, 0x61, 0x41,
	0xe7, 0xa7, 0x8f, 0x4e, 0x54, 0xac, 0x25, 0xbb, 0x00, 0x57, 0xa3, 0xc0, 0x54, 0x93, 0x51, 0x6b,
	0x72, 0x3c, 0xa2, 0x7d, 0x46, 0x8f, 0xd4, 0xe4, 0x05, 0x98, 0xf9, 0xa2, 0x42, 0x11, 0x15, 0x13,
	0xe6, 0x05, 0x1b, 0x83, 0x2b, 0x63, 0xcc, 0x06, 0xb7, 0x07, 0xf6, 0xb0, 0x35, 0xe9, 0x16, 0xe2,
	0xea, 0xb7, 0x66, 0x04, 0xf3, 0x82, 0xe6, 0x5f, 0x41, 0xb7, 0xc2, 0xb8, 0x5c, 0xab, 0x37, 0x33,
	0x9a, 0x34, 0x0d, 0x17, 0x65, 0x34, 0xd9, 0xdb, 0xbf, 0x84, 0x76, 0x45, 0x9f, 0x4b, 0xfc, 0x8d,
	0xab, 0xcc, 0xa8, 0x88, 0xfb, 0xbf, 0x0d, 0xab, 0x5d, 0x1a, 0x44, 0xfb, 0x6b, 0x97, 0xc9, 0x67,
	0x03, 0xa0, 0x02, 0xd9, 0x0d, 0x38, 0xd9, 0xa5, 0x59, 0xbf, 0xa6, 0x33, 0xfe, 0x00, 0x5e, 0xbb,
	0x40, 0xa7, 0x91, 0x88, 0x57, 0xe5, 0x88, 0xfe, 0x01, 0x9b, 0x81, 0x93, 0x9d, 0x65, 0x87, 0xda,
	0xb8, 0xba, 0x77, 0xb6, 0x07, 0xdd, 0xda, 0xdc, 0x82, 0x9b, 0xa7, 0xca, 0xce, 0x6b, 0xd4, 0x6f,
	0x71, 0xef, 0x1d, 0x64, 0x0a, 0xf6, 0x5c, 0x22, 0x3b, 0xad, 0xc9, 0xab, 0xf8, 0xbd, 0xfe, 0x6e,
	0xb0, 0xf4, 0x78, 0x72, 0xe9, 0x2b, 0xb8, 0xfe, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xc6, 0xe7, 0x0f,
	0x2d, 0x46, 0x03, 0x00, 0x00,
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/group/member.proto

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

// 加入成员的请求
type MemberAddRequest struct {
	Collection           string   `protobuf:"bytes,1,opt,name=collection,proto3" json:"collection,omitempty"`
	Element              string   `protobuf:"bytes,2,opt,name=element,proto3" json:"element,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemberAddRequest) Reset()         { *m = MemberAddRequest{} }
func (m *MemberAddRequest) String() string { return proto.CompactTextString(m) }
func (*MemberAddRequest) ProtoMessage()    {}
func (*MemberAddRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_640ebdb6edb0325d, []int{0}
}

func (m *MemberAddRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberAddRequest.Unmarshal(m, b)
}
func (m *MemberAddRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberAddRequest.Marshal(b, m, deterministic)
}
func (m *MemberAddRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberAddRequest.Merge(m, src)
}
func (m *MemberAddRequest) XXX_Size() int {
	return xxx_messageInfo_MemberAddRequest.Size(m)
}
func (m *MemberAddRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberAddRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MemberAddRequest proto.InternalMessageInfo

func (m *MemberAddRequest) GetCollection() string {
	if m != nil {
		return m.Collection
	}
	return ""
}

func (m *MemberAddRequest) GetElement() string {
	if m != nil {
		return m.Element
	}
	return ""
}

// 列举成员的请求
type MemberListRequest struct {
	Offset               int64    `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Count                int64    `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemberListRequest) Reset()         { *m = MemberListRequest{} }
func (m *MemberListRequest) String() string { return proto.CompactTextString(m) }
func (*MemberListRequest) ProtoMessage()    {}
func (*MemberListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_640ebdb6edb0325d, []int{1}
}

func (m *MemberListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberListRequest.Unmarshal(m, b)
}
func (m *MemberListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberListRequest.Marshal(b, m, deterministic)
}
func (m *MemberListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberListRequest.Merge(m, src)
}
func (m *MemberListRequest) XXX_Size() int {
	return xxx_messageInfo_MemberListRequest.Size(m)
}
func (m *MemberListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MemberListRequest proto.InternalMessageInfo

func (m *MemberListRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *MemberListRequest) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

// 列举成员的回复
type MemberListResponse struct {
	Status               *Status         `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Total                uint64          `protobuf:"varint,2,opt,name=total,proto3" json:"total,omitempty"`
	Entity               []*MemberEntity `protobuf:"bytes,3,rep,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MemberListResponse) Reset()         { *m = MemberListResponse{} }
func (m *MemberListResponse) String() string { return proto.CompactTextString(m) }
func (*MemberListResponse) ProtoMessage()    {}
func (*MemberListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_640ebdb6edb0325d, []int{2}
}

func (m *MemberListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberListResponse.Unmarshal(m, b)
}
func (m *MemberListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberListResponse.Marshal(b, m, deterministic)
}
func (m *MemberListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberListResponse.Merge(m, src)
}
func (m *MemberListResponse) XXX_Size() int {
	return xxx_messageInfo_MemberListResponse.Size(m)
}
func (m *MemberListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MemberListResponse proto.InternalMessageInfo

func (m *MemberListResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *MemberListResponse) GetTotal() uint64 {
	if m != nil {
		return m.Total
	}
	return 0
}

func (m *MemberListResponse) GetEntity() []*MemberEntity {
	if m != nil {
		return m.Entity
	}
	return nil
}

// 移除成员的请求
type MemberRemoveRequest struct {
	Element              string   `protobuf:"bytes,1,opt,name=element,proto3" json:"element,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemberRemoveRequest) Reset()         { *m = MemberRemoveRequest{} }
func (m *MemberRemoveRequest) String() string { return proto.CompactTextString(m) }
func (*MemberRemoveRequest) ProtoMessage()    {}
func (*MemberRemoveRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_640ebdb6edb0325d, []int{3}
}

func (m *MemberRemoveRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberRemoveRequest.Unmarshal(m, b)
}
func (m *MemberRemoveRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberRemoveRequest.Marshal(b, m, deterministic)
}
func (m *MemberRemoveRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberRemoveRequest.Merge(m, src)
}
func (m *MemberRemoveRequest) XXX_Size() int {
	return xxx_messageInfo_MemberRemoveRequest.Size(m)
}
func (m *MemberRemoveRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberRemoveRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MemberRemoveRequest proto.InternalMessageInfo

func (m *MemberRemoveRequest) GetElement() string {
	if m != nil {
		return m.Element
	}
	return ""
}

// 获取成员的请求
type MemberGetRequest struct {
	Element              string   `protobuf:"bytes,1,opt,name=element,proto3" json:"element,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MemberGetRequest) Reset()         { *m = MemberGetRequest{} }
func (m *MemberGetRequest) String() string { return proto.CompactTextString(m) }
func (*MemberGetRequest) ProtoMessage()    {}
func (*MemberGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_640ebdb6edb0325d, []int{4}
}

func (m *MemberGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberGetRequest.Unmarshal(m, b)
}
func (m *MemberGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberGetRequest.Marshal(b, m, deterministic)
}
func (m *MemberGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberGetRequest.Merge(m, src)
}
func (m *MemberGetRequest) XXX_Size() int {
	return xxx_messageInfo_MemberGetRequest.Size(m)
}
func (m *MemberGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MemberGetRequest proto.InternalMessageInfo

func (m *MemberGetRequest) GetElement() string {
	if m != nil {
		return m.Element
	}
	return ""
}

// 获取成员信息的回复
type MemberGetResponse struct {
	Status               *Status       `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Entity               *MemberEntity `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *MemberGetResponse) Reset()         { *m = MemberGetResponse{} }
func (m *MemberGetResponse) String() string { return proto.CompactTextString(m) }
func (*MemberGetResponse) ProtoMessage()    {}
func (*MemberGetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_640ebdb6edb0325d, []int{5}
}

func (m *MemberGetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MemberGetResponse.Unmarshal(m, b)
}
func (m *MemberGetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MemberGetResponse.Marshal(b, m, deterministic)
}
func (m *MemberGetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MemberGetResponse.Merge(m, src)
}
func (m *MemberGetResponse) XXX_Size() int {
	return xxx_messageInfo_MemberGetResponse.Size(m)
}
func (m *MemberGetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MemberGetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MemberGetResponse proto.InternalMessageInfo

func (m *MemberGetResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *MemberGetResponse) GetEntity() *MemberEntity {
	if m != nil {
		return m.Entity
	}
	return nil
}

func init() {
	proto.RegisterType((*MemberAddRequest)(nil), "group.MemberAddRequest")
	proto.RegisterType((*MemberListRequest)(nil), "group.MemberListRequest")
	proto.RegisterType((*MemberListResponse)(nil), "group.MemberListResponse")
	proto.RegisterType((*MemberRemoveRequest)(nil), "group.MemberRemoveRequest")
	proto.RegisterType((*MemberGetRequest)(nil), "group.MemberGetRequest")
	proto.RegisterType((*MemberGetResponse)(nil), "group.MemberGetResponse")
}

func init() {
	proto.RegisterFile("proto/group/member.proto", fileDescriptor_640ebdb6edb0325d)
}

var fileDescriptor_640ebdb6edb0325d = []byte{
	// 349 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4d, 0x4b, 0xc3, 0x40,
	0x10, 0x35, 0x4d, 0x1b, 0x71, 0x8a, 0xa0, 0xdb, 0xa2, 0x6b, 0x0e, 0x52, 0x02, 0x42, 0x41, 0x69,
	0xa1, 0x82, 0x07, 0xd1, 0x43, 0x05, 0xe9, 0xa5, 0x5e, 0xd6, 0x5f, 0x90, 0x36, 0xd3, 0x5a, 0x4c,
	0xb2, 0x35, 0x3b, 0x11, 0xbc, 0xf8, 0xbf, 0xbd, 0x49, 0x77, 0xb3, 0x26, 0xf1, 0x13, 0x8f, 0x33,
	0xef, 0x23, 0xfb, 0xde, 0x04, 0xf8, 0x3a, 0x93, 0x24, 0x87, 0xcb, 0x4c, 0xe6, 0xeb, 0x61, 0x82,
	0xc9, 0x0c, 0xb3, 0x81, 0x5e, 0xb1, 0x96, 0xde, 0xf9, 0x35, 0x82, 0x7a, 0x08, 0x33, 0x8c, 0x0c,
	0x21, 0x98, 0xc2, 0xde, 0x9d, 0x16, 0x8c, 0xa3, 0x48, 0xe0, 0x53, 0x8e, 0x8a, 0xd8, 0x31, 0xc0,
	0x5c, 0xc6, 0x31, 0xce, 0x69, 0x25, 0x53, 0xee, 0xf4, 0x9c, 0xfe, 0x8e, 0xa8, 0x6c, 0x18, 0x87,
	0x6d, 0x8c, 0x31, 0xc1, 0x94, 0x78, 0x43, 0x83, 0x76, 0x0c, 0xc6, 0xb0, 0x6f, 0xdc, 0xa6, 0x2b,
	0x45, 0xd6, 0xee, 0x00, 0x3c, 0xb9, 0x58, 0x28, 0x24, 0x6d, 0xe5, 0x8a, 0x62, 0x62, 0x5d, 0x68,
	0xcd, 0x65, 0x5e, 0x98, 0xb8, 0xc2, 0x0c, 0xc1, 0x2b, 0xb0, 0xaa, 0x85, 0x5a, 0xcb, 0x54, 0x21,
	0x3b, 0x01, 0x4f, 0x51, 0x48, 0xb9, 0xd2, 0x1e, 0xed, 0xd1, 0xee, 0x40, 0x67, 0x19, 0xdc, 0xeb,
	0xa5, 0x28, 0xc0, 0x8d, 0x25, 0x49, 0x0a, 0x63, 0x6d, 0xd9, 0x14, 0x66, 0x60, 0xa7, 0xe0, 0x61,
	0x4a, 0x2b, 0x7a, 0xe1, 0x6e, 0xcf, 0xed, 0xb7, 0x47, 0x9d, 0x42, 0x6c, 0xbe, 0x73, 0xab, 0x21,
	0x51, 0x50, 0x82, 0x21, 0x74, 0xcc, 0x5e, 0x60, 0x22, 0x9f, 0xd1, 0x86, 0xa8, 0x64, 0x76, 0xea,
	0x99, 0xcf, 0x6c, 0x83, 0x13, 0xa4, 0xbf, 0xd9, 0x4b, 0xdb, 0x90, 0x66, 0xff, 0x2f, 0x5d, 0x99,
	0xa3, 0xa1, 0x69, 0xbf, 0xe5, 0x18, 0xbd, 0x39, 0xe0, 0x19, 0x80, 0x5d, 0x80, 0x3b, 0x8e, 0x22,
	0x76, 0x58, 0xa3, 0x97, 0xf7, 0xf6, 0xbb, 0x05, 0x70, 0x13, 0x87, 0xe9, 0xa3, 0x7d, 0x54, 0xb0,
	0xc5, 0xae, 0xa1, 0xb9, 0x39, 0x02, 0xe3, 0x35, 0x61, 0xe5, 0xb4, 0xfe, 0xd1, 0x37, 0xc8, 0x87,
	0xfc, 0x0a, 0x3c, 0xd3, 0x21, 0xf3, 0x6b, 0xb4, 0x5a, 0xb1, 0x3f, 0x7e, 0xfc, 0x12, 0xdc, 0x09,
	0xd2, 0xa7, 0x47, 0x97, 0x15, 0xfb, 0xfc, 0x2b, 0x60, 0xb5, 0x33, 0x4f, 0xff, 0xdb, 0xe7, 0xef,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x6e, 0x7e, 0xa1, 0x00, 0x18, 0x03, 0x00, 0x00,
}
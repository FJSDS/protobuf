// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: imports/test_a_2/m3.proto

package test_a_2

import (
	fmt "fmt"
	proto "github.com/FJSDS/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type M3 struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *M3) Reset()         { *m = M3{} }
func (m *M3) String() string { return proto.CompactTextString(m) }
func (*M3) ProtoMessage()    {}
func (*M3) Descriptor() ([]byte, []int) {
	return fileDescriptor_ff9d8f834875c9c5, []int{0}
}
func (m *M3) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M3.Unmarshal(m, b)
}
func (m *M3) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M3.Marshal(b, m, deterministic)
}
func (m *M3) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M3.Merge(m, src)
}
func (m *M3) XXX_Size() int {
	return xxx_messageInfo_M3.Size(m)
}
func (m *M3) XXX_DiscardUnknown() {
	xxx_messageInfo_M3.DiscardUnknown(m)
}

var xxx_messageInfo_M3 proto.InternalMessageInfo

func init() {
	proto.RegisterType((*M3)(nil), "test.a.M3")
}

func init() { proto.RegisterFile("imports/test_a_2/m3.proto", fileDescriptor_ff9d8f834875c9c5) }

var fileDescriptor_ff9d8f834875c9c5 = []byte{
	// 112 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0xcc, 0x2d, 0xc8,
	0x2f, 0x2a, 0x29, 0xd6, 0x2f, 0x49, 0x2d, 0x2e, 0x89, 0x4f, 0x8c, 0x37, 0xd2, 0xcf, 0x35, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x03, 0x09, 0xe9, 0x25, 0x2a, 0xb1, 0x70, 0x31, 0xf9,
	0x1a, 0x3b, 0xb9, 0x44, 0x39, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea,
	0xa7, 0xe7, 0xa7, 0xe7, 0xeb, 0x83, 0x95, 0x25, 0x95, 0xa6, 0x41, 0x18, 0xc9, 0xba, 0xe9, 0xa9,
	0x79, 0xba, 0x60, 0x09, 0x90, 0xc6, 0x94, 0xc4, 0x92, 0x44, 0x7d, 0x74, 0xc3, 0x93, 0xd8, 0xc0,
	0x4a, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x11, 0xfd, 0xd0, 0xcb, 0x77, 0x00, 0x00, 0x00,
}

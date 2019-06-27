// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: md.proto

package moredefaults

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/FJSDS/protobuf/gogoproto"
	proto "github.com/FJSDS/protobuf/proto"
	example "github.com/FJSDS/protobuf/test/example"
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

type MoreDefaultsB struct {
	Field1               *string  `protobuf:"bytes,1,opt,name=Field1" json:"Field1,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MoreDefaultsB) Reset()         { *m = MoreDefaultsB{} }
func (m *MoreDefaultsB) String() string { return proto.CompactTextString(m) }
func (*MoreDefaultsB) ProtoMessage()    {}
func (*MoreDefaultsB) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0e149d9fdc447d0, []int{0}
}
func (m *MoreDefaultsB) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoreDefaultsB.Unmarshal(m, b)
}
func (m *MoreDefaultsB) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoreDefaultsB.Marshal(b, m, deterministic)
}
func (m *MoreDefaultsB) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoreDefaultsB.Merge(m, src)
}
func (m *MoreDefaultsB) XXX_Size() int {
	return xxx_messageInfo_MoreDefaultsB.Size(m)
}
func (m *MoreDefaultsB) XXX_DiscardUnknown() {
	xxx_messageInfo_MoreDefaultsB.DiscardUnknown(m)
}

var xxx_messageInfo_MoreDefaultsB proto.InternalMessageInfo

func (m *MoreDefaultsB) GetField1() string {
	if m != nil && m.Field1 != nil {
		return *m.Field1
	}
	return ""
}

type MoreDefaultsA struct {
	Field1               *int64         `protobuf:"varint,1,opt,name=Field1,def=1234" json:"Field1,omitempty"`
	Field2               int64          `protobuf:"varint,2,opt,name=Field2" json:"Field2"`
	B1                   *MoreDefaultsB `protobuf:"bytes,3,opt,name=B1" json:"B1,omitempty"`
	B2                   MoreDefaultsB  `protobuf:"bytes,4,opt,name=B2" json:"B2"`
	A1                   *example.A     `protobuf:"bytes,5,opt,name=A1" json:"A1,omitempty"`
	A2                   example.A      `protobuf:"bytes,6,opt,name=A2" json:"A2"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *MoreDefaultsA) Reset()         { *m = MoreDefaultsA{} }
func (m *MoreDefaultsA) String() string { return proto.CompactTextString(m) }
func (*MoreDefaultsA) ProtoMessage()    {}
func (*MoreDefaultsA) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0e149d9fdc447d0, []int{1}
}
func (m *MoreDefaultsA) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MoreDefaultsA.Unmarshal(m, b)
}
func (m *MoreDefaultsA) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MoreDefaultsA.Marshal(b, m, deterministic)
}
func (m *MoreDefaultsA) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MoreDefaultsA.Merge(m, src)
}
func (m *MoreDefaultsA) XXX_Size() int {
	return xxx_messageInfo_MoreDefaultsA.Size(m)
}
func (m *MoreDefaultsA) XXX_DiscardUnknown() {
	xxx_messageInfo_MoreDefaultsA.DiscardUnknown(m)
}

var xxx_messageInfo_MoreDefaultsA proto.InternalMessageInfo

const Default_MoreDefaultsA_Field1 int64 = 1234

func (m *MoreDefaultsA) GetField1() int64 {
	if m != nil && m.Field1 != nil {
		return *m.Field1
	}
	return Default_MoreDefaultsA_Field1
}

func (m *MoreDefaultsA) GetField2() int64 {
	if m != nil {
		return m.Field2
	}
	return 0
}

func (m *MoreDefaultsA) GetB1() *MoreDefaultsB {
	if m != nil {
		return m.B1
	}
	return nil
}

func (m *MoreDefaultsA) GetB2() MoreDefaultsB {
	if m != nil {
		return m.B2
	}
	return MoreDefaultsB{}
}

func (m *MoreDefaultsA) GetA1() *example.A {
	if m != nil {
		return m.A1
	}
	return nil
}

func (m *MoreDefaultsA) GetA2() example.A {
	if m != nil {
		return m.A2
	}
	return example.A{}
}

func init() {
	proto.RegisterType((*MoreDefaultsB)(nil), "moredefaults.MoreDefaultsB")
	proto.RegisterType((*MoreDefaultsA)(nil), "moredefaults.MoreDefaultsA")
}

func init() { proto.RegisterFile("md.proto", fileDescriptor_e0e149d9fdc447d0) }

var fileDescriptor_e0e149d9fdc447d0 = []byte{
	// 258 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xc8, 0x4d, 0xd1, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xc9, 0xcd, 0x2f, 0x4a, 0x4d, 0x49, 0x4d, 0x4b, 0x2c, 0xcd,
	0x29, 0x29, 0x96, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f,
	0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0x2b, 0x4a, 0x2a, 0x4d, 0x03, 0xf3, 0xc0, 0x1c, 0x30, 0x0b, 0xa2,
	0x59, 0xca, 0x18, 0xa7, 0xf2, 0x92, 0xd4, 0xe2, 0x12, 0xfd, 0xd4, 0x8a, 0xc4, 0xdc, 0x82, 0x9c,
	0x54, 0x18, 0x0d, 0xd1, 0xa4, 0xa4, 0xce, 0xc5, 0xeb, 0x9b, 0x5f, 0x94, 0xea, 0x02, 0xb5, 0xd3,
	0x49, 0x48, 0x8c, 0x8b, 0xcd, 0x2d, 0x33, 0x35, 0x27, 0xc5, 0x50, 0x82, 0x51, 0x81, 0x51, 0x83,
	0x33, 0x08, 0xca, 0x53, 0x7a, 0xcc, 0x88, 0xaa, 0xd2, 0x51, 0x48, 0x06, 0x45, 0x25, 0xb3, 0x15,
	0x8b, 0xa1, 0x91, 0xb1, 0x09, 0x4c, 0x3d, 0x5c, 0xd6, 0x48, 0x82, 0x09, 0x24, 0xeb, 0xc4, 0x72,
	0xe2, 0x9e, 0x3c, 0x03, 0x54, 0xd6, 0x48, 0x48, 0x9b, 0x8b, 0xc9, 0xc9, 0x50, 0x82, 0x59, 0x81,
	0x51, 0x83, 0xdb, 0x48, 0x5a, 0x0f, 0xd9, 0xd7, 0x7a, 0x28, 0xce, 0x09, 0x62, 0x72, 0x32, 0x14,
	0x32, 0xe4, 0x62, 0x72, 0x32, 0x92, 0x60, 0x21, 0xa8, 0x18, 0x6a, 0x07, 0x93, 0x93, 0x91, 0x90,
	0x38, 0x17, 0x93, 0xa3, 0xa1, 0x04, 0x2b, 0x58, 0x0b, 0xbb, 0x1e, 0xc8, 0xff, 0x7a, 0x8e, 0x41,
	0x4c, 0x8e, 0x86, 0x42, 0xb2, 0x5c, 0x4c, 0x8e, 0x46, 0x12, 0x6c, 0x28, 0x12, 0x30, 0x7d, 0x8e,
	0x46, 0x4e, 0x02, 0x27, 0x1e, 0xca, 0x31, 0xfe, 0x78, 0x28, 0xc7, 0xb8, 0xe2, 0x91, 0x1c, 0xe3,
	0x8e, 0x47, 0x72, 0x8c, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf2, 0xf0, 0x3f, 0xeb, 0x9d, 0x01,
	0x00, 0x00,
}

func (this *MoreDefaultsB) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MoreDefaultsB)
	if !ok {
		that2, ok := that.(MoreDefaultsB)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Field1 != nil && that1.Field1 != nil {
		if *this.Field1 != *that1.Field1 {
			return false
		}
	} else if this.Field1 != nil {
		return false
	} else if that1.Field1 != nil {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *MoreDefaultsA) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MoreDefaultsA)
	if !ok {
		that2, ok := that.(MoreDefaultsA)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Field1 != nil && that1.Field1 != nil {
		if *this.Field1 != *that1.Field1 {
			return false
		}
	} else if this.Field1 != nil {
		return false
	} else if that1.Field1 != nil {
		return false
	}
	if this.Field2 != that1.Field2 {
		return false
	}
	if !this.B1.Equal(that1.B1) {
		return false
	}
	if !this.B2.Equal(&that1.B2) {
		return false
	}
	if !this.A1.Equal(that1.A1) {
		return false
	}
	if !this.A2.Equal(&that1.A2) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func NewPopulatedMoreDefaultsB(r randyMd, easy bool) *MoreDefaultsB {
	this := &MoreDefaultsB{}
	if r.Intn(10) != 0 {
		v1 := string(randStringMd(r))
		this.Field1 = &v1
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedMd(r, 2)
	}
	return this
}

func NewPopulatedMoreDefaultsA(r randyMd, easy bool) *MoreDefaultsA {
	this := &MoreDefaultsA{}
	if r.Intn(10) != 0 {
		v2 := int64(r.Int63())
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		this.Field1 = &v2
	}
	this.Field2 = int64(r.Int63())
	if r.Intn(2) == 0 {
		this.Field2 *= -1
	}
	if r.Intn(10) != 0 {
		this.B1 = NewPopulatedMoreDefaultsB(r, easy)
	}
	v3 := NewPopulatedMoreDefaultsB(r, easy)
	this.B2 = *v3
	if r.Intn(10) != 0 {
		this.A1 = example.NewPopulatedA(r, easy)
	}
	v4 := example.NewPopulatedA(r, easy)
	this.A2 = *v4
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedMd(r, 7)
	}
	return this
}

type randyMd interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneMd(r randyMd) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringMd(r randyMd) string {
	v5 := r.Intn(100)
	tmps := make([]rune, v5)
	for i := 0; i < v5; i++ {
		tmps[i] = randUTF8RuneMd(r)
	}
	return string(tmps)
}
func randUnrecognizedMd(r randyMd, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldMd(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldMd(dAtA []byte, r randyMd, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateMd(dAtA, uint64(key))
		v6 := r.Int63()
		if r.Intn(2) == 0 {
			v6 *= -1
		}
		dAtA = encodeVarintPopulateMd(dAtA, uint64(v6))
	case 1:
		dAtA = encodeVarintPopulateMd(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateMd(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateMd(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateMd(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateMd(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}

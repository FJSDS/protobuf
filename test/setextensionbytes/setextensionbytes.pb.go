// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: setextensionbytes.proto

package setextensionbytes

import (
	bytes "bytes"
	fmt "fmt"
	_ "github.com/FJSDS/protobuf/gogoproto"
	github_com_gogo_protobuf_proto "github.com/FJSDS/protobuf/proto"
	proto "github.com/FJSDS/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

type MyExtendable struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_extensions       []byte   `protobuf:"bytes,0,opt" json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MyExtendable) Reset()      { *m = MyExtendable{} }
func (*MyExtendable) ProtoMessage() {}
func (*MyExtendable) Descriptor() ([]byte, []int) {
	return fileDescriptor_41b8000ad7d06fe7, []int{0}
}

var extRange_MyExtendable = []proto.ExtensionRange{
	{Start: 1, End: 10},
}

func (*MyExtendable) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_MyExtendable
}

func (m *MyExtendable) GetExtensions() *[]byte {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make([]byte, 0)
	}
	return &m.XXX_extensions
}
func (m *MyExtendable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MyExtendable.Unmarshal(m, b)
}
func (m *MyExtendable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MyExtendable.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MyExtendable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MyExtendable.Merge(m, src)
}
func (m *MyExtendable) XXX_Size() int {
	return m.Size()
}
func (m *MyExtendable) XXX_DiscardUnknown() {
	xxx_messageInfo_MyExtendable.DiscardUnknown(m)
}

var xxx_messageInfo_MyExtendable proto.InternalMessageInfo

type Foo struct {
	IntFoo               int64    `protobuf:"varint,1,opt,name=intFoo" json:"intFoo"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Foo) Reset()      { *m = Foo{} }
func (*Foo) ProtoMessage() {}
func (*Foo) Descriptor() ([]byte, []int) {
	return fileDescriptor_41b8000ad7d06fe7, []int{1}
}
func (m *Foo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Foo.Unmarshal(m, b)
}
func (m *Foo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Foo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Foo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Foo.Merge(m, src)
}
func (m *Foo) XXX_Size() int {
	return m.Size()
}
func (m *Foo) XXX_DiscardUnknown() {
	xxx_messageInfo_Foo.DiscardUnknown(m)
}

var xxx_messageInfo_Foo proto.InternalMessageInfo

var E_Foos = &proto.ExtensionDesc{
	ExtendedType:  (*MyExtendable)(nil),
	ExtensionType: (*Foo)(nil),
	Field:         2,
	Name:          "setextensionbytes.Foos",
	Tag:           "bytes,2,opt,name=Foos",
	Filename:      "setextensionbytes.proto",
}

func init() {
	proto.RegisterType((*MyExtendable)(nil), "setextensionbytes.MyExtendable")
	proto.RegisterType((*Foo)(nil), "setextensionbytes.Foo")
	proto.RegisterExtension(E_Foos)
}

func init() { proto.RegisterFile("setextensionbytes.proto", fileDescriptor_41b8000ad7d06fe7) }

var fileDescriptor_41b8000ad7d06fe7 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x4e, 0x2d, 0x49,
	0xad, 0x28, 0x49, 0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0x4b, 0xaa, 0x2c, 0x49, 0x2d, 0xd6, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x12, 0xc4, 0x90, 0x90, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2,
	0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf, 0x4f, 0xcf, 0xd7, 0x07, 0xab, 0x4c, 0x2a, 0x4d, 0x03, 0xf3,
	0xc0, 0x1c, 0x30, 0x0b, 0x62, 0x82, 0x92, 0x1c, 0x17, 0x8f, 0x6f, 0xa5, 0x2b, 0xc8, 0x88, 0x94,
	0xc4, 0xa4, 0x9c, 0x54, 0x2d, 0x16, 0x0e, 0x46, 0x01, 0x6e, 0x2b, 0x8e, 0x0d, 0x0b, 0xe4, 0x19,
	0x4e, 0x2c, 0x94, 0x67, 0x50, 0x52, 0xe6, 0x62, 0x76, 0xcb, 0xcf, 0x17, 0x92, 0xe1, 0x62, 0xcb,
	0xcc, 0x2b, 0x71, 0xcb, 0xcf, 0x97, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x76, 0x62, 0x39, 0x71, 0x4f,
	0x9e, 0x21, 0x08, 0x2a, 0x66, 0xe5, 0xcd, 0xc5, 0xe2, 0x96, 0x9f, 0x5f, 0x2c, 0x24, 0xaf, 0x87,
	0xe9, 0x50, 0x64, 0xd3, 0x25, 0x98, 0x14, 0x18, 0x35, 0xb8, 0x8d, 0xc4, 0xb0, 0x28, 0x73, 0xcb,
	0xcf, 0x0f, 0x02, 0x1b, 0xe2, 0x24, 0x73, 0xe2, 0xa1, 0x1c, 0xc3, 0x8d, 0x87, 0x72, 0x0c, 0x0d,
	0x8f, 0xe4, 0x18, 0x57, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0xf1, 0xc7, 0x23, 0x39, 0xc6,
	0x86, 0xc7, 0x72, 0x8c, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x32, 0xc1, 0xe9, 0x65, 0x0b, 0x01,
	0x00, 0x00,
}

func (this *MyExtendable) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MyExtendable)
	if !ok {
		that2, ok := that.(MyExtendable)
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
	if !bytes.Equal(this.XXX_extensions, that1.XXX_extensions) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Foo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Foo)
	if !ok {
		that2, ok := that.(Foo)
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
	if this.IntFoo != that1.IntFoo {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (m *MyExtendable) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MyExtendable) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.XXX_extensions != nil {
		i += copy(dAtA[i:], m.XXX_extensions)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Foo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Foo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0x8
	i++
	i = encodeVarintSetextensionbytes(dAtA, i, uint64(m.IntFoo))
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintSetextensionbytes(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *MyExtendable) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.XXX_extensions != nil {
		n += len(m.XXX_extensions)
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Foo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	n += 1 + sovSetextensionbytes(uint64(m.IntFoo))
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSetextensionbytes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSetextensionbytes(x uint64) (n int) {
	return sovSetextensionbytes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *MyExtendable) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&MyExtendable{`,
		`XXX_extensions:` + github_com_gogo_protobuf_proto.StringFromExtensionsBytes(this.XXX_extensions) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Foo) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Foo{`,
		`IntFoo:` + fmt.Sprintf("%v", this.IntFoo) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringSetextensionbytes(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *MyExtendable) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSetextensionbytes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MyExtendable: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MyExtendable: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			if (fieldNum >= 1) && (fieldNum < 11) {
				var sizeOfWire int
				for {
					sizeOfWire++
					wire >>= 7
					if wire == 0 {
						break
					}
				}
				iNdEx -= sizeOfWire
				skippy, err := skipSetextensionbytes(dAtA[iNdEx:])
				if err != nil {
					return err
				}
				if skippy < 0 {
					return ErrInvalidLengthSetextensionbytes
				}
				if (iNdEx + skippy) < 0 {
					return ErrInvalidLengthSetextensionbytes
				}
				if (iNdEx + skippy) > l {
					return io.ErrUnexpectedEOF
				}
				github_com_gogo_protobuf_proto.AppendExtension(m, int32(fieldNum), dAtA[iNdEx:iNdEx+skippy])
				iNdEx += skippy
			} else {
				iNdEx = preIndex
				skippy, err := skipSetextensionbytes(dAtA[iNdEx:])
				if err != nil {
					return err
				}
				if skippy < 0 {
					return ErrInvalidLengthSetextensionbytes
				}
				if (iNdEx + skippy) < 0 {
					return ErrInvalidLengthSetextensionbytes
				}
				if (iNdEx + skippy) > l {
					return io.ErrUnexpectedEOF
				}
				m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
				iNdEx += skippy
			}
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Foo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSetextensionbytes
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Foo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Foo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IntFoo", wireType)
			}
			m.IntFoo = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSetextensionbytes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IntFoo |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSetextensionbytes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSetextensionbytes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthSetextensionbytes
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSetextensionbytes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSetextensionbytes
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSetextensionbytes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSetextensionbytes
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthSetextensionbytes
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthSetextensionbytes
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSetextensionbytes
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipSetextensionbytes(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthSetextensionbytes
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthSetextensionbytes = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSetextensionbytes   = fmt.Errorf("proto: integer overflow")
)

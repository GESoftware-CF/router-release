// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sidecar.proto

package models

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Sidecar struct {
	Action   *Action `protobuf:"bytes,1,opt,name=action,proto3" json:"action,omitempty"`
	DiskMb   int32   `protobuf:"varint,2,opt,name=disk_mb,json=diskMb,proto3" json:"disk_mb"`
	MemoryMb int32   `protobuf:"varint,3,opt,name=memory_mb,json=memoryMb,proto3" json:"memory_mb"`
}

func (m *Sidecar) Reset()      { *m = Sidecar{} }
func (*Sidecar) ProtoMessage() {}
func (*Sidecar) Descriptor() ([]byte, []int) {
	return fileDescriptor_179ad3b13e6397ec, []int{0}
}
func (m *Sidecar) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Sidecar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Sidecar.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Sidecar) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Sidecar.Merge(m, src)
}
func (m *Sidecar) XXX_Size() int {
	return m.Size()
}
func (m *Sidecar) XXX_DiscardUnknown() {
	xxx_messageInfo_Sidecar.DiscardUnknown(m)
}

var xxx_messageInfo_Sidecar proto.InternalMessageInfo

func (m *Sidecar) GetAction() *Action {
	if m != nil {
		return m.Action
	}
	return nil
}

func (m *Sidecar) GetDiskMb() int32 {
	if m != nil {
		return m.DiskMb
	}
	return 0
}

func (m *Sidecar) GetMemoryMb() int32 {
	if m != nil {
		return m.MemoryMb
	}
	return 0
}

func init() {
	proto.RegisterType((*Sidecar)(nil), "models.Sidecar")
}

func init() { proto.RegisterFile("sidecar.proto", fileDescriptor_179ad3b13e6397ec) }

var fileDescriptor_179ad3b13e6397ec = []byte{
	// 239 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0xce, 0x4c, 0x49,
	0x4d, 0x4e, 0x2c, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xcb, 0xcd, 0x4f, 0x49, 0xcd,
	0x29, 0x96, 0xd2, 0x4d, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x4f, 0xcf,
	0x4f, 0xcf, 0xd7, 0x07, 0x4b, 0x27, 0x95, 0xa6, 0x81, 0x79, 0x60, 0x0e, 0x98, 0x05, 0xd1, 0x26,
	0xc5, 0x9b, 0x98, 0x5c, 0x92, 0x99, 0x9f, 0x57, 0x0c, 0xe1, 0x2a, 0x35, 0x33, 0x72, 0xb1, 0x07,
	0x43, 0xcc, 0x15, 0x52, 0xe3, 0x62, 0x83, 0x48, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x1b, 0xf1,
	0xe9, 0x41, 0xac, 0xd0, 0x73, 0x04, 0x8b, 0x06, 0x41, 0x65, 0x85, 0x54, 0xb8, 0xd8, 0x53, 0x32,
	0x8b, 0xb3, 0xe3, 0x73, 0x93, 0x24, 0x98, 0x14, 0x18, 0x35, 0x58, 0x9d, 0xb8, 0x5f, 0xdd, 0x93,
	0x87, 0x09, 0x05, 0xb1, 0x81, 0x18, 0xbe, 0x49, 0x42, 0x5a, 0x5c, 0x9c, 0xb9, 0xa9, 0xb9, 0xf9,
	0x45, 0x95, 0x20, 0x75, 0xcc, 0x60, 0x75, 0xbc, 0xaf, 0xee, 0xc9, 0x23, 0x04, 0x83, 0x38, 0x20,
	0x4c, 0xdf, 0x24, 0x27, 0x93, 0x0b, 0x0f, 0xe5, 0x18, 0x6e, 0x3c, 0x94, 0x63, 0xf8, 0xf0, 0x50,
	0x8e, 0xb1, 0xe1, 0x91, 0x1c, 0xe3, 0x8a, 0x47, 0x72, 0x8c, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78,
	0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x8b, 0x47, 0x72, 0x0c, 0x1f, 0x1e, 0xc9, 0x31, 0x4e,
	0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1, 0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x49, 0x6c, 0x60, 0x2f,
	0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x28, 0x34, 0x4f, 0x19, 0x01, 0x00, 0x00,
}

func (this *Sidecar) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Sidecar)
	if !ok {
		that2, ok := that.(Sidecar)
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
	if !this.Action.Equal(that1.Action) {
		return false
	}
	if this.DiskMb != that1.DiskMb {
		return false
	}
	if this.MemoryMb != that1.MemoryMb {
		return false
	}
	return true
}
func (this *Sidecar) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 7)
	s = append(s, "&models.Sidecar{")
	if this.Action != nil {
		s = append(s, "Action: "+fmt.Sprintf("%#v", this.Action)+",\n")
	}
	s = append(s, "DiskMb: "+fmt.Sprintf("%#v", this.DiskMb)+",\n")
	s = append(s, "MemoryMb: "+fmt.Sprintf("%#v", this.MemoryMb)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringSidecar(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *Sidecar) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Sidecar) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Sidecar) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MemoryMb != 0 {
		i = encodeVarintSidecar(dAtA, i, uint64(m.MemoryMb))
		i--
		dAtA[i] = 0x18
	}
	if m.DiskMb != 0 {
		i = encodeVarintSidecar(dAtA, i, uint64(m.DiskMb))
		i--
		dAtA[i] = 0x10
	}
	if m.Action != nil {
		{
			size, err := m.Action.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintSidecar(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintSidecar(dAtA []byte, offset int, v uint64) int {
	offset -= sovSidecar(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Sidecar) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Action != nil {
		l = m.Action.Size()
		n += 1 + l + sovSidecar(uint64(l))
	}
	if m.DiskMb != 0 {
		n += 1 + sovSidecar(uint64(m.DiskMb))
	}
	if m.MemoryMb != 0 {
		n += 1 + sovSidecar(uint64(m.MemoryMb))
	}
	return n
}

func sovSidecar(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSidecar(x uint64) (n int) {
	return sovSidecar(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *Sidecar) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Sidecar{`,
		`Action:` + strings.Replace(fmt.Sprintf("%v", this.Action), "Action", "Action", 1) + `,`,
		`DiskMb:` + fmt.Sprintf("%v", this.DiskMb) + `,`,
		`MemoryMb:` + fmt.Sprintf("%v", this.MemoryMb) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringSidecar(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *Sidecar) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSidecar
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
			return fmt.Errorf("proto: Sidecar: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Sidecar: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Action", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSidecar
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSidecar
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthSidecar
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Action == nil {
				m.Action = &Action{}
			}
			if err := m.Action.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DiskMb", wireType)
			}
			m.DiskMb = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSidecar
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DiskMb |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MemoryMb", wireType)
			}
			m.MemoryMb = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSidecar
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MemoryMb |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSidecar(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSidecar
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSidecar(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSidecar
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
					return 0, ErrIntOverflowSidecar
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSidecar
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
				return 0, ErrInvalidLengthSidecar
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSidecar
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSidecar
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSidecar        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSidecar          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSidecar = fmt.Errorf("proto: unexpected end of group")
)

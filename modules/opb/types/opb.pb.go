// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: opb/opb.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

// Params defines the parameters for the OPB module.
type Params struct {
	BaseTokenDenom            string `protobuf:"bytes,1,opt,name=base_token_denom,json=baseTokenDenom,proto3" json:"base_token_denom,omitempty" yaml:"base_token_denom"`
	PointTokenDenom           string `protobuf:"bytes,2,opt,name=point_token_denom,json=pointTokenDenom,proto3" json:"point_token_denom,omitempty" yaml:"point_token_denom"`
	BaseTokenManager          string `protobuf:"bytes,3,opt,name=base_token_manager,json=baseTokenManager,proto3" json:"base_token_manager,omitempty" yaml:"base_token_manager"`
	UnrestrictedTokenTransfer bool   `protobuf:"varint,4,opt,name=unrestricted_token_transfer,json=unrestrictedTokenTransfer,proto3" json:"unrestricted_token_transfer,omitempty" yaml:"unrestricted_token_transfer"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_1cbfaa920b6e27d9, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "metaos.opb.Params")
}

func init() { proto.RegisterFile("opb/opb.proto", fileDescriptor_1cbfaa920b6e27d9) }

var fileDescriptor_1cbfaa920b6e27d9 = []byte{
	// 330 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x31, 0x4b, 0xf3, 0x40,
	0x18, 0xc7, 0x73, 0x7d, 0x5f, 0x4a, 0x3d, 0x50, 0x6b, 0x10, 0x4c, 0xad, 0x5e, 0xca, 0x0d, 0x52,
	0x10, 0x9a, 0xc1, 0xad, 0x63, 0x51, 0x51, 0x54, 0x90, 0xd0, 0xc9, 0xa5, 0xdc, 0xb5, 0xd7, 0x18,
	0xec, 0xe5, 0x09, 0x77, 0xd7, 0xa1, 0xdf, 0xc2, 0xd1, 0xb1, 0x1f, 0xa7, 0x63, 0xc7, 0x4e, 0x41,
	0xdb, 0xc5, 0x39, 0x9f, 0x40, 0x92, 0x14, 0xa9, 0x16, 0xdc, 0x8e, 0xdf, 0xf3, 0xbf, 0xdf, 0xf3,
	0xc0, 0x1f, 0xef, 0x42, 0xcc, 0x3d, 0x88, 0x79, 0x2b, 0x56, 0x60, 0xc0, 0xc6, 0x52, 0x18, 0x06,
	0xba, 0x05, 0x31, 0x3f, 0x3e, 0x0c, 0x20, 0x80, 0x1c, 0x7b, 0xd9, 0xab, 0x48, 0xd0, 0x45, 0x09,
	0x97, 0x1f, 0x99, 0x62, 0x52, 0xdb, 0x57, 0xb8, 0xca, 0x99, 0x16, 0x3d, 0x03, 0x2f, 0x22, 0xea,
	0x0d, 0x44, 0x04, 0xd2, 0x41, 0x0d, 0xd4, 0xdc, 0xe9, 0xd4, 0xd3, 0xc4, 0x3d, 0x9a, 0x30, 0x39,
	0x6a, 0xd3, 0xdf, 0x09, 0xea, 0xef, 0x65, 0xa8, 0x9b, 0x91, 0xcb, 0x0c, 0xd8, 0x37, 0xf8, 0x20,
	0x86, 0x30, 0x32, 0x3f, 0x3c, 0xa5, 0xdc, 0x73, 0x92, 0x26, 0xae, 0x53, 0x78, 0xb6, 0x22, 0xd4,
	0xdf, 0xcf, 0xd9, 0x86, 0xe9, 0x0e, 0xdb, 0x1b, 0xeb, 0x24, 0x8b, 0x58, 0x20, 0x94, 0xf3, 0x2f,
	0x57, 0x9d, 0xa6, 0x89, 0x5b, 0xdb, 0x3a, 0x69, 0x9d, 0xa1, 0x7e, 0xf5, 0xfb, 0xa8, 0x87, 0x02,
	0xd9, 0x43, 0x5c, 0x1f, 0x47, 0x4a, 0x68, 0xa3, 0xc2, 0xbe, 0x11, 0x83, 0xf5, 0x07, 0xa3, 0x58,
	0xa4, 0x87, 0x42, 0x39, 0xff, 0x1b, 0xa8, 0x59, 0xe9, 0x9c, 0xa5, 0x89, 0x4b, 0x0b, 0xeb, 0x1f,
	0x61, 0xea, 0xd7, 0x36, 0xa7, 0xf9, 0x9a, 0xee, 0x7a, 0xd6, 0xae, 0xbc, 0x4d, 0x5d, 0xeb, 0x73,
	0xea, 0xa2, 0xce, 0xed, 0xec, 0x83, 0x58, 0xb3, 0x25, 0x41, 0xf3, 0x25, 0x41, 0xef, 0x4b, 0x82,
	0x5e, 0x57, 0xc4, 0x9a, 0xaf, 0x88, 0xb5, 0x58, 0x11, 0xeb, 0xe9, 0x3c, 0x08, 0xcd, 0xf3, 0x98,
	0xb7, 0xfa, 0x20, 0xbd, 0xeb, 0x30, 0xbc, 0x67, 0x5c, 0x7b, 0x45, 0x5b, 0x9e, 0x84, 0xc1, 0x78,
	0x24, 0x74, 0xd6, 0xa3, 0x67, 0x26, 0xb1, 0xd0, 0xbc, 0x9c, 0x97, 0x75, 0xf1, 0x15, 0x00, 0x00,
	0xff, 0xff, 0xa2, 0x52, 0xbe, 0x8c, 0xdf, 0x01, 0x00, 0x00,
}

func (this *Params) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Params)
	if !ok {
		that2, ok := that.(Params)
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
	if this.BaseTokenDenom != that1.BaseTokenDenom {
		return false
	}
	if this.PointTokenDenom != that1.PointTokenDenom {
		return false
	}
	if this.BaseTokenManager != that1.BaseTokenManager {
		return false
	}
	if this.UnrestrictedTokenTransfer != that1.UnrestrictedTokenTransfer {
		return false
	}
	return true
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.UnrestrictedTokenTransfer {
		i--
		if m.UnrestrictedTokenTransfer {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.BaseTokenManager) > 0 {
		i -= len(m.BaseTokenManager)
		copy(dAtA[i:], m.BaseTokenManager)
		i = encodeVarintOpb(dAtA, i, uint64(len(m.BaseTokenManager)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.PointTokenDenom) > 0 {
		i -= len(m.PointTokenDenom)
		copy(dAtA[i:], m.PointTokenDenom)
		i = encodeVarintOpb(dAtA, i, uint64(len(m.PointTokenDenom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.BaseTokenDenom) > 0 {
		i -= len(m.BaseTokenDenom)
		copy(dAtA[i:], m.BaseTokenDenom)
		i = encodeVarintOpb(dAtA, i, uint64(len(m.BaseTokenDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOpb(dAtA []byte, offset int, v uint64) int {
	offset -= sovOpb(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BaseTokenDenom)
	if l > 0 {
		n += 1 + l + sovOpb(uint64(l))
	}
	l = len(m.PointTokenDenom)
	if l > 0 {
		n += 1 + l + sovOpb(uint64(l))
	}
	l = len(m.BaseTokenManager)
	if l > 0 {
		n += 1 + l + sovOpb(uint64(l))
	}
	if m.UnrestrictedTokenTransfer {
		n += 2
	}
	return n
}

func sovOpb(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOpb(x uint64) (n int) {
	return sovOpb(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOpb
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseTokenDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOpb
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOpb
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseTokenDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PointTokenDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOpb
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOpb
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PointTokenDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseTokenManager", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthOpb
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOpb
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BaseTokenManager = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field UnrestrictedTokenTransfer", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOpb
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.UnrestrictedTokenTransfer = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipOpb(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthOpb
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthOpb
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
func skipOpb(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOpb
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
					return 0, ErrIntOverflowOpb
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
					return 0, ErrIntOverflowOpb
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
				return 0, ErrInvalidLengthOpb
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOpb
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOpb
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOpb        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOpb          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOpb = fmt.Errorf("proto: unexpected end of group")
)

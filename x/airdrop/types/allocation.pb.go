// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nxtpop/airdrop/v1beta1/allocation.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

// AirdropAllocation defines the user's airdrop allocation.
type AirdropAllocation struct {
	Chain         string                                 `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain,omitempty"`
	Address       string                                 `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	Amount        github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,3,opt,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"amount"`
	ClaimedAmount github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,4,opt,name=claimed_amount,json=claimedAmount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"claimed_amount"`
}

func (m *AirdropAllocation) Reset()         { *m = AirdropAllocation{} }
func (m *AirdropAllocation) String() string { return proto.CompactTextString(m) }
func (*AirdropAllocation) ProtoMessage()    {}
func (*AirdropAllocation) Descriptor() ([]byte, []int) {
	return fileDescriptor_33f443a60d83adb6, []int{0}
}
func (m *AirdropAllocation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AirdropAllocation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AirdropAllocation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AirdropAllocation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AirdropAllocation.Merge(m, src)
}
func (m *AirdropAllocation) XXX_Size() int {
	return m.Size()
}
func (m *AirdropAllocation) XXX_DiscardUnknown() {
	xxx_messageInfo_AirdropAllocation.DiscardUnknown(m)
}

var xxx_messageInfo_AirdropAllocation proto.InternalMessageInfo

func (m *AirdropAllocation) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *AirdropAllocation) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*AirdropAllocation)(nil), "nxtpop.airdrop.v1beta1.AirdropAllocation")
}

func init() {
	proto.RegisterFile("nxtpop/airdrop/v1beta1/allocation.proto", fileDescriptor_33f443a60d83adb6)
}

var fileDescriptor_33f443a60d83adb6 = []byte{
	// 293 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0x1b, 0xff, 0x4c, 0x0c, 0x28, 0x58, 0x86, 0xd4, 0x1d, 0x32, 0xf1, 0xa0, 0x5e, 0x96,
	0x30, 0xbc, 0x78, 0xdd, 0x04, 0xc1, 0x93, 0xa3, 0xe2, 0xc5, 0xcb, 0x48, 0xd3, 0xd2, 0x15, 0xdb,
	0xbc, 0x25, 0xc9, 0x64, 0x7e, 0x0b, 0x3f, 0xd6, 0x8e, 0x3b, 0x8a, 0x87, 0x21, 0xed, 0x17, 0x91,
	0x25, 0x71, 0x78, 0xf6, 0x94, 0x3c, 0x79, 0x9e, 0xfc, 0x1e, 0xf2, 0x06, 0x5f, 0xc9, 0x85, 0xa9,
	0xa1, 0x66, 0xbc, 0x50, 0xa9, 0x82, 0x9a, 0xbd, 0x0d, 0x93, 0xcc, 0xf0, 0x21, 0xe3, 0x65, 0x09,
	0x82, 0x9b, 0x02, 0x24, 0xad, 0x15, 0x18, 0x08, 0x4f, 0x5d, 0x90, 0xfa, 0x20, 0xf5, 0xc1, 0x5e,
	0x37, 0x87, 0x1c, 0x6c, 0x84, 0x6d, 0x76, 0x2e, 0xdd, 0x3b, 0x13, 0xa0, 0x2b, 0xd0, 0x53, 0x67,
	0x38, 0xe1, 0xac, 0x8b, 0x06, 0xe1, 0x93, 0x91, 0x83, 0x8c, 0xb6, 0x25, 0x61, 0x17, 0xef, 0x8b,
	0x19, 0x2f, 0x64, 0x84, 0xce, 0xd1, 0xf5, 0x61, 0xec, 0x44, 0x18, 0xe1, 0x03, 0x9e, 0xa6, 0x2a,
	0xd3, 0x3a, 0xda, 0xb1, 0xe7, 0xbf, 0x32, 0xbc, 0xc7, 0x1d, 0x5e, 0xc1, 0x5c, 0x9a, 0x68, 0x77,
	0x63, 0x8c, 0xe9, 0x72, 0xdd, 0x0f, 0xbe, 0xd6, 0xfd, 0xcb, 0xbc, 0x30, 0xb3, 0x79, 0x42, 0x05,
	0x54, 0xbe, 0xd6, 0x2f, 0x03, 0x9d, 0xbe, 0x32, 0xf3, 0x5e, 0x67, 0x9a, 0x3e, 0x48, 0x13, 0xfb,
	0xdb, 0xe1, 0x33, 0x3e, 0x16, 0x25, 0x2f, 0xaa, 0x2c, 0x9d, 0x7a, 0xde, 0xde, 0xbf, 0x78, 0x47,
	0x9e, 0x32, 0xb2, 0x90, 0x71, 0xbc, 0x6c, 0x08, 0x5a, 0x35, 0x04, 0x7d, 0x37, 0x04, 0x7d, 0xb4,
	0x24, 0x58, 0xb5, 0x24, 0xf8, 0x6c, 0x49, 0xf0, 0x72, 0xfb, 0x07, 0x38, 0x79, 0x9c, 0x3c, 0x55,
	0x5c, 0x99, 0x3b, 0x90, 0x46, 0x71, 0x61, 0x98, 0x9b, 0xf1, 0xc0, 0x3e, 0x9e, 0x2d, 0xb6, 0x9f,
	0x62, 0x6b, 0x92, 0x8e, 0x9d, 0xdf, 0xcd, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x40, 0x8b, 0x33,
	0xce, 0xb3, 0x01, 0x00, 0x00,
}

func (m *AirdropAllocation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AirdropAllocation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AirdropAllocation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.ClaimedAmount.Size()
		i -= size
		if _, err := m.ClaimedAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAllocation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintAllocation(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintAllocation(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintAllocation(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAllocation(dAtA []byte, offset int, v uint64) int {
	offset -= sovAllocation(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AirdropAllocation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovAllocation(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovAllocation(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovAllocation(uint64(l))
	l = m.ClaimedAmount.Size()
	n += 1 + l + sovAllocation(uint64(l))
	return n
}

func sovAllocation(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAllocation(x uint64) (n int) {
	return sovAllocation(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AirdropAllocation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAllocation
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
			return fmt.Errorf("proto: AirdropAllocation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AirdropAllocation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllocation
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
				return ErrInvalidLengthAllocation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAllocation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllocation
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
				return ErrInvalidLengthAllocation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAllocation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllocation
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
				return ErrInvalidLengthAllocation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAllocation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimedAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllocation
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
				return ErrInvalidLengthAllocation
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAllocation
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ClaimedAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAllocation(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAllocation
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
func skipAllocation(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAllocation
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
					return 0, ErrIntOverflowAllocation
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
					return 0, ErrIntOverflowAllocation
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
				return 0, ErrInvalidLengthAllocation
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAllocation
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAllocation
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAllocation        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAllocation          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAllocation = fmt.Errorf("proto: unexpected end of group")
)

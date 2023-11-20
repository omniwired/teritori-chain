// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: teritori/mint/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// MsgBurnTokens defines an sdk.Msg type that burn tokens
type MsgBurnTokens struct {
	Sender string                                    `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	Amount []github_com_cosmos_cosmos_sdk_types.Coin `protobuf:"bytes,2,rep,name=amount,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Coin" json:"amount"`
}

func (m *MsgBurnTokens) Reset()         { *m = MsgBurnTokens{} }
func (m *MsgBurnTokens) String() string { return proto.CompactTextString(m) }
func (*MsgBurnTokens) ProtoMessage()    {}
func (*MsgBurnTokens) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2bf5271f1525b13, []int{0}
}
func (m *MsgBurnTokens) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBurnTokens) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBurnTokens.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBurnTokens) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBurnTokens.Merge(m, src)
}
func (m *MsgBurnTokens) XXX_Size() int {
	return m.Size()
}
func (m *MsgBurnTokens) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBurnTokens.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBurnTokens proto.InternalMessageInfo

func (m *MsgBurnTokens) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

// MsgBurnTokensResponse defines the Msg/BurnTokens response type.
type MsgBurnTokensResponse struct {
}

func (m *MsgBurnTokensResponse) Reset()         { *m = MsgBurnTokensResponse{} }
func (m *MsgBurnTokensResponse) String() string { return proto.CompactTextString(m) }
func (*MsgBurnTokensResponse) ProtoMessage()    {}
func (*MsgBurnTokensResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f2bf5271f1525b13, []int{1}
}
func (m *MsgBurnTokensResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgBurnTokensResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgBurnTokensResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgBurnTokensResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgBurnTokensResponse.Merge(m, src)
}
func (m *MsgBurnTokensResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgBurnTokensResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgBurnTokensResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgBurnTokensResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgBurnTokens)(nil), "teritori.mint.v1beta1.MsgBurnTokens")
	proto.RegisterType((*MsgBurnTokensResponse)(nil), "teritori.mint.v1beta1.MsgBurnTokensResponse")
}

func init() { proto.RegisterFile("teritori/mint/v1beta1/tx.proto", fileDescriptor_f2bf5271f1525b13) }

var fileDescriptor_f2bf5271f1525b13 = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xbf, 0x6a, 0xf3, 0x30,
	0x14, 0xc5, 0xed, 0x2f, 0x10, 0x88, 0xe0, 0x5b, 0x4c, 0xd3, 0xa6, 0x19, 0x94, 0x10, 0x0a, 0xcd,
	0xd0, 0xe8, 0x92, 0xf6, 0x0d, 0x52, 0x4a, 0xf1, 0x10, 0x0a, 0xc6, 0x53, 0xa7, 0xda, 0x8e, 0xaa,
	0x88, 0xd4, 0xba, 0xc6, 0x92, 0x4b, 0xf2, 0x16, 0x7d, 0xac, 0x8c, 0x19, 0x4b, 0x87, 0x50, 0xec,
	0x17, 0x29, 0xf1, 0x1f, 0x48, 0xa0, 0x43, 0x27, 0xe9, 0xf2, 0x3b, 0x3a, 0x47, 0x47, 0x22, 0xd4,
	0xf0, 0x54, 0x1a, 0x4c, 0x25, 0xc4, 0x52, 0x19, 0x78, 0x9f, 0x86, 0xdc, 0x04, 0x53, 0x30, 0x6b,
	0x96, 0xa4, 0x68, 0xd0, 0xe9, 0x36, 0x9c, 0x1d, 0x38, 0xab, 0x79, 0xff, 0x4c, 0xa0, 0xc0, 0x52,
	0x01, 0x87, 0x5d, 0x25, 0xee, 0x5f, 0x0a, 0x44, 0xf1, 0xc6, 0xa1, 0x9c, 0xc2, 0xec, 0x15, 0x02,
	0xb5, 0xa9, 0xd0, 0x28, 0x21, 0xff, 0xe7, 0x5a, 0xcc, 0xb2, 0x54, 0xf9, 0xb8, 0xe2, 0x4a, 0x3b,
	0xe7, 0xa4, 0xad, 0xb9, 0x5a, 0xf0, 0xb4, 0x67, 0x0f, 0xed, 0x71, 0xc7, 0xab, 0x27, 0xe7, 0x91,
	0xb4, 0x83, 0x18, 0x33, 0x65, 0x7a, 0xff, 0x86, 0xad, 0x71, 0x67, 0x06, 0xdb, 0xfd, 0xc0, 0xfa,
	0xda, 0x0f, 0xae, 0x85, 0x34, 0xcb, 0x2c, 0x64, 0x11, 0xc6, 0x10, 0xa1, 0x8e, 0x51, 0xd7, 0xcb,
	0x44, 0x2f, 0x56, 0x60, 0x36, 0x09, 0xd7, 0xec, 0x1e, 0xa5, 0xf2, 0xea, 0xe3, 0xa3, 0x0b, 0xd2,
	0x3d, 0x49, 0xf4, 0xb8, 0x4e, 0x50, 0x69, 0x7e, 0x2b, 0x48, 0x6b, 0xae, 0x85, 0xf3, 0x42, 0xc8,
	0xd1, 0x75, 0xae, 0xd8, 0xaf, 0x45, 0xd9, 0x89, 0x45, 0xff, 0xe6, 0x2f, 0xaa, 0x26, 0x68, 0xe6,
	0x6e, 0x73, 0x6a, 0xef, 0x72, 0x6a, 0x7f, 0xe7, 0xd4, 0xfe, 0x28, 0xa8, 0xb5, 0x2b, 0xa8, 0xf5,
	0x59, 0x50, 0xeb, 0x19, 0x8e, 0xca, 0xf8, 0x0f, 0x9e, 0xeb, 0x3f, 0x79, 0x2e, 0x34, 0xd6, 0x93,
	0x68, 0x19, 0x48, 0x05, 0xeb, 0xea, 0x47, 0xca, 0x66, 0x61, 0xbb, 0x7c, 0xc5, 0xbb, 0x9f, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x63, 0xdd, 0xd9, 0x2d, 0xaf, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// BurnTokens defines a method to burn tokens
	BurnTokens(ctx context.Context, in *MsgBurnTokens, opts ...grpc.CallOption) (*MsgBurnTokensResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) BurnTokens(ctx context.Context, in *MsgBurnTokens, opts ...grpc.CallOption) (*MsgBurnTokensResponse, error) {
	out := new(MsgBurnTokensResponse)
	err := c.cc.Invoke(ctx, "/teritori.mint.v1beta1.Msg/BurnTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// BurnTokens defines a method to burn tokens
	BurnTokens(context.Context, *MsgBurnTokens) (*MsgBurnTokensResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) BurnTokens(ctx context.Context, req *MsgBurnTokens) (*MsgBurnTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BurnTokens not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_BurnTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgBurnTokens)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).BurnTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/teritori.mint.v1beta1.Msg/BurnTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).BurnTokens(ctx, req.(*MsgBurnTokens))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "teritori.mint.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BurnTokens",
			Handler:    _Msg_BurnTokens_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "teritori/mint/v1beta1/tx.proto",
}

func (m *MsgBurnTokens) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBurnTokens) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBurnTokens) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.Amount[iNdEx].Size()
				i -= size
				if _, err := m.Amount[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgBurnTokensResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgBurnTokensResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgBurnTokensResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgBurnTokens) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	return n
}

func (m *MsgBurnTokensResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgBurnTokens) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgBurnTokens: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBurnTokens: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_cosmos_cosmos_sdk_types.Coin
			m.Amount = append(m.Amount, v)
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgBurnTokensResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgBurnTokensResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgBurnTokensResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)

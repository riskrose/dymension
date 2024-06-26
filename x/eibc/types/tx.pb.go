// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dymension/eibc/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
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

// MsgFulfillOrder defines the FulfillOrder request type.
type MsgFulfillOrder struct {
	// fulfiller_address is the bech32-encoded address of the account which the message was sent from.
	FulfillerAddress string `protobuf:"bytes,1,opt,name=fulfiller_address,json=fulfillerAddress,proto3" json:"fulfiller_address,omitempty"`
	OrderId          string `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
}

func (m *MsgFulfillOrder) Reset()         { *m = MsgFulfillOrder{} }
func (m *MsgFulfillOrder) String() string { return proto.CompactTextString(m) }
func (*MsgFulfillOrder) ProtoMessage()    {}
func (*MsgFulfillOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6b186a723f445fa, []int{0}
}
func (m *MsgFulfillOrder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgFulfillOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgFulfillOrder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgFulfillOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgFulfillOrder.Merge(m, src)
}
func (m *MsgFulfillOrder) XXX_Size() int {
	return m.Size()
}
func (m *MsgFulfillOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgFulfillOrder.DiscardUnknown(m)
}

var xxx_messageInfo_MsgFulfillOrder proto.InternalMessageInfo

func (m *MsgFulfillOrder) GetFulfillerAddress() string {
	if m != nil {
		return m.FulfillerAddress
	}
	return ""
}

func (m *MsgFulfillOrder) GetOrderId() string {
	if m != nil {
		return m.OrderId
	}
	return ""
}

// MsgFulfillOrderResponse defines the FulfillOrder response type.
type MsgFulfillOrderResponse struct {
}

func (m *MsgFulfillOrderResponse) Reset()         { *m = MsgFulfillOrderResponse{} }
func (m *MsgFulfillOrderResponse) String() string { return proto.CompactTextString(m) }
func (*MsgFulfillOrderResponse) ProtoMessage()    {}
func (*MsgFulfillOrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_b6b186a723f445fa, []int{1}
}
func (m *MsgFulfillOrderResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgFulfillOrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgFulfillOrderResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgFulfillOrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgFulfillOrderResponse.Merge(m, src)
}
func (m *MsgFulfillOrderResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgFulfillOrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgFulfillOrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgFulfillOrderResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgFulfillOrder)(nil), "dymensionxyz.dymension.eibc.MsgFulfillOrder")
	proto.RegisterType((*MsgFulfillOrderResponse)(nil), "dymensionxyz.dymension.eibc.MsgFulfillOrderResponse")
}

func init() { proto.RegisterFile("dymension/eibc/tx.proto", fileDescriptor_b6b186a723f445fa) }

var fileDescriptor_b6b186a723f445fa = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0xa9, 0xcc, 0x4d,
	0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0xd3, 0x4f, 0xcd, 0x4c, 0x4a, 0xd6, 0x2f, 0xa9, 0xd0, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x92, 0x86, 0x4b, 0x54, 0x54, 0x56, 0xe9, 0xc1, 0x39, 0x7a, 0x20, 0x55,
	0x4a, 0x91, 0x5c, 0xfc, 0xbe, 0xc5, 0xe9, 0x6e, 0xa5, 0x39, 0x69, 0x99, 0x39, 0x39, 0xfe, 0x45,
	0x29, 0xa9, 0x45, 0x42, 0xda, 0x5c, 0x82, 0x69, 0x10, 0x7e, 0x6a, 0x51, 0x7c, 0x62, 0x4a, 0x4a,
	0x51, 0x6a, 0x71, 0xb1, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x90, 0x00, 0x5c, 0xc2, 0x11, 0x22,
	0x2e, 0x24, 0xc9, 0xc5, 0x91, 0x0f, 0xd2, 0x15, 0x9f, 0x99, 0x22, 0xc1, 0x04, 0x56, 0xc3, 0x0e,
	0xe6, 0x7b, 0xa6, 0x28, 0x49, 0x72, 0x89, 0xa3, 0x19, 0x1d, 0x94, 0x5a, 0x5c, 0x90, 0x9f, 0x57,
	0x9c, 0x6a, 0x54, 0xcd, 0xc5, 0xec, 0x5b, 0x9c, 0x2e, 0x54, 0xc2, 0xc5, 0x83, 0x62, 0xb3, 0x8e,
	0x1e, 0x1e, 0xa7, 0xea, 0xa1, 0x19, 0x26, 0x65, 0x42, 0x8a, 0x6a, 0x98, 0xd5, 0x4a, 0x0c, 0x4e,
	0xde, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7,
	0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x65, 0x98, 0x9e, 0x59, 0x92,
	0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x8f, 0x6c, 0x36, 0x82, 0xa3, 0x5f, 0x66, 0xac, 0x5f,
	0x01, 0x0d, 0xdf, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0x70, 0x18, 0x1b, 0x03, 0x02, 0x00, 0x00,
	0xff, 0xff, 0x15, 0xb2, 0xf7, 0xb3, 0x7e, 0x01, 0x00, 0x00,
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
	FulfillOrder(ctx context.Context, in *MsgFulfillOrder, opts ...grpc.CallOption) (*MsgFulfillOrderResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) FulfillOrder(ctx context.Context, in *MsgFulfillOrder, opts ...grpc.CallOption) (*MsgFulfillOrderResponse, error) {
	out := new(MsgFulfillOrderResponse)
	err := c.cc.Invoke(ctx, "/dymensionxyz.dymension.eibc.Msg/FulfillOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	FulfillOrder(context.Context, *MsgFulfillOrder) (*MsgFulfillOrderResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) FulfillOrder(ctx context.Context, req *MsgFulfillOrder) (*MsgFulfillOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FulfillOrder not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_FulfillOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgFulfillOrder)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).FulfillOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dymensionxyz.dymension.eibc.Msg/FulfillOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).FulfillOrder(ctx, req.(*MsgFulfillOrder))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dymensionxyz.dymension.eibc.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FulfillOrder",
			Handler:    _Msg_FulfillOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dymension/eibc/tx.proto",
}

func (m *MsgFulfillOrder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgFulfillOrder) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgFulfillOrder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.OrderId) > 0 {
		i -= len(m.OrderId)
		copy(dAtA[i:], m.OrderId)
		i = encodeVarintTx(dAtA, i, uint64(len(m.OrderId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.FulfillerAddress) > 0 {
		i -= len(m.FulfillerAddress)
		copy(dAtA[i:], m.FulfillerAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.FulfillerAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgFulfillOrderResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgFulfillOrderResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgFulfillOrderResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgFulfillOrder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FulfillerAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.OrderId)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgFulfillOrderResponse) Size() (n int) {
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
func (m *MsgFulfillOrder) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgFulfillOrder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgFulfillOrder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FulfillerAddress", wireType)
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
			m.FulfillerAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderId", wireType)
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
			m.OrderId = string(dAtA[iNdEx:postIndex])
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
func (m *MsgFulfillOrderResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgFulfillOrderResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgFulfillOrderResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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

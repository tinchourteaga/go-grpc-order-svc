// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal/pb/order/order.proto

package order

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type CreateRequest struct {
	ProductId            int64    `protobuf:"varint,1,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity             int64    `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	UserId               int64    `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateRequest) Reset()         { *m = CreateRequest{} }
func (m *CreateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()    {}
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f91fa151f379689c, []int{0}
}

func (m *CreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateRequest.Unmarshal(m, b)
}
func (m *CreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateRequest.Marshal(b, m, deterministic)
}
func (m *CreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateRequest.Merge(m, src)
}
func (m *CreateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateRequest.Size(m)
}
func (m *CreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateRequest proto.InternalMessageInfo

func (m *CreateRequest) GetProductId() int64 {
	if m != nil {
		return m.ProductId
	}
	return 0
}

func (m *CreateRequest) GetQuantity() int64 {
	if m != nil {
		return m.Quantity
	}
	return 0
}

func (m *CreateRequest) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

type CreateResponse struct {
	Status               int64    `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Id                   int64    `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f91fa151f379689c, []int{1}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

func (m *CreateResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *CreateResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *CreateResponse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "order.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "order.CreateResponse")
}

func init() { proto.RegisterFile("internal/pb/order/order.proto", fileDescriptor_f91fa151f379689c) }

var fileDescriptor_f91fa151f379689c = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0xed, 0x96, 0xae, 0x76, 0xd4, 0x1e, 0xc6, 0x56, 0x4b, 0xa1, 0x20, 0x7b, 0xf2, 0xb4,
	0x0b, 0x7a, 0xf5, 0xa4, 0xa7, 0x82, 0x28, 0xac, 0x37, 0x2f, 0x92, 0x6e, 0xe6, 0x10, 0x90, 0x24,
	0x9d, 0x4c, 0x04, 0xff, 0xbd, 0x6c, 0xb2, 0x2e, 0xa8, 0x97, 0xc0, 0xf7, 0x86, 0xbc, 0x99, 0xf7,
	0x60, 0x6b, 0xac, 0x10, 0x5b, 0xf5, 0xd1, 0xf8, 0x7d, 0xe3, 0x58, 0x13, 0xe7, 0xb7, 0xf6, 0xec,
	0xc4, 0xe1, 0x2c, 0x41, 0xd5, 0xc1, 0xf9, 0x23, 0x93, 0x12, 0x6a, 0xe9, 0x10, 0x29, 0x08, 0x6e,
	0x01, 0x3c, 0x3b, 0x1d, 0x3b, 0x79, 0x37, 0x7a, 0x3d, 0xb9, 0x9e, 0xdc, 0x4c, 0xdb, 0xf9, 0xa0,
	0xec, 0x34, 0x6e, 0xe0, 0xe4, 0x10, 0x95, 0x15, 0x23, 0x5f, 0xeb, 0x22, 0x0d, 0x47, 0xc6, 0x2b,
	0x38, 0x8e, 0x81, 0xb8, 0xff, 0x37, 0x4d, 0xa3, 0xb2, 0xc7, 0x9d, 0xae, 0x9e, 0x61, 0xf1, 0xb3,
	0x24, 0x78, 0x67, 0x03, 0xe1, 0x25, 0x94, 0x41, 0x94, 0xc4, 0x30, 0x6c, 0x18, 0x08, 0x97, 0x30,
	0x23, 0x66, 0xc7, 0xc9, 0x7b, 0xde, 0x66, 0xc0, 0x05, 0x14, 0xa3, 0x67, 0x61, 0xf4, 0xed, 0x13,
	0x9c, 0xbd, 0xf4, 0xd7, 0xbf, 0x12, 0x7f, 0x9a, 0x8e, 0xf0, 0x1e, 0x4e, 0xb3, 0x7f, 0x52, 0x71,
	0x59, 0xe7, 0xa0, 0xbf, 0x82, 0x6d, 0x56, 0x7f, 0xd4, 0x7c, 0x49, 0x75, 0xf4, 0xb0, 0x7a, 0xbb,
	0xa8, 0x9b, 0x7f, 0x65, 0xed, 0xcb, 0xd4, 0xd3, 0xdd, 0x77, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe8,
	0x40, 0xcf, 0x7c, 0x48, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderServiceClient interface {
	CreateOrder(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
}

type orderServiceClient struct {
	cc *grpc.ClientConn
}

func NewOrderServiceClient(cc *grpc.ClientConn) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) CreateOrder(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/order.OrderService/CreateOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
type OrderServiceServer interface {
	CreateOrder(context.Context, *CreateRequest) (*CreateResponse, error)
}

// UnimplementedOrderServiceServer can be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (*UnimplementedOrderServiceServer) CreateOrder(ctx context.Context, req *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}

func RegisterOrderServiceServer(s *grpc.Server, srv OrderServiceServer) {
	s.RegisterService(&_OrderService_serviceDesc, srv)
}

func _OrderService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.OrderService/CreateOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateOrder(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _OrderService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "order.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _OrderService_CreateOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/pb/order/order.proto",
}

// Code generated by protoc-gen-go.
// source: add.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	add.proto

It has these top-level messages:
	SumRequest
	SumReply
	ConcatRequest
	ConcatReply
*/
package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// The sum request contains two parameters.
type SumRequest struct {
	A int64 `protobuf:"varint,1,opt,name=a" json:"a,omitempty"`
	B int64 `protobuf:"varint,2,opt,name=b" json:"b,omitempty"`
}

func (m *SumRequest) Reset()                    { *m = SumRequest{} }
func (m *SumRequest) String() string            { return proto.CompactTextString(m) }
func (*SumRequest) ProtoMessage()               {}
func (*SumRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// The sum response contains the result of the calculation.
type SumReply struct {
	V int64 `protobuf:"varint,1,opt,name=v" json:"v,omitempty"`
}

func (m *SumReply) Reset()                    { *m = SumReply{} }
func (m *SumReply) String() string            { return proto.CompactTextString(m) }
func (*SumReply) ProtoMessage()               {}
func (*SumReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// The Concat request contains two parameters.
type ConcatRequest struct {
	A string `protobuf:"bytes,1,opt,name=a" json:"a,omitempty"`
	B string `protobuf:"bytes,2,opt,name=b" json:"b,omitempty"`
}

func (m *ConcatRequest) Reset()                    { *m = ConcatRequest{} }
func (m *ConcatRequest) String() string            { return proto.CompactTextString(m) }
func (*ConcatRequest) ProtoMessage()               {}
func (*ConcatRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// The Concat response contains the result of the concatenation.
type ConcatReply struct {
	V string `protobuf:"bytes,1,opt,name=v" json:"v,omitempty"`
}

func (m *ConcatReply) Reset()                    { *m = ConcatReply{} }
func (m *ConcatReply) String() string            { return proto.CompactTextString(m) }
func (*ConcatReply) ProtoMessage()               {}
func (*ConcatReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func init() {
	proto.RegisterType((*SumRequest)(nil), "pb.SumRequest")
	proto.RegisterType((*SumReply)(nil), "pb.SumReply")
	proto.RegisterType((*ConcatRequest)(nil), "pb.ConcatRequest")
	proto.RegisterType((*ConcatReply)(nil), "pb.ConcatReply")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion2

// Client API for Add service

type AddClient interface {
	// Sums two integers.
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumReply, error)
	// Concatenates two strings
	Concat(ctx context.Context, in *ConcatRequest, opts ...grpc.CallOption) (*ConcatReply, error)
}

type addClient struct {
	cc *grpc.ClientConn
}

func NewAddClient(cc *grpc.ClientConn) AddClient {
	return &addClient{cc}
}

func (c *addClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumReply, error) {
	out := new(SumReply)
	err := grpc.Invoke(ctx, "/pb.Add/Sum", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addClient) Concat(ctx context.Context, in *ConcatRequest, opts ...grpc.CallOption) (*ConcatReply, error) {
	out := new(ConcatReply)
	err := grpc.Invoke(ctx, "/pb.Add/Concat", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Add service

type AddServer interface {
	// Sums two integers.
	Sum(context.Context, *SumRequest) (*SumReply, error)
	// Concatenates two strings
	Concat(context.Context, *ConcatRequest) (*ConcatReply, error)
}

func RegisterAddServer(s *grpc.Server, srv AddServer) {
	s.RegisterService(&_Add_serviceDesc, srv)
}

func _Add_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Add/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Add_Concat_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConcatRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddServer).Concat(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Add/Concat",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddServer).Concat(ctx, req.(*ConcatRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Add_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Add",
	HandlerType: (*AddServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _Add_Sum_Handler,
		},
		{
			MethodName: "Concat",
			Handler:    _Add_Concat_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}

var fileDescriptor0 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x4c, 0x49, 0xd1,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xd2, 0xe0, 0xe2, 0x0a, 0x2e, 0xcd,
	0x0d, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0xe2, 0xe1, 0x62, 0x4c, 0x94, 0x60, 0x54, 0x60,
	0xd4, 0x60, 0x0e, 0x62, 0x4c, 0x04, 0xf1, 0x92, 0x24, 0x98, 0x20, 0xbc, 0x24, 0x25, 0x09, 0x2e,
	0x0e, 0xb0, 0xca, 0x82, 0x9c, 0x4a, 0x90, 0x4c, 0x19, 0x4c, 0x5d, 0x99, 0x92, 0x36, 0x17, 0xaf,
	0x73, 0x7e, 0x5e, 0x72, 0x62, 0x09, 0x86, 0x31, 0x9c, 0x28, 0xc6, 0x70, 0x82, 0x8c, 0x91, 0xe6,
	0xe2, 0x86, 0x29, 0x46, 0x31, 0x09, 0x28, 0x59, 0x66, 0x14, 0xc3, 0xc5, 0xec, 0x98, 0x92, 0x22,
	0xa4, 0xca, 0xc5, 0x0c, 0xb4, 0x4a, 0x88, 0x4f, 0xaf, 0x20, 0x49, 0x0f, 0xe1, 0x3a, 0x29, 0x1e,
	0x38, 0x1f, 0xa8, 0x53, 0x89, 0x41, 0x48, 0x8f, 0x8b, 0x0d, 0x62, 0x94, 0x90, 0x20, 0x48, 0x06,
	0xc5, 0x0d, 0x52, 0xfc, 0xc8, 0x42, 0x60, 0xf5, 0x49, 0x6c, 0x60, 0x6f, 0x1b, 0x03, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xb4, 0xc9, 0xe7, 0x58, 0x03, 0x01, 0x00, 0x00,
}

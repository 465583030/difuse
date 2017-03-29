// Code generated by protoc-gen-go.
// source: rpc.proto
// DO NOT EDIT!

/*
Package rpc is a generated protocol buffer package.

It is generated from these files:
	rpc.proto

It has these top-level messages:
*/
package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import types "github.com/ipkg/difuse/types"

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
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DifuseRPC service

type DifuseRPCClient interface {
	GetTxBlockServe(ctx context.Context, in *types.VnodeBytes, opts ...grpc.CallOption) (*types.TxBlock, error)
	GetTxServe(ctx context.Context, in *types.VnodeBytes, opts ...grpc.CallOption) (*types.Tx, error)
	NewTxServe(ctx context.Context, in *types.VnodeBytes, opts ...grpc.CallOption) (*types.Tx, error)
	ProposeTxServe(ctx context.Context, in *types.VnodeTx, opts ...grpc.CallOption) (*types.VnodeBytes, error)
	TakeoverTxBlocksServe(ctx context.Context, opts ...grpc.CallOption) (DifuseRPC_TakeoverTxBlocksServeClient, error)
}

type difuseRPCClient struct {
	cc *grpc.ClientConn
}

func NewDifuseRPCClient(cc *grpc.ClientConn) DifuseRPCClient {
	return &difuseRPCClient{cc}
}

func (c *difuseRPCClient) GetTxBlockServe(ctx context.Context, in *types.VnodeBytes, opts ...grpc.CallOption) (*types.TxBlock, error) {
	out := new(types.TxBlock)
	err := grpc.Invoke(ctx, "/rpc.DifuseRPC/GetTxBlockServe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *difuseRPCClient) GetTxServe(ctx context.Context, in *types.VnodeBytes, opts ...grpc.CallOption) (*types.Tx, error) {
	out := new(types.Tx)
	err := grpc.Invoke(ctx, "/rpc.DifuseRPC/GetTxServe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *difuseRPCClient) NewTxServe(ctx context.Context, in *types.VnodeBytes, opts ...grpc.CallOption) (*types.Tx, error) {
	out := new(types.Tx)
	err := grpc.Invoke(ctx, "/rpc.DifuseRPC/NewTxServe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *difuseRPCClient) ProposeTxServe(ctx context.Context, in *types.VnodeTx, opts ...grpc.CallOption) (*types.VnodeBytes, error) {
	out := new(types.VnodeBytes)
	err := grpc.Invoke(ctx, "/rpc.DifuseRPC/ProposeTxServe", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *difuseRPCClient) TakeoverTxBlocksServe(ctx context.Context, opts ...grpc.CallOption) (DifuseRPC_TakeoverTxBlocksServeClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_DifuseRPC_serviceDesc.Streams[0], c.cc, "/rpc.DifuseRPC/TakeoverTxBlocksServe", opts...)
	if err != nil {
		return nil, err
	}
	x := &difuseRPCTakeoverTxBlocksServeClient{stream}
	return x, nil
}

type DifuseRPC_TakeoverTxBlocksServeClient interface {
	Send(*types.TxBlock) error
	CloseAndRecv() (*types.VnodeBytes, error)
	grpc.ClientStream
}

type difuseRPCTakeoverTxBlocksServeClient struct {
	grpc.ClientStream
}

func (x *difuseRPCTakeoverTxBlocksServeClient) Send(m *types.TxBlock) error {
	return x.ClientStream.SendMsg(m)
}

func (x *difuseRPCTakeoverTxBlocksServeClient) CloseAndRecv() (*types.VnodeBytes, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(types.VnodeBytes)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for DifuseRPC service

type DifuseRPCServer interface {
	GetTxBlockServe(context.Context, *types.VnodeBytes) (*types.TxBlock, error)
	GetTxServe(context.Context, *types.VnodeBytes) (*types.Tx, error)
	NewTxServe(context.Context, *types.VnodeBytes) (*types.Tx, error)
	ProposeTxServe(context.Context, *types.VnodeTx) (*types.VnodeBytes, error)
	TakeoverTxBlocksServe(DifuseRPC_TakeoverTxBlocksServeServer) error
}

func RegisterDifuseRPCServer(s *grpc.Server, srv DifuseRPCServer) {
	s.RegisterService(&_DifuseRPC_serviceDesc, srv)
}

func _DifuseRPC_GetTxBlockServe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.VnodeBytes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DifuseRPCServer).GetTxBlockServe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.DifuseRPC/GetTxBlockServe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DifuseRPCServer).GetTxBlockServe(ctx, req.(*types.VnodeBytes))
	}
	return interceptor(ctx, in, info, handler)
}

func _DifuseRPC_GetTxServe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.VnodeBytes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DifuseRPCServer).GetTxServe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.DifuseRPC/GetTxServe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DifuseRPCServer).GetTxServe(ctx, req.(*types.VnodeBytes))
	}
	return interceptor(ctx, in, info, handler)
}

func _DifuseRPC_NewTxServe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.VnodeBytes)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DifuseRPCServer).NewTxServe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.DifuseRPC/NewTxServe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DifuseRPCServer).NewTxServe(ctx, req.(*types.VnodeBytes))
	}
	return interceptor(ctx, in, info, handler)
}

func _DifuseRPC_ProposeTxServe_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(types.VnodeTx)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DifuseRPCServer).ProposeTxServe(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.DifuseRPC/ProposeTxServe",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DifuseRPCServer).ProposeTxServe(ctx, req.(*types.VnodeTx))
	}
	return interceptor(ctx, in, info, handler)
}

func _DifuseRPC_TakeoverTxBlocksServe_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DifuseRPCServer).TakeoverTxBlocksServe(&difuseRPCTakeoverTxBlocksServeServer{stream})
}

type DifuseRPC_TakeoverTxBlocksServeServer interface {
	SendAndClose(*types.VnodeBytes) error
	Recv() (*types.TxBlock, error)
	grpc.ServerStream
}

type difuseRPCTakeoverTxBlocksServeServer struct {
	grpc.ServerStream
}

func (x *difuseRPCTakeoverTxBlocksServeServer) SendAndClose(m *types.VnodeBytes) error {
	return x.ServerStream.SendMsg(m)
}

func (x *difuseRPCTakeoverTxBlocksServeServer) Recv() (*types.TxBlock, error) {
	m := new(types.TxBlock)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _DifuseRPC_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.DifuseRPC",
	HandlerType: (*DifuseRPCServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTxBlockServe",
			Handler:    _DifuseRPC_GetTxBlockServe_Handler,
		},
		{
			MethodName: "GetTxServe",
			Handler:    _DifuseRPC_GetTxServe_Handler,
		},
		{
			MethodName: "NewTxServe",
			Handler:    _DifuseRPC_NewTxServe_Handler,
		},
		{
			MethodName: "ProposeTxServe",
			Handler:    _DifuseRPC_ProposeTxServe_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "TakeoverTxBlocksServe",
			Handler:       _DifuseRPC_TakeoverTxBlocksServe_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "rpc.proto",
}

func init() { proto.RegisterFile("rpc.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 196 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2c, 0x2a, 0x48, 0xd6,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2e, 0x2a, 0x48, 0x96, 0xd2, 0x48, 0xcf, 0x2c, 0xc9,
	0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xcf, 0x2c, 0xc8, 0x4e, 0xd7, 0x4f, 0xc9, 0x4c, 0x2b,
	0x2d, 0x4e, 0xd5, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0x86, 0x90, 0x10, 0xe5, 0x46, 0xd3, 0x98, 0xb8,
	0x38, 0x5d, 0xc0, 0x92, 0x41, 0x01, 0xce, 0x42, 0x66, 0x5c, 0xfc, 0xee, 0xa9, 0x25, 0x21, 0x15,
	0x4e, 0x39, 0xf9, 0xc9, 0xd9, 0xc1, 0xa9, 0x45, 0x65, 0xa9, 0x42, 0x82, 0x7a, 0x10, 0xe5, 0x61,
	0x79, 0xf9, 0x29, 0xa9, 0x4e, 0x95, 0x25, 0xa9, 0xc5, 0x52, 0x7c, 0x50, 0x21, 0xa8, 0x3a, 0x25,
	0x06, 0x21, 0x1d, 0x2e, 0x2e, 0xb0, 0x3e, 0x9c, 0x5a, 0x38, 0xe1, 0x5a, 0x20, 0xaa, 0xfd, 0x52,
	0xcb, 0x89, 0x55, 0x6d, 0xca, 0xc5, 0x17, 0x50, 0x94, 0x5f, 0x90, 0x5f, 0x9c, 0x0a, 0xd3, 0xc1,
	0x87, 0xac, 0x23, 0xa4, 0x42, 0x0a, 0xd3, 0x04, 0x25, 0x06, 0x21, 0x3b, 0x2e, 0xd1, 0x90, 0xc4,
	0xec, 0xd4, 0xfc, 0xb2, 0xd4, 0x22, 0xa8, 0x3b, 0x8b, 0x51, 0x75, 0x43, 0x45, 0xb1, 0xea, 0xd6,
	0x60, 0x4c, 0x62, 0x03, 0x87, 0x8f, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xa3, 0xd2, 0xca, 0xc3,
	0x5b, 0x01, 0x00, 0x00,
}

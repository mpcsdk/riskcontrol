// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: tfa/v1/tfa.proto

package v1

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// TFAClient is the client API for TFA service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TFAClient interface {
	CallTFAInfo(ctx context.Context, in *TFAReq, opts ...grpc.CallOption) (*TFARes, error)
}

type tFAClient struct {
	cc grpc.ClientConnInterface
}

func NewTFAClient(cc grpc.ClientConnInterface) TFAClient {
	return &tFAClient{cc}
}

func (c *tFAClient) CallTFAInfo(ctx context.Context, in *TFAReq, opts ...grpc.CallOption) (*TFARes, error) {
	out := new(TFARes)
	err := c.cc.Invoke(ctx, "/tfa.TFA/CallTFAInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TFAServer is the server API for TFA service.
// All implementations must embed UnimplementedTFAServer
// for forward compatibility
type TFAServer interface {
	CallTFAInfo(context.Context, *TFAReq) (*TFARes, error)
	mustEmbedUnimplementedTFAServer()
}

// UnimplementedTFAServer must be embedded to have forward compatible implementations.
type UnimplementedTFAServer struct {
}

func (UnimplementedTFAServer) CallTFAInfo(context.Context, *TFAReq) (*TFARes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CallTFAInfo not implemented")
}
func (UnimplementedTFAServer) mustEmbedUnimplementedTFAServer() {}

// UnsafeTFAServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TFAServer will
// result in compilation errors.
type UnsafeTFAServer interface {
	mustEmbedUnimplementedTFAServer()
}

func RegisterTFAServer(s grpc.ServiceRegistrar, srv TFAServer) {
	s.RegisterService(&TFA_ServiceDesc, srv)
}

func _TFA_CallTFAInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TFAReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TFAServer).CallTFAInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tfa.TFA/CallTFAInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TFAServer).CallTFAInfo(ctx, req.(*TFAReq))
	}
	return interceptor(ctx, in, info, handler)
}

// TFA_ServiceDesc is the grpc.ServiceDesc for TFA service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TFA_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tfa.TFA",
	HandlerType: (*TFAServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CallTFAInfo",
			Handler:    _TFA_CallTFAInfo_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tfa/v1/tfa.proto",
}

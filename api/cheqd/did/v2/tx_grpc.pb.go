// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: cheqd/did/v2/tx.proto

package didv2

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

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MsgClient interface {
	CreateDidDoc(ctx context.Context, in *MsgCreateDidDoc, opts ...grpc.CallOption) (*MsgCreateDidDocResponse, error)
	UpdateDidDoc(ctx context.Context, in *MsgUpdateDidDoc, opts ...grpc.CallOption) (*MsgUpdateDidDocResponse, error)
	DeactivateDidDoc(ctx context.Context, in *MsgDeactivateDidDoc, opts ...grpc.CallOption) (*MsgDeactivateDidDocResponse, error)
}

type msgClient struct {
	cc grpc.ClientConnInterface
}

func NewMsgClient(cc grpc.ClientConnInterface) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateDidDoc(ctx context.Context, in *MsgCreateDidDoc, opts ...grpc.CallOption) (*MsgCreateDidDocResponse, error) {
	out := new(MsgCreateDidDocResponse)
	err := c.cc.Invoke(ctx, "/cheqd.did.v2.Msg/CreateDidDoc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateDidDoc(ctx context.Context, in *MsgUpdateDidDoc, opts ...grpc.CallOption) (*MsgUpdateDidDocResponse, error) {
	out := new(MsgUpdateDidDocResponse)
	err := c.cc.Invoke(ctx, "/cheqd.did.v2.Msg/UpdateDidDoc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeactivateDidDoc(ctx context.Context, in *MsgDeactivateDidDoc, opts ...grpc.CallOption) (*MsgDeactivateDidDocResponse, error) {
	out := new(MsgDeactivateDidDocResponse)
	err := c.cc.Invoke(ctx, "/cheqd.did.v2.Msg/DeactivateDidDoc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
// All implementations must embed UnimplementedMsgServer
// for forward compatibility
type MsgServer interface {
	CreateDidDoc(context.Context, *MsgCreateDidDoc) (*MsgCreateDidDocResponse, error)
	UpdateDidDoc(context.Context, *MsgUpdateDidDoc) (*MsgUpdateDidDocResponse, error)
	DeactivateDidDoc(context.Context, *MsgDeactivateDidDoc) (*MsgDeactivateDidDocResponse, error)
	mustEmbedUnimplementedMsgServer()
}

// UnimplementedMsgServer must be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (UnimplementedMsgServer) CreateDidDoc(context.Context, *MsgCreateDidDoc) (*MsgCreateDidDocResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDidDoc not implemented")
}
func (UnimplementedMsgServer) UpdateDidDoc(context.Context, *MsgUpdateDidDoc) (*MsgUpdateDidDocResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDidDoc not implemented")
}
func (UnimplementedMsgServer) DeactivateDidDoc(context.Context, *MsgDeactivateDidDoc) (*MsgDeactivateDidDocResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeactivateDidDoc not implemented")
}
func (UnimplementedMsgServer) mustEmbedUnimplementedMsgServer() {}

// UnsafeMsgServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MsgServer will
// result in compilation errors.
type UnsafeMsgServer interface {
	mustEmbedUnimplementedMsgServer()
}

func RegisterMsgServer(s grpc.ServiceRegistrar, srv MsgServer) {
	s.RegisterService(&Msg_ServiceDesc, srv)
}

func _Msg_CreateDidDoc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateDidDoc)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateDidDoc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheqd.did.v2.Msg/CreateDidDoc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateDidDoc(ctx, req.(*MsgCreateDidDoc))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateDidDoc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateDidDoc)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateDidDoc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheqd.did.v2.Msg/UpdateDidDoc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateDidDoc(ctx, req.(*MsgUpdateDidDoc))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeactivateDidDoc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeactivateDidDoc)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeactivateDidDoc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cheqd.did.v2.Msg/DeactivateDidDoc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeactivateDidDoc(ctx, req.(*MsgDeactivateDidDoc))
	}
	return interceptor(ctx, in, info, handler)
}

// Msg_ServiceDesc is the grpc.ServiceDesc for Msg service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Msg_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "cheqd.did.v2.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDidDoc",
			Handler:    _Msg_CreateDidDoc_Handler,
		},
		{
			MethodName: "UpdateDidDoc",
			Handler:    _Msg_UpdateDidDoc_Handler,
		},
		{
			MethodName: "DeactivateDidDoc",
			Handler:    _Msg_DeactivateDidDoc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cheqd/did/v2/tx.proto",
}

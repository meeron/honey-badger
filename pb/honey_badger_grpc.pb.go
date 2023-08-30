// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.2
// source: pb/honey_badger.proto

package pb

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

// HoneyBadgerClient is the client API for HoneyBadger service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type HoneyBadgerClient interface {
	Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*Result, error)
	Get(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*GetResult, error)
	GetByPrefix(ctx context.Context, in *PrefixRequest, opts ...grpc.CallOption) (*PrefixResult, error)
	Delete(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*Result, error)
	DeleteByPrefix(ctx context.Context, in *PrefixRequest, opts ...grpc.CallOption) (*Result, error)
}

type honeyBadgerClient struct {
	cc grpc.ClientConnInterface
}

func NewHoneyBadgerClient(cc grpc.ClientConnInterface) HoneyBadgerClient {
	return &honeyBadgerClient{cc}
}

func (c *honeyBadgerClient) Set(ctx context.Context, in *SetRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/pb.HoneyBadger/Set", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *honeyBadgerClient) Get(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*GetResult, error) {
	out := new(GetResult)
	err := c.cc.Invoke(ctx, "/pb.HoneyBadger/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *honeyBadgerClient) GetByPrefix(ctx context.Context, in *PrefixRequest, opts ...grpc.CallOption) (*PrefixResult, error) {
	out := new(PrefixResult)
	err := c.cc.Invoke(ctx, "/pb.HoneyBadger/GetByPrefix", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *honeyBadgerClient) Delete(ctx context.Context, in *KeyRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/pb.HoneyBadger/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *honeyBadgerClient) DeleteByPrefix(ctx context.Context, in *PrefixRequest, opts ...grpc.CallOption) (*Result, error) {
	out := new(Result)
	err := c.cc.Invoke(ctx, "/pb.HoneyBadger/DeleteByPrefix", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// HoneyBadgerServer is the server API for HoneyBadger service.
// All implementations must embed UnimplementedHoneyBadgerServer
// for forward compatibility
type HoneyBadgerServer interface {
	Set(context.Context, *SetRequest) (*Result, error)
	Get(context.Context, *KeyRequest) (*GetResult, error)
	GetByPrefix(context.Context, *PrefixRequest) (*PrefixResult, error)
	Delete(context.Context, *KeyRequest) (*Result, error)
	DeleteByPrefix(context.Context, *PrefixRequest) (*Result, error)
	mustEmbedUnimplementedHoneyBadgerServer()
}

// UnimplementedHoneyBadgerServer must be embedded to have forward compatible implementations.
type UnimplementedHoneyBadgerServer struct {
}

func (UnimplementedHoneyBadgerServer) Set(context.Context, *SetRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Set not implemented")
}
func (UnimplementedHoneyBadgerServer) Get(context.Context, *KeyRequest) (*GetResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedHoneyBadgerServer) GetByPrefix(context.Context, *PrefixRequest) (*PrefixResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByPrefix not implemented")
}
func (UnimplementedHoneyBadgerServer) Delete(context.Context, *KeyRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedHoneyBadgerServer) DeleteByPrefix(context.Context, *PrefixRequest) (*Result, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteByPrefix not implemented")
}
func (UnimplementedHoneyBadgerServer) mustEmbedUnimplementedHoneyBadgerServer() {}

// UnsafeHoneyBadgerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to HoneyBadgerServer will
// result in compilation errors.
type UnsafeHoneyBadgerServer interface {
	mustEmbedUnimplementedHoneyBadgerServer()
}

func RegisterHoneyBadgerServer(s grpc.ServiceRegistrar, srv HoneyBadgerServer) {
	s.RegisterService(&HoneyBadger_ServiceDesc, srv)
}

func _HoneyBadger_Set_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoneyBadgerServer).Set(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HoneyBadger/Set",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoneyBadgerServer).Set(ctx, req.(*SetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HoneyBadger_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoneyBadgerServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HoneyBadger/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoneyBadgerServer).Get(ctx, req.(*KeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HoneyBadger_GetByPrefix_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrefixRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoneyBadgerServer).GetByPrefix(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HoneyBadger/GetByPrefix",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoneyBadgerServer).GetByPrefix(ctx, req.(*PrefixRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HoneyBadger_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(KeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoneyBadgerServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HoneyBadger/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoneyBadgerServer).Delete(ctx, req.(*KeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _HoneyBadger_DeleteByPrefix_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrefixRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(HoneyBadgerServer).DeleteByPrefix(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.HoneyBadger/DeleteByPrefix",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(HoneyBadgerServer).DeleteByPrefix(ctx, req.(*PrefixRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// HoneyBadger_ServiceDesc is the grpc.ServiceDesc for HoneyBadger service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var HoneyBadger_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.HoneyBadger",
	HandlerType: (*HoneyBadgerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Set",
			Handler:    _HoneyBadger_Set_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _HoneyBadger_Get_Handler,
		},
		{
			MethodName: "GetByPrefix",
			Handler:    _HoneyBadger_GetByPrefix_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _HoneyBadger_Delete_Handler,
		},
		{
			MethodName: "DeleteByPrefix",
			Handler:    _HoneyBadger_DeleteByPrefix_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/honey_badger.proto",
}

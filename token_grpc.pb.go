// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package go_tokenmgmt_grpc

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

// TknClient is the client API for Tkn service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TknClient interface {
	CreateNewToken(ctx context.Context, in *NewToken, opts ...grpc.CallOption) (*Token, error)
	GetToekns(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Token, error)
	DropToken(ctx context.Context, in *TokenInfo, opts ...grpc.CallOption) (*EmptyToken, error)
	WriteToken(ctx context.Context, in *NewToken, opts ...grpc.CallOption) (*Token, error)
}

type tknClient struct {
	cc grpc.ClientConnInterface
}

func NewTknClient(cc grpc.ClientConnInterface) TknClient {
	return &tknClient{cc}
}

func (c *tknClient) CreateNewToken(ctx context.Context, in *NewToken, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/tokenmgmt.Tkn/CreateNewToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tknClient) GetToekns(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/tokenmgmt.Tkn/GetToekns", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tknClient) DropToken(ctx context.Context, in *TokenInfo, opts ...grpc.CallOption) (*EmptyToken, error) {
	out := new(EmptyToken)
	err := c.cc.Invoke(ctx, "/tokenmgmt.Tkn/DropToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tknClient) WriteToken(ctx context.Context, in *NewToken, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/tokenmgmt.Tkn/WriteToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TknServer is the server API for Tkn service.
// All implementations must embed UnimplementedTknServer
// for forward compatibility
type TknServer interface {
	CreateNewToken(context.Context, *NewToken) (*Token, error)
	GetToekns(context.Context, *Token) (*Token, error)
	DropToken(context.Context, *TokenInfo) (*EmptyToken, error)
	WriteToken(context.Context, *NewToken) (*Token, error)
	mustEmbedUnimplementedTknServer()
}

// UnimplementedTknServer must be embedded to have forward compatible implementations.
type UnimplementedTknServer struct {
}

func (UnimplementedTknServer) CreateNewToken(context.Context, *NewToken) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewToken not implemented")
}
func (UnimplementedTknServer) GetToekns(context.Context, *Token) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetToekns not implemented")
}
func (UnimplementedTknServer) DropToken(context.Context, *TokenInfo) (*EmptyToken, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DropToken not implemented")
}
func (UnimplementedTknServer) WriteToken(context.Context, *NewToken) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteToken not implemented")
}
func (UnimplementedTknServer) mustEmbedUnimplementedTknServer() {}

// UnsafeTknServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TknServer will
// result in compilation errors.
type UnsafeTknServer interface {
	mustEmbedUnimplementedTknServer()
}

func RegisterTknServer(s grpc.ServiceRegistrar, srv TknServer) {
	s.RegisterService(&Tkn_ServiceDesc, srv)
}

func _Tkn_CreateNewToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TknServer).CreateNewToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenmgmt.Tkn/CreateNewToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TknServer).CreateNewToken(ctx, req.(*NewToken))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tkn_GetToekns_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TknServer).GetToekns(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenmgmt.Tkn/GetToekns",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TknServer).GetToekns(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tkn_DropToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenInfo)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TknServer).DropToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenmgmt.Tkn/DropToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TknServer).DropToken(ctx, req.(*TokenInfo))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tkn_WriteToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewToken)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TknServer).WriteToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tokenmgmt.Tkn/WriteToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TknServer).WriteToken(ctx, req.(*NewToken))
	}
	return interceptor(ctx, in, info, handler)
}

// Tkn_ServiceDesc is the grpc.ServiceDesc for Tkn service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Tkn_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tokenmgmt.Tkn",
	HandlerType: (*TknServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNewToken",
			Handler:    _Tkn_CreateNewToken_Handler,
		},
		{
			MethodName: "GetToekns",
			Handler:    _Tkn_GetToekns_Handler,
		},
		{
			MethodName: "DropToken",
			Handler:    _Tkn_DropToken_Handler,
		},
		{
			MethodName: "WriteToken",
			Handler:    _Tkn_WriteToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/token.proto",
}
// Version

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: modules/auth/authPb/authPb.proto

package hello_sekai_shop_tutorial

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

const (
	AuthGrpcService_AccessTokenSearch_FullMethodName = "/AuthGrpcService/AccessTokenSearch"
	AuthGrpcService_RolesCount_FullMethodName        = "/AuthGrpcService/RolesCount"
)

// AuthGrpcServiceClient is the client API for AuthGrpcService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AuthGrpcServiceClient interface {
	AccessTokenSearch(ctx context.Context, in *AccessTokenSearchReq, opts ...grpc.CallOption) (*AccessTokenSearchRes, error)
	RolesCount(ctx context.Context, in *RolesCountReq, opts ...grpc.CallOption) (*RolesCountRes, error)
}

type authGrpcServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAuthGrpcServiceClient(cc grpc.ClientConnInterface) AuthGrpcServiceClient {
	return &authGrpcServiceClient{cc}
}

func (c *authGrpcServiceClient) AccessTokenSearch(ctx context.Context, in *AccessTokenSearchReq, opts ...grpc.CallOption) (*AccessTokenSearchRes, error) {
	out := new(AccessTokenSearchRes)
	err := c.cc.Invoke(ctx, AuthGrpcService_AccessTokenSearch_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authGrpcServiceClient) RolesCount(ctx context.Context, in *RolesCountReq, opts ...grpc.CallOption) (*RolesCountRes, error) {
	out := new(RolesCountRes)
	err := c.cc.Invoke(ctx, AuthGrpcService_RolesCount_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthGrpcServiceServer is the server API for AuthGrpcService service.
// All implementations must embed UnimplementedAuthGrpcServiceServer
// for forward compatibility
type AuthGrpcServiceServer interface {
	AccessTokenSearch(context.Context, *AccessTokenSearchReq) (*AccessTokenSearchRes, error)
	RolesCount(context.Context, *RolesCountReq) (*RolesCountRes, error)
	mustEmbedUnimplementedAuthGrpcServiceServer()
}

// UnimplementedAuthGrpcServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAuthGrpcServiceServer struct {
}

func (UnimplementedAuthGrpcServiceServer) AccessTokenSearch(context.Context, *AccessTokenSearchReq) (*AccessTokenSearchRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AccessTokenSearch not implemented")
}
func (UnimplementedAuthGrpcServiceServer) RolesCount(context.Context, *RolesCountReq) (*RolesCountRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RolesCount not implemented")
}
func (UnimplementedAuthGrpcServiceServer) mustEmbedUnimplementedAuthGrpcServiceServer() {}

// UnsafeAuthGrpcServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AuthGrpcServiceServer will
// result in compilation errors.
type UnsafeAuthGrpcServiceServer interface {
	mustEmbedUnimplementedAuthGrpcServiceServer()
}

func RegisterAuthGrpcServiceServer(s grpc.ServiceRegistrar, srv AuthGrpcServiceServer) {
	s.RegisterService(&AuthGrpcService_ServiceDesc, srv)
}

func _AuthGrpcService_AccessTokenSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AccessTokenSearchReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthGrpcServiceServer).AccessTokenSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthGrpcService_AccessTokenSearch_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthGrpcServiceServer).AccessTokenSearch(ctx, req.(*AccessTokenSearchReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _AuthGrpcService_RolesCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RolesCountReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthGrpcServiceServer).RolesCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AuthGrpcService_RolesCount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthGrpcServiceServer).RolesCount(ctx, req.(*RolesCountReq))
	}
	return interceptor(ctx, in, info, handler)
}

// AuthGrpcService_ServiceDesc is the grpc.ServiceDesc for AuthGrpcService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AuthGrpcService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "AuthGrpcService",
	HandlerType: (*AuthGrpcServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AccessTokenSearch",
			Handler:    _AuthGrpcService_AccessTokenSearch_Handler,
		},
		{
			MethodName: "RolesCount",
			Handler:    _AuthGrpcService_RolesCount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "modules/auth/authPb/authPb.proto",
}
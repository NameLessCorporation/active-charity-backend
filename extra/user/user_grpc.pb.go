// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: user/user.proto

package user

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
	User_CreateUserV1_FullMethodName             = "/proto.user.User/CreateUserV1"
	User_JoinUserToOrganizationV1_FullMethodName = "/proto.user.User/JoinUserToOrganizationV1"
)

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	CreateUserV1(ctx context.Context, in *CreateUserV1Request, opts ...grpc.CallOption) (*CreateUserV1Response, error)
	JoinUserToOrganizationV1(ctx context.Context, in *JoinUserToOrganizationV1Request, opts ...grpc.CallOption) (*JoinUserToOrganizationV1Response, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) CreateUserV1(ctx context.Context, in *CreateUserV1Request, opts ...grpc.CallOption) (*CreateUserV1Response, error) {
	out := new(CreateUserV1Response)
	err := c.cc.Invoke(ctx, User_CreateUserV1_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) JoinUserToOrganizationV1(ctx context.Context, in *JoinUserToOrganizationV1Request, opts ...grpc.CallOption) (*JoinUserToOrganizationV1Response, error) {
	out := new(JoinUserToOrganizationV1Response)
	err := c.cc.Invoke(ctx, User_JoinUserToOrganizationV1_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations should embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	CreateUserV1(context.Context, *CreateUserV1Request) (*CreateUserV1Response, error)
	JoinUserToOrganizationV1(context.Context, *JoinUserToOrganizationV1Request) (*JoinUserToOrganizationV1Response, error)
}

// UnimplementedUserServer should be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) CreateUserV1(context.Context, *CreateUserV1Request) (*CreateUserV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUserV1 not implemented")
}
func (UnimplementedUserServer) JoinUserToOrganizationV1(context.Context, *JoinUserToOrganizationV1Request) (*JoinUserToOrganizationV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinUserToOrganizationV1 not implemented")
}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_CreateUserV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateUserV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_CreateUserV1_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateUserV1(ctx, req.(*CreateUserV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_JoinUserToOrganizationV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinUserToOrganizationV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).JoinUserToOrganizationV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_JoinUserToOrganizationV1_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).JoinUserToOrganizationV1(ctx, req.(*JoinUserToOrganizationV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUserV1",
			Handler:    _User_CreateUserV1_Handler,
		},
		{
			MethodName: "JoinUserToOrganizationV1",
			Handler:    _User_JoinUserToOrganizationV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/user.proto",
}

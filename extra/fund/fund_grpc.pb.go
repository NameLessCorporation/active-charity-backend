// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.3
// source: fund/fund.proto

package fund

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
	Fund_CreateFundV1_FullMethodName = "/proto.fund.Fund/CreateFundV1"
	Fund_GetFundV1_FullMethodName    = "/proto.fund.Fund/GetFundV1"
	Fund_GetFundsV1_FullMethodName   = "/proto.fund.Fund/GetFundsV1"
)

// FundClient is the client API for Fund service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FundClient interface {
	CreateFundV1(ctx context.Context, in *CreateFundV1Request, opts ...grpc.CallOption) (*CreateFundV1Response, error)
	GetFundV1(ctx context.Context, in *GetFundV1Request, opts ...grpc.CallOption) (*GetFundV1Response, error)
	GetFundsV1(ctx context.Context, in *GetFundsV1Request, opts ...grpc.CallOption) (*GetFundsV1Response, error)
}

type fundClient struct {
	cc grpc.ClientConnInterface
}

func NewFundClient(cc grpc.ClientConnInterface) FundClient {
	return &fundClient{cc}
}

func (c *fundClient) CreateFundV1(ctx context.Context, in *CreateFundV1Request, opts ...grpc.CallOption) (*CreateFundV1Response, error) {
	out := new(CreateFundV1Response)
	err := c.cc.Invoke(ctx, Fund_CreateFundV1_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fundClient) GetFundV1(ctx context.Context, in *GetFundV1Request, opts ...grpc.CallOption) (*GetFundV1Response, error) {
	out := new(GetFundV1Response)
	err := c.cc.Invoke(ctx, Fund_GetFundV1_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *fundClient) GetFundsV1(ctx context.Context, in *GetFundsV1Request, opts ...grpc.CallOption) (*GetFundsV1Response, error) {
	out := new(GetFundsV1Response)
	err := c.cc.Invoke(ctx, Fund_GetFundsV1_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FundServer is the server API for Fund service.
// All implementations should embed UnimplementedFundServer
// for forward compatibility
type FundServer interface {
	CreateFundV1(context.Context, *CreateFundV1Request) (*CreateFundV1Response, error)
	GetFundV1(context.Context, *GetFundV1Request) (*GetFundV1Response, error)
	GetFundsV1(context.Context, *GetFundsV1Request) (*GetFundsV1Response, error)
}

// UnimplementedFundServer should be embedded to have forward compatible implementations.
type UnimplementedFundServer struct {
}

func (UnimplementedFundServer) CreateFundV1(context.Context, *CreateFundV1Request) (*CreateFundV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFundV1 not implemented")
}
func (UnimplementedFundServer) GetFundV1(context.Context, *GetFundV1Request) (*GetFundV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFundV1 not implemented")
}
func (UnimplementedFundServer) GetFundsV1(context.Context, *GetFundsV1Request) (*GetFundsV1Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFundsV1 not implemented")
}

// UnsafeFundServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FundServer will
// result in compilation errors.
type UnsafeFundServer interface {
	mustEmbedUnimplementedFundServer()
}

func RegisterFundServer(s grpc.ServiceRegistrar, srv FundServer) {
	s.RegisterService(&Fund_ServiceDesc, srv)
}

func _Fund_CreateFundV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateFundV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FundServer).CreateFundV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Fund_CreateFundV1_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FundServer).CreateFundV1(ctx, req.(*CreateFundV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fund_GetFundV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFundV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FundServer).GetFundV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Fund_GetFundV1_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FundServer).GetFundV1(ctx, req.(*GetFundV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

func _Fund_GetFundsV1_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFundsV1Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FundServer).GetFundsV1(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Fund_GetFundsV1_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FundServer).GetFundsV1(ctx, req.(*GetFundsV1Request))
	}
	return interceptor(ctx, in, info, handler)
}

// Fund_ServiceDesc is the grpc.ServiceDesc for Fund service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Fund_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.fund.Fund",
	HandlerType: (*FundServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFundV1",
			Handler:    _Fund_CreateFundV1_Handler,
		},
		{
			MethodName: "GetFundV1",
			Handler:    _Fund_GetFundV1_Handler,
		},
		{
			MethodName: "GetFundsV1",
			Handler:    _Fund_GetFundsV1_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "fund/fund.proto",
}
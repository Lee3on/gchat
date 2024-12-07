// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.3
// source: pkg/protocol/proto/user.int.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserIntClient is the client API for UserInt service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserIntClient interface {
	Auth(ctx context.Context, in *AuthReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserResp, error)
	GetUsers(ctx context.Context, in *GetUsersReq, opts ...grpc.CallOption) (*GetUsersResp, error)
}

type userIntClient struct {
	cc grpc.ClientConnInterface
}

func NewUserIntClient(cc grpc.ClientConnInterface) UserIntClient {
	return &userIntClient{cc}
}

func (c *userIntClient) Auth(ctx context.Context, in *AuthReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.UserInt/Auth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userIntClient) GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserResp, error) {
	out := new(GetUserResp)
	err := c.cc.Invoke(ctx, "/pb.UserInt/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userIntClient) GetUsers(ctx context.Context, in *GetUsersReq, opts ...grpc.CallOption) (*GetUsersResp, error) {
	out := new(GetUsersResp)
	err := c.cc.Invoke(ctx, "/pb.UserInt/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserIntServer is the server API for UserInt service.
// All implementations must embed UnimplementedUserIntServer
// for forward compatibility
type UserIntServer interface {
	Auth(context.Context, *AuthReq) (*emptypb.Empty, error)
	GetUser(context.Context, *GetUserReq) (*GetUserResp, error)
	GetUsers(context.Context, *GetUsersReq) (*GetUsersResp, error)
	mustEmbedUnimplementedUserIntServer()
}

// UnimplementedUserIntServer must be embedded to have forward compatible implementations.
type UnimplementedUserIntServer struct {
}

func (UnimplementedUserIntServer) Auth(context.Context, *AuthReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Auth not implemented")
}
func (UnimplementedUserIntServer) GetUser(context.Context, *GetUserReq) (*GetUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserIntServer) GetUsers(context.Context, *GetUsersReq) (*GetUsersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedUserIntServer) mustEmbedUnimplementedUserIntServer() {}

// UnsafeUserIntServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserIntServer will
// result in compilation errors.
type UnsafeUserIntServer interface {
	mustEmbedUnimplementedUserIntServer()
}

func RegisterUserIntServer(s grpc.ServiceRegistrar, srv UserIntServer) {
	s.RegisterService(&UserInt_ServiceDesc, srv)
}

func _UserInt_Auth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserIntServer).Auth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserInt/Auth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserIntServer).Auth(ctx, req.(*AuthReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserInt_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserIntServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserInt/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserIntServer).GetUser(ctx, req.(*GetUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserInt_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserIntServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserInt/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserIntServer).GetUsers(ctx, req.(*GetUsersReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserInt_ServiceDesc is the grpc.ServiceDesc for UserInt service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserInt_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserInt",
	HandlerType: (*UserIntServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Auth",
			Handler:    _UserInt_Auth_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserInt_GetUser_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _UserInt_GetUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/protocol/proto/user.int.proto",
}

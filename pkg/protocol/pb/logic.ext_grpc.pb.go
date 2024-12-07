// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.3
// source: pkg/protocol/proto/logic.ext.proto

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

// LogicExtClient is the client API for LogicExt service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LogicExtClient interface {
	// Push a message to the room
	PushRoom(ctx context.Context, in *PushRoomReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Send a message to a friend
	SendMessageToFriend(ctx context.Context, in *SendMessageReq, opts ...grpc.CallOption) (*SendMessageResp, error)
	// Send a group message
	SendMessageToGroup(ctx context.Context, in *SendMessageReq, opts ...grpc.CallOption) (*SendMessageResp, error)
}

type logicExtClient struct {
	cc grpc.ClientConnInterface
}

func NewLogicExtClient(cc grpc.ClientConnInterface) LogicExtClient {
	return &logicExtClient{cc}
}

func (c *logicExtClient) PushRoom(ctx context.Context, in *PushRoomReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.LogicExt/PushRoom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicExtClient) SendMessageToFriend(ctx context.Context, in *SendMessageReq, opts ...grpc.CallOption) (*SendMessageResp, error) {
	out := new(SendMessageResp)
	err := c.cc.Invoke(ctx, "/pb.LogicExt/SendMessageToFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicExtClient) SendMessageToGroup(ctx context.Context, in *SendMessageReq, opts ...grpc.CallOption) (*SendMessageResp, error) {
	out := new(SendMessageResp)
	err := c.cc.Invoke(ctx, "/pb.LogicExt/SendMessageToGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogicExtServer is the server API for LogicExt service.
// All implementations must embed UnimplementedLogicExtServer
// for forward compatibility
type LogicExtServer interface {
	// Push a message to the room
	PushRoom(context.Context, *PushRoomReq) (*emptypb.Empty, error)
	// Send a message to a friend
	SendMessageToFriend(context.Context, *SendMessageReq) (*SendMessageResp, error)
	// Send a group message
	SendMessageToGroup(context.Context, *SendMessageReq) (*SendMessageResp, error)
	mustEmbedUnimplementedLogicExtServer()
}

// UnimplementedLogicExtServer must be embedded to have forward compatible implementations.
type UnimplementedLogicExtServer struct {
}

func (UnimplementedLogicExtServer) PushRoom(context.Context, *PushRoomReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushRoom not implemented")
}
func (UnimplementedLogicExtServer) SendMessageToFriend(context.Context, *SendMessageReq) (*SendMessageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessageToFriend not implemented")
}
func (UnimplementedLogicExtServer) SendMessageToGroup(context.Context, *SendMessageReq) (*SendMessageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessageToGroup not implemented")
}
func (UnimplementedLogicExtServer) mustEmbedUnimplementedLogicExtServer() {}

// UnsafeLogicExtServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LogicExtServer will
// result in compilation errors.
type UnsafeLogicExtServer interface {
	mustEmbedUnimplementedLogicExtServer()
}

func RegisterLogicExtServer(s grpc.ServiceRegistrar, srv LogicExtServer) {
	s.RegisterService(&LogicExt_ServiceDesc, srv)
}

func _LogicExt_PushRoom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PushRoomReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicExtServer).PushRoom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LogicExt/PushRoom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicExtServer).PushRoom(ctx, req.(*PushRoomReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicExt_SendMessageToFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicExtServer).SendMessageToFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LogicExt/SendMessageToFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicExtServer).SendMessageToFriend(ctx, req.(*SendMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicExt_SendMessageToGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendMessageReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicExtServer).SendMessageToGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LogicExt/SendMessageToGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicExtServer).SendMessageToGroup(ctx, req.(*SendMessageReq))
	}
	return interceptor(ctx, in, info, handler)
}

// LogicExt_ServiceDesc is the grpc.ServiceDesc for LogicExt service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LogicExt_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.LogicExt",
	HandlerType: (*LogicExtServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PushRoom",
			Handler:    _LogicExt_PushRoom_Handler,
		},
		{
			MethodName: "SendMessageToFriend",
			Handler:    _LogicExt_SendMessageToFriend_Handler,
		},
		{
			MethodName: "SendMessageToGroup",
			Handler:    _LogicExt_SendMessageToGroup_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/protocol/proto/logic.ext.proto",
}

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
	// Register device
	RegisterDevice(ctx context.Context, in *RegisterDeviceReq, opts ...grpc.CallOption) (*RegisterDeviceResp, error)
	// Push a message to the room
	PushRoom(ctx context.Context, in *PushRoomReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Send a message to a friend
	SendMessageToFriend(ctx context.Context, in *SendMessageReq, opts ...grpc.CallOption) (*SendMessageResp, error)
	// Add a friend
	AddFriend(ctx context.Context, in *AddFriendReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Agree to add a friend
	AgreeAddFriend(ctx context.Context, in *AgreeAddFriendReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Set friend information
	SetFriend(ctx context.Context, in *SetFriendReq, opts ...grpc.CallOption) (*SetFriendResp, error)
	// Get the friend list
	GetFriends(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetFriendsResp, error)
	// Send a group message
	SendMessageToGroup(ctx context.Context, in *SendMessageReq, opts ...grpc.CallOption) (*SendMessageResp, error)
	// Create a group
	CreateGroup(ctx context.Context, in *CreateGroupReq, opts ...grpc.CallOption) (*CreateGroupResp, error)
	// Update group information
	UpdateGroup(ctx context.Context, in *UpdateGroupReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Get group information
	GetGroup(ctx context.Context, in *GetGroupReq, opts ...grpc.CallOption) (*GetGroupResp, error)
	// Get all groups the user has joined
	GetGroups(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetGroupsResp, error)
	// Add group members
	AddGroupMembers(ctx context.Context, in *AddGroupMembersReq, opts ...grpc.CallOption) (*AddGroupMembersResp, error)
	// Update group member information
	UpdateGroupMember(ctx context.Context, in *UpdateGroupMemberReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Delete group members
	DeleteGroupMember(ctx context.Context, in *DeleteGroupMemberReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Get group members
	GetGroupMembers(ctx context.Context, in *GetGroupMembersReq, opts ...grpc.CallOption) (*GetGroupMembersResp, error)
}

type logicExtClient struct {
	cc grpc.ClientConnInterface
}

func NewLogicExtClient(cc grpc.ClientConnInterface) LogicExtClient {
	return &logicExtClient{cc}
}

func (c *logicExtClient) RegisterDevice(ctx context.Context, in *RegisterDeviceReq, opts ...grpc.CallOption) (*RegisterDeviceResp, error) {
	out := new(RegisterDeviceResp)
	err := c.cc.Invoke(ctx, "/pb.LogicExt/RegisterDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
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

func (c *logicExtClient) AddFriend(ctx context.Context, in *AddFriendReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.LogicExt/AddFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicExtClient) AgreeAddFriend(ctx context.Context, in *AgreeAddFriendReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.LogicExt/AgreeAddFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicExtClient) SetFriend(ctx context.Context, in *SetFriendReq, opts ...grpc.CallOption) (*SetFriendResp, error) {
	out := new(SetFriendResp)
	err := c.cc.Invoke(ctx, "/pb.LogicExt/SetFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicExtClient) GetFriends(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetFriendsResp, error) {
	out := new(GetFriendsResp)
	err := c.cc.Invoke(ctx, "/pb.LogicExt/GetFriends", in, out, opts...)
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

func (c *logicExtClient) CreateGroup(ctx context.Context, in *CreateGroupReq, opts ...grpc.CallOption) (*CreateGroupResp, error) {
	out := new(CreateGroupResp)
	err := c.cc.Invoke(ctx, "/pb.LogicExt/CreateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicExtClient) UpdateGroup(ctx context.Context, in *UpdateGroupReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.LogicExt/UpdateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicExtClient) GetGroup(ctx context.Context, in *GetGroupReq, opts ...grpc.CallOption) (*GetGroupResp, error) {
	out := new(GetGroupResp)
	err := c.cc.Invoke(ctx, "/pb.LogicExt/GetGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicExtClient) GetGroups(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetGroupsResp, error) {
	out := new(GetGroupsResp)
	err := c.cc.Invoke(ctx, "/pb.LogicExt/GetGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicExtClient) AddGroupMembers(ctx context.Context, in *AddGroupMembersReq, opts ...grpc.CallOption) (*AddGroupMembersResp, error) {
	out := new(AddGroupMembersResp)
	err := c.cc.Invoke(ctx, "/pb.LogicExt/AddGroupMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicExtClient) UpdateGroupMember(ctx context.Context, in *UpdateGroupMemberReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.LogicExt/UpdateGroupMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicExtClient) DeleteGroupMember(ctx context.Context, in *DeleteGroupMemberReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.LogicExt/DeleteGroupMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *logicExtClient) GetGroupMembers(ctx context.Context, in *GetGroupMembersReq, opts ...grpc.CallOption) (*GetGroupMembersResp, error) {
	out := new(GetGroupMembersResp)
	err := c.cc.Invoke(ctx, "/pb.LogicExt/GetGroupMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LogicExtServer is the server API for LogicExt service.
// All implementations must embed UnimplementedLogicExtServer
// for forward compatibility
type LogicExtServer interface {
	// Register device
	RegisterDevice(context.Context, *RegisterDeviceReq) (*RegisterDeviceResp, error)
	// Push a message to the room
	PushRoom(context.Context, *PushRoomReq) (*emptypb.Empty, error)
	// Send a message to a friend
	SendMessageToFriend(context.Context, *SendMessageReq) (*SendMessageResp, error)
	// Add a friend
	AddFriend(context.Context, *AddFriendReq) (*emptypb.Empty, error)
	// Agree to add a friend
	AgreeAddFriend(context.Context, *AgreeAddFriendReq) (*emptypb.Empty, error)
	// Set friend information
	SetFriend(context.Context, *SetFriendReq) (*SetFriendResp, error)
	// Get the friend list
	GetFriends(context.Context, *emptypb.Empty) (*GetFriendsResp, error)
	// Send a group message
	SendMessageToGroup(context.Context, *SendMessageReq) (*SendMessageResp, error)
	// Create a group
	CreateGroup(context.Context, *CreateGroupReq) (*CreateGroupResp, error)
	// Update group information
	UpdateGroup(context.Context, *UpdateGroupReq) (*emptypb.Empty, error)
	// Get group information
	GetGroup(context.Context, *GetGroupReq) (*GetGroupResp, error)
	// Get all groups the user has joined
	GetGroups(context.Context, *emptypb.Empty) (*GetGroupsResp, error)
	// Add group members
	AddGroupMembers(context.Context, *AddGroupMembersReq) (*AddGroupMembersResp, error)
	// Update group member information
	UpdateGroupMember(context.Context, *UpdateGroupMemberReq) (*emptypb.Empty, error)
	// Delete group members
	DeleteGroupMember(context.Context, *DeleteGroupMemberReq) (*emptypb.Empty, error)
	// Get group members
	GetGroupMembers(context.Context, *GetGroupMembersReq) (*GetGroupMembersResp, error)
	mustEmbedUnimplementedLogicExtServer()
}

// UnimplementedLogicExtServer must be embedded to have forward compatible implementations.
type UnimplementedLogicExtServer struct {
}

func (UnimplementedLogicExtServer) RegisterDevice(context.Context, *RegisterDeviceReq) (*RegisterDeviceResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterDevice not implemented")
}
func (UnimplementedLogicExtServer) PushRoom(context.Context, *PushRoomReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PushRoom not implemented")
}
func (UnimplementedLogicExtServer) SendMessageToFriend(context.Context, *SendMessageReq) (*SendMessageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessageToFriend not implemented")
}
func (UnimplementedLogicExtServer) AddFriend(context.Context, *AddFriendReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFriend not implemented")
}
func (UnimplementedLogicExtServer) AgreeAddFriend(context.Context, *AgreeAddFriendReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AgreeAddFriend not implemented")
}
func (UnimplementedLogicExtServer) SetFriend(context.Context, *SetFriendReq) (*SetFriendResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetFriend not implemented")
}
func (UnimplementedLogicExtServer) GetFriends(context.Context, *emptypb.Empty) (*GetFriendsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriends not implemented")
}
func (UnimplementedLogicExtServer) SendMessageToGroup(context.Context, *SendMessageReq) (*SendMessageResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendMessageToGroup not implemented")
}
func (UnimplementedLogicExtServer) CreateGroup(context.Context, *CreateGroupReq) (*CreateGroupResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (UnimplementedLogicExtServer) UpdateGroup(context.Context, *UpdateGroupReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroup not implemented")
}
func (UnimplementedLogicExtServer) GetGroup(context.Context, *GetGroupReq) (*GetGroupResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroup not implemented")
}
func (UnimplementedLogicExtServer) GetGroups(context.Context, *emptypb.Empty) (*GetGroupsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroups not implemented")
}
func (UnimplementedLogicExtServer) AddGroupMembers(context.Context, *AddGroupMembersReq) (*AddGroupMembersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGroupMembers not implemented")
}
func (UnimplementedLogicExtServer) UpdateGroupMember(context.Context, *UpdateGroupMemberReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroupMember not implemented")
}
func (UnimplementedLogicExtServer) DeleteGroupMember(context.Context, *DeleteGroupMemberReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroupMember not implemented")
}
func (UnimplementedLogicExtServer) GetGroupMembers(context.Context, *GetGroupMembersReq) (*GetGroupMembersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupMembers not implemented")
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

func _LogicExt_RegisterDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterDeviceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicExtServer).RegisterDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LogicExt/RegisterDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicExtServer).RegisterDevice(ctx, req.(*RegisterDeviceReq))
	}
	return interceptor(ctx, in, info, handler)
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

func _LogicExt_AddFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFriendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicExtServer).AddFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LogicExt/AddFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicExtServer).AddFriend(ctx, req.(*AddFriendReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicExt_AgreeAddFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgreeAddFriendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicExtServer).AgreeAddFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LogicExt/AgreeAddFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicExtServer).AgreeAddFriend(ctx, req.(*AgreeAddFriendReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicExt_SetFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetFriendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicExtServer).SetFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LogicExt/SetFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicExtServer).SetFriend(ctx, req.(*SetFriendReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicExt_GetFriends_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicExtServer).GetFriends(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LogicExt/GetFriends",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicExtServer).GetFriends(ctx, req.(*emptypb.Empty))
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

func _LogicExt_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicExtServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LogicExt/CreateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicExtServer).CreateGroup(ctx, req.(*CreateGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicExt_UpdateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicExtServer).UpdateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LogicExt/UpdateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicExtServer).UpdateGroup(ctx, req.(*UpdateGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicExt_GetGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicExtServer).GetGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LogicExt/GetGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicExtServer).GetGroup(ctx, req.(*GetGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicExt_GetGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicExtServer).GetGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LogicExt/GetGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicExtServer).GetGroups(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicExt_AddGroupMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddGroupMembersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicExtServer).AddGroupMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LogicExt/AddGroupMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicExtServer).AddGroupMembers(ctx, req.(*AddGroupMembersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicExt_UpdateGroupMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupMemberReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicExtServer).UpdateGroupMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LogicExt/UpdateGroupMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicExtServer).UpdateGroupMember(ctx, req.(*UpdateGroupMemberReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicExt_DeleteGroupMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroupMemberReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicExtServer).DeleteGroupMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LogicExt/DeleteGroupMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicExtServer).DeleteGroupMember(ctx, req.(*DeleteGroupMemberReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _LogicExt_GetGroupMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupMembersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LogicExtServer).GetGroupMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.LogicExt/GetGroupMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LogicExtServer).GetGroupMembers(ctx, req.(*GetGroupMembersReq))
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
			MethodName: "RegisterDevice",
			Handler:    _LogicExt_RegisterDevice_Handler,
		},
		{
			MethodName: "PushRoom",
			Handler:    _LogicExt_PushRoom_Handler,
		},
		{
			MethodName: "SendMessageToFriend",
			Handler:    _LogicExt_SendMessageToFriend_Handler,
		},
		{
			MethodName: "AddFriend",
			Handler:    _LogicExt_AddFriend_Handler,
		},
		{
			MethodName: "AgreeAddFriend",
			Handler:    _LogicExt_AgreeAddFriend_Handler,
		},
		{
			MethodName: "SetFriend",
			Handler:    _LogicExt_SetFriend_Handler,
		},
		{
			MethodName: "GetFriends",
			Handler:    _LogicExt_GetFriends_Handler,
		},
		{
			MethodName: "SendMessageToGroup",
			Handler:    _LogicExt_SendMessageToGroup_Handler,
		},
		{
			MethodName: "CreateGroup",
			Handler:    _LogicExt_CreateGroup_Handler,
		},
		{
			MethodName: "UpdateGroup",
			Handler:    _LogicExt_UpdateGroup_Handler,
		},
		{
			MethodName: "GetGroup",
			Handler:    _LogicExt_GetGroup_Handler,
		},
		{
			MethodName: "GetGroups",
			Handler:    _LogicExt_GetGroups_Handler,
		},
		{
			MethodName: "AddGroupMembers",
			Handler:    _LogicExt_AddGroupMembers_Handler,
		},
		{
			MethodName: "UpdateGroupMember",
			Handler:    _LogicExt_UpdateGroupMember_Handler,
		},
		{
			MethodName: "DeleteGroupMember",
			Handler:    _LogicExt_DeleteGroupMember_Handler,
		},
		{
			MethodName: "GetGroupMembers",
			Handler:    _LogicExt_GetGroupMembers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/protocol/proto/logic.ext.proto",
}
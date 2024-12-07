// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.3
// source: pkg/protocol/proto/user.ext.proto

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

// UserExtClient is the client API for UserExt service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserExtClient interface {
	// Register device
	RegisterDevice(ctx context.Context, in *RegisterDeviceReq, opts ...grpc.CallOption) (*RegisterDeviceResp, error)
	SignIn(ctx context.Context, in *SignInReq, opts ...grpc.CallOption) (*SignInResp, error)
	GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserResp, error)
	UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Add a friend
	AddFriend(ctx context.Context, in *AddFriendReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Agree to add a friend
	AgreeAddFriend(ctx context.Context, in *AgreeAddFriendReq, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// Set friend information
	SetFriend(ctx context.Context, in *SetFriendReq, opts ...grpc.CallOption) (*SetFriendResp, error)
	// Get the friend list
	GetFriends(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetFriendsResp, error)
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

type userExtClient struct {
	cc grpc.ClientConnInterface
}

func NewUserExtClient(cc grpc.ClientConnInterface) UserExtClient {
	return &userExtClient{cc}
}

func (c *userExtClient) RegisterDevice(ctx context.Context, in *RegisterDeviceReq, opts ...grpc.CallOption) (*RegisterDeviceResp, error) {
	out := new(RegisterDeviceResp)
	err := c.cc.Invoke(ctx, "/pb.UserExt/RegisterDevice", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userExtClient) SignIn(ctx context.Context, in *SignInReq, opts ...grpc.CallOption) (*SignInResp, error) {
	out := new(SignInResp)
	err := c.cc.Invoke(ctx, "/pb.UserExt/SignIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userExtClient) GetUser(ctx context.Context, in *GetUserReq, opts ...grpc.CallOption) (*GetUserResp, error) {
	out := new(GetUserResp)
	err := c.cc.Invoke(ctx, "/pb.UserExt/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userExtClient) UpdateUser(ctx context.Context, in *UpdateUserReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.UserExt/UpdateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userExtClient) AddFriend(ctx context.Context, in *AddFriendReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.UserExt/AddFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userExtClient) AgreeAddFriend(ctx context.Context, in *AgreeAddFriendReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.UserExt/AgreeAddFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userExtClient) SetFriend(ctx context.Context, in *SetFriendReq, opts ...grpc.CallOption) (*SetFriendResp, error) {
	out := new(SetFriendResp)
	err := c.cc.Invoke(ctx, "/pb.UserExt/SetFriend", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userExtClient) GetFriends(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetFriendsResp, error) {
	out := new(GetFriendsResp)
	err := c.cc.Invoke(ctx, "/pb.UserExt/GetFriends", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userExtClient) CreateGroup(ctx context.Context, in *CreateGroupReq, opts ...grpc.CallOption) (*CreateGroupResp, error) {
	out := new(CreateGroupResp)
	err := c.cc.Invoke(ctx, "/pb.UserExt/CreateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userExtClient) UpdateGroup(ctx context.Context, in *UpdateGroupReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.UserExt/UpdateGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userExtClient) GetGroup(ctx context.Context, in *GetGroupReq, opts ...grpc.CallOption) (*GetGroupResp, error) {
	out := new(GetGroupResp)
	err := c.cc.Invoke(ctx, "/pb.UserExt/GetGroup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userExtClient) GetGroups(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*GetGroupsResp, error) {
	out := new(GetGroupsResp)
	err := c.cc.Invoke(ctx, "/pb.UserExt/GetGroups", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userExtClient) AddGroupMembers(ctx context.Context, in *AddGroupMembersReq, opts ...grpc.CallOption) (*AddGroupMembersResp, error) {
	out := new(AddGroupMembersResp)
	err := c.cc.Invoke(ctx, "/pb.UserExt/AddGroupMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userExtClient) UpdateGroupMember(ctx context.Context, in *UpdateGroupMemberReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.UserExt/UpdateGroupMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userExtClient) DeleteGroupMember(ctx context.Context, in *DeleteGroupMemberReq, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/pb.UserExt/DeleteGroupMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userExtClient) GetGroupMembers(ctx context.Context, in *GetGroupMembersReq, opts ...grpc.CallOption) (*GetGroupMembersResp, error) {
	out := new(GetGroupMembersResp)
	err := c.cc.Invoke(ctx, "/pb.UserExt/GetGroupMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserExtServer is the server API for UserExt service.
// All implementations must embed UnimplementedUserExtServer
// for forward compatibility
type UserExtServer interface {
	// Register device
	RegisterDevice(context.Context, *RegisterDeviceReq) (*RegisterDeviceResp, error)
	SignIn(context.Context, *SignInReq) (*SignInResp, error)
	GetUser(context.Context, *GetUserReq) (*GetUserResp, error)
	UpdateUser(context.Context, *UpdateUserReq) (*emptypb.Empty, error)
	// Add a friend
	AddFriend(context.Context, *AddFriendReq) (*emptypb.Empty, error)
	// Agree to add a friend
	AgreeAddFriend(context.Context, *AgreeAddFriendReq) (*emptypb.Empty, error)
	// Set friend information
	SetFriend(context.Context, *SetFriendReq) (*SetFriendResp, error)
	// Get the friend list
	GetFriends(context.Context, *emptypb.Empty) (*GetFriendsResp, error)
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
	mustEmbedUnimplementedUserExtServer()
}

// UnimplementedUserExtServer must be embedded to have forward compatible implementations.
type UnimplementedUserExtServer struct {
}

func (UnimplementedUserExtServer) RegisterDevice(context.Context, *RegisterDeviceReq) (*RegisterDeviceResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterDevice not implemented")
}
func (UnimplementedUserExtServer) SignIn(context.Context, *SignInReq) (*SignInResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SignIn not implemented")
}
func (UnimplementedUserExtServer) GetUser(context.Context, *GetUserReq) (*GetUserResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserExtServer) UpdateUser(context.Context, *UpdateUserReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUser not implemented")
}
func (UnimplementedUserExtServer) AddFriend(context.Context, *AddFriendReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddFriend not implemented")
}
func (UnimplementedUserExtServer) AgreeAddFriend(context.Context, *AgreeAddFriendReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AgreeAddFriend not implemented")
}
func (UnimplementedUserExtServer) SetFriend(context.Context, *SetFriendReq) (*SetFriendResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetFriend not implemented")
}
func (UnimplementedUserExtServer) GetFriends(context.Context, *emptypb.Empty) (*GetFriendsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFriends not implemented")
}
func (UnimplementedUserExtServer) CreateGroup(context.Context, *CreateGroupReq) (*CreateGroupResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGroup not implemented")
}
func (UnimplementedUserExtServer) UpdateGroup(context.Context, *UpdateGroupReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroup not implemented")
}
func (UnimplementedUserExtServer) GetGroup(context.Context, *GetGroupReq) (*GetGroupResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroup not implemented")
}
func (UnimplementedUserExtServer) GetGroups(context.Context, *emptypb.Empty) (*GetGroupsResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroups not implemented")
}
func (UnimplementedUserExtServer) AddGroupMembers(context.Context, *AddGroupMembersReq) (*AddGroupMembersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddGroupMembers not implemented")
}
func (UnimplementedUserExtServer) UpdateGroupMember(context.Context, *UpdateGroupMemberReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateGroupMember not implemented")
}
func (UnimplementedUserExtServer) DeleteGroupMember(context.Context, *DeleteGroupMemberReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteGroupMember not implemented")
}
func (UnimplementedUserExtServer) GetGroupMembers(context.Context, *GetGroupMembersReq) (*GetGroupMembersResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupMembers not implemented")
}
func (UnimplementedUserExtServer) mustEmbedUnimplementedUserExtServer() {}

// UnsafeUserExtServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserExtServer will
// result in compilation errors.
type UnsafeUserExtServer interface {
	mustEmbedUnimplementedUserExtServer()
}

func RegisterUserExtServer(s grpc.ServiceRegistrar, srv UserExtServer) {
	s.RegisterService(&UserExt_ServiceDesc, srv)
}

func _UserExt_RegisterDevice_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterDeviceReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserExtServer).RegisterDevice(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserExt/RegisterDevice",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserExtServer).RegisterDevice(ctx, req.(*RegisterDeviceReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserExt_SignIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignInReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserExtServer).SignIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserExt/SignIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserExtServer).SignIn(ctx, req.(*SignInReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserExt_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserExtServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserExt/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserExtServer).GetUser(ctx, req.(*GetUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserExt_UpdateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUserReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserExtServer).UpdateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserExt/UpdateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserExtServer).UpdateUser(ctx, req.(*UpdateUserReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserExt_AddFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddFriendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserExtServer).AddFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserExt/AddFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserExtServer).AddFriend(ctx, req.(*AddFriendReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserExt_AgreeAddFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AgreeAddFriendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserExtServer).AgreeAddFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserExt/AgreeAddFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserExtServer).AgreeAddFriend(ctx, req.(*AgreeAddFriendReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserExt_SetFriend_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetFriendReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserExtServer).SetFriend(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserExt/SetFriend",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserExtServer).SetFriend(ctx, req.(*SetFriendReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserExt_GetFriends_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserExtServer).GetFriends(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserExt/GetFriends",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserExtServer).GetFriends(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserExt_CreateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserExtServer).CreateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserExt/CreateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserExtServer).CreateGroup(ctx, req.(*CreateGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserExt_UpdateGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserExtServer).UpdateGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserExt/UpdateGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserExtServer).UpdateGroup(ctx, req.(*UpdateGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserExt_GetGroup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserExtServer).GetGroup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserExt/GetGroup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserExtServer).GetGroup(ctx, req.(*GetGroupReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserExt_GetGroups_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserExtServer).GetGroups(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserExt/GetGroups",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserExtServer).GetGroups(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserExt_AddGroupMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddGroupMembersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserExtServer).AddGroupMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserExt/AddGroupMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserExtServer).AddGroupMembers(ctx, req.(*AddGroupMembersReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserExt_UpdateGroupMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateGroupMemberReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserExtServer).UpdateGroupMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserExt/UpdateGroupMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserExtServer).UpdateGroupMember(ctx, req.(*UpdateGroupMemberReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserExt_DeleteGroupMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteGroupMemberReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserExtServer).DeleteGroupMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserExt/DeleteGroupMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserExtServer).DeleteGroupMember(ctx, req.(*DeleteGroupMemberReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserExt_GetGroupMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupMembersReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserExtServer).GetGroupMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.UserExt/GetGroupMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserExtServer).GetGroupMembers(ctx, req.(*GetGroupMembersReq))
	}
	return interceptor(ctx, in, info, handler)
}

// UserExt_ServiceDesc is the grpc.ServiceDesc for UserExt service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserExt_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.UserExt",
	HandlerType: (*UserExtServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterDevice",
			Handler:    _UserExt_RegisterDevice_Handler,
		},
		{
			MethodName: "SignIn",
			Handler:    _UserExt_SignIn_Handler,
		},
		{
			MethodName: "GetUser",
			Handler:    _UserExt_GetUser_Handler,
		},
		{
			MethodName: "UpdateUser",
			Handler:    _UserExt_UpdateUser_Handler,
		},
		{
			MethodName: "AddFriend",
			Handler:    _UserExt_AddFriend_Handler,
		},
		{
			MethodName: "AgreeAddFriend",
			Handler:    _UserExt_AgreeAddFriend_Handler,
		},
		{
			MethodName: "SetFriend",
			Handler:    _UserExt_SetFriend_Handler,
		},
		{
			MethodName: "GetFriends",
			Handler:    _UserExt_GetFriends_Handler,
		},
		{
			MethodName: "CreateGroup",
			Handler:    _UserExt_CreateGroup_Handler,
		},
		{
			MethodName: "UpdateGroup",
			Handler:    _UserExt_UpdateGroup_Handler,
		},
		{
			MethodName: "GetGroup",
			Handler:    _UserExt_GetGroup_Handler,
		},
		{
			MethodName: "GetGroups",
			Handler:    _UserExt_GetGroups_Handler,
		},
		{
			MethodName: "AddGroupMembers",
			Handler:    _UserExt_AddGroupMembers_Handler,
		},
		{
			MethodName: "UpdateGroupMember",
			Handler:    _UserExt_UpdateGroupMember_Handler,
		},
		{
			MethodName: "DeleteGroupMember",
			Handler:    _UserExt_DeleteGroupMember_Handler,
		},
		{
			MethodName: "GetGroupMembers",
			Handler:    _UserExt_GetGroupMembers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/protocol/proto/user.ext.proto",
}

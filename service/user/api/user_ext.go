package api

import (
	"context"
	"gchat/pkg/grpclib"
	"gchat/pkg/protocol/pb"
	app2 "gchat/service/user/domain/app"
	"gchat/service/user/domain/device"
	"gchat/service/user/domain/friend"
	"gchat/service/user/domain/group"

	"google.golang.org/protobuf/types/known/emptypb"
)

type UserExtServer struct {
	pb.UnsafeUserExtServer
}

// RegisterDevice register device
func (*UserExtServer) RegisterDevice(ctx context.Context, in *pb.RegisterDeviceReq) (*pb.RegisterDeviceResp, error) {
	deviceId, err := device.App.Register(ctx, in)
	return &pb.RegisterDeviceResp{DeviceId: deviceId}, err
}

func (s *UserExtServer) SignIn(ctx context.Context, req *pb.SignInReq) (*pb.SignInResp, error) {
	isNew, userId, token, err := app2.AuthApp.SignIn(ctx, req.PhoneNumber, req.Code, req.DeviceId)
	if err != nil {
		return nil, err
	}
	return &pb.SignInResp{
		IsNew:  isNew,
		UserId: userId,
		Token:  token,
	}, nil
}

func (s *UserExtServer) GetUser(ctx context.Context, req *pb.GetUserReq) (*pb.GetUserResp, error) {
	userId, _, err := grpclib.GetCtxData(ctx)
	if err != nil {
		return nil, err
	}

	user, err := app2.UserApp.Get(ctx, userId)
	return &pb.GetUserResp{User: user}, err
}

func (s *UserExtServer) UpdateUser(ctx context.Context, req *pb.UpdateUserReq) (*emptypb.Empty, error) {
	userId, _, err := grpclib.GetCtxData(ctx)
	if err != nil {
		return nil, err
	}

	return new(emptypb.Empty), app2.UserApp.Update(ctx, userId, req)
}

// AddFriend add a new friend
func (s *UserExtServer) AddFriend(ctx context.Context, in *pb.AddFriendReq) (*emptypb.Empty, error) {
	userId, _, err := grpclib.GetCtxData(ctx)
	if err != nil {
		return nil, err
	}

	err = friend.App.AddFriend(ctx, userId, in.FriendId, in.Remarks, in.Description)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// AgreeAddFriend agree an add friend request
func (s *UserExtServer) AgreeAddFriend(ctx context.Context, in *pb.AgreeAddFriendReq) (*emptypb.Empty, error) {
	userId, _, err := grpclib.GetCtxData(ctx)
	if err != nil {
		return nil, err
	}

	err = friend.App.AgreeAddFriend(ctx, userId, in.UserId, in.Remarks)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// SetFriend set friend information
func (s *UserExtServer) SetFriend(ctx context.Context, req *pb.SetFriendReq) (*pb.SetFriendResp, error) {
	userId, _, err := grpclib.GetCtxData(ctx)
	if err != nil {
		return nil, err
	}

	err = friend.App.SetFriend(ctx, userId, req)
	if err != nil {
		return nil, err
	}
	return &pb.SetFriendResp{}, nil
}

// GetFriends get all friends
func (s *UserExtServer) GetFriends(ctx context.Context, in *emptypb.Empty) (*pb.GetFriendsResp, error) {
	userId, _, err := grpclib.GetCtxData(ctx)
	if err != nil {
		return nil, err
	}
	friends, err := friend.App.List(ctx, userId)
	return &pb.GetFriendsResp{Friends: friends}, err
}

// CreateGroup create a new group
func (*UserExtServer) CreateGroup(ctx context.Context, in *pb.CreateGroupReq) (*pb.CreateGroupResp, error) {
	userId, _, err := grpclib.GetCtxData(ctx)
	if err != nil {
		return nil, err
	}

	groupId, err := group.App.CreateGroup(ctx, userId, in)
	return &pb.CreateGroupResp{GroupId: groupId}, err
}

// UpdateGroup update group information
func (*UserExtServer) UpdateGroup(ctx context.Context, in *pb.UpdateGroupReq) (*emptypb.Empty, error) {
	userId, _, err := grpclib.GetCtxData(ctx)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, group.App.Update(ctx, userId, in)
}

// GetGroup get group information
func (*UserExtServer) GetGroup(ctx context.Context, in *pb.GetGroupReq) (*pb.GetGroupResp, error) {
	group, err := group.App.GetGroup(ctx, in.GroupId)
	return &pb.GetGroupResp{Group: group}, err
}

// GetGroups get all groups a user belongs to
func (*UserExtServer) GetGroups(ctx context.Context, in *emptypb.Empty) (*pb.GetGroupsResp, error) {
	userId, _, err := grpclib.GetCtxData(ctx)
	if err != nil {
		return nil, err
	}

	groups, err := group.App.GetUserGroups(ctx, userId)
	return &pb.GetGroupsResp{Groups: groups}, err
}

// AddGroupMembers add members to a group
func (s *UserExtServer) AddGroupMembers(ctx context.Context, in *pb.AddGroupMembersReq) (*pb.AddGroupMembersResp, error) {
	userId, _, err := grpclib.GetCtxData(ctx)
	if err != nil {
		return nil, err
	}

	userIds, err := group.App.AddMembers(ctx, userId, in.GroupId, in.UserIds)
	return &pb.AddGroupMembersResp{UserIds: userIds}, err
}

// UpdateGroupMember update group member information
func (*UserExtServer) UpdateGroupMember(ctx context.Context, in *pb.UpdateGroupMemberReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, group.App.UpdateMember(ctx, in)
}

// DeleteGroupMember delete a group member
func (*UserExtServer) DeleteGroupMember(ctx context.Context, in *pb.DeleteGroupMemberReq) (*emptypb.Empty, error) {
	userId, _, err := grpclib.GetCtxData(ctx)
	if err != nil {
		return nil, err
	}

	err = group.App.DeleteMember(ctx, in.GroupId, in.UserId, userId)
	return &emptypb.Empty{}, err
}

// GetGroupMembers get all members of a group
func (s *UserExtServer) GetGroupMembers(ctx context.Context, in *pb.GetGroupMembersReq) (*pb.GetGroupMembersResp, error) {
	members, err := group.App.GetMembers(ctx, in.GroupId)
	return &pb.GetGroupMembersResp{Members: members}, err
}

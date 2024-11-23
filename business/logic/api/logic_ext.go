package api

import (
	"context"
	"gchat/business/logic/domain/device"
	"gchat/business/logic/domain/friend"
	"gchat/business/logic/domain/group"
	"gchat/business/logic/domain/room"
	"gchat/pkg/grpclib"
	"gchat/pkg/protocol/pb"

	"google.golang.org/protobuf/types/known/emptypb"
)

type LogicExtServer struct {
	pb.UnsafeLogicExtServer
}

// RegisterDevice register device
func (*LogicExtServer) RegisterDevice(ctx context.Context, in *pb.RegisterDeviceReq) (*pb.RegisterDeviceResp, error) {
	deviceId, err := device.App.Register(ctx, in)
	return &pb.RegisterDeviceResp{DeviceId: deviceId}, err
}

// PushRoom push room message
func (s *LogicExtServer) PushRoom(ctx context.Context, req *pb.PushRoomReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, room.App.Push(ctx, req)
}

// SendMessageToFriend send message to friend
func (*LogicExtServer) SendMessageToFriend(ctx context.Context, in *pb.SendMessageReq) (*pb.SendMessageResp, error) {
	userId, deviceId, err := grpclib.GetCtxData(ctx)
	if err != nil {
		return nil, err
	}

	seq, err := friend.App.SendToFriend(ctx, deviceId, userId, in)
	if err != nil {
		return nil, err
	}
	return &pb.SendMessageResp{Seq: seq}, nil
}

// AddFriend add a new friend
func (s *LogicExtServer) AddFriend(ctx context.Context, in *pb.AddFriendReq) (*emptypb.Empty, error) {
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
func (s *LogicExtServer) AgreeAddFriend(ctx context.Context, in *pb.AgreeAddFriendReq) (*emptypb.Empty, error) {
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
func (s *LogicExtServer) SetFriend(ctx context.Context, req *pb.SetFriendReq) (*pb.SetFriendResp, error) {
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
func (s *LogicExtServer) GetFriends(ctx context.Context, in *emptypb.Empty) (*pb.GetFriendsResp, error) {
	userId, _, err := grpclib.GetCtxData(ctx)
	if err != nil {
		return nil, err
	}
	friends, err := friend.App.List(ctx, userId)
	return &pb.GetFriendsResp{Friends: friends}, err
}

// SendMessageToGroup send a message to a group
func (*LogicExtServer) SendMessageToGroup(ctx context.Context, in *pb.SendMessageReq) (*pb.SendMessageResp, error) {
	userId, deviceId, err := grpclib.GetCtxData(ctx)
	if err != nil {
		return nil, err
	}

	seq, err := group.App.SendMessage(ctx, deviceId, userId, in)
	if err != nil {
		return nil, err
	}
	return &pb.SendMessageResp{Seq: seq}, nil
}

// CreateGroup create a new group
func (*LogicExtServer) CreateGroup(ctx context.Context, in *pb.CreateGroupReq) (*pb.CreateGroupResp, error) {
	userId, _, err := grpclib.GetCtxData(ctx)
	if err != nil {
		return nil, err
	}

	groupId, err := group.App.CreateGroup(ctx, userId, in)
	return &pb.CreateGroupResp{GroupId: groupId}, err
}

// UpdateGroup update group information
func (*LogicExtServer) UpdateGroup(ctx context.Context, in *pb.UpdateGroupReq) (*emptypb.Empty, error) {
	userId, _, err := grpclib.GetCtxData(ctx)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, group.App.Update(ctx, userId, in)
}

// GetGroup get group information
func (*LogicExtServer) GetGroup(ctx context.Context, in *pb.GetGroupReq) (*pb.GetGroupResp, error) {
	group, err := group.App.GetGroup(ctx, in.GroupId)
	return &pb.GetGroupResp{Group: group}, err
}

// GetGroups get all groups a user belongs to
func (*LogicExtServer) GetGroups(ctx context.Context, in *emptypb.Empty) (*pb.GetGroupsResp, error) {
	userId, _, err := grpclib.GetCtxData(ctx)
	if err != nil {
		return nil, err
	}

	groups, err := group.App.GetUserGroups(ctx, userId)
	return &pb.GetGroupsResp{Groups: groups}, err
}

// AddGroupMembers add members to a group
func (s *LogicExtServer) AddGroupMembers(ctx context.Context, in *pb.AddGroupMembersReq) (*pb.AddGroupMembersResp, error) {
	userId, _, err := grpclib.GetCtxData(ctx)
	if err != nil {
		return nil, err
	}

	userIds, err := group.App.AddMembers(ctx, userId, in.GroupId, in.UserIds)
	return &pb.AddGroupMembersResp{UserIds: userIds}, err
}

// UpdateGroupMember update group member information
func (*LogicExtServer) UpdateGroupMember(ctx context.Context, in *pb.UpdateGroupMemberReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, group.App.UpdateMember(ctx, in)
}

// DeleteGroupMember delete a group member
func (*LogicExtServer) DeleteGroupMember(ctx context.Context, in *pb.DeleteGroupMemberReq) (*emptypb.Empty, error) {
	userId, _, err := grpclib.GetCtxData(ctx)
	if err != nil {
		return nil, err
	}

	err = group.App.DeleteMember(ctx, in.GroupId, in.UserId, userId)
	return &emptypb.Empty{}, err
}

// GetGroupMembers get all members of a group
func (s *LogicExtServer) GetGroupMembers(ctx context.Context, in *pb.GetGroupMembersReq) (*pb.GetGroupMembersResp, error) {
	members, err := group.App.GetMembers(ctx, in.GroupId)
	return &pb.GetGroupMembersResp{Members: members}, err
}

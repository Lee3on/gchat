package api

import (
	"context"
	"gchat/business/logic/domain/device"
	"gchat/business/logic/domain/message"
	"gchat/business/logic/domain/room"
	"gchat/business/logic/proxy"
	"gchat/pkg/logger"
	"gchat/pkg/protocol/pb"

	"google.golang.org/protobuf/types/known/emptypb"
)

type LogicIntServer struct {
	pb.UnsafeLogicIntServer
}

// ConnSignIn sign in a device
func (*LogicIntServer) ConnSignIn(ctx context.Context, req *pb.ConnSignInReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{},
		device.App.SignIn(ctx, req.UserId, req.DeviceId, req.Token, req.ConnAddr, req.ClientAddr)
}

// Sync sync messages among devices
func (*LogicIntServer) Sync(ctx context.Context, req *pb.SyncReq) (*pb.SyncResp, error) {
	return message.App.Sync(ctx, req.UserId, req.Seq)
}

// MessageACK device receives an ACK message
func (*LogicIntServer) MessageACK(ctx context.Context, req *pb.MessageACKReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, message.App.MessageAck(ctx, req.UserId, req.DeviceId, req.DeviceAck)
}

// Offline offline a divice
func (*LogicIntServer) Offline(ctx context.Context, req *pb.OfflineReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, device.App.Offline(ctx, req.DeviceId, req.ClientAddr)
}

// SubscribeRoom subscribe a room
func (s *LogicIntServer) SubscribeRoom(ctx context.Context, req *pb.SubscribeRoomReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, room.App.SubscribeRoom(ctx, req)
}

// Push push messages
func (*LogicIntServer) Push(ctx context.Context, req *pb.PushReq) (*pb.PushResp, error) {
	seq, err := proxy.PushToUserBytes(ctx, req.UserId, req.Code, req.Content, req.IsPersist)
	if err != nil {
		return nil, err
	}
	return &pb.PushResp{Seq: seq}, nil
}

// PushRoom push room messages
func (s *LogicIntServer) PushRoom(ctx context.Context, req *pb.PushRoomReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, room.App.Push(ctx, req)
}

// PushAll push messages to all users
func (s *LogicIntServer) PushAll(ctx context.Context, req *pb.PushAllReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, message.App.PushAll(ctx, req)
}

// GetDevice get device information
func (*LogicIntServer) GetDevice(ctx context.Context, req *pb.GetDeviceReq) (*pb.GetDeviceResp, error) {
	device, err := device.App.GetDevice(ctx, req.DeviceId)
	return &pb.GetDeviceResp{Device: device}, err
}

// ServerStop stop a server
func (s *LogicIntServer) ServerStop(ctx context.Context, in *pb.ServerStopReq) (*emptypb.Empty, error) {
	go func() {
		err := device.App.ServerStop(ctx, in.ConnAddr)
		if err != nil {
			logger.Sugar.Error(err)
		}
	}()
	return &emptypb.Empty{}, nil
}

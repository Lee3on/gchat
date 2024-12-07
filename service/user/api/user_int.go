package api

import (
	"context"
	"gchat/pkg/protocol/pb"
	app2 "gchat/service/user/domain/app"

	"google.golang.org/protobuf/types/known/emptypb"
)

type UserIntServer struct {
	pb.UnsafeUserIntServer
}

func (*UserIntServer) Auth(ctx context.Context, req *pb.AuthReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, app2.AuthApp.Auth(ctx, req.UserId, req.DeviceId, req.Token)
}

func (*UserIntServer) GetUser(ctx context.Context, req *pb.GetUserReq) (*pb.GetUserResp, error) {
	user, err := app2.UserApp.Get(ctx, req.UserId)
	return &pb.GetUserResp{User: user}, err
}

func (*UserIntServer) GetUsers(ctx context.Context, req *pb.GetUsersReq) (*pb.GetUsersResp, error) {
	var userIds = make([]int64, 0, len(req.UserIds))
	for k := range req.UserIds {
		userIds = append(userIds, k)
	}

	users, err := app2.UserApp.GetByIds(ctx, userIds)
	return &pb.GetUsersResp{Users: users}, err
}

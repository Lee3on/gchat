package api

import (
	"context"
	"fmt"
	"gchat/pkg/logger"
	"gchat/pkg/protocol/pb"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func getLogicIntClient() pb.LogicIntClient {
	conn, err := grpc.Dial("127.0.0.1:8010", grpc.WithInsecure())
	if err != nil {
		logger.Sugar.Error(err)
		return nil
	}
	return pb.NewLogicIntClient(conn)
}

func TestLogicIntServer_SignIn(t *testing.T) {
	token := ""

	resp, err := getLogicIntClient().ConnSignIn(context.TODO(),
		&pb.ConnSignInReq{
			DeviceId: 1,
			UserId:   1,
			Token:    token,
			ConnAddr: "127.0.0.1:5000",
		})
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	logger.Sugar.Info(resp)
}

func TestLogicIntServer_Sync(t *testing.T) {
	resp, err := getLogicIntClient().Sync(metadata.NewOutgoingContext(context.TODO(), metadata.Pairs("key", "val")),
		&pb.SyncReq{
			UserId:   1,
			DeviceId: 1,
			Seq:      0,
		})
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	logger.Sugar.Info(resp)
}

func TestLogicIntServer_MessageACK(t *testing.T) {
	resp, err := getLogicIntClient().MessageACK(metadata.NewOutgoingContext(context.TODO(), metadata.Pairs("key", "val")),
		&pb.MessageACKReq{
			UserId:      1,
			DeviceId:    1,
			DeviceAck:   1,
			ReceiveTime: 1,
		})
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	logger.Sugar.Info(resp)
}

func TestLogicIntServer_Offline(t *testing.T) {
	resp, err := getLogicIntClient().Offline(metadata.NewOutgoingContext(context.TODO(), metadata.Pairs("key", "val")),
		&pb.OfflineReq{
			UserId:   1,
			DeviceId: 1,
		})
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	logger.Sugar.Info(resp)
}

func TestLogicIntServer_PushRoom(t *testing.T) {
	resp, err := getLogicIntClient().PushRoom(getCtx(),
		&pb.PushRoomReq{
			RoomId:  1,
			Code:    1,
			Content: []byte("hahaha"),
		})
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestLogicIntServer_PushAll(t *testing.T) {
	resp, err := getLogicIntClient().PushAll(getCtx(),
		&pb.PushAllReq{
			Code:    1,
			Content: []byte("hahaha"),
		})
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	fmt.Printf("%+v\n", resp)
}

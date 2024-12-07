package api

import (
	"context"
	"fmt"
	"gchat/pkg/protocol/pb"
	"gchat/pkg/util"
	"strconv"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func getLogicExtClient() pb.LogicExtClient {
	conn, err := grpc.Dial("127.0.0.1:8010", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return pb.NewLogicExtClient(conn)
}

func getCtx() context.Context {
	token := "0"
	return metadata.NewOutgoingContext(context.TODO(), metadata.Pairs(
		"user_id", "2",
		"device_id", "2",
		"token", token,
		"request_id", strconv.FormatInt(time.Now().UnixNano(), 10)))
}

func TestLogicExtServer_SendMessageToFriend(t *testing.T) {
	resp, err := getLogicExtClient().SendMessageToFriend(getCtx(),
		&pb.SendMessageReq{
			ReceiverId: 2,
			Content:    []byte("test1to2-c"),
			SendTime:   util.UnixMilliTime(time.Now()),
		})
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestLogicExtServer_SendMessageToGroup(t *testing.T) {
	resp, err := getLogicExtClient().SendMessageToGroup(getCtx(),
		&pb.SendMessageReq{
			ReceiverId: 1,
			Content:    []byte("group message-d"),
			SendTime:   util.UnixMilliTime(time.Now()),
		})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", resp)
}

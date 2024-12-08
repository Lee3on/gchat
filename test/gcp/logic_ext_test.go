package gcp

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"gchat/pkg/protocol/pb"
	"gchat/pkg/util"
	"google.golang.org/grpc/credentials"
	"strconv"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func GetLogicExtClient() pb.LogicExtClient {
	host := "logic-service-653320394232.us-central1.run.app:443"
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithAuthority(host))
	systemRoots, err := x509.SystemCertPool()
	if err != nil {
		return nil
	}
	cred := credentials.NewTLS(&tls.Config{
		RootCAs: systemRoots,
	})
	opts = append(opts, grpc.WithTransportCredentials(cred))
	conn, err := grpc.Dial(host, opts...)
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
	resp, err := GetLogicExtClient().SendMessageToFriend(getCtx(),
		&pb.SendMessageReq{
			ReceiverId: 1,
			Content:    []byte("test2to1-test1"),
			SendTime:   util.UnixMilliTime(time.Now()),
		})
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestLogicExtServer_SendMessageToGroup(t *testing.T) {
	resp, err := GetLogicExtClient().SendMessageToGroup(getCtx(),
		&pb.SendMessageReq{
			ReceiverId: 4,
			Content:    []byte("group message "),
			SendTime:   util.UnixMilliTime(time.Now()),
		})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", resp)
}

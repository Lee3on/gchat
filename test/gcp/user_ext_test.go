package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"gchat/pkg/protocol/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"testing"
)

func getUserExtClient() pb.UserExtClient {
	host := "user-service-653320394232.us-central1.run.app:443"
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
	return pb.NewUserExtClient(conn)
}

func TestUserExtServer_SignIn(t *testing.T) {
	resp, err := getUserExtClient().SignIn(getCtx(), &pb.SignInReq{
		PhoneNumber: "22222222222",
		Code:        "0",
		DeviceId:    2,
	})
	if err != nil {
		fmt.Println(err)
		t.Fatalf("err: %v", err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestUserExtServer_GetUser(t *testing.T) {
	resp, err := getUserExtClient().GetUser(getCtx(), &pb.GetUserReq{UserId: 1})
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestUserExtServer_UpdateUser(t *testing.T) {
	resp, err := getUserExtClient().UpdateUser(getCtx(), &pb.UpdateUserReq{
		Nickname: "test",
	})
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	fmt.Printf("%+v\n", resp)
}

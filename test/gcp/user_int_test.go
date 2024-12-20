package gcp

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"gchat/pkg/protocol/pb"
	"google.golang.org/grpc/credentials"
	"testing"

	"google.golang.org/grpc"
)

func getUserIntClient() pb.UserIntClient {
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
	return pb.NewUserIntClient(conn)
}

func TestUserIntServer_Auth(t *testing.T) {
	_, err := getUserIntClient().Auth(getCtx(), &pb.AuthReq{
		UserId:   1,
		DeviceId: 1,
		Token:    "0",
	})
	fmt.Println(err)
}

func TestUserIntServer_GetUsers(t *testing.T) {
	resp, err := getUserIntClient().GetUsers(getCtx(), &pb.GetUsersReq{
		UserIds: map[int64]int32{
			1: 0,
			2: 0,
			3: 0,
		},
	})
	if err != nil {
		t.Fatalf("err: %v", err)
	}

	for k, v := range resp.Users {
		fmt.Printf("%+-5v  %+v\n", k, v)
	}
}

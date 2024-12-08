package gcp

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"gchat/pkg/protocol/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
	"strconv"
	"testing"
	"time"
)

func GetUserExtClient() pb.UserExtClient {
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

func TestUserExtServer_RegisterDevice(t *testing.T) {
	resp, err := GetUserExtClient().RegisterDevice(context.TODO(),
		&pb.RegisterDeviceReq{
			Type:          2,
			Brand:         "apple",
			Model:         "iphone 16pro",
			SystemVersion: "1.0.0",
			SdkVersion:    "1.0.0",
		})
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestUserExtServer_SignIn(t *testing.T) {
	resp, err := GetUserExtClient().SignIn(getCtx(), &pb.SignInReq{
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
	resp, err := GetUserExtClient().GetUser(getCtx(), &pb.GetUserReq{UserId: 1})
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestUserExtServer_UpdateUsers(t *testing.T) {
	for i := 1; i <= 2000; i++ {
		ctx := metadata.NewOutgoingContext(context.TODO(), metadata.Pairs(
			"user_id", fmt.Sprintf("%d", i),
			"device_id", fmt.Sprintf("%d", i),
			"token", "0",
			"request_id", strconv.FormatInt(time.Now().UnixNano(), 10)))
		resp, err := GetUserExtClient().UpdateUser(ctx, &pb.UpdateUserReq{
			PhoneNumber: fmt.Sprintf("%d", i),
		})
		if err != nil {
			t.Fatalf("err: %v", err)
		}
		fmt.Printf("%+v\n", resp)
	}
}

func TestUserExtServer_UpdateUser(t *testing.T) {
	resp, err := GetUserExtClient().UpdateUser(getCtx(), &pb.UpdateUserReq{
		PhoneNumber: "3033",
	})
	if err != nil {
		t.Fatalf("err: %v", err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestLogicExtServer_CreateGroup(t *testing.T) {
	resp, err := GetUserExtClient().CreateGroup(getCtx(),
		&pb.CreateGroupReq{
			Name:         "10",
			Introduction: "10",
			Extra:        "10",
		})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", resp)
}

func TestLogicExtServer_UpdateGroup(t *testing.T) {
	resp, err := GetUserExtClient().UpdateGroup(getCtx(),
		&pb.UpdateGroupReq{
			GroupId:      2,
			Name:         "11",
			Introduction: "11",
			Extra:        "11",
		})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", resp)
}

func TestLogicExtServer_GetGroup(t *testing.T) {
	resp, err := GetUserExtClient().GetGroup(getCtx(),
		&pb.GetGroupReq{
			GroupId: 2,
		})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", resp)
}

func TestLogicExtServer_GetUserGroups(t *testing.T) {
	resp, err := GetUserExtClient().GetGroups(getCtx(), &emptypb.Empty{})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", resp)
}

func TestLogicExtServer_AddGroupMember(t *testing.T) {
	resp, err := GetUserExtClient().AddGroupMembers(getCtx(),
		&pb.AddGroupMembersReq{
			GroupId: 2,
		})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", resp)
}

func TestLogicExtServer_UpdateGroupMember(t *testing.T) {
	resp, err := GetUserExtClient().UpdateGroupMember(getCtx(),
		&pb.UpdateGroupMemberReq{
			GroupId: 2,
			UserId:  3,
			Remarks: "2",
			Extra:   "2",
		})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", resp)
}

func TestLogicExtServer_DeleteGroupMember(t *testing.T) {
	resp, err := GetUserExtClient().DeleteGroupMember(getCtx(),
		&pb.DeleteGroupMemberReq{
			GroupId: 10,
			UserId:  1,
		})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", resp)
}

func TestLogicExtServer_GetGroupMembers(t *testing.T) {
	resp, err := GetUserExtClient().GetGroupMembers(getCtx(),
		&pb.GetGroupMembersReq{
			GroupId: 2,
		})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", resp)
}

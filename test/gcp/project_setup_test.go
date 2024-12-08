package gcp

import (
	"context"
	"gchat/pkg/protocol/pb"
	"testing"
)

func Test_RegisterDevice(t *testing.T) {
	for i := 0; i < 1988; i++ {
		_, err := GetUserExtClient().RegisterDevice(context.TODO(),
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
	}
}

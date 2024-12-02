package config

import (
	"context"
	"gchat/pkg/gerrors"
	"gchat/pkg/protocol/pb"
	"os"

	"google.golang.org/grpc"
)

var builders = map[string]Builder{
	"default": &defaultBuilder{},
	"gcp":     &gcpBuilder{},
}

var Config Configuration

type Builder interface {
	Build() Configuration
}

type Configuration struct {
	MySQL                string
	RedisHost            string
	RedisPassword        string
	PushRoomSubscribeNum int
	PushAllSubscribeNum  int

	ConnectLocalAddr     string
	ConnectRPCListenAddr string
	ConnectTCPListenAddr string
	ConnectWSListenAddr  string

	LogicRPCListenAddr string
	UserRPCListenAddr  string
	FileHTTPListenAddr string

	ConnectIntClientBuilder func() pb.ConnectIntClient
	LogicIntClientBuilder   func() pb.LogicIntClient
	UserIntClientBuilder    func() pb.BusinessIntClient
}

func init() {
	env := os.Getenv("GCHAT_ENV")
	builder, ok := builders[env]
	if !ok {
		builder = new(defaultBuilder)
	}
	Config = builder.Build()

}

func interceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	err := invoker(ctx, method, req, reply, cc, opts...)
	return gerrors.WrapRPCError(err)
}

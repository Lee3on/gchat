package main

import (
	"context"
	"gchat/config"
	"gchat/pkg/interceptor"
	"gchat/pkg/logger"
	"gchat/pkg/protocol/pb"
	"gchat/pkg/rpc"
	"gchat/service/connect"
	"net"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	// Start the TCP long connection server
	go func() {
		connect.StartTCPServer(config.Config.ConnectTCPListenAddr)
	}()

	// Start the WebSocket long connection server
	go func() {
		connect.StartWSServer(config.Config.ConnectWSListenAddr)
	}()

	// Start the service subscription
	connect.StartSubscribe()

	server := grpc.NewServer(grpc.UnaryInterceptor(interceptor.NewInterceptor("connect_interceptor", nil)))

	// Listen for service shutdown signals and perform a graceful restart of the service
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGTERM)
		s := <-c
		logger.Logger.Info("server stop start", zap.Any("signal", s))
		_, _ = rpc.GetLogicIntClient().ServerStop(context.TODO(), &pb.ServerStopReq{ConnAddr: config.Config.ConnectLocalAddr})
		logger.Logger.Info("server stop end")

		server.GracefulStop()
	}()

	// Register the service
	pb.RegisterConnectIntServer(server, &connect.ConnIntServer{})
	listener, err := net.Listen("tcp", config.Config.ConnectRPCListenAddr)
	if err != nil {
		panic(err)
	}

	logger.Logger.Info("Connect server start", zap.String("addr", config.Config.ConnectRPCListenAddr))
	err = server.Serve(listener)
	if err != nil {
		logger.Logger.Error("serve error", zap.Error(err))
	}
}

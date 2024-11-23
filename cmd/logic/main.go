package main

import (
	"gchat/business/logic/api"
	"gchat/business/logic/domain/device"
	"gchat/business/logic/domain/message"
	"gchat/business/logic/proxy"
	"gchat/config"
	"gchat/pkg/interceptor"
	"gchat/pkg/logger"
	"gchat/pkg/protocol/pb"
	"gchat/pkg/urlwhitelist"
	"net"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func init() {
	proxy.MessageProxy = message.App
	proxy.DeviceProxy = device.App
}

func main() {
	server := grpc.NewServer(grpc.UnaryInterceptor(interceptor.NewInterceptor("logic_interceptor", urlwhitelist.Logic)))

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGTERM)
		s := <-c
		logger.Logger.Info("server stop", zap.Any("signal", s))
		server.GracefulStop()
	}()

	pb.RegisterLogicIntServer(server, &api.LogicIntServer{})
	pb.RegisterLogicExtServer(server, &api.LogicExtServer{})
	listen, err := net.Listen("tcp", config.Config.LogicRPCListenAddr)
	if err != nil {
		panic(err)
	}

	logger.Logger.Info("Logic server start", zap.String("addr", config.Config.LogicRPCListenAddr))
	err = server.Serve(listen)
	if err != nil {
		logger.Logger.Error("serve error", zap.Error(err))
	}
}

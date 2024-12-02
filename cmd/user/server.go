package main

import (
	"gchat/business/user/api"
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

func main() {
	server := grpc.NewServer(grpc.UnaryInterceptor(interceptor.NewInterceptor("business_interceptor", urlwhitelist.Business)))

	// Listen for service shutdown signals, service graceful restart
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGTERM)
		s := <-c
		logger.Logger.Info("server stop", zap.Any("signal", s))
		server.GracefulStop()
	}()

	pb.RegisterBusinessIntServer(server, &api.BusinessIntServer{})
	pb.RegisterBusinessExtServer(server, &api.BusinessExtServer{})
	listen, err := net.Listen("tcp", config.Config.UserRPCListenAddr)
	if err != nil {
		panic(err)
	}

	logger.Logger.Info("user server start", zap.String("addr", config.Config.UserRPCListenAddr))
	err = server.Serve(listen)
	if err != nil {
		logger.Logger.Error("serve error", zap.Error(err))
	}
}

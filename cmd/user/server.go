package main

import (
	"gchat/config"
	"gchat/pkg/interceptor"
	"gchat/pkg/logger"
	"gchat/pkg/protocol/pb"
	"gchat/pkg/urlwhitelist"
	"gchat/service/user/api"
	"net"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer(grpc.UnaryInterceptor(interceptor.NewInterceptor("user_interceptor", urlwhitelist.User)))

	// Listen for service shutdown signals, service graceful restart
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGTERM)
		s := <-c
		logger.Logger.Info("server stop", zap.Any("signal", s))
		server.GracefulStop()
	}()

	pb.RegisterUserIntServer(server, &api.UserIntServer{})
	pb.RegisterUserExtServer(server, &api.UserExtServer{})
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

package config

import (
	"context"
	"fmt"
	"gchat/pkg/grpclib/picker"
	_ "gchat/pkg/grpclib/resolver/addrs"
	"gchat/pkg/logger"
	"gchat/pkg/protocol/pb"
	"os"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
)

type defaultBuilder struct{}

func (*defaultBuilder) Build() Configuration {
	logger.Level = zap.DebugLevel
	logger.Target = logger.Console
	mysqlHost := os.Getenv("MYSQLHOST")
	mysqlPort := os.Getenv("MYSQLPORT")
	mysqlPassword := os.Getenv("MYSQLPASSWORD")
	mysqlUser := os.Getenv("MYSQLUSER")
	if mysqlHost == "" {
		mysqlHost = "127.0.0.1" // Default MySQL host
	}
	if mysqlPort == "" {
		mysqlPort = "3306" // Default MySQL port
	}
	if mysqlUser == "" {
		mysqlUser = "root"
	}
	if mysqlPassword == "" {
		mysqlPassword = "jason123456"
	}
	mysqlAddr := os.Getenv("MYSQLADDR")
	if mysqlAddr == "" {
		mysqlAddr = fmt.Sprintf("%s:%s", mysqlHost, mysqlPort)
	}

	redisHost := os.Getenv("REDISHOST")
	redisPort := os.Getenv("REDISPORT")
	if redisHost == "" {
		redisHost = "127.0.0.1" // Default Redis host
	}
	if redisPort == "" {
		redisPort = "6379" // Default Redis port
	}
	redisPassword := os.Getenv("REDISPASSWORD")

	connector := os.Getenv("DB_CONNECTOR")
	if connector == "" {
		connector = "tcp"
	}

	connectPort := os.Getenv("CONNECT_PORT")
	if connectPort == "" {
		connectPort = "8000"
	}

	return Configuration{
		MySQL:                fmt.Sprintf("%s:%s@%s(%s)/gchat?charset=utf8&parseTime=true", mysqlUser, mysqlPassword, connector, mysqlAddr),
		RedisHost:            fmt.Sprintf("%s:%s", redisHost, redisPort),
		RedisPassword:        redisPassword,
		PushRoomSubscribeNum: 100,
		PushAllSubscribeNum:  100,

		ConnectLocalAddr:     "127.0.0.1:8000",
		ConnectRPCListenAddr: ":" + connectPort,
		ConnectTCPListenAddr: ":8001",
		ConnectWSListenAddr:  ":8002",

		LogicRPCListenAddr: ":8010",
		UserRPCListenAddr:  ":8020",
		FileHTTPListenAddr: "8030",

		ConnectIntClientBuilder: func() pb.ConnectIntClient {
			conn, err := grpc.DialContext(context.TODO(), "addrs:///127.0.0.1:8000", grpc.WithInsecure(), grpc.WithUnaryInterceptor(interceptor),
				grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, picker.AddrPickerName)))
			if err != nil {
				panic(err)
			}
			return pb.NewConnectIntClient(conn)
		},
		LogicIntClientBuilder: func() pb.LogicIntClient {
			conn, err := grpc.DialContext(context.TODO(), "addrs:///127.0.0.1:8010", grpc.WithInsecure(), grpc.WithUnaryInterceptor(interceptor),
				grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)))
			if err != nil {
				panic(err)
			}
			return pb.NewLogicIntClient(conn)
		},
		UserIntClientBuilder: func() pb.UserIntClient {
			conn, err := grpc.DialContext(context.TODO(), "addrs:///127.0.0.1:8020", grpc.WithInsecure(), grpc.WithUnaryInterceptor(interceptor),
				grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)))
			if err != nil {
				panic(err)
			}
			return pb.NewUserIntClient(conn)
		},
	}
}

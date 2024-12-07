package config

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"gchat/pkg/grpclib/picker"
	"gchat/pkg/k8sutil"
	"gchat/pkg/logger"
	"gchat/pkg/protocol/pb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"os"
	"strconv"
)

type gcpBuilder struct{}

func (*gcpBuilder) Build() Configuration {
	// Get mysql configuration from env
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
	mysqlAddr := os.Getenv("MYSQLADDR")
	if mysqlAddr == "" {
		mysqlAddr = fmt.Sprintf("%s:%s", mysqlHost, mysqlPort)
	}
	// Get connector configuration from env
	connector := os.Getenv("DB_CONNECTOR")
	if connector == "" {
		connector = "tcp"
	}
	mysqlUrl := fmt.Sprintf("%s:%s@%s(%s)/gchat?charset=utf8&parseTime=true", mysqlUser, mysqlPassword, connector, mysqlAddr)

	// Get redis configuration from env
	redisHost := os.Getenv("REDISHOST")
	redisPort := os.Getenv("REDISPORT")
	if redisHost == "" {
		redisHost = "127.0.0.1" // Default Redis host
	}
	if redisPort == "" {
		redisPort = "6379" // Default Redis port
	}
	redisHost = fmt.Sprintf("%s:%s", redisHost, redisPort)
	redisPassword := os.Getenv("REDISPASSWORD")

	pushRoomSubscribeNum := 100
	pushAllSubscribeNum := 100
	// Get cloud service configuration from env
	cloudService := os.Getenv("CLOUD_SERVICE")
	if cloudService == "k8s" {
		const namespace = "default"

		k8sClient, err := k8sutil.GetK8sClient()
		if err != nil {
			panic(err)
		}
		configmap, err := k8sClient.CoreV1().ConfigMaps(namespace).Get(context.TODO(), "config", metav1.GetOptions{})
		if err != nil {
			panic(err)
		}

		mysqlUrl = configmap.Data["mysql"]
		redisHost = configmap.Data["redisIP"]
		redisPassword = configmap.Data["redisPassword"]
		pushRoomSubscribeNum = getInt(configmap.Data, "pushRoomSubscribeNum")
		pushAllSubscribeNum = getInt(configmap.Data, "pushAllSubscribeNum")
	}

	const (
		RPCListenAddr = ":8000"
		RPCDialAddr   = "8000"
	)

	logger.Level = zap.DebugLevel
	logger.Target = logger.Console

	return Configuration{
		MySQL:                mysqlUrl,
		RedisHost:            redisHost,
		RedisPassword:        redisPassword,
		PushRoomSubscribeNum: pushRoomSubscribeNum,
		PushAllSubscribeNum:  pushAllSubscribeNum,

		ConnectLocalAddr:     os.Getenv("POD_IP") + RPCListenAddr,
		ConnectRPCListenAddr: ":8000",
		ConnectTCPListenAddr: ":8001",
		ConnectWSListenAddr:  ":8002",

		LogicRPCListenAddr: ":8010",
		UserRPCListenAddr:  ":8020",

		ConnectIntClientBuilder: func() pb.ConnectIntClient {
			conn, err := grpc.Dial("34.44.91.248:"+RPCDialAddr,
				grpc.WithTransportCredentials(insecure.NewCredentials()),
				grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, picker.AddrPickerName)))
			if err != nil {
				panic(err)
			}
			return pb.NewConnectIntClient(conn)
		},
		LogicIntClientBuilder: func() pb.LogicIntClient {
			host := "logic-service-653320394232.us-central1.run.app:443"
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
			opts = append(opts, grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)))
			conn, err := grpc.Dial(host, opts...)
			if err != nil {
				panic(err)
			}
			return pb.NewLogicIntClient(conn)
		},
		UserIntClientBuilder: func() pb.UserIntClient {
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
			opts = append(opts, grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, roundrobin.Name)))
			conn, err := grpc.Dial(host, opts...)
			if err != nil {
				panic(err)
			}
			return pb.NewUserIntClient(conn)
		},
	}
}

func getInt(m map[string]string, key string) int {
	value, _ := strconv.Atoi(m[key])
	return value
}

package user

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"gchat/pkg/protocol/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

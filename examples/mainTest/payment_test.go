package mainTest

import (
	"github.com/entysquare/payment-sdk-go"
	"github.com/entysquare/payment-sdk-go/lib"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"testing"
)

func TestServer(t *testing.T) {
	creds, err := credentials.NewServerTLSFromFile("../testkeys/server.pem", "../testkeys/server.key")
	lis, err := net.Listen("tcp", ":8028") //监听所有网卡8028端口的TCP连接
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}
	s := grpc.NewServer(grpc.Creds(creds)) //创建gRPC服务
	//config.RegisterWaiterServer(s, &payApiClien{cc})
	lib.RegisterPayApiServer(s, &payment.MockPayServer{})
	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

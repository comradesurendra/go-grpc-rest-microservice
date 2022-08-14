package rpcserver

import (
	"GRPC-SERVE/app/home"
	jdconfig "GRPC-SERVE/config"
	"GRPC-SERVE/middleware"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// StartServer Main rpc server Initialize
func StartServer() {
	port := jdconfig.Cfg.Ports.RpcPort
	ip := jdconfig.Cfg.Ports.RpcIp
	tcp := ip + ":" + port
	lis, err := net.Listen("tcp", tcp)
	fmt.Println("Listening on", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(middleware.UnaryInterceptor), grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: 5 * time.Minute, // <--- This fixes it!
	}))

	home.RegisterLoanServiceServer(s, &home.HomeStruct{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

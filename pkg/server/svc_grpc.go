package server

import (
	"fmt"
	"github.com/dev2choiz/hello/pkg/grpc_handlers"
	"github.com/dev2choiz/hello/pkg/protobuf/pingpb"
	"google.golang.org/grpc"
	"log"
	"net"
)

func ExecuteSvcGrpc(conf *Config) {
	grpcServer := grpc.NewServer()
	pingpb.RegisterPingServer(grpcServer, &grpc_handlers.PingServer{ SvcName: conf.Name })

	lis, err :=net.Listen("tcp", fmt.Sprintf(":%s", conf.GrpcPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer lis.Close()

	// start gRPC server
	log.Println("starting gRPC server...")
	log.Println(fmt.Sprintf(":%s", conf.GrpcPort))
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}


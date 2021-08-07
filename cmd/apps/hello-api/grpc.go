package main

import (
	"fmt"
	"github.com/dev2choiz/hello/pkg/grpc_handlers"
	"github.com/dev2choiz/hello/pkg/protobuf/healthpb"
	"github.com/dev2choiz/hello/pkg/protobuf/notifypb"
	"github.com/dev2choiz/hello/pkg/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func executeApiGrpc(conf *server.Config) {
	grpcServer := grpc.NewServer()
	healthpb.RegisterHealthServer(grpcServer, &grpc_handlers.HealthServer{})
	notifypb.RegisterNotifyServer(grpcServer, &grpc_handlers.NotifyServer{})

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


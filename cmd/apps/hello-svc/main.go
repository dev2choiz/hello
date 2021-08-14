package main

import (
	"github.com/dev2choiz/hello/pkg/handlers"
	"github.com/dev2choiz/hello/pkg/protobuf/pingpb"
	"github.com/dev2choiz/hello/pkg/server"
	"google.golang.org/grpc"
)

func main() {
	conf := server.GetConfig()
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(server.LogInterceptor))
	pingpb.RegisterPingServer(grpcServer, &handlers.PingServer{ SvcName: conf.Name })

	server.RunGrpcServer(grpcServer, conf)
}

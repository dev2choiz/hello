package main

import (
	"github.com/dev2choiz/hello/pkg/handlers"
	"github.com/dev2choiz/hello/pkg/protobuf/healthpb"
	"github.com/dev2choiz/hello/pkg/protobuf/notifypb"
	"github.com/dev2choiz/hello/pkg/server"
	"google.golang.org/grpc"
)

func main() {
	conf := server.GetConfig()
	executeApiGrpc(conf)
}

func executeApiGrpc(conf *server.Config) {
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(server.LogInterceptor))
	healthpb.RegisterHealthServer(grpcServer, &handlers.HealthServer{})
	notifypb.RegisterNotifyServer(grpcServer, &handlers.NotifyServer{})

	server.RunGrpcServer(grpcServer, conf)
}

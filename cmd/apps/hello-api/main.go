package main

import (
	"github.com/dev2choiz/hello/internal/server"
	"github.com/dev2choiz/hello/pkg/handlers"
	"github.com/dev2choiz/hello/pkg/protobuf/healthpb"
	"github.com/dev2choiz/hello/pkg/protobuf/notifypb"
	"github.com/dev2choiz/hello/pkg/protobuf/sandboxpb"
	"google.golang.org/grpc"
)

func main() {
	conf := server.GetConfig()
	conf.Name = "hello-api"
	executeApiGrpc(conf)
}

func executeApiGrpc(conf *server.Config) {
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(server.LogInterceptor))
	healthpb.RegisterHealthServer(grpcServer, &handlers.HealthServer{})
	notifypb.RegisterNotifyServer(grpcServer, &handlers.NotifyServer{})
	sandboxpb.RegisterSandboxServer(grpcServer, &handlers.SandboxServer{})

	server.RunGrpcServer(grpcServer, conf)
}

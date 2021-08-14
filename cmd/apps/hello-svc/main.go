package main

import (
	"context"
	"fmt"
	"github.com/dev2choiz/hello/internal/logger"
	"github.com/dev2choiz/hello/pkg/handlers"
	"github.com/dev2choiz/hello/pkg/protobuf/pingpb"
	"github.com/dev2choiz/hello/pkg/server"
	"google.golang.org/grpc"
	"net"
)

func main() {
	conf := server.GetConfig()
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		logger.Info(info.FullMethod)
		h, err := handler(ctx, req)
		if err != nil {
			logger.Error(err.Error())
		}
		return h, err
	}))
	pingpb.RegisterPingServer(grpcServer, &handlers.PingServer{ SvcName: conf.Name })

	lis, err :=net.Listen("tcp", fmt.Sprintf(":%s", conf.Port))
	if err != nil {
		logger.Fatal("Failed to listen: " + err.Error())
	}
	defer lis.Close()

	// start gRPC server
	logger.Infof("starting gRPC server on :%s", conf.Port)
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}

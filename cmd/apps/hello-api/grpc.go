package main

import (
	"context"
	"fmt"
	"github.com/dev2choiz/hello/internal/logger"
	"github.com/dev2choiz/hello/pkg/grpc_handlers"
	"github.com/dev2choiz/hello/pkg/protobuf/healthpb"
	"github.com/dev2choiz/hello/pkg/protobuf/notifypb"
	"github.com/dev2choiz/hello/pkg/server"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

func executeApiGrpc(conf *server.Config) {
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		logger.Info(info.FullMethod)
		logger.RInfo(info.FullMethod)
		h, err := handler(ctx, req)
		if err != nil {
			logger.Error(err)
		}
		return h, err
	}))
	healthpb.RegisterHealthServer(grpcServer, &grpc_handlers.HealthServer{})
	notifypb.RegisterNotifyServer(grpcServer, &grpc_handlers.NotifyServer{})

	lis, err :=net.Listen("tcp", fmt.Sprintf(":%s", conf.GrpcPort))
	if err != nil {
		logger.Fatalf("Failed to listen: %v", err)
	}
	defer lis.Close()

	// start gRPC server
	logger.Info("starting gRPC server...")
	logger.Infof(":%s", conf.GrpcPort)

	logger.RInfo("starting gRPC server...")
	logger.RInfo(":%s", zap.String("port", conf.GrpcPort))
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
}

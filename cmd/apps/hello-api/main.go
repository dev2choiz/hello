package main

import (
	"context"
	"fmt"
	"github.com/dev2choiz/hello/internal/logger"
	"github.com/dev2choiz/hello/pkg/handlers"
	"github.com/dev2choiz/hello/pkg/protobuf/healthpb"
	"github.com/dev2choiz/hello/pkg/protobuf/notifypb"
	"github.com/dev2choiz/hello/pkg/server"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	conf := server.GetConfig()
	executeApiGrpc(conf)
}

func executeApiGrpc(conf *server.Config) {
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		logger.Info(info.FullMethod)
		h, err := handler(ctx, req)
		if err != nil {
			logger.Error(err.Error())
		}
		return h, err
	}))
	healthpb.RegisterHealthServer(grpcServer, &handlers.HealthServer{})
	notifypb.RegisterNotifyServer(grpcServer, &handlers.NotifyServer{})

	lis, err :=net.Listen("tcp", fmt.Sprintf(":%s", conf.Port))
	if err != nil {
		logger.Fatal("Failed to listen: " + err.Error())
	}
	defer lis.Close()

	errChan := make(chan error)
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		logger.Infof("starting gRPC server on :%s", conf.Port)
		if err = grpcServer.Serve(lis); err != nil {
			errChan <- err
		}
	}()
	defer func() {
		logger.Warn("Terminating...")
		grpcServer.GracefulStop()
	}()

	select {
		case err := <-errChan:
			logger.Error("Fatal error: " + err.Error())
		case <-stopChan:
	}
}

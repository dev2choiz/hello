package server

import (
	"context"
	"fmt"
	"github.com/dev2choiz/hello/pkg/logger"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Config struct {
	Name string
	Port string
}

var RunConfig = &Config{}

func GetConfig() *Config {
	return RunConfig
}

func RunGrpcServer(grpcServer *grpc.Server, conf *Config) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", conf.Port))
	if err != nil {
		logger.Fatal("Failed to listen: " + err.Error())
	}
	defer lis.Close()

	errChan := make(chan error)
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		logger.Infof("starting %s gRPC server on :%s", conf.Name, conf.Port)
		if err = grpcServer.Serve(lis); err != nil {
			errChan <- err
		}
	}()
	defer func() {
		logger.Warnf("will stop %s grpc server gracefully...", conf.Name)
		grpcServer.GracefulStop()
		logger.Warnf("%s grpc server stopped", conf.Name)
	}()

	select {
	case err := <-errChan:
		logger.Error("Fatal error: " + err.Error())
	case <-stopChan:
	}
}

func LogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	logger.Info(info.FullMethod)
	h, err := handler(ctx, req)
	if err != nil {
		logger.Error(err.Error())
	}
	return h, err
}

func LogStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	logger.Info(info.FullMethod)
	err := handler(srv, ss)
	if err != nil {
		logger.Error(err.Error())
	}
	return err
}

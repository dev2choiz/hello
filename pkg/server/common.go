package server

import (
	"context"
	"flag"
	"fmt"
	"github.com/dev2choiz/hello/internal/logger"
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

func init() {
	flag.StringVar(&RunConfig.Name, "name", "", "svc name")
	flag.StringVar(&RunConfig.Port, "port", "", "grpc port")
	flag.Parse()

	logger.Infof("config name=%s port=%s", RunConfig.Name, RunConfig.Port)
}

func GetConfig() *Config {
	return RunConfig
}

func RunGrpcServer(grpcServer * grpc.Server, conf *Config) {
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

func LogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	logger.Info(info.FullMethod)
	h, err := handler(ctx, req)
	if err != nil {
		logger.Error(err.Error())
	}
	return h, err
}

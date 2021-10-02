package server

import (
	"context"
	"fmt"
	"github.com/dev2choiz/hello/pkg/config"
	"github.com/dev2choiz/hello/pkg/logger"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// RunGrpcServer handle listening and graceful stop of a grpc server
func RunGrpcServer(grpcServer *grpc.Server, conf *config.Config) {
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

// LogInterceptor a grpc unary interceptor to log requests
func LogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	logger.Info(info.FullMethod)
	h, err := handler(ctx, req)
	if err != nil {
		logger.Error(err.Error())
	}
	return h, err
}

// LogStreamInterceptor a grpc stream interceptor to log requests
func LogStreamInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
	logger.Info(info.FullMethod)
	err := handler(srv, ss)
	if err != nil {
		logger.Error(err.Error())
	}
	return err
}

// RunHttpServer handle listening and graceful stop of a http server
func RunHttpServer(server *http.Server, conf *config.Config) {
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
		if err = server.Serve(lis); err != nil {
			errChan <- err
		}
	}()

	select {
	case err := <-errChan:
		if err != nil && err != http.ErrServerClosed {
			logger.Fatalf("error: %s", err.Error())
		}
	case <-stopChan:
		logger.Warnf("will stop %s grpc server gracefully...", conf.Name)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(ctx); err != nil {
		logger.Fatalf("server Shutdown Failed: %s", err.Error())
	}
	logger.Warnf("%s grpc server stopped", conf.Name)
}

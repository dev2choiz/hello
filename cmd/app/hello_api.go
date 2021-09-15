package app

import (
	"github.com/dev2choiz/hello/internal/server"
	"github.com/dev2choiz/hello/pkg/handlers"
	"github.com/dev2choiz/hello/pkg/logger"
	"github.com/dev2choiz/hello/pkg/protobuf/healthpb"
	"github.com/dev2choiz/hello/pkg/protobuf/notifypb"
	"github.com/dev2choiz/hello/pkg/protobuf/sandboxpb"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var helloApiCmd = &cobra.Command{
	Use:   "hello-api",
	Short: "Run hello-api server",
	Long: "Run hello-api server",
	Run: func(cmd *cobra.Command, args []string) {
		executeHelloApi()
	},
}

func init() {
	helloApiCmd.PersistentFlags().StringVar(&server.RunConfig.Port, "port", "", "grpc port")
	helloApiCmd.PersistentFlags().StringVar(&server.RunConfig.Name, "name", "", "service name")
}

func executeHelloApi() {
	conf := server.GetConfig()
	conf.Name = "hello-api"
	logger.Infof("config name=%s port=%s", conf.Name, conf.Port)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(server.LogInterceptor),
		grpc.StreamInterceptor(server.LogStreamInterceptor),
	)
	healthpb.RegisterHealthServer(grpcServer, &handlers.HealthServer{})
	notifypb.RegisterNotifyServer(grpcServer, &handlers.NotifyServer{})
	sandboxpb.RegisterSandboxServer(grpcServer, &handlers.SandboxServer{})

	server.RunGrpcServer(grpcServer, conf)
}

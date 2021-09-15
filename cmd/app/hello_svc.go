package app

import (
	"github.com/dev2choiz/hello/internal/server"
	"github.com/dev2choiz/hello/pkg/handlers"
	"github.com/dev2choiz/hello/pkg/logger"
	"github.com/dev2choiz/hello/pkg/protobuf/pingpb"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var helloSvcCmd = &cobra.Command{
	Use:   "hello-svc",
	Short: "Run hello-svc server",
	Long: "Run hello-svc server",
	Run: func(cmd *cobra.Command, args []string) {
		executeHelloSvc()
	},
}

func init() {
	helloSvcCmd.PersistentFlags().StringVar(&server.RunConfig.Port, "port", "", "grpc port")
	helloSvcCmd.PersistentFlags().StringVar(&server.RunConfig.Name, "name", "", "service name")
}

func executeHelloSvc() {
	conf := server.GetConfig()
	logger.Infof("config name=%s port=%s", conf.Name, conf.Port)

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(server.LogInterceptor))
	pingpb.RegisterPingServer(grpcServer, &handlers.PingServer{SvcName: conf.Name})

	server.RunGrpcServer(grpcServer, conf)
}

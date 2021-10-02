package app

import (
	"github.com/dev2choiz/hello/internal/server"
	"github.com/dev2choiz/hello/pkg/app_wire"
	"github.com/dev2choiz/hello/pkg/config"
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
	helloSvcCmd.PersistentFlags().StringVar(&config.Conf.Port, "port", "", "grpc port")
	helloSvcCmd.PersistentFlags().StringVar(&config.Conf.Name, "name", "", "service name")
}

// executeHelloSvc is the entrypoint for the generic micro-services
func executeHelloSvc() {
	// initialize dependencies
	app_wire.InitializeLogger()
	conf := config.Conf
	logger.Infof("config name=%s port=%s", conf.Name, conf.Port)

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(server.LogInterceptor))
	pingpb.RegisterPingServer(grpcServer, &handlers.PingServer{SvcName: conf.Name})

	server.RunGrpcServer(grpcServer, conf)
}

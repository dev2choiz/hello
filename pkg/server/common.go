package server

import (
	"flag"
	"github.com/dev2choiz/hello/internal/logger"
	"go.uber.org/zap"
)

type Config struct {
	Name string
	Mode string
	GrpcPort string
	HttpPort string
}

var RunConfig = &Config{}

func init() {
	flag.StringVar(&RunConfig.Name, "name", "", "svc name")
	flag.StringVar(&RunConfig.Mode, "mode", "", "grpc or http")
	flag.StringVar(&RunConfig.GrpcPort, "grpc-port", "", "grpc port")
	flag.StringVar(&RunConfig.HttpPort, "http-port", "", "http port")
	flag.Parse()

	logger.Info("config",
		zap.String("name", RunConfig.Name),
		zap.String("mode", RunConfig.Mode),
		zap.String("grpc-port", RunConfig.GrpcPort),
		zap.String("http-port", RunConfig.HttpPort))
}

func GetConfig() *Config {
	return RunConfig
}
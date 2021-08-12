package server

import (
	"flag"
	"github.com/dev2choiz/hello/internal/logger"
	"go.uber.org/zap"
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

	logger.Info("config",
		zap.String("name", RunConfig.Name),
		zap.String("grpc-port", RunConfig.Port))
}

func GetConfig() *Config {
	return RunConfig
}
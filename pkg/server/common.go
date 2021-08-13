package server

import (
	"flag"
	"github.com/dev2choiz/hello/internal/logger"
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
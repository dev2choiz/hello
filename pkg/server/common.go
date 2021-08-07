package server

import (
	"flag"
	"log"
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
	log.Println("name", RunConfig.Name)
	log.Println("mode", RunConfig.Mode)
	log.Println("grpc-port", RunConfig.GrpcPort)
	log.Println("http-port", RunConfig.HttpPort)
}

func GetConfig() *Config {
	return RunConfig
}
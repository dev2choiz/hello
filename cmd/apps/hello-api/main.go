package main

import (
	"github.com/dev2choiz/hello/pkg/server"
)

func main() {
	conf := server.GetConfig()
	if conf.Mode == "http" {
		executeApiHttp(conf)
	} else {
		executeApiGrpc(conf)
	}
}

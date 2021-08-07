package main

import (
	"github.com/dev2choiz/hello/pkg/server"
)

func main() {
	conf := server.GetConfig()
	if conf.Mode == "http" {
		server.ExecuteSvcHttp(conf)
	} else {
		server.ExecuteSvcGrpc(conf)
	}
}

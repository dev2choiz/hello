package main

import (
	"github.com/dev2choiz/hello/cmd/app"
	"github.com/dev2choiz/hello/pkg/logger"
)

func main() {
	if err := app.Execute(); err != nil {
		logger.Fatalf(err.Error())
	}
}
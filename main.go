package main

import (
	"github.com/dev2choiz/hello/cmd/app"
	"github.com/dev2choiz/hello/pkg/logger"
)

func main() {
	// run the application using cobra
	if err := app.Execute(); err != nil {
		logger.Fatalf(err.Error())
	}
}
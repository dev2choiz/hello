package main

import (
	"github.com/dev2choiz/hello/pkg/app"
)

func main() {
	err := app.Execute()
	if err != nil {
		panic(err)
	}
}

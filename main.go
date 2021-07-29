package main

import (
	"github.com/dev2choiz/hello/pkg/app"
)

func main() {
	err := app.ExecuteGin()
	if err != nil {
		panic(err)
	}
}

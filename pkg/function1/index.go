package function1

import (
	"fmt"
	"github.com/dev2choiz/hello/pkg/version"
	"time"
)

func DoFunction1(name string) error {
	sec := 20
	fmt.Println("Hello ", name, ".\nVersion = ", version.Get())
	<-time.After(time.Duration(sec) * time.Second)
	fmt.Println("Log after waiting ", sec, " seconds")

	return nil
}

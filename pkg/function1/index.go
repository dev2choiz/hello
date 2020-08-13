package function1

import (
	"fmt"
	"github.com/dev2choiz/hello/pkg/version"
)

func DoFunction1() error {
	fmt.Println("Function1 executed. version = ", version.Get())
	return nil
}

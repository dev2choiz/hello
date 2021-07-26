package helloworld

import (
	"fmt"
	"github.com/dev2choiz/hello/pkg/version"
)

func Say() string {
	return fmt.Sprintf(`version=%s`, version.Get())
}

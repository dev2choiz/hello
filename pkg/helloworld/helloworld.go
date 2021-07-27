package helloworld

import (
	"github.com/dev2choiz/hello/pkg/version"
)

func Say() map[string]string {
	return map[string]string{
		"version": version.Get(),
	}
}

package function1

import (
	"github.com/dev2choiz/hello/internal/logger"
	"github.com/dev2choiz/hello/pkg/version"
	"go.uber.org/zap"
	"time"
)

func DoFunction1(name string) error {
	sec := 20
	logger.Info("hello " + name, zap.String("version", version.Get()))
	<-time.After(time.Duration(sec) * time.Second)
	logger.Info("bye " + name)

	return nil
}

//go:generate wire
//+build wireinject

package wire

import (
	"github.com/dev2choiz/hello/pkg/logger"
	"github.com/google/wire"
)

func InitializeLogger() logger.LoggerInterface {
	wire.Build(logger.NewLogger)
	return &logger.Logger{}
}
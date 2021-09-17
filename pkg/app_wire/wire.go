//go:generate wire
//+build wireinject

package app_wire

import (
	"github.com/dev2choiz/hello/pkg/logger"
	"github.com/google/wire"
)

func InitializeLogger() logger.LoggerInterface {
	wire.Build(logger.NewLogger)
	return &logger.Logger{}
}

func InitializeLoggerMock() logger.LoggerInterface {
	wire.Build(logger.NewLoggerMock)
	return &logger.Logger{}
}
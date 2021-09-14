package logger

import (
	"go.uber.org/zap"
)

type LoggerMock struct {}

func NewLoggerTest() LoggerInterface {
	if instance == nil {
		instance = &LoggerMock{}
	}
	return instance
}

func (l *LoggerMock) Debug(msg string, fields ...zap.Field) {
	// do nothing
}

func (l *LoggerMock) Info(msg string, fields ...zap.Field) {
	// do nothing
}

func (l *LoggerMock) Warn(msg string, fields ...zap.Field) {
	// do nothing
}

func (l *LoggerMock) Error(msg string, fields ...zap.Field) {
	// do nothing
}

func (l *LoggerMock) Fatal(msg string, fields ...zap.Field) {
	// do nothing
}

func (l *LoggerMock) Debugf(msg string, args ...interface{}) {
	// do nothing
}

func (l *LoggerMock) Infof(msg string, args ...interface{}) {
	// do nothing
}

func (l *LoggerMock) Warnf(msg string, args ...interface{}) {
	// do nothing
}

func (l *LoggerMock) Errorf(msg string, args ...interface{}) {
	// do nothing
}

func (l *LoggerMock) Fatalf(msg string, args ...interface{}) {
	// do nothing
}

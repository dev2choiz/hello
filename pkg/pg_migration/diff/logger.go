package diff

import (
	"context"
	"gorm.io/gorm/logger"
	"time"
)

type migLogger struct {
	Default  logger.Interface
	LogLevel logger.LogLevel
	Sql      []string
}

func newMigLogger() *migLogger {
	return &migLogger{
		Default: logger.Default,
		LogLevel: logger.Info,
	}
}

func (l* migLogger) LogMode(lvl logger.LogLevel) logger.Interface {
	l.LogLevel = lvl
	return l
}

func (l* migLogger) Info(ctx context.Context, m string, d ...interface{}) {
	l.Default.LogMode(l.LogLevel).Info(ctx, m, d)
}

func (l* migLogger) Warn(ctx context.Context, m string, d ...interface{}) {
	l.Default.LogMode(l.LogLevel).Warn(ctx, m, d)
}

func (l* migLogger) Error(ctx context.Context, m string, d ...interface{}) {
	l.Default.LogMode(l.LogLevel).Error(ctx, m, d)
}

func (l* migLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	l.Default.LogMode(l.LogLevel).Trace(ctx, begin, fc, err)
	if err != nil {
		return
	}

	sql, rows := fc()
	if rows == -1 {
		return
	}

	l.Sql = append(l.Sql, sql)
}

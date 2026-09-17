package logger

import (
	"context"
	"errors"
	"time"

	"go.uber.org/zap"
	gormlogger "gorm.io/gorm/logger"
)

type GormLogConfig struct {
	SlowThreshold        time.Duration
	IgnoreRecordNotFound bool
}

type GormZapLogger struct {
	zap    *zap.Logger
	config GormLogConfig
}

func NewGormZapLogger(zapLogger *zap.Logger, config GormLogConfig) *GormZapLogger {
	return &GormZapLogger{zap: zapLogger, config: config}
}

func (l *GormZapLogger) LogMode(gormlogger.LogLevel) gormlogger.Interface {
	return l
}

func (l *GormZapLogger) Info(_ context.Context, msg string, args ...interface{}) {
	l.zap.Sugar().Infof(msg, args...)
}

func (l *GormZapLogger) Warn(_ context.Context, msg string, args ...interface{}) {
	l.zap.Sugar().Warnf(msg, args...)
}

func (l *GormZapLogger) Error(_ context.Context, msg string, args ...interface{}) {
	l.zap.Sugar().Errorf(msg, args...)
}

func (l *GormZapLogger) Trace(_ context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()

	if err != nil && !(l.config.IgnoreRecordNotFound && errors.Is(err, gormlogger.ErrRecordNotFound)) {
		l.zap.Error("gorm query error", zap.Error(err), zap.String("sql", sql), zap.Int64("rows", rows), zap.Duration("elapsed", elapsed))
		return
	}

	if l.config.SlowThreshold != 0 && elapsed > l.config.SlowThreshold {
		l.zap.Warn("gorm slow query", zap.String("sql", sql), zap.Int64("rows", rows), zap.Duration("elapsed", elapsed))
	}
}

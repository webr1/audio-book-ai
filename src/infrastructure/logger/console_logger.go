package logger

import "go.uber.org/zap"

func NewConsoleLogger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	return logger
}

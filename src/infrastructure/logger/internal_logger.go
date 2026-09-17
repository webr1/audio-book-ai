package logger

import "go.uber.org/zap"

// InternalLogger is a package-level singleton for use from anywhere that
// doesn't have DI-injected logger access (e.g. infrastructure/errors).
var InternalLogger *zap.Logger = NewConsoleLogger()

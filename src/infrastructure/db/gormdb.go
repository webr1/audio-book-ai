package db

import (
	"fmt"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"audio-book-ai/src/infrastructure/env"
	"audio-book-ai/src/infrastructure/logger"
)

// @inject
func NewGormDB(e *env.Env, zapLogger *zap.Logger) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		e.DBHost, e.DBUser, e.DBPassword, e.DBName, e.DBPort, e.DBSSLMode,
	)

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.NewGormZapLogger(zapLogger, logger.GormLogConfig{
			SlowThreshold:        250 * time.Millisecond,
			IgnoreRecordNotFound: true,
		}),
	})
	if err != nil {
		panic(err)
	}

	return gormDB
}

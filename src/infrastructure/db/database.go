package db

import "gorm.io/gorm"

type Database struct {
	DB *gorm.DB
}

// @inject
func NewDatabase(gormDB *gorm.DB) *Database {
	return &Database{DB: gormDB}
}

package repository

import (
	"gorm.io/gorm"

	"audio-book-ai/src/infrastructure/db"
)

type BaseRepository struct {
	database *db.Database
}

// @inject
func NewBaseRepository(database *db.Database) *BaseRepository {
	return &BaseRepository{database: database}
}

func (this *BaseRepository) db() *gorm.DB {
	return this.database.DB
}

package repository

import (
	"context"

	"audio-book-ai/src/core/domain/entity"
	"audio-book-ai/src/core/domain/entity/enum"
)

type BookRepository interface {
	Create(ctx context.Context, book *entity.BookEntity) error
	Get(ctx context.Context, id uint) (*entity.BookEntity, error)
	UpdateStatus(ctx context.Context, id uint, status enum.BookStatus, errMsg string) error
	UpdateProgress(ctx context.Context, id uint, chunksDone, chunksTotal int) error
	SetCounts(ctx context.Context, id uint, wordCount, charCount, chunksTotal int) error
	SetOutputPath(ctx context.Context, id uint, path string) error
}

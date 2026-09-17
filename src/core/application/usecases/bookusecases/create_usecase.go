package bookusecases

import (
	"context"
	"encoding/json"
	"io"
	"path/filepath"

	"audio-book-ai/src/core/domain/entity"
	"audio-book-ai/src/core/domain/ports/repository"
	"audio-book-ai/src/core/domain/ports/storageport"
	"audio-book-ai/src/core/domain/ports/taskport"
)

const TaskTypeProcessBook = "book:process"

type ProcessBookPayload struct {
	BookID uint `json:"book_id"`
}

type BookCreateUseCase struct {
	repository repository.BookRepository
	store      storageport.Store
	enqueuer   taskport.Enqueuer
}

// @inject
func NewBookCreateUseCase(repository repository.BookRepository, store storageport.Store, enqueuer taskport.Enqueuer) *BookCreateUseCase {
	return &BookCreateUseCase{repository: repository, store: store, enqueuer: enqueuer}
}

func (this *BookCreateUseCase) Invoke(ctx context.Context, userID uint, filename string, content io.Reader) (*entity.BookEntity, error) {
	book := entity.NewBookEntity(userID, filename)
	if err := this.repository.Create(ctx, book); err != nil {
		return nil, err
	}

	ext := filepath.Ext(filename)
	if _, err := this.store.SaveUpload(ctx, book.ID, ext, content); err != nil {
		return nil, err
	}

	payload, err := json.Marshal(ProcessBookPayload{BookID: book.ID})
	if err != nil {
		return nil, err
	}
	if err := this.enqueuer.Enqueue(ctx, TaskTypeProcessBook, payload); err != nil {
		return nil, err
	}

	return book, nil
}

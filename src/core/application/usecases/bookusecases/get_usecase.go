package bookusecases

import (
	"context"

	"audio-book-ai/src/core/application/response"
	"audio-book-ai/src/core/domain/entity"
	"audio-book-ai/src/core/domain/ports/repository"
)

type BookGetUseCase struct {
	repository repository.BookRepository
}

// @inject
func NewBookGetUseCase(repository repository.BookRepository) *BookGetUseCase {
	return &BookGetUseCase{repository: repository}
}

// Invoke returns response.NotFoundError for a book owned by someone else,
// same as for a missing id, so we never reveal that a book with that id exists.
func (this *BookGetUseCase) Invoke(ctx context.Context, id uint, requestingUserID uint) (*entity.BookEntity, error) {
	book, err := this.repository.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	if book.UserID != requestingUserID {
		return nil, response.NotFoundError
	}

	return book, nil
}

package bookusecases

import (
	"context"
	"io"

	"audio-book-ai/src/core/domain/ports/repository"
	"audio-book-ai/src/core/domain/ports/storageport"
)

type UploadReceiptUseCase struct {
	repository repository.BookRepository
	store      storageport.Store
}

// @inject
func NewUploadReceiptUseCase(repository repository.BookRepository, store storageport.Store) *UploadReceiptUseCase {
	return &UploadReceiptUseCase{repository: repository, store: store}
}

func (this *UploadReceiptUseCase) Invoke(ctx context.Context, bookID uint, ext string, content io.Reader) error {
	path, err := this.store.SaveReceipt(ctx, bookID, ext, content)
	if err != nil {
		return err
	}
	return this.repository.SetPaymentReceipt(ctx, bookID, path)
}

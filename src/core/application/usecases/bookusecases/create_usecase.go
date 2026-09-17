package bookusecases

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"audio-book-ai/src/core/domain/entity"
	"audio-book-ai/src/core/domain/ports/ingestport"
	"audio-book-ai/src/core/domain/ports/repository"
	"audio-book-ai/src/core/domain/ports/storageport"
	"audio-book-ai/src/core/domain/ports/taskport"
	"audio-book-ai/src/infrastructure/env"
)

const TaskTypeProcessBook = "book:process"

type ProcessBookPayload struct {
	BookID uint `json:"book_id"`
}

type BookCreateUseCase struct {
	repository   repository.BookRepository
	store        storageport.Store
	extractor    ingestport.Extractor
	enqueuer     taskport.Enqueuer
	pricePerChar float64
}

// @inject
func NewBookCreateUseCase(repository repository.BookRepository, store storageport.Store, extractor ingestport.Extractor, enqueuer taskport.Enqueuer, e *env.Env) *BookCreateUseCase {
	return &BookCreateUseCase{repository: repository, store: store, extractor: extractor, enqueuer: enqueuer, pricePerChar: e.PricePerChar}
}

// Invoke counts words/chars synchronously (cheap — plain text parsing) so
// the caller sees them immediately, without waiting for the async TTS
// pipeline to reach that step.
func (this *BookCreateUseCase) Invoke(ctx context.Context, userID uint, filename string, content io.Reader) (*entity.BookEntity, error) {
	data, err := io.ReadAll(content)
	if err != nil {
		return nil, fmt.Errorf("read uploaded file: %w", err)
	}

	text, err := this.extractor.Extract(ctx, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	wordCount := len(strings.Fields(text))
	charCount := utf8.RuneCountInString(text)
	price := float64(charCount) * this.pricePerChar

	book := entity.NewBookEntity(userID, filename)
	if err := this.repository.Create(ctx, book); err != nil {
		return nil, err
	}
	if err := this.repository.SetCounts(ctx, book.ID, wordCount, charCount, 0); err != nil {
		return nil, err
	}
	if err := this.repository.SetPrice(ctx, book.ID, price); err != nil {
		return nil, err
	}
	book.WordCount = wordCount
	book.CharCount = charCount
	book.Price = price

	ext := filepath.Ext(filename)
	if _, err := this.store.SaveUpload(ctx, book.ID, ext, bytes.NewReader(data)); err != nil {
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

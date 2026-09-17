package bookusecases

import (
	"context"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"audio-book-ai/src/core/domain/entity/enum"
	"audio-book-ai/src/core/domain/ports/audioport"
	"audio-book-ai/src/core/domain/ports/chunkport"
	"audio-book-ai/src/core/domain/ports/ingestport"
	"audio-book-ai/src/core/domain/ports/repository"
	"audio-book-ai/src/core/domain/ports/storageport"
	"audio-book-ai/src/core/domain/ports/ttsport"
)

type ProcessBookUseCase struct {
	repository    repository.BookRepository
	store         storageport.Store
	extractor     ingestport.Extractor
	chunker       chunkport.Chunker
	ttsProvider   ttsport.TTSProvider
	concatenator  audioport.Concatenator
	maxChunkChars int
}

// @inject
func NewProcessBookUseCase(
	repository repository.BookRepository,
	store storageport.Store,
	extractor ingestport.Extractor,
	chunker chunkport.Chunker,
	ttsProvider ttsport.TTSProvider,
	concatenator audioport.Concatenator,
) *ProcessBookUseCase {
	return &ProcessBookUseCase{
		repository:    repository,
		store:         store,
		extractor:     extractor,
		chunker:       chunker,
		ttsProvider:   ttsProvider,
		concatenator:  concatenator,
		maxChunkChars: 1000,
	}
}

func (this *ProcessBookUseCase) Invoke(ctx context.Context, bookID uint) error {
	book, err := this.repository.Get(ctx, bookID)
	if err != nil {
		return err
	}

	if err := this.repository.UpdateStatus(ctx, bookID, enum.BookStatusProcessing, ""); err != nil {
		return err
	}

	if err := this.run(ctx, bookID, book.Filename); err != nil {
		_ = this.repository.UpdateStatus(ctx, bookID, enum.BookStatusFailed, err.Error())
		return err
	}

	return this.repository.UpdateStatus(ctx, bookID, enum.BookStatusDone, "")
}

func (this *ProcessBookUseCase) run(ctx context.Context, bookID uint, filename string) error {
	ext := filepath.Ext(filename)

	source, err := this.store.OpenUpload(ctx, bookID, ext)
	if err != nil {
		return err
	}
	defer source.Close()

	text, err := this.extractor.Extract(ctx, source)
	if err != nil {
		return err
	}

	wordCount := len(strings.Fields(text))
	charCount := utf8.RuneCountInString(text)

	chunks := this.chunker.Chunk(text, chunkport.Options{MaxChars: this.maxChunkChars})
	if err := this.repository.SetCounts(ctx, bookID, wordCount, charCount, len(chunks)); err != nil {
		return err
	}

	for i, chunk := range chunks {
		audio, err := this.ttsProvider.Synthesize(ctx, chunk, ttsport.VoiceOptions{})
		if err != nil {
			return err
		}
		if _, err := this.store.WriteChunk(ctx, bookID, i, audio); err != nil {
			return err
		}
		if err := this.repository.UpdateProgress(ctx, bookID, i+1, len(chunks)); err != nil {
			return err
		}
	}

	chunkPaths := this.store.ChunkPaths(bookID, len(chunks))
	outputPath := this.store.OutputPath(bookID)
	if err := this.concatenator.Concatenate(chunkPaths, outputPath); err != nil {
		return err
	}

	return this.repository.SetOutputPath(ctx, bookID, outputPath)
}

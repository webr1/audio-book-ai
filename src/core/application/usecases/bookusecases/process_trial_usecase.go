package bookusecases

import (
	"context"
	"path/filepath"
	"strings"

	"audio-book-ai/src/core/domain/entity/enum"
	"audio-book-ai/src/core/domain/ports/audioport"
	"audio-book-ai/src/core/domain/ports/chunkport"
	"audio-book-ai/src/core/domain/ports/ingestport"
	"audio-book-ai/src/core/domain/ports/repository"
	"audio-book-ai/src/core/domain/ports/storageport"
	"audio-book-ai/src/core/domain/ports/ttsport"
	"audio-book-ai/src/infrastructure/env"
)

// ProcessTrialUseCase synthesizes only the first TrialChars characters of
// the book — a free sample so the user can hear the voice quality before
// paying for the full book. It reuses the exact same chunker/TTS/concat
// pipeline as ProcessBookUseCase, just on a truncated slice of the text
// and writing to the Trial* storage paths/DB columns instead.
type ProcessTrialUseCase struct {
	repository    repository.BookRepository
	store         storageport.Store
	extractor     ingestport.Extractor
	chunker       chunkport.Chunker
	ttsProvider   ttsport.TTSProvider
	concatenator  audioport.Concatenator
	maxChunkChars int
	trialChars    int
}

// @inject
func NewProcessTrialUseCase(
	repository repository.BookRepository,
	store storageport.Store,
	extractor ingestport.Extractor,
	chunker chunkport.Chunker,
	ttsProvider ttsport.TTSProvider,
	concatenator audioport.Concatenator,
	e *env.Env,
) *ProcessTrialUseCase {
	return &ProcessTrialUseCase{
		repository:    repository,
		store:         store,
		extractor:     extractor,
		chunker:       chunker,
		ttsProvider:   ttsProvider,
		concatenator:  concatenator,
		maxChunkChars: e.MaxChunkChars,
		trialChars:    e.TrialChars,
	}
}

func (this *ProcessTrialUseCase) Invoke(ctx context.Context, bookID uint) error {
	book, err := this.repository.Get(ctx, bookID)
	if err != nil {
		return err
	}

	if err := this.repository.UpdateTrialStatus(ctx, bookID, enum.BookStatusProcessing, ""); err != nil {
		return err
	}

	if err := this.run(ctx, bookID, book.Filename); err != nil {
		_ = this.repository.UpdateTrialStatus(ctx, bookID, enum.BookStatusFailed, err.Error())
		return err
	}

	return this.repository.UpdateTrialStatus(ctx, bookID, enum.BookStatusDone, "")
}

func (this *ProcessTrialUseCase) run(ctx context.Context, bookID uint, filename string) error {
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
	text = truncateToChars(text, this.trialChars)

	chunks := this.chunker.Chunk(text, chunkport.Options{MaxChars: this.maxChunkChars})
	if err := this.repository.UpdateTrialProgress(ctx, bookID, 0, len(chunks)); err != nil {
		return err
	}

	for i, chunk := range chunks {
		audio, err := this.ttsProvider.Synthesize(ctx, chunk, ttsport.VoiceOptions{})
		if err != nil {
			return err
		}
		if _, err := this.store.WriteTrialChunk(ctx, bookID, i, audio); err != nil {
			return err
		}
		if err := this.repository.UpdateTrialProgress(ctx, bookID, i+1, len(chunks)); err != nil {
			return err
		}
	}

	chunkPaths := this.store.TrialChunkPaths(bookID, len(chunks))
	outputPath := this.store.TrialOutputPath(bookID)
	if err := this.concatenator.Concatenate(chunkPaths, outputPath); err != nil {
		return err
	}

	return this.repository.SetTrialOutputPath(ctx, bookID, outputPath)
}

// truncateToChars cuts text to at most maxChars runes, backing off to the
// last whitespace boundary so it never splits mid-word.
func truncateToChars(text string, maxChars int) string {
	runes := []rune(text)
	if len(runes) <= maxChars {
		return text
	}

	cut := string(runes[:maxChars])
	if idx := strings.LastIndexAny(cut, " \n\t"); idx > 0 {
		cut = cut[:idx]
	}
	return strings.TrimSpace(cut)
}

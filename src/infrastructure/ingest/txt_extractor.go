package ingest

import (
	"context"
	"fmt"
	"io"
	"unicode/utf8"

	"audio-book-ai/src/core/domain/ports/ingestport"
)

type TxtExtractor struct{}

// @inject
func NewTxtExtractor() ingestport.Extractor {
	return &TxtExtractor{}
}

func (e *TxtExtractor) Extract(_ context.Context, r io.Reader) (string, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return "", fmt.Errorf("read txt: %w", err)
	}
	if !utf8.Valid(data) {
		return "", fmt.Errorf("file is not valid UTF-8 text")
	}
	return string(data), nil
}

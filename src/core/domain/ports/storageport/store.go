package storageport

import (
	"context"
	"io"
)

// Store persists everything the pipeline reads/writes on disk: the
// uploaded source file, each synthesized chunk (full and trial), the
// payment receipt, and path lookups needed to build final outputs.
// Local disk is the only implementation today (src/infrastructure/storage);
// MinIO can be added later as a second implementation without touching
// any caller.
type Store interface {
	SaveUpload(ctx context.Context, bookID uint, ext string, r io.Reader) (path string, err error)
	OpenUpload(ctx context.Context, bookID uint, ext string) (io.ReadCloser, error)

	WriteChunk(ctx context.Context, bookID uint, idx int, data []byte) (path string, err error)
	ChunkPaths(bookID uint, count int) []string
	OutputPath(bookID uint) string

	WriteTrialChunk(ctx context.Context, bookID uint, idx int, data []byte) (path string, err error)
	TrialChunkPaths(bookID uint, count int) []string
	TrialOutputPath(bookID uint) string

	SaveReceipt(ctx context.Context, bookID uint, ext string, r io.Reader) (path string, err error)
}

package storage

import (
	"context"
	"fmt"
	"io"
	"os"

	"audio-book-ai/src/core/domain/ports/storageport"
	"audio-book-ai/src/infrastructure/env"
)

type FileStore struct {
	dataDir string
}

// @inject
func NewFileStore(e *env.Env) storageport.Store {
	return &FileStore{dataDir: e.DataDir}
}

func (s *FileStore) SaveUpload(_ context.Context, bookID uint, ext string, r io.Reader) (string, error) {
	if err := os.MkdirAll(UploadDir(s.dataDir, bookID), 0o755); err != nil {
		return "", fmt.Errorf("create upload dir: %w", err)
	}

	path := UploadPath(s.dataDir, bookID, ext)
	f, err := os.Create(path)
	if err != nil {
		return "", fmt.Errorf("create upload file: %w", err)
	}
	defer f.Close()

	if _, err := io.Copy(f, r); err != nil {
		return "", fmt.Errorf("write upload file: %w", err)
	}
	return path, nil
}

func (s *FileStore) OpenUpload(_ context.Context, bookID uint, ext string) (io.ReadCloser, error) {
	return os.Open(UploadPath(s.dataDir, bookID, ext))
}

func (s *FileStore) WriteChunk(_ context.Context, bookID uint, idx int, data []byte) (string, error) {
	if err := os.MkdirAll(ChunkDir(s.dataDir, bookID), 0o755); err != nil {
		return "", fmt.Errorf("create chunk dir: %w", err)
	}

	path := ChunkPath(s.dataDir, bookID, idx)
	if err := os.WriteFile(path, data, 0o644); err != nil {
		return "", fmt.Errorf("write chunk file: %w", err)
	}
	return path, nil
}

func (s *FileStore) ChunkPaths(bookID uint, count int) []string {
	paths := make([]string, count)
	for i := 0; i < count; i++ {
		paths[i] = ChunkPath(s.dataDir, bookID, i)
	}
	return paths
}

func (s *FileStore) OutputPath(bookID uint) string {
	if err := os.MkdirAll(OutputDir(s.dataDir, bookID), 0o755); err != nil {
		panic(err)
	}
	return OutputPath(s.dataDir, bookID)
}

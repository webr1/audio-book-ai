package storage

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func EnsureDataDirs(dataDir string) error {
	dirs := []string{
		filepath.Join(dataDir, "uploads"),
		filepath.Join(dataDir, "chunks"),
		filepath.Join(dataDir, "output"),
	}
	for _, d := range dirs {
		if err := os.MkdirAll(d, 0o755); err != nil {
			return fmt.Errorf("create dir %s: %w", d, err)
		}
	}
	return nil
}

func bookDir(dataDir string, bookID uint) string {
	return strconv.FormatUint(uint64(bookID), 10)
}

func UploadDir(dataDir string, bookID uint) string {
	return filepath.Join(dataDir, "uploads", bookDir(dataDir, bookID))
}

func UploadPath(dataDir string, bookID uint, ext string) string {
	return filepath.Join(UploadDir(dataDir, bookID), "source"+ext)
}

func ChunkDir(dataDir string, bookID uint) string {
	return filepath.Join(dataDir, "chunks", bookDir(dataDir, bookID))
}

func ChunkPath(dataDir string, bookID uint, idx int) string {
	return filepath.Join(ChunkDir(dataDir, bookID), fmt.Sprintf("%04d.wav", idx))
}

func OutputDir(dataDir string, bookID uint) string {
	return filepath.Join(dataDir, "output", bookDir(dataDir, bookID))
}

func OutputPath(dataDir string, bookID uint) string {
	return filepath.Join(OutputDir(dataDir, bookID), "audiobook.wav")
}

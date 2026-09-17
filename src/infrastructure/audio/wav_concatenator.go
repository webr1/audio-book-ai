package audio

import (
	"encoding/binary"
	"fmt"
	"os"

	"audio-book-ai/src/core/domain/ports/audioport"
)

const wavHeaderSize = 44

type WAVConcatenator struct{}

// @inject
func NewWAVConcatenator() audioport.Concatenator {
	return &WAVConcatenator{}
}

// Concatenate joins WAV chunk files that all share the same format
// (guaranteed by the mock provider today) into a single WAV file:
// one shared header, followed by every chunk's raw `data` payload back to back.
func (c *WAVConcatenator) Concatenate(chunkPaths []string, outputPath string) error {
	if len(chunkPaths) == 0 {
		return fmt.Errorf("no chunks to concatenate")
	}

	out, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("create output file: %w", err)
	}
	defer out.Close()

	// Placeholder header, patched with the real data size at the end.
	if _, err := out.Write(make([]byte, wavHeaderSize)); err != nil {
		return fmt.Errorf("write header placeholder: %w", err)
	}

	var format []byte
	var totalDataLen uint32

	for _, path := range chunkPaths {
		raw, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("read chunk %s: %w", path, err)
		}
		if len(raw) < wavHeaderSize {
			return fmt.Errorf("chunk %s is not a valid WAV file", path)
		}

		fmtChunk := raw[20:36] // audio format, channels, sample rate, byte rate, block align, bits per sample
		if format == nil {
			format = fmtChunk
		} else if string(format) != string(fmtChunk) {
			return fmt.Errorf("chunk %s has a different WAV format than earlier chunks", path)
		}

		data := raw[wavHeaderSize:]
		if _, err := out.Write(data); err != nil {
			return fmt.Errorf("write chunk %s data: %w", path, err)
		}
		totalDataLen += uint32(len(data))
	}

	if _, err := out.Seek(0, 0); err != nil {
		return fmt.Errorf("seek to header: %w", err)
	}
	if err := writeWAVHeader(out, format, totalDataLen); err != nil {
		return fmt.Errorf("write final header: %w", err)
	}

	return nil
}

func writeWAVHeader(w *os.File, format []byte, dataLen uint32) error {
	buf := make([]byte, 0, wavHeaderSize)

	buf = append(buf, []byte("RIFF")...)
	buf = binary.LittleEndian.AppendUint32(buf, 36+dataLen)
	buf = append(buf, []byte("WAVE")...)

	buf = append(buf, []byte("fmt ")...)
	buf = binary.LittleEndian.AppendUint32(buf, 16)
	buf = append(buf, format...)

	buf = append(buf, []byte("data")...)
	buf = binary.LittleEndian.AppendUint32(buf, dataLen)

	_, err := w.Write(buf)
	return err
}

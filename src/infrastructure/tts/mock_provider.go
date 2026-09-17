package tts

import (
	"bytes"
	"context"
	"encoding/binary"
	"math"
	"time"

	"audio-book-ai/src/core/domain/ports/ttsport"
)

const (
	sampleRate    = 22050
	charsPerSec   = 15.0
	minDuration   = 300 * time.Millisecond
	beepDuration  = 200 * time.Millisecond
	beepFrequency = 440.0
)

// MockProvider is a stand-in for a real TTS engine (e.g. the future
// uzbekvoice.ai adapter). It produces a real, valid, playable mono
// 16-bit PCM WAV file per chunk — silence sized to the chunk's text
// length, with a short beep marking the start of each chunk so the
// concatenated output is audibly checkable during manual testing.
type MockProvider struct{}

// @inject
func NewMockProvider() ttsport.TTSProvider {
	return &MockProvider{}
}

func (p *MockProvider) Format() string {
	return "wav"
}

func (p *MockProvider) Synthesize(ctx context.Context, text string, _ ttsport.VoiceOptions) ([]byte, error) {
	duration := time.Duration(float64(len(text)) / charsPerSec * float64(time.Second))
	if duration < minDuration {
		duration = minDuration
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-time.After(100 * time.Millisecond):
	}

	samples := generateSamples(duration)
	return encodeWAV(samples), nil
}

func generateSamples(duration time.Duration) []int16 {
	total := int(duration.Seconds() * sampleRate)
	beepSamples := int(beepDuration.Seconds() * sampleRate)
	if beepSamples > total {
		beepSamples = total
	}

	samples := make([]int16, total)
	for i := 0; i < beepSamples; i++ {
		t := float64(i) / sampleRate
		samples[i] = int16(8000 * math.Sin(2*math.Pi*beepFrequency*t))
	}
	// remaining samples stay zero (silence)
	return samples
}

func encodeWAV(samples []int16) []byte {
	dataLen := len(samples) * 2
	buf := new(bytes.Buffer)

	buf.WriteString("RIFF")
	binary.Write(buf, binary.LittleEndian, uint32(36+dataLen))
	buf.WriteString("WAVE")

	buf.WriteString("fmt ")
	binary.Write(buf, binary.LittleEndian, uint32(16))
	binary.Write(buf, binary.LittleEndian, uint16(1)) // PCM
	binary.Write(buf, binary.LittleEndian, uint16(1)) // mono
	binary.Write(buf, binary.LittleEndian, uint32(sampleRate))
	binary.Write(buf, binary.LittleEndian, uint32(sampleRate*2)) // byte rate
	binary.Write(buf, binary.LittleEndian, uint16(2))            // block align
	binary.Write(buf, binary.LittleEndian, uint16(16))           // bits per sample

	buf.WriteString("data")
	binary.Write(buf, binary.LittleEndian, uint32(dataLen))
	binary.Write(buf, binary.LittleEndian, samples)

	return buf.Bytes()
}

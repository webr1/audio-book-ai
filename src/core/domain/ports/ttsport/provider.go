package ttsport

import "context"

type VoiceOptions struct {
	VoiceID string
	Speed   float64
	Lang    string
}

// TTSProvider synthesizes speech for a single chunk of text.
// The uzbekvoice.ai adapter (once its API is documented) will be a second
// implementation of this interface — nothing else in the pipeline changes.
type TTSProvider interface {
	Synthesize(ctx context.Context, text string, opts VoiceOptions) ([]byte, error)
	Format() string
}

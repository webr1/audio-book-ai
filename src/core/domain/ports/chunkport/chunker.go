package chunkport

type Options struct {
	MaxChars int
}

// Chunker splits text into ordered pieces, each <= MaxChars, preferring
// paragraph then sentence boundaries, never splitting mid-word.
type Chunker interface {
	Chunk(text string, opts Options) []string
}

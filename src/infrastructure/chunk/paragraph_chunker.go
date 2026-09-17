package chunk

import (
	"regexp"
	"strings"

	"audio-book-ai/src/core/domain/ports/chunkport"
)

var (
	paragraphSplitRe = regexp.MustCompile(`\n\s*\n`)
	sentenceSplitRe  = regexp.MustCompile(`([.!?])\s+`)
)

type ParagraphChunker struct{}

// @inject
func NewParagraphChunker() chunkport.Chunker {
	return &ParagraphChunker{}
}

func (c *ParagraphChunker) Chunk(text string, opts chunkport.Options) []string {
	maxChars := opts.MaxChars
	if maxChars <= 0 {
		maxChars = 1000
	}

	var chunks []string
	var current strings.Builder

	flush := func() {
		if s := strings.TrimSpace(current.String()); s != "" {
			chunks = append(chunks, s)
		}
		current.Reset()
	}

	for _, paragraph := range paragraphSplitRe.Split(text, -1) {
		paragraph = strings.TrimSpace(paragraph)
		if paragraph == "" {
			continue
		}

		for _, piece := range splitToFit(paragraph, maxChars) {
			if current.Len() > 0 && current.Len()+1+len(piece) > maxChars {
				flush()
			}
			if current.Len() > 0 {
				current.WriteString(" ")
			}
			current.WriteString(piece)
		}
	}
	flush()

	return chunks
}

// splitToFit breaks a paragraph into pieces <= maxChars, preferring
// sentence boundaries, falling back to a hard split on whitespace.
func splitToFit(paragraph string, maxChars int) []string {
	if len(paragraph) <= maxChars {
		return []string{paragraph}
	}

	var pieces []string
	for _, sentence := range splitSentences(paragraph) {
		if len(sentence) <= maxChars {
			pieces = append(pieces, sentence)
			continue
		}
		pieces = append(pieces, hardSplit(sentence, maxChars)...)
	}
	return pieces
}

func splitSentences(text string) []string {
	indices := sentenceSplitRe.FindAllStringIndex(text, -1)
	if len(indices) == 0 {
		return []string{text}
	}

	var sentences []string
	start := 0
	for _, idx := range indices {
		end := idx[1]
		sentences = append(sentences, strings.TrimSpace(text[start:end]))
		start = end
	}
	if start < len(text) {
		sentences = append(sentences, strings.TrimSpace(text[start:]))
	}
	return sentences
}

// hardSplit breaks text on the last whitespace before maxChars, never mid-word.
func hardSplit(text string, maxChars int) []string {
	var pieces []string
	for len(text) > maxChars {
		cut := strings.LastIndex(text[:maxChars], " ")
		if cut <= 0 {
			cut = maxChars
		}
		pieces = append(pieces, strings.TrimSpace(text[:cut]))
		text = strings.TrimSpace(text[cut:])
	}
	if text != "" {
		pieces = append(pieces, text)
	}
	return pieces
}

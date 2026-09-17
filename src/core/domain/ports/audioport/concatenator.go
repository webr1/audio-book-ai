package audioport

// Concatenator joins ordered chunk audio files into a single output file.
type Concatenator interface {
	Concatenate(chunkPaths []string, outputPath string) error
}

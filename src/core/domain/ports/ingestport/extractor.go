package ingestport

import (
	"context"
	"io"
)

// Extractor turns a raw uploaded file into plain text.
// Registered per file extension; adding .epub/.pdf later means adding
// another implementation, not touching this interface or its callers.
type Extractor interface {
	Extract(ctx context.Context, r io.Reader) (string, error)
}

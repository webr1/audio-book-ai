package httpport

import "audio-book-ai/src/core/domain/ports/httpport/ctx"

type Handler func(c ctx.Context) error

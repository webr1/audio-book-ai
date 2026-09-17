package wire

import (
	"github.com/google/wire"

	"audio-book-ai/src/entrypoint/asynctask"
	"audio-book-ai/src/entrypoint/http"
)

func InitHttpApp() *http.App {
	wire.Build(ProviderSet)
	return nil
}

func InitAsyncApp() *asynctask.App {
	wire.Build(ProviderSet)
	return nil
}

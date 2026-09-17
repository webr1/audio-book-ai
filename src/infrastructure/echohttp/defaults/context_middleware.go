package defaults

import (
	"github.com/labstack/echo/v4"

	echoContext "audio-book-ai/src/infrastructure/echohttp/context"
)

// ContextMiddleware must run before any other middleware/handler: it
// replaces echo.Context with *echoContext.EchoContext, whose dynamic type
// satisfies both echo.Context and ctx.Context for the rest of the chain.
func ContextMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		wrapped := echoContext.NewEchoContext(c)
		return next(wrapped.(echo.Context))
	}
}

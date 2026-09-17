package mapper

import (
	"github.com/labstack/echo/v4"

	"audio-book-ai/src/core/domain/ports/httpport"
	"audio-book-ai/src/core/domain/ports/httpport/ctx"
)

// ToEchoHandler assumes defaults.ContextMiddleware already ran, replacing
// echo.Context's dynamic type with *context.EchoContext, which implements
// ctx.Context — so this assertion always succeeds downstream of it.
func ToEchoHandler(h httpport.Handler) echo.HandlerFunc {
	return func(c echo.Context) error {
		return h(c.(ctx.Context))
	}
}

func ToEchoMiddleware(mw httpport.Middleware) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			handler := func(hc ctx.Context) error {
				return next(hc.(echo.Context))
			}
			return mw(handler)(c.(ctx.Context))
		}
	}
}

func ToEchoMiddlewares(mws []httpport.Middleware) []echo.MiddlewareFunc {
	out := make([]echo.MiddlewareFunc, len(mws))
	for i, mw := range mws {
		out[i] = ToEchoMiddleware(mw)
	}
	return out
}

package permissions

import (
	"audio-book-ai/src/core/application/response"
	"audio-book-ai/src/core/domain/ports/httpport"
	"audio-book-ai/src/core/domain/ports/httpport/ctx"
)

func UserAuthenticatedPermission(next httpport.Handler) httpport.Handler {
	return func(c ctx.Context) error {
		if c.User() == nil {
			return response.UnauthorizedError
		}
		return next(c)
	}
}

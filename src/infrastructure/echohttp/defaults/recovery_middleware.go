package defaults

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/labstack/echo/v4"

	"audio-book-ai/src/core/application/response"
	"audio-book-ai/src/infrastructure/logger"
)

// RecoveryMiddleware turns a panicked *response.Response (e.g. from
// ctx.GetBody) into the matching JSON error response, and anything else
// into a 500, instead of letting Echo's default panic handling take over.
func RecoveryMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		defer func() {
			if r := recover(); r != nil {
				debug.PrintStack()

				if resp, ok := r.(*response.Response); ok {
					_ = c.JSON(resp.Status, resp)
					return
				}

				logger.InternalLogger.Error(fmt.Sprintf("panic: %v", r))
				_ = c.JSON(http.StatusInternalServerError, map[string]any{
					"status":  http.StatusInternalServerError,
					"message": fmt.Sprintf("%v", r),
				})
			}
		}()
		return next(c)
	}
}

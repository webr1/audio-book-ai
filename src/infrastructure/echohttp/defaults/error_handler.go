package defaults

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"audio-book-ai/src/core/application/response"
)

// HTTPErrorHandler renders a *response.Response returned (not panicked)
// by a handler in the same JSON shape as the success path, and falls
// back to a generic 500 for anything else.
func HTTPErrorHandler(err error, c echo.Context) {
	if c.Response().Committed {
		return
	}

	if resp, ok := err.(*response.Response); ok {
		_ = c.JSON(resp.Status, resp)
		return
	}

	if httpErr, ok := err.(*echo.HTTPError); ok {
		_ = c.JSON(httpErr.Code, map[string]any{"status": httpErr.Code, "message": httpErr.Message})
		return
	}

	_ = c.JSON(http.StatusInternalServerError, map[string]any{
		"status":  http.StatusInternalServerError,
		"message": err.Error(),
	})
}

package auth

import (
	"net/http"

	"audio-book-ai/src/core/domain/ports/httpport/ctx"
)

type MeHandler struct{}

// @inject
func NewMeHandler() *MeHandler {
	return &MeHandler{}
}

// Handle godoc
// @Tags		auth
// @Produce		json
// @Security	BearerAuth
// @Success		200	{object}	map[string]any
// @Router		/auth/me [get]
func (this *MeHandler) Handle(c ctx.Context) error {
	user := c.User()

	return c.JsonResponse(http.StatusOK, map[string]any{
		"id":      user.ID,
		"email":   user.Email,
		"name":    user.Name,
		"picture": user.Picture,
	})
}

package auth

import (
	"net/http"

	"audio-book-ai/src/core/application/response"
	"audio-book-ai/src/core/application/usecases/authusecases"
	"audio-book-ai/src/core/domain/ports/httpport/ctx"
)

type GoogleLoginRequest struct {
	IDToken string `json:"id_token" validate:"required"`
}

type GoogleLoginHandler struct {
	uc *authusecases.GoogleLoginUseCase
}

// @inject
func NewGoogleLoginHandler(uc *authusecases.GoogleLoginUseCase) *GoogleLoginHandler {
	return &GoogleLoginHandler{uc: uc}
}

// Handle godoc
// @Tags		auth
// @Accept		json
// @Produce		json
// @Param		body	body	GoogleLoginRequest	true	"Google ID token"
// @Success		200	{object}	map[string]any
// @Router		/auth/google [post]
func (this *GoogleLoginHandler) Handle(c ctx.Context) error {
	var body GoogleLoginRequest
	if err := c.Bind(&body); err != nil || body.IDToken == "" {
		return response.NewFailResponse(http.StatusBadRequest, "id_token is required")
	}

	access, refresh, user, err := this.uc.Invoke(c.GetContext(), body.IDToken)
	if err != nil {
		return err
	}

	return c.JsonResponse(http.StatusOK, map[string]any{
		"access_token":  access,
		"refresh_token": refresh,
		"user": map[string]any{
			"id":      user.ID,
			"email":   user.Email,
			"name":    user.Name,
			"picture": user.Picture,
		},
	})
}

package auth

import (
	"net/http"

	"audio-book-ai/src/core/application/response"
	"audio-book-ai/src/core/application/usecases/authusecases"
	"audio-book-ai/src/core/domain/ports/httpport/ctx"
)

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" validate:"required"`
}

type RefreshTokenHandler struct {
	uc *authusecases.RefreshTokenUseCase
}

// @inject
func NewRefreshTokenHandler(uc *authusecases.RefreshTokenUseCase) *RefreshTokenHandler {
	return &RefreshTokenHandler{uc: uc}
}

// Handle godoc
// @Tags		auth
// @Accept		json
// @Produce		json
// @Param		body	body	RefreshTokenRequest	true	"Refresh token"
// @Success		200	{object}	map[string]any
// @Router		/auth/refresh [post]
func (this *RefreshTokenHandler) Handle(c ctx.Context) error {
	var body RefreshTokenRequest
	if err := c.Bind(&body); err != nil || body.RefreshToken == "" {
		return response.NewFailResponse(http.StatusBadRequest, "refresh_token is required")
	}

	access, err := this.uc.Invoke(c.GetContext(), body.RefreshToken)
	if err != nil {
		return err
	}

	return c.JsonResponse(http.StatusOK, map[string]any{
		"access_token":  access,
		"refresh_token": body.RefreshToken,
	})
}

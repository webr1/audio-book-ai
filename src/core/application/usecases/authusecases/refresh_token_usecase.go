package authusecases

import (
	"context"

	"audio-book-ai/src/core/application/services"
	"audio-book-ai/src/core/domain/entity/enum"
)

type RefreshTokenUseCase struct {
	auth *services.UserAuthTokenService
}

// @inject
func NewRefreshTokenUseCase(auth *services.UserAuthTokenService) *RefreshTokenUseCase {
	return &RefreshTokenUseCase{auth: auth}
}

func (this *RefreshTokenUseCase) Invoke(ctx context.Context, refreshToken string) (string, error) {
	user, err := this.auth.VerifyToken(ctx, refreshToken, enum.TokenTypeRefresh)
	if err != nil {
		return "", err
	}
	return this.auth.GenerateAccessToken(user.ID), nil
}

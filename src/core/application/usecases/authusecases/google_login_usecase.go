package authusecases

import (
	"context"
	"errors"

	"audio-book-ai/src/core/application/response"
	"audio-book-ai/src/core/application/services"
	"audio-book-ai/src/core/domain/entity"
	"audio-book-ai/src/core/domain/ports/gateway"
	"audio-book-ai/src/core/domain/ports/repository"
)

type GoogleLoginUseCase struct {
	gateway    gateway.GoogleOAuthGateway
	repository repository.UserRepository
	auth       *services.UserAuthTokenService
}

// @inject
func NewGoogleLoginUseCase(gateway gateway.GoogleOAuthGateway, repository repository.UserRepository, auth *services.UserAuthTokenService) *GoogleLoginUseCase {
	return &GoogleLoginUseCase{gateway: gateway, repository: repository, auth: auth}
}

func (this *GoogleLoginUseCase) Invoke(ctx context.Context, idToken string) (access, refresh string, user *entity.UserEntity, err error) {
	info, err := this.gateway.VerifyIDToken(ctx, idToken)
	if err != nil {
		return "", "", nil, response.NewFailResponse(401, "invalid google id token")
	}

	user, err = this.repository.FindByGoogleSub(ctx, info.Sub)
	if errors.Is(err, response.NotFoundError) {
		user = entity.NewUserEntity(info.Sub, info.Email, info.Name, info.Picture)
		if err := this.repository.Create(ctx, user); err != nil {
			return "", "", nil, err
		}
	} else if err != nil {
		return "", "", nil, err
	}

	access, refresh = this.auth.GenerateToken(user.ID)
	return access, refresh, user, nil
}

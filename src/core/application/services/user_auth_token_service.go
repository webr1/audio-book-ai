package services

import (
	"context"
	"strconv"
	"time"

	"audio-book-ai/src/core/application/response"
	"audio-book-ai/src/core/domain/entity"
	"audio-book-ai/src/core/domain/entity/enum"
	"audio-book-ai/src/core/domain/ports/repository"
	"audio-book-ai/src/core/domain/ports/security"
	"audio-book-ai/src/infrastructure/env"
)

type UserAuthTokenService struct {
	tokenService  security.TokenServicePort
	repository    repository.UserRepository
	accessExpire  time.Duration
	refreshExpire time.Duration
}

// @inject
func NewUserAuthTokenService(tokenService security.TokenServicePort, repository repository.UserRepository, e *env.Env) *UserAuthTokenService {
	return &UserAuthTokenService{
		tokenService:  tokenService,
		repository:    repository,
		accessExpire:  time.Duration(e.JwtAccessExpireMinutes) * time.Minute,
		refreshExpire: time.Duration(e.JwtRefreshExpireMinutes) * time.Minute,
	}
}

func (this *UserAuthTokenService) GenerateToken(userID uint) (access string, refresh string) {
	return this.GenerateAccessToken(userID), this.generateRefreshToken(userID)
}

func (this *UserAuthTokenService) GenerateAccessToken(userID uint) string {
	claim := entity.NewTokenEntity(
		time.Now().Add(this.accessExpire),
		strconv.Itoa(int(userID)),
		map[string]interface{}{"user_id": userID, "type": string(enum.TokenTypeAccess)},
	)
	return this.tokenService.Encode(claim)
}

func (this *UserAuthTokenService) generateRefreshToken(userID uint) string {
	claim := entity.NewTokenEntity(
		time.Now().Add(this.refreshExpire),
		strconv.Itoa(int(userID)),
		map[string]interface{}{"user_id": userID, "type": string(enum.TokenTypeRefresh)},
	)
	return this.tokenService.Encode(claim)
}

// VerifyToken decodes the JWT, checks its declared type, and always
// re-fetches the current user from the DB rather than trusting claims.
func (this *UserAuthTokenService) VerifyToken(ctx context.Context, tokenString string, tokenType enum.TokenType) (*entity.UserEntity, error) {
	token, err := this.tokenService.Decode(tokenString)
	if err != nil {
		return nil, err
	}

	if token.Payload["type"] != string(tokenType) {
		return nil, response.InvalidTokenError
	}

	userID, err := strconv.Atoi(token.Subject)
	if err != nil {
		return nil, response.InvalidTokenError
	}

	return this.repository.GetByID(ctx, uint(userID))
}

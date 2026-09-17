package security

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"

	"audio-book-ai/src/core/application/response"
	"audio-book-ai/src/core/domain/entity"
	"audio-book-ai/src/core/domain/ports/security"
	"audio-book-ai/src/infrastructure/env"
)

type JwtTokenService struct {
	secret []byte
}

// @inject
func NewJwtTokenService(e *env.Env) security.TokenServicePort {
	return &JwtTokenService{secret: []byte(e.JwtSecret)}
}

func (this *JwtTokenService) Encode(data *entity.TokenEntity) string {
	claim := NewJwtTokenClaim(
		jwt.RegisteredClaims{
			Subject:   data.Subject,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(data.Exp),
			ID:        uuid.New().String(),
		},
		data.Payload,
	)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signed, err := token.SignedString(this.secret)
	if err != nil {
		return ""
	}
	return signed
}

func (this *JwtTokenService) Decode(tokenString string) (*entity.TokenEntity, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtTokenClaim{}, func(token *jwt.Token) (interface{}, error) {
		return this.secret, nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, response.ExpiredTokenError
		}
		return nil, response.InvalidTokenError
	}

	claim := token.Claims.(*JwtTokenClaim)

	return entity.NewTokenEntity(claim.ExpiresAt.Time, claim.Subject, claim.Payload), nil
}

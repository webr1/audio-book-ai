package security

import "audio-book-ai/src/core/domain/entity"

// TokenServicePort signs and verifies opaque JWTs carrying a TokenEntity.
// Implemented over golang-jwt/jwt/v5 (src/infrastructure/security).
type TokenServicePort interface {
	Encode(data *entity.TokenEntity) string
	Decode(tokenString string) (*entity.TokenEntity, error)
}

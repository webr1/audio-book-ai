package security

import "github.com/golang-jwt/jwt/v5"

type JwtTokenClaim struct {
	jwt.RegisteredClaims
	Payload map[string]interface{} `json:"payload"`
}

func NewJwtTokenClaim(registeredClaims jwt.RegisteredClaims, payload map[string]interface{}) *JwtTokenClaim {
	return &JwtTokenClaim{RegisteredClaims: registeredClaims, Payload: payload}
}

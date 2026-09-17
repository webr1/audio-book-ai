package gateway

import (
	"context"
	"fmt"

	"google.golang.org/api/idtoken"

	"audio-book-ai/src/core/domain/ports/gateway"
	"audio-book-ai/src/infrastructure/env"
)

type GoogleOAuthGatewayImpl struct {
	clientID string
}

// @inject
func NewGoogleOAuthGatewayImpl(e *env.Env) gateway.GoogleOAuthGateway {
	return &GoogleOAuthGatewayImpl{clientID: e.GoogleClientID}
}

func (this *GoogleOAuthGatewayImpl) VerifyIDToken(ctx context.Context, idTokenString string) (*gateway.GoogleUserInfo, error) {
	payload, err := idtoken.Validate(ctx, idTokenString, this.clientID)
	if err != nil {
		return nil, fmt.Errorf("verify google id token: %w", err)
	}

	return &gateway.GoogleUserInfo{
		Sub:     stringClaim(payload.Claims, "sub"),
		Email:   stringClaim(payload.Claims, "email"),
		Name:    stringClaim(payload.Claims, "name"),
		Picture: stringClaim(payload.Claims, "picture"),
	}, nil
}

func stringClaim(claims map[string]interface{}, key string) string {
	if v, ok := claims[key].(string); ok {
		return v
	}
	return ""
}

package gateway

import "context"

type GoogleUserInfo struct {
	Sub     string
	Email   string
	Name    string
	Picture string
}

// GoogleOAuthGateway verifies a Google Sign-In ID token and extracts the
// user's profile from its claims. Implemented over the official
// google.golang.org/api/idtoken package (src/infrastructure/gateway),
// which verifies the token's signature against Google's public keys and
// checks its audience against our configured OAuth client ID.
type GoogleOAuthGateway interface {
	VerifyIDToken(ctx context.Context, idToken string) (*GoogleUserInfo, error)
}

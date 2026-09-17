package middlewares

import (
	"strings"

	"audio-book-ai/src/core/application/services"
	"audio-book-ai/src/core/domain/entity/enum"
	"audio-book-ai/src/core/domain/ports/httpport"
	"audio-book-ai/src/core/domain/ports/httpport/ctx"
)

// JwtAuthMiddleware never blocks the request itself: it just sets
// c.SetUser() when a valid Bearer access token is present. Enforcing
// "must be authenticated" is left to permissions.UserAuthenticatedPermission
// applied per-route/group.
type JwtAuthMiddleware struct {
	authService *services.UserAuthTokenService
}

// @inject
func NewJwtAuthMiddleware(authService *services.UserAuthTokenService) *JwtAuthMiddleware {
	return &JwtAuthMiddleware{authService: authService}
}

func (this *JwtAuthMiddleware) Call(next httpport.Handler) httpport.Handler {
	return func(c ctx.Context) error {
		this.authenticate(c)
		return next(c)
	}
}

func (this *JwtAuthMiddleware) authenticate(c ctx.Context) {
	authHeader := c.Header("Authorization")
	if !strings.HasPrefix(authHeader, "Bearer ") {
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	user, err := this.authService.VerifyToken(c.GetContext(), token, enum.TokenTypeAccess)
	if err != nil {
		return
	}
	c.SetUser(user)
}

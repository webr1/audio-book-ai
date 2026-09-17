package groups

import (
	"audio-book-ai/src/core/domain/ports/httpport"
	"audio-book-ai/src/entrypoint/http/handlers/auth"
	"audio-book-ai/src/entrypoint/http/interceptor/permissions"
)

type AuthGroup struct {
	googleLoginHandler  *auth.GoogleLoginHandler
	refreshTokenHandler *auth.RefreshTokenHandler
	meHandler           *auth.MeHandler
}

// @inject
func NewAuthGroup(googleLoginHandler *auth.GoogleLoginHandler, refreshTokenHandler *auth.RefreshTokenHandler, meHandler *auth.MeHandler) *AuthGroup {
	return &AuthGroup{googleLoginHandler: googleLoginHandler, refreshTokenHandler: refreshTokenHandler, meHandler: meHandler}
}

func (this *AuthGroup) RegisterRoutes(g httpport.Group) {
	g.POST("/google", this.googleLoginHandler.Handle)
	g.POST("/refresh", this.refreshTokenHandler.Handle)
	g.GET("/me", this.meHandler.Handle, permissions.UserAuthenticatedPermission)
}

package http

import (
	"net/http"

	"audio-book-ai/src/core/domain/ports/httpport"
	"audio-book-ai/src/core/domain/ports/httpport/ctx"
	"audio-book-ai/src/entrypoint/http/groups"
	"audio-book-ai/src/entrypoint/http/interceptor/middlewares"
	"audio-book-ai/src/entrypoint/http/interceptor/permissions"
	"audio-book-ai/src/infrastructure/env"
	"audio-book-ai/src/infrastructure/storage"
)

type App struct {
	server            httpport.HTTPServer
	bookGroup         *groups.BookGroup
	authGroup         *groups.AuthGroup
	jwtAuthMiddleware *middlewares.JwtAuthMiddleware
	env               *env.Env
}

// @inject
func NewApp(server httpport.HTTPServer, bookGroup *groups.BookGroup, authGroup *groups.AuthGroup, jwtAuthMiddleware *middlewares.JwtAuthMiddleware, e *env.Env) *App {
	return &App{server: server, bookGroup: bookGroup, authGroup: authGroup, jwtAuthMiddleware: jwtAuthMiddleware, env: e}
}

func (this *App) Init() {
	if err := storage.EnsureDataDirs(this.env.DataDir); err != nil {
		panic(err)
	}

	this.server.Init()
	this.initMiddlewares()
	this.initGroups()
}

func (this *App) initMiddlewares() {
	this.server.Use(this.jwtAuthMiddleware.Call)
}

func (this *App) initGroups() {
	api := this.server.Group("/api")
	api.GET("/health", healthHandler)

	this.authGroup.RegisterRoutes(this.group("/auth"))
	this.bookGroup.RegisterRoutes(this.group("/book", permissions.UserAuthenticatedPermission))
}

func (this *App) group(prefix string, mws ...httpport.Middleware) httpport.Group {
	return this.server.Group("/api"+prefix, mws...)
}

func (this *App) Start() {
	if err := this.server.Run(":" + this.env.Port); err != nil {
		panic(err)
	}
}

func healthHandler(c ctx.Context) error {
	return c.JsonResponse(http.StatusOK, map[string]any{"status": "ok"})
}

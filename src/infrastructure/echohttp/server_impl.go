package echohttp

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"audio-book-ai/src/core/domain/ports/httpport"
	"audio-book-ai/src/infrastructure/echohttp/defaults"
	"audio-book-ai/src/infrastructure/echohttp/mapper"
	"audio-book-ai/src/infrastructure/echohttp/validator"
)

type ServerImpl struct {
	echo *echo.Echo
}

// @inject
func NewServerImpl() httpport.HTTPServer {
	e := echo.New()
	e.Validator = validator.New()
	e.HTTPErrorHandler = defaults.HTTPErrorHandler
	return &ServerImpl{echo: e}
}

func (this *ServerImpl) Init() {
	this.echo.Use(middleware.CORS())
	this.echo.Use(defaults.ContextMiddleware)
	this.echo.Use(defaults.RecoveryMiddleware)
}

func (this *ServerImpl) Use(middlewares ...httpport.Middleware) {
	for _, mw := range middlewares {
		this.echo.Use(mapper.ToEchoMiddleware(mw))
	}
}

func (this *ServerImpl) Group(prefix string, middlewares ...httpport.Middleware) httpport.Group {
	g := this.echo.Group(prefix, mapper.ToEchoMiddlewares(middlewares)...)
	return NewGroupImpl(g)
}

func (this *ServerImpl) Run(address string) error {
	return this.echo.Start(address)
}

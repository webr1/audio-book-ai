package echohttp

import (
	"github.com/labstack/echo/v4"

	"audio-book-ai/src/core/domain/ports/httpport"
	"audio-book-ai/src/infrastructure/echohttp/mapper"
)

type GroupImpl struct {
	group *echo.Group
}

func NewGroupImpl(g *echo.Group) httpport.Group {
	return &GroupImpl{group: g}
}

func (this *GroupImpl) Use(middlewares ...httpport.Middleware) {
	for _, mw := range middlewares {
		this.group.Use(mapper.ToEchoMiddleware(mw))
	}
}

func (this *GroupImpl) GET(path string, handler httpport.Handler, middlewares ...httpport.Middleware) {
	this.group.GET(path, mapper.ToEchoHandler(handler), mapper.ToEchoMiddlewares(middlewares)...)
}

func (this *GroupImpl) POST(path string, handler httpport.Handler, middlewares ...httpport.Middleware) {
	this.group.POST(path, mapper.ToEchoHandler(handler), mapper.ToEchoMiddlewares(middlewares)...)
}

func (this *GroupImpl) PUT(path string, handler httpport.Handler, middlewares ...httpport.Middleware) {
	this.group.PUT(path, mapper.ToEchoHandler(handler), mapper.ToEchoMiddlewares(middlewares)...)
}

func (this *GroupImpl) DELETE(path string, handler httpport.Handler, middlewares ...httpport.Middleware) {
	this.group.DELETE(path, mapper.ToEchoHandler(handler), mapper.ToEchoMiddlewares(middlewares)...)
}

func (this *GroupImpl) PATCH(path string, handler httpport.Handler, middlewares ...httpport.Middleware) {
	this.group.PATCH(path, mapper.ToEchoHandler(handler), mapper.ToEchoMiddlewares(middlewares)...)
}

func (this *GroupImpl) Any(path string, handler httpport.Handler, middlewares ...httpport.Middleware) {
	this.group.Any(path, mapper.ToEchoHandler(handler), mapper.ToEchoMiddlewares(middlewares)...)
}

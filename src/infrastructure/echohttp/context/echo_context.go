package context

import (
	"context"
	"mime/multipart"

	"github.com/labstack/echo/v4"

	"audio-book-ai/src/core/application/response"
	"audio-book-ai/src/core/domain/entity"
	"audio-book-ai/src/core/domain/ports/httpport/ctx"
)

// EchoContext embeds echo.Context so a single value satisfies both
// echo.Context (needed by the Echo framework/middleware chain) and
// ctx.Context (needed by everything above the framework boundary).
type EchoContext struct {
	echo.Context
	user *entity.UserEntity
}

// @inject
func NewEchoContext(c echo.Context) ctx.Context {
	return &EchoContext{Context: c}
}

func (this *EchoContext) GetContext() context.Context {
	return this.Context.Request().Context()
}

func (this *EchoContext) JSON(status int, i interface{}) error {
	return this.Context.JSON(status, i)
}

func (this *EchoContext) JsonResponse(status int, i interface{}) error {
	return this.Context.JSON(status, response.NewResponse(status, 0, "", i))
}

func (this *EchoContext) File(path string) error {
	return this.Context.File(path)
}

func (this *EchoContext) Param(name string) string {
	return this.Context.Param(name)
}

func (this *EchoContext) QueryParam(name string) string {
	return this.Context.QueryParam(name)
}

func (this *EchoContext) Header(name string) string {
	return this.Context.Request().Header.Get(name)
}

func (this *EchoContext) User() *entity.UserEntity {
	return this.user
}

func (this *EchoContext) SetUser(user *entity.UserEntity) {
	this.user = user
}

func (this *EchoContext) Bind(i interface{}) error {
	return this.Context.Bind(i)
}

func (this *EchoContext) Validate(i interface{}) error {
	return this.Context.Validate(i)
}

func (this *EchoContext) FormFile(name string) (*multipart.FileHeader, error) {
	return this.Context.FormFile(name)
}

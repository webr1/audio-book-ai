package ctx

import (
	"context"
	"mime/multipart"

	"audio-book-ai/src/core/domain/entity"
)

// Context abstracts the HTTP framework away from handlers/usecases.
// Echo is the only adapter today (src/infrastructure/echohttp), but
// nothing above this port imports echo directly.
type Context interface {
	GetContext() context.Context
	JSON(status int, i interface{}) error
	JsonResponse(status int, i interface{}) error
	File(path string) error
	Param(name string) string
	QueryParam(name string) string
	Header(name string) string
	Bind(i interface{}) error
	Validate(i interface{}) error
	FormFile(name string) (*multipart.FileHeader, error)
	User() *entity.UserEntity
	SetUser(user *entity.UserEntity)
}

package repository

import (
	"context"

	"audio-book-ai/src/core/domain/entity"
)

type UserRepository interface {
	FindByGoogleSub(ctx context.Context, googleSub string) (*entity.UserEntity, error)
	Create(ctx context.Context, user *entity.UserEntity) error
	GetByID(ctx context.Context, id uint) (*entity.UserEntity, error)
}

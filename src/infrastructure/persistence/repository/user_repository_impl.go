package repository

import (
	"context"

	"audio-book-ai/src/core/domain/entity"
	"audio-book-ai/src/core/domain/ports/repository"
	infraError "audio-book-ai/src/infrastructure/errors"
	"audio-book-ai/src/infrastructure/persistence/mapper"
	"audio-book-ai/src/infrastructure/persistence/models"
)

type UserRepositoryImpl struct {
	*BaseRepository
}

// @inject
func NewUserRepositoryImpl(baseRepository *BaseRepository) repository.UserRepository {
	return &UserRepositoryImpl{BaseRepository: baseRepository}
}

func (this *UserRepositoryImpl) FindByGoogleSub(ctx context.Context, googleSub string) (*entity.UserEntity, error) {
	db := this.db().WithContext(ctx)

	var model models.UserModel
	if err := db.Where("google_sub = ?", googleSub).First(&model).Error; err != nil {
		return nil, infraError.Wrap(err)
	}

	return mapper.UserModelToEntity(&model), nil
}

func (this *UserRepositoryImpl) Create(ctx context.Context, user *entity.UserEntity) error {
	db := this.db().WithContext(ctx)

	model := mapper.UserEntityToModel(user)
	if err := db.Create(model).Error; err != nil {
		return infraError.Wrap(err)
	}

	user.ID = model.ID
	user.CreatedAt = model.CreatedAt
	user.UpdatedAt = model.UpdatedAt
	return nil
}

func (this *UserRepositoryImpl) GetByID(ctx context.Context, id uint) (*entity.UserEntity, error) {
	db := this.db().WithContext(ctx)

	var model models.UserModel
	if err := db.Where("id = ?", id).First(&model).Error; err != nil {
		return nil, infraError.Wrap(err)
	}

	return mapper.UserModelToEntity(&model), nil
}

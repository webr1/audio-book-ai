package mapper

import (
	"audio-book-ai/src/core/domain/entity"
	"audio-book-ai/src/infrastructure/persistence/models"
)

func UserEntityToModel(e *entity.UserEntity) *models.UserModel {
	return models.NewUserModel(e.GoogleSub, e.Email, e.Name, e.Picture)
}

func UserModelToEntity(m *models.UserModel) *entity.UserEntity {
	return &entity.UserEntity{
		ID:        m.ID,
		GoogleSub: m.GoogleSub,
		Email:     m.Email,
		Name:      m.Name,
		Picture:   m.Picture,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

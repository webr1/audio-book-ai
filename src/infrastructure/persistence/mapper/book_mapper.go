package mapper

import (
	"audio-book-ai/src/core/domain/entity"
	"audio-book-ai/src/core/domain/entity/enum"
	"audio-book-ai/src/infrastructure/persistence/models"
)

func BookEntityToModel(e *entity.BookEntity) *models.BookModel {
	return models.NewBookModel(e.UserID, e.Filename, string(e.Status))
}

func BookModelToEntity(m *models.BookModel) *entity.BookEntity {
	return &entity.BookEntity{
		ID:          m.ID,
		UserID:      m.UserID,
		Filename:    m.Filename,
		Status:      enum.BookStatus(m.Status),
		WordCount:   m.WordCount,
		CharCount:   m.CharCount,
		ChunksTotal: m.ChunksTotal,
		ChunksDone:  m.ChunksDone,
		ErrorMsg:    m.ErrorMsg,
		OutputPath:  m.OutputPath,
		CreatedAt:   m.CreatedAt,
		UpdatedAt:   m.UpdatedAt,
	}
}

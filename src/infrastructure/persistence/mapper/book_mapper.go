package mapper

import (
	"audio-book-ai/src/core/domain/entity"
	"audio-book-ai/src/core/domain/entity/enum"
	"audio-book-ai/src/infrastructure/persistence/models"
)

func BookEntityToModel(e *entity.BookEntity) *models.BookModel {
	return models.NewBookModel(e.UserID, e.Filename, string(e.Status), string(e.TrialStatus), string(e.PaymentStatus))
}

func BookModelToEntity(m *models.BookModel) *entity.BookEntity {
	return &entity.BookEntity{
		ID:          m.ID,
		UserID:      m.UserID,
		Filename:    m.Filename,
		Status:      enum.BookStatus(m.Status),
		WordCount:   m.WordCount,
		CharCount:   m.CharCount,
		Price:       m.Price,
		ChunksTotal: m.ChunksTotal,
		ChunksDone:  m.ChunksDone,
		ErrorMsg:    m.ErrorMsg,
		OutputPath:  m.OutputPath,

		TrialStatus:      enum.BookStatus(m.TrialStatus),
		TrialChunksTotal: m.TrialChunksTotal,
		TrialChunksDone:  m.TrialChunksDone,
		TrialErrorMsg:    m.TrialErrorMsg,
		TrialOutputPath:  m.TrialOutputPath,

		PaymentStatus:      enum.PaymentStatus(m.PaymentStatus),
		PaymentReceiptPath: m.PaymentReceiptPath,

		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}

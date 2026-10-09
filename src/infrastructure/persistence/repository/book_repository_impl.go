package repository

import (
	"context"

	"gorm.io/gorm"

	"audio-book-ai/src/core/domain/entity"
	"audio-book-ai/src/core/domain/entity/enum"
	"audio-book-ai/src/core/domain/ports/repository"
	infraError "audio-book-ai/src/infrastructure/errors"
	"audio-book-ai/src/infrastructure/persistence/mapper"
	"audio-book-ai/src/infrastructure/persistence/models"
)

var gormRecordNotFound = gorm.ErrRecordNotFound

type BookRepositoryImpl struct {
	*BaseRepository
}

// @inject
func NewBookRepositoryImpl(baseRepository *BaseRepository) repository.BookRepository {
	return &BookRepositoryImpl{BaseRepository: baseRepository}
}

func (this *BookRepositoryImpl) Create(ctx context.Context, book *entity.BookEntity) error {
	db := this.db().WithContext(ctx)

	model := mapper.BookEntityToModel(book)
	if err := db.Create(model).Error; err != nil {
		return infraError.Wrap(err)
	}

	book.ID = model.ID
	book.CreatedAt = model.CreatedAt
	book.UpdatedAt = model.UpdatedAt
	return nil
}

func (this *BookRepositoryImpl) Get(ctx context.Context, id uint) (*entity.BookEntity, error) {
	db := this.db().WithContext(ctx)

	var model models.BookModel
	if err := db.Where("id = ?", id).First(&model).Error; err != nil {
		return nil, infraError.Wrap(err)
	}

	return mapper.BookModelToEntity(&model), nil
}

func (this *BookRepositoryImpl) UpdateStatus(ctx context.Context, id uint, status enum.BookStatus, errMsg string) error {
	db := this.db().WithContext(ctx)

	result := db.Model(&models.BookModel{}).Where("id = ?", id).Updates(map[string]interface{}{
		"status":    string(status),
		"error_msg": errMsg,
	})
	if result.Error != nil {
		return infraError.Wrap(result.Error)
	}
	if result.RowsAffected == 0 {
		return infraError.Wrap(gormRecordNotFound)
	}
	return nil
}

func (this *BookRepositoryImpl) UpdateProgress(ctx context.Context, id uint, chunksDone, chunksTotal int) error {
	db := this.db().WithContext(ctx)

	result := db.Model(&models.BookModel{}).Where("id = ?", id).Updates(map[string]interface{}{
		"chunks_done":  chunksDone,
		"chunks_total": chunksTotal,
	})
	if result.Error != nil {
		return infraError.Wrap(result.Error)
	}
	if result.RowsAffected == 0 {
		return infraError.Wrap(gormRecordNotFound)
	}
	return nil
}

func (this *BookRepositoryImpl) SetCounts(ctx context.Context, id uint, wordCount, charCount, chunksTotal int) error {
	db := this.db().WithContext(ctx)

	result := db.Model(&models.BookModel{}).Where("id = ?", id).Updates(map[string]interface{}{
		"word_count":   wordCount,
		"char_count":   charCount,
		"chunks_total": chunksTotal,
	})
	if result.Error != nil {
		return infraError.Wrap(result.Error)
	}
	if result.RowsAffected == 0 {
		return infraError.Wrap(gormRecordNotFound)
	}
	return nil
}

func (this *BookRepositoryImpl) SetPrice(ctx context.Context, id uint, price float64) error {
	db := this.db().WithContext(ctx)

	result := db.Model(&models.BookModel{}).Where("id = ?", id).Update("price", price)
	if result.Error != nil {
		return infraError.Wrap(result.Error)
	}
	if result.RowsAffected == 0 {
		return infraError.Wrap(gormRecordNotFound)
	}
	return nil
}

func (this *BookRepositoryImpl) SetOutputPath(ctx context.Context, id uint, path string) error {
	db := this.db().WithContext(ctx)

	result := db.Model(&models.BookModel{}).Where("id = ?", id).Update("output_path", path)
	if result.Error != nil {
		return infraError.Wrap(result.Error)
	}
	if result.RowsAffected == 0 {
		return infraError.Wrap(gormRecordNotFound)
	}
	return nil
}

func (this *BookRepositoryImpl) UpdateTrialStatus(ctx context.Context, id uint, status enum.BookStatus, errMsg string) error {
	db := this.db().WithContext(ctx)

	result := db.Model(&models.BookModel{}).Where("id = ?", id).Updates(map[string]interface{}{
		"trial_status":    string(status),
		"trial_error_msg": errMsg,
	})
	if result.Error != nil {
		return infraError.Wrap(result.Error)
	}
	if result.RowsAffected == 0 {
		return infraError.Wrap(gormRecordNotFound)
	}
	return nil
}

func (this *BookRepositoryImpl) UpdateTrialProgress(ctx context.Context, id uint, chunksDone, chunksTotal int) error {
	db := this.db().WithContext(ctx)

	result := db.Model(&models.BookModel{}).Where("id = ?", id).Updates(map[string]interface{}{
		"trial_chunks_done":  chunksDone,
		"trial_chunks_total": chunksTotal,
	})
	if result.Error != nil {
		return infraError.Wrap(result.Error)
	}
	if result.RowsAffected == 0 {
		return infraError.Wrap(gormRecordNotFound)
	}
	return nil
}

func (this *BookRepositoryImpl) SetTrialOutputPath(ctx context.Context, id uint, path string) error {
	db := this.db().WithContext(ctx)

	result := db.Model(&models.BookModel{}).Where("id = ?", id).Update("trial_output_path", path)
	if result.Error != nil {
		return infraError.Wrap(result.Error)
	}
	if result.RowsAffected == 0 {
		return infraError.Wrap(gormRecordNotFound)
	}
	return nil
}

func (this *BookRepositoryImpl) SetPaymentReceipt(ctx context.Context, id uint, path string) error {
	db := this.db().WithContext(ctx)

	result := db.Model(&models.BookModel{}).Where("id = ?", id).Updates(map[string]interface{}{
		"payment_receipt_path": path,
		"payment_status":       string(enum.PaymentStatusPendingReview),
	})
	if result.Error != nil {
		return infraError.Wrap(result.Error)
	}
	if result.RowsAffected == 0 {
		return infraError.Wrap(gormRecordNotFound)
	}
	return nil
}

func (this *BookRepositoryImpl) SetPaymentStatus(ctx context.Context, id uint, status enum.PaymentStatus) error {
	db := this.db().WithContext(ctx)

	result := db.Model(&models.BookModel{}).Where("id = ?", id).Update("payment_status", string(status))
	if result.Error != nil {
		return infraError.Wrap(result.Error)
	}
	if result.RowsAffected == 0 {
		return infraError.Wrap(gormRecordNotFound)
	}
	return nil
}

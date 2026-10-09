package models

import "gorm.io/gorm"

type BookModel struct {
	gorm.Model

	UserID      uint    `gorm:"column:user_id;not null;index"`
	Filename    string  `gorm:"column:filename;type:varchar(255);not null"`
	Status      string  `gorm:"column:status;type:varchar(32);not null;default:pending"`
	WordCount   int     `gorm:"column:word_count;not null;default:0"`
	CharCount   int     `gorm:"column:char_count;not null;default:0"`
	Price       float64 `gorm:"column:price;type:double precision;not null;default:0"`
	ChunksTotal int     `gorm:"column:chunks_total;not null;default:0"`
	ChunksDone  int     `gorm:"column:chunks_done;not null;default:0"`
	ErrorMsg    string  `gorm:"column:error_msg;type:text;not null;default:''"`
	OutputPath  string  `gorm:"column:output_path;type:varchar(512);not null;default:''"`

	TrialStatus      string `gorm:"column:trial_status;type:varchar(32);not null;default:pending"`
	TrialChunksTotal int    `gorm:"column:trial_chunks_total;not null;default:0"`
	TrialChunksDone  int    `gorm:"column:trial_chunks_done;not null;default:0"`
	TrialErrorMsg    string `gorm:"column:trial_error_msg;type:text;not null;default:''"`
	TrialOutputPath  string `gorm:"column:trial_output_path;type:varchar(512);not null;default:''"`

	PaymentStatus      string `gorm:"column:payment_status;type:varchar(32);not null;default:none"`
	PaymentReceiptPath string `gorm:"column:payment_receipt_path;type:varchar(512);not null;default:''"`
}

func (BookModel) TableName() string {
	return "books"
}

func NewBookModel(userID uint, filename, status, trialStatus, paymentStatus string) *BookModel {
	return &BookModel{
		UserID:        userID,
		Filename:      filename,
		Status:        status,
		TrialStatus:   trialStatus,
		PaymentStatus: paymentStatus,
	}
}

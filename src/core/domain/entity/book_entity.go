package entity

import (
	"time"

	"audio-book-ai/src/core/domain/entity/enum"
)

type BookEntity struct {
	ID          uint
	UserID      uint
	Filename    string
	Status      enum.BookStatus
	WordCount   int
	CharCount   int
	Price       float64
	ChunksTotal int
	ChunksDone  int
	ErrorMsg    string
	OutputPath  string

	TrialStatus      enum.BookStatus
	TrialChunksTotal int
	TrialChunksDone  int
	TrialErrorMsg    string
	TrialOutputPath  string

	PaymentStatus      enum.PaymentStatus
	PaymentReceiptPath string

	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewBookEntity(userID uint, filename string) *BookEntity {
	return &BookEntity{
		UserID:        userID,
		Filename:      filename,
		Status:        enum.BookStatusPending,
		TrialStatus:   enum.BookStatusPending,
		PaymentStatus: enum.PaymentStatusNone,
	}
}

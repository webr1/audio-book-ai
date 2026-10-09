package bookusecases

import (
	"context"
	"encoding/json"

	"audio-book-ai/src/core/domain/entity/enum"
	"audio-book-ai/src/core/domain/ports/repository"
	"audio-book-ai/src/core/domain/ports/taskport"
)

// ConfirmPaymentUseCase marks a book's receipt as approved and kicks off
// the full-book synthesis pipeline. It is intentionally not owner-scoped —
// the caller (handler) authorizes this via an operator-only check, since
// the book owner can never confirm their own payment.
type ConfirmPaymentUseCase struct {
	repository repository.BookRepository
	enqueuer   taskport.Enqueuer
}

// @inject
func NewConfirmPaymentUseCase(repository repository.BookRepository, enqueuer taskport.Enqueuer) *ConfirmPaymentUseCase {
	return &ConfirmPaymentUseCase{repository: repository, enqueuer: enqueuer}
}

func (this *ConfirmPaymentUseCase) Invoke(ctx context.Context, bookID uint) error {
	if err := this.repository.SetPaymentStatus(ctx, bookID, enum.PaymentStatusPaid); err != nil {
		return err
	}

	payload, err := json.Marshal(ProcessBookPayload{BookID: bookID})
	if err != nil {
		return err
	}
	return this.enqueuer.Enqueue(ctx, TaskTypeProcessBook, payload)
}

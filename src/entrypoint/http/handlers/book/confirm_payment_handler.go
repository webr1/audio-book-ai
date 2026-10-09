package book

import (
	"net/http"

	"audio-book-ai/src/core/application/response"
	"audio-book-ai/src/core/application/usecases/bookusecases"
	"audio-book-ai/src/core/domain/entity/enum"
	"audio-book-ai/src/core/domain/ports/httpport/ctx"
	"audio-book-ai/src/infrastructure/env"
)

// ConfirmPaymentHandler is operator-only: it is gated by a shared secret
// (X-Admin-Key) rather than by being the book's owner, since confirming
// your own payment would defeat the point. There is no roles/admin system
// yet — this is a deliberately simple stand-in until one exists.
type ConfirmPaymentHandler struct {
	uc  *bookusecases.ConfirmPaymentUseCase
	env *env.Env
}

// @inject
func NewConfirmPaymentHandler(uc *bookusecases.ConfirmPaymentUseCase, e *env.Env) *ConfirmPaymentHandler {
	return &ConfirmPaymentHandler{uc: uc, env: e}
}

// Handle godoc
// @Tags		book
// @Produce		json
// @Param		id	path	int	true	"Book ID"
// @Param		X-Admin-Key	header	string	true	"Operator key"
// @Success		200	{object}	map[string]any
// @Router		/book/{id}/confirm-payment [post]
func (this *ConfirmPaymentHandler) Handle(c ctx.Context) error {
	if this.env.AdminAPIKey == "" || c.Header("X-Admin-Key") != this.env.AdminAPIKey {
		return response.NewFailResponse(http.StatusForbidden, "not authorized to confirm payments")
	}

	id := ctx.GetUintPathParam(c, "id")

	if err := this.uc.Invoke(c.GetContext(), id); err != nil {
		return err
	}

	return c.JsonResponse(http.StatusOK, map[string]any{
		"payment_status": enum.PaymentStatusPaid,
	})
}

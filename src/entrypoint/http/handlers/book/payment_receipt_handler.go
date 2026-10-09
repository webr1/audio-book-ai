package book

import (
	"net/http"
	"path/filepath"
	"strings"

	"audio-book-ai/src/core/application/response"
	"audio-book-ai/src/core/application/usecases/bookusecases"
	"audio-book-ai/src/core/domain/entity/enum"
	"audio-book-ai/src/core/domain/ports/httpport/ctx"
)

var allowedReceiptExts = map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".pdf": true}

type PaymentReceiptHandler struct {
	getUC    *bookusecases.BookGetUseCase
	uploadUC *bookusecases.UploadReceiptUseCase
}

// @inject
func NewPaymentReceiptHandler(getUC *bookusecases.BookGetUseCase, uploadUC *bookusecases.UploadReceiptUseCase) *PaymentReceiptHandler {
	return &PaymentReceiptHandler{getUC: getUC, uploadUC: uploadUC}
}

// Handle godoc
// @Tags		book
// @Accept		multipart/form-data
// @Produce		json
// @Param		id		path		int		true	"Book ID"
// @Param		file	formData	file	true	"Payment receipt photo/scan (.jpg/.jpeg/.png/.pdf)"
// @Success		200	{object}	map[string]any
// @Router		/book/{id}/payment-receipt [post]
func (this *PaymentReceiptHandler) Handle(c ctx.Context) error {
	id := ctx.GetUintPathParam(c, "id")

	// Ownership check — BookGetUseCase returns NotFound for someone else's book.
	if _, err := this.getUC.Invoke(c.GetContext(), id, c.User().ID); err != nil {
		return err
	}

	fileHeader, err := c.FormFile("file")
	if err != nil {
		return response.NewFailResponse(http.StatusBadRequest, "file is required")
	}

	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	if !allowedReceiptExts[ext] {
		return response.NewFailResponse(http.StatusBadRequest, "unsupported file extension: "+ext+" (allowed: .jpg, .jpeg, .png, .pdf)")
	}

	file, err := fileHeader.Open()
	if err != nil {
		return response.NewFailResponse(http.StatusBadRequest, "could not read uploaded file")
	}
	defer file.Close()

	if err := this.uploadUC.Invoke(c.GetContext(), id, ext, file); err != nil {
		return err
	}

	return c.JsonResponse(http.StatusOK, map[string]any{
		"payment_status": enum.PaymentStatusPendingReview,
	})
}

package book

import (
	"net/http"
	"path/filepath"

	"audio-book-ai/src/core/application/response"
	"audio-book-ai/src/core/application/usecases/bookusecases"
	"audio-book-ai/src/core/domain/ports/httpport/ctx"
)

type CreateHandler struct {
	uc *bookusecases.BookCreateUseCase
}

// @inject
func NewCreateHandler(uc *bookusecases.BookCreateUseCase) *CreateHandler {
	return &CreateHandler{uc: uc}
}

// Handle godoc
// @Tags		book
// @Accept		multipart/form-data
// @Produce		json
// @Param		file	formData	file	true	"Book text file (.txt)"
// @Success		202	{object}	map[string]any
// @Router		/book/create [post]
func (this *CreateHandler) Handle(c ctx.Context) error {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		return response.NewFailResponse(http.StatusBadRequest, "file is required")
	}

	if ext := filepath.Ext(fileHeader.Filename); ext != ".txt" {
		return response.NewFailResponse(http.StatusBadRequest, "unsupported file extension: "+ext+" (only .txt for now)")
	}

	file, err := fileHeader.Open()
	if err != nil {
		return response.NewFailResponse(http.StatusBadRequest, "could not read uploaded file")
	}
	defer file.Close()

	book, err := this.uc.Invoke(c.GetContext(), c.User().ID, fileHeader.Filename, file)
	if err != nil {
		return err
	}

	return c.JsonResponse(http.StatusAccepted, map[string]any{
		"id":     book.ID,
		"status": book.Status,
	})
}

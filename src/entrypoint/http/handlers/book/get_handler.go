package book

import (
	"net/http"

	"audio-book-ai/src/core/application/usecases/bookusecases"
	"audio-book-ai/src/core/domain/ports/httpport/ctx"
)

type GetHandler struct {
	uc *bookusecases.BookGetUseCase
}

// @inject
func NewGetHandler(uc *bookusecases.BookGetUseCase) *GetHandler {
	return &GetHandler{uc: uc}
}

// Handle godoc
// @Tags		book
// @Produce		json
// @Param		id	path	int	true	"Book ID"
// @Success		200	{object}	map[string]any
// @Router		/book/{id} [get]
func (this *GetHandler) Handle(c ctx.Context) error {
	id := ctx.GetUintPathParam(c, "id")

	book, err := this.uc.Invoke(c.GetContext(), id, c.User().ID)
	if err != nil {
		return err
	}

	return c.JsonResponse(http.StatusOK, map[string]any{
		"id":           book.ID,
		"filename":     book.Filename,
		"status":       book.Status,
		"word_count":   book.WordCount,
		"char_count":   book.CharCount,
		"price":        book.Price,
		"chunks_done":  book.ChunksDone,
		"chunks_total": book.ChunksTotal,
		"error":        book.ErrorMsg,
		"created_at":   book.CreatedAt,
		"updated_at":   book.UpdatedAt,
	})
}

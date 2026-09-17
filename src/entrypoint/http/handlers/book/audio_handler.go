package book

import (
	"net/http"

	"audio-book-ai/src/core/application/response"
	"audio-book-ai/src/core/application/usecases/bookusecases"
	"audio-book-ai/src/core/domain/entity/enum"
	"audio-book-ai/src/core/domain/ports/httpport/ctx"
)

type AudioHandler struct {
	uc *bookusecases.BookGetUseCase
}

// @inject
func NewAudioHandler(uc *bookusecases.BookGetUseCase) *AudioHandler {
	return &AudioHandler{uc: uc}
}

// Handle godoc
// @Tags		book
// @Produce		audio/wav
// @Param		id	path	int	true	"Book ID"
// @Success		200	{file}		binary
// @Router		/book/{id}/audio [get]
func (this *AudioHandler) Handle(c ctx.Context) error {
	id := ctx.GetUintPathParam(c, "id")

	book, err := this.uc.Invoke(c.GetContext(), id, c.User().ID)
	if err != nil {
		return err
	}

	if book.Status != enum.BookStatusDone {
		return response.NewFailResponse(http.StatusConflict, "book not ready, status="+string(book.Status))
	}

	return c.File(book.OutputPath)
}

package book

import (
	"net/http"

	"audio-book-ai/src/core/application/response"
	"audio-book-ai/src/core/application/usecases/bookusecases"
	"audio-book-ai/src/core/domain/entity/enum"
	"audio-book-ai/src/core/domain/ports/httpport/ctx"
)

type TrialAudioHandler struct {
	uc *bookusecases.BookGetUseCase
}

// @inject
func NewTrialAudioHandler(uc *bookusecases.BookGetUseCase) *TrialAudioHandler {
	return &TrialAudioHandler{uc: uc}
}

// Handle godoc
// @Tags		book
// @Produce		audio/wav
// @Param		id	path	int	true	"Book ID"
// @Success		200	{file}		binary
// @Router		/book/{id}/trial-audio [get]
func (this *TrialAudioHandler) Handle(c ctx.Context) error {
	id := ctx.GetUintPathParam(c, "id")

	book, err := this.uc.Invoke(c.GetContext(), id, c.User().ID)
	if err != nil {
		return err
	}

	if book.TrialStatus != enum.BookStatusDone {
		return response.NewFailResponse(http.StatusConflict, "trial not ready, status="+string(book.TrialStatus))
	}

	return c.File(book.TrialOutputPath)
}

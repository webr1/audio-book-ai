package groups

import (
	"audio-book-ai/src/core/domain/ports/httpport"
	"audio-book-ai/src/entrypoint/http/handlers/book"
)

type BookGroup struct {
	createHandler *book.CreateHandler
	getHandler    *book.GetHandler
	audioHandler  *book.AudioHandler
}

// @inject
func NewBookGroup(createHandler *book.CreateHandler, getHandler *book.GetHandler, audioHandler *book.AudioHandler) *BookGroup {
	return &BookGroup{createHandler: createHandler, getHandler: getHandler, audioHandler: audioHandler}
}

func (this *BookGroup) RegisterRoutes(g httpport.Group) {
	g.POST("/create", this.createHandler.Handle)
	g.GET("/:id", this.getHandler.Handle)
	g.GET("/:id/audio", this.audioHandler.Handle)
}

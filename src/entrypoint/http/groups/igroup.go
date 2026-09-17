package groups

import "audio-book-ai/src/core/domain/ports/httpport"

type IGroup interface {
	RegisterRoutes(g httpport.Group)
}

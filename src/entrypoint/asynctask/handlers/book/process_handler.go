package book

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"

	"audio-book-ai/src/core/application/usecases/bookusecases"
)

type ProcessHandler struct {
	uc *bookusecases.ProcessBookUseCase
}

// @inject
func NewProcessHandler(uc *bookusecases.ProcessBookUseCase) *ProcessHandler {
	return &ProcessHandler{uc: uc}
}

func (this *ProcessHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var payload bookusecases.ProcessBookPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return fmt.Errorf("unmarshal process-book payload: %w", err)
	}
	return this.uc.Invoke(ctx, payload.BookID)
}

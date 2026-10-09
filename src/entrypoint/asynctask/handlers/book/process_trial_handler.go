package book

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"

	"audio-book-ai/src/core/application/usecases/bookusecases"
)

type ProcessTrialHandler struct {
	uc *bookusecases.ProcessTrialUseCase
}

// @inject
func NewProcessTrialHandler(uc *bookusecases.ProcessTrialUseCase) *ProcessTrialHandler {
	return &ProcessTrialHandler{uc: uc}
}

func (this *ProcessTrialHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {
	var payload bookusecases.ProcessBookPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return fmt.Errorf("unmarshal process-trial payload: %w", err)
	}
	return this.uc.Invoke(ctx, payload.BookID)
}

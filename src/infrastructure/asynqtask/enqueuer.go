package asynqtask

import (
	"context"

	"github.com/hibiken/asynq"

	"audio-book-ai/src/core/domain/ports/taskport"
)

type Enqueuer struct {
	client *asynq.Client
}

// @inject
func NewEnqueuer(client *asynq.Client) taskport.Enqueuer {
	return &Enqueuer{client: client}
}

func (e *Enqueuer) Enqueue(ctx context.Context, taskType string, payload []byte) error {
	task := asynq.NewTask(taskType, payload)
	_, err := e.client.EnqueueContext(ctx, task)
	return err
}

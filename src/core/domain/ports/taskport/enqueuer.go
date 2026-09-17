package taskport

import "context"

// Enqueuer schedules a background task by type name with an opaque payload.
// Implemented over asynq; usecases depend on this port, never on asynq directly.
type Enqueuer interface {
	Enqueue(ctx context.Context, taskType string, payload []byte) error
}

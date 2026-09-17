package asynqtask

import (
	"github.com/hibiken/asynq"

	"audio-book-ai/src/infrastructure/env"
)

// @inject
func NewAsynqClient(e *env.Env) *asynq.Client {
	return asynq.NewClient(asynq.RedisClientOpt{
		Addr:     e.RedisAddr,
		Password: e.RedisPassword,
	})
}

func RedisConnOpt(e *env.Env) asynq.RedisClientOpt {
	return asynq.RedisClientOpt{Addr: e.RedisAddr, Password: e.RedisPassword}
}

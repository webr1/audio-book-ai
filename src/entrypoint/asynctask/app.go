package asynctask

import (
	"github.com/hibiken/asynq"

	"audio-book-ai/src/core/application/usecases/bookusecases"
	"audio-book-ai/src/entrypoint/asynctask/handlers/book"
	"audio-book-ai/src/infrastructure/asynqtask"
	"audio-book-ai/src/infrastructure/env"
	"audio-book-ai/src/infrastructure/storage"
)

type App struct {
	server         *asynq.Server
	mux            *asynq.ServeMux
	processHandler *book.ProcessHandler
	env            *env.Env
}

// @inject
func NewApp(e *env.Env, processHandler *book.ProcessHandler) *App {
	server := asynq.NewServer(asynqtask.RedisConnOpt(e), asynq.Config{Concurrency: 5})
	return &App{server: server, mux: asynq.NewServeMux(), processHandler: processHandler, env: e}
}

func (this *App) Init() {
	if err := storage.EnsureDataDirs(this.env.DataDir); err != nil {
		panic(err)
	}
	this.mux.HandleFunc(bookusecases.TaskTypeProcessBook, this.processHandler.ProcessTask)
}

func (this *App) Start() {
	if err := this.server.Run(this.mux); err != nil {
		panic(err)
	}
}

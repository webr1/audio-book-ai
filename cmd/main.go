package main

import (
	"flag"
	"log"

	container "audio-book-ai/cmd/app"
)

func main() {
	mode := flag.String("mode", "http", "run mode: http | task-worker")
	flag.Parse()

	switch *mode {
	case "http":
		app := container.InitHttpApp()
		app.Init()
		app.Start()
	case "task-worker":
		app := container.InitAsyncApp()
		app.Init()
		app.Start()
	default:
		log.Fatalf("unknown -mode %q, expected http or task-worker", *mode)
	}
}

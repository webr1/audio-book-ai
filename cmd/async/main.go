package main

import (
	container "audio-book-ai/cmd/app"
)

func main() {
	app := container.InitAsyncApp()
	app.Init()
	app.Start()
}

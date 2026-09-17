// @title Audio Book AI API
// @version 1.0
// @BasePath /api
package main

import (
	container "audio-book-ai/cmd/app"
)

func main() {
	app := container.InitHttpApp()
	app.Init()
	app.Start()
}

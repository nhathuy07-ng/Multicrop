package main

import (
	"embed"

	"log"

	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed all:frontend/dist
var assets embed.FS

func init() {

}

// main function serves as the application's entry point. It initializes the application, creates a window,
// and starts a goroutine that emits a time-based event every second. It subsequently runs the application and
// logs any error that might occur.
func main() {

	app := application.New(application.Options{
		Name:        "multicrop",
		Description: "A demo of using raw HTML & CSS",
		Services:    []application.Service{},
		Assets: application.AssetOptions{
			Handler: application.AssetFileServerFS(assets),
		},
		Mac: application.MacOptions{
			ApplicationShouldTerminateAfterLastWindowClosed: true,
		},
	})

	app.Window.NewWithOptions(application.WebviewWindowOptions{
		EnableFileDrop: true,
		Width:          1024,
		Height:         768,
		Title:          "Multicrop",
		URL:            "/",
	})

	err := app.Run()

	if err != nil {
		log.Fatal(err)
	}
}

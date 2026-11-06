package main

import (
	"embed"
	"fmt"
	"log/slog"

	"log"

	"github.com/wailsapp/wails/v3/pkg/application"
	"github.com/wailsapp/wails/v3/pkg/events"
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
		LogLevel: slog.LevelInfo,
	})

	mainWindow := app.Window.NewWithOptions(application.WebviewWindowOptions{
		EnableFileDrop: true,
		Width:          1024,
		Height:         768,
		Title:          "Multicrop",
		URL:            "/",
	})

	mainWindow.OnWindowEvent(events.Common.WindowFilesDropped, func(event *application.WindowEvent) {

		files := event.Context().DroppedFiles()
		for _, file := range files {
			fmt.Println("Dropped:", file)
		}
	})

	err := app.Run()

	if err != nil {
		log.Fatal(err)
	}
}

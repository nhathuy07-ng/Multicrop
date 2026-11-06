package main

import (
	"embed"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Start libvips in background
	vips.Startup(nil)

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "multicrop",
		Width:  850,
		Height: 600,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		OnStartup: app.startup,
		Bind: []interface{}{
			app,
		},

		DragAndDrop: &options.DragAndDrop{
			EnableFileDrop:     true,
			DisableWebViewDrop: false,
			CSSDropProperty:    "--wails-drop-target",
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

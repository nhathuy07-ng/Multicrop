package main

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// interface{} is an empty interface, meaning all data types can be stored
func onFileDropped(app *App, eventData []any) {
	file := eventData[2].([]string)[0]
	fmt.Println(file)
	filepath, err := GenerateWebpProxy(file)

	if err != nil {
		runtime.MessageDialog(app.ctx, runtime.MessageDialogOptions{
			Type:    runtime.ErrorDialog,
			Message: fmt.Sprintf("Error while opening file: %s", err),
		})
		return
	}

	runtime.MessageDialog(app.ctx, runtime.MessageDialogOptions{
		Type:    runtime.InfoDialog,
		Message: fmt.Sprintf("File parsed: %s", filepath),
	})
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {

	a.ctx = ctx
	runtime.EventsOn(ctx, "wails:file-drop", func(optionalData ...any) {
		runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
			Type:    runtime.InfoDialog,
			Message: "file dropped",
		})
		onFileDropped(a, optionalData)
	})
}

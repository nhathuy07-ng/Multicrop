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
// ...param means variadic parameters (much like python's *args)
func onFileDropped(eventData ...interface{}) {
	file := eventData[2].([]string)[0]
	fmt.Println(file)
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	runtime.EventsOn(ctx, "wails:file-drop", onFileDropped)
}

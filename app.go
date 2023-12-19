package main

import (
	"changeme/Global"
	"changeme/Settings"
	"context"
	"os"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	Global.Global_ConText = ctx
}

func (a *App) SetExportFileUrl() string {
	return Settings.SetExportFileUrl()
}

func (a *App) GetPath() string {
	str, _ := os.Getwd()
	return str
}

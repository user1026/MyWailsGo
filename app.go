package main

import (
	"MyWailsGo/Global"
	"MyWailsGo/Settings"
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
func (a *App) GetSettingsJson() map[string]string {
	return Settings.ReadSettings()
}
func (a *App) SetExportFileType(value string) bool {
	return Settings.SetExportFileType(value)
}
func (a *App) ExportFile(FileType string, data string) bool {
	return Settings.ExportFile(FileType, data)
}
func (a *App) ExportFile2(FileType string, data Global.PCHardwork) bool {
	return Settings.ExportFile2(FileType, data)
}

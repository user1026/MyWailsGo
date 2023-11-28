package main

import (
	"changeme/Computer"
	"changeme/Global"
	"changeme/Settings"
	"context"
	"fmt"
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

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetCpuInfo() Computer.CPUInfo {
	return Computer.GetCpuInfo()
}

func (a *App) GetUsingCpuInfo() float64 {
	return Computer.GetUsingCpuInfo()
}
func (a *App) GetRamInfo() Computer.RamInfo {
	return Computer.GetRamInfo()
}
func (a *App) SetExportFileUrl() string {
	return Settings.SetExportFileUrl()
}
func (a *App) GetCpuJSONData() interface{} {
	return Computer.GetCpuJSONData()
}
func openFile(data ...interface{}) {

}

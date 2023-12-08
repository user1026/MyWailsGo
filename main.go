package main

import (
	"changeme/AppMenu"
	"changeme/Computer"
	"changeme/Event"
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

// /go:embed all:frontend/dist
//
//go:embed frontend/dist/*
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	var event = Event.EventList
	// Create application with options
	err := wails.Run(&options.App{
		//Title:  "MyWailsGo",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Frameless:        false,
		MinWidth:         700,
		MinHeight:        700,
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		OnDomReady:       event.SetContext,
		Menu:             AppMenu.AllMenuList(),
		Bind: []interface{}{
			app,
			Computer.NewCPU(),
			Computer.NewPower(),
			Computer.NewRAM(),
			Computer.NewHDD(),
			Computer.NewGpu(),
			Computer.NewChassis(),
			Computer.NewRadiator(),
			Computer.NewSSD(),
			Computer.NewMainBoard(),
		},
	})

	//e.EventList()

	if err != nil {
		println("Error:", err.Error())
	}
}

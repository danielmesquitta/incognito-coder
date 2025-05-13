package main

import (
	root "github.com/danielmesquitta/incognito-coder"
	"github.com/danielmesquitta/incognito-coder/internal/app"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

func main() {
	app := app.New()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "incognito-coder",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: root.Assets,
		},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 70},
		OnStartup:        app.Run,
		Bind: []any{
			app,
		},
		AlwaysOnTop: true,
		Windows: &windows.Options{
			WindowIsTranslucent: true,
		},
		Linux: &linux.Options{
			WindowIsTranslucent: true,
		},
		Mac: &mac.Options{
			WindowIsTranslucent: true,
		},
	})

	if err != nil {
		panic(err.Error())
	}
}

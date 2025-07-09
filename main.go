package main

import (
	"SI/app"
	"SI/controller"
	"SI/repository"
	"SI/service"
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/build/client
var assets embed.FS

func main() {
	// Create an instance of the app structure
	appView := NewApp()

	//init database
	app.Connexion()
	db := app.DB

	// Dépendances
	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	ctrl := controller.NewController(db, repo, svc)

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "SI",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        appView.startup,
		Bind: []interface{}{
			appView,
			ctrl,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}

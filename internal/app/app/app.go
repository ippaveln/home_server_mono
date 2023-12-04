package app

import (
	"log/slog"
	"sync"

	"github.com/ippaveln/home_server_mono/internal/app/config"
	"github.com/ippaveln/home_server_mono/internal/services/connector"
)

func New(config *config.Config, log *slog.Logger, wg *sync.WaitGroup) *App {
	return &App{
		wg:        wg,
		log:       log,
		config:    config,
		connector: connector.New(log, wg),
		services:  listService(),
	}
}

func (app *App) Run() {
	app.log.Info("App starting")
	app.connector.Run()
	for _, service := range app.services {
		service.Run(app.config, app.connector, app.log, app.wg)
	}
	app.log.Info("App started")
}

func (app *App) Stop() {
	app.log.Info("App stopping")
	app.connector.Stop()
	for _, service := range app.services {
		service.Stop()
	}
	app.log.Info("App stopped")
}

package main

import (
	"log/slog"
	"os"
	"sync"

	"github.com/ippaveln/home_server_mono/internal/app/app"
	"github.com/ippaveln/home_server_mono/internal/app/config"
)

func main() {

	config := config.MustGetConfig()

	log := setupLogger(&config)

	wg := &sync.WaitGroup{}
	wg.Add(1)

	app := app.New(&config, log, wg)

	app.Run()

	wg.Wait()
	app.Stop()
}

func setupLogger(c *config.Config) *slog.Logger {
	var log *slog.Logger

	if c != nil && c.Debug {
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	} else {
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}

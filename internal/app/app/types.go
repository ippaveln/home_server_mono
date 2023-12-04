package app

import (
	"log/slog"
	"sync"

	"github.com/ippaveln/home_server_mono/internal/app/config"
	"github.com/ippaveln/home_server_mono/internal/services/connector"
	"github.com/ippaveln/home_server_mono/internal/services/service"
)

type App struct {
	wg        *sync.WaitGroup
	log       *slog.Logger
	config    *config.Config
	connector *connector.Connector
	services  []service.Service
}

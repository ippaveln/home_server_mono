package http_server

import (
	"log/slog"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/ippaveln/home_server_mono/internal/services/connector"
	"github.com/ippaveln/home_server_mono/internal/services/service"
)

type HttpServer struct {
	port      int
	status    service.Status
	log       *slog.Logger
	wg        *sync.WaitGroup
	connector *connector.Connector
	router    *chi.Mux
}

package service

import (
	"log/slog"
	"sync"

	"github.com/ippaveln/home_server_mono/internal/app/config"
	"github.com/ippaveln/home_server_mono/internal/services/connector"
)

type Status int

const (
	Init Status = iota
	Run
	Fail
	Stopping
	Stop
)

type Service interface {
	Run(config *config.Config, conn *connector.Connector, log *slog.Logger, wg *sync.WaitGroup)
	Stop()
	Status() Status
}

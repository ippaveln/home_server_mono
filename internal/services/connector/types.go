package connector

import (
	"log/slog"
	"net/http"
	"sync"
)

type Connector struct {
	log                  *slog.Logger
	wg                   *sync.WaitGroup
	GetStatusHa          func() (*http.Response, error)
	SendTextToTTSHa      func(string) error
	SendCmdToYaStationHa func(string) error
}

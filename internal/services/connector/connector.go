package connector

import (
	"log/slog"
	"sync"
)

func New(log *slog.Logger, wg *sync.WaitGroup) *Connector {
	return &Connector{
		log: log,
		wg:  wg,
	}
}

func (c *Connector) Run() {
	c.log.Info("connector starting")
	c.wg.Add(1)
	c.log.Info("connector started")
}

func (c *Connector) Stop() {
	c.log.Info("connector stopping")
	c.wg.Done()
	c.log.Info("connector stopped")
}

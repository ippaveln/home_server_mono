package tgBot

import (
	"log/slog"
	"sync"

	"github.com/ippaveln/home_server_mono/internal/services/connector"
	"github.com/ippaveln/home_server_mono/internal/services/service"
	tele "gopkg.in/telebot.v3"
)

type TgBot struct {
	token      string
	bot        *tele.Bot
	log        *slog.Logger
	status     service.Status
	wg         *sync.WaitGroup
	ippavelnID int64
	connector  *connector.Connector
}

package app

import (
	home_assistant "github.com/ippaveln/home_server_mono/internal/services/homeassistant"
	http_server "github.com/ippaveln/home_server_mono/internal/services/httpserver"
	"github.com/ippaveln/home_server_mono/internal/services/service"
	"github.com/ippaveln/home_server_mono/internal/services/tgBot"
)

func listService() []service.Service {
	return []service.Service{
		&tgBot.TgBot{},
		&home_assistant.HomeAssistantClient{},
		&http_server.HttpServer{},
	}
}

package app

import (
	"github.com/ippaveln/home_server_mono/internal/services/home_assistant"
	"github.com/ippaveln/home_server_mono/internal/services/service"
	"github.com/ippaveln/home_server_mono/internal/services/tgBot"
)

func listService() []service.Service {
	return []service.Service{
		&tgBot.TgBot{},
		&home_assistant.HomeAssistantClient{},
	}
}

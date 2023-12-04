package home_assistant

import (
	"net/http"

	"github.com/ippaveln/home_server_mono/internal/services/service"
)

type HomeAssistantClient struct {
	port   int
	host   string
	token  string
	status service.Status
	client http.Client
}

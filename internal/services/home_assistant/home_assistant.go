package home_assistant

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"

	"github.com/ippaveln/home_server_mono/internal/app/config"
	"github.com/ippaveln/home_server_mono/internal/services/connector"
	"github.com/ippaveln/home_server_mono/internal/services/service"
)

func (ha *HomeAssistantClient) Run(config *config.Config, conn *connector.Connector, log *slog.Logger, wg *sync.WaitGroup) {
	ha.port = 8123
	ha.host = "192.168.0.99"
	ha.token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiI2MzA0MzE3MDJiNmY0ZjI5YTU0MzJhZTQyNzZmYTFiZSIsImlhdCI6MTY5ODU4MTEwMiwiZXhwIjoyMDEzOTQxMTAyfQ.9E6chjLPCwbpBrzFXIuCB_RtFCCy7XII8f3WxrQpVxU"
	ha.client = http.Client{Timeout: 6 * time.Second}

	conn.GetStatusHa = ha.GetStatus
	conn.SendTextToTTSHa = ha.SendTextSpeechYaStation
	conn.SendCmdToYaStationHa = ha.SendCommandYaStation

}

func (ha *HomeAssistantClient) Status() service.Status {
	return ha.status
}

func (ha *HomeAssistantClient) Stop() {}

func (ha *HomeAssistantClient) GetStatus() (*http.Response, error) {
	u := url.URL{
		Scheme: "http",
		Host:   ha.host + ":" + strconv.Itoa(ha.port),
		Path:   "api/",
	}
	req, _ := http.NewRequest(
		http.MethodGet,
		u.String(),
		nil,
	)
	req.Header.Add("Authorization", "Bearer "+ha.token)

	return ha.client.Do(req)
}

func (ha *HomeAssistantClient) SendTextSpeechYaStation(text string) error {
	return ha.sendDataToYaStation(text, "text")
}

func (ha *HomeAssistantClient) SendCommandYaStation(cmd string) error {
	return ha.sendDataToYaStation(cmd, "command")
}

func (ha *HomeAssistantClient) sendDataToYaStation(data, typeAction string) error {
	u := url.URL{
		Scheme: "http",
		Host:   ha.host + ":" + strconv.Itoa(ha.port),
		Path:   "api/services/media_player/play_media",
	}
	b, err := json.Marshal(buildYaStationRequest(data, typeAction))
	if err != nil {
		return err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		u.String(),
		bytes.NewReader(b),
	)

	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "Bearer "+ha.token)
	req.Header.Add("Content-Type", "application/json")

	resp, err := ha.client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status response not ok, this is %s", resp.Status)
	}
	return nil
}

type yandexStationRequest struct {
	EntityID         string `json:"entity_id"`
	MediaContentId   string `json:"media_content_id"`
	MediaContentType string `json:"media_content_type"`
}

func buildYaStationRequest(data, typeData string) *yandexStationRequest {
	return &yandexStationRequest{
		EntityID:         "media_player.yandex_station_742078e2880c08040810",
		MediaContentId:   data,
		MediaContentType: typeData,
	}
}

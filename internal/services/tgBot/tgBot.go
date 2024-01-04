package tgBot

import (
	"errors"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"sync"
	"time"

	tele "gopkg.in/telebot.v3"
	"gopkg.in/telebot.v3/middleware"

	"github.com/ippaveln/home_server_mono/internal/app/config"
	"github.com/ippaveln/home_server_mono/internal/services/connector"
	"github.com/ippaveln/home_server_mono/internal/services/service"
)

func (bot *TgBot) Run(config *config.Config, conn *connector.Connector, log *slog.Logger, wg *sync.WaitGroup) {
	wg.Add(1)
	if wg == nil {
		panic("waitGroup is nil!")
	}
	bot.wg = wg
	bot.connector = conn
	log.Info("telegram bot starting")
	if bot == nil {
		panic("telegram bot is nil")
	}
	bot.status = service.Init
	bot.log = log

	bot.token = config.TgBot.Token

	pref := tele.Settings{
		Token:  bot.token,
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)

	if err != nil {
		log.Error("telegram bot don't create , err: %w", err)
		bot.status = service.Fail
		panic("telegram bot creating fail")
	}
	bot.bot = b

	levelBotDebug := slog.LevelInfo
	if config.TgBot.Debug {
		levelBotDebug = slog.LevelDebug
	}

	handerLog := log.Handler()
	oldLog := slog.NewLogLogger(handerLog, levelBotDebug)
	bot.bot.Use(middleware.Logger(oldLog))
	bot.ippavelnID = config.TgBot.IppavelnID

	bot.startHandlers()

	go bot.bot.Start()

	bot.status = service.Run
	log.Info("telegram bot started")
	wg.Done()
}

func (bot *TgBot) Stop() {
	bot.bot.Stop()
	bot.wg.Done()
}

func (bot *TgBot) Status() service.Status {
	return bot.status
}

func (bot *TgBot) startHandlers() {
	bot.bot.Handle("/ha", func(ctx tele.Context) error {
		resp, err := bot.connector.GetStatusHa()
		if err != nil || resp.StatusCode != http.StatusOK {
			return ctx.Reply(strconv.Itoa(int(ctx.Sender().ID)))
		}

		defer resp.Body.Close()
		// data
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return ctx.Reply(strconv.Itoa(int(ctx.Sender().ID)) + err.Error())
		}

		return ctx.Reply(string(bodyBytes))
	},
		middleware.Whitelist(bot.ippavelnID),
	)
	bot.bot.Handle(tele.OnText, func(ctx tele.Context) error {
		return ctx.Reply(ctx.Text())
	})

	bot.bot.Handle("/ha_tts", func(c tele.Context) error {
		payload := c.Message().Payload
		if payload == "" {
			return errors.New("empty payload")
		}
		err := bot.connector.SendTextToTTSHa(payload)
		if err != nil {
			return err
		}
		return c.Reply("ok")
	})

	bot.bot.Handle("/ha_cmd", func(c tele.Context) error {
		payload := c.Message().Payload
		if payload == "" {
			return errors.New("empty payload")
		}
		err := bot.connector.SendCmdToYaStationHa(payload)
		if err != nil {
			return err
		}
		return c.Reply("ok")
	})

}

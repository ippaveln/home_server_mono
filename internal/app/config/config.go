package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func MustGetConfig() Config {
	loadEnv()

	return parseEnv()
}

func loadEnv() {
	err := godotenv.Load(".env.secrets")
	if err != nil {
		log.Panicf("error loading secrets, err: %s", err.Error())
	}
}

func parseEnv() Config {
	res := Config{}

	parseCommon(&res)
	parseTgBot(&res)
	parseServer(&res)

	return res
}

func parseCommon(config *Config) {
	config.Debug = os.Getenv("DEBUG") == "true"
}

func parseTgBot(config *Config) {
	config.TgBot.Token = os.Getenv("BOT_TOKEN")
	config.TgBot.Debug = os.Getenv("BOT_DEBUG") == "true"
	i, err := strconv.Atoi(os.Getenv("TG_ID_IPPAVELN"))
	if err != nil {
		log.Panicf("error parse tg id ippaveln, err: %s: ", err.Error())
	}
	config.TgBot.IppavelnID = int64(i)
}

func parseServer(config *Config) {
	config.Server.Debug = os.Getenv("SERVER_DEBUG") == "true"

	i, err := strconv.Atoi(os.Getenv("SERVER_HTTP_PORT"))
	if err != nil {
		log.Panicf("error parse http port, err: %s", err.Error())
	}
	config.Server.Ports.Http = int64(i)

	i, err = strconv.Atoi(os.Getenv("SERVER_GRPC_PORT"))
	if err != nil {
		log.Panicf("error parse grpc port, err: %s", err.Error())
	}
	config.Server.Ports.Grpc = int64(i)

	i, err = strconv.Atoi(os.Getenv("SERVER_SWAGGER_PORT"))
	if err != nil {
		log.Panicf("error parse swagger port, err: %s", err.Error())
	}
	config.Server.Ports.Swagger = int64(i)
}

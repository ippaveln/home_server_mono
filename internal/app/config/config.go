package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func GetConfig() Config {
	loadEnv()

	return parseEnv()
}

func loadEnv() {
	err := godotenv.Load(".env.secrets")
	if err != nil {
		log.Panicf("error loading secrets, err: %w", err)
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
}

func parseServer(config *Config) {
	config.Server.Debug = os.Getenv("SERVER_DEBUG") == "true"

	i, err := strconv.Atoi(os.Getenv("SERVER_HTTP_PORT"))
	if err != nil {
		log.Panicf("error parse http port, err: %w", err)
	}
	config.Server.Ports.Http = int64(i)

	i, err = strconv.Atoi(os.Getenv("SERVER_GRPC_PORT"))
	if err != nil {
		log.Panicf("error parse grpc port, err: %w", err)
	}
	config.Server.Ports.Grpc = int64(i)

	i, err = strconv.Atoi(os.Getenv("SERVER_SWAGGER_PORT"))
	if err != nil {
		log.Panicf("error parse swagger port, err: %w", err)
	}
	config.Server.Ports.Swagger = int64(i)
}

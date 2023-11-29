package main

import (
	"log"

	"github.com/ippaveln/home_server_mono/internal/app/config"
)

func main() {

	config := config.GetConfig()
	log.Println(config)

	// TODO: init logger

	// TODO: init app

	// TODO: init http server

	// TODO: init grpc server

}

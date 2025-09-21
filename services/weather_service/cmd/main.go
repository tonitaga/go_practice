package main

import (
	"log"

	"github.com/tonitaga/weather_service/internal/config"
	"github.com/tonitaga/weather_service/internal/server"
)

func main() {
	server := server.NewServer(config.NewConfig())
	if err := server.Run(); err != nil {
		log.Fatalln(err)
	}
}

package app

import (
	"flag"

	"github.com/tonitaga/some_service/internal/config"
	"github.com/tonitaga/some_service/internal/database"
	"github.com/tonitaga/some_service/internal/database/postgres"
)

var (
	configPath *string = flag.String("config", "config/config.json", "path to the application config file")
)

func HandleStorage(storage database.Storage) {
	storage.AutoMigrate()
}

func Run() {
	flag.Parse()

	config := config.NewConfig(*configPath)
	storage, err := postgres.NewPostgresStorage(config.Postgres)
	if err == nil {
		HandleStorage(storage)
	}
}

package app

import (
	"flag"
	"fmt"

	"github.com/tonitaga/some_service/internal/config"
)

var (
	configPath *string = flag.String("config", "config/config.json", "path to the application config file")
)

func Run() {
	flag.Parse()

	config := config.NewConfig(*configPath)
	fmt.Println(*config)
}

package config

import (
	"flag"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Host   string
	Port   int
	ApiKey string
}

func NewConfig() *Config {
	config := &Config{}
	config.parse()

	return config
}

func (c *Config) parse() {
	flag.StringVar(&c.ApiKey, "api", "", "API key for requests to http://api.openweathermap.org")

	flag.StringVar(&c.Host, "host", "localhost", "listening host")
	flag.IntVar(&c.Port, "port", 8080, "listening port")
	flag.Parse()

	if c.ApiKey == "" {
		err := godotenv.Load()
		if err == nil {
			c.ApiKey = os.Getenv("API_KEY")
		}
	}
}

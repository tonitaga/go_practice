package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Postgres PostgresConfig `json:"postgres"`
}

func NewDefault() *Config {
	return &Config{
		Postgres: PostgresConfig{
			Host:     "localhost",
			Port:     5432,
			Name:     "postgres",
			User:     "postgres",
			Password: "postgres",
		},
	}
}

func NewConfig(path string) *Config {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	config := &Config{}
	if err := json.NewDecoder(file).Decode(config); err != nil {
		panic(err)
	}

	return config
}

type PostgresConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
}

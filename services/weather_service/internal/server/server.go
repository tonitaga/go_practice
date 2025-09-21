package server

import (
	"fmt"
	"net/http"

	"github.com/tonitaga/weather_service/internal/config"
)

type server struct {
	config      *config.Config
	httpHandler *http.ServeMux
}

func NewServer(config *config.Config) *server {
	return &server{
		config:      config,
		httpHandler: &http.ServeMux{},
	}
}

func (s *server) Run() error {
	s.initHandlers()

	address := fmt.Sprintf("%s:%d", s.config.Host, s.config.Port)
	return http.ListenAndServe(address, s.httpHandler)
}

func (s *server) initHandlers() {
	s.httpHandler.HandleFunc("GET /", s.executeTemplateIndex)
	s.httpHandler.HandleFunc("POST /api/weather", s.handleWeatherRequest)
}

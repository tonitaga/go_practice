package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tonitaga/weather_service/internal/config"
)

type server struct {
	engine      *gin.Engine
	config      *config.Config
	httpHandler *http.ServeMux
}

func NewServer(config *config.Config) *server {
	return &server{
		engine:      gin.Default(),
		config:      config,
		httpHandler: &http.ServeMux{},
	}
}

func (s *server) Run() error {
	s.initHandlers()
	return s.engine.Run(fmt.Sprintf("%s:%d", s.config.Host, s.config.Port))
}

func (s *server) initHandlers() {
	s.engine.Static("/", "template")
	s.engine.POST("/api/weather", s.getWeahterHandler)
}

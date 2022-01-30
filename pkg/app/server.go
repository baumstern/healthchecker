package app

import (
	"healthchecker/pkg/api"
	"healthchecker/pkg/config"

	"github.com/labstack/echo"
)

type Server struct {
	cfg          *config.Config
	router       *echo.Echo
	watchService api.WatchService
}

func NewServer(cfg *config.Config, watchService api.WatchService) *Server {
	return &Server{
		cfg:          cfg,
		router:       echo.New(),
		watchService: watchService}
}

func (s *Server) Start() {
	s.Routes()

	// TODO: make config
	port := "8080"
	if s.cfg.Server.Port != "" {
		port = s.cfg.Server.Port
	}

	s.router.Logger.Fatal((s.router.Start((":" + port))))
}

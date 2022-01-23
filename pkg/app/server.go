package app

import (
	"errors"
	"healthchecker/pkg/api"
	"healthchecker/pkg/config"
	"log"
	"net/http"
)

type Server struct {
	cfg          *config.Config
	watchService api.WatchService
}

func NewServer(cfg *config.Config, watchService api.WatchService) *Server {
	return &Server{
		cfg:          cfg,
		watchService: watchService}
}

func (s *Server) Start() {
	s.Routes()

	// TODO: make config
	port := "8080"
	if s.cfg.Server.Port != "" {
		port = s.cfg.Server.Port
	}

	if err := http.ListenAndServe(":"+port, nil); errors.Is(err, http.ErrServerClosed) {
		log.Fatalln("Web server has shut down")
	} else {
		log.Fatalln("Web server has shut down unexpectedly")
	}
}

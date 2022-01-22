package main

import (
	"healthcheck/pkg/api"
	"healthcheck/pkg/app"
	"healthcheck/pkg/config"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Println("failed to load config file:", err)
	}

	watchService := api.NewWatchService(cfg)
	err = watchService.Start("klaytn")
	if err != nil {
		log.Fatalln("failed to start watch")
	}

	s := app.NewServer(cfg, watchService)
	s.Start()
}

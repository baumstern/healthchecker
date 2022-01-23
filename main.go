package main

import (
	"healthchecker/pkg/api"
	"healthchecker/pkg/app"
	"healthchecker/pkg/config"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Println("failed to load config file:", err)
		cfg = &config.Config{}
	}

	watchService := api.NewWatchService(cfg)
	err = watchService.Start("klaytn")
	if err != nil {
		log.Fatalln("failed to start watch")
	}
	err = watchService.Start("ethereum")
	if err != nil {
		log.Fatalln("failed to start watch")
	}

	s := app.NewServer(cfg, watchService)
	s.Start()
}

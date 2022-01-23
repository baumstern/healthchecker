package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port string `yaml:"port"`
	}
	Ethereum struct {
		ApiKey        string `yaml:"api_key"`
		WatchInterval int    `yaml:"watch_interval"`
	}
	Klaytn struct {
		AccessToken   string `yaml:"access_token"`
		WatchInterval int    `yaml:"watch_interval"`
	}
}

func LoadConfig() (*Config, error) {
	f, err := os.Open("config.yaml")
	if err != nil {
		log.Println("failed to open config.yaml file:", err)
		return nil, err
	}
	defer f.Close()

	var cfg Config
	err = yaml.NewDecoder(f).Decode(&cfg)
	if err != nil {
		log.Fatalln("failed to decode config.yaml file:", err)
		return nil, err
	}
	return &cfg, nil
}

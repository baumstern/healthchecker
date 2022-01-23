package config

import (
	"errors"
	"log"
	"os"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

// TODO: check watch interval keyword is correct by lookup reference code(e.g. Kubernetess)
type Config struct {
	Server struct {
		Port string `yaml:"port" envconfig:"port" default:"8080"`
	}
	Ethereum struct {
		ApiKey        string `yaml:"api_key" envconfig:"api_key"`
		WatchInterval int    `yaml:"watch_interval" envconfig:"watch_interval" default:"7"`
	}
	Klaytn struct {
		AccessToken   string `yaml:"access_token" envconfig:"access_token"`
		WatchInterval int    `yaml:"watch_interval" envconfig:"watch_interval" default:"1"`
	}
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}

	err := cfg.loadFromYaml()
	if err != nil {
		log.Println("failed to get config from yaml file:", err)
	}
	err = cfg.loadFromEnv()
	if err != nil {
		return nil, err
	}

	err = cfg.checkApiTokens()
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) loadFromYaml() error {
	f, err := os.Open("config.yaml")
	if err != nil {
		return err
	}
	defer f.Close()

	err = yaml.NewDecoder(f).Decode(c)
	if err != nil {
		return err
	}
	return nil
}

func (c *Config) loadFromEnv() error {
	err := envconfig.Process("healthchecker", c)
	if err != nil {
		log.Fatalln("failed to get configuration:", err)
		return err
	}

	return nil
}

func (c *Config) checkApiTokens() error {
	if c.Ethereum.ApiKey == "" || c.Klaytn.AccessToken == "" {
		return errors.New("couldn't get blockchain status: API token didn't provided")
	}
	return nil
}

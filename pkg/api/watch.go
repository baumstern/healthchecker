package api

import (
	"fmt"
	"healthcheck/pkg/collector"
	"healthcheck/pkg/collector/klaytn"
	"healthcheck/pkg/config"
)

type cancel func()

type klaytnClient struct {
	client      klaytn.Client
	cancelWatch cancel
}

type WatchService interface {
	Start(blockchainType string) error
	Stop(blockchainType string) error
	GetLatestBlock(blockchainType string) (*collector.LatestBlock, error)
}

type watchService struct {
	klaytn klaytnClient
}

func NewWatchService(cfg *config.Config) WatchService {
	return &watchService{klaytn: klaytnClient{*klaytn.NewClient(cfg.Klaytn.AccessToken), nil}}
}

func (w *watchService) Start(blockchainType string) error {
	switch blockchainType {
	case "klaytn":
		w.klaytn.cancelWatch = w.klaytn.client.Watch()
		return nil
	default:
		return fmt.Errorf("failed to start watching blockchain: %s is not supported", blockchainType)
	}
}

func (w *watchService) Stop(blockchainType string) error {
	switch blockchainType {
	case "klaytn":
		w.klaytn.cancelWatch()
		return nil
	default:
		return fmt.Errorf("failed to stop watching blockchain: %s is not supported", blockchainType)
	}
}

func (w *watchService) GetLatestBlock(blockchainType string) (*collector.LatestBlock, error) {
	switch blockchainType {
	case "klaytn":
		return w.klaytn.client.GetLatestBlock(), nil
	default:
		return nil, fmt.Errorf("failed to get latest block number: %s is not supported", blockchainType)
	}
}

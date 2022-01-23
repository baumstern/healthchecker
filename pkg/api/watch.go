package api

import (
	"fmt"
	"healthchecker/pkg/collector"
	"healthchecker/pkg/collector/ethereum"
	"healthchecker/pkg/collector/klaytn"
	"healthchecker/pkg/config"
)

type cancel func()

type klaytnClient struct {
	client      klaytn.Client
	cancelWatch cancel
}

type ethereumClient struct {
	client      ethereum.Client
	cancelWatch cancel
}

type WatchService interface {
	Start(blockchainType string) error
	Stop(blockchainType string) error
	GetLatestBlock(blockchainType string) (*collector.LatestBlock, error)
}

type watchService struct {
	ethereum ethereumClient
	klaytn   klaytnClient
}

func NewWatchService(cfg *config.Config) WatchService {
	return &watchService{
		ethereum: ethereumClient{*ethereum.NewClient(cfg), nil},
		klaytn:   klaytnClient{*klaytn.NewClient(cfg), nil},
	}
}

func (w *watchService) Start(blockchainType string) error {
	switch blockchainType {
	case "ethereum":
		w.ethereum.cancelWatch = w.ethereum.client.Watch()
		return nil
	case "klaytn":
		w.klaytn.cancelWatch = w.klaytn.client.Watch()
		return nil
	default:
		return fmt.Errorf("failed to start watching blockchain: %s is not supported", blockchainType)
	}
}

func (w *watchService) Stop(blockchainType string) error {
	switch blockchainType {
	case "ethereum":
		w.ethereum.cancelWatch()
		return nil
	case "klaytn":
		w.klaytn.cancelWatch()
		return nil
	default:
		return fmt.Errorf("failed to stop watching blockchain: %s is not supported", blockchainType)
	}
}

func (w *watchService) GetLatestBlock(blockchainType string) (*collector.LatestBlock, error) {
	switch blockchainType {
	case "ethereum":
		return w.ethereum.client.GetLatestBlock(), nil
	case "klaytn":
		return w.klaytn.client.GetLatestBlock(), nil
	default:
		return nil, fmt.Errorf("failed to get latest block number: %s is not supported", blockchainType)
	}
}

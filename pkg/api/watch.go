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
	isWatching  bool
}

type ethereumClient struct {
	client      ethereum.Client
	cancelWatch cancel
	isWatching  bool
}

type WatchService interface {
	Start(blockchainType string) error
	Stop(blockchainType string) error
	GetLatestBlock(blockchainType string) (*collector.LatestBlock, error)
	IsWatching(blockchainType string) (bool, error)
}

type watchService struct {
	ethereum ethereumClient
	klaytn   klaytnClient
}

func NewWatchService(cfg *config.Config) WatchService {
	return &watchService{
		ethereum: ethereumClient{*ethereum.NewClient(cfg), nil, false},
		klaytn:   klaytnClient{*klaytn.NewClient(cfg), nil, false},
	}
}

func (w *watchService) Start(blockchainType string) error {
	switch blockchainType {
	case "ethereum":
		if !w.ethereum.isWatching {
			w.ethereum.cancelWatch = w.ethereum.client.Watch()
			w.ethereum.isWatching = true
		}
		return nil

	case "klaytn":
		if !w.klaytn.isWatching {
			w.klaytn.cancelWatch = w.klaytn.client.Watch()
			w.klaytn.isWatching = true
		}

		return nil

	default:
		return fmt.Errorf("failed to start watching blockchain: %s is not supported", blockchainType)
	}
}

func (w *watchService) Stop(blockchainType string) error {
	switch blockchainType {
	case "ethereum":
		if w.ethereum.isWatching {
			w.ethereum.cancelWatch()
			w.ethereum.isWatching = false
		}

		return nil
	case "klaytn":
		if w.klaytn.isWatching {
			w.klaytn.cancelWatch()
			w.klaytn.isWatching = false
		}

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

func (w *watchService) IsWatching(blockchainType string) (bool, error) {
	switch blockchainType {
	case "ethereum":
		return w.ethereum.isWatching, nil
	case "klaytn":
		return w.klaytn.isWatching, nil
	default:
		return false, fmt.Errorf("failed to get watch status: %s is not supported", blockchainType)
	}
}

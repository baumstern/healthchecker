package ethereum

import (
	"bytes"
	"encoding/json"
	"fmt"
	"healthchecker/pkg/collector"
	"healthchecker/pkg/config"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type ResponseBlockNumber struct {
	JsonRpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Result  string `json:"result"`
}

type Client struct {
	latestBlock   *collector.LatestBlock
	apiKey        string
	watchInterval int
}

func NewClient(cfg *config.Config) *Client {
	if cfg.Ethereum.ApiKey == "" {
		log.Fatalln("API key for Ethereum is not provied")
	}
	return &Client{
		latestBlock:   &collector.LatestBlock{},
		apiKey:        cfg.Ethereum.ApiKey,
		watchInterval: cfg.Ethereum.WatchInterval,
	}
}

func (c *Client) GetLatestBlock() *collector.LatestBlock {
	return c.latestBlock
}

func (c *Client) Watch() func() {
	done := make(chan struct{})
	cancel := func() {
		close(done)
	}

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("break")
				return
			default:
				latestBlockNum, err := c.getLatestBlock()
				if err != nil {
					log.Println("failed to check latest block:", err)
				}

				c.latestBlock.Num = latestBlockNum
				c.latestBlock.Timestamp = time.Now()

				time.Sleep(time.Duration(c.watchInterval) * time.Second)
			}
		}
	}()
	return cancel
}

func (c *Client) getLatestBlock() (int64, error) {
	// "https://api.opensea.io/api/v1/assets?owner=0xdfccbfe7ddce6878ea4542cc343be1ce4ed4516e&order_direction=desc&offset=0&limit=20"
	baseURL := "https://eth-mainnet.alchemyapi.io/v2/"

	body := "{\"jsonrpc\":\"2.0\",\"method\":\"eth_blockNumber\",\"params\":[],\"id\":1}"

	req, err := http.NewRequest(http.MethodPost, baseURL+c.apiKey, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return -1, err
	}
	req.Header = map[string][]string{
		"Content-Type": {"application/json"},
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return -1, err
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return -1, err
	}

	block := &ResponseBlockNumber{}
	err = json.Unmarshal(data, block)
	if err != nil {
		return -1, err
	}

	result, err := strconv.ParseInt(block.Result[2:], 16, 64)
	if err != nil {
		return -1, err
	}

	return result, nil
}

package klaytn

import (
	"bytes"
	"encoding/json"
	"fmt"
	"healthcheck/pkg/collector"
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
	latestBlock *collector.LatestBlock
	accessToken string
}

func NewClient(accessToken string) *Client {
	return &Client{
		latestBlock: &collector.LatestBlock{},
		accessToken: accessToken,
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

				time.Sleep(1 * time.Second)
			}
		}
	}()
	return cancel
}

func (c *Client) getLatestBlock() (int64, error) {
	// "https://api.opensea.io/api/v1/assets?owner=0xdfccbfe7ddce6878ea4542cc343be1ce4ed4516e&order_direction=desc&offset=0&limit=20"
	baseURL := "https://node-api.klaytnapi.com/v1/klaytn"

	body := "{\"jsonrpc\":\"2.0\",\"method\":\"klay_blockNumber\",\"params\":[],\"id\":1}"

	req, err := http.NewRequest(http.MethodPost, baseURL, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return -1, err
	}
	req.Header = map[string][]string{
		"x-chain-id":    {"8217"}, // 8217 is id of Cypress network(mainnet)
		"Content-Type":  {"application/json"},
		"Authorization": {"Basic " + c.accessToken},
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

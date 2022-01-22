package collector

import "time"

type Collector interface {
	GetLatestBlock() LatestBlock
}

type LatestBlock struct {
	Num       int64     `json:"block_num"`
	Timestamp time.Time `json:"timestamp"`
}

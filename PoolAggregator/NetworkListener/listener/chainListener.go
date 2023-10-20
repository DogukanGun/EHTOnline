package listener

import (
	"PoolAggregator/NetworkListener/yield"
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

type NetworkListener struct {
	url                     string
	contractAddress         common.Address
	crossChainUrl           string
	crossChainTargetAddress string
}

func (c *NetworkListener) listenNetwork() {
	client, err := ethclient.Dial(c.url)
	if err != nil {
		log.Fatal(err)
	}
	query := ethereum.FilterQuery{
		Addresses: []common.Address{c.contractAddress},
	}
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			if vLog.Topics[0] == common.HexToHash("0xe633b43bc0097e50fdbee62a7bb46bfeaeb94cb9ed08f0e549a8160fa1d197f9") {
				yield.StartYieldAggregator()
			}
		}
	}
}

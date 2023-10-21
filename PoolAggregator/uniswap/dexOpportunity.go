package uniswap

import (
	"PoolAggregator/data/response"
	"PoolAggregator/uniswap/oracle"
	"PoolAggregator/utils"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
)

func DexOpportunity(token0 string, token1 string) (res response.UniswapPriceResponse) {

	// Access the first and second arguments.
	envNames := []string{"ETH_URL"}
	status, _, envMap := utils.InitializeENV(envNames, "dex.env")
	if !status {
		fmt.Println("Error in env")
	}
	client, err := ethclient.Dial(envMap["ETH_URL"])
	if err != nil {
		fmt.Println(fmt.Sprintf("Error in client connection: %s", err))
	}
	price, _, _, err := UniswapV3PriceOracle(token0, token1, []int64{500, 3000, 5000}, client)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error in price calculation: %s", err))
	}
	responseForToken0 := oracle.ChainlinkETHUSDOracle(token0, client)
	responseForToken1 := oracle.ChainlinkETHUSDOracle(token1, client)
	priceForToken0 := responseForToken0.Price
	priceForToken1 := responseForToken1.Price
	res.Price = price * priceForToken1
	if priceForToken0 < res.Price {
		res.ShouldSell = true
		res.Advantage = (res.Price - priceForToken0) / 100
		fmt.Println(fmt.Sprintf("Oppurtinaty to sell %s", token0))
	} else {
		res.ShouldSell = false
		res.Advantage = (priceForToken0 - res.Price) / 100
		fmt.Println(fmt.Sprintf("Oppurtinaty to buy %s", token0))
	}
	fmt.Println(fmt.Sprintf("%.7f", res.Price))
	return
}

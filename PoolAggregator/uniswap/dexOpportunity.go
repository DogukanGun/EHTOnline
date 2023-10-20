package uniswap

import (
	"PoolAggregator/data/response"
	"PoolAggregator/uniswap/oracle"
	"PoolAggregator/utils"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"os"
)

func DexOpportunity(token0 string, token1 string) (res response.UniswapPriceResponse) {
	if len(os.Args) < 3 {
		fmt.Println("Please provide at least two arguments.")
		return
	}

	// Access the first and second arguments.
	envNames := []string{"ETH_URL"}
	status, _, envMap := utils.InitializeENV(envNames, ".env")
	if !status {
		fmt.Println("Error in env")
		os.Exit(1)
	}
	client, err := ethclient.Dial(envMap["ETH_URL"])

	if err != nil {
		fmt.Println(fmt.Sprintf("Error in client connection: %s", err))
		os.Exit(1)
	}
	price, _, _, err := UniswapV3PriceOracle(token0, token1, []int64{500, 3000, 5000}, client)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error in price calculation: %s", err))
		os.Exit(1)
	}
	responseForToken0 := oracle.ChainlinkETHUSDOracle(token0, client)
	responseForToken1 := oracle.ChainlinkETHUSDOracle(token0, client)
	priceForToken0 := responseForToken0.Price
	priceForToken1 := responseForToken1.Price
	priceFromPool := (price * priceForToken0) - priceForToken1
	if priceForToken0 < priceFromPool {
		res.ShouldSell = true
		fmt.Println(fmt.Sprintf("Oppurtinaty to sell %s", token0))
	} else {
		res.ShouldSell = false
		fmt.Println(fmt.Sprintf("Oppurtinaty to buy %s", token0))
	}
	fmt.Println(fmt.Sprintf("%.7f", (priceFromPool-priceForToken0)/100))
	res.Price = (priceFromPool - priceForToken0) / 100
	return
}

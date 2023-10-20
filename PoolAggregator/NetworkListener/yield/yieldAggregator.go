package yield

import (
	"PoolAggregator/NetworkListener/yield/uniswap"
	"PoolAggregator/data"
	"PoolAggregator/utils"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"os"
	"strings"
)

func StartYieldAggregator() {
	//TODO Send contracts a request to get is there a stake there
	//If stake is in chain a then check chain b on opposite situation check chain b
	//Get current stake
	envNames := []string{"ETH_URL"}
	status, _, envMap := utils.InitializeENV(envNames, ".env")
	if !status {
		fmt.Println("Error in env")
		os.Exit(1)
	}
	client, _ := ethclient.Dial(envMap["ETH_URL"])
	chain := "binance"
	usdtAddress := "0xdac17f958d2ee523a2206206994597c13d831ec7"
	var yields []data.YieldCalculation
	if chain == "binance" {
		for key, _ := range supportedCoins {
			yields = append(yields, uniswap.GetUniswapYield(key, usdtAddress, client))
		}
	} else if chain == "goerli" {
		//TODO @CemDenizsel add pancake yield calculation here
	}
	maximumYield := getMaxYield(yields)
	fmt.Println("Max yield in ", maximumYield.ChainName, " and adrees is ", maximumYield.PoolAddress)
	//TODO Send other contract this info and make the comparison in the smart contract
}

func getMaxYield(yields []data.YieldCalculation) data.YieldCalculation {
	max := yields[0]
	for _, num := range yields {
		if num.YieldValue.Cmp(max.YieldValue) > 0 {
			max = num
		}
	}
	return max
}

var supportedCoins = map[string]string{
	// WETH
	strings.ToLower("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"): strings.ToLower("0x5f4ec3df9cbd43714fe2740f5e3616155c5b8419"),
	// WBTC
	strings.ToLower("0x2260fac5e5542a773aa44fbcfedf7c193bc2c599"): strings.ToLower("0xf4030086522a5beea4988f8ca5b36dbc97bee88c"),
	// LINK
	strings.ToLower("0x514910771af9ca656af840dff83e8264ecf986ca"): strings.ToLower("0x2c1d072e956affc0d435cb7ac38ef18d24d9127c"),
	// UNI
	strings.ToLower("0x1f9840a85d5af5bf1d1762f925bdaddc4201f984"): strings.ToLower("0x553303d460ee0afb37edff9be42922d8ff63220e"),
}

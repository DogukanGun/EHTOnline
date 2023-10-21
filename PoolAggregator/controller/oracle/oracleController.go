package oracleController

import (
	"PoolAggregator/uniswap/ChainlinkAggregator"
	"PoolAggregator/uniswap/oracle"
	"PoolAggregator/utils"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func Handler() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Access the first and second arguments.
		envNames := []string{"ETH_URL"}
		status, _, envMap := utils.InitializeENV(envNames, "dex.env")
		if !status {
			fmt.Println("Error in env")
			os.Exit(1)
		}
		client, err := ethclient.Dial(envMap["ETH_URL"])
		if err != nil {
			context.JSON(http.StatusNotFound, map[string]interface{}{"error": true, "message": err})
		}
		var responses []oracle.PriceOracleResponse
		for key, _ := range ChainlinkAggregator.CHAINLINK_ETH_PRICE_FEEDS {
			res := oracle.ChainlinkETHUSDOracle(key, client)
			responses = append(responses, res)
		}
		context.JSON(http.StatusOK, map[string]interface{}{"data": responses})
	}
}

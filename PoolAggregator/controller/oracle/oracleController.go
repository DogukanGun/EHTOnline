package oracleController

import (
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
		tokenAddress := context.Param("tokenAddress")
		// Access the first and second arguments.
		envNames := []string{"ETH_URL"}
		status, _, envMap := utils.InitializeENV(envNames, ".env")
		if !status {
			fmt.Println("Error in env")
			os.Exit(1)
		}
		client, err := ethclient.Dial(envMap["ETH_URL"])
		if err != nil {
			context.JSON(http.StatusNotFound, map[string]interface{}{"error": true, "message": err})
		}
		response := oracle.ChainlinkETHUSDOracle(tokenAddress, client)
		context.JSON(http.StatusOK, response)
	}
}

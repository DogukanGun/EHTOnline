package dex

import (
	"PoolAggregator/data/response"
	"PoolAggregator/uniswap"
	"PoolAggregator/uniswap/ChainlinkAggregator"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handler() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Param("token")
		var responses []response.UniswapPriceResponse
		for key, _ := range ChainlinkAggregator.CHAINLINK_ETH_PRICE_FEEDS {
			res := uniswap.DexOpportunity(token, key)
			if res.Price > 0.0 {
				responses = append(responses, res)
			}
		}
		context.JSON(http.StatusOK, map[string][]response.UniswapPriceResponse{
			"data": responses,
		})
	}
}

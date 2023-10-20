package dex

import (
	"PoolAggregator/uniswap"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Handler() gin.HandlerFunc {
	return func(context *gin.Context) {
		token0 := context.Param("token0")
		token1 := context.Param("token1")

		res := uniswap.DexOpportunity(token0, token1)

		context.JSON(http.StatusOK, res)
	}
}

package router

import (
	"PoolAggregator/controller/dex"
	"PoolAggregator/controller/oracle"
	"encoding/gob"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

// New registers the routes and returns the router.
func New() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "DELETE", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/dex/:token", dex.Handler())
	router.GET("/price/:token", oracleController.Handler())

	gob.Register(map[string]interface{}{})

	return router
}

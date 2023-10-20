package main

import (
	"PoolAggregator/router"
	"fmt"
	"net/http"
)

func main() {
	rtr := router.New()

	err := http.ListenAndServe("0.0.0.0:3000", rtr)

	if err != nil {
		fmt.Printf("There was an error with the http server: %v\n", err)
	}
}

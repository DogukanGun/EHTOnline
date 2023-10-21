package response

type UniswapPriceResponse struct {
	ShouldSell  bool    `json:"should_sell"`
	Price       float64 `json:"price"`
	Advantage   float64 `json:"advantage"`
	PoolAddress string  `json:"pool_address"`
	PoolAssets  string  `json:"pool_assets"`
}

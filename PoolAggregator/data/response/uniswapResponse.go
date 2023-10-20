package response

type UniswapPriceResponse struct {
	ShouldSell bool    `json:"should_sell"`
	Price      float64 `json:"price"`
}

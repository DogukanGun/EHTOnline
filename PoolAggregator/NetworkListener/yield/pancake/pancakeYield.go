package pancake

import (
	"PoolAggregator/data"
	"encoding/json"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"math/big"
	"net/http"
)

func GetPancakeYield(token0 string, token1 string, client *ethclient.Client) (res data.YieldCalculation) {
	res.ChainName = "Binance"
	resp, err := http.Post(GRAPHQL_ENDPOINT, "application/json", nil)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var data Query3Response
	if err := json.Unmarshal(body, &data); err != nil {
		return
	}

	feesUSD := data.PoolDayDatas[0].Pool.FeesUSD
	tvlUSD := data.PoolDayDatas[0].Pool.TotalValueLockedUSD
	dailyFeeAPR := (feesUSD / tvlUSD) * 100
	res.YieldValue = big.NewInt(int64(dailyFeeAPR))
	return
}

const GRAPHQL_ENDPOINT = "https://open-platform.nodereal.io/c394aca2e9cf46f8874641fec15b9061/pancakeswap-v3/graphql"

type PoolData struct {
	Token0 struct {
		Symbol           string `json:"symbol"`
		Name             string `json:"name"`
		TotalSupply      string `json:"totalSupply"`
		TotalValueLocked string `json:"totalValueLocked"`
	} `json:"token0"`
	Token1 struct {
		Symbol           string `json:"symbol"`
		Name             string `json:"name"`
		TotalSupply      string `json:"totalSupply"`
		TotalValueLocked string `json:"totalValueLocked"`
	} `json:"token1"`
	Token0Price            string  `json:"token0Price"`
	Token1Price            string  `json:"token1Price"`
	FeesUSD                float64 `json:"fees_usd"`
	TotalValueLockedToken0 string  `json:"totalValueLockedToken0"`
	TotalValueLockedToken1 string  `json:"totalValueLockedToken1"`
	Liquidity              string  `json:"liquidity"`
	ProtocolFeesUSD        string  `json:"protocolFeesUSD"`
	CollectedFeesToken0    string  `json:"collectedFeesToken0"`
	CollectedFeesToken1    string  `json:"collectedFeesToken1"`
	CollectedFeesUSD       string  `json:"collectedFeesUSD"`
	LiquidityProviderCount string  `json:"liquidityProviderCount"`
	TotalValueLockedUSD    float64 `json:"totalValueLockedUSD"`
	TotalValueLockedETH    string  `json:"totalValueLockedETH"`
}

type PoolDayData struct {
	ID              string   `json:"id"`
	Date            string   `json:"date"`
	Pool            PoolData `json:"pool"`
	Liquidity       string   `json:"liquidity"`
	Token0Price     string   `json:"token0Price"`
	Token1Price     string   `json:"token1Price"`
	TvlUSD          string   `json:"tvlUSD"`
	ProtocolFeesUSD string   `json:"protocolFeesUSD"`
	FeesUSD         string   `json:"feesUSD"`
}

type Query3Response struct {
	PoolDayDatas []PoolDayData `json:"poolDayDatas"`
}

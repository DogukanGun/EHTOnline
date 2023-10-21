package oracle

import (
	contracts "PoolAggregator/uniswap/ChainlinkAggregator"
	"PoolAggregator/uniswap/erc20"
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"io"
	"math"
	"math/big"
	"net/http"
	"strings"
)

type PriceOracleResponse struct {
	Price            float64 `json:"price"`
	TokenDecimal     uint8   `json:"token_decimal"`
	OracleIdentifier string  `json:"oracle_identifier"`
	ExtraData        string  `json:"extra_data"`
	Error            error   `json:"error"`
}

func RedStonePriceOracle(tokenAddress string, client *ethclient.Client, channel chan PriceOracleResponse) (response PriceOracleResponse) {
	// If the channel is closed the golang will panic while trying to send a value to that channel
	// but in this case the inner workings of the library is designed for this situation thus we need to handle the
	// panic
	defer panicHandler(response)

	// Initialize the result object
	response = PriceOracleResponse{
		Price:            0,
		TokenDecimal:     0,
		OracleIdentifier: "RedStone",
		ExtraData:        "",
		Error:            nil,
	}

	// Get the ERC20 token's decimal
	erc20Instance, err := erc20.NewErc20Caller(common.HexToAddress(tokenAddress), client)
	if err != nil {
		response.Error = err
		channel <- response
		return
	}
	response.TokenDecimal, response.Error = erc20Instance.Decimals(nil)

	if response.Error != nil {
		channel <- response
		return
	}

	// Get token abbreviation
	tokenAbb := REDSTONE_PRICE_FEEDS[tokenAddress]
	baseURL := "https://api.redstone.finance/prices/?symbol=" + tokenAbb + "&provider=redstone"

	resp, err := http.Get(baseURL)

	if err != nil {
		response.Error = err
		channel <- response
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		response.Error = err
		channel <- response
		return
	}

	var redStonePriceObject RedStonePricePayload
	err = json.Unmarshal(body, &redStonePriceObject)

	if err != nil {
		response.Error = err
		channel <- response
		return
	}

	if len(redStonePriceObject) == 0 {
		response.Error = errors.New("redstone returned empty response for the given token")
		channel <- response
		return
	}

	response.Price = redStonePriceObject[0].Value
	channel <- response

	return
}

/*
panicHandler closing a channel on the receiver side will cause panic. This function handles this
*/
func panicHandler(input PriceOracleResponse) (output PriceOracleResponse) {
	input = output
	if r := recover(); r != nil {
		return
	}
	return
}

/*
chainlinkETHUSDOracle returns the price of the given token in terms of USD. Utilizes Chainlink Ethereum Oracles
*/
func ChainlinkETHUSDOracle(tokenAddress string, client *ethclient.Client) (response PriceOracleResponse) {
	// If the channel is closed the golang will panic while trying to send a value to that channel
	// but in this case the inner workings of the library is designed for this situation thus we need to handle the
	// panic
	defer panicHandler(response)

	// Initialize the result object
	response = PriceOracleResponse{
		Price:            0,
		TokenDecimal:     0,
		OracleIdentifier: "Chainlink",
		ExtraData:        REDSTONE_PRICE_FEEDS[tokenAddress],
		Error:            nil,
	}

	// Convert the token address to lower case
	tokenAddress = strings.ToLower(tokenAddress)

	priceFeedAddress, isOk := contracts.CHAINLINK_ETH_PRICE_FEEDS[tokenAddress]

	if !isOk {
		response.Error = errors.New("price feed for the given token could not be found on chainlink eth oracles")
		return
	}

	// Get the ERC20 token's decimal
	erc20Instance, err := erc20.NewErc20Caller(common.HexToAddress(tokenAddress), client)

	if err != nil {
		response.Error = err
		return
	}

	response.TokenDecimal, response.Error = erc20Instance.Decimals(nil)

	if response.Error != nil {
		return
	}

	// Initialize the pricefeed instance
	priceFeedInstance, err := contracts.NewAggregatorV3InterfaceCaller(common.HexToAddress(priceFeedAddress), client)

	if err != nil {
		response.Error = err
		return
	}

	// Get the latest round data
	latestRoundData, err := priceFeedInstance.LatestRoundData(nil)

	if err != nil {
		response.Error = err
		return
	}

	// Get the decimals to divide it
	answerDecimal, err := priceFeedInstance.Decimals(nil)

	if err != nil {
		response.Error = err
		return
	}

	// Sanity check
	if answerDecimal <= 0 {
		response.Error = errors.New("chainlink price oracle returned a invalid decimal number")
		return
	}

	oracleAnswer := big.NewFloat(0).SetInt(latestRoundData.Answer)
	divisor := big.NewFloat(math.Pow(10, float64(answerDecimal)))

	// Calculate the price
	response.Price, _ = big.NewFloat(0).Quo(oracleAnswer, divisor).Float64()

	return
}

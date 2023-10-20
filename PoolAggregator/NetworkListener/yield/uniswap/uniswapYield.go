package uniswap

import (
	"PoolAggregator/data"
	"PoolAggregator/uniswap/factoryV3"
	"PoolAggregator/uniswap/oracle"
	"PoolAggregator/uniswap/poolV3"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func GetUniswapYield(token0 string, token1 string, client *ethclient.Client) data.YieldCalculation {
	v3Factory, _ := factoryV3.NewUniswapFactoryV3(common.HexToAddress("0x1F98431c8aD98523631AE4a59f267346ea31F984"),
		client)
	currentAddress, _ := v3Factory.GetPool(nil, common.HexToAddress(token0), common.HexToAddress(token1),
		big.NewInt(100))
	v3Pool, _ := poolV3.NewUniswapPoolV3Caller(currentAddress, client)
	var temp [32]byte
	position, err := v3Pool.Positions(nil, temp)
	if err != nil {
		fmt.Println(err)
		return data.YieldCalculation{}
	}
	feesEarned := calculateFeesEarned(
		position.FeeGrowthInside0LastX128,
		position.FeeGrowthInside0LastX128.Mul(position.FeeGrowthInside0LastX128, big.NewInt(10)),
		position.TokensOwed0,
		position.FeeGrowthInside1LastX128,
		position.FeeGrowthInside1LastX128.Mul(position.FeeGrowthInside1LastX128, big.NewInt(10)),
		position.TokensOwed1,
	)
	price := oracle.ChainlinkETHUSDOracle(token0, client)
	yield := getCalculationYield(
		position.Liquidity,
		position.TokensOwed0,
		big.NewInt(int64(int(price.Price))),
		position.TokensOwed1,
		feesEarned,
	)
	return data.YieldCalculation{
		YieldValue:  yield,
		PoolAddress: currentAddress.String(),
		ChainName:   "goerli",
	}
}

func calculateFeesEarned(FeeGrowthInside0LastX128 *big.Int, FeeGrowthInside0PriorX128 *big.Int, TokensOwed0 *big.Int, FeeGrowthInside1LastX128 *big.Int, FeeGrowthInside1PriorX128 *big.Int, TokensOwed1 *big.Int) *big.Int {
	// Calculate the square roots of fee growth rates
	sqrtFeeGrowth0Last := new(big.Float).SetInt(FeeGrowthInside0LastX128)
	sqrtFeeGrowth0Prior := new(big.Float).SetInt(FeeGrowthInside0PriorX128)
	sqrtFeeGrowth1Last := new(big.Float).SetInt(FeeGrowthInside1LastX128)
	sqrtFeeGrowth1Prior := new(big.Float).SetInt(FeeGrowthInside1PriorX128)

	sqrtFeeGrowth0Last.Sqrt(sqrtFeeGrowth0Last)
	sqrtFeeGrowth0Prior.Sqrt(sqrtFeeGrowth0Prior)
	sqrtFeeGrowth1Last.Sqrt(sqrtFeeGrowth1Last)
	sqrtFeeGrowth1Prior.Sqrt(sqrtFeeGrowth1Prior)

	// Calculate the fees earned
	feesEarned := new(big.Int)

	// Calculate fees earned for token0
	diffFeeGrowth0 := new(big.Float).Sub(sqrtFeeGrowth0Last, sqrtFeeGrowth0Prior)
	diffFeeGrowth0.Mul(diffFeeGrowth0, new(big.Float).SetInt(TokensOwed0))
	diffFeeGrowth0Int, _ := diffFeeGrowth0.Int(nil)
	feesEarned.Add(feesEarned, diffFeeGrowth0Int)

	// Calculate fees earned for token1
	diffFeeGrowth1 := new(big.Float).Sub(sqrtFeeGrowth1Last, sqrtFeeGrowth1Prior)
	diffFeeGrowth1.Mul(diffFeeGrowth1, new(big.Float).SetInt(TokensOwed1))
	diffFeeGrowth1Int, _ := diffFeeGrowth1.Int(nil)
	feesEarned.Add(feesEarned, diffFeeGrowth1Int)

	return feesEarned
}

func getCalculationYield(
	Liquidity *big.Int,
	TokensOwed0 *big.Int,
	Token0Price *big.Int,
	TokensOwed1 *big.Int,
	FeesEarned *big.Int,
) *big.Int {
	// Value = (Liquidity + (TokensOwed0 * currentPrice) + TokensOwed1) - FeesEarned
	temp := Liquidity.Add(Liquidity, TokensOwed1.Add(TokensOwed1, TokensOwed0.Mul(TokensOwed0, Token0Price)))
	return temp.Sub(temp, FeesEarned)
}

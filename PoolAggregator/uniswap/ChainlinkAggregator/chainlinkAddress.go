package ChainlinkAggregator

import "strings"

// Maps token address to the chainlink feed address. KEY -> ERC20 ADDRESS VALUE -> CHAINLINK ADDRESS
var CHAINLINK_ETH_PRICE_FEEDS = map[string]string{
	// WETH
	strings.ToLower("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"): strings.ToLower("0x5f4ec3df9cbd43714fe2740f5e3616155c5b8419"),
	// WBTC
	strings.ToLower("0x2260fac5e5542a773aa44fbcfedf7c193bc2c599"): strings.ToLower("0xf4030086522a5beea4988f8ca5b36dbc97bee88c"),
	// LINK
	strings.ToLower("0x514910771af9ca656af840dff83e8264ecf986ca"): strings.ToLower("0x2c1d072e956affc0d435cb7ac38ef18d24d9127c"),
	// AAVE
	strings.ToLower("0x7Fc66500c84A76Ad7e9c93437bFc5Ac33E2DDaE9"): strings.ToLower("0x547a514d5e3769680ce22b2361c10ea13619e8a9"),
	// UNI
	strings.ToLower("0x1f9840a85d5af5bf1d1762f925bdaddc4201f984"): strings.ToLower("0x553303d460ee0afb37edff9be42922d8ff63220e"),
	// MATIC
	strings.ToLower("0x7D1AfA7B718fb893dB30A3aBc0Cfc608AaCfeBB0"): strings.ToLower("0x7bac85a8a13a4bcd8abb3eb7d6b4d632c5a57676"),
	// WBNB
	strings.ToLower("0x418D75f65a02b3D53B2418FB8E1fe493759c7605"): strings.ToLower("0x14e613ac84a31f709eadbdf89c6cc390fdc9540a"),
	// SNX
	strings.ToLower("0xc011a73ee8576fb46f5e1c5751ca3b9fe0af2a6f"): strings.ToLower("0xdc3ea94cd0ac27d9a86c180091e7f78c683d3699"),
	// 1INCH
	strings.ToLower("0x111111111117dc0aa78b770fa6a738034120c302"): strings.ToLower("0xc929ad75b72593967de83e7f7cda0493458261d9"),
	// CRV
	strings.ToLower("0xD533a949740bb3306d119CC777fa900bA034cd52"): strings.ToLower("0xcd627aa160a6fa45eb793d19ef54f5062f20f33f"),
	// SOL
	strings.ToLower("0xD31a59c85aE9D8edEFeC411D448f90841571b89c"): strings.ToLower("0x4ffc43a60e009b551865a93d232e33fce9f01507"),
	// stETH
	strings.ToLower("0xae7ab96520de3a18e5e111b5eaab095312d7fe84"): strings.ToLower("0xcfe54b5cd566ab89272946f602d76ea879cab4a8"),
	// DAI
	strings.ToLower("0x6b175474e89094c44da98b954eedeac495271d0f"): strings.ToLower("0xAed0c38402a5d19df6E4c03F4E2DceD6e29c1ee9"),
	// USDT
	strings.ToLower("0xdac17f958d2ee523a2206206994597c13d831ec7"): strings.ToLower("0x3E7d1eAB13ad0104d2750B8863b489D65364e32D"),
}

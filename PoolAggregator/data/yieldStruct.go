package data

import "math/big"

type YieldCalculation struct {
	ChainName   string   `json:"chain_name"`
	YieldValue  *big.Int `json:"yield_value"`
	PoolAddress string   `json:"pool_address"`
}

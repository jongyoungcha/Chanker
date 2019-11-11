package types


import (
	"sync/atomic"
	"math/big"
)


type Transaction struct {
	data txdata
	hash atomic.Value
	size atomic.Value
	from atomic.Value
}


type txdata struct {
	AccountNonce uint64
	Price        *big.Int
	Amount       *big.Int
}









package ethereum

import (
	"math/big"
)

var one = big.NewInt(1)

// IncNonce increments nonce by one
func IncNonce(nonce *big.Int) {
	nonce.Add(nonce, one)
}

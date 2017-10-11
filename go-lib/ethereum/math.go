package ethereum

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

var one = big.NewInt(1)

// IncNonce increments nonce by one and returns updated nonce
func IncNonce(nonce *big.Int) *big.Int {
	return nonce.Add(nonce, one)
}

// IncTxoNonce increments transaction options nonce
func IncTxoNonce(txo *bind.TransactOpts, tx *types.Transaction) {
	if txo.Nonce == nil {
		txo.Nonce = big.NewInt(int64(tx.Nonce()))
	}
	IncNonce(txo.Nonce)
}

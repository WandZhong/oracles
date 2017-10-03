package ethereum

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/robert-zaremba/errstack"
	bat "github.com/robert-zaremba/go-bat"
)

var weiRatio *big.Int
var gweiRatio *big.Int
var one = big.NewInt(1)
var decimalSuffixes [10]string

func init() {
	var accuracy big.Accuracy
	weiRatio, accuracy = big.NewFloat(1e18).Int(weiRatio)
	if accuracy != big.Exact {
		logger.Fatal("Wrong wei accuracy", "accuracy", accuracy)
	}
	gweiRatio, accuracy = big.NewFloat(1e9).Int(gweiRatio)
	if accuracy != big.Exact {
		logger.Fatal("Wrong wei accuracy", "accuracy", accuracy)
	}
	s := ""
	for i := 9; i >= 0; i-- {
		decimalSuffixes[i] = s
		s += "0"
	}
}

// IncNonce increments nonce by one
func IncNonce(nonce *big.Int) {
	nonce.Add(nonce, one)
}

// IncTxoNonce increments transaction options nonce
func IncTxoNonce(txo *bind.TransactOpts, tx *types.Transaction) {
	if txo.Nonce == nil {
		txo.Nonce = big.NewInt(int64(tx.Nonce()))
	}
	IncNonce(txo.Nonce)
}

// ToWei converts integer (Ether units) to wei
func ToWei(amount int64) *big.Int {
	a := big.NewInt(amount)
	return a.Mul(a, weiRatio)
}

// WeiToInt converts wei to integers (Ether units - 1e18)
func WeiToInt(wei *big.Int) int64 {
	var i = new(big.Int)
	i.Set(wei)
	return i.Div(wei, weiRatio).Int64()
}

// AfToWei takes float number in Ascii, with max  9 digits after comman and converts it to Wei.
func AfToWei(amount string, errp errstack.Putter) *big.Int {
	amount, err := afToGigaInt(amount)
	if err != nil {
		errp.Put(err)
		return nil
	}
	v, e := bat.Atoi64(amount)
	if e != nil {
		errp.Put("Can't parse gwei amount")
		return nil
	}
	a := big.NewInt(v)
	return a.Mul(a, gweiRatio)
}

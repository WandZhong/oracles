package ethereum

import (
	"math/big"
)

var weiRatio *big.Int

func init() {
	var accuracy big.Accuracy
	weiRatio, accuracy = big.NewFloat(1e18).Int(weiRatio)
	if accuracy != big.Exact {
		logger.Fatal("Wrong wei accuracy", "accuracy", accuracy)
	}
}

// ToWei converts integer to wei
func ToWei(amount int64) *big.Int {
	a := big.NewInt(amount)
	return a.Mul(a, weiRatio)
}

// WeiToInt converts wei to integers denominated in 1e18
func WeiToInt(wei *big.Int) int64 {
	var i = new(big.Int)
	i.Set(wei)
	return i.Div(wei, weiRatio).Int64()
}

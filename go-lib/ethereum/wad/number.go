package wad

import (
	"math/big"

	"github.com/robert-zaremba/errstack"
)

var oneCoin *big.Int
var oneGwei *big.Int

func init() {
	var accuracy big.Accuracy
	oneCoin, accuracy = big.NewFloat(1e18).Int(oneCoin)
	if accuracy != big.Exact {
		logger.Fatal("Wrong wei accuracy", "accuracy", accuracy)
	}
	oneGwei, accuracy = big.NewFloat(1e9).Int(oneGwei)
	if accuracy != big.Exact {
		logger.Fatal("Wrong wei accuracy", "accuracy", accuracy)
	}
}

// ToWei converts integer (Ether units) to wei
func ToWei(amount uint64) *big.Int {
	var a = new(big.Int)
	a.SetUint64(amount)
	return a.Mul(a, oneCoin)
}

// WeiToInt converts wei to integers (Ether units - 1e18)
func WeiToInt(wei *big.Int) uint64 {
	var i = new(big.Int)
	i.Set(wei)
	return i.Div(wei, oneCoin).Uint64()
}

func parseDec9(amount string, isPositive bool, errp errstack.Putter) *big.Int {
	amount, err := afToCoinStr(amount)
	if err != nil {
		errp.Put(err)
		return nil
	}
	if isPositive && amount == "0" {
		errp.Put("must be positive")
		return nil
	}
	var wei = new(big.Int)
	_, ok := wei.SetString(amount, 10)
	if !ok {
		errp.Put("Can't parse decimal number")
		return nil
	}
	return wei
}

// AfToWei takes float number in Ascii, with max  9 digits after comman and converts it to Wei.
func AfToWei(amount string, errp errstack.Putter) *big.Int {
	return parseDec9(amount, false, errp)
}

// AfToPosWei takes float number in Ascii, with max  9 digits after comman and converts it to
// Wei. It puts an error if amount is zero.
func AfToPosWei(amount string, errp errstack.Putter) *big.Int {
	return parseDec9(amount, true, errp)
}

package wad

import (
	"math/big"

	etherutils "github.com/orinocopay/go-etherutils"
)

// WeiToString turns a number of Wei in to a string.
func WeiToString(wei *big.Int) string {
	s := etherutils.WeiToString(wei, true)
	l := len(s) - 5
	if l > 1 && s[l:] == "Ether" {
		return s[:l] + "Coin"
	} else if s == "0" {
		return "0 Wei"
	}
	return s
}

package wad

import (
	"math/big"
	"testing"
)

func TestWeiToString(t *testing.T) {
	mul := func(a *big.Int, m int64) *big.Int {
		var out = big.NewInt(m)
		return out.Mul(out, a)
	}
	var cases = []struct {
		v        *big.Int
		expected string
	}{
		{oneCoin, "1 Coin"},
		{mul(oneCoin, 10), "10 Coin"},
		{oneGwei, "1 GWei"},
		{mul(oneGwei, 10), "10 GWei"},
		{big.NewInt(0), "0 Wei"},
		{big.NewInt(1), "1 Wei"},
		{big.NewInt(123), "123 Wei"},
	}
	for _, c := range cases {
		if s := WeiToString(c.v); s != c.expected {
			t.Error("Got", s, "expected", c.expected)
		}

	}
}

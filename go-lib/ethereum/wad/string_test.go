// Copyright (c) 2017 Sweetbridge Stiftung (Sweetbridge Foundation)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

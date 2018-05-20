// Copyright (c) 2017 Sweetbridge Inc.
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

	etherutils "github.com/orinocopay/go-etherutils"
)

// WeiToString turns a number of Wei in to a string.
func WeiToString(wei *big.Int) string {
	if wei == nil {
		return "undefined"
	}
	s := etherutils.WeiToString(wei, true)
	l := len(s) - 5
	if l > 1 && s[l:] == "Ether" {
		return s[:l] + "Coin"
	} else if s == "0" {
		return "0 Wei"
	}
	return s
}

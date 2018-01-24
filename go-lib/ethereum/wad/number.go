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

	"github.com/robert-zaremba/errstack"
)

var oneCoinF = big.NewFloat(1e18)
var oneCoin *big.Int
var oneGwei *big.Int

func init() {
	var accuracy big.Accuracy
	oneCoin, accuracy = oneCoinF.Int(nil)
	if accuracy != big.Exact {
		logger.Fatal("Wrong wei accuracy", "accuracy", accuracy)
	}
	oneGwei, accuracy = big.NewFloat(1e9).Int(nil)
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

// FToWei transforms float64 coin 1e16 denominated into wei.
func FToWei(amount float64) *big.Int {
	w := big.NewFloat(amount)
	w = w.Mul(w, oneCoinF)
	i, _ := w.Int(nil)
	return i
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

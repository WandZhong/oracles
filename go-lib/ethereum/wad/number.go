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
	"strings"

	"github.com/robert-zaremba/errstack"
)

var oneCoinF = big.NewFloat(1e18)
var oneCoin *big.Int
var oneGwei *big.Int

type numberType int

const (
	anyNumber numberType = iota
	negative
	notNegative
	positive
)

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

// WeiToInt converts wei to coins unit (Ether units - 1e18)
func WeiToInt(wei *big.Int) uint64 {
	var i = new(big.Int)
	i.Set(wei)
	return i.Div(wei, oneCoin).Uint64()
}

// GweiToWei converts gwei to wei amount
func GweiToWei(n uint64) *big.Int {
	gwei := new(big.Int)
	gwei.SetUint64(n)
	return gwei.Mul(gwei, oneGwei)
}

func parseDec9(amount string, numberT numberType, errp errstack.Putter) *big.Int {
	amount, err := afToCoinStr(amount)
	if err != nil {
		errp.Put(err)
		return nil
	}
	switch numberT {
	case negative:
		if !strings.HasPrefix(amount, "-") {
			errp.Put("must be negative")
			return nil
		}
	case notNegative:
		if strings.HasPrefix(amount, "-") {
			errp.Put("must not be negative")
			return nil
		}
	case positive:
		if amount == "0" || strings.HasPrefix(amount, "-") {
			errp.Put("must be positive")
			return nil
		}
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
	return parseDec9(amount, anyNumber, errp)
}

// AfToNotNegWei takes float number in Ascii, with max  9 digits after comman and converts it to
// Wei. It puts an error if amount less  then zero.
func AfToNotNegWei(amount string, errp errstack.Putter) *big.Int {
	return parseDec9(amount, notNegative, errp)
}

// AfToPosWei takes float number in Ascii, with max  9 digits after comman and converts it to
// Wei. It puts an error if amount less or equal zero.
func AfToPosWei(amount string, errp errstack.Putter) *big.Int {
	return parseDec9(amount, positive, errp)
}

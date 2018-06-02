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
	. "github.com/scale-it/checkers"
	. "gopkg.in/check.v1"
)

var gweiZeros = "000000000"
var coinZeros = gweiZeros + gweiZeros

type numberCase struct {
	num      uint64
	expected string
}
type NumberSuite struct{}

func (suite *NumberSuite) TestDecimalSuffixes(c *C) {
	var cases = []numberCase{
		{0, coinZeros},
		{9, gweiZeros},
		{16, "00"},
		{18, ""}}
	for _, x := range cases {
		c.Check(decimalSuffixes[x.num], Equals, x.expected, Comment(x))
	}
}

func (suite *NumberSuite) TestToWei(c *C) {
	var cases = []numberCase{
		{0, "0"},
		{1, "1" + coinZeros},
		{999, "999" + coinZeros}}
	for _, x := range cases {
		w := ToWei(x.num)
		wstr := w.String()
		a := WeiToInt(w)
		c.Check(wstr, Equals, x.expected, Comment(x))
		c.Check(a, Equals, x.num, Comment("Convertion back to int64 doesn't work", x))
	}
}

func (suite *NumberSuite) TestGweiToWei(c *C) {
	gwei := GweiToWei(31)
	c.Check(gwei.String(), Equals, "31"+gweiZeros)
}

func (suite *NumberSuite) TestAfToWei(c *C) {
	var cases = []struct {
		str, expected string
		hasErr        bool
	}{
		{"0", "0", false},
		{"0000", "0", false},
		{"0.0000", "0", false},
		{"1", "1" + coinZeros, false},
		{"0.001", "1000000" + gweiZeros, false},
		{"1230.00123", "1230001230000" + gweiZeros, false},
		{"001230.0012300", "1230001230000" + gweiZeros, false},
		{"123456789001230.0012300", "123456789001230001230000" + gweiZeros, false},
		{"-1.2", "-1200000000" + gweiZeros, false},
	}
	runner := func(parser func(amount string, errp errstack.Putter) *big.Int) {
		for _, x := range cases {
			expected := new(big.Int)
			_, ok := expected.SetString(x.expected, 10)
			c.Assert(ok, IsTrue, Comment("Can't parse the test case: ", x))
			errb := errstack.NewBuilder()
			wei := parser(x.str, errb.Putter("amount"))
			if errb.NotNil() != x.hasErr {
				c.Error("Case ", x, " has error=", x.hasErr, ". ", errb.ToReqErr())
			} else if !x.hasErr && wei.Cmp(expected) != 0 {
				c.Errorf("%v should equal\n\t   %v for %q", wei, expected, x.str)
			}
		}
	}

	runner(AfToWei)

	cases[len(cases)-1].hasErr = true
	runner(AfToNotNegWei)

	// check positive values
	for i := 0; i < 3; i++ {
		cases[i].hasErr = true
	}
	runner(AfToPosWei)
}

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

	for i := 0; i < 3; i++ {
		cases[i].hasErr = true
	}
	runner(AfToPosWei)
}

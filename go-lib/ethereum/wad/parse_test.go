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
	. "github.com/scale-it/checkers"
	. "gopkg.in/check.v1"
)

type parseCase struct{ str, expected string }
type ParseSuite struct{}

func (suite *ParseSuite) checkCases(c *C, cases []parseCase, operator func(string) string) {
	for _, x := range cases {
		s := operator(x.str)
		c.Check(s, Equals, x.expected)
	}
}

func (suite *ParseSuite) TestDropLeadingZeros(c *C) {
	var cases = []parseCase{
		{"", ""},
		{"0", ""},
		{"00000000000000000000000000", ""},
		{"1", "1"},
		{"3", "3"},
		{"7519820", "7519820"},
		{"0000002", "2"},
		{"00200", "200"},
		{"200", "200"}}
	suite.checkCases(c, cases, dropLeadingZeros)
}

func (suite *ParseSuite) TestDropLastZeros(c *C) {
	var cases = []parseCase{
		{"", ""},
		{"0", ""},
		{"00000000000000000000000000", ""},
		{"1", "1"},
		{"3", "3"},
		{"7519820", "751982"},
		{"0000002", "0000002"},
		{"00200", "002"},
		{"200", "2"}}
	suite.checkCases(c, cases, dropLastZeros)
}

func (suite *ParseSuite) TestAfToCoinStr(c *C) {
	var cases = []parseCase{
		{"0", "0"},
		{"00", "0"},
		{"00.000", "0"},
		{"1", "1000000000" + gweiZeros},
		{"2", "2000000000" + gweiZeros},
		{"20", "20000000000" + gweiZeros},
		{"20.0", "20000000000" + gweiZeros},
		{"0.1", "100000000" + gweiZeros},
		{"1.2", "1200000000" + gweiZeros},
		{"1230.00123", "1230001230000" + gweiZeros},
		{"001230.0012300", "1230001230000" + gweiZeros},
		{"0.123456789", "123456789" + gweiZeros},
		{"0.123456789123456789", "123456789123456789"},
		{"22.123456789123456789", "22123456789123456789"},
	}
	for _, x := range cases {
		s, err := afToCoinStr(x.str)
		c.Assert(err, IsNil, Comment(x))
		c.Check(s, Equals, x.expected)
	}

	var errCases = []string{"0.1234567891234567891", "", " ", "1a", "1 2", ".",
		".1", "1.", ".0", "0."}
	for _, s := range errCases {
		_, err := afToCoinStr(s)
		c.Check(err, NotNil, Comment(s))
	}
}

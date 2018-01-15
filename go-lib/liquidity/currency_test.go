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

package liquidity

import (
	"encoding/json"

	. "gopkg.in/check.v1"
)

type CurrencySuite struct{}

func (suite CurrencySuite) TestJSONMarshal(c *C) {
	curr := CurrUSD
	data, err := json.Marshal(curr)
	c.Assert(err, IsNil)
	var curr2 Currency
	err = json.Unmarshal(data, &curr2)
	c.Assert(err, IsNil)
	c.Assert(curr, Equals, curr2)
}

func (suite CurrencySuite) TestJSONMarshalMap(c *C) {
	m := map[Currency]int{CurrUSD: 2}
	_, err := json.Marshal(m)
	c.Assert(err, IsNil)
}

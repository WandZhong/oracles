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

package trancheq

import (
	"math/big"

	"bitbucket.org/sweetbridge/oracles/go-lib/liquidity"
	"bitbucket.org/sweetbridge/oracles/go-lib/test/ethereumt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	. "gopkg.in/check.v1"
)

type EventSuite struct {
	eventDirectPledge abi.Event
}

func (suite *EventSuite) SetUpSuite(c *C) {
	suite.eventDirectPledge = LogSWCqueueDirectPledge()
}

func (suite EventSuite) TestLogSWCqueueDirectPledge(c *C) {
	// event LogSWCqueueDirectPledge(address who, uint128 wad, bytes3 currency);
	var data = `{
      "address":"0xfdada2f6dfdd969b5f9297f07f3622e4dfe462d6",
      "topics":["0x3a7f5663aac61ec71b13259e6a35978a2fba9256cd41bfacd1840a8b1c6bdd43"],
      "data":"0x00000000000000000000000000ce0d46d924cc8437c806721496599fc3ffa2680000000000000000000000000000000000000000000000008ac7230489e800007573640000000000000000000000000000000000000000000000000000000000",
      "blockNumber":"0x13b45",
      "transactionHash":"0xb0b1d0d020b3962f7246cb2978a02d2755a69eb163e6e544a54cc573738f428f",
      "transactionIndex":"0x1",
      "blockHash":"0xe3483334b3ee3a97bef4d959b5798881e82ecbf724f51e43661942b8d67e2d61",
      "logIndex":"0x1",
      "removed":false}`

	log := ethereumt.ParseLog(data, c)
	var expected = EventDirectPledge{
		Who:      common.HexToAddress("0x00Ce0d46d924CC8437c806721496599FC3FFA268"),
		Wad:      big.NewInt(0),
		Currency: liquidity.Currency{'u', 's', 'd'},
	}
	expected.Wad.SetString("10000000000000000000", 10)

	var o EventDirectPledge
	err := o.Unmarshal(log)
	c.Assert(err, IsNil)
	c.Check(o.Who, Equals, expected.Who)
	c.Check(o.Wad.String(), Equals, expected.Wad.String())
	c.Check(o.Currency, Equals, expected.Currency)
}

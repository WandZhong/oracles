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
	"math/big"

	"bitbucket.org/sweetbridge/oracles/go-lib/test/ethereumt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	. "gopkg.in/check.v1"
)

type EventSuite struct {
	eventTransferBRG abi.Event
}

func (suite *EventSuite) SetUpSuite(c *C) {
	suite.eventTransferBRG = LogBRGusdTransfer()
}

func (suite EventSuite) TestLogTransferBRG(c *C) {
	// event Transfer( address indexed from, address indexed to, uint256 value);
	var data = `{
      "logIndex": "0x0",
      "transactionIndex": "0x0",
      "transactionHash": "0xd79f2b7b08dab487b51b9f7557ce736b069195677e445926d58b9b174614d051",
      "blockHash": "0x0248c44667ce668023a7e17b60ea2c67f2911b33630004e677fe259bea7e20a7",
      "blockNumber": "0x21",
      "address": "0xef03118e3e60d9003b6c622a2e94c39ebb6985f2",
      "data": "0x00000000000000000000000000000000000000000000000000000000000f4240",
      "topics":
       [ "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef",
         "0x0000000000000000000000000f21f6fb13310ac0e17205840a91da93119efbec",
         "0x000000000000000000000000d435bbbaa004889f95f54e8232575d87793b42df" ],
      "type": "mined" }`

	log := ethereumt.ParseLog(data, c)
	var expected = EventTokenTransfer{
		From:  common.HexToAddress("0x0f21F6fB13310AC0E17205840a91dA93119efbec"),
		To:    common.HexToAddress("0xd435bbBAA004889F95F54E8232575d87793B42df"),
		Value: big.NewInt(1000000)}

	var o EventTokenTransfer
	err := o.Unmarshal(log)
	c.Assert(err, IsNil)
	c.Check(o, DeepEquals, expected)
}

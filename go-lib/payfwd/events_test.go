package payfwd

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	. "gopkg.in/check.v1"
)

type PaymentForwarderS struct{}

func (suite *PaymentForwarderS) TestUnmarshalLog(c *C) {
	var data = `{
		"address": "0x372f5aef2ce6aef17992a7f4e6036fd46528c29d",
		"topics": ["0xb98c7eedd7391b63c06f053a9d009a25952086d6ac2c28a62c3673b9a8aa32c5"],
		"data": "0x000000000000000000000000b57b73449b9db70aee8eae5324ff9600670a8005",
		"blockNumber": "0x380a0",
		"transactionHash": "0x5c30a832febc70ec210f65149693a32549cacffda122b278513af690b397a184",
		"transactionIndex": "0x0",
		"blockHash": "0xa7e85b05996cf6842a208362c5578bdf77d728c0c869698e0def64ad33df7286",
		"logIndex": "0x0",
		"removed": false
	} `

	var log = new(types.Log)
	err := log.UnmarshalJSON([]byte(data))
	c.Assert(err, IsNil)

	var expected = EventForwarderCreated{
		Forwarder: common.HexToAddress("0xb57b73449b9db70aee8eae5324ff9600670a8005"),
	}

	var o EventForwarderCreated
	err = o.Unmarshal(*log)
	logger.Error("error", err)
	c.Assert(err, IsNil)

	c.Check(o.Forwarder, Equals, expected.Forwarder)
}

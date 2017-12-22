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

package main

import (
	"context"
	"math/big"
	"time"

	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"bitbucket.org/sweetbridge/oracles/go-lib/liquidity"
	"bitbucket.org/sweetbridge/oracles/go-lib/trancheq"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	. "github.com/scale-it/checkers"
	. "gopkg.in/check.v1"
)

type PledgeS struct {
	txo *bind.TransactOpts
}

func (suite *PledgeS) SetUpSuite(c *C) {
	if !*flagIntegration {
		c.Skip("-integration not provided")
	}
	suite.txo = cf.Txo()
	tx, err := brgC.MintFor(suite.txo, suite.txo.From, big.NewInt(1000))
	c.Assert(err, IsNil)
	checkBalance(suite.txo.From, c)
	ethereum.IncTxoNonce(suite.txo, tx)
}

type pledgeChecker struct {
	expected  liquidity.EventTokenTransfer
	txHash    string
	ctx       context.Context
	ctxCancel context.CancelFunc
	c         *C
}

func (pc *pledgeChecker) check(p trancheq.Pledge) error {
	pc.c.Check(p.Tx, Equals, pc.txHash)
	pc.c.Check(p.WalletAddr.Hex(), Equals, pc.expected.From.Hex())
	pc.c.Check(p.CtrAddr.Hex(), Equals, pc.expected.To.Hex())
	pc.c.Check(p.Wad.String(), Equals, pc.expected.Value.String())
	pc.c.Check(p.Currency, Equals, liquidity.CurrUSD)
	pc.c.Check(p.Direct, IsFalse)

	pc.c.Check(createPledge(p), IsNil, Comment("Pledge should be inserted in DB"))
	pc.ctxCancel()
	return nil
}

func (suite *PledgeS) TestDirectPledge(c *C) {
	var expected = liquidity.EventTokenTransfer{
		From:  common.HexToAddress("0x0f21f6fb13310ac0e17205840a91da93119efbec"),
		To:    addrSWCq,
		Value: big.NewInt(24),
	}

	var pc = pledgeChecker{expected: expected, c: c}
	pc.ctx, pc.ctxCancel = context.WithTimeout(context.Background(), 50*time.Second)
	go subscribeTransfer(pc.ctx, pc.check)

	tx, err := brgC.Transfer(suite.txo, expected.To, expected.Value)
	if err != nil {
		pc.ctxCancel()
		c.Fatal("Should be able to send BRG", err)
	}
	ethereum.LogTx("BRG transferred", tx)
	pc.txHash = tx.Hash().Hex()
	checkBalance(suite.txo.From, c)

	<-pc.ctx.Done()
	c.Assert(pc.ctx.Err().Error(), Equals, "context canceled")
}

func checkBalance(who common.Address, c *C) {
	balance, err := brgC.BalanceOf(nil, who)
	c.Assert(err, IsNil)
	c.Log("Balance", balance)
}

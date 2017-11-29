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

	contracts "bitbucket.org/sweetbridge/oracles/go-contracts"
	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"bitbucket.org/sweetbridge/oracles/go-lib/log"
	"bitbucket.org/sweetbridge/oracles/go-lib/setup"
	"bitbucket.org/sweetbridge/oracles/go-lib/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-pg/pg"
	"github.com/robert-zaremba/log15/rollbar"
)

type mainFlags struct {
	setup.PgFlags
	setup.BaseOracleFlags
}

var flags = mainFlags{PgFlags: setup.NewPgFlags(),
	BaseOracleFlags: setup.NewBaseOracleFlags()}
var logger = log.Root()
var (
	brgC   *contracts.BridgeToken
	swcqC  *contracts.SWCqueue
	client *ethclient.Client
	db     *pg.DB

	addrBrg, addrSWCq common.Address
)

func init() {
	setup.FlagSimpleInit("tranche-queue", *flags.Rollbar, flags.BaseOracleFlags, flags.PgFlags)
	db = flags.MustConnect()
}

func setupContracts() (cf ethereum.ContractFactory) {
	var err error
	client, cf = flags.MustEthFactory()
	brgC, addrBrg, err = cf.GetBRG()
	utils.Assert(err, "Can't instantiate BRG contract")
	swcqC, addrSWCq, err = cf.GetSWCqueue()
	utils.Assert(err, "Can't instantiate SWCqueue contract")
	logger.Debug("Contract addresses:", "BRG", addrBrg.Hex(),
		"SWCqueue", addrSWCq.Hex())
	return
}

func main() {
	defer rollbar.WaitForRollbar(logger)

	ctx := context.Background()
	setupContracts()
	go listenDirectPledge(ctx, createPledge)
	listenTransfer(ctx, createPledge)
}

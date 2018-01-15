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

package main

import (
	"flag"

	"bitbucket.org/sweetbridge/oracles/go-contracts"
	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"bitbucket.org/sweetbridge/oracles/go-lib/log"
	"bitbucket.org/sweetbridge/oracles/go-lib/middleware"
	"bitbucket.org/sweetbridge/oracles/go-lib/setup"
	"bitbucket.org/sweetbridge/oracles/go-lib/utils"
	"github.com/blockcypher/gobcy"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/robert-zaremba/errstack"
	"github.com/robert-zaremba/log15/rollbar"
)

type mainFlags struct {
	setup.BaseOracleFlags
	port      *string
	apiKey    *string
	bcyNet    *string
	txTimeout *int
	btcPool   *string
	ethPool   *string
}

var (
	bcyAPI  gobcy.API
	ethPool common.Address
)

func (f mainFlags) Check() error {
	if len(*f.bcyNet) < 4 {
		return errstack.NewReq("`-bcy-net` must be at least 4 characters long")
	}

	bcyNet := networks[*flags.bcyNet]
	bcyAPI = gobcy.API{*flags.apiKey, bcyNet.coin, bcyNet.chain}
	if _, err := bcyAPI.GetChain(); err != nil {
		return errstack.NewReq("could not initialise BCY API")
	}

	var err error
	_, err = bcyAPI.GetAddr(*flags.btcPool, nil)
	if err != nil {
		return errstack.NewReq("`-btc-pool` must be a valid bitcoin address")
	}
	ethPool, err = ethereum.ParseAddress(*f.ethPool)
	if err != nil {
		return errstack.NewReq("`-eth-pool` must be a valid ethereum address")
	}

	return f.BaseOracleFlags.Check()
}

var (
	flags                   mainFlags
	logger                  = log.Root()
	client                  *ethclient.Client
	cf                      ethereum.ContractFactory
	forwarderFactory        *contracts.ForwarderFactory
	forwarderFactoryAddress common.Address
)

const serviceName = "payment-forwarder"

func setupFlags() {
	flags = mainFlags{
		BaseOracleFlags: setup.NewBaseOracleFlags(),
		apiKey:          flag.String("bcy-key", "", "BlockCypher API Token [required]"),
		bcyNet:          flag.String("bcy-net", "", "BlockCypher network (main or test) [required]"),
		btcPool:         flag.String("btc-pool", "", "the default pool address where all BTC will be forwarded to [required]"),
		ethPool:         flag.String("eth-pool", "", "the default pool address where all ETH will be forwarded to [required]"),
		port:            flag.String("port", "8000", "The HTTP listening port"),
		txTimeout:       flag.Int("tx-timeout", 600, "how many seconds should the daemon wait for the transaction receipt?"),
	}
	setup.FlagSimpleInit(serviceName, "", flags.Rollbar, flags)
}

func setupContracts() {
	var err error
	client, cf = flags.MustEthFactory()
	forwarderFactory, forwarderFactoryAddress, err = cf.GetForwarderFactory()
	utils.Assert(err, "Can't instantiate ForwardFactory contract")
}

func main() {
	setupFlags()
	defer rollbar.WaitForRollbar(logger)
	setupContracts()

	handler, r := middleware.StdRouter(serviceName)
	r.Post("/btc", handleBtcCreate)
	r.Post("/eth", handleEthCreate)
	setup.HTTPServer(serviceName, *flags.port, handler)
}

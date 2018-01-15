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
	"net/http"

	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"bitbucket.org/sweetbridge/oracles/go-lib/log"
	"bitbucket.org/sweetbridge/oracles/go-lib/middleware"
	"bitbucket.org/sweetbridge/oracles/go-lib/setup"
	"bitbucket.org/sweetbridge/oracles/go-lib/trancheq"
	"bitbucket.org/sweetbridge/oracles/go-lib/utils"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/robert-zaremba/log15/rollbar"
)

type mainFlags struct {
	setup.BaseOracleFlags
	port *string
}

var flags mainFlags
var logger = log.Root()
var pledger trancheq.Pledger
var client *ethclient.Client

const serviceName = "direct-pledge"

func setupFlags() {
	flags = mainFlags{BaseOracleFlags: setup.NewBaseOracleFlags(),
		port: flag.String("port", "8000", "The HTTP listening port")}

	setup.FlagSimpleInit(serviceName, "", flags.Rollbar, flags)
}

func setupContracts() {
	var err error
	var cf ethereum.ContractFactory
	client, cf = flags.MustEthFactory()
	pledger, err = trancheq.NewPledger(cf)
	utils.Assert(err, "Can't instantiate BRG contract")
}

func main() {
	setupFlags()
	defer rollbar.WaitForRollbar(logger)
	setupContracts()

	handler, r := middleware.StdRouter(serviceName)
	r.Post("/pledge", httpPostPledge)
	logger.Info("direct-pledge listening at", "port", *flags.port)
	if err := http.ListenAndServe(":"+*flags.port, handler); err != nil {
		logger.Error("Can't initiate HTTP service", err)
	}
}

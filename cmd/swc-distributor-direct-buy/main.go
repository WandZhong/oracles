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

	"bitbucket.org/sweetbridge/oracles/go-lib/directbuy"
	"bitbucket.org/sweetbridge/oracles/go-lib/log"
	"bitbucket.org/sweetbridge/oracles/go-lib/setup"
	"github.com/go-pg/pg"
	"github.com/robert-zaremba/errstack"
	"github.com/robert-zaremba/go-bat"
	"github.com/robert-zaremba/log15/rollbar"
)

var (
	db     *pg.DB
	logger = log.Root()
)

type mainFlags struct {
	setup.BaseOracleFlags
	*setup.GasPriceFlags
	PgFlags     setup.PgFlags
	dryRun      *bool
	tranche     *uint64
	gasPriceStr *uint64
}

func (f mainFlags) Check() error {
	var err error
	*f.tranche, err = bat.Atoui64(flag.Arg(0))
	if err != nil || *f.tranche == 0 {
		return errstack.NewReq("Tranche ID must be specified (not 0)")
	}
	return setup.FlagCheckMany(f.BaseOracleFlags, f.GasPriceFlags)
}

var flags = mainFlags{BaseOracleFlags: setup.NewBaseOracleFlags(),
	PgFlags:       setup.NewPgFlags(),
	GasPriceFlags: setup.NewGasPriceFlags(),
	dryRun:        flag.Bool("dry-run", false, "Make a dry run - if set, not transaction is executed"),
	tranche:       new(uint64)}

func init() {
	setup.FlagSimpleInit("swc-distributor", "tranche-id", flags.Rollbar, flags)
	db = flags.PgFlags.MustConnect()
}

func main() {
	defer rollbar.WaitForRollbar(logger)

	records, err := directbuy.CreateDirectBuyReport("TGE", *flags.tranche, db)
	if err != nil {
		logger.Error("Bad request", err)
		return
	}
	_, cf := flags.MustEthFactory()
	summaries, err := directbuy.ReportRecordsToSummary(records)
	if err != nil {
		logger.Error("Wrong direct-buy data", err)
		return
	}
	err = directbuy.DistributeSWC(*flags.dryRun, summaries, flags.GasPrice, cf, db)
	if err != nil {
		logger.Error("Can't perform SWC distribution", err)
	}
}

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
	"github.com/robert-zaremba/errstack"
	"github.com/robert-zaremba/log15/rollbar"
)

var logger = log.Root()

type mainFlags struct {
	setup.BaseOracleFlags
	*setup.GasPriceFlags
	dryRun      *bool
	expectedMd5 *string
	maxSWC      *uint64
}

func (f mainFlags) Check() error {
	if *f.maxSWC <= 0 {
		return errstack.NewReq("`-max-swc` must be a positive number")
	}
	return setup.FlagCheckMany(f.BaseOracleFlags, f.GasPriceFlags)
}

var flags = mainFlags{BaseOracleFlags: setup.NewBaseOracleFlags(),
	GasPriceFlags: setup.NewGasPriceFlags(),
	dryRun:        flag.Bool("dry-run", false, "Make a dry run - if set, not transaction is executed"),
	expectedMd5:   flag.String("md5sum", "", "If specified the application will check if the input file matches the given control sum."),
	maxSWC:        flag.Uint64("max-swc", 0, "Max SWC amount per row [required]")}

func init() {
	setup.FlagSimpleInit("swc-distributor", "source_file.csv", flags.Rollbar, flags)
}

func main() {
	defer rollbar.WaitForRollbar(logger)

	records, err := readRecords(flag.Arg(0))
	if err != nil {
		logger.Error("Bad request", err)
		return
	}
	_, cf := flags.MustEthFactory()
	if err = directbuy.DistributeSWC(*flags.dryRun, records, flags.GasPrice, cf, nil); err != nil {
		logger.Error("Can't perform SWC distribution", err)
	}
}

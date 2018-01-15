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
	"os"

	"bitbucket.org/sweetbridge/oracles/go-lib/log"
	"bitbucket.org/sweetbridge/oracles/go-lib/setup"
	"github.com/robert-zaremba/errstack"
	"github.com/robert-zaremba/log15/rollbar"
)

var logger = log.Root()

type mainFlags struct {
	setup.BaseOracleFlags
	dryRun      *bool
	expectedMd5 *string
	maxSWC      *uint64
}

func (f mainFlags) Check() error {
	if flag.NArg() < 1 {
		return errstack.NewReq("command argument is required")
	}
	if *f.maxSWC <= 0 {
		return errstack.NewReq("`-max-swc` must be a positive number")
	}
	return f.BaseOracleFlags.Check()
}

var flags mainFlags

func init() {
	flags = mainFlags{BaseOracleFlags: setup.NewBaseOracleFlags(),
		dryRun:      flag.Bool("dry-run", false, "Make a dry run - if set, not transaction is executed"),
		expectedMd5: flag.String("md5sum", "", "If specified the application will check if the input file matches the given control sum."),
		maxSWC:      flag.Uint64("max-swc", 0, "Max SWC amount per row [required]")}

	setup.FlagSimpleInit("swc-distributor", "source_file.csv", flags.Rollbar, flags)
}

func main() {
	defer rollbar.WaitForRollbar(logger)

	records, err := readRecords(flag.Arg(0))
	checkOK(err)
	distributeSWC(records)
}

func checkOK(err error) {
	if err != nil {
		logger.Error("Bad request", err)
		rollbar.WaitForRollbar(logger)
		os.Exit(2)
	}
}

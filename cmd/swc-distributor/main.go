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
}

func (f mainFlags) Check() error {
	if flag.NArg() < 1 {
		return errstack.NewReq("command argument is required")
	}
	return f.BaseOracleFlags.Check()
}

var flags mainFlags

func init() {
	flags = mainFlags{BaseOracleFlags: setup.NewBaseOracleFlags(),
		dryRun:      flag.Bool("dry-run", false, "Make a dry run - if set, not transaction is executed"),
		expectedMd5: flag.String("md5sum", "", "If specified the application will check if the input file matches the given control sum.")}

	setup.Flag("source_file.csv")
	setup.FlagValidate(flags)
	setup.MustLogger("swc-distributor", *flags.Rollbar)
}

func main() {
	defer rollbar.WaitForRollbar(logger)

	records, err := readRecords(flag.Arg(0))
	checkOK(err)
	checkOK(validate(records))
	distributeSWC(records)
}

func checkOK(err error) {
	if err != nil {
		logger.Error("Bad request", err)
		rollbar.WaitForRollbar(logger)
		os.Exit(2)
	}
}

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
		expectedMd5: flag.String("md5sum", "", "If specified the application will check if the input file matches the given control sum.")}

	setup.Flag("source_file.csv")
	setup.FlagValidate(flags)
	setup.MustLogger("swc-distributor", *flags.Rollbar)
}

func main() {
	defer rollbar.WaitForRollbar(logger)

	distributeSWC(flag.Arg(0))
}

func distributeSWC(fname string) {
	records, ok := readRecords(fname)
	checkOK(ok)
	checkOK(validate(records))
	transferSWC(records)
}

func checkOK(ok bool) {
	if !ok {
		rollbar.WaitForRollbar(logger)
		os.Exit(2)
	}
}

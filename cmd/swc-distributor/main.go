package main

import (
	"flag"
	"os"

	"bitbucket.org/sweetbridge/oracles/go-lib/log"
	"bitbucket.org/sweetbridge/oracles/go-lib/setup"
)

var logger = log.Root()

type mainFlags struct {
	setup.EthFlags
	ExpectedMd5 *string
}

func (f mainFlags) Check() bool {
	return f.EthFlags.Check() || flag.NArg() < 1
}

var flags mainFlags

func flagsSetup() {
	flags = mainFlags{EthFlags: setup.NewEthFlags(),
		ExpectedMd5: flag.String("md5sum", "", "If specified the application will check if the input file matches the given control sum.")}

	setup.Flag("source_file.csv")
	setup.FlagValidate(flags)
}

func main() {
	flagsSetup()
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
		os.Exit(2)
	}
}

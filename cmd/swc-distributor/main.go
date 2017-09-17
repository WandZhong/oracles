package main

import (
	"flag"
	"os"

	"github.com/ethereum/go-ethereum/common"

	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
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
	sp := ethereum.NewSchemaProvider(*flags.ContractsPath, *flags.Network)
	_, swcAddr := sp.MustReadGetAddress("SweetToken")
	distributeSWC(flag.Arg(0), swcAddr)
}

func distributeSWC(fname string, swcAddr common.Address) {
	logger.Debug("SWC addr: " + swcAddr.Hex())
	records, ok := readRecords(fname)
	checkOK(ok)
	checkOK(validate(records))
	transferSWC(records, swcAddr)
}

func checkOK(ok bool) {
	if !ok {
		os.Exit(2)
	}
}

package main

import (
	"flag"
	"os"

	"github.com/ethereum/go-ethereum/common"

	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"bitbucket.org/sweetbridge/oracles/go-lib/log"
	"bitbucket.org/sweetbridge/oracles/go-lib/setup"
	"bitbucket.org/sweetbridge/oracles/go-lib/utils"
)

var logger = log.Root()

var (
	pkFile        = flag.String("pk", "", "path to the private key json file [required]")
	pkPwd         = flag.String("pwd", "", "key file password [required]")
	contractsPath = flag.String("contracts", "", "path to contract schemas directory [required]")
	network       = flag.String("network", "", "blockchain network name [required]")
	ethHost       = flag.String("host", "localhost:8545", "ethereum node address. 'http' prefix added automatically.")
	expectedMd5   = flag.String("md5sum", "", "If specified the application will check if the input file matches the given control sum.")
)

func flagsSetup() {
	setup.Flag("source_file.csv")
	utils.AssertIsFile(*pkFile, "-pk")
	if *pkPwd == "" || *ethHost == "" || *contractsPath == "" || *network == "" ||
		flag.NArg() < 1 {
		setup.FlagFail()
	}
}

func main() {
	flagsSetup()
	if stat, err := os.Stat(*pkFile); err != nil || stat.IsDir() {
		logger.Fatal("-pk must be a valid file path.", "-pk", *pkFile, err)
	}
	if *pkPwd == "" || *ethHost == "" || flag.NArg() < 0 {
		logger.Error("Wrong CMD parameters")
		flag.Usage()
		return
	}

	_, swcAddr := ethereum.MustReadSchemaAndAddress(*contractsPath, "SweetToken", *network)
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

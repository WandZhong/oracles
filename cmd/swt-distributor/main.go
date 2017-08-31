package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/common"

	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"bitbucket.org/sweetbridge/oracles/go-lib/log"
)

var logger = log.Root()

var (
	pkFile        = flag.String("pk", "", "path to the private key json file [required]")
	pkPwd         = flag.String("pwd", "", "key file password [required]")
	contractsPath = flag.String("contracts", "", "path to contract schemas directory [required]")
	network       = flag.String("network", "", "blockchain network name [required]")
	ethHost       = flag.String("host", "localhost:8545", "ethereum node address. 'http' prefix added automatically.")
)

func flagsSetup() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr,
			"Usage of %s parameters source_file.csv\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
	if stat, err := os.Stat(*pkFile); err != nil || stat.IsDir() {
		logger.Fatal("-pk must be a valid file path.", "-pk", *pkFile, err)
	}
	if *pkPwd == "" || *ethHost == "" || *contractsPath == "" || *network == "" ||
		flag.NArg() < 1 {
		logger.Error("Wrong CMD parameters")
		flag.Usage()
		os.Exit(1)
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

	swtSchema, err := ethereum.ReadSchema(*contractsPath, "SweetToken")
	if err != nil {
		logger.Fatal("Can't read SWT contract schema", err)
	}
	swtAddr, err := swtSchema.Address(*network)
	if err != nil {
		logger.Fatal("Can't get SWT address", err)
	}
	distributeSWT(flag.Arg(0), swtAddr)
}

func distributeSWT(fname string, swtAddr common.Address) {
	logger.Debug("SWT addr: " + swtAddr.Hex())
	records, ok := readRecords(fname)
	checkOK(ok)
	checkOK(validate(records))
	transferSWT(records, swtAddr)
}

func checkOK(ok bool) {
	if !ok {
		os.Exit(2)
	}
}

package main

import (
	"flag"
	"net/http"
	"os"

	contracts "bitbucket.org/sweetbridge/oracles/go-contracts"
	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"bitbucket.org/sweetbridge/oracles/go-lib/log"
	"bitbucket.org/sweetbridge/oracles/go-lib/middleware"
	"bitbucket.org/sweetbridge/oracles/go-lib/setup"
	"bitbucket.org/sweetbridge/oracles/go-lib/utils"
	"github.com/ethereum/go-ethereum/ethclient"
)

var logger = log.Root()
var client *ethclient.Client
var brgC *contracts.BridgeToken
var (
	pkFile        = flag.String("pk", "", "path to the private key json file [required]")
	pkPwd         = flag.String("pwd", "", "key file password [required]")
	contractsPath = flag.String("contracts", "", "path to contract schemas directory [required]")
	network       = flag.String("network", "", "blockchain network name [required]")
	ethHost       = flag.String("host", "localhost:8545", "ethereum node address. 'http' prefix added automatically.")
	port          = flag.String("port", "8000", "The HTTP listening port")
)

func flagsSetup() {
	setup.Flag("")
	utils.AssertIsFile(*pkFile, "-pk")
	if *pkPwd == "" || *ethHost == "" || *contractsPath == "" || *network == "" {
		logger.Error("Wrong CMD parameters")
		flag.Usage()
		os.Exit(1)
	}
}

func setupContracts() {
	client = setup.EthClient(*ethHost)
	_, brgAddr := ethereum.MustReadSchemaAndAddress(*contractsPath, "BridgeToken", *network)
	_, swcQueueAddr := ethereum.MustReadSchemaAndAddress(*contractsPath, "SWCqueue", *network)
	logger.Debug("Contract addresses:", "brg", brgAddr.Hex(), "swc-queue", swcQueueAddr.Hex())
	var err error
	if brgC, err = contracts.NewBridgeToken(brgAddr, client); err != nil {
		logger.Fatal("Can't instantiate Bridgecoin contract", err)
	}
}

func main() {
	flagsSetup()
	setupContracts()

	r := middleware.StdRouter()
	r.Post("/pledge", httpPostPledge)
	logger.Info("brg-src-pledge listening at", "port", *port)
	if err := http.ListenAndServe(":"+*port, r); err != nil {
		logger.Error("Can't initiate HTTP service", err)
	}
}

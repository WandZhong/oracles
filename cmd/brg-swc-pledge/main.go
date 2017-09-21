package main

import (
	"flag"
	"net/http"

	contracts "bitbucket.org/sweetbridge/oracles/go-contracts"
	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"bitbucket.org/sweetbridge/oracles/go-lib/log"
	"bitbucket.org/sweetbridge/oracles/go-lib/middleware"
	"bitbucket.org/sweetbridge/oracles/go-lib/setup"
	"bitbucket.org/sweetbridge/oracles/go-lib/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var logger = log.Root()
var client *ethclient.Client
var brgC *contracts.BridgeToken
var txrFactory ethereum.TxrFactory

type mainFlags struct {
	setup.EthFlags
	Port *string
}

var flags mainFlags

func flagsSetup() {
	flags = mainFlags{EthFlags: setup.NewEthFlags(),
		Port: flag.String("port", "8000", "The HTTP listening port")}

	setup.Flag("")
	setup.FlagValidate(flags.EthFlags)
}

func setupContracts() {
	var err error
	var addrBrg common.Address
	client = setup.EthClient(*flags.Network)
	cf := ethereum.MustNewContractFactorySF(client, *flags.ContractsPath, *flags.Network)
	brgC, addrBrg, err = cf.GetBRG()
	utils.Assert(err, "Can't instantiate BRG contract")
	logger.Debug("Contract addresses:", "brg", addrBrg.Hex())
	txrFactory = flags.MustNewTxrFactory()
}

func main() {
	flagsSetup()
	setupContracts()

	r := middleware.StdRouter()
	r.Post("/pledge", httpPostPledge)
	logger.Info("brg-src-pledge listening at", "port", *flags.Port)
	if err := http.ListenAndServe(":"+*flags.Port, r); err != nil {
		logger.Error("Can't initiate HTTP service", err)
	}
}

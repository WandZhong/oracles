package main

import (
	contracts "bitbucket.org/sweetbridge/oracles/go-contracts"
	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"bitbucket.org/sweetbridge/oracles/go-lib/log"
	"bitbucket.org/sweetbridge/oracles/go-lib/setup"
	"bitbucket.org/sweetbridge/oracles/go-lib/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/robert-zaremba/log15/rollbar"
)

var logger = log.Root()
var client *ethclient.Client
var txrFactory ethereum.TxrFactory
var brgC *contracts.BridgeToken
var swcQC *contracts.SWCqueue
var addrBrg, addrSWCq common.Address

var flags setup.BaseOracleFlags

func init() {
	flags = setup.NewBaseOracleFlags()
	setup.Flag("")
	setup.FlagValidate(flags)
	setup.MustLogger("swc-queue", *flags.Rollbar)
}

func setupContracts() {
	var err error
	client = setup.EthClient(*flags.Network)
	cf := ethereum.MustNewContractFactorySF(client, *flags.ContractsPath, *flags.Network)
	brgC, addrBrg, err = cf.GetBRG()
	utils.Assert(err, "Can't instantiate BRG contract")
	swcQC, addrSWCq, err = cf.GetSWCqueue()
	utils.Assert(err, "Can't instantiate SWCqueue contract")
	logger.Debug("Contract addresses:", "BRG", addrBrg.Hex(),
		"SWCqueue", addrSWCq.Hex())
	txrFactory = flags.MustNewTxrFactory()
}

func main() {
	defer rollbar.WaitForRollbar(logger)
	setupContracts()

	var stopChan = make(chan struct{})
	listenPledge(stopChan)
}

package main

import (
	contracts "bitbucket.org/sweetbridge/oracles/go-contracts"
	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"bitbucket.org/sweetbridge/oracles/go-lib/log"
	"bitbucket.org/sweetbridge/oracles/go-lib/setup"
	"bitbucket.org/sweetbridge/oracles/go-lib/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-pg/pg"
	"github.com/robert-zaremba/log15/rollbar"
)

var logger = log.Root()
var flags setup.BaseOracleFlags
var (
	brgC   *contracts.BridgeToken
	swcqC  *contracts.SWCqueue
	client *ethclient.Client
	db     *pg.DB

	addrBrg, addrSWCq common.Address
)

func init() {
	flags = setup.NewBaseOracleFlags()
	setup.Flag("")
	setup.FlagValidate(flags)
	setup.MustLogger("swc-queue", *flags.Rollbar)
	db = setup.MustPsql()
}

func setupContracts() {
	var err error
	var cf ethereum.ContractFactory
	client, cf = flags.MustEthFactory()
	brgC, addrBrg, err = cf.GetBRG()
	utils.Assert(err, "Can't instantiate BRG contract")
	swcqC, addrSWCq, err = cf.GetSWCqueue()
	utils.Assert(err, "Can't instantiate SWCqueue contract")
	logger.Debug("Contract addresses:", "BRG", addrBrg.Hex(),
		"SWCqueue", addrSWCq.Hex())
}

func main() {
	defer rollbar.WaitForRollbar(logger)
	// createPledge()
	// find()

	setupContracts()
	listenPledge()
}

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
	"github.com/robert-zaremba/log15/rollbar"
)

var logger = log.Root()
var txrFactory ethereum.TxrFactory
var (
	brgC     *contracts.BridgeToken
	swcqC    *contracts.SWCqueue
	swcqAddr common.Address
)

type mainFlags struct {
	setup.BaseOracleFlags
	port *string
}

var flags mainFlags

func init() {
	flags = mainFlags{BaseOracleFlags: setup.NewBaseOracleFlags(),
		port: flag.String("port", "8000", "The HTTP listening port")}

	setup.Flag("")
	setup.FlagValidate(flags)
	setup.MustLogger("brg-swc-pledge", *flags.Rollbar)
}

func setupContracts() {
	var err error
	var addrBrg common.Address
	_, cf := flags.MustEthFactory()
	brgC, addrBrg, err = cf.GetBRG()
	utils.Assert(err, "Can't instantiate BRG contract")
	swcqC, swcqAddr, err = cf.GetSWCqueue()
	logger.Debug("Contract addresses:", "brg", addrBrg.Hex(), "swcq", swcqAddr.Hex())
	txrFactory = flags.MustNewTxrFactory()
}

func main() {
	defer rollbar.WaitForRollbar(logger)
	setupContracts()

	r := middleware.StdRouter()
	r.Post("/pledge", httpPostPledge)
	logger.Info("brg-src-pledge listening at", "port", *flags.port)
	if err := http.ListenAndServe(":"+*flags.port, r); err != nil {
		logger.Error("Can't initiate HTTP service", err)
	}
}

package main

import (
	"math/big"
	"net/http"

	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"bitbucket.org/sweetbridge/oracles/go-lib/liquidity"
	"github.com/go-ozzo/ozzo-routing"
	"github.com/robert-zaremba/errstack"
)

func httpPostPledge(c *routing.Context) (err error) {
	brgParam, currParam := c.PostForm("brg"), c.PostForm("currency")
	errb := errstack.NewBuilder()
	wei := ethereum.AfToWei(brgParam, errb.Putter("brg"))
	currency := liquidity.ToCurrency(currParam, errb.Putter("currency"))
	addr := ethereum.ToAddressErrp(c.PostForm("address"), errb.Putter("address"))
	if errb.NotNil() {
		return routing.NewHTTPError(http.StatusBadRequest,
			errb.ToReqErr().Error())
	}
	logger.Debug("BRG-SWC pledge request received", "brg-wei", wei, "currency", currParam,
		"address", addr.Hex())

	txo := txrFactory.Txo()
	tx, err := brgC.MintFor(txo, swcqAddr, big.NewInt(45))
	if err != nil {
		logger.Error("Can't mint BRG", err)
		return err
	}
	ethereum.LogTx("BRG minted for SWCq", tx)
	ethereum.IncTxoNonce(txo, tx)
	tx2, err := swcqC.DirectPledge(txo, addr, wei, currency)
	if err != nil {
		logger.Error("Can't pledge ", err)
		return err
	}
	ethereum.LogTx("SWCq direct pledge log recorded", tx2)
	ethereum.FlogTx(c.Response, "BRG minted for SWCq", tx)
	ethereum.FlogTx(c.Response, "SWCq direct pledge log recorded", tx2)
	return nil
}

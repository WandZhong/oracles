package main

import (
	"net/http"

	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum/wad"
	"bitbucket.org/sweetbridge/oracles/go-lib/liquidity"
	"github.com/go-ozzo/ozzo-routing"
	"github.com/robert-zaremba/errstack"
)

func httpPostPledge(c *routing.Context) (err error) {
	errb := errstack.NewBuilder()
	addr := ethereum.ToAddressErrp(c.Request.PostFormValue("address"), errb.Putter("address"))
	wei := wad.AfToWei(c.Request.PostFormValue("brg"), errb.Putter("brg"))
	currParam := c.Request.PostFormValue("currency")
	currency := liquidity.ParseCurrencyErrp(currParam, errb.Putter("currency"))
	if errb.NotNil() {
		return routing.NewHTTPError(http.StatusBadRequest,
			errb.ToReqErr().Error())
	}
	logger.Info("BRG-SWC pledge request received", "brg-wei", wei,
		"currency", currParam, "currencyCode", currency,
		"address", addr.Hex())

	txMint, txPledge, err := pledger.Post(addr, wei, currency)
	if err != nil {
		return err
	}
	ethereum.FlogTx(c.Response, "BRG minted for SWCq", txMint)
	ethereum.FlogTx(c.Response, "SWCq direct pledge log recorded", txPledge)
	return nil
}

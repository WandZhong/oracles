package main

import (
	"net/http"

	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"github.com/go-ozzo/ozzo-routing"
)

func httpPostPledge(c *routing.Context) (err error) {
	usdStr := c.PostForm("usd")
	addr, err := ethereum.ToAddress(c.PostForm("address"))
	if err != nil {
		return routing.NewHTTPError(http.StatusBadRequest,
			"Wrong request, expecting correct user account address. "+err.Error())
	}
	usdWei, err := ethereum.AfToWei(usdStr)
	if err != nil {
		return routing.NewHTTPError(http.StatusBadRequest,
			"Wrong request, expecting correct usd amount. "+err.Error())
	}
	logger.Debug("BRG-SWC pledge request received", "usd", usdStr, "address", addr.Hex())

	txo := ethereum.MustFileTxr(*flags.PkFile, *flags.PkPwd)

	// TODO - move this code to the lib
	tx, err := brgC.MintFor(txo, addr, usdWei)
	if err != nil {
		logger.Error("Can't mint BRG", err)
		return err
	}
	ethereum.LogTx("BRG minted", tx)

	// TODO how to transfer for the user?
	// ethereum.IncTxoNonce(txo, tx)
	// tx, err = brgC.Transfer(txo, ...)
	// ethereum.LogTx("BRG minted", tx)
	return nil
}

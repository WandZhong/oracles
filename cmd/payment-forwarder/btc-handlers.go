package main

import (
	"encoding/json"

	"github.com/blockcypher/gobcy"
	"github.com/go-ozzo/ozzo-routing"
	"github.com/robert-zaremba/errstack"
)

type network struct {
	coin  string
	chain string
}

var (
	networks = map[string]network{
		"main": {"btc", "main"},
		"test": {"bcy", "test"},
	}
)

var (
	bcyApi gobcy.API
)

func initBcyApi() {
	bcyNet := networks[*flags.bcyNet]
	bcyApi = gobcy.API{*flags.apiKey, bcyNet.coin, bcyNet.chain}
	if _, err := bcyApi.GetChain(); err != nil {
		logger.Error("could not initialise BCY API", err)
	}
	return
}

type bcyReturnData struct {
	Address string
	Id      string
}

func handleBtcCreate(ctx *routing.Context) error {
	callBack := ctx.Request.PostFormValue("callBack")

	forwardAddress := ctx.Request.PostFormValue("toAddress")
	payfwd, err := bcyApi.CreatePayFwd(gobcy.PayFwd{Destination: forwardAddress, CallbackURL: callBack})
	if err != nil {
		return errstack.WrapAsInf(err, "creating BTC forwarder")
	}
	js, err := json.Marshal(bcyReturnData{Address: payfwd.InputAddr, Id: payfwd.ID})
	if err != nil {
		return errstack.WrapAsInf(err, "unmarshall response")
	}
	return ctx.Write(string(js))
}

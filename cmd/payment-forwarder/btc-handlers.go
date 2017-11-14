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
	bcyAPI gobcy.API
)

func initBcyAPI() {
	bcyNet := networks[*flags.bcyNet]
	bcyAPI = gobcy.API{*flags.apiKey, bcyNet.coin, bcyNet.chain}
	if _, err := bcyAPI.GetChain(); err != nil {
		logger.Error("could not initialise BCY API", err)
	}
}

type bcyReturnData struct {
	Address string
	ID      string
}

func handleBtcCreate(ctx *routing.Context) error {
	callBack := ctx.Request.PostFormValue("callBack")

	forwardAddress := ctx.Request.PostFormValue("toAddress")
	payfwd, err := bcyAPI.CreatePayFwd(gobcy.PayFwd{Destination: forwardAddress, CallbackURL: callBack})
	if err != nil {
		return errstack.WrapAsInf(err, "creating BTC forwarder")
	}
	js, err := json.Marshal(bcyReturnData{payfwd.InputAddr, payfwd.ID})
	if err != nil {
		return errstack.WrapAsInf(err, "unmarshall response")
	}
	return ctx.Write(string(js))
}

package main

import (
	"context"
	"encoding/json"
	"time"

	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"bitbucket.org/sweetbridge/oracles/go-lib/payfwd"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-ozzo/ozzo-routing"
	"github.com/robert-zaremba/errstack"
)

func getForwarderAddress(txHash common.Hash) (*common.Address, errstack.E) {
	timeout := *flags.txTimeout
	for i := 0; i < (timeout / 2); i++ {
		time.Sleep(time.Second * 2)
		receipt, err := client.TransactionReceipt(context.Background(), txHash)
		if err == nil {
			var e = payfwd.EventForwarderCreated{}
			e.Unmarshal(*receipt.Logs[0])
			return &e.Forwarder, nil
		}
		if err.Error() != "not found" {
			return nil, errstack.WrapAsInfF(err, "can't get transaction receipt for [%s]", txHash.Hex())
		}
	}
	return nil, errstack.NewInfF("timed out transaction receipt for [%s]", txHash.Hex())
}

func handleEthCreate(ctx *routing.Context) error {
	toAddress, err := ethereum.ToAddress(ctx.Request.PostFormValue("toAddress"))
	if err != nil {
		return err
	}
	txo := cf.Txo()
	owner, txErr := forwarderFactory.Owner(nil)
	if txErr != nil {
		return errstack.WrapAsInf(txErr, "reading owner from ForwarderFactory")
	}

	if txo.From != owner {
		isAdmin, txErr := forwarderFactory.SenderHasRole(nil, "admin")
		if txErr != nil {
			return errstack.WrapAsInf(txErr, "checking role in ForwarderFactory")
		}
		if !isAdmin {
			return errstack.NewReqF("User [%s] is not authorised to create a forwarder ", txo.From)
		}
	}
	txFwdr, txErr := forwarderFactory.CreateForwarder(txo, toAddress)
	if txErr != nil {
		return errstack.WrapAsInf(txErr, "creating new Forwarder")
	}
	forwarder, err := getForwarderAddress(txFwdr.Hash())
	if txErr != nil {
		return err
	}

	js, txErr := json.Marshal(payfwd.EventForwarderCreated{Forwarder: *forwarder})
	if txErr != nil {
		return errstack.WrapAsDomain(txErr, "unmarshall event forwarder created")
	}

	return ctx.Write(js)
}

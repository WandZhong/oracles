package main

import (
	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func mustBeAnAddress(addrStr, comment string) common.Address {
	addr, err := ethereum.ToAddress(addrStr)
	if err != nil {
		logger.Fatal(comment, err)
	}
	return addr
}

func mustMkTx() *bind.TransactOpts {
	tx, err := ethereum.JKeyTransactor(*pkFile, *pkPwd)
	if err != nil {
		logger.Fatal("Can't create transactor", err)
	}
	return tx
}

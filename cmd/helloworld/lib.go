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

func mkTxr() *bind.TransactOpts {
	p, err := ethereum.NewJSONTxrFactory(*pkFile, *pkPwd)
	if err != nil {
		logger.Fatal("Can't create TxrFactory based on JSON file", "filename", *pkFile)
	}
	return p.Txo()
}

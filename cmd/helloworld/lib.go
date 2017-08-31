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
	return ethereum.MustFileTxr(*pkFile, *pkPwd)
}

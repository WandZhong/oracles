package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func connect() {
	if !strings.HasPrefix(*ethHost, "http") {
		*ethHost = "http://" + *ethHost
	}
	var err error
	if client, err = ethclient.Dial(*ethHost); err != nil {
		fmt.Println("Can't connect to the node " + *ethHost)
		panic(err)
	}
}

func toAddress(addr string) (a common.Address, err error) {
	if addr == "" {
		return a, errors.New("Address is empty")
	}
	if !common.IsHexAddress(addr) {
		return a, errors.New("Malformed address")
	}
	return common.HexToAddress(addr), nil
}

func mustBeAnAddress(addrStr, comment string) common.Address {
	addr, err := toAddress(addrStr)
	if err != nil {
		panic(comment + " " + err.Error())
	}
	return addr
}

func mkTransactor(pkFile string, password string) (res *bind.TransactOpts) {
	kr, err := os.Open(pkFile)
	if err != nil {
		panic(err)
	}
	if res, err = bind.NewTransactor(kr, password); err != nil {
		panic(err)
	}
	return
}

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"bitbucket.org/sweetbridge/oracles/go-contracts"
	"bitbucket.org/sweetbridge/oracles/go-lib/log"
	"bitbucket.org/sweetbridge/oracles/go-lib/setup"
)

var logger = log.Root()
var client *ethclient.Client
var rootAddr common.Address
var pkFile = flag.String("pk", "", "path to the private key json file [required]")
var pkPwd = flag.String("pwd", "", "key file password [required]")
var ethHost = flag.String("host", "localhost:8545", "ethereum node address. 'http' prefix added automatically. [required]")

func main() {
	setup.Init()
	logger.Info("Hello World!", "version", setup.GitVersion)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr,
			"Usage of %s options command [command option]:\nPARAMETERS:\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, `COMMANDS:
  check-user <user address>
      prints Root owner and checks if user is regitesred in the Root contract.
  register
      deploys new user directory for the user represented by the '-pk'.
`)
	}

	flag.Parse()
	rootAddr = mustBeAnAddress(os.Getenv("SB_ETH_ROOT"),
		"Define correct SB_ETH_ROOT env to root address.")
	if stat, err := os.Stat(*pkFile); err != nil || stat.IsDir() {
		logger.Fatal("-pk must be a valid file path.", "-pk", *pkFile, err)
	}
	if *pkPwd == "" || *ethHost == "" || flag.NArg() < 0 {
		logger.Error("Wrong CMD parameters")
		flag.Usage()
		return
	}

	client = setup.EthClient(*ethHost)

	switch flag.Arg(0) {
	case "register":
		registerUser()
	case "check-user":
		checkUser(flag.Arg(1))
	default:
		logger.Error("Wrong command")
		flag.Usage()
	}
}

func registerUser() {
	logger.Info("Registering user")

	curr := [3]byte{'U', 'S', 'D'} // [85 83 68]
	addr, tx, ud, err := contracts.DeployUserDirectory(mustMkTx(), client, rootAddr, curr)
	if err != nil {
		fmt.Println("Can't deploy UserDirectory", err)
		os.Exit(1)
	}
	fmt.Println("Address: ", addr.Hex(), tx, ud)
	fmt.Println("Gas: ", tx.Gas(), "\ngas price", tx.GasPrice())
}

func checkUser(addrStr string) {
	logger.Info("Checking user addStr")

	userAddr := mustBeAnAddress(addrStr, "<user address> must be specified correctly")
	root, err := contracts.NewRoot(rootAddr, client)
	if err != nil {
		panic("Can't get root contract" + err.Error())
	}
	opts := new(bind.CallOpts)

	addrOut, err := root.Owner(opts)
	if err != nil {
		panic("Can't check root owner: " + err.Error())
	}
	fmt.Println("Root owner is", addrOut.String())

	if addrOut, err = root.UserDirectories(opts, userAddr); err != nil {
		panic("Error:" + err.Error())
	} else {
		fmt.Println("Is user registered? ", addrOut != common.Address{})
	}
}

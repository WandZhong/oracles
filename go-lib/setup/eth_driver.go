package setup

import (
	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"bitbucket.org/sweetbridge/oracles/go-lib/utils"
	"github.com/ethereum/go-ethereum/ethclient"
)

type networkMap struct {
	id   int
	addr string
}

// netMap lists all available networks and addresses
var ethNetworkMap = map[string]networkMap{}

// MustEthClient creates new eth client and ContractFactory based on the network name.
func MustEthClient(networkName, contractsPath string) (*ethclient.Client, ethereum.ContractFactory) {
	n, ok := ethNetworkMap[networkName]
	if !ok {
		logger.Fatal("Unknown network name", "network", networkName)
	}
	logger.Debug("Creating Eth Client", "address", n.addr, "id", n.id)
	client, err := ethclient.Dial(n.addr)
	utils.Assert(err, "Can't connect to the node "+n.addr)
	sf, err := ethereum.NewContractFactorySF(client, contractsPath, n.id)
	utils.Assert(err, "wrong contractsPath")
	return client, sf
}

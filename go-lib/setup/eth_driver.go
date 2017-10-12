package setup

import (
	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"bitbucket.org/sweetbridge/oracles/go-lib/utils"
	"github.com/ethereum/go-ethereum/ethclient"
)

func init() {
	for k := range ethNetworkMap {
		ethNetworks = append(ethNetworks, k)
	}
}

type networkMap struct {
	id   int
	addr string
}

// netMap lists all available networks and addresses
var ethNetworkMap = map[string]networkMap{}

const testrpcID = 9

var ethNetworks []string

// MustEthFactory creates new eth client and ContractFactory based on the network name.
func MustEthFactory(networkName, contractsPath string, txrF ethereum.TxrFactory) (
	*ethclient.Client, ethereum.ContractFactory) {

	n, ok := ethNetworkMap[networkName]
	if !ok {
		logger.Fatal("Unknown network name", "network", networkName)
	}
	logger.Debug("Creating Eth Client", "address", n.addr, "id", n.id)
	client, err := ethclient.Dial(n.addr)
	utils.Assert(err, "Can't connect to the node "+n.addr)
	sf, err := ethereum.NewSchemaFactory(contractsPath, n.id)
	utils.Assert(err, "wrong contractsPath")
	cf := ethereum.NewContractFactory(client, sf, txrF, n.id == testrpcID)
	return client, cf
}

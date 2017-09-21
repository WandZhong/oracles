package setup

import (
	"github.com/ethereum/go-ethereum/ethclient"
)

type addrStruct struct {
	isSSL                  bool
	host, httpPort, wsPort string
}

// netMap lists all available networks and addresses
var ethNetworkMap = map[string]networkMap{}

// EthClient creates new eth client. Addr represent a network URL.
// The http prefix is optional. If it's https then it have to be included in addr.
func EthClient(networkName string) *ethclient.Client {
	addr := ethNetworkAddrs[networkName]
	logger.Debug("Creating Eth Clienet", "address", addr)
	client, err := ethclient.Dial(addr)
	if err != nil {
		logger.Fatal("Can't connect to the node " + addr)
	}
	return client
}

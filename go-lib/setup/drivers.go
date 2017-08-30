package setup

import (
	"strings"

	"github.com/ethereum/go-ethereum/ethclient"
)

// EthClient creates new eth client. Addr represent a network URL.
// The http prefix is optional. If it's https then it have to be included in addr.
func EthClient(addr string) *ethclient.Client {
	if !strings.HasPrefix(addr, "http") {
		addr = "http://" + addr
	}
	client, err := ethclient.Dial(addr)
	if err != nil {
		logger.Fatal("Can't connect to the node " + addr)
	}
	return client
}

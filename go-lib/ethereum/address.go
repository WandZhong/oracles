package ethereum

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/robert-zaremba/errstack"
)

// ToAddress converts hex string to Ethereum address
func ToAddress(addr string) (a common.Address, err error) {
	if addr == "" {
		return a, errstack.NewReq("Address is not specifed")
	}
	if !common.IsHexAddress(addr) {
		return a, errstack.NewReq("Malformed address")
	}
	return common.HexToAddress(addr), nil
}

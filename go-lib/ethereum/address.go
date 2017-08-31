package ethereum

import (
	"bytes"

	"github.com/ethereum/go-ethereum/common"
	"github.com/robert-zaremba/errstack"
)

// ZeroAddress represents Ethereum unknown or invalid address
var ZeroAddress = common.HexToAddress("00")
var zeroAddressSlice = ZeroAddress.Bytes()

// ToAddress converts hex string to Ethereum address
func ToAddress(addr string) (a common.Address, err errstack.E) {
	if addr == "" {
		return a, errstack.NewReq("Address is not specifed")
	}
	if !common.IsHexAddress(addr) {
		return a, errstack.NewReq("Invalid address")
	}
	return common.HexToAddress(addr), nil
}

// IsZeroAddr check if `a` is zero or invalid address
func IsZeroAddr(a common.Address) bool {
	return bytes.Compare(a.Bytes(), zeroAddressSlice) == 0
}

package ethereum

import (
	"bytes"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/robert-zaremba/errstack"
)

// ZeroAddress represents Ethereum unknown or invalid address
var ZeroAddress = common.HexToAddress("00")
var zeroAddressSlice = ZeroAddress.Bytes()

// ToAddress converts hex string to Ethereum address
func ToAddress(addr string) (a common.Address, err errstack.E) {
	if addr == "" {
		return a, errstack.NewReq("can't be empty")
	}
	if !strings.HasPrefix(addr, "0x") {
		return a, errstack.NewReq("must have 0x prefix")
	}
	if !common.IsHexAddress(addr) {
		return a, errstack.NewReq("Invalid address")
	}
	return common.HexToAddress(addr), nil
}

// ToAddressErrp calls ToAddress and sets the error in the putter
func ToAddressErrp(addr string, errp errstack.Putter) common.Address {
	a, err := ToAddress(addr)
	if err != nil {
		errp.Put(err)
	}
	return a
}

// IsZeroAddr check if `a` is zero or invalid address
func IsZeroAddr(a common.Address) bool {
	return bytes.Equal(a.Bytes(), zeroAddressSlice)
}

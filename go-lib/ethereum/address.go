package ethereum

import (
	"database/sql/driver"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/robert-zaremba/errstack"
	bat "github.com/robert-zaremba/go-bat"
)

// ZeroAddress represents Ethereum unknown or invalid address
var ZeroAddress = common.HexToAddress("00")

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

// HashToAddress encodes address from hash data
func HashToAddress(h common.Hash) (common.Address, errstack.E) {
	addr := common.BytesToAddress(h.Bytes())
	if IsZeroAddr(addr) {
		return addr, errstack.NewReq("must have 0x prefix")
	}
	return addr, nil
}

// IsZeroAddr check if `a` is zero or invalid address
func IsZeroAddr(a common.Address) bool {
	return AddrEqual(a, ZeroAddress)
}

// AddrEqual check if both addresses are equal
func AddrEqual(a, b common.Address) bool {
	for i := 0; i < common.AddressLength; i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// PgtAddress is a ethereum Address wrapper to provide DB interface
type PgtAddress struct {
	common.Address
}

// Scan implements sql.Sanner interface
func (a *PgtAddress) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	s, err := bat.UnsafeToString(src)
	if err != nil {
		return err
	}
	if !common.IsHexAddress(s) {
		return errstack.NewReq("Invalid address")
	}
	a.Address = common.HexToAddress(s)
	return nil
}

// Value implements sql/driver.Valuer
func (a PgtAddress) Value() (driver.Value, error) {
	return strings.ToLower(a.Hex()), nil
}

package ethereum

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

// MustParseABI parses abi string
func MustParseABI(name string, ctrABI string) abi.ABI {
	a, err := abi.JSON(strings.NewReader(ctrABI))
	if err != nil {
		logger.Fatal("Can't parse contract abi", "name", name, err)
	}
	return a
}

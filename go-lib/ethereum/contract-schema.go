package ethereum

import (
	"path"

	"github.com/ethereum/go-ethereum/common"
	"github.com/robert-zaremba/errstack"
	"github.com/robert-zaremba/go-bat"
)

// NetMap lists all available networks
var NetMap = map[string]int{
	"development": 256,
	"backstage":   200,
	"localhost":   201,
}

// Contracts is a list of available contracts
var Contracts = []string{
	"BridgeToken",
	"Root",
	"SWCqueue",
	"SweetToken",
	"SweetTokenLogic",
	"TokenLogic",
	"Treasury",
	"UserDirectory",
	"Vault",
	"VaultConfig",
	"Wallet",
}

// Schema is a type representing truffle-schema contract file
type Schema struct {
	Name          string `json:"contract_name"`
	Networks      map[int]NetSchema
	SchemaVersion string `json:"schema_version"`
	UpdatedAt     int    `json:"updated_at"`
}

// NetSchema is a type representing truffle-schema network description
type NetSchema struct {
	Address   string
	UpdatedAt int `json:"updated_at"`
}

// Address is a handy method which returns smart contract address deployed for given network
func (s Schema) Address(networkName string) (a common.Address, e errstack.E) {
	netID, ok := NetMap[networkName]
	if !ok {
		return a, errstack.NewReqF("Wrong blockchain network=%q", networkName)
	}
	n, ok := s.Networks[netID]
	if !ok {
		return a, errstack.NewReqF("Smart-Contract not deployed on network=%q", networkName)
	}
	return ToAddress(n.Address)
}

// ReadSchema reads truffle-schema file. The name should not finish with ".json"
func ReadSchema(dir, name string) (s Schema, e errstack.E) {
	if bat.StrSliceIdx(Contracts, name) < 0 {
		return s, errstack.NewReq("Unknown contract " + name)
	}
	e = bat.DecodeJSONFile(path.Join(dir, name+".json"), &s, logger)
	return
}

// ReadSchemaAndAddress reads schema using `ReadSchema` and extract contract address
// using then network identifier.
func ReadSchemaAndAddress(dir, name, network string) (s Schema, a common.Address, err errstack.E) {
	if s, err = ReadSchema(dir, name); err != nil {
		return
	}
	a, err = s.Address(network)
	return
}

// MustReadSchemaAndAddress returns contract schema and address and panics on error
func MustReadSchemaAndAddress(dir, name, network string) (Schema, common.Address) {
	s, a, err := ReadSchemaAndAddress(dir, name, network)
	if err != nil {
		logger.Fatal("Can't read schema or address", "dir", dir, "contract", name, "network", network, err)
	}
	return s, a
}

package ethereum

import (
	"path"

	"github.com/ethereum/go-ethereum/common"
	"github.com/robert-zaremba/errstack"
	"github.com/robert-zaremba/go-bat"
)

// NetMap lists all available networks
var NetMap = map[string]int{
	"develop":   256,
	"backstage": 10,
}

// Contracts is a list of available contracts
var Contracts = []string{
	"Assets",
	"BridgeToken",
	"Root",
	"SweetToken",
	"Token",
	"Treasury",
	"UserDirectory",
	"Vault",
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
func ReadSchema(dir string, name string) (s Schema, e errstack.E) {
	if bat.StrSliceIdx(Contracts, name) < 0 {
		return s, errstack.NewReq("Unknown contract " + name)
	}
	e = bat.DecodeJSONFile(path.Join(dir, name+".json"), &s, logger)
	return
}

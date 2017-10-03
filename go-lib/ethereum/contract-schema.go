package ethereum

import (
	"path"

	"github.com/ethereum/go-ethereum/common"
	"github.com/robert-zaremba/errstack"
	"github.com/robert-zaremba/go-bat"
)

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
func (s Schema) Address(networkID int) (a common.Address, e errstack.E) {
	n, ok := s.Networks[networkID]
	if !ok {
		return a, errstack.NewReqF("Smart-Contract not deployed on network=%v", networkID)
	}
	return ToAddress(n.Address)
}

// SchemaFactory is a structure which provides contract schema functions and data
type SchemaFactory struct {
	Dir     string
	Network int
}

// NewSchemaFactory creates new SchemaFactory.
func NewSchemaFactory(contractsPath string, network int) (SchemaFactory, errstack.E) {
	return SchemaFactory{contractsPath, network},
		bat.IsDir(contractsPath)
}

// Read reads truffle-schema file. The name should not finish with ".json"
func (sf SchemaFactory) Read(name string) (s Schema, e errstack.E) {
	if bat.StrSliceIdx(availableContracts, name) < 0 {
		return s, errstack.NewReq("Unknown contract " + name)
	}
	e = bat.DecodeJSONFile(path.Join(sf.Dir, name+".json"), &s, logger)
	return
}

// ReadGetAddress reads schema using `Read` and extract contract address
// using then network identifier.
func (sf SchemaFactory) ReadGetAddress(name string) (s Schema, a common.Address, err errstack.E) {
	if s, err = sf.Read(name); err != nil {
		return
	}
	a, err = s.Address(sf.Network)
	return
}

// MustReadGetAddress returns contract schema and address and panics on error
func (sf SchemaFactory) MustReadGetAddress(name string) (Schema, common.Address) {
	s, a, err := sf.ReadGetAddress(name)
	if err != nil {
		logger.Fatal("Can't read schema or address", "contract", name,
			"dir", sf.Dir, "network", sf.Network, err)
	}
	return s, a
}

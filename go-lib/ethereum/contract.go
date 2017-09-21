package ethereum

import (
	contracts "bitbucket.org/sweetbridge/oracles/go-contracts"
	"bitbucket.org/sweetbridge/oracles/go-lib/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/robert-zaremba/errstack"
)

// Contract names
const (
	CtrBRG           = "BridgeToken"
	CtrRoot          = "Root"
	CtrSWC           = "SweetToken"
	CtrSWCqueue      = "SWCqueue"
	CtrSWCLogic      = "SweetTokenLogic"
	CtrTokenLogic    = "TokenLogic"
	CtrTreasury      = "Treasury"
	CtrUserDirectory = "UserDirectory"
	CtrVault         = "Vault"
	CtrVaultConfig   = "VaultConfig"
	CtrWallet        = "Wallet"
)

// Contracts is a list of available contracts
var availableContracts = []string{
	CtrBRG, CtrRoot, CtrSWC, CtrSWCqueue, CtrSWCLogic, CtrTokenLogic, CtrTreasury, CtrTreasury,
	CtrUserDirectory, CtrVault, CtrVaultConfig, CtrWallet}

// ContractFactory delivers methods to easily construct contracts
type ContractFactory interface {
	GetBRG() (*contracts.BridgeToken, common.Address, errstack.E)
	GetSWC() (*contracts.SweetToken, common.Address, errstack.E)
	GetSWCqueue() (*contracts.SWCqueue, common.Address, errstack.E)
}

type schemaAddress struct {
	schema Schema
	addr   common.Address
}

type contractFactory struct {
	client *ethclient.Client
	sf     SchemaFactory
	addrs  map[string]common.Address
}

// NewContractFactory is a default contract provider based on truffle schema files.
func NewContractFactory(c *ethclient.Client, sf SchemaFactory) ContractFactory {
	return contractFactory{c, sf, map[string]common.Address{}}
}

// MustNewContractFactorySF is a default contract provider using default SchemaFactory.
// Panics in case of error.
func MustNewContractFactorySF(c *ethclient.Client, contractsPath, network string) ContractFactory {
	sf, err := NewSchemaFactory(contractsPath, network)
	utils.Assert(err, "wrong contractsPath")
	return contractFactory{c, sf, map[string]common.Address{}}
}

func (cf contractFactory) getSchemaAddres(contractName string) (common.Address, errstack.E) {
	addr, ok := cf.addrs[contractName]
	if !ok {
		var err errstack.E
		_, addr, err = cf.sf.ReadGetAddress(contractName)
		if err != nil {
			return addr, err
		}
		cf.addrs[contractName] = addr
	}
	return addr, nil
}

func (cf contractFactory) GetBRG() (*contracts.BridgeToken, common.Address, errstack.E) {
	addr, errE := cf.getSchemaAddres(CtrBRG)
	if errE != nil {
		return nil, addr, errE
	}
	ctr, err := contracts.NewBridgeToken(addr, cf.client)
	if err != nil {
		return nil, addr, errstack.WrapAsInf(err, "Can't construct BRG contract instance")
	}
	return ctr, addr, nil
}

func (cf contractFactory) GetSWC() (*contracts.SweetToken, common.Address, errstack.E) {
	addr, errE := cf.getSchemaAddres(CtrSWC)
	if errE != nil {
		return nil, addr, errE
	}
	ctr, err := contracts.NewSweetToken(addr, cf.client)
	if err != nil {
		return nil, addr, errstack.WrapAsInf(err, "Can't construct SWC contract instance")
	}
	return ctr, addr, nil
}

func (cf contractFactory) GetSWCqueue() (*contracts.SWCqueue, common.Address, errstack.E) {
	addr, errE := cf.getSchemaAddres(CtrSWCqueue)
	if errE != nil {
		return nil, addr, errE
	}
	ctr, err := contracts.NewSWCqueue(addr, cf.client)
	if err != nil {
		return nil, addr, errstack.WrapAsInf(err, "Can't construct SWCQUEUE contract instance")
	}
	return ctr, addr, nil
}

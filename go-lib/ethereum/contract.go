package ethereum

import (
	contracts "bitbucket.org/sweetbridge/oracles/go-contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
	CtrSWClogic      = "SweetTokenLogic"
	CtrTokenLogic    = "TokenLogic"
	CtrTreasury      = "Treasury"
	CtrUserDirectory = "UserDirectory"
	CtrVault         = "Vault"
	CtrVaultConfig   = "VaultConfig"
	CtrWallet        = "Wallet"
)

// Contracts is a list of available contracts
var availableContracts = []string{
	CtrBRG, CtrRoot, CtrSWC, CtrSWCqueue, CtrSWClogic, CtrTokenLogic, CtrTreasury, CtrTreasury,
	CtrUserDirectory, CtrVault, CtrVaultConfig, CtrWallet}

// ContractFactory delivers methods to easily construct contracts
type ContractFactory interface {
	GetBRG() (*contracts.BridgeToken, common.Address, errstack.E)
	GetSWC() (*contracts.SweetToken, common.Address, errstack.E)
	GetSWClogic() (*contracts.SweetTokenLogic, common.Address, errstack.E)
	GetSWCqueue() (*contracts.SWCqueue, common.Address, errstack.E)

	TxrFactory
}

type schemaAddress struct {
	schema Schema
	addr   common.Address
}

type contractFactory struct {
	client    *ethclient.Client
	sf        SchemaFactory
	txrF      TxrFactory
	isTestRPC bool

	addrs map[string]common.Address
}

// NewContractFactory is a default contract provider based on truffle schema files.
func NewContractFactory(c *ethclient.Client, sf SchemaFactory, txrF TxrFactory, isTestRPC bool) ContractFactory {
	return contractFactory{c, sf, txrF, isTestRPC,
		map[string]common.Address{}}
}

// Txo implements TxrFactory interface
func (cf contractFactory) Txo() *bind.TransactOpts {
	txo := cf.txrF.Txo()
	// nonce, err := cf.client.PendingNonceAt(context.TODO(), txo.From)
	// logger.Info("nonce", "val", nonce)
	// if err != nil {
	// 	logger.Error("Can't get pending nonce", err)
	// 	txo.Nonce = big.NewInt(1)
	// 	return txo
	// }
	// if !cf.isTestRPC {
	// 	nonce++
	// }
	// txo.Nonce = big.NewInt(int64(nonce))

	return txo
}

// Addr returns signer address
func (cf contractFactory) Addr() common.Address {
	return cf.txrF.Addr()
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

func (cf contractFactory) GetBRG() (c *contracts.BridgeToken, addr common.Address, err errstack.E) {
	addr, err = cf.mkContract(CtrBRG, func(addr common.Address) (err2 error) {
		c, err2 = contracts.NewBridgeToken(addr, cf.client)
		return
	})
	return
}

func (cf contractFactory) GetSWC() (c *contracts.SweetToken, addr common.Address, err errstack.E) {
	addr, err = cf.mkContract(CtrSWC, func(addr common.Address) (err2 error) {
		c, err2 = contracts.NewSweetToken(addr, cf.client)
		return
	})
	return
}

func (cf contractFactory) GetSWClogic() (c *contracts.SweetTokenLogic, addr common.Address, err errstack.E) {
	addr, err = cf.mkContract(CtrSWClogic, func(addr common.Address) (err2 error) {
		c, err2 = contracts.NewSweetTokenLogic(addr, cf.client)
		return
	})
	return
}

func (cf contractFactory) GetSWCqueue() (c *contracts.SWCqueue, addr common.Address, err errstack.E) {
	addr, err = cf.mkContract(CtrSWCqueue, func(addr common.Address) (err2 error) {
		c, err2 = contracts.NewSWCqueue(addr, cf.client)
		return
	})
	return
}

func (cf contractFactory) mkContract(ctrName string, constructor func(common.Address) error) (common.Address, errstack.E) {
	addr, errE := cf.getSchemaAddres(ctrName)
	if errE != nil {
		return addr, errE
	}
	if err := constructor(addr); err != nil {
		return addr, errstack.WrapAsInfF(err, "Can't construct %s contract instance", ctrName)
	}
	return addr, nil
}

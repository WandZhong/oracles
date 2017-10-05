// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// VaultConfigABI is the input ABI used to generate the binding from.
const VaultConfigABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxUOU\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"uouFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"days_\",\"type\":\"uint32\"}],\"name\":\"setMaxUOUdays\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"ratio\",\"type\":\"uint32\"}],\"name\":\"setMaxUOU\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address[]\"},{\"name\":\"ratio\",\"type\":\"uint32[]\"}],\"name\":\"setMaxUOUs\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxUOUdays\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"}]"

// VaultConfigBin is the compiled bytecode used for deploying new contracts.
const VaultConfigBin = `undefined`

// DeployVaultConfig deploys a new Ethereum contract, binding an instance of VaultConfig to it.
func DeployVaultConfig(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VaultConfig, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultConfigABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VaultConfigBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VaultConfig{VaultConfigCaller: VaultConfigCaller{contract: contract}, VaultConfigTransactor: VaultConfigTransactor{contract: contract}}, nil
}

// VaultConfig is an auto generated Go binding around an Ethereum contract.
type VaultConfig struct {
	VaultConfigCaller     // Read-only binding to the contract
	VaultConfigTransactor // Write-only binding to the contract
}

// VaultConfigCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultConfigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultConfigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultConfigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultConfigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultConfigSession struct {
	Contract     *VaultConfig      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultConfigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultConfigCallerSession struct {
	Contract *VaultConfigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// VaultConfigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultConfigTransactorSession struct {
	Contract     *VaultConfigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// VaultConfigRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultConfigRaw struct {
	Contract *VaultConfig // Generic contract binding to access the raw methods on
}

// VaultConfigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultConfigCallerRaw struct {
	Contract *VaultConfigCaller // Generic read-only contract binding to access the raw methods on
}

// VaultConfigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultConfigTransactorRaw struct {
	Contract *VaultConfigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVaultConfig creates a new instance of VaultConfig, bound to a specific deployed contract.
func NewVaultConfig(address common.Address, backend bind.ContractBackend) (*VaultConfig, error) {
	contract, err := bindVaultConfig(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VaultConfig{VaultConfigCaller: VaultConfigCaller{contract: contract}, VaultConfigTransactor: VaultConfigTransactor{contract: contract}}, nil
}

// NewVaultConfigCaller creates a new read-only instance of VaultConfig, bound to a specific deployed contract.
func NewVaultConfigCaller(address common.Address, caller bind.ContractCaller) (*VaultConfigCaller, error) {
	contract, err := bindVaultConfig(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &VaultConfigCaller{contract: contract}, nil
}

// NewVaultConfigTransactor creates a new write-only instance of VaultConfig, bound to a specific deployed contract.
func NewVaultConfigTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultConfigTransactor, error) {
	contract, err := bindVaultConfig(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &VaultConfigTransactor{contract: contract}, nil
}

// bindVaultConfig binds a generic wrapper to an already deployed contract.
func bindVaultConfig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultConfigABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultConfig *VaultConfigRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VaultConfig.Contract.VaultConfigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultConfig *VaultConfigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultConfig.Contract.VaultConfigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultConfig *VaultConfigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultConfig.Contract.VaultConfigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultConfig *VaultConfigCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VaultConfig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultConfig *VaultConfigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultConfig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultConfig *VaultConfigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultConfig.Contract.contract.Transact(opts, method, params...)
}

// MaxUOU is a free data retrieval call binding the contract method 0x0729bfe8.
//
// Solidity: function maxUOU( address) constant returns(uint32)
func (_VaultConfig *VaultConfigCaller) MaxUOU(opts *bind.CallOpts, arg0 common.Address) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _VaultConfig.contract.Call(opts, out, "maxUOU", arg0)
	return *ret0, err
}

// MaxUOU is a free data retrieval call binding the contract method 0x0729bfe8.
//
// Solidity: function maxUOU( address) constant returns(uint32)
func (_VaultConfig *VaultConfigSession) MaxUOU(arg0 common.Address) (uint32, error) {
	return _VaultConfig.Contract.MaxUOU(&_VaultConfig.CallOpts, arg0)
}

// MaxUOU is a free data retrieval call binding the contract method 0x0729bfe8.
//
// Solidity: function maxUOU( address) constant returns(uint32)
func (_VaultConfig *VaultConfigCallerSession) MaxUOU(arg0 common.Address) (uint32, error) {
	return _VaultConfig.Contract.MaxUOU(&_VaultConfig.CallOpts, arg0)
}

// MaxUOUdays is a free data retrieval call binding the contract method 0xc8031b81.
//
// Solidity: function maxUOUdays() constant returns(uint32)
func (_VaultConfig *VaultConfigCaller) MaxUOUdays(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _VaultConfig.contract.Call(opts, out, "maxUOUdays")
	return *ret0, err
}

// MaxUOUdays is a free data retrieval call binding the contract method 0xc8031b81.
//
// Solidity: function maxUOUdays() constant returns(uint32)
func (_VaultConfig *VaultConfigSession) MaxUOUdays() (uint32, error) {
	return _VaultConfig.Contract.MaxUOUdays(&_VaultConfig.CallOpts)
}

// MaxUOUdays is a free data retrieval call binding the contract method 0xc8031b81.
//
// Solidity: function maxUOUdays() constant returns(uint32)
func (_VaultConfig *VaultConfigCallerSession) MaxUOUdays() (uint32, error) {
	return _VaultConfig.Contract.MaxUOUdays(&_VaultConfig.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_VaultConfig *VaultConfigCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VaultConfig.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_VaultConfig *VaultConfigSession) Owner() (common.Address, error) {
	return _VaultConfig.Contract.Owner(&_VaultConfig.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_VaultConfig *VaultConfigCallerSession) Owner() (common.Address, error) {
	return _VaultConfig.Contract.Owner(&_VaultConfig.CallOpts)
}

// UouFee is a free data retrieval call binding the contract method 0x3bcdd3cf.
//
// Solidity: function uouFee() constant returns(uint256)
func (_VaultConfig *VaultConfigCaller) UouFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VaultConfig.contract.Call(opts, out, "uouFee")
	return *ret0, err
}

// UouFee is a free data retrieval call binding the contract method 0x3bcdd3cf.
//
// Solidity: function uouFee() constant returns(uint256)
func (_VaultConfig *VaultConfigSession) UouFee() (*big.Int, error) {
	return _VaultConfig.Contract.UouFee(&_VaultConfig.CallOpts)
}

// UouFee is a free data retrieval call binding the contract method 0x3bcdd3cf.
//
// Solidity: function uouFee() constant returns(uint256)
func (_VaultConfig *VaultConfigCallerSession) UouFee() (*big.Int, error) {
	return _VaultConfig.Contract.UouFee(&_VaultConfig.CallOpts)
}

// SetMaxUOU is a paid mutator transaction binding the contract method 0x98bd594e.
//
// Solidity: function setMaxUOU(token address, ratio uint32) returns()
func (_VaultConfig *VaultConfigTransactor) SetMaxUOU(opts *bind.TransactOpts, token common.Address, ratio uint32) (*types.Transaction, error) {
	return _VaultConfig.contract.Transact(opts, "setMaxUOU", token, ratio)
}

// SetMaxUOU is a paid mutator transaction binding the contract method 0x98bd594e.
//
// Solidity: function setMaxUOU(token address, ratio uint32) returns()
func (_VaultConfig *VaultConfigSession) SetMaxUOU(token common.Address, ratio uint32) (*types.Transaction, error) {
	return _VaultConfig.Contract.SetMaxUOU(&_VaultConfig.TransactOpts, token, ratio)
}

// SetMaxUOU is a paid mutator transaction binding the contract method 0x98bd594e.
//
// Solidity: function setMaxUOU(token address, ratio uint32) returns()
func (_VaultConfig *VaultConfigTransactorSession) SetMaxUOU(token common.Address, ratio uint32) (*types.Transaction, error) {
	return _VaultConfig.Contract.SetMaxUOU(&_VaultConfig.TransactOpts, token, ratio)
}

// SetMaxUOUdays is a paid mutator transaction binding the contract method 0x40b9e5ad.
//
// Solidity: function setMaxUOUdays(days_ uint32) returns()
func (_VaultConfig *VaultConfigTransactor) SetMaxUOUdays(opts *bind.TransactOpts, days_ uint32) (*types.Transaction, error) {
	return _VaultConfig.contract.Transact(opts, "setMaxUOUdays", days_)
}

// SetMaxUOUdays is a paid mutator transaction binding the contract method 0x40b9e5ad.
//
// Solidity: function setMaxUOUdays(days_ uint32) returns()
func (_VaultConfig *VaultConfigSession) SetMaxUOUdays(days_ uint32) (*types.Transaction, error) {
	return _VaultConfig.Contract.SetMaxUOUdays(&_VaultConfig.TransactOpts, days_)
}

// SetMaxUOUdays is a paid mutator transaction binding the contract method 0x40b9e5ad.
//
// Solidity: function setMaxUOUdays(days_ uint32) returns()
func (_VaultConfig *VaultConfigTransactorSession) SetMaxUOUdays(days_ uint32) (*types.Transaction, error) {
	return _VaultConfig.Contract.SetMaxUOUdays(&_VaultConfig.TransactOpts, days_)
}

// SetMaxUOUs is a paid mutator transaction binding the contract method 0xa714f0d1.
//
// Solidity: function setMaxUOUs(token address[], ratio uint32[]) returns()
func (_VaultConfig *VaultConfigTransactor) SetMaxUOUs(opts *bind.TransactOpts, token []common.Address, ratio []uint32) (*types.Transaction, error) {
	return _VaultConfig.contract.Transact(opts, "setMaxUOUs", token, ratio)
}

// SetMaxUOUs is a paid mutator transaction binding the contract method 0xa714f0d1.
//
// Solidity: function setMaxUOUs(token address[], ratio uint32[]) returns()
func (_VaultConfig *VaultConfigSession) SetMaxUOUs(token []common.Address, ratio []uint32) (*types.Transaction, error) {
	return _VaultConfig.Contract.SetMaxUOUs(&_VaultConfig.TransactOpts, token, ratio)
}

// SetMaxUOUs is a paid mutator transaction binding the contract method 0xa714f0d1.
//
// Solidity: function setMaxUOUs(token address[], ratio uint32[]) returns()
func (_VaultConfig *VaultConfigTransactorSession) SetMaxUOUs(token []common.Address, ratio []uint32) (*types.Transaction, error) {
	return _VaultConfig.Contract.SetMaxUOUs(&_VaultConfig.TransactOpts, token, ratio)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_VaultConfig *VaultConfigTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _VaultConfig.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_VaultConfig *VaultConfigSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _VaultConfig.Contract.SetOwner(&_VaultConfig.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_VaultConfig *VaultConfigTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _VaultConfig.Contract.SetOwner(&_VaultConfig.TransactOpts, owner_)
}

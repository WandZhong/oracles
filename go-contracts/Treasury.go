// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// TreasuryABI is the input ABI used to generate the binding from.
const TreasuryABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"vaults\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"root\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"v\",\"type\":\"address\"}],\"name\":\"addVault\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"r\",\"type\":\"address\"}],\"payable\":false,\"type\":\"constructor\"}]"

// TreasuryBin is the compiled bytecode used for deploying new contracts.
const TreasuryBin = `"0x6060604052341561000f57600080fd5b6040516020806101f4833981016040528080519150505b600160a060020a038116151561003b57600080fd5b60018054600160a060020a031916600160a060020a0383161790555b505b61018c806100686000396000f300606060405263ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663a622ee7c8114610053578063ebf0c7171461008e578063ec3a7823146100bd575b600080fd5b341561005e57600080fd5b610072600160a060020a03600435166100e4565b604051600160a060020a03909116815260200160405180910390f35b341561009957600080fd5b6100726100ff565b604051600160a060020a03909116815260200160405180910390f35b34156100c857600080fd5b6100e2600160a060020a036004358116906024351661010e565b005b600060208190529081526040902054600160a060020a031681565b600154600160a060020a031681565b600160a060020a038216151561012057fe5b600160a060020a038181166000908152602081905260409020805473ffffffffffffffffffffffffffffffffffffffff19169184169190911790555b50505600a165627a7a723058208f8c4c23bafe0ad3e111721a7e9a5675c5ae55a821851e9095a58fd281dfa5e50029"`

// DeployTreasury deploys a new Ethereum contract, binding an instance of Treasury to it.
func DeployTreasury(auth *bind.TransactOpts, backend bind.ContractBackend, r common.Address) (common.Address, *types.Transaction, *Treasury, error) {
	parsed, err := abi.JSON(strings.NewReader(TreasuryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TreasuryBin), backend, r)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Treasury{TreasuryCaller: TreasuryCaller{contract: contract}, TreasuryTransactor: TreasuryTransactor{contract: contract}}, nil
}

// Treasury is an auto generated Go binding around an Ethereum contract.
type Treasury struct {
	TreasuryCaller     // Read-only binding to the contract
	TreasuryTransactor // Write-only binding to the contract
}

// TreasuryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TreasuryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TreasuryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TreasuryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TreasurySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TreasurySession struct {
	Contract     *Treasury         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TreasuryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TreasuryCallerSession struct {
	Contract *TreasuryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TreasuryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TreasuryTransactorSession struct {
	Contract     *TreasuryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TreasuryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TreasuryRaw struct {
	Contract *Treasury // Generic contract binding to access the raw methods on
}

// TreasuryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TreasuryCallerRaw struct {
	Contract *TreasuryCaller // Generic read-only contract binding to access the raw methods on
}

// TreasuryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TreasuryTransactorRaw struct {
	Contract *TreasuryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTreasury creates a new instance of Treasury, bound to a specific deployed contract.
func NewTreasury(address common.Address, backend bind.ContractBackend) (*Treasury, error) {
	contract, err := bindTreasury(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Treasury{TreasuryCaller: TreasuryCaller{contract: contract}, TreasuryTransactor: TreasuryTransactor{contract: contract}}, nil
}

// NewTreasuryCaller creates a new read-only instance of Treasury, bound to a specific deployed contract.
func NewTreasuryCaller(address common.Address, caller bind.ContractCaller) (*TreasuryCaller, error) {
	contract, err := bindTreasury(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TreasuryCaller{contract: contract}, nil
}

// NewTreasuryTransactor creates a new write-only instance of Treasury, bound to a specific deployed contract.
func NewTreasuryTransactor(address common.Address, transactor bind.ContractTransactor) (*TreasuryTransactor, error) {
	contract, err := bindTreasury(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TreasuryTransactor{contract: contract}, nil
}

// bindTreasury binds a generic wrapper to an already deployed contract.
func bindTreasury(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TreasuryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Treasury *TreasuryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Treasury.Contract.TreasuryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Treasury *TreasuryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Treasury.Contract.TreasuryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Treasury *TreasuryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Treasury.Contract.TreasuryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Treasury *TreasuryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Treasury.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Treasury *TreasuryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Treasury.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Treasury *TreasuryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Treasury.Contract.contract.Transact(opts, method, params...)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() constant returns(address)
func (_Treasury *TreasuryCaller) Root(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Treasury.contract.Call(opts, out, "root")
	return *ret0, err
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() constant returns(address)
func (_Treasury *TreasurySession) Root() (common.Address, error) {
	return _Treasury.Contract.Root(&_Treasury.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() constant returns(address)
func (_Treasury *TreasuryCallerSession) Root() (common.Address, error) {
	return _Treasury.Contract.Root(&_Treasury.CallOpts)
}

// Vaults is a free data retrieval call binding the contract method 0xa622ee7c.
//
// Solidity: function vaults( address) constant returns(address)
func (_Treasury *TreasuryCaller) Vaults(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Treasury.contract.Call(opts, out, "vaults", arg0)
	return *ret0, err
}

// Vaults is a free data retrieval call binding the contract method 0xa622ee7c.
//
// Solidity: function vaults( address) constant returns(address)
func (_Treasury *TreasurySession) Vaults(arg0 common.Address) (common.Address, error) {
	return _Treasury.Contract.Vaults(&_Treasury.CallOpts, arg0)
}

// Vaults is a free data retrieval call binding the contract method 0xa622ee7c.
//
// Solidity: function vaults( address) constant returns(address)
func (_Treasury *TreasuryCallerSession) Vaults(arg0 common.Address) (common.Address, error) {
	return _Treasury.Contract.Vaults(&_Treasury.CallOpts, arg0)
}

// AddVault is a paid mutator transaction binding the contract method 0xec3a7823.
//
// Solidity: function addVault(owner address, v address) returns()
func (_Treasury *TreasuryTransactor) AddVault(opts *bind.TransactOpts, owner common.Address, v common.Address) (*types.Transaction, error) {
	return _Treasury.contract.Transact(opts, "addVault", owner, v)
}

// AddVault is a paid mutator transaction binding the contract method 0xec3a7823.
//
// Solidity: function addVault(owner address, v address) returns()
func (_Treasury *TreasurySession) AddVault(owner common.Address, v common.Address) (*types.Transaction, error) {
	return _Treasury.Contract.AddVault(&_Treasury.TransactOpts, owner, v)
}

// AddVault is a paid mutator transaction binding the contract method 0xec3a7823.
//
// Solidity: function addVault(owner address, v address) returns()
func (_Treasury *TreasuryTransactorSession) AddVault(owner common.Address, v common.Address) (*types.Transaction, error) {
	return _Treasury.Contract.AddVault(&_Treasury.TransactOpts, owner, v)
}

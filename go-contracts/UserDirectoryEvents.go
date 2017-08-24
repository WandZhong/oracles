// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package contracts

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// UserDirectoryEventsABI is the input ABI used to generate the binding from.
const UserDirectoryEventsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newWallet\",\"type\":\"address\"}],\"name\":\"LogWalletAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"removedWallet\",\"type\":\"address\"}],\"name\":\"LogWalletRemoved\",\"type\":\"event\"}]"

// UserDirectoryEventsBin is the compiled bytecode used for deploying new contracts.
const UserDirectoryEventsBin = `"0x60606040523415600e57600080fd5b5b603680601c6000396000f30060606040525b600080fd00a165627a7a72305820331d5bffd992731a61e28cc00ee0c6f4d2d2a7a9c61543348be3cf91a53b14bf0029"`

// DeployUserDirectoryEvents deploys a new Ethereum contract, binding an instance of UserDirectoryEvents to it.
func DeployUserDirectoryEvents(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UserDirectoryEvents, error) {
	parsed, err := abi.JSON(strings.NewReader(UserDirectoryEventsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UserDirectoryEventsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UserDirectoryEvents{UserDirectoryEventsCaller: UserDirectoryEventsCaller{contract: contract}, UserDirectoryEventsTransactor: UserDirectoryEventsTransactor{contract: contract}}, nil
}

// UserDirectoryEvents is an auto generated Go binding around an Ethereum contract.
type UserDirectoryEvents struct {
	UserDirectoryEventsCaller     // Read-only binding to the contract
	UserDirectoryEventsTransactor // Write-only binding to the contract
}

// UserDirectoryEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type UserDirectoryEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserDirectoryEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UserDirectoryEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserDirectoryEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UserDirectoryEventsSession struct {
	Contract     *UserDirectoryEvents // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// UserDirectoryEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UserDirectoryEventsCallerSession struct {
	Contract *UserDirectoryEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// UserDirectoryEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UserDirectoryEventsTransactorSession struct {
	Contract     *UserDirectoryEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// UserDirectoryEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type UserDirectoryEventsRaw struct {
	Contract *UserDirectoryEvents // Generic contract binding to access the raw methods on
}

// UserDirectoryEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UserDirectoryEventsCallerRaw struct {
	Contract *UserDirectoryEventsCaller // Generic read-only contract binding to access the raw methods on
}

// UserDirectoryEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UserDirectoryEventsTransactorRaw struct {
	Contract *UserDirectoryEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUserDirectoryEvents creates a new instance of UserDirectoryEvents, bound to a specific deployed contract.
func NewUserDirectoryEvents(address common.Address, backend bind.ContractBackend) (*UserDirectoryEvents, error) {
	contract, err := bindUserDirectoryEvents(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UserDirectoryEvents{UserDirectoryEventsCaller: UserDirectoryEventsCaller{contract: contract}, UserDirectoryEventsTransactor: UserDirectoryEventsTransactor{contract: contract}}, nil
}

// NewUserDirectoryEventsCaller creates a new read-only instance of UserDirectoryEvents, bound to a specific deployed contract.
func NewUserDirectoryEventsCaller(address common.Address, caller bind.ContractCaller) (*UserDirectoryEventsCaller, error) {
	contract, err := bindUserDirectoryEvents(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &UserDirectoryEventsCaller{contract: contract}, nil
}

// NewUserDirectoryEventsTransactor creates a new write-only instance of UserDirectoryEvents, bound to a specific deployed contract.
func NewUserDirectoryEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*UserDirectoryEventsTransactor, error) {
	contract, err := bindUserDirectoryEvents(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &UserDirectoryEventsTransactor{contract: contract}, nil
}

// bindUserDirectoryEvents binds a generic wrapper to an already deployed contract.
func bindUserDirectoryEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UserDirectoryEventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserDirectoryEvents *UserDirectoryEventsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UserDirectoryEvents.Contract.UserDirectoryEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserDirectoryEvents *UserDirectoryEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserDirectoryEvents.Contract.UserDirectoryEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserDirectoryEvents *UserDirectoryEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserDirectoryEvents.Contract.UserDirectoryEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserDirectoryEvents *UserDirectoryEventsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UserDirectoryEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserDirectoryEvents *UserDirectoryEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserDirectoryEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserDirectoryEvents *UserDirectoryEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserDirectoryEvents.Contract.contract.Transact(opts, method, params...)
}

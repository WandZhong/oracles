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

// AssetEventsABI is the input ABI used to generate the binding from.
const AssetEventsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"AssetAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"AssetRemoved\",\"type\":\"event\"}]"

// AssetEventsBin is the compiled bytecode used for deploying new contracts.
const AssetEventsBin = `"0x60606040523415600e57600080fd5b5b603680601c6000396000f30060606040525b600080fd00a165627a7a723058204f225ea42535e20673610ee316389777294072bc479894bf6459141a9abc12230029"`

// DeployAssetEvents deploys a new Ethereum contract, binding an instance of AssetEvents to it.
func DeployAssetEvents(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AssetEvents, error) {
	parsed, err := abi.JSON(strings.NewReader(AssetEventsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AssetEventsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AssetEvents{AssetEventsCaller: AssetEventsCaller{contract: contract}, AssetEventsTransactor: AssetEventsTransactor{contract: contract}}, nil
}

// AssetEvents is an auto generated Go binding around an Ethereum contract.
type AssetEvents struct {
	AssetEventsCaller     // Read-only binding to the contract
	AssetEventsTransactor // Write-only binding to the contract
}

// AssetEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type AssetEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AssetEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AssetEventsSession struct {
	Contract     *AssetEvents      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AssetEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AssetEventsCallerSession struct {
	Contract *AssetEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// AssetEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AssetEventsTransactorSession struct {
	Contract     *AssetEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// AssetEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type AssetEventsRaw struct {
	Contract *AssetEvents // Generic contract binding to access the raw methods on
}

// AssetEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AssetEventsCallerRaw struct {
	Contract *AssetEventsCaller // Generic read-only contract binding to access the raw methods on
}

// AssetEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AssetEventsTransactorRaw struct {
	Contract *AssetEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAssetEvents creates a new instance of AssetEvents, bound to a specific deployed contract.
func NewAssetEvents(address common.Address, backend bind.ContractBackend) (*AssetEvents, error) {
	contract, err := bindAssetEvents(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AssetEvents{AssetEventsCaller: AssetEventsCaller{contract: contract}, AssetEventsTransactor: AssetEventsTransactor{contract: contract}}, nil
}

// NewAssetEventsCaller creates a new read-only instance of AssetEvents, bound to a specific deployed contract.
func NewAssetEventsCaller(address common.Address, caller bind.ContractCaller) (*AssetEventsCaller, error) {
	contract, err := bindAssetEvents(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &AssetEventsCaller{contract: contract}, nil
}

// NewAssetEventsTransactor creates a new write-only instance of AssetEvents, bound to a specific deployed contract.
func NewAssetEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*AssetEventsTransactor, error) {
	contract, err := bindAssetEvents(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &AssetEventsTransactor{contract: contract}, nil
}

// bindAssetEvents binds a generic wrapper to an already deployed contract.
func bindAssetEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AssetEventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetEvents *AssetEventsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AssetEvents.Contract.AssetEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetEvents *AssetEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetEvents.Contract.AssetEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetEvents *AssetEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetEvents.Contract.AssetEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AssetEvents *AssetEventsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AssetEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AssetEvents *AssetEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AssetEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AssetEvents *AssetEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AssetEvents.Contract.contract.Transact(opts, method, params...)
}

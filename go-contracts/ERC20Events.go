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

// ERC20EventsABI is the input ABI used to generate the binding from.
const ERC20EventsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// ERC20EventsBin is the compiled bytecode used for deploying new contracts.
const ERC20EventsBin = `"0x60606040523415600e57600080fd5b5b603680601c6000396000f30060606040525b600080fd00a165627a7a72305820a8edb8ce6f381591f202eb598e5c650525494693931f955b6b16edb948bcd0ce0029"`

// DeployERC20Events deploys a new Ethereum contract, binding an instance of ERC20Events to it.
func DeployERC20Events(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20Events, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20EventsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20EventsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20Events{ERC20EventsCaller: ERC20EventsCaller{contract: contract}, ERC20EventsTransactor: ERC20EventsTransactor{contract: contract}}, nil
}

// ERC20Events is an auto generated Go binding around an Ethereum contract.
type ERC20Events struct {
	ERC20EventsCaller     // Read-only binding to the contract
	ERC20EventsTransactor // Write-only binding to the contract
}

// ERC20EventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20EventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20EventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20EventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20EventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20EventsSession struct {
	Contract     *ERC20Events      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20EventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20EventsCallerSession struct {
	Contract *ERC20EventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ERC20EventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20EventsTransactorSession struct {
	Contract     *ERC20EventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ERC20EventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20EventsRaw struct {
	Contract *ERC20Events // Generic contract binding to access the raw methods on
}

// ERC20EventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20EventsCallerRaw struct {
	Contract *ERC20EventsCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20EventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20EventsTransactorRaw struct {
	Contract *ERC20EventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Events creates a new instance of ERC20Events, bound to a specific deployed contract.
func NewERC20Events(address common.Address, backend bind.ContractBackend) (*ERC20Events, error) {
	contract, err := bindERC20Events(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Events{ERC20EventsCaller: ERC20EventsCaller{contract: contract}, ERC20EventsTransactor: ERC20EventsTransactor{contract: contract}}, nil
}

// NewERC20EventsCaller creates a new read-only instance of ERC20Events, bound to a specific deployed contract.
func NewERC20EventsCaller(address common.Address, caller bind.ContractCaller) (*ERC20EventsCaller, error) {
	contract, err := bindERC20Events(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20EventsCaller{contract: contract}, nil
}

// NewERC20EventsTransactor creates a new write-only instance of ERC20Events, bound to a specific deployed contract.
func NewERC20EventsTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20EventsTransactor, error) {
	contract, err := bindERC20Events(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ERC20EventsTransactor{contract: contract}, nil
}

// bindERC20Events binds a generic wrapper to an already deployed contract.
func bindERC20Events(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20EventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Events *ERC20EventsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Events.Contract.ERC20EventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Events *ERC20EventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Events.Contract.ERC20EventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Events *ERC20EventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Events.Contract.ERC20EventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Events *ERC20EventsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Events.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Events *ERC20EventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Events.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Events *ERC20EventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Events.Contract.contract.Transact(opts, method, params...)
}

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

// AuthEventsABI is the input ABI used to generate the binding from.
const AuthEventsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"UnauthorizedAccess\",\"type\":\"event\"}]"

// AuthEventsBin is the compiled bytecode used for deploying new contracts.
const AuthEventsBin = `"0x60606040523415600e57600080fd5b5b603680601c6000396000f30060606040525b600080fd00a165627a7a72305820a1eae9f9a301983d234147f581597ce280da4ae567987e6ecd228e39008698a40029"`

// DeployAuthEvents deploys a new Ethereum contract, binding an instance of AuthEvents to it.
func DeployAuthEvents(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AuthEvents, error) {
	parsed, err := abi.JSON(strings.NewReader(AuthEventsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AuthEventsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AuthEvents{AuthEventsCaller: AuthEventsCaller{contract: contract}, AuthEventsTransactor: AuthEventsTransactor{contract: contract}}, nil
}

// AuthEvents is an auto generated Go binding around an Ethereum contract.
type AuthEvents struct {
	AuthEventsCaller     // Read-only binding to the contract
	AuthEventsTransactor // Write-only binding to the contract
}

// AuthEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuthEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuthEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuthEventsSession struct {
	Contract     *AuthEvents       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuthEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuthEventsCallerSession struct {
	Contract *AuthEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// AuthEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuthEventsTransactorSession struct {
	Contract     *AuthEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AuthEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuthEventsRaw struct {
	Contract *AuthEvents // Generic contract binding to access the raw methods on
}

// AuthEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuthEventsCallerRaw struct {
	Contract *AuthEventsCaller // Generic read-only contract binding to access the raw methods on
}

// AuthEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuthEventsTransactorRaw struct {
	Contract *AuthEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuthEvents creates a new instance of AuthEvents, bound to a specific deployed contract.
func NewAuthEvents(address common.Address, backend bind.ContractBackend) (*AuthEvents, error) {
	contract, err := bindAuthEvents(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AuthEvents{AuthEventsCaller: AuthEventsCaller{contract: contract}, AuthEventsTransactor: AuthEventsTransactor{contract: contract}}, nil
}

// NewAuthEventsCaller creates a new read-only instance of AuthEvents, bound to a specific deployed contract.
func NewAuthEventsCaller(address common.Address, caller bind.ContractCaller) (*AuthEventsCaller, error) {
	contract, err := bindAuthEvents(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &AuthEventsCaller{contract: contract}, nil
}

// NewAuthEventsTransactor creates a new write-only instance of AuthEvents, bound to a specific deployed contract.
func NewAuthEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*AuthEventsTransactor, error) {
	contract, err := bindAuthEvents(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &AuthEventsTransactor{contract: contract}, nil
}

// bindAuthEvents binds a generic wrapper to an already deployed contract.
func bindAuthEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AuthEventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuthEvents *AuthEventsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AuthEvents.Contract.AuthEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuthEvents *AuthEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuthEvents.Contract.AuthEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuthEvents *AuthEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuthEvents.Contract.AuthEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AuthEvents *AuthEventsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AuthEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AuthEvents *AuthEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AuthEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AuthEvents *AuthEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AuthEvents.Contract.contract.Transact(opts, method, params...)
}

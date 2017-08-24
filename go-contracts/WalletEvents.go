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

// WalletEventsABI is the input ABI used to generate the binding from.
const WalletEventsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"LogVaultAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"LogVaultRemoved\",\"type\":\"event\"}]"

// WalletEventsBin is the compiled bytecode used for deploying new contracts.
const WalletEventsBin = `"0x60606040523415600e57600080fd5b5b603680601c6000396000f30060606040525b600080fd00a165627a7a72305820a1485155f174e6dc526574dab0260fc5a11ab371e69716e4f4435853d05ea6fc0029"`

// DeployWalletEvents deploys a new Ethereum contract, binding an instance of WalletEvents to it.
func DeployWalletEvents(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WalletEvents, error) {
	parsed, err := abi.JSON(strings.NewReader(WalletEventsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WalletEventsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WalletEvents{WalletEventsCaller: WalletEventsCaller{contract: contract}, WalletEventsTransactor: WalletEventsTransactor{contract: contract}}, nil
}

// WalletEvents is an auto generated Go binding around an Ethereum contract.
type WalletEvents struct {
	WalletEventsCaller     // Read-only binding to the contract
	WalletEventsTransactor // Write-only binding to the contract
}

// WalletEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletEventsSession struct {
	Contract     *WalletEvents     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletEventsCallerSession struct {
	Contract *WalletEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// WalletEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletEventsTransactorSession struct {
	Contract     *WalletEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// WalletEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletEventsRaw struct {
	Contract *WalletEvents // Generic contract binding to access the raw methods on
}

// WalletEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletEventsCallerRaw struct {
	Contract *WalletEventsCaller // Generic read-only contract binding to access the raw methods on
}

// WalletEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletEventsTransactorRaw struct {
	Contract *WalletEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWalletEvents creates a new instance of WalletEvents, bound to a specific deployed contract.
func NewWalletEvents(address common.Address, backend bind.ContractBackend) (*WalletEvents, error) {
	contract, err := bindWalletEvents(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WalletEvents{WalletEventsCaller: WalletEventsCaller{contract: contract}, WalletEventsTransactor: WalletEventsTransactor{contract: contract}}, nil
}

// NewWalletEventsCaller creates a new read-only instance of WalletEvents, bound to a specific deployed contract.
func NewWalletEventsCaller(address common.Address, caller bind.ContractCaller) (*WalletEventsCaller, error) {
	contract, err := bindWalletEvents(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &WalletEventsCaller{contract: contract}, nil
}

// NewWalletEventsTransactor creates a new write-only instance of WalletEvents, bound to a specific deployed contract.
func NewWalletEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletEventsTransactor, error) {
	contract, err := bindWalletEvents(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &WalletEventsTransactor{contract: contract}, nil
}

// bindWalletEvents binds a generic wrapper to an already deployed contract.
func bindWalletEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WalletEventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletEvents *WalletEventsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WalletEvents.Contract.WalletEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletEvents *WalletEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletEvents.Contract.WalletEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletEvents *WalletEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletEvents.Contract.WalletEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletEvents *WalletEventsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WalletEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletEvents *WalletEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletEvents *WalletEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletEvents.Contract.contract.Transact(opts, method, params...)
}

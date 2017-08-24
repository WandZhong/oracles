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

// VaultEventsABI is the input ABI used to generate the binding from.
const VaultEventsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"brgAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"uouIndex\",\"type\":\"uint256\"}],\"name\":\"UouRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"brgAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"uouIndex\",\"type\":\"uint256\"}],\"name\":\"UouRequestApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"brgAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"uouIndex\",\"type\":\"uint256\"}],\"name\":\"UouRequestDeclined\",\"type\":\"event\"}]"

// VaultEventsBin is the compiled bytecode used for deploying new contracts.
const VaultEventsBin = `"0x60606040523415600e57600080fd5b5b603680601c6000396000f30060606040525b600080fd00a165627a7a72305820501d9c76a64385ba6c2303b2d460019f0ba306c16778c650ec78640e2d02e60b0029"`

// DeployVaultEvents deploys a new Ethereum contract, binding an instance of VaultEvents to it.
func DeployVaultEvents(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VaultEvents, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultEventsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VaultEventsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VaultEvents{VaultEventsCaller: VaultEventsCaller{contract: contract}, VaultEventsTransactor: VaultEventsTransactor{contract: contract}}, nil
}

// VaultEvents is an auto generated Go binding around an Ethereum contract.
type VaultEvents struct {
	VaultEventsCaller     // Read-only binding to the contract
	VaultEventsTransactor // Write-only binding to the contract
}

// VaultEventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultEventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultEventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultEventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultEventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultEventsSession struct {
	Contract     *VaultEvents      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultEventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultEventsCallerSession struct {
	Contract *VaultEventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// VaultEventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultEventsTransactorSession struct {
	Contract     *VaultEventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// VaultEventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultEventsRaw struct {
	Contract *VaultEvents // Generic contract binding to access the raw methods on
}

// VaultEventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultEventsCallerRaw struct {
	Contract *VaultEventsCaller // Generic read-only contract binding to access the raw methods on
}

// VaultEventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultEventsTransactorRaw struct {
	Contract *VaultEventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVaultEvents creates a new instance of VaultEvents, bound to a specific deployed contract.
func NewVaultEvents(address common.Address, backend bind.ContractBackend) (*VaultEvents, error) {
	contract, err := bindVaultEvents(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VaultEvents{VaultEventsCaller: VaultEventsCaller{contract: contract}, VaultEventsTransactor: VaultEventsTransactor{contract: contract}}, nil
}

// NewVaultEventsCaller creates a new read-only instance of VaultEvents, bound to a specific deployed contract.
func NewVaultEventsCaller(address common.Address, caller bind.ContractCaller) (*VaultEventsCaller, error) {
	contract, err := bindVaultEvents(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &VaultEventsCaller{contract: contract}, nil
}

// NewVaultEventsTransactor creates a new write-only instance of VaultEvents, bound to a specific deployed contract.
func NewVaultEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultEventsTransactor, error) {
	contract, err := bindVaultEvents(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &VaultEventsTransactor{contract: contract}, nil
}

// bindVaultEvents binds a generic wrapper to an already deployed contract.
func bindVaultEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultEventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultEvents *VaultEventsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VaultEvents.Contract.VaultEventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultEvents *VaultEventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultEvents.Contract.VaultEventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultEvents *VaultEventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultEvents.Contract.VaultEventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultEvents *VaultEventsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VaultEvents.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultEvents *VaultEventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultEvents.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultEvents *VaultEventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultEvents.Contract.contract.Transact(opts, method, params...)
}

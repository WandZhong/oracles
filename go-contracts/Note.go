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

// NoteABI is the input ABI used to generate the binding from.
const NoteABI = "[{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"foo\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"bar\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fax\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"}]"

// NoteBin is the compiled bytecode used for deploying new contracts.
const NoteBin = `"0x60606040523415600e57600080fd5b5b603680601c6000396000f30060606040525b600080fd00a165627a7a72305820c1983ee4f37a133e6466e85e9c53c9bc8d2268c00e74dec8ca8d74f24a1454570029"`

// DeployNote deploys a new Ethereum contract, binding an instance of Note to it.
func DeployNote(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Note, error) {
	parsed, err := abi.JSON(strings.NewReader(NoteABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NoteBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Note{NoteCaller: NoteCaller{contract: contract}, NoteTransactor: NoteTransactor{contract: contract}}, nil
}

// Note is an auto generated Go binding around an Ethereum contract.
type Note struct {
	NoteCaller     // Read-only binding to the contract
	NoteTransactor // Write-only binding to the contract
}

// NoteCaller is an auto generated read-only Go binding around an Ethereum contract.
type NoteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NoteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NoteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NoteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NoteSession struct {
	Contract     *Note             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NoteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NoteCallerSession struct {
	Contract *NoteCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NoteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NoteTransactorSession struct {
	Contract     *NoteTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NoteRaw is an auto generated low-level Go binding around an Ethereum contract.
type NoteRaw struct {
	Contract *Note // Generic contract binding to access the raw methods on
}

// NoteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NoteCallerRaw struct {
	Contract *NoteCaller // Generic read-only contract binding to access the raw methods on
}

// NoteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NoteTransactorRaw struct {
	Contract *NoteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNote creates a new instance of Note, bound to a specific deployed contract.
func NewNote(address common.Address, backend bind.ContractBackend) (*Note, error) {
	contract, err := bindNote(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Note{NoteCaller: NoteCaller{contract: contract}, NoteTransactor: NoteTransactor{contract: contract}}, nil
}

// NewNoteCaller creates a new read-only instance of Note, bound to a specific deployed contract.
func NewNoteCaller(address common.Address, caller bind.ContractCaller) (*NoteCaller, error) {
	contract, err := bindNote(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &NoteCaller{contract: contract}, nil
}

// NewNoteTransactor creates a new write-only instance of Note, bound to a specific deployed contract.
func NewNoteTransactor(address common.Address, transactor bind.ContractTransactor) (*NoteTransactor, error) {
	contract, err := bindNote(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &NoteTransactor{contract: contract}, nil
}

// bindNote binds a generic wrapper to an already deployed contract.
func bindNote(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NoteABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Note *NoteRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Note.Contract.NoteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Note *NoteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Note.Contract.NoteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Note *NoteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Note.Contract.NoteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Note *NoteCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Note.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Note *NoteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Note.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Note *NoteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Note.Contract.contract.Transact(opts, method, params...)
}

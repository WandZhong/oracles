// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package contracts

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// MigrationsABI is the input ABI used to generate the binding from.
const MigrationsABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"upgrade\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"last_completed_migration\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"completed\",\"type\":\"uint256\"}],\"name\":\"setCompleted\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"type\":\"constructor\"}]"

// MigrationsBin is the compiled bytecode used for deploying new contracts.
const MigrationsBin = `"0x6060604052341561000f57600080fd5b5b60008054600160a060020a03191633600160a060020a03161790555b5b6101e58061003c6000396000f300606060405263ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630900f010811461005e578063445df0ac1461007f5780638da5cb5b146100a4578063fdacd576146100d3575b600080fd5b341561006957600080fd5b61007d600160a060020a03600435166100eb565b005b341561008a57600080fd5b610092610182565b60405190815260200160405180910390f35b34156100af57600080fd5b6100b7610188565b604051600160a060020a03909116815260200160405180910390f35b34156100de57600080fd5b61007d600435610197565b005b6000805433600160a060020a039081169116141561017c5781905080600160a060020a031663fdacd5766001546040517c010000000000000000000000000000000000000000000000000000000063ffffffff84160281526004810191909152602401600060405180830381600087803b151561016757600080fd5b6102c65a03f1151561017857600080fd5b5050505b5b5b5050565b60015481565b600054600160a060020a031681565b60005433600160a060020a03908116911614156101b45760018190555b5b5b505600a165627a7a723058208e2e4f30a11e1da8d8d32bcc21afbdfdac4796c4989e6254855ea7235d52822b0029"`

// DeployMigrations deploys a new Ethereum contract, binding an instance of Migrations to it.
func DeployMigrations(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Migrations, error) {
	parsed, err := abi.JSON(strings.NewReader(MigrationsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MigrationsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Migrations{MigrationsCaller: MigrationsCaller{contract: contract}, MigrationsTransactor: MigrationsTransactor{contract: contract}}, nil
}

// Migrations is an auto generated Go binding around an Ethereum contract.
type Migrations struct {
	MigrationsCaller     // Read-only binding to the contract
	MigrationsTransactor // Write-only binding to the contract
}

// MigrationsCaller is an auto generated read-only Go binding around an Ethereum contract.
type MigrationsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MigrationsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MigrationsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MigrationsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MigrationsSession struct {
	Contract     *Migrations       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MigrationsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MigrationsCallerSession struct {
	Contract *MigrationsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// MigrationsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MigrationsTransactorSession struct {
	Contract     *MigrationsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MigrationsRaw is an auto generated low-level Go binding around an Ethereum contract.
type MigrationsRaw struct {
	Contract *Migrations // Generic contract binding to access the raw methods on
}

// MigrationsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MigrationsCallerRaw struct {
	Contract *MigrationsCaller // Generic read-only contract binding to access the raw methods on
}

// MigrationsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MigrationsTransactorRaw struct {
	Contract *MigrationsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMigrations creates a new instance of Migrations, bound to a specific deployed contract.
func NewMigrations(address common.Address, backend bind.ContractBackend) (*Migrations, error) {
	contract, err := bindMigrations(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Migrations{MigrationsCaller: MigrationsCaller{contract: contract}, MigrationsTransactor: MigrationsTransactor{contract: contract}}, nil
}

// NewMigrationsCaller creates a new read-only instance of Migrations, bound to a specific deployed contract.
func NewMigrationsCaller(address common.Address, caller bind.ContractCaller) (*MigrationsCaller, error) {
	contract, err := bindMigrations(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &MigrationsCaller{contract: contract}, nil
}

// NewMigrationsTransactor creates a new write-only instance of Migrations, bound to a specific deployed contract.
func NewMigrationsTransactor(address common.Address, transactor bind.ContractTransactor) (*MigrationsTransactor, error) {
	contract, err := bindMigrations(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &MigrationsTransactor{contract: contract}, nil
}

// bindMigrations binds a generic wrapper to an already deployed contract.
func bindMigrations(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MigrationsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Migrations *MigrationsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Migrations.Contract.MigrationsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Migrations *MigrationsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Migrations.Contract.MigrationsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Migrations *MigrationsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Migrations.Contract.MigrationsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Migrations *MigrationsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Migrations.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Migrations *MigrationsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Migrations.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Migrations *MigrationsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Migrations.Contract.contract.Transact(opts, method, params...)
}

// Last_completed_migration is a free data retrieval call binding the contract method 0x445df0ac.
//
// Solidity: function last_completed_migration() constant returns(uint256)
func (_Migrations *MigrationsCaller) Last_completed_migration(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Migrations.contract.Call(opts, out, "last_completed_migration")
	return *ret0, err
}

// Last_completed_migration is a free data retrieval call binding the contract method 0x445df0ac.
//
// Solidity: function last_completed_migration() constant returns(uint256)
func (_Migrations *MigrationsSession) Last_completed_migration() (*big.Int, error) {
	return _Migrations.Contract.Last_completed_migration(&_Migrations.CallOpts)
}

// Last_completed_migration is a free data retrieval call binding the contract method 0x445df0ac.
//
// Solidity: function last_completed_migration() constant returns(uint256)
func (_Migrations *MigrationsCallerSession) Last_completed_migration() (*big.Int, error) {
	return _Migrations.Contract.Last_completed_migration(&_Migrations.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Migrations *MigrationsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Migrations.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Migrations *MigrationsSession) Owner() (common.Address, error) {
	return _Migrations.Contract.Owner(&_Migrations.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Migrations *MigrationsCallerSession) Owner() (common.Address, error) {
	return _Migrations.Contract.Owner(&_Migrations.CallOpts)
}

// SetCompleted is a paid mutator transaction binding the contract method 0xfdacd576.
//
// Solidity: function setCompleted(completed uint256) returns()
func (_Migrations *MigrationsTransactor) SetCompleted(opts *bind.TransactOpts, completed *big.Int) (*types.Transaction, error) {
	return _Migrations.contract.Transact(opts, "setCompleted", completed)
}

// SetCompleted is a paid mutator transaction binding the contract method 0xfdacd576.
//
// Solidity: function setCompleted(completed uint256) returns()
func (_Migrations *MigrationsSession) SetCompleted(completed *big.Int) (*types.Transaction, error) {
	return _Migrations.Contract.SetCompleted(&_Migrations.TransactOpts, completed)
}

// SetCompleted is a paid mutator transaction binding the contract method 0xfdacd576.
//
// Solidity: function setCompleted(completed uint256) returns()
func (_Migrations *MigrationsTransactorSession) SetCompleted(completed *big.Int) (*types.Transaction, error) {
	return _Migrations.Contract.SetCompleted(&_Migrations.TransactOpts, completed)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(newAddress address) returns()
func (_Migrations *MigrationsTransactor) Upgrade(opts *bind.TransactOpts, newAddress common.Address) (*types.Transaction, error) {
	return _Migrations.contract.Transact(opts, "upgrade", newAddress)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(newAddress address) returns()
func (_Migrations *MigrationsSession) Upgrade(newAddress common.Address) (*types.Transaction, error) {
	return _Migrations.Contract.Upgrade(&_Migrations.TransactOpts, newAddress)
}

// Upgrade is a paid mutator transaction binding the contract method 0x0900f010.
//
// Solidity: function upgrade(newAddress address) returns()
func (_Migrations *MigrationsTransactorSession) Upgrade(newAddress common.Address) (*types.Transaction, error) {
	return _Migrations.Contract.Upgrade(&_Migrations.TransactOpts, newAddress)
}

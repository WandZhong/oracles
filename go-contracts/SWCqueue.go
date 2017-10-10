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

// SWCqueueABI is the input ABI used to generate the binding from.
const SWCqueueABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"senderHasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"restart\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint128\"},{\"name\":\"currency\",\"type\":\"bytes3\"}],\"name\":\"directPledge\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"logTrancheRelease\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"roles\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ratio\",\"type\":\"uint256\"}],\"name\":\"setNextBRGSWTratio\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"brg\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"brg_\",\"type\":\"address\"}],\"name\":\"setBRG\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wad\",\"type\":\"uint128\"},{\"name\":\"currency\",\"type\":\"bytes3\"}],\"name\":\"cancel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"brgAmount\",\"type\":\"uint128\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextBRGSWTratio\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"roles_\",\"type\":\"address\"}],\"name\":\"setRolesContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"hasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"brg_\",\"type\":\"address\"},{\"name\":\"nextBRGSWTratio_\",\"type\":\"uint256\"},{\"name\":\"rolesContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint128\"},{\"indexed\":false,\"name\":\"currency\",\"type\":\"bytes3\"}],\"name\":\"LogSWCqueueCancel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint128\"},{\"indexed\":false,\"name\":\"currency\",\"type\":\"bytes3\"}],\"name\":\"LogSWCqueueDirectPledge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"LogSWCqueueTranchRelease\",\"type\":\"event\"}]"

// SWCqueue is an auto generated Go binding around an Ethereum contract.
type SWCqueue struct {
	SWCqueueCaller     // Read-only binding to the contract
	SWCqueueTransactor // Write-only binding to the contract
}

// SWCqueueCaller is an auto generated read-only Go binding around an Ethereum contract.
type SWCqueueCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SWCqueueTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SWCqueueTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SWCqueueSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SWCqueueSession struct {
	Contract     *SWCqueue         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SWCqueueCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SWCqueueCallerSession struct {
	Contract *SWCqueueCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SWCqueueTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SWCqueueTransactorSession struct {
	Contract     *SWCqueueTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SWCqueueRaw is an auto generated low-level Go binding around an Ethereum contract.
type SWCqueueRaw struct {
	Contract *SWCqueue // Generic contract binding to access the raw methods on
}

// SWCqueueCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SWCqueueCallerRaw struct {
	Contract *SWCqueueCaller // Generic read-only contract binding to access the raw methods on
}

// SWCqueueTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SWCqueueTransactorRaw struct {
	Contract *SWCqueueTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSWCqueue creates a new instance of SWCqueue, bound to a specific deployed contract.
func NewSWCqueue(address common.Address, backend bind.ContractBackend) (*SWCqueue, error) {
	contract, err := bindSWCqueue(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SWCqueue{SWCqueueCaller: SWCqueueCaller{contract: contract}, SWCqueueTransactor: SWCqueueTransactor{contract: contract}}, nil
}

// NewSWCqueueCaller creates a new read-only instance of SWCqueue, bound to a specific deployed contract.
func NewSWCqueueCaller(address common.Address, caller bind.ContractCaller) (*SWCqueueCaller, error) {
	contract, err := bindSWCqueue(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &SWCqueueCaller{contract: contract}, nil
}

// NewSWCqueueTransactor creates a new write-only instance of SWCqueue, bound to a specific deployed contract.
func NewSWCqueueTransactor(address common.Address, transactor bind.ContractTransactor) (*SWCqueueTransactor, error) {
	contract, err := bindSWCqueue(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &SWCqueueTransactor{contract: contract}, nil
}

// bindSWCqueue binds a generic wrapper to an already deployed contract.
func bindSWCqueue(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SWCqueueABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SWCqueue *SWCqueueRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SWCqueue.Contract.SWCqueueCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SWCqueue *SWCqueueRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SWCqueue.Contract.SWCqueueTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SWCqueue *SWCqueueRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SWCqueue.Contract.SWCqueueTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SWCqueue *SWCqueueCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SWCqueue.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SWCqueue *SWCqueueTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SWCqueue.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SWCqueue *SWCqueueTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SWCqueue.Contract.contract.Transact(opts, method, params...)
}

// Brg is a free data retrieval call binding the contract method 0x4f9c8fe8.
//
// Solidity: function brg() constant returns(address)
func (_SWCqueue *SWCqueueCaller) Brg(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SWCqueue.contract.Call(opts, out, "brg")
	return *ret0, err
}

// Brg is a free data retrieval call binding the contract method 0x4f9c8fe8.
//
// Solidity: function brg() constant returns(address)
func (_SWCqueue *SWCqueueSession) Brg() (common.Address, error) {
	return _SWCqueue.Contract.Brg(&_SWCqueue.CallOpts)
}

// Brg is a free data retrieval call binding the contract method 0x4f9c8fe8.
//
// Solidity: function brg() constant returns(address)
func (_SWCqueue *SWCqueueCallerSession) Brg() (common.Address, error) {
	return _SWCqueue.Contract.Brg(&_SWCqueue.CallOpts)
}

// ContractHash is a free data retrieval call binding the contract method 0x904c6094.
//
// Solidity: function contractHash() constant returns(bytes32)
func (_SWCqueue *SWCqueueCaller) ContractHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _SWCqueue.contract.Call(opts, out, "contractHash")
	return *ret0, err
}

// ContractHash is a free data retrieval call binding the contract method 0x904c6094.
//
// Solidity: function contractHash() constant returns(bytes32)
func (_SWCqueue *SWCqueueSession) ContractHash() ([32]byte, error) {
	return _SWCqueue.Contract.ContractHash(&_SWCqueue.CallOpts)
}

// ContractHash is a free data retrieval call binding the contract method 0x904c6094.
//
// Solidity: function contractHash() constant returns(bytes32)
func (_SWCqueue *SWCqueueCallerSession) ContractHash() ([32]byte, error) {
	return _SWCqueue.Contract.ContractHash(&_SWCqueue.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0xe3c33a9b.
//
// Solidity: function hasRole(roleName string) constant returns(bool)
func (_SWCqueue *SWCqueueCaller) HasRole(opts *bind.CallOpts, roleName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SWCqueue.contract.Call(opts, out, "hasRole", roleName)
	return *ret0, err
}

// HasRole is a free data retrieval call binding the contract method 0xe3c33a9b.
//
// Solidity: function hasRole(roleName string) constant returns(bool)
func (_SWCqueue *SWCqueueSession) HasRole(roleName string) (bool, error) {
	return _SWCqueue.Contract.HasRole(&_SWCqueue.CallOpts, roleName)
}

// HasRole is a free data retrieval call binding the contract method 0xe3c33a9b.
//
// Solidity: function hasRole(roleName string) constant returns(bool)
func (_SWCqueue *SWCqueueCallerSession) HasRole(roleName string) (bool, error) {
	return _SWCqueue.Contract.HasRole(&_SWCqueue.CallOpts, roleName)
}

// NextBRGSWTratio is a free data retrieval call binding the contract method 0xac63f9d9.
//
// Solidity: function nextBRGSWTratio() constant returns(uint256)
func (_SWCqueue *SWCqueueCaller) NextBRGSWTratio(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SWCqueue.contract.Call(opts, out, "nextBRGSWTratio")
	return *ret0, err
}

// NextBRGSWTratio is a free data retrieval call binding the contract method 0xac63f9d9.
//
// Solidity: function nextBRGSWTratio() constant returns(uint256)
func (_SWCqueue *SWCqueueSession) NextBRGSWTratio() (*big.Int, error) {
	return _SWCqueue.Contract.NextBRGSWTratio(&_SWCqueue.CallOpts)
}

// NextBRGSWTratio is a free data retrieval call binding the contract method 0xac63f9d9.
//
// Solidity: function nextBRGSWTratio() constant returns(uint256)
func (_SWCqueue *SWCqueueCallerSession) NextBRGSWTratio() (*big.Int, error) {
	return _SWCqueue.Contract.NextBRGSWTratio(&_SWCqueue.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SWCqueue *SWCqueueCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SWCqueue.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SWCqueue *SWCqueueSession) Owner() (common.Address, error) {
	return _SWCqueue.Contract.Owner(&_SWCqueue.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SWCqueue *SWCqueueCallerSession) Owner() (common.Address, error) {
	return _SWCqueue.Contract.Owner(&_SWCqueue.CallOpts)
}

// Roles is a free data retrieval call binding the contract method 0x392f5f64.
//
// Solidity: function roles() constant returns(address)
func (_SWCqueue *SWCqueueCaller) Roles(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SWCqueue.contract.Call(opts, out, "roles")
	return *ret0, err
}

// Roles is a free data retrieval call binding the contract method 0x392f5f64.
//
// Solidity: function roles() constant returns(address)
func (_SWCqueue *SWCqueueSession) Roles() (common.Address, error) {
	return _SWCqueue.Contract.Roles(&_SWCqueue.CallOpts)
}

// Roles is a free data retrieval call binding the contract method 0x392f5f64.
//
// Solidity: function roles() constant returns(address)
func (_SWCqueue *SWCqueueCallerSession) Roles() (common.Address, error) {
	return _SWCqueue.Contract.Roles(&_SWCqueue.CallOpts)
}

// SenderHasRole is a free data retrieval call binding the contract method 0x1ca03b8e.
//
// Solidity: function senderHasRole(roleName string) constant returns(bool)
func (_SWCqueue *SWCqueueCaller) SenderHasRole(opts *bind.CallOpts, roleName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SWCqueue.contract.Call(opts, out, "senderHasRole", roleName)
	return *ret0, err
}

// SenderHasRole is a free data retrieval call binding the contract method 0x1ca03b8e.
//
// Solidity: function senderHasRole(roleName string) constant returns(bool)
func (_SWCqueue *SWCqueueSession) SenderHasRole(roleName string) (bool, error) {
	return _SWCqueue.Contract.SenderHasRole(&_SWCqueue.CallOpts, roleName)
}

// SenderHasRole is a free data retrieval call binding the contract method 0x1ca03b8e.
//
// Solidity: function senderHasRole(roleName string) constant returns(bool)
func (_SWCqueue *SWCqueueCallerSession) SenderHasRole(roleName string) (bool, error) {
	return _SWCqueue.Contract.SenderHasRole(&_SWCqueue.CallOpts, roleName)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_SWCqueue *SWCqueueCaller) Stopped(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SWCqueue.contract.Call(opts, out, "stopped")
	return *ret0, err
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_SWCqueue *SWCqueueSession) Stopped() (bool, error) {
	return _SWCqueue.Contract.Stopped(&_SWCqueue.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_SWCqueue *SWCqueueCallerSession) Stopped() (bool, error) {
	return _SWCqueue.Contract.Stopped(&_SWCqueue.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x90bc1693.
//
// Solidity: function burn(brgAmount uint128) returns()
func (_SWCqueue *SWCqueueTransactor) Burn(opts *bind.TransactOpts, brgAmount *big.Int) (*types.Transaction, error) {
	return _SWCqueue.contract.Transact(opts, "burn", brgAmount)
}

// Burn is a paid mutator transaction binding the contract method 0x90bc1693.
//
// Solidity: function burn(brgAmount uint128) returns()
func (_SWCqueue *SWCqueueSession) Burn(brgAmount *big.Int) (*types.Transaction, error) {
	return _SWCqueue.Contract.Burn(&_SWCqueue.TransactOpts, brgAmount)
}

// Burn is a paid mutator transaction binding the contract method 0x90bc1693.
//
// Solidity: function burn(brgAmount uint128) returns()
func (_SWCqueue *SWCqueueTransactorSession) Burn(brgAmount *big.Int) (*types.Transaction, error) {
	return _SWCqueue.Contract.Burn(&_SWCqueue.TransactOpts, brgAmount)
}

// Cancel is a paid mutator transaction binding the contract method 0x90a160d4.
//
// Solidity: function cancel(wad uint128, currency bytes3) returns()
func (_SWCqueue *SWCqueueTransactor) Cancel(opts *bind.TransactOpts, wad *big.Int, currency [3]byte) (*types.Transaction, error) {
	return _SWCqueue.contract.Transact(opts, "cancel", wad, currency)
}

// Cancel is a paid mutator transaction binding the contract method 0x90a160d4.
//
// Solidity: function cancel(wad uint128, currency bytes3) returns()
func (_SWCqueue *SWCqueueSession) Cancel(wad *big.Int, currency [3]byte) (*types.Transaction, error) {
	return _SWCqueue.Contract.Cancel(&_SWCqueue.TransactOpts, wad, currency)
}

// Cancel is a paid mutator transaction binding the contract method 0x90a160d4.
//
// Solidity: function cancel(wad uint128, currency bytes3) returns()
func (_SWCqueue *SWCqueueTransactorSession) Cancel(wad *big.Int, currency [3]byte) (*types.Transaction, error) {
	return _SWCqueue.Contract.Cancel(&_SWCqueue.TransactOpts, wad, currency)
}

// DirectPledge is a paid mutator transaction binding the contract method 0x209c4986.
//
// Solidity: function directPledge(who address, wad uint128, currency bytes3) returns()
func (_SWCqueue *SWCqueueTransactor) DirectPledge(opts *bind.TransactOpts, who common.Address, wad *big.Int, currency [3]byte) (*types.Transaction, error) {
	return _SWCqueue.contract.Transact(opts, "directPledge", who, wad, currency)
}

// DirectPledge is a paid mutator transaction binding the contract method 0x209c4986.
//
// Solidity: function directPledge(who address, wad uint128, currency bytes3) returns()
func (_SWCqueue *SWCqueueSession) DirectPledge(who common.Address, wad *big.Int, currency [3]byte) (*types.Transaction, error) {
	return _SWCqueue.Contract.DirectPledge(&_SWCqueue.TransactOpts, who, wad, currency)
}

// DirectPledge is a paid mutator transaction binding the contract method 0x209c4986.
//
// Solidity: function directPledge(who address, wad uint128, currency bytes3) returns()
func (_SWCqueue *SWCqueueTransactorSession) DirectPledge(who common.Address, wad *big.Int, currency [3]byte) (*types.Transaction, error) {
	return _SWCqueue.Contract.DirectPledge(&_SWCqueue.TransactOpts, who, wad, currency)
}

// LogTrancheRelease is a paid mutator transaction binding the contract method 0x30311449.
//
// Solidity: function logTrancheRelease(wad uint128) returns()
func (_SWCqueue *SWCqueueTransactor) LogTrancheRelease(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _SWCqueue.contract.Transact(opts, "logTrancheRelease", wad)
}

// LogTrancheRelease is a paid mutator transaction binding the contract method 0x30311449.
//
// Solidity: function logTrancheRelease(wad uint128) returns()
func (_SWCqueue *SWCqueueSession) LogTrancheRelease(wad *big.Int) (*types.Transaction, error) {
	return _SWCqueue.Contract.LogTrancheRelease(&_SWCqueue.TransactOpts, wad)
}

// LogTrancheRelease is a paid mutator transaction binding the contract method 0x30311449.
//
// Solidity: function logTrancheRelease(wad uint128) returns()
func (_SWCqueue *SWCqueueTransactorSession) LogTrancheRelease(wad *big.Int) (*types.Transaction, error) {
	return _SWCqueue.Contract.LogTrancheRelease(&_SWCqueue.TransactOpts, wad)
}

// Restart is a paid mutator transaction binding the contract method 0x1ef3755d.
//
// Solidity: function restart() returns()
func (_SWCqueue *SWCqueueTransactor) Restart(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SWCqueue.contract.Transact(opts, "restart")
}

// Restart is a paid mutator transaction binding the contract method 0x1ef3755d.
//
// Solidity: function restart() returns()
func (_SWCqueue *SWCqueueSession) Restart() (*types.Transaction, error) {
	return _SWCqueue.Contract.Restart(&_SWCqueue.TransactOpts)
}

// Restart is a paid mutator transaction binding the contract method 0x1ef3755d.
//
// Solidity: function restart() returns()
func (_SWCqueue *SWCqueueTransactorSession) Restart() (*types.Transaction, error) {
	return _SWCqueue.Contract.Restart(&_SWCqueue.TransactOpts)
}

// SetBRG is a paid mutator transaction binding the contract method 0x679ccc6d.
//
// Solidity: function setBRG(brg_ address) returns()
func (_SWCqueue *SWCqueueTransactor) SetBRG(opts *bind.TransactOpts, brg_ common.Address) (*types.Transaction, error) {
	return _SWCqueue.contract.Transact(opts, "setBRG", brg_)
}

// SetBRG is a paid mutator transaction binding the contract method 0x679ccc6d.
//
// Solidity: function setBRG(brg_ address) returns()
func (_SWCqueue *SWCqueueSession) SetBRG(brg_ common.Address) (*types.Transaction, error) {
	return _SWCqueue.Contract.SetBRG(&_SWCqueue.TransactOpts, brg_)
}

// SetBRG is a paid mutator transaction binding the contract method 0x679ccc6d.
//
// Solidity: function setBRG(brg_ address) returns()
func (_SWCqueue *SWCqueueTransactorSession) SetBRG(brg_ common.Address) (*types.Transaction, error) {
	return _SWCqueue.Contract.SetBRG(&_SWCqueue.TransactOpts, brg_)
}

// SetNextBRGSWTratio is a paid mutator transaction binding the contract method 0x429ad4f2.
//
// Solidity: function setNextBRGSWTratio(ratio uint256) returns()
func (_SWCqueue *SWCqueueTransactor) SetNextBRGSWTratio(opts *bind.TransactOpts, ratio *big.Int) (*types.Transaction, error) {
	return _SWCqueue.contract.Transact(opts, "setNextBRGSWTratio", ratio)
}

// SetNextBRGSWTratio is a paid mutator transaction binding the contract method 0x429ad4f2.
//
// Solidity: function setNextBRGSWTratio(ratio uint256) returns()
func (_SWCqueue *SWCqueueSession) SetNextBRGSWTratio(ratio *big.Int) (*types.Transaction, error) {
	return _SWCqueue.Contract.SetNextBRGSWTratio(&_SWCqueue.TransactOpts, ratio)
}

// SetNextBRGSWTratio is a paid mutator transaction binding the contract method 0x429ad4f2.
//
// Solidity: function setNextBRGSWTratio(ratio uint256) returns()
func (_SWCqueue *SWCqueueTransactorSession) SetNextBRGSWTratio(ratio *big.Int) (*types.Transaction, error) {
	return _SWCqueue.Contract.SetNextBRGSWTratio(&_SWCqueue.TransactOpts, ratio)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_SWCqueue *SWCqueueTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _SWCqueue.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_SWCqueue *SWCqueueSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _SWCqueue.Contract.SetOwner(&_SWCqueue.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_SWCqueue *SWCqueueTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _SWCqueue.Contract.SetOwner(&_SWCqueue.TransactOpts, owner_)
}

// SetRolesContract is a paid mutator transaction binding the contract method 0xafa202ac.
//
// Solidity: function setRolesContract(roles_ address) returns()
func (_SWCqueue *SWCqueueTransactor) SetRolesContract(opts *bind.TransactOpts, roles_ common.Address) (*types.Transaction, error) {
	return _SWCqueue.contract.Transact(opts, "setRolesContract", roles_)
}

// SetRolesContract is a paid mutator transaction binding the contract method 0xafa202ac.
//
// Solidity: function setRolesContract(roles_ address) returns()
func (_SWCqueue *SWCqueueSession) SetRolesContract(roles_ common.Address) (*types.Transaction, error) {
	return _SWCqueue.Contract.SetRolesContract(&_SWCqueue.TransactOpts, roles_)
}

// SetRolesContract is a paid mutator transaction binding the contract method 0xafa202ac.
//
// Solidity: function setRolesContract(roles_ address) returns()
func (_SWCqueue *SWCqueueTransactorSession) SetRolesContract(roles_ common.Address) (*types.Transaction, error) {
	return _SWCqueue.Contract.SetRolesContract(&_SWCqueue.TransactOpts, roles_)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_SWCqueue *SWCqueueTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SWCqueue.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_SWCqueue *SWCqueueSession) Stop() (*types.Transaction, error) {
	return _SWCqueue.Contract.Stop(&_SWCqueue.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_SWCqueue *SWCqueueTransactorSession) Stop() (*types.Transaction, error) {
	return _SWCqueue.Contract.Stop(&_SWCqueue.TransactOpts)
}

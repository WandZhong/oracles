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

// SweetTokenABI is the input ABI used to generate the binding from.
const SweetTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guy\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"senderHasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"restart\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"mintFor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"roles\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"logic_\",\"type\":\"address\"}],\"name\":\"setLogic\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"pull\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"roles_\",\"type\":\"address\"}],\"name\":\"setRolesContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name_\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"logic\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"_allowance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"hasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"name_\",\"type\":\"string\"},{\"name\":\"symbol_\",\"type\":\"bytes32\"},{\"name\":\"rolesContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"src\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"LogBurn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"src\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"LogMint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// SweetToken is an auto generated Go binding around an Ethereum contract.
type SweetToken struct {
	SweetTokenCaller     // Read-only binding to the contract
	SweetTokenTransactor // Write-only binding to the contract
}

// SweetTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type SweetTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SweetTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SweetTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SweetTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SweetTokenSession struct {
	Contract     *SweetToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SweetTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SweetTokenCallerSession struct {
	Contract *SweetTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SweetTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SweetTokenTransactorSession struct {
	Contract     *SweetTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SweetTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type SweetTokenRaw struct {
	Contract *SweetToken // Generic contract binding to access the raw methods on
}

// SweetTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SweetTokenCallerRaw struct {
	Contract *SweetTokenCaller // Generic read-only contract binding to access the raw methods on
}

// SweetTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SweetTokenTransactorRaw struct {
	Contract *SweetTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSweetToken creates a new instance of SweetToken, bound to a specific deployed contract.
func NewSweetToken(address common.Address, backend bind.ContractBackend) (*SweetToken, error) {
	contract, err := bindSweetToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SweetToken{SweetTokenCaller: SweetTokenCaller{contract: contract}, SweetTokenTransactor: SweetTokenTransactor{contract: contract}}, nil
}

// NewSweetTokenCaller creates a new read-only instance of SweetToken, bound to a specific deployed contract.
func NewSweetTokenCaller(address common.Address, caller bind.ContractCaller) (*SweetTokenCaller, error) {
	contract, err := bindSweetToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &SweetTokenCaller{contract: contract}, nil
}

// NewSweetTokenTransactor creates a new write-only instance of SweetToken, bound to a specific deployed contract.
func NewSweetTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*SweetTokenTransactor, error) {
	contract, err := bindSweetToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &SweetTokenTransactor{contract: contract}, nil
}

// bindSweetToken binds a generic wrapper to an already deployed contract.
func bindSweetToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SweetTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SweetToken *SweetTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SweetToken.Contract.SweetTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SweetToken *SweetTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SweetToken.Contract.SweetTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SweetToken *SweetTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SweetToken.Contract.SweetTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SweetToken *SweetTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SweetToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SweetToken *SweetTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SweetToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SweetToken *SweetTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SweetToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(_allowance uint256)
func (_SweetToken *SweetTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SweetToken.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(_allowance uint256)
func (_SweetToken *SweetTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SweetToken.Contract.Allowance(&_SweetToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(_allowance uint256)
func (_SweetToken *SweetTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SweetToken.Contract.Allowance(&_SweetToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(value uint256)
func (_SweetToken *SweetTokenCaller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SweetToken.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(value uint256)
func (_SweetToken *SweetTokenSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _SweetToken.Contract.BalanceOf(&_SweetToken.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(value uint256)
func (_SweetToken *SweetTokenCallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _SweetToken.Contract.BalanceOf(&_SweetToken.CallOpts, who)
}

// ContractHash is a free data retrieval call binding the contract method 0x904c6094.
//
// Solidity: function contractHash() constant returns(bytes32)
func (_SweetToken *SweetTokenCaller) ContractHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _SweetToken.contract.Call(opts, out, "contractHash")
	return *ret0, err
}

// ContractHash is a free data retrieval call binding the contract method 0x904c6094.
//
// Solidity: function contractHash() constant returns(bytes32)
func (_SweetToken *SweetTokenSession) ContractHash() ([32]byte, error) {
	return _SweetToken.Contract.ContractHash(&_SweetToken.CallOpts)
}

// ContractHash is a free data retrieval call binding the contract method 0x904c6094.
//
// Solidity: function contractHash() constant returns(bytes32)
func (_SweetToken *SweetTokenCallerSession) ContractHash() ([32]byte, error) {
	return _SweetToken.Contract.ContractHash(&_SweetToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_SweetToken *SweetTokenCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SweetToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_SweetToken *SweetTokenSession) Decimals() (*big.Int, error) {
	return _SweetToken.Contract.Decimals(&_SweetToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_SweetToken *SweetTokenCallerSession) Decimals() (*big.Int, error) {
	return _SweetToken.Contract.Decimals(&_SweetToken.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0xe3c33a9b.
//
// Solidity: function hasRole(roleName string) constant returns(bool)
func (_SweetToken *SweetTokenCaller) HasRole(opts *bind.CallOpts, roleName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SweetToken.contract.Call(opts, out, "hasRole", roleName)
	return *ret0, err
}

// HasRole is a free data retrieval call binding the contract method 0xe3c33a9b.
//
// Solidity: function hasRole(roleName string) constant returns(bool)
func (_SweetToken *SweetTokenSession) HasRole(roleName string) (bool, error) {
	return _SweetToken.Contract.HasRole(&_SweetToken.CallOpts, roleName)
}

// HasRole is a free data retrieval call binding the contract method 0xe3c33a9b.
//
// Solidity: function hasRole(roleName string) constant returns(bool)
func (_SweetToken *SweetTokenCallerSession) HasRole(roleName string) (bool, error) {
	return _SweetToken.Contract.HasRole(&_SweetToken.CallOpts, roleName)
}

// Logic is a free data retrieval call binding the contract method 0xd7dfa0dd.
//
// Solidity: function logic() constant returns(address)
func (_SweetToken *SweetTokenCaller) Logic(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SweetToken.contract.Call(opts, out, "logic")
	return *ret0, err
}

// Logic is a free data retrieval call binding the contract method 0xd7dfa0dd.
//
// Solidity: function logic() constant returns(address)
func (_SweetToken *SweetTokenSession) Logic() (common.Address, error) {
	return _SweetToken.Contract.Logic(&_SweetToken.CallOpts)
}

// Logic is a free data retrieval call binding the contract method 0xd7dfa0dd.
//
// Solidity: function logic() constant returns(address)
func (_SweetToken *SweetTokenCallerSession) Logic() (common.Address, error) {
	return _SweetToken.Contract.Logic(&_SweetToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SweetToken *SweetTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SweetToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SweetToken *SweetTokenSession) Name() (string, error) {
	return _SweetToken.Contract.Name(&_SweetToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SweetToken *SweetTokenCallerSession) Name() (string, error) {
	return _SweetToken.Contract.Name(&_SweetToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SweetToken *SweetTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SweetToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SweetToken *SweetTokenSession) Owner() (common.Address, error) {
	return _SweetToken.Contract.Owner(&_SweetToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SweetToken *SweetTokenCallerSession) Owner() (common.Address, error) {
	return _SweetToken.Contract.Owner(&_SweetToken.CallOpts)
}

// Roles is a free data retrieval call binding the contract method 0x392f5f64.
//
// Solidity: function roles() constant returns(address)
func (_SweetToken *SweetTokenCaller) Roles(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SweetToken.contract.Call(opts, out, "roles")
	return *ret0, err
}

// Roles is a free data retrieval call binding the contract method 0x392f5f64.
//
// Solidity: function roles() constant returns(address)
func (_SweetToken *SweetTokenSession) Roles() (common.Address, error) {
	return _SweetToken.Contract.Roles(&_SweetToken.CallOpts)
}

// Roles is a free data retrieval call binding the contract method 0x392f5f64.
//
// Solidity: function roles() constant returns(address)
func (_SweetToken *SweetTokenCallerSession) Roles() (common.Address, error) {
	return _SweetToken.Contract.Roles(&_SweetToken.CallOpts)
}

// SenderHasRole is a free data retrieval call binding the contract method 0x1ca03b8e.
//
// Solidity: function senderHasRole(roleName string) constant returns(bool)
func (_SweetToken *SweetTokenCaller) SenderHasRole(opts *bind.CallOpts, roleName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SweetToken.contract.Call(opts, out, "senderHasRole", roleName)
	return *ret0, err
}

// SenderHasRole is a free data retrieval call binding the contract method 0x1ca03b8e.
//
// Solidity: function senderHasRole(roleName string) constant returns(bool)
func (_SweetToken *SweetTokenSession) SenderHasRole(roleName string) (bool, error) {
	return _SweetToken.Contract.SenderHasRole(&_SweetToken.CallOpts, roleName)
}

// SenderHasRole is a free data retrieval call binding the contract method 0x1ca03b8e.
//
// Solidity: function senderHasRole(roleName string) constant returns(bool)
func (_SweetToken *SweetTokenCallerSession) SenderHasRole(roleName string) (bool, error) {
	return _SweetToken.Contract.SenderHasRole(&_SweetToken.CallOpts, roleName)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_SweetToken *SweetTokenCaller) Stopped(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SweetToken.contract.Call(opts, out, "stopped")
	return *ret0, err
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_SweetToken *SweetTokenSession) Stopped() (bool, error) {
	return _SweetToken.Contract.Stopped(&_SweetToken.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_SweetToken *SweetTokenCallerSession) Stopped() (bool, error) {
	return _SweetToken.Contract.Stopped(&_SweetToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(bytes32)
func (_SweetToken *SweetTokenCaller) Symbol(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _SweetToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(bytes32)
func (_SweetToken *SweetTokenSession) Symbol() ([32]byte, error) {
	return _SweetToken.Contract.Symbol(&_SweetToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(bytes32)
func (_SweetToken *SweetTokenCallerSession) Symbol() ([32]byte, error) {
	return _SweetToken.Contract.Symbol(&_SweetToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SweetToken *SweetTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SweetToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SweetToken *SweetTokenSession) TotalSupply() (*big.Int, error) {
	return _SweetToken.Contract.TotalSupply(&_SweetToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SweetToken *SweetTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _SweetToken.Contract.TotalSupply(&_SweetToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(guy address, wad uint256) returns(bool)
func (_SweetToken *SweetTokenTransactor) Approve(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetToken.contract.Transact(opts, "approve", guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(guy address, wad uint256) returns(bool)
func (_SweetToken *SweetTokenSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetToken.Contract.Approve(&_SweetToken.TransactOpts, guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(guy address, wad uint256) returns(bool)
func (_SweetToken *SweetTokenTransactorSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetToken.Contract.Approve(&_SweetToken.TransactOpts, guy, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x90bc1693.
//
// Solidity: function burn(wad uint128) returns()
func (_SweetToken *SweetTokenTransactor) Burn(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _SweetToken.contract.Transact(opts, "burn", wad)
}

// Burn is a paid mutator transaction binding the contract method 0x90bc1693.
//
// Solidity: function burn(wad uint128) returns()
func (_SweetToken *SweetTokenSession) Burn(wad *big.Int) (*types.Transaction, error) {
	return _SweetToken.Contract.Burn(&_SweetToken.TransactOpts, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x90bc1693.
//
// Solidity: function burn(wad uint128) returns()
func (_SweetToken *SweetTokenTransactorSession) Burn(wad *big.Int) (*types.Transaction, error) {
	return _SweetToken.Contract.Burn(&_SweetToken.TransactOpts, wad)
}

// MintFor is a paid mutator transaction binding the contract method 0x27cfc219.
//
// Solidity: function mintFor(recipient address, wad uint128) returns()
func (_SweetToken *SweetTokenTransactor) MintFor(opts *bind.TransactOpts, recipient common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetToken.contract.Transact(opts, "mintFor", recipient, wad)
}

// MintFor is a paid mutator transaction binding the contract method 0x27cfc219.
//
// Solidity: function mintFor(recipient address, wad uint128) returns()
func (_SweetToken *SweetTokenSession) MintFor(recipient common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetToken.Contract.MintFor(&_SweetToken.TransactOpts, recipient, wad)
}

// MintFor is a paid mutator transaction binding the contract method 0x27cfc219.
//
// Solidity: function mintFor(recipient address, wad uint128) returns()
func (_SweetToken *SweetTokenTransactorSession) MintFor(recipient common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetToken.Contract.MintFor(&_SweetToken.TransactOpts, recipient, wad)
}

// Pull is a paid mutator transaction binding the contract method 0x8402181f.
//
// Solidity: function pull(src address, wad uint128) returns(bool)
func (_SweetToken *SweetTokenTransactor) Pull(opts *bind.TransactOpts, src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetToken.contract.Transact(opts, "pull", src, wad)
}

// Pull is a paid mutator transaction binding the contract method 0x8402181f.
//
// Solidity: function pull(src address, wad uint128) returns(bool)
func (_SweetToken *SweetTokenSession) Pull(src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetToken.Contract.Pull(&_SweetToken.TransactOpts, src, wad)
}

// Pull is a paid mutator transaction binding the contract method 0x8402181f.
//
// Solidity: function pull(src address, wad uint128) returns(bool)
func (_SweetToken *SweetTokenTransactorSession) Pull(src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetToken.Contract.Pull(&_SweetToken.TransactOpts, src, wad)
}

// Restart is a paid mutator transaction binding the contract method 0x1ef3755d.
//
// Solidity: function restart() returns()
func (_SweetToken *SweetTokenTransactor) Restart(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SweetToken.contract.Transact(opts, "restart")
}

// Restart is a paid mutator transaction binding the contract method 0x1ef3755d.
//
// Solidity: function restart() returns()
func (_SweetToken *SweetTokenSession) Restart() (*types.Transaction, error) {
	return _SweetToken.Contract.Restart(&_SweetToken.TransactOpts)
}

// Restart is a paid mutator transaction binding the contract method 0x1ef3755d.
//
// Solidity: function restart() returns()
func (_SweetToken *SweetTokenTransactorSession) Restart() (*types.Transaction, error) {
	return _SweetToken.Contract.Restart(&_SweetToken.TransactOpts)
}

// SetLogic is a paid mutator transaction binding the contract method 0x71857000.
//
// Solidity: function setLogic(logic_ address) returns(bool)
func (_SweetToken *SweetTokenTransactor) SetLogic(opts *bind.TransactOpts, logic_ common.Address) (*types.Transaction, error) {
	return _SweetToken.contract.Transact(opts, "setLogic", logic_)
}

// SetLogic is a paid mutator transaction binding the contract method 0x71857000.
//
// Solidity: function setLogic(logic_ address) returns(bool)
func (_SweetToken *SweetTokenSession) SetLogic(logic_ common.Address) (*types.Transaction, error) {
	return _SweetToken.Contract.SetLogic(&_SweetToken.TransactOpts, logic_)
}

// SetLogic is a paid mutator transaction binding the contract method 0x71857000.
//
// Solidity: function setLogic(logic_ address) returns(bool)
func (_SweetToken *SweetTokenTransactorSession) SetLogic(logic_ common.Address) (*types.Transaction, error) {
	return _SweetToken.Contract.SetLogic(&_SweetToken.TransactOpts, logic_)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(name_ string) returns()
func (_SweetToken *SweetTokenTransactor) SetName(opts *bind.TransactOpts, name_ string) (*types.Transaction, error) {
	return _SweetToken.contract.Transact(opts, "setName", name_)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(name_ string) returns()
func (_SweetToken *SweetTokenSession) SetName(name_ string) (*types.Transaction, error) {
	return _SweetToken.Contract.SetName(&_SweetToken.TransactOpts, name_)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(name_ string) returns()
func (_SweetToken *SweetTokenTransactorSession) SetName(name_ string) (*types.Transaction, error) {
	return _SweetToken.Contract.SetName(&_SweetToken.TransactOpts, name_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_SweetToken *SweetTokenTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _SweetToken.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_SweetToken *SweetTokenSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _SweetToken.Contract.SetOwner(&_SweetToken.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_SweetToken *SweetTokenTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _SweetToken.Contract.SetOwner(&_SweetToken.TransactOpts, owner_)
}

// SetRolesContract is a paid mutator transaction binding the contract method 0xafa202ac.
//
// Solidity: function setRolesContract(roles_ address) returns()
func (_SweetToken *SweetTokenTransactor) SetRolesContract(opts *bind.TransactOpts, roles_ common.Address) (*types.Transaction, error) {
	return _SweetToken.contract.Transact(opts, "setRolesContract", roles_)
}

// SetRolesContract is a paid mutator transaction binding the contract method 0xafa202ac.
//
// Solidity: function setRolesContract(roles_ address) returns()
func (_SweetToken *SweetTokenSession) SetRolesContract(roles_ common.Address) (*types.Transaction, error) {
	return _SweetToken.Contract.SetRolesContract(&_SweetToken.TransactOpts, roles_)
}

// SetRolesContract is a paid mutator transaction binding the contract method 0xafa202ac.
//
// Solidity: function setRolesContract(roles_ address) returns()
func (_SweetToken *SweetTokenTransactorSession) SetRolesContract(roles_ common.Address) (*types.Transaction, error) {
	return _SweetToken.Contract.SetRolesContract(&_SweetToken.TransactOpts, roles_)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_SweetToken *SweetTokenTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SweetToken.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_SweetToken *SweetTokenSession) Stop() (*types.Transaction, error) {
	return _SweetToken.Contract.Stop(&_SweetToken.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_SweetToken *SweetTokenTransactorSession) Stop() (*types.Transaction, error) {
	return _SweetToken.Contract.Stop(&_SweetToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(dst address, wad uint256) returns(bool)
func (_SweetToken *SweetTokenTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetToken.contract.Transact(opts, "transfer", dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(dst address, wad uint256) returns(bool)
func (_SweetToken *SweetTokenSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetToken.Contract.Transfer(&_SweetToken.TransactOpts, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(dst address, wad uint256) returns(bool)
func (_SweetToken *SweetTokenTransactorSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetToken.Contract.Transfer(&_SweetToken.TransactOpts, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_SweetToken *SweetTokenTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetToken.contract.Transact(opts, "transferFrom", src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_SweetToken *SweetTokenSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetToken.Contract.TransferFrom(&_SweetToken.TransactOpts, src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_SweetToken *SweetTokenTransactorSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetToken.Contract.TransferFrom(&_SweetToken.TransactOpts, src, dst, wad)
}

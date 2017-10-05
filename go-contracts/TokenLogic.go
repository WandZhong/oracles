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

// TokenLogicABI is the input ABI used to generate the binding from.
const TokenLogicABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token_\",\"type\":\"address\"}],\"name\":\"setToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"senderHasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"restart\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"listNamesLen\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"freeTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"mintFor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"isFree\",\"type\":\"bool\"}],\"name\":\"setFreeTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"roles\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"listName\",\"type\":\"bytes32\"}],\"name\":\"addWhiteList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"listName\",\"type\":\"bytes32\"},{\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"removeFromWhiteList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"whiteLists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"listNames\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"listName\",\"type\":\"bytes32\"},{\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"addToWhiteList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"listName\",\"type\":\"bytes32\"}],\"name\":\"listExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"roles_\",\"type\":\"address\"}],\"name\":\"setRolesContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"hasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"token_\",\"type\":\"address\"},{\"name\":\"data_\",\"type\":\"address\"},{\"name\":\"supply_\",\"type\":\"uint256\"},{\"name\":\"rolesContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"}]"

// TokenLogicBin is the compiled bytecode used for deploying new contracts.
const TokenLogicBin = `undefined`

// DeployTokenLogic deploys a new Ethereum contract, binding an instance of TokenLogic to it.
func DeployTokenLogic(auth *bind.TransactOpts, backend bind.ContractBackend, token_ common.Address, data_ common.Address, supply_ *big.Int, rolesContract common.Address) (common.Address, *types.Transaction, *TokenLogic, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenLogicABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenLogicBin), backend, token_, data_, supply_, rolesContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenLogic{TokenLogicCaller: TokenLogicCaller{contract: contract}, TokenLogicTransactor: TokenLogicTransactor{contract: contract}}, nil
}

// TokenLogic is an auto generated Go binding around an Ethereum contract.
type TokenLogic struct {
	TokenLogicCaller     // Read-only binding to the contract
	TokenLogicTransactor // Write-only binding to the contract
}

// TokenLogicCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenLogicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenLogicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenLogicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenLogicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenLogicSession struct {
	Contract     *TokenLogic       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenLogicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenLogicCallerSession struct {
	Contract *TokenLogicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TokenLogicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenLogicTransactorSession struct {
	Contract     *TokenLogicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TokenLogicRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenLogicRaw struct {
	Contract *TokenLogic // Generic contract binding to access the raw methods on
}

// TokenLogicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenLogicCallerRaw struct {
	Contract *TokenLogicCaller // Generic read-only contract binding to access the raw methods on
}

// TokenLogicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenLogicTransactorRaw struct {
	Contract *TokenLogicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenLogic creates a new instance of TokenLogic, bound to a specific deployed contract.
func NewTokenLogic(address common.Address, backend bind.ContractBackend) (*TokenLogic, error) {
	contract, err := bindTokenLogic(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenLogic{TokenLogicCaller: TokenLogicCaller{contract: contract}, TokenLogicTransactor: TokenLogicTransactor{contract: contract}}, nil
}

// NewTokenLogicCaller creates a new read-only instance of TokenLogic, bound to a specific deployed contract.
func NewTokenLogicCaller(address common.Address, caller bind.ContractCaller) (*TokenLogicCaller, error) {
	contract, err := bindTokenLogic(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TokenLogicCaller{contract: contract}, nil
}

// NewTokenLogicTransactor creates a new write-only instance of TokenLogic, bound to a specific deployed contract.
func NewTokenLogicTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenLogicTransactor, error) {
	contract, err := bindTokenLogic(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TokenLogicTransactor{contract: contract}, nil
}

// bindTokenLogic binds a generic wrapper to an already deployed contract.
func bindTokenLogic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenLogicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenLogic *TokenLogicRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenLogic.Contract.TokenLogicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenLogic *TokenLogicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenLogic.Contract.TokenLogicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenLogic *TokenLogicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenLogic.Contract.TokenLogicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenLogic *TokenLogicCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenLogic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenLogic *TokenLogicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenLogic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenLogic *TokenLogicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenLogic.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(src address, spender address) constant returns(uint256)
func (_TokenLogic *TokenLogicCaller) Allowance(opts *bind.CallOpts, src common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenLogic.contract.Call(opts, out, "allowance", src, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(src address, spender address) constant returns(uint256)
func (_TokenLogic *TokenLogicSession) Allowance(src common.Address, spender common.Address) (*big.Int, error) {
	return _TokenLogic.Contract.Allowance(&_TokenLogic.CallOpts, src, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(src address, spender address) constant returns(uint256)
func (_TokenLogic *TokenLogicCallerSession) Allowance(src common.Address, spender common.Address) (*big.Int, error) {
	return _TokenLogic.Contract.Allowance(&_TokenLogic.CallOpts, src, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(src address) constant returns(uint256)
func (_TokenLogic *TokenLogicCaller) BalanceOf(opts *bind.CallOpts, src common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenLogic.contract.Call(opts, out, "balanceOf", src)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(src address) constant returns(uint256)
func (_TokenLogic *TokenLogicSession) BalanceOf(src common.Address) (*big.Int, error) {
	return _TokenLogic.Contract.BalanceOf(&_TokenLogic.CallOpts, src)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(src address) constant returns(uint256)
func (_TokenLogic *TokenLogicCallerSession) BalanceOf(src common.Address) (*big.Int, error) {
	return _TokenLogic.Contract.BalanceOf(&_TokenLogic.CallOpts, src)
}

// ContractHash is a free data retrieval call binding the contract method 0x904c6094.
//
// Solidity: function contractHash() constant returns(bytes32)
func (_TokenLogic *TokenLogicCaller) ContractHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _TokenLogic.contract.Call(opts, out, "contractHash")
	return *ret0, err
}

// ContractHash is a free data retrieval call binding the contract method 0x904c6094.
//
// Solidity: function contractHash() constant returns(bytes32)
func (_TokenLogic *TokenLogicSession) ContractHash() ([32]byte, error) {
	return _TokenLogic.Contract.ContractHash(&_TokenLogic.CallOpts)
}

// ContractHash is a free data retrieval call binding the contract method 0x904c6094.
//
// Solidity: function contractHash() constant returns(bytes32)
func (_TokenLogic *TokenLogicCallerSession) ContractHash() ([32]byte, error) {
	return _TokenLogic.Contract.ContractHash(&_TokenLogic.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(address)
func (_TokenLogic *TokenLogicCaller) Data(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenLogic.contract.Call(opts, out, "data")
	return *ret0, err
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(address)
func (_TokenLogic *TokenLogicSession) Data() (common.Address, error) {
	return _TokenLogic.Contract.Data(&_TokenLogic.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(address)
func (_TokenLogic *TokenLogicCallerSession) Data() (common.Address, error) {
	return _TokenLogic.Contract.Data(&_TokenLogic.CallOpts)
}

// FreeTransfer is a free data retrieval call binding the contract method 0x211ed6c1.
//
// Solidity: function freeTransfer() constant returns(bool)
func (_TokenLogic *TokenLogicCaller) FreeTransfer(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenLogic.contract.Call(opts, out, "freeTransfer")
	return *ret0, err
}

// FreeTransfer is a free data retrieval call binding the contract method 0x211ed6c1.
//
// Solidity: function freeTransfer() constant returns(bool)
func (_TokenLogic *TokenLogicSession) FreeTransfer() (bool, error) {
	return _TokenLogic.Contract.FreeTransfer(&_TokenLogic.CallOpts)
}

// FreeTransfer is a free data retrieval call binding the contract method 0x211ed6c1.
//
// Solidity: function freeTransfer() constant returns(bool)
func (_TokenLogic *TokenLogicCallerSession) FreeTransfer() (bool, error) {
	return _TokenLogic.Contract.FreeTransfer(&_TokenLogic.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0xe3c33a9b.
//
// Solidity: function hasRole(roleName string) constant returns(bool)
func (_TokenLogic *TokenLogicCaller) HasRole(opts *bind.CallOpts, roleName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenLogic.contract.Call(opts, out, "hasRole", roleName)
	return *ret0, err
}

// HasRole is a free data retrieval call binding the contract method 0xe3c33a9b.
//
// Solidity: function hasRole(roleName string) constant returns(bool)
func (_TokenLogic *TokenLogicSession) HasRole(roleName string) (bool, error) {
	return _TokenLogic.Contract.HasRole(&_TokenLogic.CallOpts, roleName)
}

// HasRole is a free data retrieval call binding the contract method 0xe3c33a9b.
//
// Solidity: function hasRole(roleName string) constant returns(bool)
func (_TokenLogic *TokenLogicCallerSession) HasRole(roleName string) (bool, error) {
	return _TokenLogic.Contract.HasRole(&_TokenLogic.CallOpts, roleName)
}

// ListExists is a free data retrieval call binding the contract method 0x9495ad6c.
//
// Solidity: function listExists(listName bytes32) constant returns(bool)
func (_TokenLogic *TokenLogicCaller) ListExists(opts *bind.CallOpts, listName [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenLogic.contract.Call(opts, out, "listExists", listName)
	return *ret0, err
}

// ListExists is a free data retrieval call binding the contract method 0x9495ad6c.
//
// Solidity: function listExists(listName bytes32) constant returns(bool)
func (_TokenLogic *TokenLogicSession) ListExists(listName [32]byte) (bool, error) {
	return _TokenLogic.Contract.ListExists(&_TokenLogic.CallOpts, listName)
}

// ListExists is a free data retrieval call binding the contract method 0x9495ad6c.
//
// Solidity: function listExists(listName bytes32) constant returns(bool)
func (_TokenLogic *TokenLogicCallerSession) ListExists(listName [32]byte) (bool, error) {
	return _TokenLogic.Contract.ListExists(&_TokenLogic.CallOpts, listName)
}

// ListNames is a free data retrieval call binding the contract method 0x66272623.
//
// Solidity: function listNames( uint256) constant returns(bytes32)
func (_TokenLogic *TokenLogicCaller) ListNames(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _TokenLogic.contract.Call(opts, out, "listNames", arg0)
	return *ret0, err
}

// ListNames is a free data retrieval call binding the contract method 0x66272623.
//
// Solidity: function listNames( uint256) constant returns(bytes32)
func (_TokenLogic *TokenLogicSession) ListNames(arg0 *big.Int) ([32]byte, error) {
	return _TokenLogic.Contract.ListNames(&_TokenLogic.CallOpts, arg0)
}

// ListNames is a free data retrieval call binding the contract method 0x66272623.
//
// Solidity: function listNames( uint256) constant returns(bytes32)
func (_TokenLogic *TokenLogicCallerSession) ListNames(arg0 *big.Int) ([32]byte, error) {
	return _TokenLogic.Contract.ListNames(&_TokenLogic.CallOpts, arg0)
}

// ListNamesLen is a free data retrieval call binding the contract method 0x1f99c458.
//
// Solidity: function listNamesLen() constant returns(uint256)
func (_TokenLogic *TokenLogicCaller) ListNamesLen(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenLogic.contract.Call(opts, out, "listNamesLen")
	return *ret0, err
}

// ListNamesLen is a free data retrieval call binding the contract method 0x1f99c458.
//
// Solidity: function listNamesLen() constant returns(uint256)
func (_TokenLogic *TokenLogicSession) ListNamesLen() (*big.Int, error) {
	return _TokenLogic.Contract.ListNamesLen(&_TokenLogic.CallOpts)
}

// ListNamesLen is a free data retrieval call binding the contract method 0x1f99c458.
//
// Solidity: function listNamesLen() constant returns(uint256)
func (_TokenLogic *TokenLogicCallerSession) ListNamesLen() (*big.Int, error) {
	return _TokenLogic.Contract.ListNamesLen(&_TokenLogic.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenLogic *TokenLogicCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenLogic.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenLogic *TokenLogicSession) Owner() (common.Address, error) {
	return _TokenLogic.Contract.Owner(&_TokenLogic.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenLogic *TokenLogicCallerSession) Owner() (common.Address, error) {
	return _TokenLogic.Contract.Owner(&_TokenLogic.CallOpts)
}

// Roles is a free data retrieval call binding the contract method 0x392f5f64.
//
// Solidity: function roles() constant returns(address)
func (_TokenLogic *TokenLogicCaller) Roles(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenLogic.contract.Call(opts, out, "roles")
	return *ret0, err
}

// Roles is a free data retrieval call binding the contract method 0x392f5f64.
//
// Solidity: function roles() constant returns(address)
func (_TokenLogic *TokenLogicSession) Roles() (common.Address, error) {
	return _TokenLogic.Contract.Roles(&_TokenLogic.CallOpts)
}

// Roles is a free data retrieval call binding the contract method 0x392f5f64.
//
// Solidity: function roles() constant returns(address)
func (_TokenLogic *TokenLogicCallerSession) Roles() (common.Address, error) {
	return _TokenLogic.Contract.Roles(&_TokenLogic.CallOpts)
}

// SenderHasRole is a free data retrieval call binding the contract method 0x1ca03b8e.
//
// Solidity: function senderHasRole(roleName string) constant returns(bool)
func (_TokenLogic *TokenLogicCaller) SenderHasRole(opts *bind.CallOpts, roleName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenLogic.contract.Call(opts, out, "senderHasRole", roleName)
	return *ret0, err
}

// SenderHasRole is a free data retrieval call binding the contract method 0x1ca03b8e.
//
// Solidity: function senderHasRole(roleName string) constant returns(bool)
func (_TokenLogic *TokenLogicSession) SenderHasRole(roleName string) (bool, error) {
	return _TokenLogic.Contract.SenderHasRole(&_TokenLogic.CallOpts, roleName)
}

// SenderHasRole is a free data retrieval call binding the contract method 0x1ca03b8e.
//
// Solidity: function senderHasRole(roleName string) constant returns(bool)
func (_TokenLogic *TokenLogicCallerSession) SenderHasRole(roleName string) (bool, error) {
	return _TokenLogic.Contract.SenderHasRole(&_TokenLogic.CallOpts, roleName)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_TokenLogic *TokenLogicCaller) Stopped(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenLogic.contract.Call(opts, out, "stopped")
	return *ret0, err
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_TokenLogic *TokenLogicSession) Stopped() (bool, error) {
	return _TokenLogic.Contract.Stopped(&_TokenLogic.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_TokenLogic *TokenLogicCallerSession) Stopped() (bool, error) {
	return _TokenLogic.Contract.Stopped(&_TokenLogic.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_TokenLogic *TokenLogicCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenLogic.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_TokenLogic *TokenLogicSession) Token() (common.Address, error) {
	return _TokenLogic.Contract.Token(&_TokenLogic.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_TokenLogic *TokenLogicCallerSession) Token() (common.Address, error) {
	return _TokenLogic.Contract.Token(&_TokenLogic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TokenLogic *TokenLogicCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenLogic.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TokenLogic *TokenLogicSession) TotalSupply() (*big.Int, error) {
	return _TokenLogic.Contract.TotalSupply(&_TokenLogic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_TokenLogic *TokenLogicCallerSession) TotalSupply() (*big.Int, error) {
	return _TokenLogic.Contract.TotalSupply(&_TokenLogic.CallOpts)
}

// WhiteLists is a free data retrieval call binding the contract method 0x4b4d0d45.
//
// Solidity: function whiteLists( address,  bytes32) constant returns(bool)
func (_TokenLogic *TokenLogicCaller) WhiteLists(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _TokenLogic.contract.Call(opts, out, "whiteLists", arg0, arg1)
	return *ret0, err
}

// WhiteLists is a free data retrieval call binding the contract method 0x4b4d0d45.
//
// Solidity: function whiteLists( address,  bytes32) constant returns(bool)
func (_TokenLogic *TokenLogicSession) WhiteLists(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _TokenLogic.Contract.WhiteLists(&_TokenLogic.CallOpts, arg0, arg1)
}

// WhiteLists is a free data retrieval call binding the contract method 0x4b4d0d45.
//
// Solidity: function whiteLists( address,  bytes32) constant returns(bool)
func (_TokenLogic *TokenLogicCallerSession) WhiteLists(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _TokenLogic.Contract.WhiteLists(&_TokenLogic.CallOpts, arg0, arg1)
}

// AddToWhiteList is a paid mutator transaction binding the contract method 0x693382a9.
//
// Solidity: function addToWhiteList(listName bytes32, guy address) returns()
func (_TokenLogic *TokenLogicTransactor) AddToWhiteList(opts *bind.TransactOpts, listName [32]byte, guy common.Address) (*types.Transaction, error) {
	return _TokenLogic.contract.Transact(opts, "addToWhiteList", listName, guy)
}

// AddToWhiteList is a paid mutator transaction binding the contract method 0x693382a9.
//
// Solidity: function addToWhiteList(listName bytes32, guy address) returns()
func (_TokenLogic *TokenLogicSession) AddToWhiteList(listName [32]byte, guy common.Address) (*types.Transaction, error) {
	return _TokenLogic.Contract.AddToWhiteList(&_TokenLogic.TransactOpts, listName, guy)
}

// AddToWhiteList is a paid mutator transaction binding the contract method 0x693382a9.
//
// Solidity: function addToWhiteList(listName bytes32, guy address) returns()
func (_TokenLogic *TokenLogicTransactorSession) AddToWhiteList(listName [32]byte, guy common.Address) (*types.Transaction, error) {
	return _TokenLogic.Contract.AddToWhiteList(&_TokenLogic.TransactOpts, listName, guy)
}

// AddWhiteList is a paid mutator transaction binding the contract method 0x3bd04d69.
//
// Solidity: function addWhiteList(listName bytes32) returns()
func (_TokenLogic *TokenLogicTransactor) AddWhiteList(opts *bind.TransactOpts, listName [32]byte) (*types.Transaction, error) {
	return _TokenLogic.contract.Transact(opts, "addWhiteList", listName)
}

// AddWhiteList is a paid mutator transaction binding the contract method 0x3bd04d69.
//
// Solidity: function addWhiteList(listName bytes32) returns()
func (_TokenLogic *TokenLogicSession) AddWhiteList(listName [32]byte) (*types.Transaction, error) {
	return _TokenLogic.Contract.AddWhiteList(&_TokenLogic.TransactOpts, listName)
}

// AddWhiteList is a paid mutator transaction binding the contract method 0x3bd04d69.
//
// Solidity: function addWhiteList(listName bytes32) returns()
func (_TokenLogic *TokenLogicTransactorSession) AddWhiteList(listName [32]byte) (*types.Transaction, error) {
	return _TokenLogic.Contract.AddWhiteList(&_TokenLogic.TransactOpts, listName)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(src address, dst address, wad uint256) returns(bool)
func (_TokenLogic *TokenLogicTransactor) Approve(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _TokenLogic.contract.Transact(opts, "approve", src, dst, wad)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(src address, dst address, wad uint256) returns(bool)
func (_TokenLogic *TokenLogicSession) Approve(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _TokenLogic.Contract.Approve(&_TokenLogic.TransactOpts, src, dst, wad)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(src address, dst address, wad uint256) returns(bool)
func (_TokenLogic *TokenLogicTransactorSession) Approve(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _TokenLogic.Contract.Approve(&_TokenLogic.TransactOpts, src, dst, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x7261e469.
//
// Solidity: function burn(src address, wad uint128) returns()
func (_TokenLogic *TokenLogicTransactor) Burn(opts *bind.TransactOpts, src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _TokenLogic.contract.Transact(opts, "burn", src, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x7261e469.
//
// Solidity: function burn(src address, wad uint128) returns()
func (_TokenLogic *TokenLogicSession) Burn(src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _TokenLogic.Contract.Burn(&_TokenLogic.TransactOpts, src, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x7261e469.
//
// Solidity: function burn(src address, wad uint128) returns()
func (_TokenLogic *TokenLogicTransactorSession) Burn(src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _TokenLogic.Contract.Burn(&_TokenLogic.TransactOpts, src, wad)
}

// MintFor is a paid mutator transaction binding the contract method 0x27cfc219.
//
// Solidity: function mintFor(dst address, wad uint128) returns()
func (_TokenLogic *TokenLogicTransactor) MintFor(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _TokenLogic.contract.Transact(opts, "mintFor", dst, wad)
}

// MintFor is a paid mutator transaction binding the contract method 0x27cfc219.
//
// Solidity: function mintFor(dst address, wad uint128) returns()
func (_TokenLogic *TokenLogicSession) MintFor(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _TokenLogic.Contract.MintFor(&_TokenLogic.TransactOpts, dst, wad)
}

// MintFor is a paid mutator transaction binding the contract method 0x27cfc219.
//
// Solidity: function mintFor(dst address, wad uint128) returns()
func (_TokenLogic *TokenLogicTransactorSession) MintFor(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _TokenLogic.Contract.MintFor(&_TokenLogic.TransactOpts, dst, wad)
}

// RemoveFromWhiteList is a paid mutator transaction binding the contract method 0x3f32aa70.
//
// Solidity: function removeFromWhiteList(listName bytes32, guy address) returns()
func (_TokenLogic *TokenLogicTransactor) RemoveFromWhiteList(opts *bind.TransactOpts, listName [32]byte, guy common.Address) (*types.Transaction, error) {
	return _TokenLogic.contract.Transact(opts, "removeFromWhiteList", listName, guy)
}

// RemoveFromWhiteList is a paid mutator transaction binding the contract method 0x3f32aa70.
//
// Solidity: function removeFromWhiteList(listName bytes32, guy address) returns()
func (_TokenLogic *TokenLogicSession) RemoveFromWhiteList(listName [32]byte, guy common.Address) (*types.Transaction, error) {
	return _TokenLogic.Contract.RemoveFromWhiteList(&_TokenLogic.TransactOpts, listName, guy)
}

// RemoveFromWhiteList is a paid mutator transaction binding the contract method 0x3f32aa70.
//
// Solidity: function removeFromWhiteList(listName bytes32, guy address) returns()
func (_TokenLogic *TokenLogicTransactorSession) RemoveFromWhiteList(listName [32]byte, guy common.Address) (*types.Transaction, error) {
	return _TokenLogic.Contract.RemoveFromWhiteList(&_TokenLogic.TransactOpts, listName, guy)
}

// Restart is a paid mutator transaction binding the contract method 0x1ef3755d.
//
// Solidity: function restart() returns()
func (_TokenLogic *TokenLogicTransactor) Restart(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenLogic.contract.Transact(opts, "restart")
}

// Restart is a paid mutator transaction binding the contract method 0x1ef3755d.
//
// Solidity: function restart() returns()
func (_TokenLogic *TokenLogicSession) Restart() (*types.Transaction, error) {
	return _TokenLogic.Contract.Restart(&_TokenLogic.TransactOpts)
}

// Restart is a paid mutator transaction binding the contract method 0x1ef3755d.
//
// Solidity: function restart() returns()
func (_TokenLogic *TokenLogicTransactorSession) Restart() (*types.Transaction, error) {
	return _TokenLogic.Contract.Restart(&_TokenLogic.TransactOpts)
}

// SetFreeTransfer is a paid mutator transaction binding the contract method 0x2abfaf1f.
//
// Solidity: function setFreeTransfer(isFree bool) returns()
func (_TokenLogic *TokenLogicTransactor) SetFreeTransfer(opts *bind.TransactOpts, isFree bool) (*types.Transaction, error) {
	return _TokenLogic.contract.Transact(opts, "setFreeTransfer", isFree)
}

// SetFreeTransfer is a paid mutator transaction binding the contract method 0x2abfaf1f.
//
// Solidity: function setFreeTransfer(isFree bool) returns()
func (_TokenLogic *TokenLogicSession) SetFreeTransfer(isFree bool) (*types.Transaction, error) {
	return _TokenLogic.Contract.SetFreeTransfer(&_TokenLogic.TransactOpts, isFree)
}

// SetFreeTransfer is a paid mutator transaction binding the contract method 0x2abfaf1f.
//
// Solidity: function setFreeTransfer(isFree bool) returns()
func (_TokenLogic *TokenLogicTransactorSession) SetFreeTransfer(isFree bool) (*types.Transaction, error) {
	return _TokenLogic.Contract.SetFreeTransfer(&_TokenLogic.TransactOpts, isFree)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_TokenLogic *TokenLogicTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _TokenLogic.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_TokenLogic *TokenLogicSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _TokenLogic.Contract.SetOwner(&_TokenLogic.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_TokenLogic *TokenLogicTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _TokenLogic.Contract.SetOwner(&_TokenLogic.TransactOpts, owner_)
}

// SetRolesContract is a paid mutator transaction binding the contract method 0xafa202ac.
//
// Solidity: function setRolesContract(roles_ address) returns()
func (_TokenLogic *TokenLogicTransactor) SetRolesContract(opts *bind.TransactOpts, roles_ common.Address) (*types.Transaction, error) {
	return _TokenLogic.contract.Transact(opts, "setRolesContract", roles_)
}

// SetRolesContract is a paid mutator transaction binding the contract method 0xafa202ac.
//
// Solidity: function setRolesContract(roles_ address) returns()
func (_TokenLogic *TokenLogicSession) SetRolesContract(roles_ common.Address) (*types.Transaction, error) {
	return _TokenLogic.Contract.SetRolesContract(&_TokenLogic.TransactOpts, roles_)
}

// SetRolesContract is a paid mutator transaction binding the contract method 0xafa202ac.
//
// Solidity: function setRolesContract(roles_ address) returns()
func (_TokenLogic *TokenLogicTransactorSession) SetRolesContract(roles_ common.Address) (*types.Transaction, error) {
	return _TokenLogic.Contract.SetRolesContract(&_TokenLogic.TransactOpts, roles_)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(token_ address) returns()
func (_TokenLogic *TokenLogicTransactor) SetToken(opts *bind.TransactOpts, token_ common.Address) (*types.Transaction, error) {
	return _TokenLogic.contract.Transact(opts, "setToken", token_)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(token_ address) returns()
func (_TokenLogic *TokenLogicSession) SetToken(token_ common.Address) (*types.Transaction, error) {
	return _TokenLogic.Contract.SetToken(&_TokenLogic.TransactOpts, token_)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(token_ address) returns()
func (_TokenLogic *TokenLogicTransactorSession) SetToken(token_ common.Address) (*types.Transaction, error) {
	return _TokenLogic.Contract.SetToken(&_TokenLogic.TransactOpts, token_)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_TokenLogic *TokenLogicTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenLogic.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_TokenLogic *TokenLogicSession) Stop() (*types.Transaction, error) {
	return _TokenLogic.Contract.Stop(&_TokenLogic.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_TokenLogic *TokenLogicTransactorSession) Stop() (*types.Transaction, error) {
	return _TokenLogic.Contract.Stop(&_TokenLogic.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(src address, dst address, wad uint256) returns(bool)
func (_TokenLogic *TokenLogicTransactor) Transfer(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _TokenLogic.contract.Transact(opts, "transfer", src, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(src address, dst address, wad uint256) returns(bool)
func (_TokenLogic *TokenLogicSession) Transfer(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _TokenLogic.Contract.Transfer(&_TokenLogic.TransactOpts, src, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(src address, dst address, wad uint256) returns(bool)
func (_TokenLogic *TokenLogicTransactorSession) Transfer(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _TokenLogic.Contract.Transfer(&_TokenLogic.TransactOpts, src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_TokenLogic *TokenLogicTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _TokenLogic.contract.Transact(opts, "transferFrom", src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_TokenLogic *TokenLogicSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _TokenLogic.Contract.TransferFrom(&_TokenLogic.TransactOpts, src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_TokenLogic *TokenLogicTransactorSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _TokenLogic.Contract.TransferFrom(&_TokenLogic.TransactOpts, src, dst, wad)
}

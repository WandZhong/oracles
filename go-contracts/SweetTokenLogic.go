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

// SweetTokenLogicABI is the input ABI used to generate the binding from.
const SweetTokenLogicABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token_\",\"type\":\"address\"}],\"name\":\"setToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"senderHasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"restart\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"listNamesLen\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"freeTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint128\"}],\"name\":\"mintFor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"isFree\",\"type\":\"bool\"}],\"name\":\"setFreeTransfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"roles\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"listName\",\"type\":\"bytes32\"}],\"name\":\"addWhiteList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"listName\",\"type\":\"bytes32\"},{\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"removeFromWhiteList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"whiteLists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"listNames\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"listName\",\"type\":\"bytes32\"},{\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"addToWhiteList\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint128\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"listName\",\"type\":\"bytes32\"}],\"name\":\"listExists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"roles_\",\"type\":\"address\"}],\"name\":\"setRolesContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"hasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"token_\",\"type\":\"address\"},{\"name\":\"data_\",\"type\":\"address\"},{\"name\":\"supply_\",\"type\":\"uint256\"},{\"name\":\"rolesContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"}]"

// SweetTokenLogic is an auto generated Go binding around an Ethereum contract.
type SweetTokenLogic struct {
	SweetTokenLogicCaller     // Read-only binding to the contract
	SweetTokenLogicTransactor // Write-only binding to the contract
}

// SweetTokenLogicCaller is an auto generated read-only Go binding around an Ethereum contract.
type SweetTokenLogicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SweetTokenLogicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SweetTokenLogicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SweetTokenLogicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SweetTokenLogicSession struct {
	Contract     *SweetTokenLogic  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SweetTokenLogicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SweetTokenLogicCallerSession struct {
	Contract *SweetTokenLogicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// SweetTokenLogicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SweetTokenLogicTransactorSession struct {
	Contract     *SweetTokenLogicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// SweetTokenLogicRaw is an auto generated low-level Go binding around an Ethereum contract.
type SweetTokenLogicRaw struct {
	Contract *SweetTokenLogic // Generic contract binding to access the raw methods on
}

// SweetTokenLogicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SweetTokenLogicCallerRaw struct {
	Contract *SweetTokenLogicCaller // Generic read-only contract binding to access the raw methods on
}

// SweetTokenLogicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SweetTokenLogicTransactorRaw struct {
	Contract *SweetTokenLogicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSweetTokenLogic creates a new instance of SweetTokenLogic, bound to a specific deployed contract.
func NewSweetTokenLogic(address common.Address, backend bind.ContractBackend) (*SweetTokenLogic, error) {
	contract, err := bindSweetTokenLogic(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SweetTokenLogic{SweetTokenLogicCaller: SweetTokenLogicCaller{contract: contract}, SweetTokenLogicTransactor: SweetTokenLogicTransactor{contract: contract}}, nil
}

// NewSweetTokenLogicCaller creates a new read-only instance of SweetTokenLogic, bound to a specific deployed contract.
func NewSweetTokenLogicCaller(address common.Address, caller bind.ContractCaller) (*SweetTokenLogicCaller, error) {
	contract, err := bindSweetTokenLogic(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &SweetTokenLogicCaller{contract: contract}, nil
}

// NewSweetTokenLogicTransactor creates a new write-only instance of SweetTokenLogic, bound to a specific deployed contract.
func NewSweetTokenLogicTransactor(address common.Address, transactor bind.ContractTransactor) (*SweetTokenLogicTransactor, error) {
	contract, err := bindSweetTokenLogic(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &SweetTokenLogicTransactor{contract: contract}, nil
}

// bindSweetTokenLogic binds a generic wrapper to an already deployed contract.
func bindSweetTokenLogic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SweetTokenLogicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SweetTokenLogic *SweetTokenLogicRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SweetTokenLogic.Contract.SweetTokenLogicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SweetTokenLogic *SweetTokenLogicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.SweetTokenLogicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SweetTokenLogic *SweetTokenLogicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.SweetTokenLogicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SweetTokenLogic *SweetTokenLogicCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SweetTokenLogic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SweetTokenLogic *SweetTokenLogicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SweetTokenLogic *SweetTokenLogicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(src address, spender address) constant returns(uint256)
func (_SweetTokenLogic *SweetTokenLogicCaller) Allowance(opts *bind.CallOpts, src common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SweetTokenLogic.contract.Call(opts, out, "allowance", src, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(src address, spender address) constant returns(uint256)
func (_SweetTokenLogic *SweetTokenLogicSession) Allowance(src common.Address, spender common.Address) (*big.Int, error) {
	return _SweetTokenLogic.Contract.Allowance(&_SweetTokenLogic.CallOpts, src, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(src address, spender address) constant returns(uint256)
func (_SweetTokenLogic *SweetTokenLogicCallerSession) Allowance(src common.Address, spender common.Address) (*big.Int, error) {
	return _SweetTokenLogic.Contract.Allowance(&_SweetTokenLogic.CallOpts, src, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(src address) constant returns(uint256)
func (_SweetTokenLogic *SweetTokenLogicCaller) BalanceOf(opts *bind.CallOpts, src common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SweetTokenLogic.contract.Call(opts, out, "balanceOf", src)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(src address) constant returns(uint256)
func (_SweetTokenLogic *SweetTokenLogicSession) BalanceOf(src common.Address) (*big.Int, error) {
	return _SweetTokenLogic.Contract.BalanceOf(&_SweetTokenLogic.CallOpts, src)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(src address) constant returns(uint256)
func (_SweetTokenLogic *SweetTokenLogicCallerSession) BalanceOf(src common.Address) (*big.Int, error) {
	return _SweetTokenLogic.Contract.BalanceOf(&_SweetTokenLogic.CallOpts, src)
}

// ContractHash is a free data retrieval call binding the contract method 0x904c6094.
//
// Solidity: function contractHash() constant returns(bytes32)
func (_SweetTokenLogic *SweetTokenLogicCaller) ContractHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _SweetTokenLogic.contract.Call(opts, out, "contractHash")
	return *ret0, err
}

// ContractHash is a free data retrieval call binding the contract method 0x904c6094.
//
// Solidity: function contractHash() constant returns(bytes32)
func (_SweetTokenLogic *SweetTokenLogicSession) ContractHash() ([32]byte, error) {
	return _SweetTokenLogic.Contract.ContractHash(&_SweetTokenLogic.CallOpts)
}

// ContractHash is a free data retrieval call binding the contract method 0x904c6094.
//
// Solidity: function contractHash() constant returns(bytes32)
func (_SweetTokenLogic *SweetTokenLogicCallerSession) ContractHash() ([32]byte, error) {
	return _SweetTokenLogic.Contract.ContractHash(&_SweetTokenLogic.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(address)
func (_SweetTokenLogic *SweetTokenLogicCaller) Data(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SweetTokenLogic.contract.Call(opts, out, "data")
	return *ret0, err
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(address)
func (_SweetTokenLogic *SweetTokenLogicSession) Data() (common.Address, error) {
	return _SweetTokenLogic.Contract.Data(&_SweetTokenLogic.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(address)
func (_SweetTokenLogic *SweetTokenLogicCallerSession) Data() (common.Address, error) {
	return _SweetTokenLogic.Contract.Data(&_SweetTokenLogic.CallOpts)
}

// FreeTransfer is a free data retrieval call binding the contract method 0x211ed6c1.
//
// Solidity: function freeTransfer() constant returns(bool)
func (_SweetTokenLogic *SweetTokenLogicCaller) FreeTransfer(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SweetTokenLogic.contract.Call(opts, out, "freeTransfer")
	return *ret0, err
}

// FreeTransfer is a free data retrieval call binding the contract method 0x211ed6c1.
//
// Solidity: function freeTransfer() constant returns(bool)
func (_SweetTokenLogic *SweetTokenLogicSession) FreeTransfer() (bool, error) {
	return _SweetTokenLogic.Contract.FreeTransfer(&_SweetTokenLogic.CallOpts)
}

// FreeTransfer is a free data retrieval call binding the contract method 0x211ed6c1.
//
// Solidity: function freeTransfer() constant returns(bool)
func (_SweetTokenLogic *SweetTokenLogicCallerSession) FreeTransfer() (bool, error) {
	return _SweetTokenLogic.Contract.FreeTransfer(&_SweetTokenLogic.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0xe3c33a9b.
//
// Solidity: function hasRole(roleName string) constant returns(bool)
func (_SweetTokenLogic *SweetTokenLogicCaller) HasRole(opts *bind.CallOpts, roleName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SweetTokenLogic.contract.Call(opts, out, "hasRole", roleName)
	return *ret0, err
}

// HasRole is a free data retrieval call binding the contract method 0xe3c33a9b.
//
// Solidity: function hasRole(roleName string) constant returns(bool)
func (_SweetTokenLogic *SweetTokenLogicSession) HasRole(roleName string) (bool, error) {
	return _SweetTokenLogic.Contract.HasRole(&_SweetTokenLogic.CallOpts, roleName)
}

// HasRole is a free data retrieval call binding the contract method 0xe3c33a9b.
//
// Solidity: function hasRole(roleName string) constant returns(bool)
func (_SweetTokenLogic *SweetTokenLogicCallerSession) HasRole(roleName string) (bool, error) {
	return _SweetTokenLogic.Contract.HasRole(&_SweetTokenLogic.CallOpts, roleName)
}

// ListExists is a free data retrieval call binding the contract method 0x9495ad6c.
//
// Solidity: function listExists(listName bytes32) constant returns(bool)
func (_SweetTokenLogic *SweetTokenLogicCaller) ListExists(opts *bind.CallOpts, listName [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SweetTokenLogic.contract.Call(opts, out, "listExists", listName)
	return *ret0, err
}

// ListExists is a free data retrieval call binding the contract method 0x9495ad6c.
//
// Solidity: function listExists(listName bytes32) constant returns(bool)
func (_SweetTokenLogic *SweetTokenLogicSession) ListExists(listName [32]byte) (bool, error) {
	return _SweetTokenLogic.Contract.ListExists(&_SweetTokenLogic.CallOpts, listName)
}

// ListExists is a free data retrieval call binding the contract method 0x9495ad6c.
//
// Solidity: function listExists(listName bytes32) constant returns(bool)
func (_SweetTokenLogic *SweetTokenLogicCallerSession) ListExists(listName [32]byte) (bool, error) {
	return _SweetTokenLogic.Contract.ListExists(&_SweetTokenLogic.CallOpts, listName)
}

// ListNames is a free data retrieval call binding the contract method 0x66272623.
//
// Solidity: function listNames( uint256) constant returns(bytes32)
func (_SweetTokenLogic *SweetTokenLogicCaller) ListNames(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _SweetTokenLogic.contract.Call(opts, out, "listNames", arg0)
	return *ret0, err
}

// ListNames is a free data retrieval call binding the contract method 0x66272623.
//
// Solidity: function listNames( uint256) constant returns(bytes32)
func (_SweetTokenLogic *SweetTokenLogicSession) ListNames(arg0 *big.Int) ([32]byte, error) {
	return _SweetTokenLogic.Contract.ListNames(&_SweetTokenLogic.CallOpts, arg0)
}

// ListNames is a free data retrieval call binding the contract method 0x66272623.
//
// Solidity: function listNames( uint256) constant returns(bytes32)
func (_SweetTokenLogic *SweetTokenLogicCallerSession) ListNames(arg0 *big.Int) ([32]byte, error) {
	return _SweetTokenLogic.Contract.ListNames(&_SweetTokenLogic.CallOpts, arg0)
}

// ListNamesLen is a free data retrieval call binding the contract method 0x1f99c458.
//
// Solidity: function listNamesLen() constant returns(uint256)
func (_SweetTokenLogic *SweetTokenLogicCaller) ListNamesLen(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SweetTokenLogic.contract.Call(opts, out, "listNamesLen")
	return *ret0, err
}

// ListNamesLen is a free data retrieval call binding the contract method 0x1f99c458.
//
// Solidity: function listNamesLen() constant returns(uint256)
func (_SweetTokenLogic *SweetTokenLogicSession) ListNamesLen() (*big.Int, error) {
	return _SweetTokenLogic.Contract.ListNamesLen(&_SweetTokenLogic.CallOpts)
}

// ListNamesLen is a free data retrieval call binding the contract method 0x1f99c458.
//
// Solidity: function listNamesLen() constant returns(uint256)
func (_SweetTokenLogic *SweetTokenLogicCallerSession) ListNamesLen() (*big.Int, error) {
	return _SweetTokenLogic.Contract.ListNamesLen(&_SweetTokenLogic.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SweetTokenLogic *SweetTokenLogicCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SweetTokenLogic.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SweetTokenLogic *SweetTokenLogicSession) Owner() (common.Address, error) {
	return _SweetTokenLogic.Contract.Owner(&_SweetTokenLogic.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SweetTokenLogic *SweetTokenLogicCallerSession) Owner() (common.Address, error) {
	return _SweetTokenLogic.Contract.Owner(&_SweetTokenLogic.CallOpts)
}

// Roles is a free data retrieval call binding the contract method 0x392f5f64.
//
// Solidity: function roles() constant returns(address)
func (_SweetTokenLogic *SweetTokenLogicCaller) Roles(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SweetTokenLogic.contract.Call(opts, out, "roles")
	return *ret0, err
}

// Roles is a free data retrieval call binding the contract method 0x392f5f64.
//
// Solidity: function roles() constant returns(address)
func (_SweetTokenLogic *SweetTokenLogicSession) Roles() (common.Address, error) {
	return _SweetTokenLogic.Contract.Roles(&_SweetTokenLogic.CallOpts)
}

// Roles is a free data retrieval call binding the contract method 0x392f5f64.
//
// Solidity: function roles() constant returns(address)
func (_SweetTokenLogic *SweetTokenLogicCallerSession) Roles() (common.Address, error) {
	return _SweetTokenLogic.Contract.Roles(&_SweetTokenLogic.CallOpts)
}

// SenderHasRole is a free data retrieval call binding the contract method 0x1ca03b8e.
//
// Solidity: function senderHasRole(roleName string) constant returns(bool)
func (_SweetTokenLogic *SweetTokenLogicCaller) SenderHasRole(opts *bind.CallOpts, roleName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SweetTokenLogic.contract.Call(opts, out, "senderHasRole", roleName)
	return *ret0, err
}

// SenderHasRole is a free data retrieval call binding the contract method 0x1ca03b8e.
//
// Solidity: function senderHasRole(roleName string) constant returns(bool)
func (_SweetTokenLogic *SweetTokenLogicSession) SenderHasRole(roleName string) (bool, error) {
	return _SweetTokenLogic.Contract.SenderHasRole(&_SweetTokenLogic.CallOpts, roleName)
}

// SenderHasRole is a free data retrieval call binding the contract method 0x1ca03b8e.
//
// Solidity: function senderHasRole(roleName string) constant returns(bool)
func (_SweetTokenLogic *SweetTokenLogicCallerSession) SenderHasRole(roleName string) (bool, error) {
	return _SweetTokenLogic.Contract.SenderHasRole(&_SweetTokenLogic.CallOpts, roleName)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_SweetTokenLogic *SweetTokenLogicCaller) Stopped(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SweetTokenLogic.contract.Call(opts, out, "stopped")
	return *ret0, err
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_SweetTokenLogic *SweetTokenLogicSession) Stopped() (bool, error) {
	return _SweetTokenLogic.Contract.Stopped(&_SweetTokenLogic.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_SweetTokenLogic *SweetTokenLogicCallerSession) Stopped() (bool, error) {
	return _SweetTokenLogic.Contract.Stopped(&_SweetTokenLogic.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_SweetTokenLogic *SweetTokenLogicCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SweetTokenLogic.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_SweetTokenLogic *SweetTokenLogicSession) Token() (common.Address, error) {
	return _SweetTokenLogic.Contract.Token(&_SweetTokenLogic.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_SweetTokenLogic *SweetTokenLogicCallerSession) Token() (common.Address, error) {
	return _SweetTokenLogic.Contract.Token(&_SweetTokenLogic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SweetTokenLogic *SweetTokenLogicCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SweetTokenLogic.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SweetTokenLogic *SweetTokenLogicSession) TotalSupply() (*big.Int, error) {
	return _SweetTokenLogic.Contract.TotalSupply(&_SweetTokenLogic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SweetTokenLogic *SweetTokenLogicCallerSession) TotalSupply() (*big.Int, error) {
	return _SweetTokenLogic.Contract.TotalSupply(&_SweetTokenLogic.CallOpts)
}

// WhiteLists is a free data retrieval call binding the contract method 0x4b4d0d45.
//
// Solidity: function whiteLists( address,  bytes32) constant returns(bool)
func (_SweetTokenLogic *SweetTokenLogicCaller) WhiteLists(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SweetTokenLogic.contract.Call(opts, out, "whiteLists", arg0, arg1)
	return *ret0, err
}

// WhiteLists is a free data retrieval call binding the contract method 0x4b4d0d45.
//
// Solidity: function whiteLists( address,  bytes32) constant returns(bool)
func (_SweetTokenLogic *SweetTokenLogicSession) WhiteLists(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _SweetTokenLogic.Contract.WhiteLists(&_SweetTokenLogic.CallOpts, arg0, arg1)
}

// WhiteLists is a free data retrieval call binding the contract method 0x4b4d0d45.
//
// Solidity: function whiteLists( address,  bytes32) constant returns(bool)
func (_SweetTokenLogic *SweetTokenLogicCallerSession) WhiteLists(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _SweetTokenLogic.Contract.WhiteLists(&_SweetTokenLogic.CallOpts, arg0, arg1)
}

// AddToWhiteList is a paid mutator transaction binding the contract method 0x693382a9.
//
// Solidity: function addToWhiteList(listName bytes32, guy address) returns()
func (_SweetTokenLogic *SweetTokenLogicTransactor) AddToWhiteList(opts *bind.TransactOpts, listName [32]byte, guy common.Address) (*types.Transaction, error) {
	return _SweetTokenLogic.contract.Transact(opts, "addToWhiteList", listName, guy)
}

// AddToWhiteList is a paid mutator transaction binding the contract method 0x693382a9.
//
// Solidity: function addToWhiteList(listName bytes32, guy address) returns()
func (_SweetTokenLogic *SweetTokenLogicSession) AddToWhiteList(listName [32]byte, guy common.Address) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.AddToWhiteList(&_SweetTokenLogic.TransactOpts, listName, guy)
}

// AddToWhiteList is a paid mutator transaction binding the contract method 0x693382a9.
//
// Solidity: function addToWhiteList(listName bytes32, guy address) returns()
func (_SweetTokenLogic *SweetTokenLogicTransactorSession) AddToWhiteList(listName [32]byte, guy common.Address) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.AddToWhiteList(&_SweetTokenLogic.TransactOpts, listName, guy)
}

// AddWhiteList is a paid mutator transaction binding the contract method 0x3bd04d69.
//
// Solidity: function addWhiteList(listName bytes32) returns()
func (_SweetTokenLogic *SweetTokenLogicTransactor) AddWhiteList(opts *bind.TransactOpts, listName [32]byte) (*types.Transaction, error) {
	return _SweetTokenLogic.contract.Transact(opts, "addWhiteList", listName)
}

// AddWhiteList is a paid mutator transaction binding the contract method 0x3bd04d69.
//
// Solidity: function addWhiteList(listName bytes32) returns()
func (_SweetTokenLogic *SweetTokenLogicSession) AddWhiteList(listName [32]byte) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.AddWhiteList(&_SweetTokenLogic.TransactOpts, listName)
}

// AddWhiteList is a paid mutator transaction binding the contract method 0x3bd04d69.
//
// Solidity: function addWhiteList(listName bytes32) returns()
func (_SweetTokenLogic *SweetTokenLogicTransactorSession) AddWhiteList(listName [32]byte) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.AddWhiteList(&_SweetTokenLogic.TransactOpts, listName)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(src address, dst address, wad uint256) returns(bool)
func (_SweetTokenLogic *SweetTokenLogicTransactor) Approve(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetTokenLogic.contract.Transact(opts, "approve", src, dst, wad)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(src address, dst address, wad uint256) returns(bool)
func (_SweetTokenLogic *SweetTokenLogicSession) Approve(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.Approve(&_SweetTokenLogic.TransactOpts, src, dst, wad)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(src address, dst address, wad uint256) returns(bool)
func (_SweetTokenLogic *SweetTokenLogicTransactorSession) Approve(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.Approve(&_SweetTokenLogic.TransactOpts, src, dst, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x7261e469.
//
// Solidity: function burn( address,  uint128) returns()
func (_SweetTokenLogic *SweetTokenLogicTransactor) Burn(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _SweetTokenLogic.contract.Transact(opts, "burn", arg0, arg1)
}

// Burn is a paid mutator transaction binding the contract method 0x7261e469.
//
// Solidity: function burn( address,  uint128) returns()
func (_SweetTokenLogic *SweetTokenLogicSession) Burn(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.Burn(&_SweetTokenLogic.TransactOpts, arg0, arg1)
}

// Burn is a paid mutator transaction binding the contract method 0x7261e469.
//
// Solidity: function burn( address,  uint128) returns()
func (_SweetTokenLogic *SweetTokenLogicTransactorSession) Burn(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.Burn(&_SweetTokenLogic.TransactOpts, arg0, arg1)
}

// MintFor is a paid mutator transaction binding the contract method 0x27cfc219.
//
// Solidity: function mintFor( address,  uint128) returns()
func (_SweetTokenLogic *SweetTokenLogicTransactor) MintFor(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _SweetTokenLogic.contract.Transact(opts, "mintFor", arg0, arg1)
}

// MintFor is a paid mutator transaction binding the contract method 0x27cfc219.
//
// Solidity: function mintFor( address,  uint128) returns()
func (_SweetTokenLogic *SweetTokenLogicSession) MintFor(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.MintFor(&_SweetTokenLogic.TransactOpts, arg0, arg1)
}

// MintFor is a paid mutator transaction binding the contract method 0x27cfc219.
//
// Solidity: function mintFor( address,  uint128) returns()
func (_SweetTokenLogic *SweetTokenLogicTransactorSession) MintFor(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.MintFor(&_SweetTokenLogic.TransactOpts, arg0, arg1)
}

// RemoveFromWhiteList is a paid mutator transaction binding the contract method 0x3f32aa70.
//
// Solidity: function removeFromWhiteList(listName bytes32, guy address) returns()
func (_SweetTokenLogic *SweetTokenLogicTransactor) RemoveFromWhiteList(opts *bind.TransactOpts, listName [32]byte, guy common.Address) (*types.Transaction, error) {
	return _SweetTokenLogic.contract.Transact(opts, "removeFromWhiteList", listName, guy)
}

// RemoveFromWhiteList is a paid mutator transaction binding the contract method 0x3f32aa70.
//
// Solidity: function removeFromWhiteList(listName bytes32, guy address) returns()
func (_SweetTokenLogic *SweetTokenLogicSession) RemoveFromWhiteList(listName [32]byte, guy common.Address) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.RemoveFromWhiteList(&_SweetTokenLogic.TransactOpts, listName, guy)
}

// RemoveFromWhiteList is a paid mutator transaction binding the contract method 0x3f32aa70.
//
// Solidity: function removeFromWhiteList(listName bytes32, guy address) returns()
func (_SweetTokenLogic *SweetTokenLogicTransactorSession) RemoveFromWhiteList(listName [32]byte, guy common.Address) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.RemoveFromWhiteList(&_SweetTokenLogic.TransactOpts, listName, guy)
}

// Restart is a paid mutator transaction binding the contract method 0x1ef3755d.
//
// Solidity: function restart() returns()
func (_SweetTokenLogic *SweetTokenLogicTransactor) Restart(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SweetTokenLogic.contract.Transact(opts, "restart")
}

// Restart is a paid mutator transaction binding the contract method 0x1ef3755d.
//
// Solidity: function restart() returns()
func (_SweetTokenLogic *SweetTokenLogicSession) Restart() (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.Restart(&_SweetTokenLogic.TransactOpts)
}

// Restart is a paid mutator transaction binding the contract method 0x1ef3755d.
//
// Solidity: function restart() returns()
func (_SweetTokenLogic *SweetTokenLogicTransactorSession) Restart() (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.Restart(&_SweetTokenLogic.TransactOpts)
}

// SetFreeTransfer is a paid mutator transaction binding the contract method 0x2abfaf1f.
//
// Solidity: function setFreeTransfer(isFree bool) returns()
func (_SweetTokenLogic *SweetTokenLogicTransactor) SetFreeTransfer(opts *bind.TransactOpts, isFree bool) (*types.Transaction, error) {
	return _SweetTokenLogic.contract.Transact(opts, "setFreeTransfer", isFree)
}

// SetFreeTransfer is a paid mutator transaction binding the contract method 0x2abfaf1f.
//
// Solidity: function setFreeTransfer(isFree bool) returns()
func (_SweetTokenLogic *SweetTokenLogicSession) SetFreeTransfer(isFree bool) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.SetFreeTransfer(&_SweetTokenLogic.TransactOpts, isFree)
}

// SetFreeTransfer is a paid mutator transaction binding the contract method 0x2abfaf1f.
//
// Solidity: function setFreeTransfer(isFree bool) returns()
func (_SweetTokenLogic *SweetTokenLogicTransactorSession) SetFreeTransfer(isFree bool) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.SetFreeTransfer(&_SweetTokenLogic.TransactOpts, isFree)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_SweetTokenLogic *SweetTokenLogicTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _SweetTokenLogic.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_SweetTokenLogic *SweetTokenLogicSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.SetOwner(&_SweetTokenLogic.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_SweetTokenLogic *SweetTokenLogicTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.SetOwner(&_SweetTokenLogic.TransactOpts, owner_)
}

// SetRolesContract is a paid mutator transaction binding the contract method 0xafa202ac.
//
// Solidity: function setRolesContract(roles_ address) returns()
func (_SweetTokenLogic *SweetTokenLogicTransactor) SetRolesContract(opts *bind.TransactOpts, roles_ common.Address) (*types.Transaction, error) {
	return _SweetTokenLogic.contract.Transact(opts, "setRolesContract", roles_)
}

// SetRolesContract is a paid mutator transaction binding the contract method 0xafa202ac.
//
// Solidity: function setRolesContract(roles_ address) returns()
func (_SweetTokenLogic *SweetTokenLogicSession) SetRolesContract(roles_ common.Address) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.SetRolesContract(&_SweetTokenLogic.TransactOpts, roles_)
}

// SetRolesContract is a paid mutator transaction binding the contract method 0xafa202ac.
//
// Solidity: function setRolesContract(roles_ address) returns()
func (_SweetTokenLogic *SweetTokenLogicTransactorSession) SetRolesContract(roles_ common.Address) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.SetRolesContract(&_SweetTokenLogic.TransactOpts, roles_)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(token_ address) returns()
func (_SweetTokenLogic *SweetTokenLogicTransactor) SetToken(opts *bind.TransactOpts, token_ common.Address) (*types.Transaction, error) {
	return _SweetTokenLogic.contract.Transact(opts, "setToken", token_)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(token_ address) returns()
func (_SweetTokenLogic *SweetTokenLogicSession) SetToken(token_ common.Address) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.SetToken(&_SweetTokenLogic.TransactOpts, token_)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(token_ address) returns()
func (_SweetTokenLogic *SweetTokenLogicTransactorSession) SetToken(token_ common.Address) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.SetToken(&_SweetTokenLogic.TransactOpts, token_)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_SweetTokenLogic *SweetTokenLogicTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SweetTokenLogic.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_SweetTokenLogic *SweetTokenLogicSession) Stop() (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.Stop(&_SweetTokenLogic.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_SweetTokenLogic *SweetTokenLogicTransactorSession) Stop() (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.Stop(&_SweetTokenLogic.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(src address, dst address, wad uint256) returns(bool)
func (_SweetTokenLogic *SweetTokenLogicTransactor) Transfer(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetTokenLogic.contract.Transact(opts, "transfer", src, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(src address, dst address, wad uint256) returns(bool)
func (_SweetTokenLogic *SweetTokenLogicSession) Transfer(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.Transfer(&_SweetTokenLogic.TransactOpts, src, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(src address, dst address, wad uint256) returns(bool)
func (_SweetTokenLogic *SweetTokenLogicTransactorSession) Transfer(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.Transfer(&_SweetTokenLogic.TransactOpts, src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_SweetTokenLogic *SweetTokenLogicTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetTokenLogic.contract.Transact(opts, "transferFrom", src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_SweetTokenLogic *SweetTokenLogicSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.TransferFrom(&_SweetTokenLogic.TransactOpts, src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_SweetTokenLogic *SweetTokenLogicTransactorSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.TransferFrom(&_SweetTokenLogic.TransactOpts, src, dst, wad)
}

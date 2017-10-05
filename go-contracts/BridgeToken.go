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

// BridgeTokenABI is the input ABI used to generate the binding from.
const BridgeTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guy\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"brgAmount\",\"type\":\"uint128\"},{\"name\":\"vault\",\"type\":\"address\"},{\"name\":\"uouIndex\",\"type\":\"uint256\"}],\"name\":\"repayUou\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"senderHasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"restart\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"mintFor\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"roles\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"logic_\",\"type\":\"address\"}],\"name\":\"setLogic\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"stopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"pull\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"contractHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"burn\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"roles_\",\"type\":\"address\"}],\"name\":\"setRolesContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name_\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"logic\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"_allowance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"roleName\",\"type\":\"string\"},{\"name\":\"symbol_\",\"type\":\"bytes32\"},{\"name\":\"rolesContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true,\"name\":\"hasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}]},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"name_\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"symbol_\",\"type\":\"bytes32\"},{\"name\":\"rolesContract\",\"type\":\"address\"}],\"name\":\"LogBurn\",\"type\":\"constructor\",\"payable\":false,\"stateMutability\":\"nonpayable\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"src\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"LogBurn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"src\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"LogMint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// BridgeTokenBin is the compiled bytecode used for deploying new contracts.
const BridgeTokenBin = `undefined`

// DeployBridgeToken deploys a new Ethereum contract, binding an instance of BridgeToken to it.
func DeployBridgeToken(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ [32]byte, rolesContract common.Address) (common.Address, *types.Transaction, *BridgeToken, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BridgeTokenBin), backend, name_, symbol_, rolesContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeToken{BridgeTokenCaller: BridgeTokenCaller{contract: contract}, BridgeTokenTransactor: BridgeTokenTransactor{contract: contract}}, nil
}

// BridgeToken is an auto generated Go binding around an Ethereum contract.
type BridgeToken struct {
	BridgeTokenCaller     // Read-only binding to the contract
	BridgeTokenTransactor // Write-only binding to the contract
}

// BridgeTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeTokenSession struct {
	Contract     *BridgeToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeTokenCallerSession struct {
	Contract *BridgeTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BridgeTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTokenTransactorSession struct {
	Contract     *BridgeTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BridgeTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeTokenRaw struct {
	Contract *BridgeToken // Generic contract binding to access the raw methods on
}

// BridgeTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeTokenCallerRaw struct {
	Contract *BridgeTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTokenTransactorRaw struct {
	Contract *BridgeTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeToken creates a new instance of BridgeToken, bound to a specific deployed contract.
func NewBridgeToken(address common.Address, backend bind.ContractBackend) (*BridgeToken, error) {
	contract, err := bindBridgeToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeToken{BridgeTokenCaller: BridgeTokenCaller{contract: contract}, BridgeTokenTransactor: BridgeTokenTransactor{contract: contract}}, nil
}

// NewBridgeTokenCaller creates a new read-only instance of BridgeToken, bound to a specific deployed contract.
func NewBridgeTokenCaller(address common.Address, caller bind.ContractCaller) (*BridgeTokenCaller, error) {
	contract, err := bindBridgeToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTokenCaller{contract: contract}, nil
}

// NewBridgeTokenTransactor creates a new write-only instance of BridgeToken, bound to a specific deployed contract.
func NewBridgeTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTokenTransactor, error) {
	contract, err := bindBridgeToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &BridgeTokenTransactor{contract: contract}, nil
}

// bindBridgeToken binds a generic wrapper to an already deployed contract.
func bindBridgeToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeToken *BridgeTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BridgeToken.Contract.BridgeTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeToken *BridgeTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeToken.Contract.BridgeTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeToken *BridgeTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeToken.Contract.BridgeTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeToken *BridgeTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BridgeToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeToken *BridgeTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeToken *BridgeTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(_allowance uint256)
func (_BridgeToken *BridgeTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BridgeToken.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(_allowance uint256)
func (_BridgeToken *BridgeTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BridgeToken.Contract.Allowance(&_BridgeToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(_allowance uint256)
func (_BridgeToken *BridgeTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BridgeToken.Contract.Allowance(&_BridgeToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(value uint256)
func (_BridgeToken *BridgeTokenCaller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BridgeToken.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(value uint256)
func (_BridgeToken *BridgeTokenSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _BridgeToken.Contract.BalanceOf(&_BridgeToken.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(value uint256)
func (_BridgeToken *BridgeTokenCallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _BridgeToken.Contract.BalanceOf(&_BridgeToken.CallOpts, who)
}

// ContractHash is a free data retrieval call binding the contract method 0xb16a3335.
//
// Solidity: function contractHash(wad uint128) constant returns(bytes32)
func (_BridgeToken *BridgeTokenCaller) ContractHash(opts *bind.CallOpts, wad *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _BridgeToken.contract.Call(opts, out, "contractHash", wad)
	return *ret0, err
}

// ContractHash is a free data retrieval call binding the contract method 0xb16a3335.
//
// Solidity: function contractHash(wad uint128) constant returns(bytes32)
func (_BridgeToken *BridgeTokenSession) ContractHash(wad *big.Int) ([32]byte, error) {
	return _BridgeToken.Contract.ContractHash(&_BridgeToken.CallOpts, wad)
}

// ContractHash is a free data retrieval call binding the contract method 0xb16a3335.
//
// Solidity: function contractHash(wad uint128) constant returns(bytes32)
func (_BridgeToken *BridgeTokenCallerSession) ContractHash(wad *big.Int) ([32]byte, error) {
	return _BridgeToken.Contract.ContractHash(&_BridgeToken.CallOpts, wad)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_BridgeToken *BridgeTokenCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BridgeToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_BridgeToken *BridgeTokenSession) Decimals() (*big.Int, error) {
	return _BridgeToken.Contract.Decimals(&_BridgeToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_BridgeToken *BridgeTokenCallerSession) Decimals() (*big.Int, error) {
	return _BridgeToken.Contract.Decimals(&_BridgeToken.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x60999b3b.
//
// Solidity: function hasRole(roleName string, symbol_ bytes32, rolesContract address) constant returns(bool)
func (_BridgeToken *BridgeTokenCaller) HasRole(opts *bind.CallOpts, roleName string, symbol_ [32]byte, rolesContract common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BridgeToken.contract.Call(opts, out, "hasRole", roleName, symbol_, rolesContract)
	return *ret0, err
}

// HasRole is a free data retrieval call binding the contract method 0x60999b3b.
//
// Solidity: function hasRole(roleName string, symbol_ bytes32, rolesContract address) constant returns(bool)
func (_BridgeToken *BridgeTokenSession) HasRole(roleName string, symbol_ [32]byte, rolesContract common.Address) (bool, error) {
	return _BridgeToken.Contract.HasRole(&_BridgeToken.CallOpts, roleName, symbol_, rolesContract)
}

// HasRole is a free data retrieval call binding the contract method 0x60999b3b.
//
// Solidity: function hasRole(roleName string, symbol_ bytes32, rolesContract address) constant returns(bool)
func (_BridgeToken *BridgeTokenCallerSession) HasRole(roleName string, symbol_ [32]byte, rolesContract common.Address) (bool, error) {
	return _BridgeToken.Contract.HasRole(&_BridgeToken.CallOpts, roleName, symbol_, rolesContract)
}

// Logic is a free data retrieval call binding the contract method 0x2e259fbb.
//
// Solidity: function logic(owner address, spender address) constant returns(address)
func (_BridgeToken *BridgeTokenCaller) Logic(opts *bind.CallOpts, owner common.Address, spender common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeToken.contract.Call(opts, out, "logic", owner, spender)
	return *ret0, err
}

// Logic is a free data retrieval call binding the contract method 0x2e259fbb.
//
// Solidity: function logic(owner address, spender address) constant returns(address)
func (_BridgeToken *BridgeTokenSession) Logic(owner common.Address, spender common.Address) (common.Address, error) {
	return _BridgeToken.Contract.Logic(&_BridgeToken.CallOpts, owner, spender)
}

// Logic is a free data retrieval call binding the contract method 0x2e259fbb.
//
// Solidity: function logic(owner address, spender address) constant returns(address)
func (_BridgeToken *BridgeTokenCallerSession) Logic(owner common.Address, spender common.Address) (common.Address, error) {
	return _BridgeToken.Contract.Logic(&_BridgeToken.CallOpts, owner, spender)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BridgeToken *BridgeTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BridgeToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BridgeToken *BridgeTokenSession) Name() (string, error) {
	return _BridgeToken.Contract.Name(&_BridgeToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BridgeToken *BridgeTokenCallerSession) Name() (string, error) {
	return _BridgeToken.Contract.Name(&_BridgeToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BridgeToken *BridgeTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BridgeToken *BridgeTokenSession) Owner() (common.Address, error) {
	return _BridgeToken.Contract.Owner(&_BridgeToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BridgeToken *BridgeTokenCallerSession) Owner() (common.Address, error) {
	return _BridgeToken.Contract.Owner(&_BridgeToken.CallOpts)
}

// Roles is a free data retrieval call binding the contract method 0x99374642.
//
// Solidity: function roles(who address) constant returns(address)
func (_BridgeToken *BridgeTokenCaller) Roles(opts *bind.CallOpts, who common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeToken.contract.Call(opts, out, "roles", who)
	return *ret0, err
}

// Roles is a free data retrieval call binding the contract method 0x99374642.
//
// Solidity: function roles(who address) constant returns(address)
func (_BridgeToken *BridgeTokenSession) Roles(who common.Address) (common.Address, error) {
	return _BridgeToken.Contract.Roles(&_BridgeToken.CallOpts, who)
}

// Roles is a free data retrieval call binding the contract method 0x99374642.
//
// Solidity: function roles(who address) constant returns(address)
func (_BridgeToken *BridgeTokenCallerSession) Roles(who common.Address) (common.Address, error) {
	return _BridgeToken.Contract.Roles(&_BridgeToken.CallOpts, who)
}

// SenderHasRole is a free data retrieval call binding the contract method 0x1ca03b8e.
//
// Solidity: function senderHasRole(roleName string) constant returns(bool)
func (_BridgeToken *BridgeTokenCaller) SenderHasRole(opts *bind.CallOpts, roleName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BridgeToken.contract.Call(opts, out, "senderHasRole", roleName)
	return *ret0, err
}

// SenderHasRole is a free data retrieval call binding the contract method 0x1ca03b8e.
//
// Solidity: function senderHasRole(roleName string) constant returns(bool)
func (_BridgeToken *BridgeTokenSession) SenderHasRole(roleName string) (bool, error) {
	return _BridgeToken.Contract.SenderHasRole(&_BridgeToken.CallOpts, roleName)
}

// SenderHasRole is a free data retrieval call binding the contract method 0x1ca03b8e.
//
// Solidity: function senderHasRole(roleName string) constant returns(bool)
func (_BridgeToken *BridgeTokenCallerSession) SenderHasRole(roleName string) (bool, error) {
	return _BridgeToken.Contract.SenderHasRole(&_BridgeToken.CallOpts, roleName)
}

// Stopped is a free data retrieval call binding the contract method 0x4e6bdd7a.
//
// Solidity: function stopped(src address, wad uint128) constant returns(bool)
func (_BridgeToken *BridgeTokenCaller) Stopped(opts *bind.CallOpts, src common.Address, wad *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BridgeToken.contract.Call(opts, out, "stopped", src, wad)
	return *ret0, err
}

// Stopped is a free data retrieval call binding the contract method 0x4e6bdd7a.
//
// Solidity: function stopped(src address, wad uint128) constant returns(bool)
func (_BridgeToken *BridgeTokenSession) Stopped(src common.Address, wad *big.Int) (bool, error) {
	return _BridgeToken.Contract.Stopped(&_BridgeToken.CallOpts, src, wad)
}

// Stopped is a free data retrieval call binding the contract method 0x4e6bdd7a.
//
// Solidity: function stopped(src address, wad uint128) constant returns(bool)
func (_BridgeToken *BridgeTokenCallerSession) Stopped(src common.Address, wad *big.Int) (bool, error) {
	return _BridgeToken.Contract.Stopped(&_BridgeToken.CallOpts, src, wad)
}

// Symbol is a free data retrieval call binding the contract method 0xac2401b4.
//
// Solidity: function symbol(dst address, wad uint256) constant returns(bytes32)
func (_BridgeToken *BridgeTokenCaller) Symbol(opts *bind.CallOpts, dst common.Address, wad *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _BridgeToken.contract.Call(opts, out, "symbol", dst, wad)
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0xac2401b4.
//
// Solidity: function symbol(dst address, wad uint256) constant returns(bytes32)
func (_BridgeToken *BridgeTokenSession) Symbol(dst common.Address, wad *big.Int) ([32]byte, error) {
	return _BridgeToken.Contract.Symbol(&_BridgeToken.CallOpts, dst, wad)
}

// Symbol is a free data retrieval call binding the contract method 0xac2401b4.
//
// Solidity: function symbol(dst address, wad uint256) constant returns(bytes32)
func (_BridgeToken *BridgeTokenCallerSession) Symbol(dst common.Address, wad *big.Int) ([32]byte, error) {
	return _BridgeToken.Contract.Symbol(&_BridgeToken.CallOpts, dst, wad)
}

// TotalSupply is a free data retrieval call binding the contract method 0xc415db13.
//
// Solidity: function totalSupply(roleName string) constant returns(uint256)
func (_BridgeToken *BridgeTokenCaller) TotalSupply(opts *bind.CallOpts, roleName string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BridgeToken.contract.Call(opts, out, "totalSupply", roleName)
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0xc415db13.
//
// Solidity: function totalSupply(roleName string) constant returns(uint256)
func (_BridgeToken *BridgeTokenSession) TotalSupply(roleName string) (*big.Int, error) {
	return _BridgeToken.Contract.TotalSupply(&_BridgeToken.CallOpts, roleName)
}

// TotalSupply is a free data retrieval call binding the contract method 0xc415db13.
//
// Solidity: function totalSupply(roleName string) constant returns(uint256)
func (_BridgeToken *BridgeTokenCallerSession) TotalSupply(roleName string) (*big.Int, error) {
	return _BridgeToken.Contract.TotalSupply(&_BridgeToken.CallOpts, roleName)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(guy address, wad uint256) returns(bool)
func (_BridgeToken *BridgeTokenTransactor) Approve(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeToken.contract.Transact(opts, "approve", guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(guy address, wad uint256) returns(bool)
func (_BridgeToken *BridgeTokenSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.Approve(&_BridgeToken.TransactOpts, guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(guy address, wad uint256) returns(bool)
func (_BridgeToken *BridgeTokenTransactorSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.Approve(&_BridgeToken.TransactOpts, guy, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x90bc1693.
//
// Solidity: function burn(wad uint128) returns(bytes32)
func (_BridgeToken *BridgeTokenTransactor) Burn(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _BridgeToken.contract.Transact(opts, "burn", wad)
}

// Burn is a paid mutator transaction binding the contract method 0x90bc1693.
//
// Solidity: function burn(wad uint128) returns(bytes32)
func (_BridgeToken *BridgeTokenSession) Burn(wad *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.Burn(&_BridgeToken.TransactOpts, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x90bc1693.
//
// Solidity: function burn(wad uint128) returns(bytes32)
func (_BridgeToken *BridgeTokenTransactorSession) Burn(wad *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.Burn(&_BridgeToken.TransactOpts, wad)
}

// MintFor is a paid mutator transaction binding the contract method 0x27cfc219.
//
// Solidity: function mintFor(recipient address, wad uint128) returns(uint256)
func (_BridgeToken *BridgeTokenTransactor) MintFor(opts *bind.TransactOpts, recipient common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeToken.contract.Transact(opts, "mintFor", recipient, wad)
}

// MintFor is a paid mutator transaction binding the contract method 0x27cfc219.
//
// Solidity: function mintFor(recipient address, wad uint128) returns(uint256)
func (_BridgeToken *BridgeTokenSession) MintFor(recipient common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.MintFor(&_BridgeToken.TransactOpts, recipient, wad)
}

// MintFor is a paid mutator transaction binding the contract method 0x27cfc219.
//
// Solidity: function mintFor(recipient address, wad uint128) returns(uint256)
func (_BridgeToken *BridgeTokenTransactorSession) MintFor(recipient common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.MintFor(&_BridgeToken.TransactOpts, recipient, wad)
}

// Pull is a paid mutator transaction binding the contract method 0x8402181f.
//
// Solidity: function pull(src address, wad uint128) returns(bool)
func (_BridgeToken *BridgeTokenTransactor) Pull(opts *bind.TransactOpts, src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeToken.contract.Transact(opts, "pull", src, wad)
}

// Pull is a paid mutator transaction binding the contract method 0x8402181f.
//
// Solidity: function pull(src address, wad uint128) returns(bool)
func (_BridgeToken *BridgeTokenSession) Pull(src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.Pull(&_BridgeToken.TransactOpts, src, wad)
}

// Pull is a paid mutator transaction binding the contract method 0x8402181f.
//
// Solidity: function pull(src address, wad uint128) returns(bool)
func (_BridgeToken *BridgeTokenTransactorSession) Pull(src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.Pull(&_BridgeToken.TransactOpts, src, wad)
}

// RepayUou is a paid mutator transaction binding the contract method 0x0badbc55.
//
// Solidity: function repayUou(brgAmount uint128, vault address, uouIndex uint256) returns()
func (_BridgeToken *BridgeTokenTransactor) RepayUou(opts *bind.TransactOpts, brgAmount *big.Int, vault common.Address, uouIndex *big.Int) (*types.Transaction, error) {
	return _BridgeToken.contract.Transact(opts, "repayUou", brgAmount, vault, uouIndex)
}

// RepayUou is a paid mutator transaction binding the contract method 0x0badbc55.
//
// Solidity: function repayUou(brgAmount uint128, vault address, uouIndex uint256) returns()
func (_BridgeToken *BridgeTokenSession) RepayUou(brgAmount *big.Int, vault common.Address, uouIndex *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.RepayUou(&_BridgeToken.TransactOpts, brgAmount, vault, uouIndex)
}

// RepayUou is a paid mutator transaction binding the contract method 0x0badbc55.
//
// Solidity: function repayUou(brgAmount uint128, vault address, uouIndex uint256) returns()
func (_BridgeToken *BridgeTokenTransactorSession) RepayUou(brgAmount *big.Int, vault common.Address, uouIndex *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.RepayUou(&_BridgeToken.TransactOpts, brgAmount, vault, uouIndex)
}

// Restart is a paid mutator transaction binding the contract method 0xf0be5a7d.
//
// Solidity: function restart(src address, dst address, wad uint256) returns(bool)
func (_BridgeToken *BridgeTokenTransactor) Restart(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeToken.contract.Transact(opts, "restart", src, dst, wad)
}

// Restart is a paid mutator transaction binding the contract method 0xf0be5a7d.
//
// Solidity: function restart(src address, dst address, wad uint256) returns(bool)
func (_BridgeToken *BridgeTokenSession) Restart(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.Restart(&_BridgeToken.TransactOpts, src, dst, wad)
}

// Restart is a paid mutator transaction binding the contract method 0xf0be5a7d.
//
// Solidity: function restart(src address, dst address, wad uint256) returns(bool)
func (_BridgeToken *BridgeTokenTransactorSession) Restart(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.Restart(&_BridgeToken.TransactOpts, src, dst, wad)
}

// SetLogic is a paid mutator transaction binding the contract method 0x71857000.
//
// Solidity: function setLogic(logic_ address) returns(bool)
func (_BridgeToken *BridgeTokenTransactor) SetLogic(opts *bind.TransactOpts, logic_ common.Address) (*types.Transaction, error) {
	return _BridgeToken.contract.Transact(opts, "setLogic", logic_)
}

// SetLogic is a paid mutator transaction binding the contract method 0x71857000.
//
// Solidity: function setLogic(logic_ address) returns(bool)
func (_BridgeToken *BridgeTokenSession) SetLogic(logic_ common.Address) (*types.Transaction, error) {
	return _BridgeToken.Contract.SetLogic(&_BridgeToken.TransactOpts, logic_)
}

// SetLogic is a paid mutator transaction binding the contract method 0x71857000.
//
// Solidity: function setLogic(logic_ address) returns(bool)
func (_BridgeToken *BridgeTokenTransactorSession) SetLogic(logic_ common.Address) (*types.Transaction, error) {
	return _BridgeToken.Contract.SetLogic(&_BridgeToken.TransactOpts, logic_)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(name_ string) returns(address)
func (_BridgeToken *BridgeTokenTransactor) SetName(opts *bind.TransactOpts, name_ string) (*types.Transaction, error) {
	return _BridgeToken.contract.Transact(opts, "setName", name_)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(name_ string) returns(address)
func (_BridgeToken *BridgeTokenSession) SetName(name_ string) (*types.Transaction, error) {
	return _BridgeToken.Contract.SetName(&_BridgeToken.TransactOpts, name_)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(name_ string) returns(address)
func (_BridgeToken *BridgeTokenTransactorSession) SetName(name_ string) (*types.Transaction, error) {
	return _BridgeToken.Contract.SetName(&_BridgeToken.TransactOpts, name_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns(uint256)
func (_BridgeToken *BridgeTokenTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _BridgeToken.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns(uint256)
func (_BridgeToken *BridgeTokenSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _BridgeToken.Contract.SetOwner(&_BridgeToken.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns(uint256)
func (_BridgeToken *BridgeTokenTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _BridgeToken.Contract.SetOwner(&_BridgeToken.TransactOpts, owner_)
}

// SetRolesContract is a paid mutator transaction binding the contract method 0xafa202ac.
//
// Solidity: function setRolesContract(roles_ address) returns()
func (_BridgeToken *BridgeTokenTransactor) SetRolesContract(opts *bind.TransactOpts, roles_ common.Address) (*types.Transaction, error) {
	return _BridgeToken.contract.Transact(opts, "setRolesContract", roles_)
}

// SetRolesContract is a paid mutator transaction binding the contract method 0xafa202ac.
//
// Solidity: function setRolesContract(roles_ address) returns()
func (_BridgeToken *BridgeTokenSession) SetRolesContract(roles_ common.Address) (*types.Transaction, error) {
	return _BridgeToken.Contract.SetRolesContract(&_BridgeToken.TransactOpts, roles_)
}

// SetRolesContract is a paid mutator transaction binding the contract method 0xafa202ac.
//
// Solidity: function setRolesContract(roles_ address) returns()
func (_BridgeToken *BridgeTokenTransactorSession) SetRolesContract(roles_ common.Address) (*types.Transaction, error) {
	return _BridgeToken.Contract.SetRolesContract(&_BridgeToken.TransactOpts, roles_)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_BridgeToken *BridgeTokenTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeToken.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_BridgeToken *BridgeTokenSession) Stop() (*types.Transaction, error) {
	return _BridgeToken.Contract.Stop(&_BridgeToken.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_BridgeToken *BridgeTokenTransactorSession) Stop() (*types.Transaction, error) {
	return _BridgeToken.Contract.Stop(&_BridgeToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(dst address, wad uint256) returns(bool)
func (_BridgeToken *BridgeTokenTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeToken.contract.Transact(opts, "transfer", dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(dst address, wad uint256) returns(bool)
func (_BridgeToken *BridgeTokenSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.Transfer(&_BridgeToken.TransactOpts, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(dst address, wad uint256) returns(bool)
func (_BridgeToken *BridgeTokenTransactorSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.Transfer(&_BridgeToken.TransactOpts, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_BridgeToken *BridgeTokenTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeToken.contract.Transact(opts, "transferFrom", src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_BridgeToken *BridgeTokenSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.TransferFrom(&_BridgeToken.TransactOpts, src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_BridgeToken *BridgeTokenTransactorSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.TransferFrom(&_BridgeToken.TransactOpts, src, dst, wad)
}

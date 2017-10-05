// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// RootABI is the input ABI used to generate the binding from.
const RootABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"swt_\",\"type\":\"address\"}],\"name\":\"setSWT\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"senderHasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"restart\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"userDirFactory_\",\"type\":\"address\"}],\"name\":\"setUserFactory\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"roles\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vaultFactory_\",\"type\":\"address\"}],\"name\":\"setVaultFactory\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"brg\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"userDirFactory\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vc\",\"type\":\"address\"}],\"name\":\"setVaultConfig\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"brg_\",\"type\":\"address\"}],\"name\":\"setBRG\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"userDirectories\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vaultConfig\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"walletFactory_\",\"type\":\"address\"}],\"name\":\"setWalletFactory\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"roles_\",\"type\":\"address\"}],\"name\":\"setRolesContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"walletFactory\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"removeDirectory\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vaultFactory\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"hasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"swt\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"t\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"addDirectory\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"brg_\",\"type\":\"address\"},{\"name\":\"swt_\",\"type\":\"address\"},{\"name\":\"vc\",\"type\":\"address\"},{\"name\":\"userDirFactory_\",\"type\":\"address\"},{\"name\":\"walletFactory_\",\"type\":\"address\"},{\"name\":\"vaultFactory_\",\"type\":\"address\"},{\"name\":\"rolesContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newUserDirectory\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogDirectoryAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"}]"

// RootBin is the compiled bytecode used for deploying new contracts.
const RootBin = `undefined`

// DeployRoot deploys a new Ethereum contract, binding an instance of Root to it.
func DeployRoot(auth *bind.TransactOpts, backend bind.ContractBackend, brg_ common.Address, swt_ common.Address, vc common.Address, userDirFactory_ common.Address, walletFactory_ common.Address, vaultFactory_ common.Address, rolesContract common.Address) (common.Address, *types.Transaction, *Root, error) {
	parsed, err := abi.JSON(strings.NewReader(RootABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RootBin), backend, brg_, swt_, vc, userDirFactory_, walletFactory_, vaultFactory_, rolesContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Root{RootCaller: RootCaller{contract: contract}, RootTransactor: RootTransactor{contract: contract}}, nil
}

// Root is an auto generated Go binding around an Ethereum contract.
type Root struct {
	RootCaller     // Read-only binding to the contract
	RootTransactor // Write-only binding to the contract
}

// RootCaller is an auto generated read-only Go binding around an Ethereum contract.
type RootCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RootTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RootSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RootSession struct {
	Contract     *Root             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RootCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RootCallerSession struct {
	Contract *RootCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RootTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RootTransactorSession struct {
	Contract     *RootTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RootRaw is an auto generated low-level Go binding around an Ethereum contract.
type RootRaw struct {
	Contract *Root // Generic contract binding to access the raw methods on
}

// RootCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RootCallerRaw struct {
	Contract *RootCaller // Generic read-only contract binding to access the raw methods on
}

// RootTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RootTransactorRaw struct {
	Contract *RootTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRoot creates a new instance of Root, bound to a specific deployed contract.
func NewRoot(address common.Address, backend bind.ContractBackend) (*Root, error) {
	contract, err := bindRoot(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Root{RootCaller: RootCaller{contract: contract}, RootTransactor: RootTransactor{contract: contract}}, nil
}

// NewRootCaller creates a new read-only instance of Root, bound to a specific deployed contract.
func NewRootCaller(address common.Address, caller bind.ContractCaller) (*RootCaller, error) {
	contract, err := bindRoot(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &RootCaller{contract: contract}, nil
}

// NewRootTransactor creates a new write-only instance of Root, bound to a specific deployed contract.
func NewRootTransactor(address common.Address, transactor bind.ContractTransactor) (*RootTransactor, error) {
	contract, err := bindRoot(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &RootTransactor{contract: contract}, nil
}

// bindRoot binds a generic wrapper to an already deployed contract.
func bindRoot(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RootABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Root *RootRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Root.Contract.RootCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Root *RootRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Root.Contract.RootTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Root *RootRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Root.Contract.RootTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Root *RootCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Root.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Root *RootTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Root.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Root *RootTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Root.Contract.contract.Transact(opts, method, params...)
}

// Brg is a free data retrieval call binding the contract method 0x4f9c8fe8.
//
// Solidity: function brg() constant returns(address)
func (_Root *RootCaller) Brg(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Root.contract.Call(opts, out, "brg")
	return *ret0, err
}

// Brg is a free data retrieval call binding the contract method 0x4f9c8fe8.
//
// Solidity: function brg() constant returns(address)
func (_Root *RootSession) Brg() (common.Address, error) {
	return _Root.Contract.Brg(&_Root.CallOpts)
}

// Brg is a free data retrieval call binding the contract method 0x4f9c8fe8.
//
// Solidity: function brg() constant returns(address)
func (_Root *RootCallerSession) Brg() (common.Address, error) {
	return _Root.Contract.Brg(&_Root.CallOpts)
}

// ContractHash is a free data retrieval call binding the contract method 0x904c6094.
//
// Solidity: function contractHash() constant returns(bytes32)
func (_Root *RootCaller) ContractHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Root.contract.Call(opts, out, "contractHash")
	return *ret0, err
}

// ContractHash is a free data retrieval call binding the contract method 0x904c6094.
//
// Solidity: function contractHash() constant returns(bytes32)
func (_Root *RootSession) ContractHash() ([32]byte, error) {
	return _Root.Contract.ContractHash(&_Root.CallOpts)
}

// ContractHash is a free data retrieval call binding the contract method 0x904c6094.
//
// Solidity: function contractHash() constant returns(bytes32)
func (_Root *RootCallerSession) ContractHash() ([32]byte, error) {
	return _Root.Contract.ContractHash(&_Root.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0xe3c33a9b.
//
// Solidity: function hasRole(roleName string) constant returns(bool)
func (_Root *RootCaller) HasRole(opts *bind.CallOpts, roleName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Root.contract.Call(opts, out, "hasRole", roleName)
	return *ret0, err
}

// HasRole is a free data retrieval call binding the contract method 0xe3c33a9b.
//
// Solidity: function hasRole(roleName string) constant returns(bool)
func (_Root *RootSession) HasRole(roleName string) (bool, error) {
	return _Root.Contract.HasRole(&_Root.CallOpts, roleName)
}

// HasRole is a free data retrieval call binding the contract method 0xe3c33a9b.
//
// Solidity: function hasRole(roleName string) constant returns(bool)
func (_Root *RootCallerSession) HasRole(roleName string) (bool, error) {
	return _Root.Contract.HasRole(&_Root.CallOpts, roleName)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Root *RootCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Root.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Root *RootSession) Owner() (common.Address, error) {
	return _Root.Contract.Owner(&_Root.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Root *RootCallerSession) Owner() (common.Address, error) {
	return _Root.Contract.Owner(&_Root.CallOpts)
}

// Roles is a free data retrieval call binding the contract method 0x392f5f64.
//
// Solidity: function roles() constant returns(address)
func (_Root *RootCaller) Roles(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Root.contract.Call(opts, out, "roles")
	return *ret0, err
}

// Roles is a free data retrieval call binding the contract method 0x392f5f64.
//
// Solidity: function roles() constant returns(address)
func (_Root *RootSession) Roles() (common.Address, error) {
	return _Root.Contract.Roles(&_Root.CallOpts)
}

// Roles is a free data retrieval call binding the contract method 0x392f5f64.
//
// Solidity: function roles() constant returns(address)
func (_Root *RootCallerSession) Roles() (common.Address, error) {
	return _Root.Contract.Roles(&_Root.CallOpts)
}

// SenderHasRole is a free data retrieval call binding the contract method 0x1ca03b8e.
//
// Solidity: function senderHasRole(roleName string) constant returns(bool)
func (_Root *RootCaller) SenderHasRole(opts *bind.CallOpts, roleName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Root.contract.Call(opts, out, "senderHasRole", roleName)
	return *ret0, err
}

// SenderHasRole is a free data retrieval call binding the contract method 0x1ca03b8e.
//
// Solidity: function senderHasRole(roleName string) constant returns(bool)
func (_Root *RootSession) SenderHasRole(roleName string) (bool, error) {
	return _Root.Contract.SenderHasRole(&_Root.CallOpts, roleName)
}

// SenderHasRole is a free data retrieval call binding the contract method 0x1ca03b8e.
//
// Solidity: function senderHasRole(roleName string) constant returns(bool)
func (_Root *RootCallerSession) SenderHasRole(roleName string) (bool, error) {
	return _Root.Contract.SenderHasRole(&_Root.CallOpts, roleName)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_Root *RootCaller) Stopped(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Root.contract.Call(opts, out, "stopped")
	return *ret0, err
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_Root *RootSession) Stopped() (bool, error) {
	return _Root.Contract.Stopped(&_Root.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_Root *RootCallerSession) Stopped() (bool, error) {
	return _Root.Contract.Stopped(&_Root.CallOpts)
}

// Swt is a free data retrieval call binding the contract method 0xf016b6b6.
//
// Solidity: function swt() constant returns(address)
func (_Root *RootCaller) Swt(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Root.contract.Call(opts, out, "swt")
	return *ret0, err
}

// Swt is a free data retrieval call binding the contract method 0xf016b6b6.
//
// Solidity: function swt() constant returns(address)
func (_Root *RootSession) Swt() (common.Address, error) {
	return _Root.Contract.Swt(&_Root.CallOpts)
}

// Swt is a free data retrieval call binding the contract method 0xf016b6b6.
//
// Solidity: function swt() constant returns(address)
func (_Root *RootCallerSession) Swt() (common.Address, error) {
	return _Root.Contract.Swt(&_Root.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() constant returns(address)
func (_Root *RootCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Root.contract.Call(opts, out, "treasury")
	return *ret0, err
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() constant returns(address)
func (_Root *RootSession) Treasury() (common.Address, error) {
	return _Root.Contract.Treasury(&_Root.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() constant returns(address)
func (_Root *RootCallerSession) Treasury() (common.Address, error) {
	return _Root.Contract.Treasury(&_Root.CallOpts)
}

// UserDirFactory is a free data retrieval call binding the contract method 0x55f42148.
//
// Solidity: function userDirFactory() constant returns(address)
func (_Root *RootCaller) UserDirFactory(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Root.contract.Call(opts, out, "userDirFactory")
	return *ret0, err
}

// UserDirFactory is a free data retrieval call binding the contract method 0x55f42148.
//
// Solidity: function userDirFactory() constant returns(address)
func (_Root *RootSession) UserDirFactory() (common.Address, error) {
	return _Root.Contract.UserDirFactory(&_Root.CallOpts)
}

// UserDirFactory is a free data retrieval call binding the contract method 0x55f42148.
//
// Solidity: function userDirFactory() constant returns(address)
func (_Root *RootCallerSession) UserDirFactory() (common.Address, error) {
	return _Root.Contract.UserDirFactory(&_Root.CallOpts)
}

// UserDirectories is a free data retrieval call binding the contract method 0x7b2fa97d.
//
// Solidity: function userDirectories( address) constant returns(address)
func (_Root *RootCaller) UserDirectories(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Root.contract.Call(opts, out, "userDirectories", arg0)
	return *ret0, err
}

// UserDirectories is a free data retrieval call binding the contract method 0x7b2fa97d.
//
// Solidity: function userDirectories( address) constant returns(address)
func (_Root *RootSession) UserDirectories(arg0 common.Address) (common.Address, error) {
	return _Root.Contract.UserDirectories(&_Root.CallOpts, arg0)
}

// UserDirectories is a free data retrieval call binding the contract method 0x7b2fa97d.
//
// Solidity: function userDirectories( address) constant returns(address)
func (_Root *RootCallerSession) UserDirectories(arg0 common.Address) (common.Address, error) {
	return _Root.Contract.UserDirectories(&_Root.CallOpts, arg0)
}

// VaultConfig is a free data retrieval call binding the contract method 0x7cc34bb4.
//
// Solidity: function vaultConfig() constant returns(address)
func (_Root *RootCaller) VaultConfig(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Root.contract.Call(opts, out, "vaultConfig")
	return *ret0, err
}

// VaultConfig is a free data retrieval call binding the contract method 0x7cc34bb4.
//
// Solidity: function vaultConfig() constant returns(address)
func (_Root *RootSession) VaultConfig() (common.Address, error) {
	return _Root.Contract.VaultConfig(&_Root.CallOpts)
}

// VaultConfig is a free data retrieval call binding the contract method 0x7cc34bb4.
//
// Solidity: function vaultConfig() constant returns(address)
func (_Root *RootCallerSession) VaultConfig() (common.Address, error) {
	return _Root.Contract.VaultConfig(&_Root.CallOpts)
}

// VaultFactory is a free data retrieval call binding the contract method 0xd8a06f73.
//
// Solidity: function vaultFactory() constant returns(address)
func (_Root *RootCaller) VaultFactory(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Root.contract.Call(opts, out, "vaultFactory")
	return *ret0, err
}

// VaultFactory is a free data retrieval call binding the contract method 0xd8a06f73.
//
// Solidity: function vaultFactory() constant returns(address)
func (_Root *RootSession) VaultFactory() (common.Address, error) {
	return _Root.Contract.VaultFactory(&_Root.CallOpts)
}

// VaultFactory is a free data retrieval call binding the contract method 0xd8a06f73.
//
// Solidity: function vaultFactory() constant returns(address)
func (_Root *RootCallerSession) VaultFactory() (common.Address, error) {
	return _Root.Contract.VaultFactory(&_Root.CallOpts)
}

// WalletFactory is a free data retrieval call binding the contract method 0xc5c03699.
//
// Solidity: function walletFactory() constant returns(address)
func (_Root *RootCaller) WalletFactory(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Root.contract.Call(opts, out, "walletFactory")
	return *ret0, err
}

// WalletFactory is a free data retrieval call binding the contract method 0xc5c03699.
//
// Solidity: function walletFactory() constant returns(address)
func (_Root *RootSession) WalletFactory() (common.Address, error) {
	return _Root.Contract.WalletFactory(&_Root.CallOpts)
}

// WalletFactory is a free data retrieval call binding the contract method 0xc5c03699.
//
// Solidity: function walletFactory() constant returns(address)
func (_Root *RootCallerSession) WalletFactory() (common.Address, error) {
	return _Root.Contract.WalletFactory(&_Root.CallOpts)
}

// AddDirectory is a paid mutator transaction binding the contract method 0xfb244019.
//
// Solidity: function addDirectory(owner address) returns()
func (_Root *RootTransactor) AddDirectory(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _Root.contract.Transact(opts, "addDirectory", owner)
}

// AddDirectory is a paid mutator transaction binding the contract method 0xfb244019.
//
// Solidity: function addDirectory(owner address) returns()
func (_Root *RootSession) AddDirectory(owner common.Address) (*types.Transaction, error) {
	return _Root.Contract.AddDirectory(&_Root.TransactOpts, owner)
}

// AddDirectory is a paid mutator transaction binding the contract method 0xfb244019.
//
// Solidity: function addDirectory(owner address) returns()
func (_Root *RootTransactorSession) AddDirectory(owner common.Address) (*types.Transaction, error) {
	return _Root.Contract.AddDirectory(&_Root.TransactOpts, owner)
}

// RemoveDirectory is a paid mutator transaction binding the contract method 0xd25ad88b.
//
// Solidity: function removeDirectory(owner address) returns()
func (_Root *RootTransactor) RemoveDirectory(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _Root.contract.Transact(opts, "removeDirectory", owner)
}

// RemoveDirectory is a paid mutator transaction binding the contract method 0xd25ad88b.
//
// Solidity: function removeDirectory(owner address) returns()
func (_Root *RootSession) RemoveDirectory(owner common.Address) (*types.Transaction, error) {
	return _Root.Contract.RemoveDirectory(&_Root.TransactOpts, owner)
}

// RemoveDirectory is a paid mutator transaction binding the contract method 0xd25ad88b.
//
// Solidity: function removeDirectory(owner address) returns()
func (_Root *RootTransactorSession) RemoveDirectory(owner common.Address) (*types.Transaction, error) {
	return _Root.Contract.RemoveDirectory(&_Root.TransactOpts, owner)
}

// Restart is a paid mutator transaction binding the contract method 0x1ef3755d.
//
// Solidity: function restart() returns()
func (_Root *RootTransactor) Restart(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Root.contract.Transact(opts, "restart")
}

// Restart is a paid mutator transaction binding the contract method 0x1ef3755d.
//
// Solidity: function restart() returns()
func (_Root *RootSession) Restart() (*types.Transaction, error) {
	return _Root.Contract.Restart(&_Root.TransactOpts)
}

// Restart is a paid mutator transaction binding the contract method 0x1ef3755d.
//
// Solidity: function restart() returns()
func (_Root *RootTransactorSession) Restart() (*types.Transaction, error) {
	return _Root.Contract.Restart(&_Root.TransactOpts)
}

// SetBRG is a paid mutator transaction binding the contract method 0x679ccc6d.
//
// Solidity: function setBRG(brg_ address) returns()
func (_Root *RootTransactor) SetBRG(opts *bind.TransactOpts, brg_ common.Address) (*types.Transaction, error) {
	return _Root.contract.Transact(opts, "setBRG", brg_)
}

// SetBRG is a paid mutator transaction binding the contract method 0x679ccc6d.
//
// Solidity: function setBRG(brg_ address) returns()
func (_Root *RootSession) SetBRG(brg_ common.Address) (*types.Transaction, error) {
	return _Root.Contract.SetBRG(&_Root.TransactOpts, brg_)
}

// SetBRG is a paid mutator transaction binding the contract method 0x679ccc6d.
//
// Solidity: function setBRG(brg_ address) returns()
func (_Root *RootTransactorSession) SetBRG(brg_ common.Address) (*types.Transaction, error) {
	return _Root.Contract.SetBRG(&_Root.TransactOpts, brg_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_Root *RootTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _Root.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_Root *RootSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Root.Contract.SetOwner(&_Root.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_Root *RootTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Root.Contract.SetOwner(&_Root.TransactOpts, owner_)
}

// SetRolesContract is a paid mutator transaction binding the contract method 0xafa202ac.
//
// Solidity: function setRolesContract(roles_ address) returns()
func (_Root *RootTransactor) SetRolesContract(opts *bind.TransactOpts, roles_ common.Address) (*types.Transaction, error) {
	return _Root.contract.Transact(opts, "setRolesContract", roles_)
}

// SetRolesContract is a paid mutator transaction binding the contract method 0xafa202ac.
//
// Solidity: function setRolesContract(roles_ address) returns()
func (_Root *RootSession) SetRolesContract(roles_ common.Address) (*types.Transaction, error) {
	return _Root.Contract.SetRolesContract(&_Root.TransactOpts, roles_)
}

// SetRolesContract is a paid mutator transaction binding the contract method 0xafa202ac.
//
// Solidity: function setRolesContract(roles_ address) returns()
func (_Root *RootTransactorSession) SetRolesContract(roles_ common.Address) (*types.Transaction, error) {
	return _Root.Contract.SetRolesContract(&_Root.TransactOpts, roles_)
}

// SetSWT is a paid mutator transaction binding the contract method 0x0d1a8df0.
//
// Solidity: function setSWT(swt_ address) returns()
func (_Root *RootTransactor) SetSWT(opts *bind.TransactOpts, swt_ common.Address) (*types.Transaction, error) {
	return _Root.contract.Transact(opts, "setSWT", swt_)
}

// SetSWT is a paid mutator transaction binding the contract method 0x0d1a8df0.
//
// Solidity: function setSWT(swt_ address) returns()
func (_Root *RootSession) SetSWT(swt_ common.Address) (*types.Transaction, error) {
	return _Root.Contract.SetSWT(&_Root.TransactOpts, swt_)
}

// SetSWT is a paid mutator transaction binding the contract method 0x0d1a8df0.
//
// Solidity: function setSWT(swt_ address) returns()
func (_Root *RootTransactorSession) SetSWT(swt_ common.Address) (*types.Transaction, error) {
	return _Root.Contract.SetSWT(&_Root.TransactOpts, swt_)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(t address) returns()
func (_Root *RootTransactor) SetTreasury(opts *bind.TransactOpts, t common.Address) (*types.Transaction, error) {
	return _Root.contract.Transact(opts, "setTreasury", t)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(t address) returns()
func (_Root *RootSession) SetTreasury(t common.Address) (*types.Transaction, error) {
	return _Root.Contract.SetTreasury(&_Root.TransactOpts, t)
}

// SetTreasury is a paid mutator transaction binding the contract method 0xf0f44260.
//
// Solidity: function setTreasury(t address) returns()
func (_Root *RootTransactorSession) SetTreasury(t common.Address) (*types.Transaction, error) {
	return _Root.Contract.SetTreasury(&_Root.TransactOpts, t)
}

// SetUserFactory is a paid mutator transaction binding the contract method 0x30cc1543.
//
// Solidity: function setUserFactory(userDirFactory_ address) returns()
func (_Root *RootTransactor) SetUserFactory(opts *bind.TransactOpts, userDirFactory_ common.Address) (*types.Transaction, error) {
	return _Root.contract.Transact(opts, "setUserFactory", userDirFactory_)
}

// SetUserFactory is a paid mutator transaction binding the contract method 0x30cc1543.
//
// Solidity: function setUserFactory(userDirFactory_ address) returns()
func (_Root *RootSession) SetUserFactory(userDirFactory_ common.Address) (*types.Transaction, error) {
	return _Root.Contract.SetUserFactory(&_Root.TransactOpts, userDirFactory_)
}

// SetUserFactory is a paid mutator transaction binding the contract method 0x30cc1543.
//
// Solidity: function setUserFactory(userDirFactory_ address) returns()
func (_Root *RootTransactorSession) SetUserFactory(userDirFactory_ common.Address) (*types.Transaction, error) {
	return _Root.Contract.SetUserFactory(&_Root.TransactOpts, userDirFactory_)
}

// SetVaultConfig is a paid mutator transaction binding the contract method 0x5faafcaa.
//
// Solidity: function setVaultConfig(vc address) returns()
func (_Root *RootTransactor) SetVaultConfig(opts *bind.TransactOpts, vc common.Address) (*types.Transaction, error) {
	return _Root.contract.Transact(opts, "setVaultConfig", vc)
}

// SetVaultConfig is a paid mutator transaction binding the contract method 0x5faafcaa.
//
// Solidity: function setVaultConfig(vc address) returns()
func (_Root *RootSession) SetVaultConfig(vc common.Address) (*types.Transaction, error) {
	return _Root.Contract.SetVaultConfig(&_Root.TransactOpts, vc)
}

// SetVaultConfig is a paid mutator transaction binding the contract method 0x5faafcaa.
//
// Solidity: function setVaultConfig(vc address) returns()
func (_Root *RootTransactorSession) SetVaultConfig(vc common.Address) (*types.Transaction, error) {
	return _Root.Contract.SetVaultConfig(&_Root.TransactOpts, vc)
}

// SetVaultFactory is a paid mutator transaction binding the contract method 0x3ea7fbdb.
//
// Solidity: function setVaultFactory(vaultFactory_ address) returns()
func (_Root *RootTransactor) SetVaultFactory(opts *bind.TransactOpts, vaultFactory_ common.Address) (*types.Transaction, error) {
	return _Root.contract.Transact(opts, "setVaultFactory", vaultFactory_)
}

// SetVaultFactory is a paid mutator transaction binding the contract method 0x3ea7fbdb.
//
// Solidity: function setVaultFactory(vaultFactory_ address) returns()
func (_Root *RootSession) SetVaultFactory(vaultFactory_ common.Address) (*types.Transaction, error) {
	return _Root.Contract.SetVaultFactory(&_Root.TransactOpts, vaultFactory_)
}

// SetVaultFactory is a paid mutator transaction binding the contract method 0x3ea7fbdb.
//
// Solidity: function setVaultFactory(vaultFactory_ address) returns()
func (_Root *RootTransactorSession) SetVaultFactory(vaultFactory_ common.Address) (*types.Transaction, error) {
	return _Root.Contract.SetVaultFactory(&_Root.TransactOpts, vaultFactory_)
}

// SetWalletFactory is a paid mutator transaction binding the contract method 0x7ebf879c.
//
// Solidity: function setWalletFactory(walletFactory_ address) returns()
func (_Root *RootTransactor) SetWalletFactory(opts *bind.TransactOpts, walletFactory_ common.Address) (*types.Transaction, error) {
	return _Root.contract.Transact(opts, "setWalletFactory", walletFactory_)
}

// SetWalletFactory is a paid mutator transaction binding the contract method 0x7ebf879c.
//
// Solidity: function setWalletFactory(walletFactory_ address) returns()
func (_Root *RootSession) SetWalletFactory(walletFactory_ common.Address) (*types.Transaction, error) {
	return _Root.Contract.SetWalletFactory(&_Root.TransactOpts, walletFactory_)
}

// SetWalletFactory is a paid mutator transaction binding the contract method 0x7ebf879c.
//
// Solidity: function setWalletFactory(walletFactory_ address) returns()
func (_Root *RootTransactorSession) SetWalletFactory(walletFactory_ common.Address) (*types.Transaction, error) {
	return _Root.Contract.SetWalletFactory(&_Root.TransactOpts, walletFactory_)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Root *RootTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Root.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Root *RootSession) Stop() (*types.Transaction, error) {
	return _Root.Contract.Stop(&_Root.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Root *RootTransactorSession) Stop() (*types.Transaction, error) {
	return _Root.Contract.Stop(&_Root.TransactOpts)
}

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

// UserDirectoryABI is the input ABI used to generate the binding from.
const UserDirectoryABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"senderHasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"restart\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"walletCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"roles\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"brg\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"wallets\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kyc\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wallet_\",\"type\":\"address\"}],\"name\":\"removeWallet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"profile\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"roles_\",\"type\":\"address\"}],\"name\":\"setRolesContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"kyc_\",\"type\":\"bool\"}],\"name\":\"setKYC\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"hasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currency\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes3\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"root\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"addWallet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"root_\",\"type\":\"address\"},{\"name\":\"owner_\",\"type\":\"address\"},{\"name\":\"rolesContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newWallet\",\"type\":\"address\"}],\"name\":\"LogWalletAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"removedWallet\",\"type\":\"address\"}],\"name\":\"LogWalletRemoved\",\"type\":\"event\"}]"

// UserDirectoryBin is the compiled bytecode used for deploying new contracts.
const UserDirectoryBin = `undefined`

// DeployUserDirectory deploys a new Ethereum contract, binding an instance of UserDirectory to it.
func DeployUserDirectory(auth *bind.TransactOpts, backend bind.ContractBackend, root_ common.Address, owner_ common.Address, rolesContract common.Address) (common.Address, *types.Transaction, *UserDirectory, error) {
	parsed, err := abi.JSON(strings.NewReader(UserDirectoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UserDirectoryBin), backend, root_, owner_, rolesContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UserDirectory{UserDirectoryCaller: UserDirectoryCaller{contract: contract}, UserDirectoryTransactor: UserDirectoryTransactor{contract: contract}}, nil
}

// UserDirectory is an auto generated Go binding around an Ethereum contract.
type UserDirectory struct {
	UserDirectoryCaller     // Read-only binding to the contract
	UserDirectoryTransactor // Write-only binding to the contract
}

// UserDirectoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type UserDirectoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserDirectoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UserDirectoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserDirectorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UserDirectorySession struct {
	Contract     *UserDirectory    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserDirectoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UserDirectoryCallerSession struct {
	Contract *UserDirectoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// UserDirectoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UserDirectoryTransactorSession struct {
	Contract     *UserDirectoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// UserDirectoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type UserDirectoryRaw struct {
	Contract *UserDirectory // Generic contract binding to access the raw methods on
}

// UserDirectoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UserDirectoryCallerRaw struct {
	Contract *UserDirectoryCaller // Generic read-only contract binding to access the raw methods on
}

// UserDirectoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UserDirectoryTransactorRaw struct {
	Contract *UserDirectoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUserDirectory creates a new instance of UserDirectory, bound to a specific deployed contract.
func NewUserDirectory(address common.Address, backend bind.ContractBackend) (*UserDirectory, error) {
	contract, err := bindUserDirectory(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UserDirectory{UserDirectoryCaller: UserDirectoryCaller{contract: contract}, UserDirectoryTransactor: UserDirectoryTransactor{contract: contract}}, nil
}

// NewUserDirectoryCaller creates a new read-only instance of UserDirectory, bound to a specific deployed contract.
func NewUserDirectoryCaller(address common.Address, caller bind.ContractCaller) (*UserDirectoryCaller, error) {
	contract, err := bindUserDirectory(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &UserDirectoryCaller{contract: contract}, nil
}

// NewUserDirectoryTransactor creates a new write-only instance of UserDirectory, bound to a specific deployed contract.
func NewUserDirectoryTransactor(address common.Address, transactor bind.ContractTransactor) (*UserDirectoryTransactor, error) {
	contract, err := bindUserDirectory(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &UserDirectoryTransactor{contract: contract}, nil
}

// bindUserDirectory binds a generic wrapper to an already deployed contract.
func bindUserDirectory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UserDirectoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserDirectory *UserDirectoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UserDirectory.Contract.UserDirectoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserDirectory *UserDirectoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserDirectory.Contract.UserDirectoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserDirectory *UserDirectoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserDirectory.Contract.UserDirectoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserDirectory *UserDirectoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UserDirectory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserDirectory *UserDirectoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserDirectory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserDirectory *UserDirectoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserDirectory.Contract.contract.Transact(opts, method, params...)
}

// Brg is a free data retrieval call binding the contract method 0x4f9c8fe8.
//
// Solidity: function brg() constant returns(address)
func (_UserDirectory *UserDirectoryCaller) Brg(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UserDirectory.contract.Call(opts, out, "brg")
	return *ret0, err
}

// Brg is a free data retrieval call binding the contract method 0x4f9c8fe8.
//
// Solidity: function brg() constant returns(address)
func (_UserDirectory *UserDirectorySession) Brg() (common.Address, error) {
	return _UserDirectory.Contract.Brg(&_UserDirectory.CallOpts)
}

// Brg is a free data retrieval call binding the contract method 0x4f9c8fe8.
//
// Solidity: function brg() constant returns(address)
func (_UserDirectory *UserDirectoryCallerSession) Brg() (common.Address, error) {
	return _UserDirectory.Contract.Brg(&_UserDirectory.CallOpts)
}

// ContractHash is a free data retrieval call binding the contract method 0x904c6094.
//
// Solidity: function contractHash() constant returns(bytes32)
func (_UserDirectory *UserDirectoryCaller) ContractHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _UserDirectory.contract.Call(opts, out, "contractHash")
	return *ret0, err
}

// ContractHash is a free data retrieval call binding the contract method 0x904c6094.
//
// Solidity: function contractHash() constant returns(bytes32)
func (_UserDirectory *UserDirectorySession) ContractHash() ([32]byte, error) {
	return _UserDirectory.Contract.ContractHash(&_UserDirectory.CallOpts)
}

// ContractHash is a free data retrieval call binding the contract method 0x904c6094.
//
// Solidity: function contractHash() constant returns(bytes32)
func (_UserDirectory *UserDirectoryCallerSession) ContractHash() ([32]byte, error) {
	return _UserDirectory.Contract.ContractHash(&_UserDirectory.CallOpts)
}

// Currency is a free data retrieval call binding the contract method 0xe5a6b10f.
//
// Solidity: function currency() constant returns(bytes3)
func (_UserDirectory *UserDirectoryCaller) Currency(opts *bind.CallOpts) ([3]byte, error) {
	var (
		ret0 = new([3]byte)
	)
	out := ret0
	err := _UserDirectory.contract.Call(opts, out, "currency")
	return *ret0, err
}

// Currency is a free data retrieval call binding the contract method 0xe5a6b10f.
//
// Solidity: function currency() constant returns(bytes3)
func (_UserDirectory *UserDirectorySession) Currency() ([3]byte, error) {
	return _UserDirectory.Contract.Currency(&_UserDirectory.CallOpts)
}

// Currency is a free data retrieval call binding the contract method 0xe5a6b10f.
//
// Solidity: function currency() constant returns(bytes3)
func (_UserDirectory *UserDirectoryCallerSession) Currency() ([3]byte, error) {
	return _UserDirectory.Contract.Currency(&_UserDirectory.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0xe3c33a9b.
//
// Solidity: function hasRole(roleName string) constant returns(bool)
func (_UserDirectory *UserDirectoryCaller) HasRole(opts *bind.CallOpts, roleName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _UserDirectory.contract.Call(opts, out, "hasRole", roleName)
	return *ret0, err
}

// HasRole is a free data retrieval call binding the contract method 0xe3c33a9b.
//
// Solidity: function hasRole(roleName string) constant returns(bool)
func (_UserDirectory *UserDirectorySession) HasRole(roleName string) (bool, error) {
	return _UserDirectory.Contract.HasRole(&_UserDirectory.CallOpts, roleName)
}

// HasRole is a free data retrieval call binding the contract method 0xe3c33a9b.
//
// Solidity: function hasRole(roleName string) constant returns(bool)
func (_UserDirectory *UserDirectoryCallerSession) HasRole(roleName string) (bool, error) {
	return _UserDirectory.Contract.HasRole(&_UserDirectory.CallOpts, roleName)
}

// Kyc is a free data retrieval call binding the contract method 0x90d6b45f.
//
// Solidity: function kyc() constant returns(bool)
func (_UserDirectory *UserDirectoryCaller) Kyc(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _UserDirectory.contract.Call(opts, out, "kyc")
	return *ret0, err
}

// Kyc is a free data retrieval call binding the contract method 0x90d6b45f.
//
// Solidity: function kyc() constant returns(bool)
func (_UserDirectory *UserDirectorySession) Kyc() (bool, error) {
	return _UserDirectory.Contract.Kyc(&_UserDirectory.CallOpts)
}

// Kyc is a free data retrieval call binding the contract method 0x90d6b45f.
//
// Solidity: function kyc() constant returns(bool)
func (_UserDirectory *UserDirectoryCallerSession) Kyc() (bool, error) {
	return _UserDirectory.Contract.Kyc(&_UserDirectory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_UserDirectory *UserDirectoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UserDirectory.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_UserDirectory *UserDirectorySession) Owner() (common.Address, error) {
	return _UserDirectory.Contract.Owner(&_UserDirectory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_UserDirectory *UserDirectoryCallerSession) Owner() (common.Address, error) {
	return _UserDirectory.Contract.Owner(&_UserDirectory.CallOpts)
}

// Profile is a free data retrieval call binding the contract method 0xab60636c.
//
// Solidity: function profile() constant returns(string)
func (_UserDirectory *UserDirectoryCaller) Profile(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _UserDirectory.contract.Call(opts, out, "profile")
	return *ret0, err
}

// Profile is a free data retrieval call binding the contract method 0xab60636c.
//
// Solidity: function profile() constant returns(string)
func (_UserDirectory *UserDirectorySession) Profile() (string, error) {
	return _UserDirectory.Contract.Profile(&_UserDirectory.CallOpts)
}

// Profile is a free data retrieval call binding the contract method 0xab60636c.
//
// Solidity: function profile() constant returns(string)
func (_UserDirectory *UserDirectoryCallerSession) Profile() (string, error) {
	return _UserDirectory.Contract.Profile(&_UserDirectory.CallOpts)
}

// Roles is a free data retrieval call binding the contract method 0x392f5f64.
//
// Solidity: function roles() constant returns(address)
func (_UserDirectory *UserDirectoryCaller) Roles(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UserDirectory.contract.Call(opts, out, "roles")
	return *ret0, err
}

// Roles is a free data retrieval call binding the contract method 0x392f5f64.
//
// Solidity: function roles() constant returns(address)
func (_UserDirectory *UserDirectorySession) Roles() (common.Address, error) {
	return _UserDirectory.Contract.Roles(&_UserDirectory.CallOpts)
}

// Roles is a free data retrieval call binding the contract method 0x392f5f64.
//
// Solidity: function roles() constant returns(address)
func (_UserDirectory *UserDirectoryCallerSession) Roles() (common.Address, error) {
	return _UserDirectory.Contract.Roles(&_UserDirectory.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() constant returns(address)
func (_UserDirectory *UserDirectoryCaller) Root(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UserDirectory.contract.Call(opts, out, "root")
	return *ret0, err
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() constant returns(address)
func (_UserDirectory *UserDirectorySession) Root() (common.Address, error) {
	return _UserDirectory.Contract.Root(&_UserDirectory.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() constant returns(address)
func (_UserDirectory *UserDirectoryCallerSession) Root() (common.Address, error) {
	return _UserDirectory.Contract.Root(&_UserDirectory.CallOpts)
}

// SenderHasRole is a free data retrieval call binding the contract method 0x1ca03b8e.
//
// Solidity: function senderHasRole(roleName string) constant returns(bool)
func (_UserDirectory *UserDirectoryCaller) SenderHasRole(opts *bind.CallOpts, roleName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _UserDirectory.contract.Call(opts, out, "senderHasRole", roleName)
	return *ret0, err
}

// SenderHasRole is a free data retrieval call binding the contract method 0x1ca03b8e.
//
// Solidity: function senderHasRole(roleName string) constant returns(bool)
func (_UserDirectory *UserDirectorySession) SenderHasRole(roleName string) (bool, error) {
	return _UserDirectory.Contract.SenderHasRole(&_UserDirectory.CallOpts, roleName)
}

// SenderHasRole is a free data retrieval call binding the contract method 0x1ca03b8e.
//
// Solidity: function senderHasRole(roleName string) constant returns(bool)
func (_UserDirectory *UserDirectoryCallerSession) SenderHasRole(roleName string) (bool, error) {
	return _UserDirectory.Contract.SenderHasRole(&_UserDirectory.CallOpts, roleName)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_UserDirectory *UserDirectoryCaller) Stopped(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _UserDirectory.contract.Call(opts, out, "stopped")
	return *ret0, err
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_UserDirectory *UserDirectorySession) Stopped() (bool, error) {
	return _UserDirectory.Contract.Stopped(&_UserDirectory.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_UserDirectory *UserDirectoryCallerSession) Stopped() (bool, error) {
	return _UserDirectory.Contract.Stopped(&_UserDirectory.CallOpts)
}

// WalletCount is a free data retrieval call binding the contract method 0x29b57c69.
//
// Solidity: function walletCount() constant returns(uint32)
func (_UserDirectory *UserDirectoryCaller) WalletCount(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _UserDirectory.contract.Call(opts, out, "walletCount")
	return *ret0, err
}

// WalletCount is a free data retrieval call binding the contract method 0x29b57c69.
//
// Solidity: function walletCount() constant returns(uint32)
func (_UserDirectory *UserDirectorySession) WalletCount() (uint32, error) {
	return _UserDirectory.Contract.WalletCount(&_UserDirectory.CallOpts)
}

// WalletCount is a free data retrieval call binding the contract method 0x29b57c69.
//
// Solidity: function walletCount() constant returns(uint32)
func (_UserDirectory *UserDirectoryCallerSession) WalletCount() (uint32, error) {
	return _UserDirectory.Contract.WalletCount(&_UserDirectory.CallOpts)
}

// Wallets is a free data retrieval call binding the contract method 0x7ad71f72.
//
// Solidity: function wallets( uint256) constant returns(address)
func (_UserDirectory *UserDirectoryCaller) Wallets(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UserDirectory.contract.Call(opts, out, "wallets", arg0)
	return *ret0, err
}

// Wallets is a free data retrieval call binding the contract method 0x7ad71f72.
//
// Solidity: function wallets( uint256) constant returns(address)
func (_UserDirectory *UserDirectorySession) Wallets(arg0 *big.Int) (common.Address, error) {
	return _UserDirectory.Contract.Wallets(&_UserDirectory.CallOpts, arg0)
}

// Wallets is a free data retrieval call binding the contract method 0x7ad71f72.
//
// Solidity: function wallets( uint256) constant returns(address)
func (_UserDirectory *UserDirectoryCallerSession) Wallets(arg0 *big.Int) (common.Address, error) {
	return _UserDirectory.Contract.Wallets(&_UserDirectory.CallOpts, arg0)
}

// AddWallet is a paid mutator transaction binding the contract method 0xfbd51eee.
//
// Solidity: function addWallet() returns()
func (_UserDirectory *UserDirectoryTransactor) AddWallet(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserDirectory.contract.Transact(opts, "addWallet")
}

// AddWallet is a paid mutator transaction binding the contract method 0xfbd51eee.
//
// Solidity: function addWallet() returns()
func (_UserDirectory *UserDirectorySession) AddWallet() (*types.Transaction, error) {
	return _UserDirectory.Contract.AddWallet(&_UserDirectory.TransactOpts)
}

// AddWallet is a paid mutator transaction binding the contract method 0xfbd51eee.
//
// Solidity: function addWallet() returns()
func (_UserDirectory *UserDirectoryTransactorSession) AddWallet() (*types.Transaction, error) {
	return _UserDirectory.Contract.AddWallet(&_UserDirectory.TransactOpts)
}

// RemoveWallet is a paid mutator transaction binding the contract method 0xa75fe8e1.
//
// Solidity: function removeWallet(wallet_ address) returns()
func (_UserDirectory *UserDirectoryTransactor) RemoveWallet(opts *bind.TransactOpts, wallet_ common.Address) (*types.Transaction, error) {
	return _UserDirectory.contract.Transact(opts, "removeWallet", wallet_)
}

// RemoveWallet is a paid mutator transaction binding the contract method 0xa75fe8e1.
//
// Solidity: function removeWallet(wallet_ address) returns()
func (_UserDirectory *UserDirectorySession) RemoveWallet(wallet_ common.Address) (*types.Transaction, error) {
	return _UserDirectory.Contract.RemoveWallet(&_UserDirectory.TransactOpts, wallet_)
}

// RemoveWallet is a paid mutator transaction binding the contract method 0xa75fe8e1.
//
// Solidity: function removeWallet(wallet_ address) returns()
func (_UserDirectory *UserDirectoryTransactorSession) RemoveWallet(wallet_ common.Address) (*types.Transaction, error) {
	return _UserDirectory.Contract.RemoveWallet(&_UserDirectory.TransactOpts, wallet_)
}

// Restart is a paid mutator transaction binding the contract method 0x1ef3755d.
//
// Solidity: function restart() returns()
func (_UserDirectory *UserDirectoryTransactor) Restart(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserDirectory.contract.Transact(opts, "restart")
}

// Restart is a paid mutator transaction binding the contract method 0x1ef3755d.
//
// Solidity: function restart() returns()
func (_UserDirectory *UserDirectorySession) Restart() (*types.Transaction, error) {
	return _UserDirectory.Contract.Restart(&_UserDirectory.TransactOpts)
}

// Restart is a paid mutator transaction binding the contract method 0x1ef3755d.
//
// Solidity: function restart() returns()
func (_UserDirectory *UserDirectoryTransactorSession) Restart() (*types.Transaction, error) {
	return _UserDirectory.Contract.Restart(&_UserDirectory.TransactOpts)
}

// SetKYC is a paid mutator transaction binding the contract method 0xbb05c30e.
//
// Solidity: function setKYC(kyc_ bool) returns()
func (_UserDirectory *UserDirectoryTransactor) SetKYC(opts *bind.TransactOpts, kyc_ bool) (*types.Transaction, error) {
	return _UserDirectory.contract.Transact(opts, "setKYC", kyc_)
}

// SetKYC is a paid mutator transaction binding the contract method 0xbb05c30e.
//
// Solidity: function setKYC(kyc_ bool) returns()
func (_UserDirectory *UserDirectorySession) SetKYC(kyc_ bool) (*types.Transaction, error) {
	return _UserDirectory.Contract.SetKYC(&_UserDirectory.TransactOpts, kyc_)
}

// SetKYC is a paid mutator transaction binding the contract method 0xbb05c30e.
//
// Solidity: function setKYC(kyc_ bool) returns()
func (_UserDirectory *UserDirectoryTransactorSession) SetKYC(kyc_ bool) (*types.Transaction, error) {
	return _UserDirectory.Contract.SetKYC(&_UserDirectory.TransactOpts, kyc_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_UserDirectory *UserDirectoryTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _UserDirectory.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_UserDirectory *UserDirectorySession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _UserDirectory.Contract.SetOwner(&_UserDirectory.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_UserDirectory *UserDirectoryTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _UserDirectory.Contract.SetOwner(&_UserDirectory.TransactOpts, owner_)
}

// SetRolesContract is a paid mutator transaction binding the contract method 0xafa202ac.
//
// Solidity: function setRolesContract(roles_ address) returns()
func (_UserDirectory *UserDirectoryTransactor) SetRolesContract(opts *bind.TransactOpts, roles_ common.Address) (*types.Transaction, error) {
	return _UserDirectory.contract.Transact(opts, "setRolesContract", roles_)
}

// SetRolesContract is a paid mutator transaction binding the contract method 0xafa202ac.
//
// Solidity: function setRolesContract(roles_ address) returns()
func (_UserDirectory *UserDirectorySession) SetRolesContract(roles_ common.Address) (*types.Transaction, error) {
	return _UserDirectory.Contract.SetRolesContract(&_UserDirectory.TransactOpts, roles_)
}

// SetRolesContract is a paid mutator transaction binding the contract method 0xafa202ac.
//
// Solidity: function setRolesContract(roles_ address) returns()
func (_UserDirectory *UserDirectoryTransactorSession) SetRolesContract(roles_ common.Address) (*types.Transaction, error) {
	return _UserDirectory.Contract.SetRolesContract(&_UserDirectory.TransactOpts, roles_)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_UserDirectory *UserDirectoryTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserDirectory.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_UserDirectory *UserDirectorySession) Stop() (*types.Transaction, error) {
	return _UserDirectory.Contract.Stop(&_UserDirectory.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_UserDirectory *UserDirectoryTransactorSession) Stop() (*types.Transaction, error) {
	return _UserDirectory.Contract.Stop(&_UserDirectory.TransactOpts)
}

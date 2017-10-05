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

// WalletABI is the input ABI used to generate the binding from.
const WalletABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"senderHasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"restart\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"roles\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ethBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"listVaults\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"addAsset\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"cleanStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"addVault\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vaults\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"assetsLen\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vaultCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"remove\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"roles_\",\"type\":\"address\"}],\"name\":\"setRolesContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"directory\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"removeVault\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assets\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"}],\"name\":\"rmAsset\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"hasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transferEth\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hasFunds\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"directory_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"LogVaultAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"LogVaultRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"AssetAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"}],\"name\":\"AssetRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"}]"

// WalletBin is the compiled bytecode used for deploying new contracts.
const WalletBin = `undefined`

// DeployWallet deploys a new Ethereum contract, binding an instance of Wallet to it.
func DeployWallet(auth *bind.TransactOpts, backend bind.ContractBackend, directory_ common.Address) (common.Address, *types.Transaction, *Wallet, error) {
	parsed, err := abi.JSON(strings.NewReader(WalletABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WalletBin), backend, directory_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Wallet{WalletCaller: WalletCaller{contract: contract}, WalletTransactor: WalletTransactor{contract: contract}}, nil
}

// Wallet is an auto generated Go binding around an Ethereum contract.
type Wallet struct {
	WalletCaller     // Read-only binding to the contract
	WalletTransactor // Write-only binding to the contract
}

// WalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletSession struct {
	Contract     *Wallet           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletCallerSession struct {
	Contract *WalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletTransactorSession struct {
	Contract     *WalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletRaw struct {
	Contract *Wallet // Generic contract binding to access the raw methods on
}

// WalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletCallerRaw struct {
	Contract *WalletCaller // Generic read-only contract binding to access the raw methods on
}

// WalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletTransactorRaw struct {
	Contract *WalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWallet creates a new instance of Wallet, bound to a specific deployed contract.
func NewWallet(address common.Address, backend bind.ContractBackend) (*Wallet, error) {
	contract, err := bindWallet(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Wallet{WalletCaller: WalletCaller{contract: contract}, WalletTransactor: WalletTransactor{contract: contract}}, nil
}

// NewWalletCaller creates a new read-only instance of Wallet, bound to a specific deployed contract.
func NewWalletCaller(address common.Address, caller bind.ContractCaller) (*WalletCaller, error) {
	contract, err := bindWallet(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &WalletCaller{contract: contract}, nil
}

// NewWalletTransactor creates a new write-only instance of Wallet, bound to a specific deployed contract.
func NewWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletTransactor, error) {
	contract, err := bindWallet(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &WalletTransactor{contract: contract}, nil
}

// bindWallet binds a generic wrapper to an already deployed contract.
func bindWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wallet *WalletRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Wallet.Contract.WalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wallet *WalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.Contract.WalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wallet *WalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wallet.Contract.WalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wallet *WalletCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Wallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wallet *WalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wallet *WalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wallet.Contract.contract.Transact(opts, method, params...)
}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets( uint256) constant returns(address)
func (_Wallet *WalletCaller) Assets(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "assets", arg0)
	return *ret0, err
}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets( uint256) constant returns(address)
func (_Wallet *WalletSession) Assets(arg0 *big.Int) (common.Address, error) {
	return _Wallet.Contract.Assets(&_Wallet.CallOpts, arg0)
}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets( uint256) constant returns(address)
func (_Wallet *WalletCallerSession) Assets(arg0 *big.Int) (common.Address, error) {
	return _Wallet.Contract.Assets(&_Wallet.CallOpts, arg0)
}

// AssetsLen is a free data retrieval call binding the contract method 0xa37685e5.
//
// Solidity: function assetsLen() constant returns(uint256)
func (_Wallet *WalletCaller) AssetsLen(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "assetsLen")
	return *ret0, err
}

// AssetsLen is a free data retrieval call binding the contract method 0xa37685e5.
//
// Solidity: function assetsLen() constant returns(uint256)
func (_Wallet *WalletSession) AssetsLen() (*big.Int, error) {
	return _Wallet.Contract.AssetsLen(&_Wallet.CallOpts)
}

// AssetsLen is a free data retrieval call binding the contract method 0xa37685e5.
//
// Solidity: function assetsLen() constant returns(uint256)
func (_Wallet *WalletCallerSession) AssetsLen() (*big.Int, error) {
	return _Wallet.Contract.AssetsLen(&_Wallet.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(token address) constant returns(uint256)
func (_Wallet *WalletCaller) BalanceOf(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "balanceOf", token)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(token address) constant returns(uint256)
func (_Wallet *WalletSession) BalanceOf(token common.Address) (*big.Int, error) {
	return _Wallet.Contract.BalanceOf(&_Wallet.CallOpts, token)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(token address) constant returns(uint256)
func (_Wallet *WalletCallerSession) BalanceOf(token common.Address) (*big.Int, error) {
	return _Wallet.Contract.BalanceOf(&_Wallet.CallOpts, token)
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(address[], uint256[])
func (_Wallet *WalletCaller) Balances(opts *bind.CallOpts) ([]common.Address, []*big.Int, error) {
	var (
		ret0 = new([]common.Address)
		ret1 = new([]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Wallet.contract.Call(opts, out, "balances")
	return *ret0, *ret1, err
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(address[], uint256[])
func (_Wallet *WalletSession) Balances() ([]common.Address, []*big.Int, error) {
	return _Wallet.Contract.Balances(&_Wallet.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(address[], uint256[])
func (_Wallet *WalletCallerSession) Balances() ([]common.Address, []*big.Int, error) {
	return _Wallet.Contract.Balances(&_Wallet.CallOpts)
}

// ContractHash is a free data retrieval call binding the contract method 0x904c6094.
//
// Solidity: function contractHash() constant returns(bytes32)
func (_Wallet *WalletCaller) ContractHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "contractHash")
	return *ret0, err
}

// ContractHash is a free data retrieval call binding the contract method 0x904c6094.
//
// Solidity: function contractHash() constant returns(bytes32)
func (_Wallet *WalletSession) ContractHash() ([32]byte, error) {
	return _Wallet.Contract.ContractHash(&_Wallet.CallOpts)
}

// ContractHash is a free data retrieval call binding the contract method 0x904c6094.
//
// Solidity: function contractHash() constant returns(bytes32)
func (_Wallet *WalletCallerSession) ContractHash() ([32]byte, error) {
	return _Wallet.Contract.ContractHash(&_Wallet.CallOpts)
}

// Directory is a free data retrieval call binding the contract method 0xc41c2f24.
//
// Solidity: function directory() constant returns(address)
func (_Wallet *WalletCaller) Directory(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "directory")
	return *ret0, err
}

// Directory is a free data retrieval call binding the contract method 0xc41c2f24.
//
// Solidity: function directory() constant returns(address)
func (_Wallet *WalletSession) Directory() (common.Address, error) {
	return _Wallet.Contract.Directory(&_Wallet.CallOpts)
}

// Directory is a free data retrieval call binding the contract method 0xc41c2f24.
//
// Solidity: function directory() constant returns(address)
func (_Wallet *WalletCallerSession) Directory() (common.Address, error) {
	return _Wallet.Contract.Directory(&_Wallet.CallOpts)
}

// EthBalance is a free data retrieval call binding the contract method 0x4e6630b0.
//
// Solidity: function ethBalance() constant returns(uint256)
func (_Wallet *WalletCaller) EthBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "ethBalance")
	return *ret0, err
}

// EthBalance is a free data retrieval call binding the contract method 0x4e6630b0.
//
// Solidity: function ethBalance() constant returns(uint256)
func (_Wallet *WalletSession) EthBalance() (*big.Int, error) {
	return _Wallet.Contract.EthBalance(&_Wallet.CallOpts)
}

// EthBalance is a free data retrieval call binding the contract method 0x4e6630b0.
//
// Solidity: function ethBalance() constant returns(uint256)
func (_Wallet *WalletCallerSession) EthBalance() (*big.Int, error) {
	return _Wallet.Contract.EthBalance(&_Wallet.CallOpts)
}

// HasFunds is a free data retrieval call binding the contract method 0xf2cd59f6.
//
// Solidity: function hasFunds() constant returns(bool)
func (_Wallet *WalletCaller) HasFunds(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "hasFunds")
	return *ret0, err
}

// HasFunds is a free data retrieval call binding the contract method 0xf2cd59f6.
//
// Solidity: function hasFunds() constant returns(bool)
func (_Wallet *WalletSession) HasFunds() (bool, error) {
	return _Wallet.Contract.HasFunds(&_Wallet.CallOpts)
}

// HasFunds is a free data retrieval call binding the contract method 0xf2cd59f6.
//
// Solidity: function hasFunds() constant returns(bool)
func (_Wallet *WalletCallerSession) HasFunds() (bool, error) {
	return _Wallet.Contract.HasFunds(&_Wallet.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0xe3c33a9b.
//
// Solidity: function hasRole(roleName string) constant returns(bool)
func (_Wallet *WalletCaller) HasRole(opts *bind.CallOpts, roleName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "hasRole", roleName)
	return *ret0, err
}

// HasRole is a free data retrieval call binding the contract method 0xe3c33a9b.
//
// Solidity: function hasRole(roleName string) constant returns(bool)
func (_Wallet *WalletSession) HasRole(roleName string) (bool, error) {
	return _Wallet.Contract.HasRole(&_Wallet.CallOpts, roleName)
}

// HasRole is a free data retrieval call binding the contract method 0xe3c33a9b.
//
// Solidity: function hasRole(roleName string) constant returns(bool)
func (_Wallet *WalletCallerSession) HasRole(roleName string) (bool, error) {
	return _Wallet.Contract.HasRole(&_Wallet.CallOpts, roleName)
}

// ListVaults is a free data retrieval call binding the contract method 0x50cc258e.
//
// Solidity: function listVaults() constant returns(address[])
func (_Wallet *WalletCaller) ListVaults(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "listVaults")
	return *ret0, err
}

// ListVaults is a free data retrieval call binding the contract method 0x50cc258e.
//
// Solidity: function listVaults() constant returns(address[])
func (_Wallet *WalletSession) ListVaults() ([]common.Address, error) {
	return _Wallet.Contract.ListVaults(&_Wallet.CallOpts)
}

// ListVaults is a free data retrieval call binding the contract method 0x50cc258e.
//
// Solidity: function listVaults() constant returns(address[])
func (_Wallet *WalletCallerSession) ListVaults() ([]common.Address, error) {
	return _Wallet.Contract.ListVaults(&_Wallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Wallet *WalletCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Wallet *WalletSession) Owner() (common.Address, error) {
	return _Wallet.Contract.Owner(&_Wallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Wallet *WalletCallerSession) Owner() (common.Address, error) {
	return _Wallet.Contract.Owner(&_Wallet.CallOpts)
}

// Roles is a free data retrieval call binding the contract method 0x392f5f64.
//
// Solidity: function roles() constant returns(address)
func (_Wallet *WalletCaller) Roles(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "roles")
	return *ret0, err
}

// Roles is a free data retrieval call binding the contract method 0x392f5f64.
//
// Solidity: function roles() constant returns(address)
func (_Wallet *WalletSession) Roles() (common.Address, error) {
	return _Wallet.Contract.Roles(&_Wallet.CallOpts)
}

// Roles is a free data retrieval call binding the contract method 0x392f5f64.
//
// Solidity: function roles() constant returns(address)
func (_Wallet *WalletCallerSession) Roles() (common.Address, error) {
	return _Wallet.Contract.Roles(&_Wallet.CallOpts)
}

// SenderHasRole is a free data retrieval call binding the contract method 0x1ca03b8e.
//
// Solidity: function senderHasRole(roleName string) constant returns(bool)
func (_Wallet *WalletCaller) SenderHasRole(opts *bind.CallOpts, roleName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "senderHasRole", roleName)
	return *ret0, err
}

// SenderHasRole is a free data retrieval call binding the contract method 0x1ca03b8e.
//
// Solidity: function senderHasRole(roleName string) constant returns(bool)
func (_Wallet *WalletSession) SenderHasRole(roleName string) (bool, error) {
	return _Wallet.Contract.SenderHasRole(&_Wallet.CallOpts, roleName)
}

// SenderHasRole is a free data retrieval call binding the contract method 0x1ca03b8e.
//
// Solidity: function senderHasRole(roleName string) constant returns(bool)
func (_Wallet *WalletCallerSession) SenderHasRole(roleName string) (bool, error) {
	return _Wallet.Contract.SenderHasRole(&_Wallet.CallOpts, roleName)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_Wallet *WalletCaller) Stopped(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "stopped")
	return *ret0, err
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_Wallet *WalletSession) Stopped() (bool, error) {
	return _Wallet.Contract.Stopped(&_Wallet.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_Wallet *WalletCallerSession) Stopped() (bool, error) {
	return _Wallet.Contract.Stopped(&_Wallet.CallOpts)
}

// VaultCount is a free data retrieval call binding the contract method 0xa7c6a100.
//
// Solidity: function vaultCount() constant returns(uint32)
func (_Wallet *WalletCaller) VaultCount(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "vaultCount")
	return *ret0, err
}

// VaultCount is a free data retrieval call binding the contract method 0xa7c6a100.
//
// Solidity: function vaultCount() constant returns(uint32)
func (_Wallet *WalletSession) VaultCount() (uint32, error) {
	return _Wallet.Contract.VaultCount(&_Wallet.CallOpts)
}

// VaultCount is a free data retrieval call binding the contract method 0xa7c6a100.
//
// Solidity: function vaultCount() constant returns(uint32)
func (_Wallet *WalletCallerSession) VaultCount() (uint32, error) {
	return _Wallet.Contract.VaultCount(&_Wallet.CallOpts)
}

// Vaults is a free data retrieval call binding the contract method 0x8c64ea4a.
//
// Solidity: function vaults( uint256) constant returns(address)
func (_Wallet *WalletCaller) Vaults(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "vaults", arg0)
	return *ret0, err
}

// Vaults is a free data retrieval call binding the contract method 0x8c64ea4a.
//
// Solidity: function vaults( uint256) constant returns(address)
func (_Wallet *WalletSession) Vaults(arg0 *big.Int) (common.Address, error) {
	return _Wallet.Contract.Vaults(&_Wallet.CallOpts, arg0)
}

// Vaults is a free data retrieval call binding the contract method 0x8c64ea4a.
//
// Solidity: function vaults( uint256) constant returns(address)
func (_Wallet *WalletCallerSession) Vaults(arg0 *big.Int) (common.Address, error) {
	return _Wallet.Contract.Vaults(&_Wallet.CallOpts, arg0)
}

// AddAsset is a paid mutator transaction binding the contract method 0x712e6279.
//
// Solidity: function addAsset(token address, src address, wad uint128) returns()
func (_Wallet *WalletTransactor) AddAsset(opts *bind.TransactOpts, token common.Address, src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "addAsset", token, src, wad)
}

// AddAsset is a paid mutator transaction binding the contract method 0x712e6279.
//
// Solidity: function addAsset(token address, src address, wad uint128) returns()
func (_Wallet *WalletSession) AddAsset(token common.Address, src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.AddAsset(&_Wallet.TransactOpts, token, src, wad)
}

// AddAsset is a paid mutator transaction binding the contract method 0x712e6279.
//
// Solidity: function addAsset(token address, src address, wad uint128) returns()
func (_Wallet *WalletTransactorSession) AddAsset(token common.Address, src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.AddAsset(&_Wallet.TransactOpts, token, src, wad)
}

// AddVault is a paid mutator transaction binding the contract method 0x8c533411.
//
// Solidity: function addVault() returns()
func (_Wallet *WalletTransactor) AddVault(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "addVault")
}

// AddVault is a paid mutator transaction binding the contract method 0x8c533411.
//
// Solidity: function addVault() returns()
func (_Wallet *WalletSession) AddVault() (*types.Transaction, error) {
	return _Wallet.Contract.AddVault(&_Wallet.TransactOpts)
}

// AddVault is a paid mutator transaction binding the contract method 0x8c533411.
//
// Solidity: function addVault() returns()
func (_Wallet *WalletTransactorSession) AddVault() (*types.Transaction, error) {
	return _Wallet.Contract.AddVault(&_Wallet.TransactOpts)
}

// CleanStorage is a paid mutator transaction binding the contract method 0x7e84d36b.
//
// Solidity: function cleanStorage() returns()
func (_Wallet *WalletTransactor) CleanStorage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "cleanStorage")
}

// CleanStorage is a paid mutator transaction binding the contract method 0x7e84d36b.
//
// Solidity: function cleanStorage() returns()
func (_Wallet *WalletSession) CleanStorage() (*types.Transaction, error) {
	return _Wallet.Contract.CleanStorage(&_Wallet.TransactOpts)
}

// CleanStorage is a paid mutator transaction binding the contract method 0x7e84d36b.
//
// Solidity: function cleanStorage() returns()
func (_Wallet *WalletTransactorSession) CleanStorage() (*types.Transaction, error) {
	return _Wallet.Contract.CleanStorage(&_Wallet.TransactOpts)
}

// Remove is a paid mutator transaction binding the contract method 0xa7f43779.
//
// Solidity: function remove() returns()
func (_Wallet *WalletTransactor) Remove(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "remove")
}

// Remove is a paid mutator transaction binding the contract method 0xa7f43779.
//
// Solidity: function remove() returns()
func (_Wallet *WalletSession) Remove() (*types.Transaction, error) {
	return _Wallet.Contract.Remove(&_Wallet.TransactOpts)
}

// Remove is a paid mutator transaction binding the contract method 0xa7f43779.
//
// Solidity: function remove() returns()
func (_Wallet *WalletTransactorSession) Remove() (*types.Transaction, error) {
	return _Wallet.Contract.Remove(&_Wallet.TransactOpts)
}

// RemoveVault is a paid mutator transaction binding the contract method 0xceb68c23.
//
// Solidity: function removeVault(vault address) returns()
func (_Wallet *WalletTransactor) RemoveVault(opts *bind.TransactOpts, vault common.Address) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "removeVault", vault)
}

// RemoveVault is a paid mutator transaction binding the contract method 0xceb68c23.
//
// Solidity: function removeVault(vault address) returns()
func (_Wallet *WalletSession) RemoveVault(vault common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.RemoveVault(&_Wallet.TransactOpts, vault)
}

// RemoveVault is a paid mutator transaction binding the contract method 0xceb68c23.
//
// Solidity: function removeVault(vault address) returns()
func (_Wallet *WalletTransactorSession) RemoveVault(vault common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.RemoveVault(&_Wallet.TransactOpts, vault)
}

// Restart is a paid mutator transaction binding the contract method 0x1ef3755d.
//
// Solidity: function restart() returns()
func (_Wallet *WalletTransactor) Restart(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "restart")
}

// Restart is a paid mutator transaction binding the contract method 0x1ef3755d.
//
// Solidity: function restart() returns()
func (_Wallet *WalletSession) Restart() (*types.Transaction, error) {
	return _Wallet.Contract.Restart(&_Wallet.TransactOpts)
}

// Restart is a paid mutator transaction binding the contract method 0x1ef3755d.
//
// Solidity: function restart() returns()
func (_Wallet *WalletTransactorSession) Restart() (*types.Transaction, error) {
	return _Wallet.Contract.Restart(&_Wallet.TransactOpts)
}

// RmAsset is a paid mutator transaction binding the contract method 0xd1da3e7b.
//
// Solidity: function rmAsset(token address, dst address) returns(bool)
func (_Wallet *WalletTransactor) RmAsset(opts *bind.TransactOpts, token common.Address, dst common.Address) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "rmAsset", token, dst)
}

// RmAsset is a paid mutator transaction binding the contract method 0xd1da3e7b.
//
// Solidity: function rmAsset(token address, dst address) returns(bool)
func (_Wallet *WalletSession) RmAsset(token common.Address, dst common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.RmAsset(&_Wallet.TransactOpts, token, dst)
}

// RmAsset is a paid mutator transaction binding the contract method 0xd1da3e7b.
//
// Solidity: function rmAsset(token address, dst address) returns(bool)
func (_Wallet *WalletTransactorSession) RmAsset(token common.Address, dst common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.RmAsset(&_Wallet.TransactOpts, token, dst)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_Wallet *WalletTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_Wallet *WalletSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.SetOwner(&_Wallet.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_Wallet *WalletTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.SetOwner(&_Wallet.TransactOpts, owner_)
}

// SetRolesContract is a paid mutator transaction binding the contract method 0xafa202ac.
//
// Solidity: function setRolesContract(roles_ address) returns()
func (_Wallet *WalletTransactor) SetRolesContract(opts *bind.TransactOpts, roles_ common.Address) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "setRolesContract", roles_)
}

// SetRolesContract is a paid mutator transaction binding the contract method 0xafa202ac.
//
// Solidity: function setRolesContract(roles_ address) returns()
func (_Wallet *WalletSession) SetRolesContract(roles_ common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.SetRolesContract(&_Wallet.TransactOpts, roles_)
}

// SetRolesContract is a paid mutator transaction binding the contract method 0xafa202ac.
//
// Solidity: function setRolesContract(roles_ address) returns()
func (_Wallet *WalletTransactorSession) SetRolesContract(roles_ common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.SetRolesContract(&_Wallet.TransactOpts, roles_)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Wallet *WalletTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Wallet *WalletSession) Stop() (*types.Transaction, error) {
	return _Wallet.Contract.Stop(&_Wallet.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Wallet *WalletTransactorSession) Stop() (*types.Transaction, error) {
	return _Wallet.Contract.Stop(&_Wallet.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0x56e596a8.
//
// Solidity: function transfer(token address, dst address, wad uint128) returns()
func (_Wallet *WalletTransactor) Transfer(opts *bind.TransactOpts, token common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "transfer", token, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0x56e596a8.
//
// Solidity: function transfer(token address, dst address, wad uint128) returns()
func (_Wallet *WalletSession) Transfer(token common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.Transfer(&_Wallet.TransactOpts, token, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0x56e596a8.
//
// Solidity: function transfer(token address, dst address, wad uint128) returns()
func (_Wallet *WalletTransactorSession) Transfer(token common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.Transfer(&_Wallet.TransactOpts, token, dst, wad)
}

// TransferEth is a paid mutator transaction binding the contract method 0xe9bb84c2.
//
// Solidity: function transferEth(dst address, wad uint256) returns()
func (_Wallet *WalletTransactor) TransferEth(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "transferEth", dst, wad)
}

// TransferEth is a paid mutator transaction binding the contract method 0xe9bb84c2.
//
// Solidity: function transferEth(dst address, wad uint256) returns()
func (_Wallet *WalletSession) TransferEth(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.TransferEth(&_Wallet.TransactOpts, dst, wad)
}

// TransferEth is a paid mutator transaction binding the contract method 0xe9bb84c2.
//
// Solidity: function transferEth(dst address, wad uint256) returns()
func (_Wallet *WalletTransactorSession) TransferEth(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.TransferEth(&_Wallet.TransactOpts, dst, wad)
}

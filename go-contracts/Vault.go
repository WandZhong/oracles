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

// VaultABI is the input ABI used to generate the binding from.
const VaultABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"brgAmount\",\"type\":\"uint128\"},{\"name\":\"uouIndex\",\"type\":\"uint256\"}],\"name\":\"repayUou\",\"outputs\":[{\"name\":\"\",\"type\":\"uint128\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"senderHasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"restart\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uouIndex\",\"type\":\"uint256\"}],\"name\":\"rejectUouRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"brgBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"amountDue\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"roles\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ethBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"uouCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"takeOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isEmpty\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"addAsset\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"cleanStorage\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uous\",\"outputs\":[{\"name\":\"initialAmount\",\"type\":\"uint128\"},{\"name\":\"repaidAmount\",\"type\":\"uint128\"},{\"name\":\"fee\",\"type\":\"uint128\"},{\"name\":\"time\",\"type\":\"uint256\"},{\"name\":\"decision\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"assetsLen\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"brgAmount\",\"type\":\"uint128\"}],\"name\":\"requestUou\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"remove\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"roles_\",\"type\":\"address\"}],\"name\":\"setRolesContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assets\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"}],\"name\":\"rmAsset\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uouIndex\",\"type\":\"uint256\"}],\"name\":\"acceptUouRequest\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"hasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transferEth\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hasFunds\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"wallet_\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"brgAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"uouIndex\",\"type\":\"uint256\"}],\"name\":\"UouRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"brgAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"uouIndex\",\"type\":\"uint256\"}],\"name\":\"UouRequestApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"brgAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"uouIndex\",\"type\":\"uint256\"}],\"name\":\"UouRequestDeclined\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"AssetAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"}],\"name\":\"AssetRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"}]"

// VaultBin is the compiled bytecode used for deploying new contracts.
const VaultBin = `undefined`

// DeployVault deploys a new Ethereum contract, binding an instance of Vault to it.
func DeployVault(auth *bind.TransactOpts, backend bind.ContractBackend, wallet_ common.Address) (common.Address, *types.Transaction, *Vault, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VaultBin), backend, wallet_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Vault{VaultCaller: VaultCaller{contract: contract}, VaultTransactor: VaultTransactor{contract: contract}}, nil
}

// Vault is an auto generated Go binding around an Ethereum contract.
type Vault struct {
	VaultCaller     // Read-only binding to the contract
	VaultTransactor // Write-only binding to the contract
}

// VaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultSession struct {
	Contract     *Vault            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultCallerSession struct {
	Contract *VaultCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultTransactorSession struct {
	Contract     *VaultTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultRaw struct {
	Contract *Vault // Generic contract binding to access the raw methods on
}

// VaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultCallerRaw struct {
	Contract *VaultCaller // Generic read-only contract binding to access the raw methods on
}

// VaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultTransactorRaw struct {
	Contract *VaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVault creates a new instance of Vault, bound to a specific deployed contract.
func NewVault(address common.Address, backend bind.ContractBackend) (*Vault, error) {
	contract, err := bindVault(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vault{VaultCaller: VaultCaller{contract: contract}, VaultTransactor: VaultTransactor{contract: contract}}, nil
}

// NewVaultCaller creates a new read-only instance of Vault, bound to a specific deployed contract.
func NewVaultCaller(address common.Address, caller bind.ContractCaller) (*VaultCaller, error) {
	contract, err := bindVault(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &VaultCaller{contract: contract}, nil
}

// NewVaultTransactor creates a new write-only instance of Vault, bound to a specific deployed contract.
func NewVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultTransactor, error) {
	contract, err := bindVault(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &VaultTransactor{contract: contract}, nil
}

// bindVault binds a generic wrapper to an already deployed contract.
func bindVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vault *VaultRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Vault.Contract.VaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vault *VaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.Contract.VaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vault *VaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vault.Contract.VaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vault *VaultCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Vault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vault *VaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vault *VaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vault.Contract.contract.Transact(opts, method, params...)
}

// AmountDue is a free data retrieval call binding the contract method 0x2f195680.
//
// Solidity: function amountDue() constant returns(uint256)
func (_Vault *VaultCaller) AmountDue(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "amountDue")
	return *ret0, err
}

// AmountDue is a free data retrieval call binding the contract method 0x2f195680.
//
// Solidity: function amountDue() constant returns(uint256)
func (_Vault *VaultSession) AmountDue() (*big.Int, error) {
	return _Vault.Contract.AmountDue(&_Vault.CallOpts)
}

// AmountDue is a free data retrieval call binding the contract method 0x2f195680.
//
// Solidity: function amountDue() constant returns(uint256)
func (_Vault *VaultCallerSession) AmountDue() (*big.Int, error) {
	return _Vault.Contract.AmountDue(&_Vault.CallOpts)
}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets( uint256) constant returns(address)
func (_Vault *VaultCaller) Assets(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "assets", arg0)
	return *ret0, err
}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets( uint256) constant returns(address)
func (_Vault *VaultSession) Assets(arg0 *big.Int) (common.Address, error) {
	return _Vault.Contract.Assets(&_Vault.CallOpts, arg0)
}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets( uint256) constant returns(address)
func (_Vault *VaultCallerSession) Assets(arg0 *big.Int) (common.Address, error) {
	return _Vault.Contract.Assets(&_Vault.CallOpts, arg0)
}

// AssetsLen is a free data retrieval call binding the contract method 0xa37685e5.
//
// Solidity: function assetsLen() constant returns(uint256)
func (_Vault *VaultCaller) AssetsLen(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "assetsLen")
	return *ret0, err
}

// AssetsLen is a free data retrieval call binding the contract method 0xa37685e5.
//
// Solidity: function assetsLen() constant returns(uint256)
func (_Vault *VaultSession) AssetsLen() (*big.Int, error) {
	return _Vault.Contract.AssetsLen(&_Vault.CallOpts)
}

// AssetsLen is a free data retrieval call binding the contract method 0xa37685e5.
//
// Solidity: function assetsLen() constant returns(uint256)
func (_Vault *VaultCallerSession) AssetsLen() (*big.Int, error) {
	return _Vault.Contract.AssetsLen(&_Vault.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(token address) constant returns(uint256)
func (_Vault *VaultCaller) BalanceOf(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "balanceOf", token)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(token address) constant returns(uint256)
func (_Vault *VaultSession) BalanceOf(token common.Address) (*big.Int, error) {
	return _Vault.Contract.BalanceOf(&_Vault.CallOpts, token)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(token address) constant returns(uint256)
func (_Vault *VaultCallerSession) BalanceOf(token common.Address) (*big.Int, error) {
	return _Vault.Contract.BalanceOf(&_Vault.CallOpts, token)
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(address[], uint256[])
func (_Vault *VaultCaller) Balances(opts *bind.CallOpts) ([]common.Address, []*big.Int, error) {
	var (
		ret0 = new([]common.Address)
		ret1 = new([]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Vault.contract.Call(opts, out, "balances")
	return *ret0, *ret1, err
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(address[], uint256[])
func (_Vault *VaultSession) Balances() ([]common.Address, []*big.Int, error) {
	return _Vault.Contract.Balances(&_Vault.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(address[], uint256[])
func (_Vault *VaultCallerSession) Balances() ([]common.Address, []*big.Int, error) {
	return _Vault.Contract.Balances(&_Vault.CallOpts)
}

// BrgBalance is a free data retrieval call binding the contract method 0x24101d40.
//
// Solidity: function brgBalance() constant returns(uint256)
func (_Vault *VaultCaller) BrgBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "brgBalance")
	return *ret0, err
}

// BrgBalance is a free data retrieval call binding the contract method 0x24101d40.
//
// Solidity: function brgBalance() constant returns(uint256)
func (_Vault *VaultSession) BrgBalance() (*big.Int, error) {
	return _Vault.Contract.BrgBalance(&_Vault.CallOpts)
}

// BrgBalance is a free data retrieval call binding the contract method 0x24101d40.
//
// Solidity: function brgBalance() constant returns(uint256)
func (_Vault *VaultCallerSession) BrgBalance() (*big.Int, error) {
	return _Vault.Contract.BrgBalance(&_Vault.CallOpts)
}

// ContractHash is a free data retrieval call binding the contract method 0x904c6094.
//
// Solidity: function contractHash() constant returns(bytes32)
func (_Vault *VaultCaller) ContractHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "contractHash")
	return *ret0, err
}

// ContractHash is a free data retrieval call binding the contract method 0x904c6094.
//
// Solidity: function contractHash() constant returns(bytes32)
func (_Vault *VaultSession) ContractHash() ([32]byte, error) {
	return _Vault.Contract.ContractHash(&_Vault.CallOpts)
}

// ContractHash is a free data retrieval call binding the contract method 0x904c6094.
//
// Solidity: function contractHash() constant returns(bytes32)
func (_Vault *VaultCallerSession) ContractHash() ([32]byte, error) {
	return _Vault.Contract.ContractHash(&_Vault.CallOpts)
}

// EthBalance is a free data retrieval call binding the contract method 0x4e6630b0.
//
// Solidity: function ethBalance() constant returns(uint256)
func (_Vault *VaultCaller) EthBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "ethBalance")
	return *ret0, err
}

// EthBalance is a free data retrieval call binding the contract method 0x4e6630b0.
//
// Solidity: function ethBalance() constant returns(uint256)
func (_Vault *VaultSession) EthBalance() (*big.Int, error) {
	return _Vault.Contract.EthBalance(&_Vault.CallOpts)
}

// EthBalance is a free data retrieval call binding the contract method 0x4e6630b0.
//
// Solidity: function ethBalance() constant returns(uint256)
func (_Vault *VaultCallerSession) EthBalance() (*big.Int, error) {
	return _Vault.Contract.EthBalance(&_Vault.CallOpts)
}

// HasFunds is a free data retrieval call binding the contract method 0xf2cd59f6.
//
// Solidity: function hasFunds() constant returns(bool)
func (_Vault *VaultCaller) HasFunds(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "hasFunds")
	return *ret0, err
}

// HasFunds is a free data retrieval call binding the contract method 0xf2cd59f6.
//
// Solidity: function hasFunds() constant returns(bool)
func (_Vault *VaultSession) HasFunds() (bool, error) {
	return _Vault.Contract.HasFunds(&_Vault.CallOpts)
}

// HasFunds is a free data retrieval call binding the contract method 0xf2cd59f6.
//
// Solidity: function hasFunds() constant returns(bool)
func (_Vault *VaultCallerSession) HasFunds() (bool, error) {
	return _Vault.Contract.HasFunds(&_Vault.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0xe3c33a9b.
//
// Solidity: function hasRole(roleName string) constant returns(bool)
func (_Vault *VaultCaller) HasRole(opts *bind.CallOpts, roleName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "hasRole", roleName)
	return *ret0, err
}

// HasRole is a free data retrieval call binding the contract method 0xe3c33a9b.
//
// Solidity: function hasRole(roleName string) constant returns(bool)
func (_Vault *VaultSession) HasRole(roleName string) (bool, error) {
	return _Vault.Contract.HasRole(&_Vault.CallOpts, roleName)
}

// HasRole is a free data retrieval call binding the contract method 0xe3c33a9b.
//
// Solidity: function hasRole(roleName string) constant returns(bool)
func (_Vault *VaultCallerSession) HasRole(roleName string) (bool, error) {
	return _Vault.Contract.HasRole(&_Vault.CallOpts, roleName)
}

// IsEmpty is a free data retrieval call binding the contract method 0x681fe70c.
//
// Solidity: function isEmpty() constant returns(bool)
func (_Vault *VaultCaller) IsEmpty(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "isEmpty")
	return *ret0, err
}

// IsEmpty is a free data retrieval call binding the contract method 0x681fe70c.
//
// Solidity: function isEmpty() constant returns(bool)
func (_Vault *VaultSession) IsEmpty() (bool, error) {
	return _Vault.Contract.IsEmpty(&_Vault.CallOpts)
}

// IsEmpty is a free data retrieval call binding the contract method 0x681fe70c.
//
// Solidity: function isEmpty() constant returns(bool)
func (_Vault *VaultCallerSession) IsEmpty() (bool, error) {
	return _Vault.Contract.IsEmpty(&_Vault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Vault *VaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Vault *VaultSession) Owner() (common.Address, error) {
	return _Vault.Contract.Owner(&_Vault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Vault *VaultCallerSession) Owner() (common.Address, error) {
	return _Vault.Contract.Owner(&_Vault.CallOpts)
}

// Roles is a free data retrieval call binding the contract method 0x392f5f64.
//
// Solidity: function roles() constant returns(address)
func (_Vault *VaultCaller) Roles(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "roles")
	return *ret0, err
}

// Roles is a free data retrieval call binding the contract method 0x392f5f64.
//
// Solidity: function roles() constant returns(address)
func (_Vault *VaultSession) Roles() (common.Address, error) {
	return _Vault.Contract.Roles(&_Vault.CallOpts)
}

// Roles is a free data retrieval call binding the contract method 0x392f5f64.
//
// Solidity: function roles() constant returns(address)
func (_Vault *VaultCallerSession) Roles() (common.Address, error) {
	return _Vault.Contract.Roles(&_Vault.CallOpts)
}

// SenderHasRole is a free data retrieval call binding the contract method 0x1ca03b8e.
//
// Solidity: function senderHasRole(roleName string) constant returns(bool)
func (_Vault *VaultCaller) SenderHasRole(opts *bind.CallOpts, roleName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "senderHasRole", roleName)
	return *ret0, err
}

// SenderHasRole is a free data retrieval call binding the contract method 0x1ca03b8e.
//
// Solidity: function senderHasRole(roleName string) constant returns(bool)
func (_Vault *VaultSession) SenderHasRole(roleName string) (bool, error) {
	return _Vault.Contract.SenderHasRole(&_Vault.CallOpts, roleName)
}

// SenderHasRole is a free data retrieval call binding the contract method 0x1ca03b8e.
//
// Solidity: function senderHasRole(roleName string) constant returns(bool)
func (_Vault *VaultCallerSession) SenderHasRole(roleName string) (bool, error) {
	return _Vault.Contract.SenderHasRole(&_Vault.CallOpts, roleName)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_Vault *VaultCaller) Stopped(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "stopped")
	return *ret0, err
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_Vault *VaultSession) Stopped() (bool, error) {
	return _Vault.Contract.Stopped(&_Vault.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_Vault *VaultCallerSession) Stopped() (bool, error) {
	return _Vault.Contract.Stopped(&_Vault.CallOpts)
}

// UouCount is a free data retrieval call binding the contract method 0x5e762424.
//
// Solidity: function uouCount() constant returns(uint256)
func (_Vault *VaultCaller) UouCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "uouCount")
	return *ret0, err
}

// UouCount is a free data retrieval call binding the contract method 0x5e762424.
//
// Solidity: function uouCount() constant returns(uint256)
func (_Vault *VaultSession) UouCount() (*big.Int, error) {
	return _Vault.Contract.UouCount(&_Vault.CallOpts)
}

// UouCount is a free data retrieval call binding the contract method 0x5e762424.
//
// Solidity: function uouCount() constant returns(uint256)
func (_Vault *VaultCallerSession) UouCount() (*big.Int, error) {
	return _Vault.Contract.UouCount(&_Vault.CallOpts)
}

// Uous is a free data retrieval call binding the contract method 0x99133141.
//
// Solidity: function uous( uint256) constant returns(initialAmount uint128, repaidAmount uint128, fee uint128, time uint256, decision uint8)
func (_Vault *VaultCaller) Uous(opts *bind.CallOpts, arg0 *big.Int) (struct {
	InitialAmount *big.Int
	RepaidAmount  *big.Int
	Fee           *big.Int
	Time          *big.Int
	Decision      uint8
}, error) {
	ret := new(struct {
		InitialAmount *big.Int
		RepaidAmount  *big.Int
		Fee           *big.Int
		Time          *big.Int
		Decision      uint8
	})
	out := ret
	err := _Vault.contract.Call(opts, out, "uous", arg0)
	return *ret, err
}

// Uous is a free data retrieval call binding the contract method 0x99133141.
//
// Solidity: function uous( uint256) constant returns(initialAmount uint128, repaidAmount uint128, fee uint128, time uint256, decision uint8)
func (_Vault *VaultSession) Uous(arg0 *big.Int) (struct {
	InitialAmount *big.Int
	RepaidAmount  *big.Int
	Fee           *big.Int
	Time          *big.Int
	Decision      uint8
}, error) {
	return _Vault.Contract.Uous(&_Vault.CallOpts, arg0)
}

// Uous is a free data retrieval call binding the contract method 0x99133141.
//
// Solidity: function uous( uint256) constant returns(initialAmount uint128, repaidAmount uint128, fee uint128, time uint256, decision uint8)
func (_Vault *VaultCallerSession) Uous(arg0 *big.Int) (struct {
	InitialAmount *big.Int
	RepaidAmount  *big.Int
	Fee           *big.Int
	Time          *big.Int
	Decision      uint8
}, error) {
	return _Vault.Contract.Uous(&_Vault.CallOpts, arg0)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_Vault *VaultCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "wallet")
	return *ret0, err
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_Vault *VaultSession) Wallet() (common.Address, error) {
	return _Vault.Contract.Wallet(&_Vault.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_Vault *VaultCallerSession) Wallet() (common.Address, error) {
	return _Vault.Contract.Wallet(&_Vault.CallOpts)
}

// AcceptUouRequest is a paid mutator transaction binding the contract method 0xdc2628fa.
//
// Solidity: function acceptUouRequest(uouIndex uint256) returns()
func (_Vault *VaultTransactor) AcceptUouRequest(opts *bind.TransactOpts, uouIndex *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "acceptUouRequest", uouIndex)
}

// AcceptUouRequest is a paid mutator transaction binding the contract method 0xdc2628fa.
//
// Solidity: function acceptUouRequest(uouIndex uint256) returns()
func (_Vault *VaultSession) AcceptUouRequest(uouIndex *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.AcceptUouRequest(&_Vault.TransactOpts, uouIndex)
}

// AcceptUouRequest is a paid mutator transaction binding the contract method 0xdc2628fa.
//
// Solidity: function acceptUouRequest(uouIndex uint256) returns()
func (_Vault *VaultTransactorSession) AcceptUouRequest(uouIndex *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.AcceptUouRequest(&_Vault.TransactOpts, uouIndex)
}

// AddAsset is a paid mutator transaction binding the contract method 0x712e6279.
//
// Solidity: function addAsset(token address, src address, wad uint128) returns()
func (_Vault *VaultTransactor) AddAsset(opts *bind.TransactOpts, token common.Address, src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "addAsset", token, src, wad)
}

// AddAsset is a paid mutator transaction binding the contract method 0x712e6279.
//
// Solidity: function addAsset(token address, src address, wad uint128) returns()
func (_Vault *VaultSession) AddAsset(token common.Address, src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.AddAsset(&_Vault.TransactOpts, token, src, wad)
}

// AddAsset is a paid mutator transaction binding the contract method 0x712e6279.
//
// Solidity: function addAsset(token address, src address, wad uint128) returns()
func (_Vault *VaultTransactorSession) AddAsset(token common.Address, src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.AddAsset(&_Vault.TransactOpts, token, src, wad)
}

// CleanStorage is a paid mutator transaction binding the contract method 0x7e84d36b.
//
// Solidity: function cleanStorage() returns()
func (_Vault *VaultTransactor) CleanStorage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "cleanStorage")
}

// CleanStorage is a paid mutator transaction binding the contract method 0x7e84d36b.
//
// Solidity: function cleanStorage() returns()
func (_Vault *VaultSession) CleanStorage() (*types.Transaction, error) {
	return _Vault.Contract.CleanStorage(&_Vault.TransactOpts)
}

// CleanStorage is a paid mutator transaction binding the contract method 0x7e84d36b.
//
// Solidity: function cleanStorage() returns()
func (_Vault *VaultTransactorSession) CleanStorage() (*types.Transaction, error) {
	return _Vault.Contract.CleanStorage(&_Vault.TransactOpts)
}

// RejectUouRequest is a paid mutator transaction binding the contract method 0x22d8fe40.
//
// Solidity: function rejectUouRequest(uouIndex uint256) returns()
func (_Vault *VaultTransactor) RejectUouRequest(opts *bind.TransactOpts, uouIndex *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "rejectUouRequest", uouIndex)
}

// RejectUouRequest is a paid mutator transaction binding the contract method 0x22d8fe40.
//
// Solidity: function rejectUouRequest(uouIndex uint256) returns()
func (_Vault *VaultSession) RejectUouRequest(uouIndex *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.RejectUouRequest(&_Vault.TransactOpts, uouIndex)
}

// RejectUouRequest is a paid mutator transaction binding the contract method 0x22d8fe40.
//
// Solidity: function rejectUouRequest(uouIndex uint256) returns()
func (_Vault *VaultTransactorSession) RejectUouRequest(uouIndex *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.RejectUouRequest(&_Vault.TransactOpts, uouIndex)
}

// Remove is a paid mutator transaction binding the contract method 0xa7f43779.
//
// Solidity: function remove() returns()
func (_Vault *VaultTransactor) Remove(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "remove")
}

// Remove is a paid mutator transaction binding the contract method 0xa7f43779.
//
// Solidity: function remove() returns()
func (_Vault *VaultSession) Remove() (*types.Transaction, error) {
	return _Vault.Contract.Remove(&_Vault.TransactOpts)
}

// Remove is a paid mutator transaction binding the contract method 0xa7f43779.
//
// Solidity: function remove() returns()
func (_Vault *VaultTransactorSession) Remove() (*types.Transaction, error) {
	return _Vault.Contract.Remove(&_Vault.TransactOpts)
}

// RepayUou is a paid mutator transaction binding the contract method 0x1a93ebf0.
//
// Solidity: function repayUou(brgAmount uint128, uouIndex uint256) returns(uint128)
func (_Vault *VaultTransactor) RepayUou(opts *bind.TransactOpts, brgAmount *big.Int, uouIndex *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "repayUou", brgAmount, uouIndex)
}

// RepayUou is a paid mutator transaction binding the contract method 0x1a93ebf0.
//
// Solidity: function repayUou(brgAmount uint128, uouIndex uint256) returns(uint128)
func (_Vault *VaultSession) RepayUou(brgAmount *big.Int, uouIndex *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.RepayUou(&_Vault.TransactOpts, brgAmount, uouIndex)
}

// RepayUou is a paid mutator transaction binding the contract method 0x1a93ebf0.
//
// Solidity: function repayUou(brgAmount uint128, uouIndex uint256) returns(uint128)
func (_Vault *VaultTransactorSession) RepayUou(brgAmount *big.Int, uouIndex *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.RepayUou(&_Vault.TransactOpts, brgAmount, uouIndex)
}

// RequestUou is a paid mutator transaction binding the contract method 0xa52d7ffb.
//
// Solidity: function requestUou(brgAmount uint128) returns()
func (_Vault *VaultTransactor) RequestUou(opts *bind.TransactOpts, brgAmount *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "requestUou", brgAmount)
}

// RequestUou is a paid mutator transaction binding the contract method 0xa52d7ffb.
//
// Solidity: function requestUou(brgAmount uint128) returns()
func (_Vault *VaultSession) RequestUou(brgAmount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.RequestUou(&_Vault.TransactOpts, brgAmount)
}

// RequestUou is a paid mutator transaction binding the contract method 0xa52d7ffb.
//
// Solidity: function requestUou(brgAmount uint128) returns()
func (_Vault *VaultTransactorSession) RequestUou(brgAmount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.RequestUou(&_Vault.TransactOpts, brgAmount)
}

// Restart is a paid mutator transaction binding the contract method 0x1ef3755d.
//
// Solidity: function restart() returns()
func (_Vault *VaultTransactor) Restart(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "restart")
}

// Restart is a paid mutator transaction binding the contract method 0x1ef3755d.
//
// Solidity: function restart() returns()
func (_Vault *VaultSession) Restart() (*types.Transaction, error) {
	return _Vault.Contract.Restart(&_Vault.TransactOpts)
}

// Restart is a paid mutator transaction binding the contract method 0x1ef3755d.
//
// Solidity: function restart() returns()
func (_Vault *VaultTransactorSession) Restart() (*types.Transaction, error) {
	return _Vault.Contract.Restart(&_Vault.TransactOpts)
}

// RmAsset is a paid mutator transaction binding the contract method 0xd1da3e7b.
//
// Solidity: function rmAsset(token address, dst address) returns(bool)
func (_Vault *VaultTransactor) RmAsset(opts *bind.TransactOpts, token common.Address, dst common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "rmAsset", token, dst)
}

// RmAsset is a paid mutator transaction binding the contract method 0xd1da3e7b.
//
// Solidity: function rmAsset(token address, dst address) returns(bool)
func (_Vault *VaultSession) RmAsset(token common.Address, dst common.Address) (*types.Transaction, error) {
	return _Vault.Contract.RmAsset(&_Vault.TransactOpts, token, dst)
}

// RmAsset is a paid mutator transaction binding the contract method 0xd1da3e7b.
//
// Solidity: function rmAsset(token address, dst address) returns(bool)
func (_Vault *VaultTransactorSession) RmAsset(token common.Address, dst common.Address) (*types.Transaction, error) {
	return _Vault.Contract.RmAsset(&_Vault.TransactOpts, token, dst)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_Vault *VaultTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_Vault *VaultSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetOwner(&_Vault.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_Vault *VaultTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetOwner(&_Vault.TransactOpts, owner_)
}

// SetRolesContract is a paid mutator transaction binding the contract method 0xafa202ac.
//
// Solidity: function setRolesContract(roles_ address) returns()
func (_Vault *VaultTransactor) SetRolesContract(opts *bind.TransactOpts, roles_ common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setRolesContract", roles_)
}

// SetRolesContract is a paid mutator transaction binding the contract method 0xafa202ac.
//
// Solidity: function setRolesContract(roles_ address) returns()
func (_Vault *VaultSession) SetRolesContract(roles_ common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetRolesContract(&_Vault.TransactOpts, roles_)
}

// SetRolesContract is a paid mutator transaction binding the contract method 0xafa202ac.
//
// Solidity: function setRolesContract(roles_ address) returns()
func (_Vault *VaultTransactorSession) SetRolesContract(roles_ common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetRolesContract(&_Vault.TransactOpts, roles_)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Vault *VaultTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Vault *VaultSession) Stop() (*types.Transaction, error) {
	return _Vault.Contract.Stop(&_Vault.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Vault *VaultTransactorSession) Stop() (*types.Transaction, error) {
	return _Vault.Contract.Stop(&_Vault.TransactOpts)
}

// TakeOwnership is a paid mutator transaction binding the contract method 0x60536172.
//
// Solidity: function takeOwnership() returns()
func (_Vault *VaultTransactor) TakeOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "takeOwnership")
}

// TakeOwnership is a paid mutator transaction binding the contract method 0x60536172.
//
// Solidity: function takeOwnership() returns()
func (_Vault *VaultSession) TakeOwnership() (*types.Transaction, error) {
	return _Vault.Contract.TakeOwnership(&_Vault.TransactOpts)
}

// TakeOwnership is a paid mutator transaction binding the contract method 0x60536172.
//
// Solidity: function takeOwnership() returns()
func (_Vault *VaultTransactorSession) TakeOwnership() (*types.Transaction, error) {
	return _Vault.Contract.TakeOwnership(&_Vault.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0x56e596a8.
//
// Solidity: function transfer(token address, dst address, wad uint128) returns()
func (_Vault *VaultTransactor) Transfer(opts *bind.TransactOpts, token common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "transfer", token, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0x56e596a8.
//
// Solidity: function transfer(token address, dst address, wad uint128) returns()
func (_Vault *VaultSession) Transfer(token common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Transfer(&_Vault.TransactOpts, token, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0x56e596a8.
//
// Solidity: function transfer(token address, dst address, wad uint128) returns()
func (_Vault *VaultTransactorSession) Transfer(token common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Transfer(&_Vault.TransactOpts, token, dst, wad)
}

// TransferEth is a paid mutator transaction binding the contract method 0xe9bb84c2.
//
// Solidity: function transferEth(dst address, wad uint256) returns()
func (_Vault *VaultTransactor) TransferEth(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "transferEth", dst, wad)
}

// TransferEth is a paid mutator transaction binding the contract method 0xe9bb84c2.
//
// Solidity: function transferEth(dst address, wad uint256) returns()
func (_Vault *VaultSession) TransferEth(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.TransferEth(&_Vault.TransactOpts, dst, wad)
}

// TransferEth is a paid mutator transaction binding the contract method 0xe9bb84c2.
//
// Solidity: function transferEth(dst address, wad uint256) returns()
func (_Vault *VaultTransactorSession) TransferEth(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.TransferEth(&_Vault.TransactOpts, dst, wad)
}

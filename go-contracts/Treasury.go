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

// TreasuryABI is the input ABI used to generate the binding from.
const TreasuryABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"senderHasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"restart\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"v\",\"type\":\"address\"}],\"name\":\"addVault\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"roles\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"vaults\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"roles_\",\"type\":\"address\"}],\"name\":\"setRolesContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"hasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"root\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"r\",\"type\":\"address\"},{\"name\":\"rolesContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"}]"

// TreasuryBin is the compiled bytecode used for deploying new contracts.
const TreasuryBin = `undefined`

// DeployTreasury deploys a new Ethereum contract, binding an instance of Treasury to it.
func DeployTreasury(auth *bind.TransactOpts, backend bind.ContractBackend, r common.Address, rolesContract common.Address) (common.Address, *types.Transaction, *Treasury, error) {
	parsed, err := abi.JSON(strings.NewReader(TreasuryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TreasuryBin), backend, r, rolesContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Treasury{TreasuryCaller: TreasuryCaller{contract: contract}, TreasuryTransactor: TreasuryTransactor{contract: contract}}, nil
}

// Treasury is an auto generated Go binding around an Ethereum contract.
type Treasury struct {
	TreasuryCaller     // Read-only binding to the contract
	TreasuryTransactor // Write-only binding to the contract
}

// TreasuryCaller is an auto generated read-only Go binding around an Ethereum contract.
type TreasuryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TreasuryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TreasuryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TreasurySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TreasurySession struct {
	Contract     *Treasury         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TreasuryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TreasuryCallerSession struct {
	Contract *TreasuryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TreasuryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TreasuryTransactorSession struct {
	Contract     *TreasuryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TreasuryRaw is an auto generated low-level Go binding around an Ethereum contract.
type TreasuryRaw struct {
	Contract *Treasury // Generic contract binding to access the raw methods on
}

// TreasuryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TreasuryCallerRaw struct {
	Contract *TreasuryCaller // Generic read-only contract binding to access the raw methods on
}

// TreasuryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TreasuryTransactorRaw struct {
	Contract *TreasuryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTreasury creates a new instance of Treasury, bound to a specific deployed contract.
func NewTreasury(address common.Address, backend bind.ContractBackend) (*Treasury, error) {
	contract, err := bindTreasury(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Treasury{TreasuryCaller: TreasuryCaller{contract: contract}, TreasuryTransactor: TreasuryTransactor{contract: contract}}, nil
}

// NewTreasuryCaller creates a new read-only instance of Treasury, bound to a specific deployed contract.
func NewTreasuryCaller(address common.Address, caller bind.ContractCaller) (*TreasuryCaller, error) {
	contract, err := bindTreasury(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TreasuryCaller{contract: contract}, nil
}

// NewTreasuryTransactor creates a new write-only instance of Treasury, bound to a specific deployed contract.
func NewTreasuryTransactor(address common.Address, transactor bind.ContractTransactor) (*TreasuryTransactor, error) {
	contract, err := bindTreasury(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TreasuryTransactor{contract: contract}, nil
}

// bindTreasury binds a generic wrapper to an already deployed contract.
func bindTreasury(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TreasuryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Treasury *TreasuryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Treasury.Contract.TreasuryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Treasury *TreasuryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Treasury.Contract.TreasuryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Treasury *TreasuryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Treasury.Contract.TreasuryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Treasury *TreasuryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Treasury.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Treasury *TreasuryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Treasury.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Treasury *TreasuryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Treasury.Contract.contract.Transact(opts, method, params...)
}

// ContractHash is a free data retrieval call binding the contract method 0x904c6094.
//
// Solidity: function contractHash() constant returns(bytes32)
func (_Treasury *TreasuryCaller) ContractHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Treasury.contract.Call(opts, out, "contractHash")
	return *ret0, err
}

// ContractHash is a free data retrieval call binding the contract method 0x904c6094.
//
// Solidity: function contractHash() constant returns(bytes32)
func (_Treasury *TreasurySession) ContractHash() ([32]byte, error) {
	return _Treasury.Contract.ContractHash(&_Treasury.CallOpts)
}

// ContractHash is a free data retrieval call binding the contract method 0x904c6094.
//
// Solidity: function contractHash() constant returns(bytes32)
func (_Treasury *TreasuryCallerSession) ContractHash() ([32]byte, error) {
	return _Treasury.Contract.ContractHash(&_Treasury.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0xe3c33a9b.
//
// Solidity: function hasRole(roleName string) constant returns(bool)
func (_Treasury *TreasuryCaller) HasRole(opts *bind.CallOpts, roleName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Treasury.contract.Call(opts, out, "hasRole", roleName)
	return *ret0, err
}

// HasRole is a free data retrieval call binding the contract method 0xe3c33a9b.
//
// Solidity: function hasRole(roleName string) constant returns(bool)
func (_Treasury *TreasurySession) HasRole(roleName string) (bool, error) {
	return _Treasury.Contract.HasRole(&_Treasury.CallOpts, roleName)
}

// HasRole is a free data retrieval call binding the contract method 0xe3c33a9b.
//
// Solidity: function hasRole(roleName string) constant returns(bool)
func (_Treasury *TreasuryCallerSession) HasRole(roleName string) (bool, error) {
	return _Treasury.Contract.HasRole(&_Treasury.CallOpts, roleName)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Treasury *TreasuryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Treasury.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Treasury *TreasurySession) Owner() (common.Address, error) {
	return _Treasury.Contract.Owner(&_Treasury.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Treasury *TreasuryCallerSession) Owner() (common.Address, error) {
	return _Treasury.Contract.Owner(&_Treasury.CallOpts)
}

// Roles is a free data retrieval call binding the contract method 0x392f5f64.
//
// Solidity: function roles() constant returns(address)
func (_Treasury *TreasuryCaller) Roles(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Treasury.contract.Call(opts, out, "roles")
	return *ret0, err
}

// Roles is a free data retrieval call binding the contract method 0x392f5f64.
//
// Solidity: function roles() constant returns(address)
func (_Treasury *TreasurySession) Roles() (common.Address, error) {
	return _Treasury.Contract.Roles(&_Treasury.CallOpts)
}

// Roles is a free data retrieval call binding the contract method 0x392f5f64.
//
// Solidity: function roles() constant returns(address)
func (_Treasury *TreasuryCallerSession) Roles() (common.Address, error) {
	return _Treasury.Contract.Roles(&_Treasury.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() constant returns(address)
func (_Treasury *TreasuryCaller) Root(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Treasury.contract.Call(opts, out, "root")
	return *ret0, err
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() constant returns(address)
func (_Treasury *TreasurySession) Root() (common.Address, error) {
	return _Treasury.Contract.Root(&_Treasury.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() constant returns(address)
func (_Treasury *TreasuryCallerSession) Root() (common.Address, error) {
	return _Treasury.Contract.Root(&_Treasury.CallOpts)
}

// SenderHasRole is a free data retrieval call binding the contract method 0x1ca03b8e.
//
// Solidity: function senderHasRole(roleName string) constant returns(bool)
func (_Treasury *TreasuryCaller) SenderHasRole(opts *bind.CallOpts, roleName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Treasury.contract.Call(opts, out, "senderHasRole", roleName)
	return *ret0, err
}

// SenderHasRole is a free data retrieval call binding the contract method 0x1ca03b8e.
//
// Solidity: function senderHasRole(roleName string) constant returns(bool)
func (_Treasury *TreasurySession) SenderHasRole(roleName string) (bool, error) {
	return _Treasury.Contract.SenderHasRole(&_Treasury.CallOpts, roleName)
}

// SenderHasRole is a free data retrieval call binding the contract method 0x1ca03b8e.
//
// Solidity: function senderHasRole(roleName string) constant returns(bool)
func (_Treasury *TreasuryCallerSession) SenderHasRole(roleName string) (bool, error) {
	return _Treasury.Contract.SenderHasRole(&_Treasury.CallOpts, roleName)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_Treasury *TreasuryCaller) Stopped(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Treasury.contract.Call(opts, out, "stopped")
	return *ret0, err
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_Treasury *TreasurySession) Stopped() (bool, error) {
	return _Treasury.Contract.Stopped(&_Treasury.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_Treasury *TreasuryCallerSession) Stopped() (bool, error) {
	return _Treasury.Contract.Stopped(&_Treasury.CallOpts)
}

// Vaults is a free data retrieval call binding the contract method 0xa622ee7c.
//
// Solidity: function vaults( address) constant returns(address)
func (_Treasury *TreasuryCaller) Vaults(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Treasury.contract.Call(opts, out, "vaults", arg0)
	return *ret0, err
}

// Vaults is a free data retrieval call binding the contract method 0xa622ee7c.
//
// Solidity: function vaults( address) constant returns(address)
func (_Treasury *TreasurySession) Vaults(arg0 common.Address) (common.Address, error) {
	return _Treasury.Contract.Vaults(&_Treasury.CallOpts, arg0)
}

// Vaults is a free data retrieval call binding the contract method 0xa622ee7c.
//
// Solidity: function vaults( address) constant returns(address)
func (_Treasury *TreasuryCallerSession) Vaults(arg0 common.Address) (common.Address, error) {
	return _Treasury.Contract.Vaults(&_Treasury.CallOpts, arg0)
}

// AddVault is a paid mutator transaction binding the contract method 0x256b5a02.
//
// Solidity: function addVault(v address) returns()
func (_Treasury *TreasuryTransactor) AddVault(opts *bind.TransactOpts, v common.Address) (*types.Transaction, error) {
	return _Treasury.contract.Transact(opts, "addVault", v)
}

// AddVault is a paid mutator transaction binding the contract method 0x256b5a02.
//
// Solidity: function addVault(v address) returns()
func (_Treasury *TreasurySession) AddVault(v common.Address) (*types.Transaction, error) {
	return _Treasury.Contract.AddVault(&_Treasury.TransactOpts, v)
}

// AddVault is a paid mutator transaction binding the contract method 0x256b5a02.
//
// Solidity: function addVault(v address) returns()
func (_Treasury *TreasuryTransactorSession) AddVault(v common.Address) (*types.Transaction, error) {
	return _Treasury.Contract.AddVault(&_Treasury.TransactOpts, v)
}

// Restart is a paid mutator transaction binding the contract method 0x1ef3755d.
//
// Solidity: function restart() returns()
func (_Treasury *TreasuryTransactor) Restart(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Treasury.contract.Transact(opts, "restart")
}

// Restart is a paid mutator transaction binding the contract method 0x1ef3755d.
//
// Solidity: function restart() returns()
func (_Treasury *TreasurySession) Restart() (*types.Transaction, error) {
	return _Treasury.Contract.Restart(&_Treasury.TransactOpts)
}

// Restart is a paid mutator transaction binding the contract method 0x1ef3755d.
//
// Solidity: function restart() returns()
func (_Treasury *TreasuryTransactorSession) Restart() (*types.Transaction, error) {
	return _Treasury.Contract.Restart(&_Treasury.TransactOpts)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_Treasury *TreasuryTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _Treasury.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_Treasury *TreasurySession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Treasury.Contract.SetOwner(&_Treasury.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_Treasury *TreasuryTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Treasury.Contract.SetOwner(&_Treasury.TransactOpts, owner_)
}

// SetRolesContract is a paid mutator transaction binding the contract method 0xafa202ac.
//
// Solidity: function setRolesContract(roles_ address) returns()
func (_Treasury *TreasuryTransactor) SetRolesContract(opts *bind.TransactOpts, roles_ common.Address) (*types.Transaction, error) {
	return _Treasury.contract.Transact(opts, "setRolesContract", roles_)
}

// SetRolesContract is a paid mutator transaction binding the contract method 0xafa202ac.
//
// Solidity: function setRolesContract(roles_ address) returns()
func (_Treasury *TreasurySession) SetRolesContract(roles_ common.Address) (*types.Transaction, error) {
	return _Treasury.Contract.SetRolesContract(&_Treasury.TransactOpts, roles_)
}

// SetRolesContract is a paid mutator transaction binding the contract method 0xafa202ac.
//
// Solidity: function setRolesContract(roles_ address) returns()
func (_Treasury *TreasuryTransactorSession) SetRolesContract(roles_ common.Address) (*types.Transaction, error) {
	return _Treasury.Contract.SetRolesContract(&_Treasury.TransactOpts, roles_)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Treasury *TreasuryTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Treasury.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Treasury *TreasurySession) Stop() (*types.Transaction, error) {
	return _Treasury.Contract.Stop(&_Treasury.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Treasury *TreasuryTransactorSession) Stop() (*types.Transaction, error) {
	return _Treasury.Contract.Stop(&_Treasury.TransactOpts)
}

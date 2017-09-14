// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package contracts

import (
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// RootABI is the input ABI used to generate the binding from.
const RootABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"swt_\",\"type\":\"address\"}],\"name\":\"setSWT\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"brg\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vc\",\"type\":\"address\"}],\"name\":\"setVaultConfig\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"brg_\",\"type\":\"address\"}],\"name\":\"setBRG\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"userDirectories\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vaultConfig\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"removeDirectory\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"swt\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"t\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"directory\",\"type\":\"address\"}],\"name\":\"addDirectory\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"brg_\",\"type\":\"address\"},{\"name\":\"swt_\",\"type\":\"address\"},{\"name\":\"vc\",\"type\":\"address\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"foo\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"bar\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fax\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"UnauthorizedAccess\",\"type\":\"event\"}]"

// RootBin is the compiled bytecode used for deploying new contracts.
const RootBin = `"0x6060604052341561000f57600080fd5b604051606080610ab28339810160405280805191906020018051919060200180519150505b5b60018054600160a060020a03191633600160a060020a03169081179091557fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a25b600160a060020a038316151561009557600080fd5b600160a060020a03821615156100aa57600080fd5b600160a060020a03811615156100bf57600080fd5b60028054600160a060020a03808616600160a060020a0319928316179092556003805485841690831617905560058054928416929091169190911790555b5050505b6109a2806101106000396000f300606060405236156100f95763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166307da68f581146100fe5780630d1a8df01461011357806313af4035146101345780634f9c8fe8146101555780635faafcaa1461018457806361d027b3146101a5578063679ccc6d146101d457806375f12b21146101f55780637a9e5e4b1461021c5780637b2fa97d1461023d5780637cc34bb4146102785780638da5cb5b146102a7578063be9a6555146102d6578063bf7e214f146102eb578063d25ad88b1461031a578063f016b6b61461033b578063f0f442601461036a578063f1ed4eda1461038b575b600080fd5b341561010957600080fd5b6101116103b2565b005b341561011e57600080fd5b610111600160a060020a0360043516610462565b005b341561013f57600080fd5b610111600160a060020a03600435166104a0565b005b341561016057600080fd5b610168610511565b604051600160a060020a03909116815260200160405180910390f35b341561018f57600080fd5b610111600160a060020a0360043516610520565b005b34156101b057600080fd5b61016861055e565b604051600160a060020a03909116815260200160405180910390f35b34156101df57600080fd5b610111600160a060020a036004351661056d565b005b341561020057600080fd5b6102086105ab565b604051901515815260200160405180910390f35b341561022757600080fd5b610111600160a060020a03600435166105cc565b005b341561024857600080fd5b610168600160a060020a036004351661063d565b604051600160a060020a03909116815260200160405180910390f35b341561028357600080fd5b610168610658565b604051600160a060020a03909116815260200160405180910390f35b34156102b257600080fd5b610168610667565b604051600160a060020a03909116815260200160405180910390f35b34156102e157600080fd5b610111610676565b005b34156102f657600080fd5b61016861070f565b604051600160a060020a03909116815260200160405180910390f35b341561032557600080fd5b610111600160a060020a036004351661071e565b005b341561034657600080fd5b610168610748565b604051600160a060020a03909116815260200160405180910390f35b341561037557600080fd5b610111600160a060020a0360043516610757565b005b341561039657600080fd5b610111600160a060020a0360043581169060243516610795565b005b6103c833600035600160e060020a0319166107ff565b15156103d057fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a46001805474ff00000000000000000000000000000000000000001916740100000000000000000000000000000000000000001790555b5b50505b565b61047833600035600160e060020a0319166107ff565b151561048057fe5b60038054600160a060020a031916600160a060020a0383161790555b5b50565b6104b633600035600160e060020a0319166107ff565b15156104be57fe5b60018054600160a060020a031916600160a060020a038381169190911791829055167fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a25b5b50565b600254600160a060020a031681565b61053633600035600160e060020a0319166107ff565b151561053e57fe5b60058054600160a060020a031916600160a060020a0383161790555b5b50565b600454600160a060020a031681565b61058333600035600160e060020a0319166107ff565b151561058b57fe5b60028054600160a060020a031916600160a060020a0383161790555b5b50565b60015474010000000000000000000000000000000000000000900460ff1681565b6105e233600035600160e060020a0319166107ff565b15156105ea57fe5b60008054600160a060020a031916600160a060020a038381169190911791829055167f1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada460405160405180910390a25b5b50565b600660205260009081526040902054600160a060020a031681565b600554600160a060020a031681565b600154600160a060020a031681565b61068c33600035600160e060020a0319166107ff565b151561069457fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a46001805474ff0000000000000000000000000000000000000000191690555b5b50505b565b600054600160a060020a031681565b600160a060020a03811660009081526006602052604090208054600160a060020a03191690555b50565b600354600160a060020a031681565b61076d33600035600160e060020a0319166107ff565b151561077557fe5b60048054600160a060020a031916600160a060020a0383161790555b5b50565b600160a060020a03821615156107a757fe5b600160a060020a0382811660009081526006602052604090205416156107cc57600080fd5b600160a060020a0382811660009081526006602052604090208054600160a060020a0319169183169190911790555b5050565b600030600160a060020a031683600160a060020a031614156108235750600161096d565b600154600160a060020a0384811691161480156108495750600054600160a060020a0316155b156108565750600161096d565b600054600160a060020a031615156108be577feb91385a0d70e44b915093d9ddfe6c8b41f2c56729431b1405f40aee5d874be78383604051600160a060020a039092168252600160e060020a03191660208201526040908101905180910390a150600061096d565b60008054600160a060020a03169063b700961390859030908690604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600160a060020a039384166004820152919092166024820152600160e060020a03199091166044820152606401602060405180830381600087803b151561095057600080fd5b6102c65a03f1151561096157600080fd5b50505060405180519150505b5b5b5b929150505600a165627a7a723058200b345485f07fc137dbbb31cd1f8f1beccdf23977c54327eb2d410fbe5ca441310029"`

// DeployRoot deploys a new Ethereum contract, binding an instance of Root to it.
func DeployRoot(auth *bind.TransactOpts, backend bind.ContractBackend, brg_ common.Address, swt_ common.Address, vc common.Address) (common.Address, *types.Transaction, *Root, error) {
	parsed, err := abi.JSON(strings.NewReader(RootABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RootBin), backend, brg_, swt_, vc)
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

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Root *RootCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Root.contract.Call(opts, out, "authority")
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Root *RootSession) Authority() (common.Address, error) {
	return _Root.Contract.Authority(&_Root.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Root *RootCallerSession) Authority() (common.Address, error) {
	return _Root.Contract.Authority(&_Root.CallOpts)
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

// AddDirectory is a paid mutator transaction binding the contract method 0xf1ed4eda.
//
// Solidity: function addDirectory(owner address, directory address) returns()
func (_Root *RootTransactor) AddDirectory(opts *bind.TransactOpts, owner common.Address, directory common.Address) (*types.Transaction, error) {
	return _Root.contract.Transact(opts, "addDirectory", owner, directory)
}

// AddDirectory is a paid mutator transaction binding the contract method 0xf1ed4eda.
//
// Solidity: function addDirectory(owner address, directory address) returns()
func (_Root *RootSession) AddDirectory(owner common.Address, directory common.Address) (*types.Transaction, error) {
	return _Root.Contract.AddDirectory(&_Root.TransactOpts, owner, directory)
}

// AddDirectory is a paid mutator transaction binding the contract method 0xf1ed4eda.
//
// Solidity: function addDirectory(owner address, directory address) returns()
func (_Root *RootTransactorSession) AddDirectory(owner common.Address, directory common.Address) (*types.Transaction, error) {
	return _Root.Contract.AddDirectory(&_Root.TransactOpts, owner, directory)
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

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_Root *RootTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _Root.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_Root *RootSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Root.Contract.SetAuthority(&_Root.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_Root *RootTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Root.Contract.SetAuthority(&_Root.TransactOpts, authority_)
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

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Root *RootTransactor) Start(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Root.contract.Transact(opts, "start")
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Root *RootSession) Start() (*types.Transaction, error) {
	return _Root.Contract.Start(&_Root.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Root *RootTransactorSession) Start() (*types.Transaction, error) {
	return _Root.Contract.Start(&_Root.TransactOpts)
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

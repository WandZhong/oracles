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
const RootABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"swt_\",\"type\":\"address\"}],\"name\":\"setSWT\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"brg\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"brg_\",\"type\":\"address\"}],\"name\":\"setBRG\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"userDirectories\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"removeDirectory\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"swt\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"t\",\"type\":\"address\"}],\"name\":\"setTreasury\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"directory\",\"type\":\"address\"}],\"name\":\"addDirectory\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"brg_\",\"type\":\"address\"},{\"name\":\"swt_\",\"type\":\"address\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"foo\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"bar\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fax\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"UnauthorizedAccess\",\"type\":\"event\"}]"

// RootBin is the compiled bytecode used for deploying new contracts.
const RootBin = `"0x6060604052341561000f57600080fd5b6040516040806109b083398101604052808051919060200180519150505b5b60018054600160a060020a03191633600160a060020a03169081179091557fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a25b600160a060020a038216151561008e57600080fd5b600160a060020a03811615156100a357600080fd5b60028054600160a060020a03808516600160a060020a03199283161790925560038054928416929091169190911790555b50505b6108ca806100e66000396000f300606060405236156100e35763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166307da68f581146100e85780630d1a8df0146100fd57806313af40351461011e5780634f9c8fe81461013f57806361d027b31461016e578063679ccc6d1461019d57806375f12b21146101be5780637a9e5e4b146101e55780637b2fa97d146102065780638da5cb5b14610241578063be9a655514610270578063bf7e214f14610285578063d25ad88b146102b4578063f016b6b6146102d5578063f0f4426014610304578063f1ed4eda14610325575b600080fd5b34156100f357600080fd5b6100fb61034c565b005b341561010857600080fd5b6100fb600160a060020a03600435166103fc565b005b341561012957600080fd5b6100fb600160a060020a036004351661043a565b005b341561014a57600080fd5b6101526104ab565b604051600160a060020a03909116815260200160405180910390f35b341561017957600080fd5b6101526104ba565b604051600160a060020a03909116815260200160405180910390f35b34156101a857600080fd5b6100fb600160a060020a03600435166104c9565b005b34156101c957600080fd5b6101d1610507565b604051901515815260200160405180910390f35b34156101f057600080fd5b6100fb600160a060020a0360043516610528565b005b341561021157600080fd5b610152600160a060020a0360043516610599565b604051600160a060020a03909116815260200160405180910390f35b341561024c57600080fd5b6101526105b4565b604051600160a060020a03909116815260200160405180910390f35b341561027b57600080fd5b6100fb6105c3565b005b341561029057600080fd5b61015261065c565b604051600160a060020a03909116815260200160405180910390f35b34156102bf57600080fd5b6100fb600160a060020a036004351661066b565b005b34156102e057600080fd5b610152610695565b604051600160a060020a03909116815260200160405180910390f35b341561030f57600080fd5b6100fb600160a060020a03600435166106a4565b005b341561033057600080fd5b6100fb600160a060020a03600435811690602435166106e2565b005b61036233600035600160e060020a031916610727565b151561036a57fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a46001805474ff00000000000000000000000000000000000000001916740100000000000000000000000000000000000000001790555b5b50505b565b61041233600035600160e060020a031916610727565b151561041a57fe5b60038054600160a060020a031916600160a060020a0383161790555b5b50565b61045033600035600160e060020a031916610727565b151561045857fe5b60018054600160a060020a031916600160a060020a038381169190911791829055167fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a25b5b50565b600254600160a060020a031681565b600454600160a060020a031681565b6104df33600035600160e060020a031916610727565b15156104e757fe5b60028054600160a060020a031916600160a060020a0383161790555b5b50565b60015474010000000000000000000000000000000000000000900460ff1681565b61053e33600035600160e060020a031916610727565b151561054657fe5b60008054600160a060020a031916600160a060020a038381169190911791829055167f1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada460405160405180910390a25b5b50565b600560205260009081526040902054600160a060020a031681565b600154600160a060020a031681565b6105d933600035600160e060020a031916610727565b15156105e157fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a46001805474ff0000000000000000000000000000000000000000191690555b5b50505b565b600054600160a060020a031681565b600160a060020a03811660009081526005602052604090208054600160a060020a03191690555b50565b600354600160a060020a031681565b6106ba33600035600160e060020a031916610727565b15156106c257fe5b60048054600160a060020a031916600160a060020a0383161790555b5b50565b600160a060020a03821615156106f457fe5b600160a060020a0382811660009081526005602052604090208054600160a060020a0319169183169190911790555b5050565b600030600160a060020a031683600160a060020a0316141561074b57506001610895565b600154600160a060020a0384811691161480156107715750600054600160a060020a0316155b1561077e57506001610895565b600054600160a060020a031615156107e6577feb91385a0d70e44b915093d9ddfe6c8b41f2c56729431b1405f40aee5d874be78383604051600160a060020a039092168252600160e060020a03191660208201526040908101905180910390a1506000610895565b60008054600160a060020a03169063b700961390859030908690604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600160a060020a039384166004820152919092166024820152600160e060020a03199091166044820152606401602060405180830381600087803b151561087857600080fd5b6102c65a03f1151561088957600080fd5b50505060405180519150505b5b5b5b929150505600a165627a7a72305820b7685718f6176be27e96b0c1f5456cbc71f3087a04df5bdbb74b854dd6228b370029"`

// DeployRoot deploys a new Ethereum contract, binding an instance of Root to it.
func DeployRoot(auth *bind.TransactOpts, backend bind.ContractBackend, brg_ common.Address, swt_ common.Address) (common.Address, *types.Transaction, *Root, error) {
	parsed, err := abi.JSON(strings.NewReader(RootABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(RootBin), backend, brg_, swt_)
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

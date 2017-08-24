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

// StoppableABI is the input ABI used to generate the binding from.
const StoppableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"foo\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"bar\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fax\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"UnauthorizedAccess\",\"type\":\"event\"}]"

// StoppableBin is the compiled bytecode used for deploying new contracts.
const StoppableBin = `"0x60606040525b60018054600160a060020a03191633600160a060020a03169081179091557fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a25b5b61059d806100606000396000f300606060405236156100805763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166307da68f5811461008557806313af40351461009a57806375f12b21146100bb5780637a9e5e4b146100e25780638da5cb5b14610103578063be9a655514610132578063bf7e214f14610147575b600080fd5b341561009057600080fd5b610098610176565b005b34156100a557600080fd5b610098600160a060020a0360043516610226565b005b34156100c657600080fd5b6100ce6102a4565b604051901515815260200160405180910390f35b34156100ed57600080fd5b610098600160a060020a03600435166102c5565b005b341561010e57600080fd5b610116610343565b604051600160a060020a03909116815260200160405180910390f35b341561013d57600080fd5b610098610352565b005b341561015257600080fd5b6101166103eb565b604051600160a060020a03909116815260200160405180910390f35b61018c33600035600160e060020a0319166103fa565b151561019457fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a46001805474ff00000000000000000000000000000000000000001916740100000000000000000000000000000000000000001790555b5b50505b565b61023c33600035600160e060020a0319166103fa565b151561024457fe5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a25b5b50565b60015474010000000000000000000000000000000000000000900460ff1681565b6102db33600035600160e060020a0319166103fa565b15156102e357fe5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167f1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada460405160405180910390a25b5b50565b600154600160a060020a031681565b61036833600035600160e060020a0319166103fa565b151561037057fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a46001805474ff0000000000000000000000000000000000000000191690555b5b50505b565b600054600160a060020a031681565b600030600160a060020a031683600160a060020a0316141561041e57506001610568565b600154600160a060020a0384811691161480156104445750600054600160a060020a0316155b1561045157506001610568565b600054600160a060020a031615156104b9577feb91385a0d70e44b915093d9ddfe6c8b41f2c56729431b1405f40aee5d874be78383604051600160a060020a039092168252600160e060020a03191660208201526040908101905180910390a1506000610568565b60008054600160a060020a03169063b700961390859030908690604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600160a060020a039384166004820152919092166024820152600160e060020a03199091166044820152606401602060405180830381600087803b151561054b57600080fd5b6102c65a03f1151561055c57600080fd5b50505060405180519150505b5b5b5b929150505600a165627a7a72305820b77bcf3225b5109b581af63fcd29a49559c728d94149a7552e83d4a3a9e3fd600029"`

// DeployStoppable deploys a new Ethereum contract, binding an instance of Stoppable to it.
func DeployStoppable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Stoppable, error) {
	parsed, err := abi.JSON(strings.NewReader(StoppableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StoppableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Stoppable{StoppableCaller: StoppableCaller{contract: contract}, StoppableTransactor: StoppableTransactor{contract: contract}}, nil
}

// Stoppable is an auto generated Go binding around an Ethereum contract.
type Stoppable struct {
	StoppableCaller     // Read-only binding to the contract
	StoppableTransactor // Write-only binding to the contract
}

// StoppableCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoppableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoppableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoppableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoppableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoppableSession struct {
	Contract     *Stoppable        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoppableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoppableCallerSession struct {
	Contract *StoppableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// StoppableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoppableTransactorSession struct {
	Contract     *StoppableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// StoppableRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoppableRaw struct {
	Contract *Stoppable // Generic contract binding to access the raw methods on
}

// StoppableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoppableCallerRaw struct {
	Contract *StoppableCaller // Generic read-only contract binding to access the raw methods on
}

// StoppableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoppableTransactorRaw struct {
	Contract *StoppableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStoppable creates a new instance of Stoppable, bound to a specific deployed contract.
func NewStoppable(address common.Address, backend bind.ContractBackend) (*Stoppable, error) {
	contract, err := bindStoppable(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stoppable{StoppableCaller: StoppableCaller{contract: contract}, StoppableTransactor: StoppableTransactor{contract: contract}}, nil
}

// NewStoppableCaller creates a new read-only instance of Stoppable, bound to a specific deployed contract.
func NewStoppableCaller(address common.Address, caller bind.ContractCaller) (*StoppableCaller, error) {
	contract, err := bindStoppable(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &StoppableCaller{contract: contract}, nil
}

// NewStoppableTransactor creates a new write-only instance of Stoppable, bound to a specific deployed contract.
func NewStoppableTransactor(address common.Address, transactor bind.ContractTransactor) (*StoppableTransactor, error) {
	contract, err := bindStoppable(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &StoppableTransactor{contract: contract}, nil
}

// bindStoppable binds a generic wrapper to an already deployed contract.
func bindStoppable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StoppableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stoppable *StoppableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Stoppable.Contract.StoppableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stoppable *StoppableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stoppable.Contract.StoppableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stoppable *StoppableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stoppable.Contract.StoppableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stoppable *StoppableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Stoppable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stoppable *StoppableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stoppable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stoppable *StoppableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stoppable.Contract.contract.Transact(opts, method, params...)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Stoppable *StoppableCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Stoppable.contract.Call(opts, out, "authority")
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Stoppable *StoppableSession) Authority() (common.Address, error) {
	return _Stoppable.Contract.Authority(&_Stoppable.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Stoppable *StoppableCallerSession) Authority() (common.Address, error) {
	return _Stoppable.Contract.Authority(&_Stoppable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Stoppable *StoppableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Stoppable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Stoppable *StoppableSession) Owner() (common.Address, error) {
	return _Stoppable.Contract.Owner(&_Stoppable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Stoppable *StoppableCallerSession) Owner() (common.Address, error) {
	return _Stoppable.Contract.Owner(&_Stoppable.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_Stoppable *StoppableCaller) Stopped(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Stoppable.contract.Call(opts, out, "stopped")
	return *ret0, err
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_Stoppable *StoppableSession) Stopped() (bool, error) {
	return _Stoppable.Contract.Stopped(&_Stoppable.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_Stoppable *StoppableCallerSession) Stopped() (bool, error) {
	return _Stoppable.Contract.Stopped(&_Stoppable.CallOpts)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_Stoppable *StoppableTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _Stoppable.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_Stoppable *StoppableSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Stoppable.Contract.SetAuthority(&_Stoppable.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_Stoppable *StoppableTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Stoppable.Contract.SetAuthority(&_Stoppable.TransactOpts, authority_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_Stoppable *StoppableTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _Stoppable.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_Stoppable *StoppableSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Stoppable.Contract.SetOwner(&_Stoppable.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_Stoppable *StoppableTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Stoppable.Contract.SetOwner(&_Stoppable.TransactOpts, owner_)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Stoppable *StoppableTransactor) Start(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stoppable.contract.Transact(opts, "start")
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Stoppable *StoppableSession) Start() (*types.Transaction, error) {
	return _Stoppable.Contract.Start(&_Stoppable.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Stoppable *StoppableTransactorSession) Start() (*types.Transaction, error) {
	return _Stoppable.Contract.Start(&_Stoppable.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Stoppable *StoppableTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stoppable.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Stoppable *StoppableSession) Stop() (*types.Transaction, error) {
	return _Stoppable.Contract.Stop(&_Stoppable.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Stoppable *StoppableTransactorSession) Stop() (*types.Transaction, error) {
	return _Stoppable.Contract.Stop(&_Stoppable.TransactOpts)
}

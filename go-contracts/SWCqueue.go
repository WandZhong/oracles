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

// SWCqueueABI is the input ABI used to generate the binding from.
const SWCqueueABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint128\"},{\"name\":\"currency\",\"type\":\"bytes3\"}],\"name\":\"directPledge\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"logTrancheRelease\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ratio\",\"type\":\"uint256\"}],\"name\":\"setNextBRGSWTratio\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"brg\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"brg_\",\"type\":\"address\"}],\"name\":\"setBRG\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wad\",\"type\":\"uint128\"},{\"name\":\"currency\",\"type\":\"bytes3\"}],\"name\":\"cancel\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"brgAmount\",\"type\":\"uint128\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextBRGSWTratio\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"brg_\",\"type\":\"address\"},{\"name\":\"nextBRGSWTratio_\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"UnauthorizedAccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint128\"},{\"indexed\":false,\"name\":\"currency\",\"type\":\"bytes3\"}],\"name\":\"LogSWCqueueCancel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint128\"},{\"indexed\":false,\"name\":\"currency\",\"type\":\"bytes3\"}],\"name\":\"LogSWCqueueDirectPledge\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"LogSWCqueueTranchRelease\",\"type\":\"event\"}]"

// SWCqueueBin is the compiled bytecode used for deploying new contracts.
const SWCqueueBin = `"0x6060604052341561000f57600080fd5b60405160408061083783398101604052808051919060200180519150505b5b60018054600160a060020a03191633600160a060020a03169081179091557fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a25b600160a060020a038216151561008e57600080fd5b60028054600160a060020a031916600160a060020a03841617905560038190555b50505b610776806100c16000396000f3006060604052361561009e5763ffffffff60e060020a60003504166313af403581146100a3578063209c4986146100c457806330311449146100fe578063429ad4f21461011f5780634f9c8fe814610137578063679ccc6d146101665780637a9e5e4b146101875780638da5cb5b146101a857806390a160d4146101d757806390bc169314610205578063ac63f9d914610226578063bf7e214f1461024b575b600080fd5b34156100ae57600080fd5b6100c2600160a060020a036004351661027a565b005b34156100cf57600080fd5b6100c2600160a060020a03600435166001608060020a0360243516600160e860020a0319604435166102f8565b005b341561010957600080fd5b6100c26001608060020a0360043516610361565b005b341561012a57600080fd5b6100c26004356103b7565b005b341561014257600080fd5b61014a6103ec565b604051600160a060020a03909116815260200160405180910390f35b341561017157600080fd5b6100c2600160a060020a03600435166103fb565b005b341561019257600080fd5b6100c2600160a060020a036004351661045b565b005b34156101b357600080fd5b61014a6104d9565b604051600160a060020a03909116815260200160405180910390f35b34156101e257600080fd5b6100c26001608060020a0360043516600160e860020a0319602435166104e8565b005b341561021057600080fd5b6100c26001608060020a0360043516610550565b005b341561023157600080fd5b6102396105d7565b60405190815260200160405180910390f35b341561025657600080fd5b61014a6105dd565b604051600160a060020a03909116815260200160405180910390f35b61029033600035600160e060020a0319166105ec565b151561029857fe5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a25b5b50565b7f3a7f5663aac61ec71b13259e6a35978a2fba9256cd41bfacd1840a8b1c6bdd43838383604051600160a060020a0390931683526001608060020a039091166020830152600160e860020a0319166040808301919091526060909101905180910390a15b505050565b7f8147bc48a2d3d4b6f389e16b4c6abf7cc6312c957592235402941eb29219e54b426003548360405192835260208301919091526001608060020a03166040808301919091526060909101905180910390a15b50565b6103cd33600035600160e060020a0319166105ec565b15156103d557fe5b600081116103e257600080fd5b60038190555b5b50565b600254600160a060020a031681565b61041133600035600160e060020a0319166105ec565b151561041957fe5b600160a060020a038116151561042e57600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b5b50565b61047133600035600160e060020a0319166105ec565b151561047957fe5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167f1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada460405160405180910390a25b5b50565b600154600160a060020a031681565b7fc91056ac51b7d416daa4e0c747efe9f7d991055476c69ae0f67e4b8f6377c923338383604051600160a060020a0390931683526001608060020a039091166020830152600160e860020a0319166040808301919091526060909101905180910390a15b5050565b61056633600035600160e060020a0319166105ec565b151561056e57fe5b600254600160a060020a03166390bc16938260405160e060020a63ffffffff84160281526001608060020a039091166004820152602401600060405180830381600087803b15156105be57600080fd5b6102c65a03f115156105cf57600080fd5b5050505b5b50565b60035481565b600054600160a060020a031681565b600030600160a060020a031683600160a060020a0316141561061057506001610741565b600154600160a060020a0384811691161480156106365750600054600160a060020a0316155b1561064357506001610741565b600054600160a060020a031615156106ab577feb91385a0d70e44b915093d9ddfe6c8b41f2c56729431b1405f40aee5d874be78383604051600160a060020a039092168252600160e060020a03191660208201526040908101905180910390a1506000610741565b60008054600160a060020a03169063b7009613908590309086906040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152600160e060020a03199091166044820152606401602060405180830381600087803b151561072457600080fd5b6102c65a03f1151561073557600080fd5b50505060405180519150505b5b5b5b929150505600a165627a7a723058209953af89f72d0ef1def6efaa1178439eaa7516da480c286fce61bad0c66042a50029"`

// DeploySWCqueue deploys a new Ethereum contract, binding an instance of SWCqueue to it.
func DeploySWCqueue(auth *bind.TransactOpts, backend bind.ContractBackend, brg_ common.Address, nextBRGSWTratio_ *big.Int) (common.Address, *types.Transaction, *SWCqueue, error) {
	parsed, err := abi.JSON(strings.NewReader(SWCqueueABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SWCqueueBin), backend, brg_, nextBRGSWTratio_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SWCqueue{SWCqueueCaller: SWCqueueCaller{contract: contract}, SWCqueueTransactor: SWCqueueTransactor{contract: contract}}, nil
}

// SWCqueue is an auto generated Go binding around an Ethereum contract.
type SWCqueue struct {
	SWCqueueCaller     // Read-only binding to the contract
	SWCqueueTransactor // Write-only binding to the contract
}

// SWCqueueCaller is an auto generated read-only Go binding around an Ethereum contract.
type SWCqueueCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SWCqueueTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SWCqueueTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SWCqueueSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SWCqueueSession struct {
	Contract     *SWCqueue         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SWCqueueCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SWCqueueCallerSession struct {
	Contract *SWCqueueCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SWCqueueTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SWCqueueTransactorSession struct {
	Contract     *SWCqueueTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SWCqueueRaw is an auto generated low-level Go binding around an Ethereum contract.
type SWCqueueRaw struct {
	Contract *SWCqueue // Generic contract binding to access the raw methods on
}

// SWCqueueCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SWCqueueCallerRaw struct {
	Contract *SWCqueueCaller // Generic read-only contract binding to access the raw methods on
}

// SWCqueueTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SWCqueueTransactorRaw struct {
	Contract *SWCqueueTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSWCqueue creates a new instance of SWCqueue, bound to a specific deployed contract.
func NewSWCqueue(address common.Address, backend bind.ContractBackend) (*SWCqueue, error) {
	contract, err := bindSWCqueue(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SWCqueue{SWCqueueCaller: SWCqueueCaller{contract: contract}, SWCqueueTransactor: SWCqueueTransactor{contract: contract}}, nil
}

// NewSWCqueueCaller creates a new read-only instance of SWCqueue, bound to a specific deployed contract.
func NewSWCqueueCaller(address common.Address, caller bind.ContractCaller) (*SWCqueueCaller, error) {
	contract, err := bindSWCqueue(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &SWCqueueCaller{contract: contract}, nil
}

// NewSWCqueueTransactor creates a new write-only instance of SWCqueue, bound to a specific deployed contract.
func NewSWCqueueTransactor(address common.Address, transactor bind.ContractTransactor) (*SWCqueueTransactor, error) {
	contract, err := bindSWCqueue(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &SWCqueueTransactor{contract: contract}, nil
}

// bindSWCqueue binds a generic wrapper to an already deployed contract.
func bindSWCqueue(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SWCqueueABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SWCqueue *SWCqueueRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SWCqueue.Contract.SWCqueueCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SWCqueue *SWCqueueRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SWCqueue.Contract.SWCqueueTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SWCqueue *SWCqueueRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SWCqueue.Contract.SWCqueueTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SWCqueue *SWCqueueCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SWCqueue.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SWCqueue *SWCqueueTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SWCqueue.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SWCqueue *SWCqueueTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SWCqueue.Contract.contract.Transact(opts, method, params...)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_SWCqueue *SWCqueueCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SWCqueue.contract.Call(opts, out, "authority")
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_SWCqueue *SWCqueueSession) Authority() (common.Address, error) {
	return _SWCqueue.Contract.Authority(&_SWCqueue.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_SWCqueue *SWCqueueCallerSession) Authority() (common.Address, error) {
	return _SWCqueue.Contract.Authority(&_SWCqueue.CallOpts)
}

// Brg is a free data retrieval call binding the contract method 0x4f9c8fe8.
//
// Solidity: function brg() constant returns(address)
func (_SWCqueue *SWCqueueCaller) Brg(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SWCqueue.contract.Call(opts, out, "brg")
	return *ret0, err
}

// Brg is a free data retrieval call binding the contract method 0x4f9c8fe8.
//
// Solidity: function brg() constant returns(address)
func (_SWCqueue *SWCqueueSession) Brg() (common.Address, error) {
	return _SWCqueue.Contract.Brg(&_SWCqueue.CallOpts)
}

// Brg is a free data retrieval call binding the contract method 0x4f9c8fe8.
//
// Solidity: function brg() constant returns(address)
func (_SWCqueue *SWCqueueCallerSession) Brg() (common.Address, error) {
	return _SWCqueue.Contract.Brg(&_SWCqueue.CallOpts)
}

// NextBRGSWTratio is a free data retrieval call binding the contract method 0xac63f9d9.
//
// Solidity: function nextBRGSWTratio() constant returns(uint256)
func (_SWCqueue *SWCqueueCaller) NextBRGSWTratio(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SWCqueue.contract.Call(opts, out, "nextBRGSWTratio")
	return *ret0, err
}

// NextBRGSWTratio is a free data retrieval call binding the contract method 0xac63f9d9.
//
// Solidity: function nextBRGSWTratio() constant returns(uint256)
func (_SWCqueue *SWCqueueSession) NextBRGSWTratio() (*big.Int, error) {
	return _SWCqueue.Contract.NextBRGSWTratio(&_SWCqueue.CallOpts)
}

// NextBRGSWTratio is a free data retrieval call binding the contract method 0xac63f9d9.
//
// Solidity: function nextBRGSWTratio() constant returns(uint256)
func (_SWCqueue *SWCqueueCallerSession) NextBRGSWTratio() (*big.Int, error) {
	return _SWCqueue.Contract.NextBRGSWTratio(&_SWCqueue.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SWCqueue *SWCqueueCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SWCqueue.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SWCqueue *SWCqueueSession) Owner() (common.Address, error) {
	return _SWCqueue.Contract.Owner(&_SWCqueue.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SWCqueue *SWCqueueCallerSession) Owner() (common.Address, error) {
	return _SWCqueue.Contract.Owner(&_SWCqueue.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0x90bc1693.
//
// Solidity: function burn(brgAmount uint128) returns()
func (_SWCqueue *SWCqueueTransactor) Burn(opts *bind.TransactOpts, brgAmount *big.Int) (*types.Transaction, error) {
	return _SWCqueue.contract.Transact(opts, "burn", brgAmount)
}

// Burn is a paid mutator transaction binding the contract method 0x90bc1693.
//
// Solidity: function burn(brgAmount uint128) returns()
func (_SWCqueue *SWCqueueSession) Burn(brgAmount *big.Int) (*types.Transaction, error) {
	return _SWCqueue.Contract.Burn(&_SWCqueue.TransactOpts, brgAmount)
}

// Burn is a paid mutator transaction binding the contract method 0x90bc1693.
//
// Solidity: function burn(brgAmount uint128) returns()
func (_SWCqueue *SWCqueueTransactorSession) Burn(brgAmount *big.Int) (*types.Transaction, error) {
	return _SWCqueue.Contract.Burn(&_SWCqueue.TransactOpts, brgAmount)
}

// Cancel is a paid mutator transaction binding the contract method 0x90a160d4.
//
// Solidity: function cancel(wad uint128, currency bytes3) returns()
func (_SWCqueue *SWCqueueTransactor) Cancel(opts *bind.TransactOpts, wad *big.Int, currency [3]byte) (*types.Transaction, error) {
	return _SWCqueue.contract.Transact(opts, "cancel", wad, currency)
}

// Cancel is a paid mutator transaction binding the contract method 0x90a160d4.
//
// Solidity: function cancel(wad uint128, currency bytes3) returns()
func (_SWCqueue *SWCqueueSession) Cancel(wad *big.Int, currency [3]byte) (*types.Transaction, error) {
	return _SWCqueue.Contract.Cancel(&_SWCqueue.TransactOpts, wad, currency)
}

// Cancel is a paid mutator transaction binding the contract method 0x90a160d4.
//
// Solidity: function cancel(wad uint128, currency bytes3) returns()
func (_SWCqueue *SWCqueueTransactorSession) Cancel(wad *big.Int, currency [3]byte) (*types.Transaction, error) {
	return _SWCqueue.Contract.Cancel(&_SWCqueue.TransactOpts, wad, currency)
}

// DirectPledge is a paid mutator transaction binding the contract method 0x209c4986.
//
// Solidity: function directPledge(who address, wad uint128, currency bytes3) returns()
func (_SWCqueue *SWCqueueTransactor) DirectPledge(opts *bind.TransactOpts, who common.Address, wad *big.Int, currency [3]byte) (*types.Transaction, error) {
	return _SWCqueue.contract.Transact(opts, "directPledge", who, wad, currency)
}

// DirectPledge is a paid mutator transaction binding the contract method 0x209c4986.
//
// Solidity: function directPledge(who address, wad uint128, currency bytes3) returns()
func (_SWCqueue *SWCqueueSession) DirectPledge(who common.Address, wad *big.Int, currency [3]byte) (*types.Transaction, error) {
	return _SWCqueue.Contract.DirectPledge(&_SWCqueue.TransactOpts, who, wad, currency)
}

// DirectPledge is a paid mutator transaction binding the contract method 0x209c4986.
//
// Solidity: function directPledge(who address, wad uint128, currency bytes3) returns()
func (_SWCqueue *SWCqueueTransactorSession) DirectPledge(who common.Address, wad *big.Int, currency [3]byte) (*types.Transaction, error) {
	return _SWCqueue.Contract.DirectPledge(&_SWCqueue.TransactOpts, who, wad, currency)
}

// LogTrancheRelease is a paid mutator transaction binding the contract method 0x30311449.
//
// Solidity: function logTrancheRelease(wad uint128) returns()
func (_SWCqueue *SWCqueueTransactor) LogTrancheRelease(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _SWCqueue.contract.Transact(opts, "logTrancheRelease", wad)
}

// LogTrancheRelease is a paid mutator transaction binding the contract method 0x30311449.
//
// Solidity: function logTrancheRelease(wad uint128) returns()
func (_SWCqueue *SWCqueueSession) LogTrancheRelease(wad *big.Int) (*types.Transaction, error) {
	return _SWCqueue.Contract.LogTrancheRelease(&_SWCqueue.TransactOpts, wad)
}

// LogTrancheRelease is a paid mutator transaction binding the contract method 0x30311449.
//
// Solidity: function logTrancheRelease(wad uint128) returns()
func (_SWCqueue *SWCqueueTransactorSession) LogTrancheRelease(wad *big.Int) (*types.Transaction, error) {
	return _SWCqueue.Contract.LogTrancheRelease(&_SWCqueue.TransactOpts, wad)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_SWCqueue *SWCqueueTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _SWCqueue.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_SWCqueue *SWCqueueSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _SWCqueue.Contract.SetAuthority(&_SWCqueue.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_SWCqueue *SWCqueueTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _SWCqueue.Contract.SetAuthority(&_SWCqueue.TransactOpts, authority_)
}

// SetBRG is a paid mutator transaction binding the contract method 0x679ccc6d.
//
// Solidity: function setBRG(brg_ address) returns()
func (_SWCqueue *SWCqueueTransactor) SetBRG(opts *bind.TransactOpts, brg_ common.Address) (*types.Transaction, error) {
	return _SWCqueue.contract.Transact(opts, "setBRG", brg_)
}

// SetBRG is a paid mutator transaction binding the contract method 0x679ccc6d.
//
// Solidity: function setBRG(brg_ address) returns()
func (_SWCqueue *SWCqueueSession) SetBRG(brg_ common.Address) (*types.Transaction, error) {
	return _SWCqueue.Contract.SetBRG(&_SWCqueue.TransactOpts, brg_)
}

// SetBRG is a paid mutator transaction binding the contract method 0x679ccc6d.
//
// Solidity: function setBRG(brg_ address) returns()
func (_SWCqueue *SWCqueueTransactorSession) SetBRG(brg_ common.Address) (*types.Transaction, error) {
	return _SWCqueue.Contract.SetBRG(&_SWCqueue.TransactOpts, brg_)
}

// SetNextBRGSWTratio is a paid mutator transaction binding the contract method 0x429ad4f2.
//
// Solidity: function setNextBRGSWTratio(ratio uint256) returns()
func (_SWCqueue *SWCqueueTransactor) SetNextBRGSWTratio(opts *bind.TransactOpts, ratio *big.Int) (*types.Transaction, error) {
	return _SWCqueue.contract.Transact(opts, "setNextBRGSWTratio", ratio)
}

// SetNextBRGSWTratio is a paid mutator transaction binding the contract method 0x429ad4f2.
//
// Solidity: function setNextBRGSWTratio(ratio uint256) returns()
func (_SWCqueue *SWCqueueSession) SetNextBRGSWTratio(ratio *big.Int) (*types.Transaction, error) {
	return _SWCqueue.Contract.SetNextBRGSWTratio(&_SWCqueue.TransactOpts, ratio)
}

// SetNextBRGSWTratio is a paid mutator transaction binding the contract method 0x429ad4f2.
//
// Solidity: function setNextBRGSWTratio(ratio uint256) returns()
func (_SWCqueue *SWCqueueTransactorSession) SetNextBRGSWTratio(ratio *big.Int) (*types.Transaction, error) {
	return _SWCqueue.Contract.SetNextBRGSWTratio(&_SWCqueue.TransactOpts, ratio)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_SWCqueue *SWCqueueTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _SWCqueue.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_SWCqueue *SWCqueueSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _SWCqueue.Contract.SetOwner(&_SWCqueue.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_SWCqueue *SWCqueueTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _SWCqueue.Contract.SetOwner(&_SWCqueue.TransactOpts, owner_)
}

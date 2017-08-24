// This file is an automatically generated Go binding. Do not modify as any
// change will likely be lost upon the next re-generation!

package contracts

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// TokenDataABI is the input ABI used to generate the binding from.
const TokenDataABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"supply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guy\",\"type\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"setBalances\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"logic_\",\"type\":\"address\"}],\"name\":\"setTokenLogic\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"supply_\",\"type\":\"uint256\"}],\"name\":\"setSupply\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"address\"}],\"name\":\"approvals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"guy\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"setApprovals\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"logic_\",\"type\":\"address\"},{\"name\":\"supply_\",\"type\":\"uint256\"},{\"name\":\"owner_\",\"type\":\"address\"}],\"payable\":false,\"type\":\"constructor\"}]"

// TokenDataBin is the compiled bytecode used for deploying new contracts.
const TokenDataBin = `"0x6060604052341561000f57600080fd5b6040516060806103b68339810160405280805191906020018051919060200180519150505b60038054600160a060020a03808616600160a060020a03199283161790925560008481556004805485851693169290921791829055911681526001602052604090208290555b5050505b6103298061008d6000396000f3006060604052361561008b5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663047fc9aa811461009057806327e235e3146100b557806328e69b16146100e65780633a3d523f1461010a5780633b4c4b251461012b5780638da5cb5b14610143578063a32ce11e14610172578063e32d5cf8146101a9575b600080fd5b341561009b57600080fd5b6100a36101d3565b60405190815260200160405180910390f35b34156100c057600080fd5b6100a3600160a060020a03600435166101d9565b60405190815260200160405180910390f35b34156100f157600080fd5b610108600160a060020a03600435166024356101eb565b005b341561011557600080fd5b610108600160a060020a0360043516610224565b005b341561013657600080fd5b610108600435610268565b005b341561014e57600080fd5b61015661028a565b604051600160a060020a03909116815260200160405180910390f35b341561017d57600080fd5b6100a3600160a060020a0360043581169060243516610299565b60405190815260200160405180910390f35b34156101b457600080fd5b610108600160a060020a03600435811690602435166044356102b6565b005b60005481565b60016020526000908152604090205481565b60035433600160a060020a0390811691161461020357fe5b600160a060020a03821660009081526001602052604090208190555b5b5050565b60045433600160a060020a0390811691161461023c57fe5b6003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b50565b60035433600160a060020a0390811691161461028057fe5b60008190555b5b50565b600454600160a060020a031681565b600260209081526000928352604080842090915290825290205481565b60035433600160a060020a039081169116146102ce57fe5b600160a060020a0380841660009081526002602090815260408083209386168352929052208190555b5b5050505600a165627a7a72305820b555670b782a6fbd07366fb67f18ffc34a831f86e97df3126bb5cf286ddfb1740029"`

// DeployTokenData deploys a new Ethereum contract, binding an instance of TokenData to it.
func DeployTokenData(auth *bind.TransactOpts, backend bind.ContractBackend, logic_ common.Address, supply_ *big.Int, owner_ common.Address) (common.Address, *types.Transaction, *TokenData, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenDataABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenDataBin), backend, logic_, supply_, owner_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenData{TokenDataCaller: TokenDataCaller{contract: contract}, TokenDataTransactor: TokenDataTransactor{contract: contract}}, nil
}

// TokenData is an auto generated Go binding around an Ethereum contract.
type TokenData struct {
	TokenDataCaller     // Read-only binding to the contract
	TokenDataTransactor // Write-only binding to the contract
}

// TokenDataCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenDataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenDataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenDataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenDataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenDataSession struct {
	Contract     *TokenData        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenDataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenDataCallerSession struct {
	Contract *TokenDataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TokenDataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenDataTransactorSession struct {
	Contract     *TokenDataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TokenDataRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenDataRaw struct {
	Contract *TokenData // Generic contract binding to access the raw methods on
}

// TokenDataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenDataCallerRaw struct {
	Contract *TokenDataCaller // Generic read-only contract binding to access the raw methods on
}

// TokenDataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenDataTransactorRaw struct {
	Contract *TokenDataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenData creates a new instance of TokenData, bound to a specific deployed contract.
func NewTokenData(address common.Address, backend bind.ContractBackend) (*TokenData, error) {
	contract, err := bindTokenData(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenData{TokenDataCaller: TokenDataCaller{contract: contract}, TokenDataTransactor: TokenDataTransactor{contract: contract}}, nil
}

// NewTokenDataCaller creates a new read-only instance of TokenData, bound to a specific deployed contract.
func NewTokenDataCaller(address common.Address, caller bind.ContractCaller) (*TokenDataCaller, error) {
	contract, err := bindTokenData(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TokenDataCaller{contract: contract}, nil
}

// NewTokenDataTransactor creates a new write-only instance of TokenData, bound to a specific deployed contract.
func NewTokenDataTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenDataTransactor, error) {
	contract, err := bindTokenData(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TokenDataTransactor{contract: contract}, nil
}

// bindTokenData binds a generic wrapper to an already deployed contract.
func bindTokenData(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenDataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenData *TokenDataRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenData.Contract.TokenDataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenData *TokenDataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenData.Contract.TokenDataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenData *TokenDataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenData.Contract.TokenDataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenData *TokenDataCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenData.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenData *TokenDataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenData.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenData *TokenDataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenData.Contract.contract.Transact(opts, method, params...)
}

// Approvals is a free data retrieval call binding the contract method 0xa32ce11e.
//
// Solidity: function approvals( address,  address) constant returns(uint256)
func (_TokenData *TokenDataCaller) Approvals(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenData.contract.Call(opts, out, "approvals", arg0, arg1)
	return *ret0, err
}

// Approvals is a free data retrieval call binding the contract method 0xa32ce11e.
//
// Solidity: function approvals( address,  address) constant returns(uint256)
func (_TokenData *TokenDataSession) Approvals(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TokenData.Contract.Approvals(&_TokenData.CallOpts, arg0, arg1)
}

// Approvals is a free data retrieval call binding the contract method 0xa32ce11e.
//
// Solidity: function approvals( address,  address) constant returns(uint256)
func (_TokenData *TokenDataCallerSession) Approvals(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _TokenData.Contract.Approvals(&_TokenData.CallOpts, arg0, arg1)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_TokenData *TokenDataCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenData.contract.Call(opts, out, "balances", arg0)
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_TokenData *TokenDataSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _TokenData.Contract.Balances(&_TokenData.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances( address) constant returns(uint256)
func (_TokenData *TokenDataCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _TokenData.Contract.Balances(&_TokenData.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenData *TokenDataCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _TokenData.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenData *TokenDataSession) Owner() (common.Address, error) {
	return _TokenData.Contract.Owner(&_TokenData.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_TokenData *TokenDataCallerSession) Owner() (common.Address, error) {
	return _TokenData.Contract.Owner(&_TokenData.CallOpts)
}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() constant returns(uint256)
func (_TokenData *TokenDataCaller) Supply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _TokenData.contract.Call(opts, out, "supply")
	return *ret0, err
}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() constant returns(uint256)
func (_TokenData *TokenDataSession) Supply() (*big.Int, error) {
	return _TokenData.Contract.Supply(&_TokenData.CallOpts)
}

// Supply is a free data retrieval call binding the contract method 0x047fc9aa.
//
// Solidity: function supply() constant returns(uint256)
func (_TokenData *TokenDataCallerSession) Supply() (*big.Int, error) {
	return _TokenData.Contract.Supply(&_TokenData.CallOpts)
}

// SetApprovals is a paid mutator transaction binding the contract method 0xe32d5cf8.
//
// Solidity: function setApprovals(src address, guy address, wad uint256) returns()
func (_TokenData *TokenDataTransactor) SetApprovals(opts *bind.TransactOpts, src common.Address, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _TokenData.contract.Transact(opts, "setApprovals", src, guy, wad)
}

// SetApprovals is a paid mutator transaction binding the contract method 0xe32d5cf8.
//
// Solidity: function setApprovals(src address, guy address, wad uint256) returns()
func (_TokenData *TokenDataSession) SetApprovals(src common.Address, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _TokenData.Contract.SetApprovals(&_TokenData.TransactOpts, src, guy, wad)
}

// SetApprovals is a paid mutator transaction binding the contract method 0xe32d5cf8.
//
// Solidity: function setApprovals(src address, guy address, wad uint256) returns()
func (_TokenData *TokenDataTransactorSession) SetApprovals(src common.Address, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _TokenData.Contract.SetApprovals(&_TokenData.TransactOpts, src, guy, wad)
}

// SetBalances is a paid mutator transaction binding the contract method 0x28e69b16.
//
// Solidity: function setBalances(guy address, balance uint256) returns()
func (_TokenData *TokenDataTransactor) SetBalances(opts *bind.TransactOpts, guy common.Address, balance *big.Int) (*types.Transaction, error) {
	return _TokenData.contract.Transact(opts, "setBalances", guy, balance)
}

// SetBalances is a paid mutator transaction binding the contract method 0x28e69b16.
//
// Solidity: function setBalances(guy address, balance uint256) returns()
func (_TokenData *TokenDataSession) SetBalances(guy common.Address, balance *big.Int) (*types.Transaction, error) {
	return _TokenData.Contract.SetBalances(&_TokenData.TransactOpts, guy, balance)
}

// SetBalances is a paid mutator transaction binding the contract method 0x28e69b16.
//
// Solidity: function setBalances(guy address, balance uint256) returns()
func (_TokenData *TokenDataTransactorSession) SetBalances(guy common.Address, balance *big.Int) (*types.Transaction, error) {
	return _TokenData.Contract.SetBalances(&_TokenData.TransactOpts, guy, balance)
}

// SetSupply is a paid mutator transaction binding the contract method 0x3b4c4b25.
//
// Solidity: function setSupply(supply_ uint256) returns()
func (_TokenData *TokenDataTransactor) SetSupply(opts *bind.TransactOpts, supply_ *big.Int) (*types.Transaction, error) {
	return _TokenData.contract.Transact(opts, "setSupply", supply_)
}

// SetSupply is a paid mutator transaction binding the contract method 0x3b4c4b25.
//
// Solidity: function setSupply(supply_ uint256) returns()
func (_TokenData *TokenDataSession) SetSupply(supply_ *big.Int) (*types.Transaction, error) {
	return _TokenData.Contract.SetSupply(&_TokenData.TransactOpts, supply_)
}

// SetSupply is a paid mutator transaction binding the contract method 0x3b4c4b25.
//
// Solidity: function setSupply(supply_ uint256) returns()
func (_TokenData *TokenDataTransactorSession) SetSupply(supply_ *big.Int) (*types.Transaction, error) {
	return _TokenData.Contract.SetSupply(&_TokenData.TransactOpts, supply_)
}

// SetTokenLogic is a paid mutator transaction binding the contract method 0x3a3d523f.
//
// Solidity: function setTokenLogic(logic_ address) returns()
func (_TokenData *TokenDataTransactor) SetTokenLogic(opts *bind.TransactOpts, logic_ common.Address) (*types.Transaction, error) {
	return _TokenData.contract.Transact(opts, "setTokenLogic", logic_)
}

// SetTokenLogic is a paid mutator transaction binding the contract method 0x3a3d523f.
//
// Solidity: function setTokenLogic(logic_ address) returns()
func (_TokenData *TokenDataSession) SetTokenLogic(logic_ common.Address) (*types.Transaction, error) {
	return _TokenData.Contract.SetTokenLogic(&_TokenData.TransactOpts, logic_)
}

// SetTokenLogic is a paid mutator transaction binding the contract method 0x3a3d523f.
//
// Solidity: function setTokenLogic(logic_ address) returns()
func (_TokenData *TokenDataTransactorSession) SetTokenLogic(logic_ common.Address) (*types.Transaction, error) {
	return _TokenData.Contract.SetTokenLogic(&_TokenData.TransactOpts, logic_)
}

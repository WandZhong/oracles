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

// AuthorityABI is the input ABI used to generate the binding from.
const AuthorityABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"canCall\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"}]"

// AuthorityBin is the compiled bytecode used for deploying new contracts.
const AuthorityBin = `"0x"`

// DeployAuthority deploys a new Ethereum contract, binding an instance of Authority to it.
func DeployAuthority(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Authority, error) {
	parsed, err := abi.JSON(strings.NewReader(AuthorityABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AuthorityBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Authority{AuthorityCaller: AuthorityCaller{contract: contract}, AuthorityTransactor: AuthorityTransactor{contract: contract}}, nil
}

// Authority is an auto generated Go binding around an Ethereum contract.
type Authority struct {
	AuthorityCaller     // Read-only binding to the contract
	AuthorityTransactor // Write-only binding to the contract
}

// AuthorityCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuthorityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthorityTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuthorityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuthoritySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuthoritySession struct {
	Contract     *Authority        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuthorityCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuthorityCallerSession struct {
	Contract *AuthorityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// AuthorityTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuthorityTransactorSession struct {
	Contract     *AuthorityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// AuthorityRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuthorityRaw struct {
	Contract *Authority // Generic contract binding to access the raw methods on
}

// AuthorityCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuthorityCallerRaw struct {
	Contract *AuthorityCaller // Generic read-only contract binding to access the raw methods on
}

// AuthorityTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuthorityTransactorRaw struct {
	Contract *AuthorityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuthority creates a new instance of Authority, bound to a specific deployed contract.
func NewAuthority(address common.Address, backend bind.ContractBackend) (*Authority, error) {
	contract, err := bindAuthority(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Authority{AuthorityCaller: AuthorityCaller{contract: contract}, AuthorityTransactor: AuthorityTransactor{contract: contract}}, nil
}

// NewAuthorityCaller creates a new read-only instance of Authority, bound to a specific deployed contract.
func NewAuthorityCaller(address common.Address, caller bind.ContractCaller) (*AuthorityCaller, error) {
	contract, err := bindAuthority(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &AuthorityCaller{contract: contract}, nil
}

// NewAuthorityTransactor creates a new write-only instance of Authority, bound to a specific deployed contract.
func NewAuthorityTransactor(address common.Address, transactor bind.ContractTransactor) (*AuthorityTransactor, error) {
	contract, err := bindAuthority(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &AuthorityTransactor{contract: contract}, nil
}

// bindAuthority binds a generic wrapper to an already deployed contract.
func bindAuthority(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AuthorityABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Authority *AuthorityRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Authority.Contract.AuthorityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Authority *AuthorityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Authority.Contract.AuthorityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Authority *AuthorityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Authority.Contract.AuthorityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Authority *AuthorityCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Authority.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Authority *AuthorityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Authority.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Authority *AuthorityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Authority.Contract.contract.Transact(opts, method, params...)
}

// CanCall is a free data retrieval call binding the contract method 0xb7009613.
//
// Solidity: function canCall(src address, dst address, sig bytes4) constant returns(bool)
func (_Authority *AuthorityCaller) CanCall(opts *bind.CallOpts, src common.Address, dst common.Address, sig [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Authority.contract.Call(opts, out, "canCall", src, dst, sig)
	return *ret0, err
}

// CanCall is a free data retrieval call binding the contract method 0xb7009613.
//
// Solidity: function canCall(src address, dst address, sig bytes4) constant returns(bool)
func (_Authority *AuthoritySession) CanCall(src common.Address, dst common.Address, sig [4]byte) (bool, error) {
	return _Authority.Contract.CanCall(&_Authority.CallOpts, src, dst, sig)
}

// CanCall is a free data retrieval call binding the contract method 0xb7009613.
//
// Solidity: function canCall(src address, dst address, sig bytes4) constant returns(bool)
func (_Authority *AuthorityCallerSession) CanCall(src common.Address, dst common.Address, sig [4]byte) (bool, error) {
	return _Authority.Contract.CanCall(&_Authority.CallOpts, src, dst, sig)
}

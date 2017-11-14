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

// ForwarderFactoryABI is the input ABI used to generate the binding from.
const ForwarderFactoryABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"destination\",\"type\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"emitLogPaymentForwarded\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"senderHasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"restart\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"roles\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"destination\",\"type\":\"address\"}],\"name\":\"createForwarder\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"roles_\",\"type\":\"address\"}],\"name\":\"setRolesContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"hasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"rolesContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"forwarder\",\"type\":\"address\"}],\"name\":\"LogForwarderCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"destination\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LogPaymentForwarded\",\"type\":\"event\"}]"

// ForwarderFactoryBin is the compiled bytecode used for deploying new contracts.
const ForwarderFactoryBin = `"0x60606040526000600360006101000a81548160ff021916908315150217905550341561002a57600080fd5b6040516020806112a5833981016040528080519060200190919050506040805190810160405280601081526020017f466f72776172646572466163746f72790000000000000000000000000000000081525081336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550816040518082805190602001908083835b6020831015156100f357805182526020820191506020810190506020830392506100ce565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390206002816000191690555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050506111288061017d6000396000f3006060604052600436106100ba576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806306d7561f146100bf57806307da68f51461012057806313af4035146101355780631ca03b8e1461016e5780631ef3755d146101e3578063392f5f64146101f857806375f12b211461024d5780638da5cb5b1461027a578063904c6094146102cf5780639193ba0b14610300578063afa202ac14610339578063e3c33a9b14610372575b600080fd5b34156100ca57600080fd5b61011e600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803573ffffffffffffffffffffffffffffffffffffffff169060200190919080359060200190919050506103e7565b005b341561012b57600080fd5b6101336104e3565b005b341561014057600080fd5b61016c600480803573ffffffffffffffffffffffffffffffffffffffff169060200190919050506105a2565b005b341561017957600080fd5b6101c9600480803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919050506106c4565b604051808215151515815260200191505060405180910390f35b34156101ee57600080fd5b6101f6610842565b005b341561020357600080fd5b61020b610901565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561025857600080fd5b610260610927565b604051808215151515815260200191505060405180910390f35b341561028557600080fd5b61028d61093a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156102da57600080fd5b6102e261095f565b60405180826000191660001916815260200191505060405180910390f35b341561030b57600080fd5b610337600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610965565b005b341561034457600080fd5b610370600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050610b1f565b005b341561037d57600080fd5b6103cd600480803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050610c1b565b604051808215151515815260200191505060405180910390f35b600460003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151561043f57600080fd5b7f112ccbc59137582755de13aadd93d32eaa2dff15e276915511f71012d2871f44838383604051808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001935050505060405180910390a1505050565b6040805190810160405280600781526020017f73746f70706572000000000000000000000000000000000000000000000000008152506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614806105795750610578816106c4565b5b151561058457600080fd5b6001600360006101000a81548160ff02191690831515021790555050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156105fd57600080fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed946000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b60006106cf82610c1b565b801561083b5750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638b51ca42600254846040518082805190602001908083835b60208310151561074d5780518252602082019150602081019050602083039250610728565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020336000604051602001526040518463ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180846000191660001916815260200183600019166000191681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019350505050602060405180830381600087803b151561081f57600080fd5b6102c65a03f1151561083057600080fd5b505050604051805190505b9050919050565b6040805190810160405280600981526020017f72657374617274657200000000000000000000000000000000000000000000008152506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614806108d857506108d7816106c4565b5b15156108e357600080fd5b6000600360006101000a81548160ff02191690831515021790555050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600360009054906101000a900460ff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60025481565b60006040805190810160405280600581526020017f61646d696e0000000000000000000000000000000000000000000000000000008152506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614806109fd57506109fc816106c4565b5b1515610a0857600080fd5b82610a11610d54565b808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001915050604051809103906000f0801515610a5d57600080fd5b91506001600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055507fb98c7eedd7391b63c06f053a9d009a25952086d6ac2c28a62c3673b9a8aa32c582604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a1505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610b7a57600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614151515610bd757600080fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633037cea3600254846040518082805190602001908083835b602083101515610c945780518252602082019150602081019050602083039250610c6f565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390206000604051602001526040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808360001916600019168152602001826000191660001916815260200192505050602060405180830381600087803b1515610d3257600080fd5b6102c65a03f11515610d4357600080fd5b505050604051805190509050919050565b60405161039880610d658339019056006060604052341561000f57600080fd5b6040516020806103988339810160405280805190602001909190505080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506102dc806100bc6000396000f30060606040526004361061004c576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063b269681d146101bb578063c45a015514610210575b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f1935050505015156100ae57600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166306d7561f3330346040518463ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050600060405180830381600087803b15156101a557600080fd5b6102c65a03f115156101b657600080fd5b505050005b34156101c657600080fd5b6101ce610265565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561021b57600080fd5b61022361028b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff16815600a165627a7a72305820b8a80f0e55cea09a86ef6d3b41770d085f58a36b0859f22fd350ad491ad6035c0029a165627a7a72305820b42a07aa537ae867467398f33652334aaed90aed7657a6e602146c5b743758f10029"`

// DeployForwarderFactory deploys a new Ethereum contract, binding an instance of ForwarderFactory to it.
func DeployForwarderFactory(auth *bind.TransactOpts, backend bind.ContractBackend, rolesContract common.Address) (common.Address, *types.Transaction, *ForwarderFactory, error) {
	parsed, err := abi.JSON(strings.NewReader(ForwarderFactoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ForwarderFactoryBin), backend, rolesContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ForwarderFactory{ForwarderFactoryCaller: ForwarderFactoryCaller{contract: contract}, ForwarderFactoryTransactor: ForwarderFactoryTransactor{contract: contract}}, nil
}

// ForwarderFactory is an auto generated Go binding around an Ethereum contract.
type ForwarderFactory struct {
	ForwarderFactoryCaller     // Read-only binding to the contract
	ForwarderFactoryTransactor // Write-only binding to the contract
}

// ForwarderFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type ForwarderFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForwarderFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ForwarderFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ForwarderFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ForwarderFactorySession struct {
	Contract     *ForwarderFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ForwarderFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ForwarderFactoryCallerSession struct {
	Contract *ForwarderFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ForwarderFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ForwarderFactoryTransactorSession struct {
	Contract     *ForwarderFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ForwarderFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type ForwarderFactoryRaw struct {
	Contract *ForwarderFactory // Generic contract binding to access the raw methods on
}

// ForwarderFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ForwarderFactoryCallerRaw struct {
	Contract *ForwarderFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// ForwarderFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ForwarderFactoryTransactorRaw struct {
	Contract *ForwarderFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewForwarderFactory creates a new instance of ForwarderFactory, bound to a specific deployed contract.
func NewForwarderFactory(address common.Address, backend bind.ContractBackend) (*ForwarderFactory, error) {
	contract, err := bindForwarderFactory(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ForwarderFactory{ForwarderFactoryCaller: ForwarderFactoryCaller{contract: contract}, ForwarderFactoryTransactor: ForwarderFactoryTransactor{contract: contract}}, nil
}

// NewForwarderFactoryCaller creates a new read-only instance of ForwarderFactory, bound to a specific deployed contract.
func NewForwarderFactoryCaller(address common.Address, caller bind.ContractCaller) (*ForwarderFactoryCaller, error) {
	contract, err := bindForwarderFactory(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &ForwarderFactoryCaller{contract: contract}, nil
}

// NewForwarderFactoryTransactor creates a new write-only instance of ForwarderFactory, bound to a specific deployed contract.
func NewForwarderFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*ForwarderFactoryTransactor, error) {
	contract, err := bindForwarderFactory(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &ForwarderFactoryTransactor{contract: contract}, nil
}

// bindForwarderFactory binds a generic wrapper to an already deployed contract.
func bindForwarderFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ForwarderFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ForwarderFactory *ForwarderFactoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ForwarderFactory.Contract.ForwarderFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ForwarderFactory *ForwarderFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ForwarderFactory.Contract.ForwarderFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ForwarderFactory *ForwarderFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ForwarderFactory.Contract.ForwarderFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ForwarderFactory *ForwarderFactoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ForwarderFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ForwarderFactory *ForwarderFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ForwarderFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ForwarderFactory *ForwarderFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ForwarderFactory.Contract.contract.Transact(opts, method, params...)
}

// ContractHash is a free data retrieval call binding the contract method 0x904c6094.
//
// Solidity: function contractHash() constant returns(bytes32)
func (_ForwarderFactory *ForwarderFactoryCaller) ContractHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _ForwarderFactory.contract.Call(opts, out, "contractHash")
	return *ret0, err
}

// ContractHash is a free data retrieval call binding the contract method 0x904c6094.
//
// Solidity: function contractHash() constant returns(bytes32)
func (_ForwarderFactory *ForwarderFactorySession) ContractHash() ([32]byte, error) {
	return _ForwarderFactory.Contract.ContractHash(&_ForwarderFactory.CallOpts)
}

// ContractHash is a free data retrieval call binding the contract method 0x904c6094.
//
// Solidity: function contractHash() constant returns(bytes32)
func (_ForwarderFactory *ForwarderFactoryCallerSession) ContractHash() ([32]byte, error) {
	return _ForwarderFactory.Contract.ContractHash(&_ForwarderFactory.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0xe3c33a9b.
//
// Solidity: function hasRole(roleName string) constant returns(bool)
func (_ForwarderFactory *ForwarderFactoryCaller) HasRole(opts *bind.CallOpts, roleName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ForwarderFactory.contract.Call(opts, out, "hasRole", roleName)
	return *ret0, err
}

// HasRole is a free data retrieval call binding the contract method 0xe3c33a9b.
//
// Solidity: function hasRole(roleName string) constant returns(bool)
func (_ForwarderFactory *ForwarderFactorySession) HasRole(roleName string) (bool, error) {
	return _ForwarderFactory.Contract.HasRole(&_ForwarderFactory.CallOpts, roleName)
}

// HasRole is a free data retrieval call binding the contract method 0xe3c33a9b.
//
// Solidity: function hasRole(roleName string) constant returns(bool)
func (_ForwarderFactory *ForwarderFactoryCallerSession) HasRole(roleName string) (bool, error) {
	return _ForwarderFactory.Contract.HasRole(&_ForwarderFactory.CallOpts, roleName)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ForwarderFactory *ForwarderFactoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ForwarderFactory.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ForwarderFactory *ForwarderFactorySession) Owner() (common.Address, error) {
	return _ForwarderFactory.Contract.Owner(&_ForwarderFactory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_ForwarderFactory *ForwarderFactoryCallerSession) Owner() (common.Address, error) {
	return _ForwarderFactory.Contract.Owner(&_ForwarderFactory.CallOpts)
}

// Roles is a free data retrieval call binding the contract method 0x392f5f64.
//
// Solidity: function roles() constant returns(address)
func (_ForwarderFactory *ForwarderFactoryCaller) Roles(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _ForwarderFactory.contract.Call(opts, out, "roles")
	return *ret0, err
}

// Roles is a free data retrieval call binding the contract method 0x392f5f64.
//
// Solidity: function roles() constant returns(address)
func (_ForwarderFactory *ForwarderFactorySession) Roles() (common.Address, error) {
	return _ForwarderFactory.Contract.Roles(&_ForwarderFactory.CallOpts)
}

// Roles is a free data retrieval call binding the contract method 0x392f5f64.
//
// Solidity: function roles() constant returns(address)
func (_ForwarderFactory *ForwarderFactoryCallerSession) Roles() (common.Address, error) {
	return _ForwarderFactory.Contract.Roles(&_ForwarderFactory.CallOpts)
}

// SenderHasRole is a free data retrieval call binding the contract method 0x1ca03b8e.
//
// Solidity: function senderHasRole(roleName string) constant returns(bool)
func (_ForwarderFactory *ForwarderFactoryCaller) SenderHasRole(opts *bind.CallOpts, roleName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ForwarderFactory.contract.Call(opts, out, "senderHasRole", roleName)
	return *ret0, err
}

// SenderHasRole is a free data retrieval call binding the contract method 0x1ca03b8e.
//
// Solidity: function senderHasRole(roleName string) constant returns(bool)
func (_ForwarderFactory *ForwarderFactorySession) SenderHasRole(roleName string) (bool, error) {
	return _ForwarderFactory.Contract.SenderHasRole(&_ForwarderFactory.CallOpts, roleName)
}

// SenderHasRole is a free data retrieval call binding the contract method 0x1ca03b8e.
//
// Solidity: function senderHasRole(roleName string) constant returns(bool)
func (_ForwarderFactory *ForwarderFactoryCallerSession) SenderHasRole(roleName string) (bool, error) {
	return _ForwarderFactory.Contract.SenderHasRole(&_ForwarderFactory.CallOpts, roleName)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_ForwarderFactory *ForwarderFactoryCaller) Stopped(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _ForwarderFactory.contract.Call(opts, out, "stopped")
	return *ret0, err
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_ForwarderFactory *ForwarderFactorySession) Stopped() (bool, error) {
	return _ForwarderFactory.Contract.Stopped(&_ForwarderFactory.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_ForwarderFactory *ForwarderFactoryCallerSession) Stopped() (bool, error) {
	return _ForwarderFactory.Contract.Stopped(&_ForwarderFactory.CallOpts)
}

// CreateForwarder is a paid mutator transaction binding the contract method 0x9193ba0b.
//
// Solidity: function createForwarder(destination address) returns()
func (_ForwarderFactory *ForwarderFactoryTransactor) CreateForwarder(opts *bind.TransactOpts, destination common.Address) (*types.Transaction, error) {
	return _ForwarderFactory.contract.Transact(opts, "createForwarder", destination)
}

// CreateForwarder is a paid mutator transaction binding the contract method 0x9193ba0b.
//
// Solidity: function createForwarder(destination address) returns()
func (_ForwarderFactory *ForwarderFactorySession) CreateForwarder(destination common.Address) (*types.Transaction, error) {
	return _ForwarderFactory.Contract.CreateForwarder(&_ForwarderFactory.TransactOpts, destination)
}

// CreateForwarder is a paid mutator transaction binding the contract method 0x9193ba0b.
//
// Solidity: function createForwarder(destination address) returns()
func (_ForwarderFactory *ForwarderFactoryTransactorSession) CreateForwarder(destination common.Address) (*types.Transaction, error) {
	return _ForwarderFactory.Contract.CreateForwarder(&_ForwarderFactory.TransactOpts, destination)
}

// EmitLogPaymentForwarded is a paid mutator transaction binding the contract method 0x06d7561f.
//
// Solidity: function emitLogPaymentForwarded(from address, destination address, amount uint256) returns()
func (_ForwarderFactory *ForwarderFactoryTransactor) EmitLogPaymentForwarded(opts *bind.TransactOpts, from common.Address, destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ForwarderFactory.contract.Transact(opts, "emitLogPaymentForwarded", from, destination, amount)
}

// EmitLogPaymentForwarded is a paid mutator transaction binding the contract method 0x06d7561f.
//
// Solidity: function emitLogPaymentForwarded(from address, destination address, amount uint256) returns()
func (_ForwarderFactory *ForwarderFactorySession) EmitLogPaymentForwarded(from common.Address, destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ForwarderFactory.Contract.EmitLogPaymentForwarded(&_ForwarderFactory.TransactOpts, from, destination, amount)
}

// EmitLogPaymentForwarded is a paid mutator transaction binding the contract method 0x06d7561f.
//
// Solidity: function emitLogPaymentForwarded(from address, destination address, amount uint256) returns()
func (_ForwarderFactory *ForwarderFactoryTransactorSession) EmitLogPaymentForwarded(from common.Address, destination common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ForwarderFactory.Contract.EmitLogPaymentForwarded(&_ForwarderFactory.TransactOpts, from, destination, amount)
}

// Restart is a paid mutator transaction binding the contract method 0x1ef3755d.
//
// Solidity: function restart() returns()
func (_ForwarderFactory *ForwarderFactoryTransactor) Restart(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ForwarderFactory.contract.Transact(opts, "restart")
}

// Restart is a paid mutator transaction binding the contract method 0x1ef3755d.
//
// Solidity: function restart() returns()
func (_ForwarderFactory *ForwarderFactorySession) Restart() (*types.Transaction, error) {
	return _ForwarderFactory.Contract.Restart(&_ForwarderFactory.TransactOpts)
}

// Restart is a paid mutator transaction binding the contract method 0x1ef3755d.
//
// Solidity: function restart() returns()
func (_ForwarderFactory *ForwarderFactoryTransactorSession) Restart() (*types.Transaction, error) {
	return _ForwarderFactory.Contract.Restart(&_ForwarderFactory.TransactOpts)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_ForwarderFactory *ForwarderFactoryTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _ForwarderFactory.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_ForwarderFactory *ForwarderFactorySession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _ForwarderFactory.Contract.SetOwner(&_ForwarderFactory.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_ForwarderFactory *ForwarderFactoryTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _ForwarderFactory.Contract.SetOwner(&_ForwarderFactory.TransactOpts, owner_)
}

// SetRolesContract is a paid mutator transaction binding the contract method 0xafa202ac.
//
// Solidity: function setRolesContract(roles_ address) returns()
func (_ForwarderFactory *ForwarderFactoryTransactor) SetRolesContract(opts *bind.TransactOpts, roles_ common.Address) (*types.Transaction, error) {
	return _ForwarderFactory.contract.Transact(opts, "setRolesContract", roles_)
}

// SetRolesContract is a paid mutator transaction binding the contract method 0xafa202ac.
//
// Solidity: function setRolesContract(roles_ address) returns()
func (_ForwarderFactory *ForwarderFactorySession) SetRolesContract(roles_ common.Address) (*types.Transaction, error) {
	return _ForwarderFactory.Contract.SetRolesContract(&_ForwarderFactory.TransactOpts, roles_)
}

// SetRolesContract is a paid mutator transaction binding the contract method 0xafa202ac.
//
// Solidity: function setRolesContract(roles_ address) returns()
func (_ForwarderFactory *ForwarderFactoryTransactorSession) SetRolesContract(roles_ common.Address) (*types.Transaction, error) {
	return _ForwarderFactory.Contract.SetRolesContract(&_ForwarderFactory.TransactOpts, roles_)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_ForwarderFactory *ForwarderFactoryTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ForwarderFactory.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_ForwarderFactory *ForwarderFactorySession) Stop() (*types.Transaction, error) {
	return _ForwarderFactory.Contract.Stop(&_ForwarderFactory.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_ForwarderFactory *ForwarderFactoryTransactorSession) Stop() (*types.Transaction, error) {
	return _ForwarderFactory.Contract.Stop(&_ForwarderFactory.TransactOpts)
}

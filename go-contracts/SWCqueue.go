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

// SWCqueueABI is the input ABI used to generate the binding from.
const SWCqueueABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"brgAmount\",\"type\":\"uint256\"}],\"name\":\"cancel\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"ratio\",\"type\":\"uint256\"}],\"name\":\"setNextBRGSWTratio\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"brg\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"brg_\",\"type\":\"address\"}],\"name\":\"setBRG\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"brgAmount\",\"type\":\"uint128\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextBRGSWTratio\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"brg_\",\"type\":\"address\"},{\"name\":\"nextBRGSWTratio_\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"UnauthorizedAccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"brgAmount\",\"type\":\"uint256\"}],\"name\":\"LogCancel\",\"type\":\"event\"}]"

// SWCqueueBin is the compiled bytecode used for deploying new contracts.
const SWCqueueBin = `"0x6060604052341561000f57600080fd5b6040516040806106e383398101604052808051919060200180519150505b5b60018054600160a060020a03191633600160a060020a03169081179091557fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a25b600160a060020a038216151561008e57600080fd5b60028054600160a060020a031916600160a060020a03841617905560038190555b50505b610622806100c16000396000f300606060405236156100885763ffffffff60e060020a60003504166313af4035811461008d57806340e58ee5146100ae578063429ad4f2146100c65780634f9c8fe8146100de578063679ccc6d1461010d5780637a9e5e4b1461012e5780638da5cb5b1461014f57806390bc16931461017e578063ac63f9d9146101a8578063bf7e214f146101cd575b600080fd5b341561009857600080fd5b6100ac600160a060020a03600435166101fc565b005b34156100b957600080fd5b6100ac60043561027a565b005b34156100d157600080fd5b6100ac6004356102c2565b005b34156100e957600080fd5b6100f16102f7565b604051600160a060020a03909116815260200160405180910390f35b341561011857600080fd5b6100ac600160a060020a0360043516610306565b005b341561013957600080fd5b6100ac600160a060020a0360043516610366565b005b341561015a57600080fd5b6100f16103e4565b604051600160a060020a03909116815260200160405180910390f35b341561018957600080fd5b6100ac6fffffffffffffffffffffffffffffffff600435166103f3565b005b34156101b357600080fd5b6101bb610483565b60405190815260200160405180910390f35b34156101d857600080fd5b6100f1610489565b604051600160a060020a03909116815260200160405180910390f35b61021233600035600160e060020a031916610498565b151561021a57fe5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a25b5b50565b7fe5ead59b814a782a5c21fa4b84845d70e9f65d11c66febe7ead5c234b68409ac3382604051600160a060020a03909216825260208201526040908101905180910390a15b50565b6102d833600035600160e060020a031916610498565b15156102e057fe5b600081116102ed57600080fd5b60038190555b5b50565b600254600160a060020a031681565b61031c33600035600160e060020a031916610498565b151561032457fe5b600160a060020a038116151561033957600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b5b50565b61037c33600035600160e060020a031916610498565b151561038457fe5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167f1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada460405160405180910390a25b5b50565b600154600160a060020a031681565b61040933600035600160e060020a031916610498565b151561041157fe5b600254600160a060020a03166390bc16938260405160e060020a63ffffffff84160281526fffffffffffffffffffffffffffffffff9091166004820152602401600060405180830381600087803b151561046a57600080fd5b6102c65a03f1151561047b57600080fd5b5050505b5b50565b60035481565b600054600160a060020a031681565b600030600160a060020a031683600160a060020a031614156104bc575060016105ed565b600154600160a060020a0384811691161480156104e25750600054600160a060020a0316155b156104ef575060016105ed565b600054600160a060020a03161515610557577feb91385a0d70e44b915093d9ddfe6c8b41f2c56729431b1405f40aee5d874be78383604051600160a060020a039092168252600160e060020a03191660208201526040908101905180910390a15060006105ed565b60008054600160a060020a03169063b7009613908590309086906040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152600160e060020a03199091166044820152606401602060405180830381600087803b15156105d057600080fd5b6102c65a03f115156105e157600080fd5b50505060405180519150505b5b5b5b929150505600a165627a7a72305820721e5022da567bb47a8ce4c21de5b5d5451e80fc629f7948168e2e4741dbc7b90029"`

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

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(brgAmount uint256) returns()
func (_SWCqueue *SWCqueueTransactor) Cancel(opts *bind.TransactOpts, brgAmount *big.Int) (*types.Transaction, error) {
	return _SWCqueue.contract.Transact(opts, "cancel", brgAmount)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(brgAmount uint256) returns()
func (_SWCqueue *SWCqueueSession) Cancel(brgAmount *big.Int) (*types.Transaction, error) {
	return _SWCqueue.Contract.Cancel(&_SWCqueue.TransactOpts, brgAmount)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(brgAmount uint256) returns()
func (_SWCqueue *SWCqueueTransactorSession) Cancel(brgAmount *big.Int) (*types.Transaction, error) {
	return _SWCqueue.Contract.Cancel(&_SWCqueue.TransactOpts, brgAmount)
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

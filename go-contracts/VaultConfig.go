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

// VaultConfigABI is the input ABI used to generate the binding from.
const VaultConfigABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"maxUOU\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"uouFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"days_\",\"type\":\"uint32\"}],\"name\":\"setMaxUOUdays\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"ratio\",\"type\":\"uint32\"}],\"name\":\"setMaxUOU\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address[]\"},{\"name\":\"ratio\",\"type\":\"uint32[]\"}],\"name\":\"setMaxUOUs\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxUOUdays\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"UnauthorizedAccess\",\"type\":\"event\"}]"

// VaultConfigBin is the compiled bytecode used for deploying new contracts.
const VaultConfigBin = `"0x60606040525b60018054600160a060020a03191633600160a060020a03169081179091557fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a25b5b6106ba806100606000396000f300606060405236156100a15763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630729bfe881146100a657806313af4035146100de5780633bcdd3cf146100ff57806340b9e5ad146101245780637a9e5e4b146101425780638da5cb5b1461016357806398bd594e14610192578063a714f0d1146101bc578063bf7e214f1461024d578063c8031b811461027c575b600080fd5b34156100b157600080fd5b6100c5600160a060020a03600435166102a8565b60405163ffffffff909116815260200160405180910390f35b34156100e957600080fd5b6100fd600160a060020a03600435166102c0565b005b341561010a57600080fd5b61011261033e565b60405190815260200160405180910390f35b341561012f57600080fd5b6100fd63ffffffff60043516610344565b005b341561014d57600080fd5b6100fd600160a060020a0360043516610388565b005b341561016e57600080fd5b610176610406565b604051600160a060020a03909116815260200160405180910390f35b341561019d57600080fd5b6100fd600160a060020a036004351663ffffffff60243516610415565b005b34156101c757600080fd5b6100fd60046024813581810190830135806020818102016040519081016040528093929190818152602001838360200280828437820191505050505050919080359060200190820180359060200190808060200260200160405190810160405280939291908181526020018383602002808284375094965061044595505050505050565b005b341561025857600080fd5b6101766104e4565b604051600160a060020a03909116815260200160405180910390f35b341561028757600080fd5b6100c56104f3565b60405163ffffffff909116815260200160405180910390f35b60026020526000908152604090205463ffffffff1681565b6102d633600035600160e060020a031916610517565b15156102de57fe5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a25b5b50565b60035481565b6001805477ffffffff000000000000000000000000000000000000000019167401000000000000000000000000000000000000000063ffffffff8416021790555b50565b61039e33600035600160e060020a031916610517565b15156103a657fe5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167f1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada460405160405180910390a25b5b50565b600154600160a060020a031681565b600160a060020a0382166000908152600260205260409020805463ffffffff191663ffffffff83161790555b5050565b6000815183511461045557600080fd5b5060005b82518163ffffffff1610156104de57818163ffffffff168151811061047a57fe5b9060200190602002015160026000858463ffffffff168151811061049a57fe5b90602001906020020151600160a060020a031681526020810191909152604001600020805463ffffffff191663ffffffff929092169190911790555b600101610459565b5b505050565b600054600160a060020a031681565b60015474010000000000000000000000000000000000000000900463ffffffff1681565b600030600160a060020a031683600160a060020a0316141561053b57506001610685565b600154600160a060020a0384811691161480156105615750600054600160a060020a0316155b1561056e57506001610685565b600054600160a060020a031615156105d6577feb91385a0d70e44b915093d9ddfe6c8b41f2c56729431b1405f40aee5d874be78383604051600160a060020a039092168252600160e060020a03191660208201526040908101905180910390a1506000610685565b60008054600160a060020a03169063b700961390859030908690604051602001526040517c010000000000000000000000000000000000000000000000000000000063ffffffff8616028152600160a060020a039384166004820152919092166024820152600160e060020a03199091166044820152606401602060405180830381600087803b151561066857600080fd5b6102c65a03f1151561067957600080fd5b50505060405180519150505b5b5b5b929150505600a165627a7a72305820338363111e34558863d13c2b10399f58bc2a2f03887cb9d9c436df140510195a0029"`

// DeployVaultConfig deploys a new Ethereum contract, binding an instance of VaultConfig to it.
func DeployVaultConfig(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *VaultConfig, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultConfigABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VaultConfigBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &VaultConfig{VaultConfigCaller: VaultConfigCaller{contract: contract}, VaultConfigTransactor: VaultConfigTransactor{contract: contract}}, nil
}

// VaultConfig is an auto generated Go binding around an Ethereum contract.
type VaultConfig struct {
	VaultConfigCaller     // Read-only binding to the contract
	VaultConfigTransactor // Write-only binding to the contract
}

// VaultConfigCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultConfigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultConfigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultConfigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultConfigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultConfigSession struct {
	Contract     *VaultConfig      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultConfigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultConfigCallerSession struct {
	Contract *VaultConfigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// VaultConfigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultConfigTransactorSession struct {
	Contract     *VaultConfigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// VaultConfigRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultConfigRaw struct {
	Contract *VaultConfig // Generic contract binding to access the raw methods on
}

// VaultConfigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultConfigCallerRaw struct {
	Contract *VaultConfigCaller // Generic read-only contract binding to access the raw methods on
}

// VaultConfigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultConfigTransactorRaw struct {
	Contract *VaultConfigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVaultConfig creates a new instance of VaultConfig, bound to a specific deployed contract.
func NewVaultConfig(address common.Address, backend bind.ContractBackend) (*VaultConfig, error) {
	contract, err := bindVaultConfig(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &VaultConfig{VaultConfigCaller: VaultConfigCaller{contract: contract}, VaultConfigTransactor: VaultConfigTransactor{contract: contract}}, nil
}

// NewVaultConfigCaller creates a new read-only instance of VaultConfig, bound to a specific deployed contract.
func NewVaultConfigCaller(address common.Address, caller bind.ContractCaller) (*VaultConfigCaller, error) {
	contract, err := bindVaultConfig(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &VaultConfigCaller{contract: contract}, nil
}

// NewVaultConfigTransactor creates a new write-only instance of VaultConfig, bound to a specific deployed contract.
func NewVaultConfigTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultConfigTransactor, error) {
	contract, err := bindVaultConfig(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &VaultConfigTransactor{contract: contract}, nil
}

// bindVaultConfig binds a generic wrapper to an already deployed contract.
func bindVaultConfig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultConfigABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultConfig *VaultConfigRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VaultConfig.Contract.VaultConfigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultConfig *VaultConfigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultConfig.Contract.VaultConfigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultConfig *VaultConfigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultConfig.Contract.VaultConfigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_VaultConfig *VaultConfigCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _VaultConfig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_VaultConfig *VaultConfigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _VaultConfig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_VaultConfig *VaultConfigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _VaultConfig.Contract.contract.Transact(opts, method, params...)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_VaultConfig *VaultConfigCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VaultConfig.contract.Call(opts, out, "authority")
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_VaultConfig *VaultConfigSession) Authority() (common.Address, error) {
	return _VaultConfig.Contract.Authority(&_VaultConfig.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_VaultConfig *VaultConfigCallerSession) Authority() (common.Address, error) {
	return _VaultConfig.Contract.Authority(&_VaultConfig.CallOpts)
}

// MaxUOU is a free data retrieval call binding the contract method 0x0729bfe8.
//
// Solidity: function maxUOU( address) constant returns(uint32)
func (_VaultConfig *VaultConfigCaller) MaxUOU(opts *bind.CallOpts, arg0 common.Address) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _VaultConfig.contract.Call(opts, out, "maxUOU", arg0)
	return *ret0, err
}

// MaxUOU is a free data retrieval call binding the contract method 0x0729bfe8.
//
// Solidity: function maxUOU( address) constant returns(uint32)
func (_VaultConfig *VaultConfigSession) MaxUOU(arg0 common.Address) (uint32, error) {
	return _VaultConfig.Contract.MaxUOU(&_VaultConfig.CallOpts, arg0)
}

// MaxUOU is a free data retrieval call binding the contract method 0x0729bfe8.
//
// Solidity: function maxUOU( address) constant returns(uint32)
func (_VaultConfig *VaultConfigCallerSession) MaxUOU(arg0 common.Address) (uint32, error) {
	return _VaultConfig.Contract.MaxUOU(&_VaultConfig.CallOpts, arg0)
}

// MaxUOUdays is a free data retrieval call binding the contract method 0xc8031b81.
//
// Solidity: function maxUOUdays() constant returns(uint32)
func (_VaultConfig *VaultConfigCaller) MaxUOUdays(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _VaultConfig.contract.Call(opts, out, "maxUOUdays")
	return *ret0, err
}

// MaxUOUdays is a free data retrieval call binding the contract method 0xc8031b81.
//
// Solidity: function maxUOUdays() constant returns(uint32)
func (_VaultConfig *VaultConfigSession) MaxUOUdays() (uint32, error) {
	return _VaultConfig.Contract.MaxUOUdays(&_VaultConfig.CallOpts)
}

// MaxUOUdays is a free data retrieval call binding the contract method 0xc8031b81.
//
// Solidity: function maxUOUdays() constant returns(uint32)
func (_VaultConfig *VaultConfigCallerSession) MaxUOUdays() (uint32, error) {
	return _VaultConfig.Contract.MaxUOUdays(&_VaultConfig.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_VaultConfig *VaultConfigCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _VaultConfig.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_VaultConfig *VaultConfigSession) Owner() (common.Address, error) {
	return _VaultConfig.Contract.Owner(&_VaultConfig.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_VaultConfig *VaultConfigCallerSession) Owner() (common.Address, error) {
	return _VaultConfig.Contract.Owner(&_VaultConfig.CallOpts)
}

// UouFee is a free data retrieval call binding the contract method 0x3bcdd3cf.
//
// Solidity: function uouFee() constant returns(uint256)
func (_VaultConfig *VaultConfigCaller) UouFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _VaultConfig.contract.Call(opts, out, "uouFee")
	return *ret0, err
}

// UouFee is a free data retrieval call binding the contract method 0x3bcdd3cf.
//
// Solidity: function uouFee() constant returns(uint256)
func (_VaultConfig *VaultConfigSession) UouFee() (*big.Int, error) {
	return _VaultConfig.Contract.UouFee(&_VaultConfig.CallOpts)
}

// UouFee is a free data retrieval call binding the contract method 0x3bcdd3cf.
//
// Solidity: function uouFee() constant returns(uint256)
func (_VaultConfig *VaultConfigCallerSession) UouFee() (*big.Int, error) {
	return _VaultConfig.Contract.UouFee(&_VaultConfig.CallOpts)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_VaultConfig *VaultConfigTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _VaultConfig.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_VaultConfig *VaultConfigSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _VaultConfig.Contract.SetAuthority(&_VaultConfig.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_VaultConfig *VaultConfigTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _VaultConfig.Contract.SetAuthority(&_VaultConfig.TransactOpts, authority_)
}

// SetMaxUOU is a paid mutator transaction binding the contract method 0x98bd594e.
//
// Solidity: function setMaxUOU(token address, ratio uint32) returns()
func (_VaultConfig *VaultConfigTransactor) SetMaxUOU(opts *bind.TransactOpts, token common.Address, ratio uint32) (*types.Transaction, error) {
	return _VaultConfig.contract.Transact(opts, "setMaxUOU", token, ratio)
}

// SetMaxUOU is a paid mutator transaction binding the contract method 0x98bd594e.
//
// Solidity: function setMaxUOU(token address, ratio uint32) returns()
func (_VaultConfig *VaultConfigSession) SetMaxUOU(token common.Address, ratio uint32) (*types.Transaction, error) {
	return _VaultConfig.Contract.SetMaxUOU(&_VaultConfig.TransactOpts, token, ratio)
}

// SetMaxUOU is a paid mutator transaction binding the contract method 0x98bd594e.
//
// Solidity: function setMaxUOU(token address, ratio uint32) returns()
func (_VaultConfig *VaultConfigTransactorSession) SetMaxUOU(token common.Address, ratio uint32) (*types.Transaction, error) {
	return _VaultConfig.Contract.SetMaxUOU(&_VaultConfig.TransactOpts, token, ratio)
}

// SetMaxUOUdays is a paid mutator transaction binding the contract method 0x40b9e5ad.
//
// Solidity: function setMaxUOUdays(days_ uint32) returns()
func (_VaultConfig *VaultConfigTransactor) SetMaxUOUdays(opts *bind.TransactOpts, days_ uint32) (*types.Transaction, error) {
	return _VaultConfig.contract.Transact(opts, "setMaxUOUdays", days_)
}

// SetMaxUOUdays is a paid mutator transaction binding the contract method 0x40b9e5ad.
//
// Solidity: function setMaxUOUdays(days_ uint32) returns()
func (_VaultConfig *VaultConfigSession) SetMaxUOUdays(days_ uint32) (*types.Transaction, error) {
	return _VaultConfig.Contract.SetMaxUOUdays(&_VaultConfig.TransactOpts, days_)
}

// SetMaxUOUdays is a paid mutator transaction binding the contract method 0x40b9e5ad.
//
// Solidity: function setMaxUOUdays(days_ uint32) returns()
func (_VaultConfig *VaultConfigTransactorSession) SetMaxUOUdays(days_ uint32) (*types.Transaction, error) {
	return _VaultConfig.Contract.SetMaxUOUdays(&_VaultConfig.TransactOpts, days_)
}

// SetMaxUOUs is a paid mutator transaction binding the contract method 0xa714f0d1.
//
// Solidity: function setMaxUOUs(token address[], ratio uint32[]) returns()
func (_VaultConfig *VaultConfigTransactor) SetMaxUOUs(opts *bind.TransactOpts, token []common.Address, ratio []uint32) (*types.Transaction, error) {
	return _VaultConfig.contract.Transact(opts, "setMaxUOUs", token, ratio)
}

// SetMaxUOUs is a paid mutator transaction binding the contract method 0xa714f0d1.
//
// Solidity: function setMaxUOUs(token address[], ratio uint32[]) returns()
func (_VaultConfig *VaultConfigSession) SetMaxUOUs(token []common.Address, ratio []uint32) (*types.Transaction, error) {
	return _VaultConfig.Contract.SetMaxUOUs(&_VaultConfig.TransactOpts, token, ratio)
}

// SetMaxUOUs is a paid mutator transaction binding the contract method 0xa714f0d1.
//
// Solidity: function setMaxUOUs(token address[], ratio uint32[]) returns()
func (_VaultConfig *VaultConfigTransactorSession) SetMaxUOUs(token []common.Address, ratio []uint32) (*types.Transaction, error) {
	return _VaultConfig.Contract.SetMaxUOUs(&_VaultConfig.TransactOpts, token, ratio)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_VaultConfig *VaultConfigTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _VaultConfig.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_VaultConfig *VaultConfigSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _VaultConfig.Contract.SetOwner(&_VaultConfig.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_VaultConfig *VaultConfigTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _VaultConfig.Contract.SetOwner(&_VaultConfig.TransactOpts, owner_)
}

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

// AssetsABI is the input ABI used to generate the binding from.
const AssetsABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"addAsset\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"cleanStorage\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assets\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"}],\"name\":\"rmAsset\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"UnauthorizedAccess\",\"type\":\"event\"}]"

// AssetsBin is the compiled bytecode used for deploying new contracts.
const AssetsBin = `"0x60606040525b60018054600160a060020a03191633600160a060020a03169081179091557fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a25b5b6109df806100606000396000f300606060405236156100935763ffffffff60e060020a60003504166313af4035811461009857806356e596a8146100b957806370a08231146100f5578063712e6279146101265780637a9e5e4b146101625780637bb98a68146101835780637e84d36b146101ea5780638da5cb5b146101ff578063bf7e214f1461022e578063cf35bdd01461025d578063d1da3e7b1461028f575b600080fd5b34156100a357600080fd5b6100b7600160a060020a03600435166102b6565b005b34156100c457600080fd5b6100b7600160a060020a03600435811690602435166fffffffffffffffffffffffffffffffff60443516610334565b005b341561010057600080fd5b610114600160a060020a03600435166103e3565b60405190815260200160405180910390f35b341561013157600080fd5b6100b7600160a060020a03600435811690602435166fffffffffffffffffffffffffffffffff6044351661045e565b005b341561016d57600080fd5b6100b7600160a060020a0360043516610518565b005b341561018e57600080fd5b610196610596565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156101d65780820151818401525b6020016101bd565b505050509050019250505060405180910390f35b34156101f557600080fd5b6100b76105ae565b005b341561020a57600080fd5b610212610634565b604051600160a060020a03909116815260200160405180910390f35b341561023957600080fd5b610212610643565b604051600160a060020a03909116815260200160405180910390f35b341561026857600080fd5b610212600435610652565b604051600160a060020a03909116815260200160405180910390f35b341561029a57600080fd5b6100b7600160a060020a0360043581169060243516610684565b005b6102cc33600035600160e060020a031916610724565b15156102d457fe5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a25b5b50565b61034a33600035600160e060020a031916610724565b151561035257fe5b82600160a060020a031663a9059cbb838360006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526fffffffffffffffffffffffffffffffff166024820152604401602060405180830381600087803b15156103c157600080fd5b6102c65a03f115156103d257600080fd5b505050604051805150505b5b505050565b600081600160a060020a03166370a082313060006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b151561043c57600080fd5b6102c65a03f1151561044d57600080fd5b50505060405180519150505b919050565b61047433600035600160e060020a031916610724565b151561047c57fe5b73__AssetsLib_____________________________6352d28bc8600285858560405160e060020a63ffffffff87160281526004810194909452600160a060020a039283166024850152911660448301526fffffffffffffffffffffffffffffffff16606482015260840160006040518083038186803b15156104fd57600080fd5b6102c65a03f4151561050e57600080fd5b5050505b5b505050565b61052e33600035600160e060020a031916610724565b151561053657fe5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167f1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada460405160405180910390a25b5b50565b61059e6109a1565b6105a86002610882565b90505b90565b6105c433600035600160e060020a031916610724565b15156105cc57fe5b73__AssetsLib_____________________________6376c86cd7600260405160e060020a63ffffffff8416028152600481019190915260240160006040518083038186803b151561061c57600080fd5b6102c65a03f415156103dd57600080fd5b5050505b5b565b600154600160a060020a031681565b600054600160a060020a031681565b600280548290811061066057fe5b906000526020600020900160005b915054906101000a9004600160a060020a031681565b61069a33600035600160e060020a031916610724565b15156106a257fe5b73__AssetsLib_____________________________638a7fcb2d6002848460405160e060020a63ffffffff86160281526004810193909352600160a060020a03918216602484015216604482015260640160006040518083038186803b151561070a57600080fd5b6102c65a03f4151561071b57600080fd5b5050505b5b5050565b600030600160a060020a031683600160a060020a0316141561074857506001610879565b600154600160a060020a03848116911614801561076e5750600054600160a060020a0316155b1561077b57506001610879565b600054600160a060020a031615156107e3577feb91385a0d70e44b915093d9ddfe6c8b41f2c56729431b1405f40aee5d874be78383604051600160a060020a039092168252600160e060020a03191660208201526040908101905180910390a1506000610879565b60008054600160a060020a03169063b7009613908590309086906040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152600160e060020a03199091166044820152606401602060405180830381600087803b151561085c57600080fd5b6102c65a03f1151561086d57600080fd5b50505060405180519150505b5b5b5b92915050565b61088a6109a1565b6108926109a1565b82546000906040518059106108a45750595b908082528060200260200182016040525b509150600090505b835463ffffffff8216101561099657838163ffffffff168154811015156108e057fe5b906000526020600020900160005b9054906101000a9004600160a060020a0316600160a060020a03166370a082313060006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b151561095657600080fd5b6102c65a03f1151561096757600080fd5b50505060405180519050828263ffffffff168151811061098357fe5b602090810290910101525b6001016108bd565b8192505b5050919050565b602060405190810160405260008152905600a165627a7a72305820bf46c1b3c6c2671f7b597c07868562c96af2af78bea32e00e0309c34580eea670029"`

// DeployAssets deploys a new Ethereum contract, binding an instance of Assets to it.
func DeployAssets(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Assets, error) {
	parsed, err := abi.JSON(strings.NewReader(AssetsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AssetsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Assets{AssetsCaller: AssetsCaller{contract: contract}, AssetsTransactor: AssetsTransactor{contract: contract}}, nil
}

// Assets is an auto generated Go binding around an Ethereum contract.
type Assets struct {
	AssetsCaller     // Read-only binding to the contract
	AssetsTransactor // Write-only binding to the contract
}

// AssetsCaller is an auto generated read-only Go binding around an Ethereum contract.
type AssetsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AssetsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AssetsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AssetsSession struct {
	Contract     *Assets           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AssetsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AssetsCallerSession struct {
	Contract *AssetsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AssetsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AssetsTransactorSession struct {
	Contract     *AssetsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AssetsRaw is an auto generated low-level Go binding around an Ethereum contract.
type AssetsRaw struct {
	Contract *Assets // Generic contract binding to access the raw methods on
}

// AssetsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AssetsCallerRaw struct {
	Contract *AssetsCaller // Generic read-only contract binding to access the raw methods on
}

// AssetsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AssetsTransactorRaw struct {
	Contract *AssetsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAssets creates a new instance of Assets, bound to a specific deployed contract.
func NewAssets(address common.Address, backend bind.ContractBackend) (*Assets, error) {
	contract, err := bindAssets(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Assets{AssetsCaller: AssetsCaller{contract: contract}, AssetsTransactor: AssetsTransactor{contract: contract}}, nil
}

// NewAssetsCaller creates a new read-only instance of Assets, bound to a specific deployed contract.
func NewAssetsCaller(address common.Address, caller bind.ContractCaller) (*AssetsCaller, error) {
	contract, err := bindAssets(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &AssetsCaller{contract: contract}, nil
}

// NewAssetsTransactor creates a new write-only instance of Assets, bound to a specific deployed contract.
func NewAssetsTransactor(address common.Address, transactor bind.ContractTransactor) (*AssetsTransactor, error) {
	contract, err := bindAssets(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &AssetsTransactor{contract: contract}, nil
}

// bindAssets binds a generic wrapper to an already deployed contract.
func bindAssets(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AssetsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Assets *AssetsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Assets.Contract.AssetsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Assets *AssetsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Assets.Contract.AssetsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Assets *AssetsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Assets.Contract.AssetsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Assets *AssetsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Assets.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Assets *AssetsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Assets.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Assets *AssetsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Assets.Contract.contract.Transact(opts, method, params...)
}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets( uint256) constant returns(address)
func (_Assets *AssetsCaller) Assets(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Assets.contract.Call(opts, out, "assets", arg0)
	return *ret0, err
}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets( uint256) constant returns(address)
func (_Assets *AssetsSession) Assets(arg0 *big.Int) (common.Address, error) {
	return _Assets.Contract.Assets(&_Assets.CallOpts, arg0)
}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets( uint256) constant returns(address)
func (_Assets *AssetsCallerSession) Assets(arg0 *big.Int) (common.Address, error) {
	return _Assets.Contract.Assets(&_Assets.CallOpts, arg0)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Assets *AssetsCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Assets.contract.Call(opts, out, "authority")
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Assets *AssetsSession) Authority() (common.Address, error) {
	return _Assets.Contract.Authority(&_Assets.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Assets *AssetsCallerSession) Authority() (common.Address, error) {
	return _Assets.Contract.Authority(&_Assets.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(token address) constant returns(uint256)
func (_Assets *AssetsCaller) BalanceOf(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Assets.contract.Call(opts, out, "balanceOf", token)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(token address) constant returns(uint256)
func (_Assets *AssetsSession) BalanceOf(token common.Address) (*big.Int, error) {
	return _Assets.Contract.BalanceOf(&_Assets.CallOpts, token)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(token address) constant returns(uint256)
func (_Assets *AssetsCallerSession) BalanceOf(token common.Address) (*big.Int, error) {
	return _Assets.Contract.BalanceOf(&_Assets.CallOpts, token)
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(uint256[])
func (_Assets *AssetsCaller) Balances(opts *bind.CallOpts) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Assets.contract.Call(opts, out, "balances")
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(uint256[])
func (_Assets *AssetsSession) Balances() ([]*big.Int, error) {
	return _Assets.Contract.Balances(&_Assets.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(uint256[])
func (_Assets *AssetsCallerSession) Balances() ([]*big.Int, error) {
	return _Assets.Contract.Balances(&_Assets.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Assets *AssetsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Assets.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Assets *AssetsSession) Owner() (common.Address, error) {
	return _Assets.Contract.Owner(&_Assets.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Assets *AssetsCallerSession) Owner() (common.Address, error) {
	return _Assets.Contract.Owner(&_Assets.CallOpts)
}

// AddAsset is a paid mutator transaction binding the contract method 0x712e6279.
//
// Solidity: function addAsset(token address, src address, wad uint128) returns()
func (_Assets *AssetsTransactor) AddAsset(opts *bind.TransactOpts, token common.Address, src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Assets.contract.Transact(opts, "addAsset", token, src, wad)
}

// AddAsset is a paid mutator transaction binding the contract method 0x712e6279.
//
// Solidity: function addAsset(token address, src address, wad uint128) returns()
func (_Assets *AssetsSession) AddAsset(token common.Address, src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Assets.Contract.AddAsset(&_Assets.TransactOpts, token, src, wad)
}

// AddAsset is a paid mutator transaction binding the contract method 0x712e6279.
//
// Solidity: function addAsset(token address, src address, wad uint128) returns()
func (_Assets *AssetsTransactorSession) AddAsset(token common.Address, src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Assets.Contract.AddAsset(&_Assets.TransactOpts, token, src, wad)
}

// CleanStorage is a paid mutator transaction binding the contract method 0x7e84d36b.
//
// Solidity: function cleanStorage() returns()
func (_Assets *AssetsTransactor) CleanStorage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Assets.contract.Transact(opts, "cleanStorage")
}

// CleanStorage is a paid mutator transaction binding the contract method 0x7e84d36b.
//
// Solidity: function cleanStorage() returns()
func (_Assets *AssetsSession) CleanStorage() (*types.Transaction, error) {
	return _Assets.Contract.CleanStorage(&_Assets.TransactOpts)
}

// CleanStorage is a paid mutator transaction binding the contract method 0x7e84d36b.
//
// Solidity: function cleanStorage() returns()
func (_Assets *AssetsTransactorSession) CleanStorage() (*types.Transaction, error) {
	return _Assets.Contract.CleanStorage(&_Assets.TransactOpts)
}

// RmAsset is a paid mutator transaction binding the contract method 0xd1da3e7b.
//
// Solidity: function rmAsset(token address, dst address) returns()
func (_Assets *AssetsTransactor) RmAsset(opts *bind.TransactOpts, token common.Address, dst common.Address) (*types.Transaction, error) {
	return _Assets.contract.Transact(opts, "rmAsset", token, dst)
}

// RmAsset is a paid mutator transaction binding the contract method 0xd1da3e7b.
//
// Solidity: function rmAsset(token address, dst address) returns()
func (_Assets *AssetsSession) RmAsset(token common.Address, dst common.Address) (*types.Transaction, error) {
	return _Assets.Contract.RmAsset(&_Assets.TransactOpts, token, dst)
}

// RmAsset is a paid mutator transaction binding the contract method 0xd1da3e7b.
//
// Solidity: function rmAsset(token address, dst address) returns()
func (_Assets *AssetsTransactorSession) RmAsset(token common.Address, dst common.Address) (*types.Transaction, error) {
	return _Assets.Contract.RmAsset(&_Assets.TransactOpts, token, dst)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_Assets *AssetsTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _Assets.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_Assets *AssetsSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Assets.Contract.SetAuthority(&_Assets.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_Assets *AssetsTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Assets.Contract.SetAuthority(&_Assets.TransactOpts, authority_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_Assets *AssetsTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _Assets.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_Assets *AssetsSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Assets.Contract.SetOwner(&_Assets.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_Assets *AssetsTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Assets.Contract.SetOwner(&_Assets.TransactOpts, owner_)
}

// Transfer is a paid mutator transaction binding the contract method 0x56e596a8.
//
// Solidity: function transfer(token address, dst address, wad uint128) returns()
func (_Assets *AssetsTransactor) Transfer(opts *bind.TransactOpts, token common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Assets.contract.Transact(opts, "transfer", token, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0x56e596a8.
//
// Solidity: function transfer(token address, dst address, wad uint128) returns()
func (_Assets *AssetsSession) Transfer(token common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Assets.Contract.Transfer(&_Assets.TransactOpts, token, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0x56e596a8.
//
// Solidity: function transfer(token address, dst address, wad uint128) returns()
func (_Assets *AssetsTransactorSession) Transfer(token common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Assets.Contract.Transfer(&_Assets.TransactOpts, token, dst, wad)
}

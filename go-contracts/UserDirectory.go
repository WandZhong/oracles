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

// UserDirectoryABI is the input ABI used to generate the binding from.
const UserDirectoryABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"senderHasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"restart\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"walletCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"roles\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"brg\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"wallets\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kyc\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wallet_\",\"type\":\"address\"}],\"name\":\"removeWallet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"profile\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"roles_\",\"type\":\"address\"}],\"name\":\"setRolesContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"kyc_\",\"type\":\"bool\"}],\"name\":\"setKYC\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"hasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currency\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes3\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"root\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"addWallet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"root_\",\"type\":\"address\"},{\"name\":\"owner_\",\"type\":\"address\"},{\"name\":\"rolesContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newWallet\",\"type\":\"address\"}],\"name\":\"LogWalletAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"removedWallet\",\"type\":\"address\"}],\"name\":\"LogWalletRemoved\",\"type\":\"event\"}]"

// UserDirectoryBin is the compiled bytecode used for deploying new contracts.
const UserDirectoryBin = `"0x60606040526003805460ff191690556006805462ffffff1916625553441763ff00000019169055341561003157600080fd5b604051606080610ebf83398101604052808051919060200180519190602001805191506040905080519081016040908152600d82527f557365724469726563746f727900000000000000000000000000000000000000602083015260008054600160a060020a03191633600160a060020a031617905582908290518082805190602001908083835b602083106100d85780518252601f1990920191602091820191016100b9565b6001836020036101000a038019825116818451161790925250505091909101925060409150505190819003902060025560018054600160a060020a031916600160a060020a0392831617905584161515905061013357600080fd5b600160a060020a038216151561014857600080fd5b5060008054600160a060020a03928316600160a060020a031990911617905560038054929091166101000261010060a860020a0319909216919091179055610d2a806101956000396000f300606060405236156100f65763ffffffff60e060020a60003504166307da68f581146100fb57806313af4035146101105780631ca03b8e1461012f5780631ef3755d1461019457806329b57c69146101a7578063392f5f64146101d35780634f9c8fe81461020257806375f12b21146102155780637ad71f72146102285780638da5cb5b1461023e578063904c60941461025157806390d6b45f14610276578063a75fe8e114610289578063ab60636c146102a8578063afa202ac14610332578063bb05c30e14610351578063e3c33a9b14610369578063e5a6b10f146103ba578063ebf0c71714610402578063fbd51eee14610415575b600080fd5b341561010657600080fd5b61010e610428565b005b341561011b57600080fd5b61010e600160a060020a0360043516610498565b341561013a57600080fd5b61018060046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061052295505050505050565b604051901515815260200160405180910390f35b341561019f57600080fd5b61010e61060d565b34156101b257600080fd5b6101ba61067a565b60405163ffffffff909116815260200160405180910390f35b34156101de57600080fd5b6101e6610681565b604051600160a060020a03909116815260200160405180910390f35b341561020d57600080fd5b6101e6610690565b341561022057600080fd5b6101806106fe565b341561023357600080fd5b6101e6600435610707565b341561024957600080fd5b6101e661072f565b341561025c57600080fd5b61026461073e565b60405190815260200160405180910390f35b341561028157600080fd5b610180610744565b341561029457600080fd5b61010e600160a060020a0360043516610754565b34156102b357600080fd5b6102bb610896565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156102f75780820151838201526020016102df565b50505050905090810190601f1680156103245780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561033d57600080fd5b61010e600160a060020a0360043516610934565b341561035c57600080fd5b61010e600435151561099a565b341561037457600080fd5b61018060046024813581810190830135806020601f82018190048102016040519081016040528181529291906020840183838082843750949650610a0195505050505050565b34156103c557600080fd5b6103cd610aba565b6040517fffffff0000000000000000000000000000000000000000000000000000000000909116815260200160405180910390f35b341561040d57600080fd5b6101e6610ae0565b341561042057600080fd5b61010e610af4565b60408051908101604052600781527f73746f7070657200000000000000000000000000000000000000000000000000602082015260005433600160a060020a039081169116148061047d575061047d81610522565b151561048857600080fd5b506003805460ff19166001179055565b60005433600160a060020a039081169116146104b357600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383811691909117918290557fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed949116604051600160a060020a03909116815260200160405180910390a150565b600154600254600091600160a060020a031690638b51ca4290846040518082805190602001908083835b6020831061056b5780518252601f19909201916020918201910161054c565b6001836020036101000a038019825116818451161790925250505091909101925060409150505180910390203360006040516020015260405160e060020a63ffffffff861602815260048101939093526024830191909152600160a060020a03166044820152606401602060405180830381600087803b15156105ed57600080fd5b6102c65a03f115156105fe57600080fd5b50505060405180519392505050565b60408051908101604052600981527f7265737461727465720000000000000000000000000000000000000000000000602082015260005433600160a060020a0390811691161480610662575061066281610522565b151561066d57600080fd5b506003805460ff19169055565b6007545b90565b600154600160a060020a031681565b6003546000906101009004600160a060020a0316634f9c8fe882604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156106df57600080fd5b6102c65a03f115156106f057600080fd5b505050604051805191505090565b60035460ff1681565b600780548290811061071557fe5b600091825260209091200154600160a060020a0316905081565b600054600160a060020a031681565b60025481565b6006546301000000900460ff1681565b600080548190819033600160a060020a0390811691161461077457600080fd5b6000199250600091505b6007548210156107cb5783600160a060020a03166007838154811015156107a157fe5b600091825260209091200154600160a060020a031614156107c0578192505b60019091019061077e565b6000831261089057506007805460001981019190829081106107e957fe5b60009182526020909120015460078054600160a060020a03909216918590811061080f57fe5b6000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039290921691909117905580610851600782610cb7565b507fb79624962e0e32fb056c1f0aef3a7ec011949ae012e85f598f00c445d8b23c2284604051600160a060020a03909116815260200160405180910390a15b50505050565b60058054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561092c5780601f106109015761010080835404028352916020019161092c565b820191906000526020600020905b81548152906001019060200180831161090f57829003601f168201915b505050505081565b60005433600160a060020a0390811691161461094f57600080fd5b60015430600160a060020a039081169116141561096b57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60408051908101604052600b81527f757365724d616e6167657200000000000000000000000000000000000000000060208201526109d781610522565b15156109e257600080fd5b506006805491151563010000000263ff00000019909216919091179055565b600154600254600091600160a060020a031690633037cea390846040518082805190602001908083835b60208310610a4a5780518252601f199092019160209182019101610a2b565b6001836020036101000a0380198251168184511617909252505050919091019250604091505051809103902060006040516020015260405160e060020a63ffffffff851602815260048101929092526024820152604401602060405180830381600087803b15156105ed57600080fd5b6006547d0100000000000000000000000000000000000000000000000000000000000281565b6003546101009004600160a060020a031681565b600060408051908101604052600b81527f757365724d616e61676572000000000000000000000000000000000000000000602082015260005433600160a060020a0390811691161480610b4b5750610b4b81610522565b1515610b5657600080fd5b6003546101009004600160a060020a031663c5c036996000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610ba357600080fd5b6102c65a03f11515610bb457600080fd5b50505060405180519050600160a060020a031663b054a9e83060006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b1515610c1457600080fd5b6102c65a03f11515610c2557600080fd5b505050604051805160078054919450915060018101610c448382610cb7565b506000918252602090912001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0384161790557f4bb72edbb6b4f4050e363e5db675f0f33dd63cabacdec731c7bd6fd0b87b7f8782604051600160a060020a03909116815260200160405180910390a15050565b815481835581811511610cdb57600083815260209020610cdb918101908301610ce0565b505050565b61067e91905b80821115610cfa5760008155600101610ce6565b50905600a165627a7a723058203805137560d77afdd1ba30c4ab58aecc157c51912dfce821cd7d9dbebdf7a2280029"`

// DeployUserDirectory deploys a new Ethereum contract, binding an instance of UserDirectory to it.
func DeployUserDirectory(auth *bind.TransactOpts, backend bind.ContractBackend, root_ common.Address, owner_ common.Address, rolesContract common.Address) (common.Address, *types.Transaction, *UserDirectory, error) {
	parsed, err := abi.JSON(strings.NewReader(UserDirectoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UserDirectoryBin), backend, root_, owner_, rolesContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UserDirectory{UserDirectoryCaller: UserDirectoryCaller{contract: contract}, UserDirectoryTransactor: UserDirectoryTransactor{contract: contract}}, nil
}

// UserDirectory is an auto generated Go binding around an Ethereum contract.
type UserDirectory struct {
	UserDirectoryCaller     // Read-only binding to the contract
	UserDirectoryTransactor // Write-only binding to the contract
}

// UserDirectoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type UserDirectoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserDirectoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UserDirectoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserDirectorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UserDirectorySession struct {
	Contract     *UserDirectory    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserDirectoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UserDirectoryCallerSession struct {
	Contract *UserDirectoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// UserDirectoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UserDirectoryTransactorSession struct {
	Contract     *UserDirectoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// UserDirectoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type UserDirectoryRaw struct {
	Contract *UserDirectory // Generic contract binding to access the raw methods on
}

// UserDirectoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UserDirectoryCallerRaw struct {
	Contract *UserDirectoryCaller // Generic read-only contract binding to access the raw methods on
}

// UserDirectoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UserDirectoryTransactorRaw struct {
	Contract *UserDirectoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUserDirectory creates a new instance of UserDirectory, bound to a specific deployed contract.
func NewUserDirectory(address common.Address, backend bind.ContractBackend) (*UserDirectory, error) {
	contract, err := bindUserDirectory(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UserDirectory{UserDirectoryCaller: UserDirectoryCaller{contract: contract}, UserDirectoryTransactor: UserDirectoryTransactor{contract: contract}}, nil
}

// NewUserDirectoryCaller creates a new read-only instance of UserDirectory, bound to a specific deployed contract.
func NewUserDirectoryCaller(address common.Address, caller bind.ContractCaller) (*UserDirectoryCaller, error) {
	contract, err := bindUserDirectory(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &UserDirectoryCaller{contract: contract}, nil
}

// NewUserDirectoryTransactor creates a new write-only instance of UserDirectory, bound to a specific deployed contract.
func NewUserDirectoryTransactor(address common.Address, transactor bind.ContractTransactor) (*UserDirectoryTransactor, error) {
	contract, err := bindUserDirectory(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &UserDirectoryTransactor{contract: contract}, nil
}

// bindUserDirectory binds a generic wrapper to an already deployed contract.
func bindUserDirectory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UserDirectoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserDirectory *UserDirectoryRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UserDirectory.Contract.UserDirectoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserDirectory *UserDirectoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserDirectory.Contract.UserDirectoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserDirectory *UserDirectoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserDirectory.Contract.UserDirectoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserDirectory *UserDirectoryCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UserDirectory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserDirectory *UserDirectoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserDirectory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserDirectory *UserDirectoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserDirectory.Contract.contract.Transact(opts, method, params...)
}

// Brg is a free data retrieval call binding the contract method 0x4f9c8fe8.
//
// Solidity: function brg() constant returns(address)
func (_UserDirectory *UserDirectoryCaller) Brg(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UserDirectory.contract.Call(opts, out, "brg")
	return *ret0, err
}

// Brg is a free data retrieval call binding the contract method 0x4f9c8fe8.
//
// Solidity: function brg() constant returns(address)
func (_UserDirectory *UserDirectorySession) Brg() (common.Address, error) {
	return _UserDirectory.Contract.Brg(&_UserDirectory.CallOpts)
}

// Brg is a free data retrieval call binding the contract method 0x4f9c8fe8.
//
// Solidity: function brg() constant returns(address)
func (_UserDirectory *UserDirectoryCallerSession) Brg() (common.Address, error) {
	return _UserDirectory.Contract.Brg(&_UserDirectory.CallOpts)
}

// ContractHash is a free data retrieval call binding the contract method 0x904c6094.
//
// Solidity: function contractHash() constant returns(bytes32)
func (_UserDirectory *UserDirectoryCaller) ContractHash(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _UserDirectory.contract.Call(opts, out, "contractHash")
	return *ret0, err
}

// ContractHash is a free data retrieval call binding the contract method 0x904c6094.
//
// Solidity: function contractHash() constant returns(bytes32)
func (_UserDirectory *UserDirectorySession) ContractHash() ([32]byte, error) {
	return _UserDirectory.Contract.ContractHash(&_UserDirectory.CallOpts)
}

// ContractHash is a free data retrieval call binding the contract method 0x904c6094.
//
// Solidity: function contractHash() constant returns(bytes32)
func (_UserDirectory *UserDirectoryCallerSession) ContractHash() ([32]byte, error) {
	return _UserDirectory.Contract.ContractHash(&_UserDirectory.CallOpts)
}

// Currency is a free data retrieval call binding the contract method 0xe5a6b10f.
//
// Solidity: function currency() constant returns(bytes3)
func (_UserDirectory *UserDirectoryCaller) Currency(opts *bind.CallOpts) ([3]byte, error) {
	var (
		ret0 = new([3]byte)
	)
	out := ret0
	err := _UserDirectory.contract.Call(opts, out, "currency")
	return *ret0, err
}

// Currency is a free data retrieval call binding the contract method 0xe5a6b10f.
//
// Solidity: function currency() constant returns(bytes3)
func (_UserDirectory *UserDirectorySession) Currency() ([3]byte, error) {
	return _UserDirectory.Contract.Currency(&_UserDirectory.CallOpts)
}

// Currency is a free data retrieval call binding the contract method 0xe5a6b10f.
//
// Solidity: function currency() constant returns(bytes3)
func (_UserDirectory *UserDirectoryCallerSession) Currency() ([3]byte, error) {
	return _UserDirectory.Contract.Currency(&_UserDirectory.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0xe3c33a9b.
//
// Solidity: function hasRole(roleName string) constant returns(bool)
func (_UserDirectory *UserDirectoryCaller) HasRole(opts *bind.CallOpts, roleName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _UserDirectory.contract.Call(opts, out, "hasRole", roleName)
	return *ret0, err
}

// HasRole is a free data retrieval call binding the contract method 0xe3c33a9b.
//
// Solidity: function hasRole(roleName string) constant returns(bool)
func (_UserDirectory *UserDirectorySession) HasRole(roleName string) (bool, error) {
	return _UserDirectory.Contract.HasRole(&_UserDirectory.CallOpts, roleName)
}

// HasRole is a free data retrieval call binding the contract method 0xe3c33a9b.
//
// Solidity: function hasRole(roleName string) constant returns(bool)
func (_UserDirectory *UserDirectoryCallerSession) HasRole(roleName string) (bool, error) {
	return _UserDirectory.Contract.HasRole(&_UserDirectory.CallOpts, roleName)
}

// Kyc is a free data retrieval call binding the contract method 0x90d6b45f.
//
// Solidity: function kyc() constant returns(bool)
func (_UserDirectory *UserDirectoryCaller) Kyc(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _UserDirectory.contract.Call(opts, out, "kyc")
	return *ret0, err
}

// Kyc is a free data retrieval call binding the contract method 0x90d6b45f.
//
// Solidity: function kyc() constant returns(bool)
func (_UserDirectory *UserDirectorySession) Kyc() (bool, error) {
	return _UserDirectory.Contract.Kyc(&_UserDirectory.CallOpts)
}

// Kyc is a free data retrieval call binding the contract method 0x90d6b45f.
//
// Solidity: function kyc() constant returns(bool)
func (_UserDirectory *UserDirectoryCallerSession) Kyc() (bool, error) {
	return _UserDirectory.Contract.Kyc(&_UserDirectory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_UserDirectory *UserDirectoryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UserDirectory.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_UserDirectory *UserDirectorySession) Owner() (common.Address, error) {
	return _UserDirectory.Contract.Owner(&_UserDirectory.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_UserDirectory *UserDirectoryCallerSession) Owner() (common.Address, error) {
	return _UserDirectory.Contract.Owner(&_UserDirectory.CallOpts)
}

// Profile is a free data retrieval call binding the contract method 0xab60636c.
//
// Solidity: function profile() constant returns(string)
func (_UserDirectory *UserDirectoryCaller) Profile(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _UserDirectory.contract.Call(opts, out, "profile")
	return *ret0, err
}

// Profile is a free data retrieval call binding the contract method 0xab60636c.
//
// Solidity: function profile() constant returns(string)
func (_UserDirectory *UserDirectorySession) Profile() (string, error) {
	return _UserDirectory.Contract.Profile(&_UserDirectory.CallOpts)
}

// Profile is a free data retrieval call binding the contract method 0xab60636c.
//
// Solidity: function profile() constant returns(string)
func (_UserDirectory *UserDirectoryCallerSession) Profile() (string, error) {
	return _UserDirectory.Contract.Profile(&_UserDirectory.CallOpts)
}

// Roles is a free data retrieval call binding the contract method 0x392f5f64.
//
// Solidity: function roles() constant returns(address)
func (_UserDirectory *UserDirectoryCaller) Roles(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UserDirectory.contract.Call(opts, out, "roles")
	return *ret0, err
}

// Roles is a free data retrieval call binding the contract method 0x392f5f64.
//
// Solidity: function roles() constant returns(address)
func (_UserDirectory *UserDirectorySession) Roles() (common.Address, error) {
	return _UserDirectory.Contract.Roles(&_UserDirectory.CallOpts)
}

// Roles is a free data retrieval call binding the contract method 0x392f5f64.
//
// Solidity: function roles() constant returns(address)
func (_UserDirectory *UserDirectoryCallerSession) Roles() (common.Address, error) {
	return _UserDirectory.Contract.Roles(&_UserDirectory.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() constant returns(address)
func (_UserDirectory *UserDirectoryCaller) Root(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UserDirectory.contract.Call(opts, out, "root")
	return *ret0, err
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() constant returns(address)
func (_UserDirectory *UserDirectorySession) Root() (common.Address, error) {
	return _UserDirectory.Contract.Root(&_UserDirectory.CallOpts)
}

// Root is a free data retrieval call binding the contract method 0xebf0c717.
//
// Solidity: function root() constant returns(address)
func (_UserDirectory *UserDirectoryCallerSession) Root() (common.Address, error) {
	return _UserDirectory.Contract.Root(&_UserDirectory.CallOpts)
}

// SenderHasRole is a free data retrieval call binding the contract method 0x1ca03b8e.
//
// Solidity: function senderHasRole(roleName string) constant returns(bool)
func (_UserDirectory *UserDirectoryCaller) SenderHasRole(opts *bind.CallOpts, roleName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _UserDirectory.contract.Call(opts, out, "senderHasRole", roleName)
	return *ret0, err
}

// SenderHasRole is a free data retrieval call binding the contract method 0x1ca03b8e.
//
// Solidity: function senderHasRole(roleName string) constant returns(bool)
func (_UserDirectory *UserDirectorySession) SenderHasRole(roleName string) (bool, error) {
	return _UserDirectory.Contract.SenderHasRole(&_UserDirectory.CallOpts, roleName)
}

// SenderHasRole is a free data retrieval call binding the contract method 0x1ca03b8e.
//
// Solidity: function senderHasRole(roleName string) constant returns(bool)
func (_UserDirectory *UserDirectoryCallerSession) SenderHasRole(roleName string) (bool, error) {
	return _UserDirectory.Contract.SenderHasRole(&_UserDirectory.CallOpts, roleName)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_UserDirectory *UserDirectoryCaller) Stopped(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _UserDirectory.contract.Call(opts, out, "stopped")
	return *ret0, err
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_UserDirectory *UserDirectorySession) Stopped() (bool, error) {
	return _UserDirectory.Contract.Stopped(&_UserDirectory.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_UserDirectory *UserDirectoryCallerSession) Stopped() (bool, error) {
	return _UserDirectory.Contract.Stopped(&_UserDirectory.CallOpts)
}

// WalletCount is a free data retrieval call binding the contract method 0x29b57c69.
//
// Solidity: function walletCount() constant returns(uint32)
func (_UserDirectory *UserDirectoryCaller) WalletCount(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _UserDirectory.contract.Call(opts, out, "walletCount")
	return *ret0, err
}

// WalletCount is a free data retrieval call binding the contract method 0x29b57c69.
//
// Solidity: function walletCount() constant returns(uint32)
func (_UserDirectory *UserDirectorySession) WalletCount() (uint32, error) {
	return _UserDirectory.Contract.WalletCount(&_UserDirectory.CallOpts)
}

// WalletCount is a free data retrieval call binding the contract method 0x29b57c69.
//
// Solidity: function walletCount() constant returns(uint32)
func (_UserDirectory *UserDirectoryCallerSession) WalletCount() (uint32, error) {
	return _UserDirectory.Contract.WalletCount(&_UserDirectory.CallOpts)
}

// Wallets is a free data retrieval call binding the contract method 0x7ad71f72.
//
// Solidity: function wallets( uint256) constant returns(address)
func (_UserDirectory *UserDirectoryCaller) Wallets(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UserDirectory.contract.Call(opts, out, "wallets", arg0)
	return *ret0, err
}

// Wallets is a free data retrieval call binding the contract method 0x7ad71f72.
//
// Solidity: function wallets( uint256) constant returns(address)
func (_UserDirectory *UserDirectorySession) Wallets(arg0 *big.Int) (common.Address, error) {
	return _UserDirectory.Contract.Wallets(&_UserDirectory.CallOpts, arg0)
}

// Wallets is a free data retrieval call binding the contract method 0x7ad71f72.
//
// Solidity: function wallets( uint256) constant returns(address)
func (_UserDirectory *UserDirectoryCallerSession) Wallets(arg0 *big.Int) (common.Address, error) {
	return _UserDirectory.Contract.Wallets(&_UserDirectory.CallOpts, arg0)
}

// AddWallet is a paid mutator transaction binding the contract method 0xfbd51eee.
//
// Solidity: function addWallet() returns()
func (_UserDirectory *UserDirectoryTransactor) AddWallet(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserDirectory.contract.Transact(opts, "addWallet")
}

// AddWallet is a paid mutator transaction binding the contract method 0xfbd51eee.
//
// Solidity: function addWallet() returns()
func (_UserDirectory *UserDirectorySession) AddWallet() (*types.Transaction, error) {
	return _UserDirectory.Contract.AddWallet(&_UserDirectory.TransactOpts)
}

// AddWallet is a paid mutator transaction binding the contract method 0xfbd51eee.
//
// Solidity: function addWallet() returns()
func (_UserDirectory *UserDirectoryTransactorSession) AddWallet() (*types.Transaction, error) {
	return _UserDirectory.Contract.AddWallet(&_UserDirectory.TransactOpts)
}

// RemoveWallet is a paid mutator transaction binding the contract method 0xa75fe8e1.
//
// Solidity: function removeWallet(wallet_ address) returns()
func (_UserDirectory *UserDirectoryTransactor) RemoveWallet(opts *bind.TransactOpts, wallet_ common.Address) (*types.Transaction, error) {
	return _UserDirectory.contract.Transact(opts, "removeWallet", wallet_)
}

// RemoveWallet is a paid mutator transaction binding the contract method 0xa75fe8e1.
//
// Solidity: function removeWallet(wallet_ address) returns()
func (_UserDirectory *UserDirectorySession) RemoveWallet(wallet_ common.Address) (*types.Transaction, error) {
	return _UserDirectory.Contract.RemoveWallet(&_UserDirectory.TransactOpts, wallet_)
}

// RemoveWallet is a paid mutator transaction binding the contract method 0xa75fe8e1.
//
// Solidity: function removeWallet(wallet_ address) returns()
func (_UserDirectory *UserDirectoryTransactorSession) RemoveWallet(wallet_ common.Address) (*types.Transaction, error) {
	return _UserDirectory.Contract.RemoveWallet(&_UserDirectory.TransactOpts, wallet_)
}

// Restart is a paid mutator transaction binding the contract method 0x1ef3755d.
//
// Solidity: function restart() returns()
func (_UserDirectory *UserDirectoryTransactor) Restart(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserDirectory.contract.Transact(opts, "restart")
}

// Restart is a paid mutator transaction binding the contract method 0x1ef3755d.
//
// Solidity: function restart() returns()
func (_UserDirectory *UserDirectorySession) Restart() (*types.Transaction, error) {
	return _UserDirectory.Contract.Restart(&_UserDirectory.TransactOpts)
}

// Restart is a paid mutator transaction binding the contract method 0x1ef3755d.
//
// Solidity: function restart() returns()
func (_UserDirectory *UserDirectoryTransactorSession) Restart() (*types.Transaction, error) {
	return _UserDirectory.Contract.Restart(&_UserDirectory.TransactOpts)
}

// SetKYC is a paid mutator transaction binding the contract method 0xbb05c30e.
//
// Solidity: function setKYC(kyc_ bool) returns()
func (_UserDirectory *UserDirectoryTransactor) SetKYC(opts *bind.TransactOpts, kyc_ bool) (*types.Transaction, error) {
	return _UserDirectory.contract.Transact(opts, "setKYC", kyc_)
}

// SetKYC is a paid mutator transaction binding the contract method 0xbb05c30e.
//
// Solidity: function setKYC(kyc_ bool) returns()
func (_UserDirectory *UserDirectorySession) SetKYC(kyc_ bool) (*types.Transaction, error) {
	return _UserDirectory.Contract.SetKYC(&_UserDirectory.TransactOpts, kyc_)
}

// SetKYC is a paid mutator transaction binding the contract method 0xbb05c30e.
//
// Solidity: function setKYC(kyc_ bool) returns()
func (_UserDirectory *UserDirectoryTransactorSession) SetKYC(kyc_ bool) (*types.Transaction, error) {
	return _UserDirectory.Contract.SetKYC(&_UserDirectory.TransactOpts, kyc_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_UserDirectory *UserDirectoryTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _UserDirectory.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_UserDirectory *UserDirectorySession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _UserDirectory.Contract.SetOwner(&_UserDirectory.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_UserDirectory *UserDirectoryTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _UserDirectory.Contract.SetOwner(&_UserDirectory.TransactOpts, owner_)
}

// SetRolesContract is a paid mutator transaction binding the contract method 0xafa202ac.
//
// Solidity: function setRolesContract(roles_ address) returns()
func (_UserDirectory *UserDirectoryTransactor) SetRolesContract(opts *bind.TransactOpts, roles_ common.Address) (*types.Transaction, error) {
	return _UserDirectory.contract.Transact(opts, "setRolesContract", roles_)
}

// SetRolesContract is a paid mutator transaction binding the contract method 0xafa202ac.
//
// Solidity: function setRolesContract(roles_ address) returns()
func (_UserDirectory *UserDirectorySession) SetRolesContract(roles_ common.Address) (*types.Transaction, error) {
	return _UserDirectory.Contract.SetRolesContract(&_UserDirectory.TransactOpts, roles_)
}

// SetRolesContract is a paid mutator transaction binding the contract method 0xafa202ac.
//
// Solidity: function setRolesContract(roles_ address) returns()
func (_UserDirectory *UserDirectoryTransactorSession) SetRolesContract(roles_ common.Address) (*types.Transaction, error) {
	return _UserDirectory.Contract.SetRolesContract(&_UserDirectory.TransactOpts, roles_)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_UserDirectory *UserDirectoryTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserDirectory.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_UserDirectory *UserDirectorySession) Stop() (*types.Transaction, error) {
	return _UserDirectory.Contract.Stop(&_UserDirectory.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_UserDirectory *UserDirectoryTransactorSession) Stop() (*types.Transaction, error) {
	return _UserDirectory.Contract.Stop(&_UserDirectory.TransactOpts)
}

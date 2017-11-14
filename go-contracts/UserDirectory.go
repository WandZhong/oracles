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
const UserDirectoryABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"active\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"currency\",\"type\":\"bytes3\"}],\"name\":\"addWallet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"senderHasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"restart\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeWallet\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"walletCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"roles\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"wallets\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"contractHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"profile\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"active_\",\"type\":\"bool\"}],\"name\":\"setActive\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"roles_\",\"type\":\"address\"}],\"name\":\"setRolesContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"roleName\",\"type\":\"string\"}],\"name\":\"hasRole\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"root\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"root_\",\"type\":\"address\"},{\"name\":\"owner_\",\"type\":\"address\"},{\"name\":\"rolesContract\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newWallet\",\"type\":\"address\"}],\"name\":\"LogWalletAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"removedWallet\",\"type\":\"address\"}],\"name\":\"LogWalletRemoved\",\"type\":\"event\"}]"

// UserDirectoryBin is the compiled bytecode used for deploying new contracts.
const UserDirectoryBin = `"0x60606040526000600360006101000a81548160ff0219169083151502179055506000600660006101000a81548160ff021916908315150217905550341561004557600080fd5b6040516060806117f8833981016040528080519060200190919080519060200190919080519060200190919050506040805190810160405280600d81526020017f557365724469726563746f72790000000000000000000000000000000000000081525081336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550816040518082805190602001908083835b60208310151561012057805182526020820191506020810190506020830392506100fb565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390206002816000191690555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141515156101d657600080fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415151561021257600080fd5b816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600360016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050611553806102a56000396000f300606060405260043610610107576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806302fb0c5e1461010c57806307da68f51461013957806313af40351461014e578063186caa30146101875780631ca03b8e146101ca5780631ef3755d1461023f5780632286975c1461025457806329b57c6914610269578063392f5f641461029e57806375f12b21146102f35780637ad71f72146103205780638da5cb5b14610383578063904c6094146103d8578063ab60636c14610409578063acec338a14610497578063afa202ac146104bc578063beabacc8146104f5578063e3c33a9b1461056e578063ebf0c717146105e3575b600080fd5b341561011757600080fd5b61011f610638565b604051808215151515815260200191505060405180910390f35b341561014457600080fd5b61014c61064b565b005b341561015957600080fd5b610185600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190505061070a565b005b341561019257600080fd5b6101c860048080357cffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690602001909190505061082c565b005b34156101d557600080fd5b610225600480803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091905050610b5e565b604051808215151515815260200191505060405180910390f35b341561024a57600080fd5b610252610cdc565b005b341561025f57600080fd5b610267610d9b565b005b341561027457600080fd5b61027c610f7d565b604051808263ffffffff1663ffffffff16815260200191505060405180910390f35b34156102a957600080fd5b6102b1610f8a565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156102fe57600080fd5b610306610fb0565b604051808215151515815260200191505060405180910390f35b341561032b57600080fd5b6103416004808035906020019091905050610fc3565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b341561038e57600080fd5b610396611002565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34156103e357600080fd5b6103eb611027565b60405180826000191660001916815260200191505060405180910390f35b341561041457600080fd5b61041c61102d565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561045c578082015181840152602081019050610441565b50505050905090810190601f1680156104895780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156104a257600080fd5b6104ba600480803515159060200190919050506110cb565b005b34156104c757600080fd5b6104f3600480803573ffffffffffffffffffffffffffffffffffffffff16906020019091905050611133565b005b341561050057600080fd5b610554600480803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803573ffffffffffffffffffffffffffffffffffffffff1690602001909190803590602001909190505061122f565b604051808215151515815260200191505060405180910390f35b341561057957600080fd5b6105c9600480803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509190505061134b565b604051808215151515815260200191505060405180910390f35b34156105ee57600080fd5b6105f6611484565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b600660009054906101000a900460ff1681565b6040805190810160405280600781526020017f73746f70706572000000000000000000000000000000000000000000000000008152506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614806106e157506106e081610b5e565b5b15156106ec57600080fd5b6001600360006101000a81548160ff02191690831515021790555050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561076557600080fd5b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed946000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a150565b60006040805190810160405280600b81526020017f757365724d616e616765720000000000000000000000000000000000000000008152506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614806108c457506108c381610b5e565b5b15156108cf57600080fd5b600360019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663c5c036996000604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b151561095d57600080fd5b6102c65a03f1151561096e57600080fd5b5050506040518051905073ffffffffffffffffffffffffffffffffffffffff16636aad7c4a30856000604051602001526040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001827cffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167cffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200192505050602060405180830381600087803b1515610a6257600080fd5b6102c65a03f11515610a7357600080fd5b50505060405180519050915060ff600780549050101515610a9357600080fd5b60078054806001018281610aa791906114aa565b9160005260206000209001600084909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550507f4bb72edbb6b4f4050e363e5db675f0f33dd63cabacdec731c7bd6fd0b87b7f8782604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a1505050565b6000610b698261134b565b8015610cd55750600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638b51ca42600254846040518082805190602001908083835b602083101515610be75780518252602082019150602081019050602083039250610bc2565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020336000604051602001526040518463ffffffff167c010000000000000000000000000000000000000000000000000000000002815260040180846000191660001916815260200183600019166000191681526020018273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019350505050602060405180830381600087803b1515610cb957600080fd5b6102c65a03f11515610cca57600080fd5b505050604051805190505b9050919050565b6040805190810160405280600981526020017f72657374617274657200000000000000000000000000000000000000000000008152506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480610d725750610d7181610b5e565b5b1515610d7d57600080fd5b6000600360006101000a81548160ff02191690831515021790555050565b6000806000803393507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9250600091505b600780549050821015610e5c578373ffffffffffffffffffffffffffffffffffffffff16600783815481101515610dff57fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610e4f57819250610e5c565b8180600101925050610dcc565b600083121515610f77576001600780549050039050600781815481101515610e8057fe5b906000526020600020900160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600784815481101515610ebb57fe5b906000526020600020900160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600781610f1291906114d6565b507fb79624962e0e32fb056c1f0aef3a7ec011949ae012e85f598f00c445d8b23c2284604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390a15b50505050565b6000600780549050905090565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600360009054906101000a900460ff1681565b600781815481101515610fd257fe5b90600052602060002090016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60025481565b60058054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156110c35780601f10611098576101008083540402835291602001916110c3565b820191906000526020600020905b8154815290600101906020018083116110a657829003601f168201915b505050505081565b6040805190810160405280600b81526020017f757365724d616e6167657200000000000000000000000000000000000000000081525061110a81610b5e565b151561111557600080fd5b81600660006101000a81548160ff0219169083151502179055505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561118e57600080fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff16141515156111eb57600080fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60006040805190810160405280600581526020017f61646d696e00000000000000000000000000000000000000000000000000000081525061127081610b5e565b151561127b57600080fd5b8473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb85856000604051602001526040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b151561132657600080fd5b6102c65a03f1151561133757600080fd5b505050604051805190509150509392505050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633037cea3600254846040518082805190602001908083835b6020831015156113c4578051825260208201915060208101905060208303925061139f565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390206000604051602001526040518363ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401808360001916600019168152602001826000191660001916815260200192505050602060405180830381600087803b151561146257600080fd5b6102c65a03f1151561147357600080fd5b505050604051805190509050919050565b600360019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b8154818355818115116114d1578183600052602060002091820191016114d09190611502565b5b505050565b8154818355818115116114fd578183600052602060002091820191016114fc9190611502565b5b505050565b61152491905b80821115611520576000816000905550600101611508565b5090565b905600a165627a7a723058204b550ca0f19a0ddc858a443b08f9e1ba12ae5056b333627e300e65dc38d07bc20029"`

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

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() constant returns(bool)
func (_UserDirectory *UserDirectoryCaller) Active(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _UserDirectory.contract.Call(opts, out, "active")
	return *ret0, err
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() constant returns(bool)
func (_UserDirectory *UserDirectorySession) Active() (bool, error) {
	return _UserDirectory.Contract.Active(&_UserDirectory.CallOpts)
}

// Active is a free data retrieval call binding the contract method 0x02fb0c5e.
//
// Solidity: function active() constant returns(bool)
func (_UserDirectory *UserDirectoryCallerSession) Active() (bool, error) {
	return _UserDirectory.Contract.Active(&_UserDirectory.CallOpts)
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

// AddWallet is a paid mutator transaction binding the contract method 0x186caa30.
//
// Solidity: function addWallet(currency bytes3) returns()
func (_UserDirectory *UserDirectoryTransactor) AddWallet(opts *bind.TransactOpts, currency [3]byte) (*types.Transaction, error) {
	return _UserDirectory.contract.Transact(opts, "addWallet", currency)
}

// AddWallet is a paid mutator transaction binding the contract method 0x186caa30.
//
// Solidity: function addWallet(currency bytes3) returns()
func (_UserDirectory *UserDirectorySession) AddWallet(currency [3]byte) (*types.Transaction, error) {
	return _UserDirectory.Contract.AddWallet(&_UserDirectory.TransactOpts, currency)
}

// AddWallet is a paid mutator transaction binding the contract method 0x186caa30.
//
// Solidity: function addWallet(currency bytes3) returns()
func (_UserDirectory *UserDirectoryTransactorSession) AddWallet(currency [3]byte) (*types.Transaction, error) {
	return _UserDirectory.Contract.AddWallet(&_UserDirectory.TransactOpts, currency)
}

// RemoveWallet is a paid mutator transaction binding the contract method 0x2286975c.
//
// Solidity: function removeWallet() returns()
func (_UserDirectory *UserDirectoryTransactor) RemoveWallet(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserDirectory.contract.Transact(opts, "removeWallet")
}

// RemoveWallet is a paid mutator transaction binding the contract method 0x2286975c.
//
// Solidity: function removeWallet() returns()
func (_UserDirectory *UserDirectorySession) RemoveWallet() (*types.Transaction, error) {
	return _UserDirectory.Contract.RemoveWallet(&_UserDirectory.TransactOpts)
}

// RemoveWallet is a paid mutator transaction binding the contract method 0x2286975c.
//
// Solidity: function removeWallet() returns()
func (_UserDirectory *UserDirectoryTransactorSession) RemoveWallet() (*types.Transaction, error) {
	return _UserDirectory.Contract.RemoveWallet(&_UserDirectory.TransactOpts)
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

// SetActive is a paid mutator transaction binding the contract method 0xacec338a.
//
// Solidity: function setActive(active_ bool) returns()
func (_UserDirectory *UserDirectoryTransactor) SetActive(opts *bind.TransactOpts, active_ bool) (*types.Transaction, error) {
	return _UserDirectory.contract.Transact(opts, "setActive", active_)
}

// SetActive is a paid mutator transaction binding the contract method 0xacec338a.
//
// Solidity: function setActive(active_ bool) returns()
func (_UserDirectory *UserDirectorySession) SetActive(active_ bool) (*types.Transaction, error) {
	return _UserDirectory.Contract.SetActive(&_UserDirectory.TransactOpts, active_)
}

// SetActive is a paid mutator transaction binding the contract method 0xacec338a.
//
// Solidity: function setActive(active_ bool) returns()
func (_UserDirectory *UserDirectoryTransactorSession) SetActive(active_ bool) (*types.Transaction, error) {
	return _UserDirectory.Contract.SetActive(&_UserDirectory.TransactOpts, active_)
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

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(token address, to address, value uint256) returns(bool)
func (_UserDirectory *UserDirectoryTransactor) Transfer(opts *bind.TransactOpts, token common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _UserDirectory.contract.Transact(opts, "transfer", token, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(token address, to address, value uint256) returns(bool)
func (_UserDirectory *UserDirectorySession) Transfer(token common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _UserDirectory.Contract.Transfer(&_UserDirectory.TransactOpts, token, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(token address, to address, value uint256) returns(bool)
func (_UserDirectory *UserDirectoryTransactorSession) Transfer(token common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _UserDirectory.Contract.Transfer(&_UserDirectory.TransactOpts, token, to, value)
}

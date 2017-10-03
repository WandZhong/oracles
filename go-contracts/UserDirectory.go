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
const UserDirectoryABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"walletCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"brg\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"wallets\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"kyc\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wallet_\",\"type\":\"address\"}],\"name\":\"removeWallet\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"profile\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"kyc_\",\"type\":\"bool\"}],\"name\":\"setKYC\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currency\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes3\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"root\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wallet_\",\"type\":\"address\"}],\"name\":\"addWallet\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"root_\",\"type\":\"address\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"foo\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"bar\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fax\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"UnauthorizedAccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newWallet\",\"type\":\"address\"}],\"name\":\"LogWalletAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"removedWallet\",\"type\":\"address\"}],\"name\":\"LogWalletRemoved\",\"type\":\"event\"}]"

// UserDirectoryBin is the compiled bytecode used for deploying new contracts.
const UserDirectoryBin = `"0x60606040526004805462ffffff1916625553441763ff00000019169055341561002757600080fd5b604051602080610d2e833981016040528080519150505b5b60018054600160a060020a03191633600160a060020a03169081179091557fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a25b600160a060020a038116151561009f57600080fd5b60028054600160a060020a031916600160a060020a0383811691909117918290556001549181169163f1ed4eda9116306040517c010000000000000000000000000000000000000000000000000000000063ffffffff8516028152600160a060020a03928316600482015291166024820152604401600060405180830381600087803b151561012d57600080fd5b6102c65a03f1151561013e57600080fd5b5050505b505b610bdb806101536000396000f300606060405236156100d55763ffffffff60e060020a60003504166307da68f581146100da57806313af4035146100ef57806329b57c69146101105780634f9c8fe81461013c57806375f12b211461016b5780637a9e5e4b146101925780637ad71f72146101b35780638da5cb5b146101e557806390d6b45f14610214578063a75fe8e11461023b578063ab60636c1461025c578063bb05c30e146102e7578063be9a655514610301578063bf7e214f14610316578063e5a6b10f14610345578063ebf0c7171461038d578063efeb5f1f146103bc575b600080fd5b34156100e557600080fd5b6100ed6103dd565b005b34156100fa57600080fd5b6100ed600160a060020a036004351661048d565b005b341561011b57600080fd5b61012361050b565b60405163ffffffff909116815260200160405180910390f35b341561014757600080fd5b61014f610512565b604051600160a060020a03909116815260200160405180910390f35b341561017657600080fd5b61017e61057c565b604051901515815260200160405180910390f35b341561019d57600080fd5b6100ed600160a060020a036004351661059d565b005b34156101be57600080fd5b61014f60043561061b565b604051600160a060020a03909116815260200160405180910390f35b34156101f057600080fd5b61014f61064d565b604051600160a060020a03909116815260200160405180910390f35b341561021f57600080fd5b61017e61065c565b604051901515815260200160405180910390f35b341561024657600080fd5b6100ed600160a060020a036004351661066c565b005b341561026757600080fd5b61026f6107a6565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156102ac5780820151818401525b602001610293565b50505050905090810190601f1680156102d95780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156102f257600080fd5b6100ed6004351515610844565b005b341561030c57600080fd5b6100ed61085f565b005b341561032157600080fd5b61014f6108f8565b604051600160a060020a03909116815260200160405180910390f35b341561035057600080fd5b610358610907565b6040517fffffff0000000000000000000000000000000000000000000000000000000000909116815260200160405180910390f35b341561039857600080fd5b61014f61092d565b604051600160a060020a03909116815260200160405180910390f35b34156103c757600080fd5b6100ed600160a060020a036004351661093c565b005b6103f333600035600160e060020a0319166109dc565b15156103fb57fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a46001805474ff00000000000000000000000000000000000000001916740100000000000000000000000000000000000000001790555b5b50505b565b6104a333600035600160e060020a0319166109dc565b15156104ab57fe5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a25b5b50565b6005545b90565b600254600090600160a060020a0316634f9c8fe882604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561055c57600080fd5b6102c65a03f1151561056d57600080fd5b50505060405180519150505b90565b60015474010000000000000000000000000000000000000000900460ff1681565b6105b333600035600160e060020a0319166109dc565b15156105bb57fe5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167f1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada460405160405180910390a25b5b50565b600580548290811061062957fe5b906000526020600020900160005b915054906101000a9004600160a060020a031681565b600154600160a060020a031681565b6004546301000000900460ff1681565b6000196000805b6005548210156106d55783600160a060020a031660058381548110151561069657fe5b906000526020600020900160005b9054906101000a9004600160a060020a0316600160a060020a031614156106c9578192505b5b600190910190610673565b6000831261079f57506005805460001981019190829081106106f357fe5b906000526020600020900160005b9054906101000a9004600160a060020a031660058481548110151561072257fe5b906000526020600020900160005b6101000a815481600160a060020a030219169083600160a060020a03160217905550806005816107609190610b3a565b507fb79624962e0e32fb056c1f0aef3a7ec011949ae012e85f598f00c445d8b23c2284604051600160a060020a03909116815260200160405180910390a15b5b50505050565b60038054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561083c5780601f106108115761010080835404028352916020019161083c565b820191906000526020600020905b81548152906001019060200180831161081f57829003601f168201915b505050505081565b6004805463ff00000019166301000000831515021790555b50565b61087533600035600160e060020a0319166109dc565b151561087d57fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a46001805474ff0000000000000000000000000000000000000000191690555b5b50505b565b600054600160a060020a031681565b6004547d0100000000000000000000000000000000000000000000000000000000000281565b600254600160a060020a031681565b80600160a060020a031633600160a060020a031614151561095c57600080fd5b600580546001810161096e8382610b3a565b916000526020600020900160005b8154600160a060020a038086166101009390930a92830292021916179055507f4bb72edbb6b4f4050e363e5db675f0f33dd63cabacdec731c7bd6fd0b87b7f8781604051600160a060020a03909116815260200160405180910390a15b50565b600030600160a060020a031683600160a060020a03161415610a0057506001610b31565b600154600160a060020a038481169116148015610a265750600054600160a060020a0316155b15610a3357506001610b31565b600054600160a060020a03161515610a9b577feb91385a0d70e44b915093d9ddfe6c8b41f2c56729431b1405f40aee5d874be78383604051600160a060020a039092168252600160e060020a03191660208201526040908101905180910390a1506000610b31565b60008054600160a060020a03169063b7009613908590309086906040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152600160e060020a03199091166044820152606401602060405180830381600087803b1515610b1457600080fd5b6102c65a03f11515610b2557600080fd5b50505060405180519150505b5b5b5b92915050565b815481835581811511610b5e57600083815260209020610b5e918101908301610b8e565b5b505050565b815481835581811511610b5e57600083815260209020610b5e918101908301610b8e565b5b505050565b61050f91905b80821115610ba85760008155600101610b94565b5090565b905600a165627a7a723058201ab217c29fbdfafda43ee1d05ec5beab3352d3a07e4fd15b2c5d2f898390f8220029"`

// DeployUserDirectory deploys a new Ethereum contract, binding an instance of UserDirectory to it.
func DeployUserDirectory(auth *bind.TransactOpts, backend bind.ContractBackend, root_ common.Address) (common.Address, *types.Transaction, *UserDirectory, error) {
	parsed, err := abi.JSON(strings.NewReader(UserDirectoryABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UserDirectoryBin), backend, root_)
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

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_UserDirectory *UserDirectoryCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _UserDirectory.contract.Call(opts, out, "authority")
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_UserDirectory *UserDirectorySession) Authority() (common.Address, error) {
	return _UserDirectory.Contract.Authority(&_UserDirectory.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_UserDirectory *UserDirectoryCallerSession) Authority() (common.Address, error) {
	return _UserDirectory.Contract.Authority(&_UserDirectory.CallOpts)
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

// AddWallet is a paid mutator transaction binding the contract method 0xefeb5f1f.
//
// Solidity: function addWallet(wallet_ address) returns()
func (_UserDirectory *UserDirectoryTransactor) AddWallet(opts *bind.TransactOpts, wallet_ common.Address) (*types.Transaction, error) {
	return _UserDirectory.contract.Transact(opts, "addWallet", wallet_)
}

// AddWallet is a paid mutator transaction binding the contract method 0xefeb5f1f.
//
// Solidity: function addWallet(wallet_ address) returns()
func (_UserDirectory *UserDirectorySession) AddWallet(wallet_ common.Address) (*types.Transaction, error) {
	return _UserDirectory.Contract.AddWallet(&_UserDirectory.TransactOpts, wallet_)
}

// AddWallet is a paid mutator transaction binding the contract method 0xefeb5f1f.
//
// Solidity: function addWallet(wallet_ address) returns()
func (_UserDirectory *UserDirectoryTransactorSession) AddWallet(wallet_ common.Address) (*types.Transaction, error) {
	return _UserDirectory.Contract.AddWallet(&_UserDirectory.TransactOpts, wallet_)
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

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_UserDirectory *UserDirectoryTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _UserDirectory.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_UserDirectory *UserDirectorySession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _UserDirectory.Contract.SetAuthority(&_UserDirectory.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_UserDirectory *UserDirectoryTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _UserDirectory.Contract.SetAuthority(&_UserDirectory.TransactOpts, authority_)
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

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_UserDirectory *UserDirectoryTransactor) Start(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserDirectory.contract.Transact(opts, "start")
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_UserDirectory *UserDirectorySession) Start() (*types.Transaction, error) {
	return _UserDirectory.Contract.Start(&_UserDirectory.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_UserDirectory *UserDirectoryTransactorSession) Start() (*types.Transaction, error) {
	return _UserDirectory.Contract.Start(&_UserDirectory.TransactOpts)
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

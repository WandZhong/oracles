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

// WalletABI is the input ABI used to generate the binding from.
const WalletABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vault_\",\"type\":\"address\"}],\"name\":\"addVault\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"listVaults\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"addAsset\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"cleanStorage\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"vaults\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"assetsLen\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"vaultCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"remove\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"directory\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"removeVault\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assets\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"}],\"name\":\"rmAsset\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hasFunds\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"directory_\",\"type\":\"address\"}],\"payable\":false,\"type\":\"constructor\"},{\"payable\":true,\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"LogVaultAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"LogVaultRemoved\",\"type\":\"event\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"foo\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"bar\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fax\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"AssetAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"}],\"name\":\"AssetRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"UnauthorizedAccess\",\"type\":\"event\"}]"

// WalletBin is the compiled bytecode used for deploying new contracts.
const WalletBin = `"0x6060604052341561000f57600080fd5b604051602080611695833981016040528080519150505b5b60018054600160a060020a03191633600160a060020a03169081179091557fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a25b6003805461010060a860020a031916610100600160a060020a038481168202929092179283905590910416638da5cb5b6000604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15156100f257600080fd5b6102c65a03f1151561010357600080fd5b505050604051805160018054600160a060020a031916600160a060020a03928316179055600354610100900416905063efeb5f1f306040517c010000000000000000000000000000000000000000000000000000000063ffffffff8416028152600160a060020a039091166004820152602401600060405180830381600087803b151561018f57600080fd5b6102c65a03f115156101a057600080fd5b5050505b505b6114e0806101b56000396000f300606060405236156101175763ffffffff60e060020a60003504166307da68f5811461011b57806313af403514610130578063256b5a021461015157806350cc258e1461017257806356e596a8146101d957806370a0823114610215578063712e62791461024657806375f12b21146102825780637a9e5e4b146102a95780637bb98a68146102ca5780637e84d36b146103785780638c64ea4a1461038d5780638da5cb5b146103bf578063a37685e5146103ee578063a7c6a10014610413578063a7f437791461043f578063be9a655514610454578063bf7e214f14610469578063c41c2f2414610498578063ceb68c23146104c7578063cf35bdd0146104e8578063d1da3e7b1461051a578063f2cd59f614610553575b5b5b005b341561012657600080fd5b61011761057a565b005b341561013b57600080fd5b610117600160a060020a0360043516610602565b005b341561015c57600080fd5b610117600160a060020a0360043516610680565b005b341561017d57600080fd5b610185610720565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156101c55780820151818401525b6020016101ac565b505050509050019250505060405180910390f35b34156101e457600080fd5b610117600160a060020a03600435811690602435166fffffffffffffffffffffffffffffffff604435166107db565b005b341561022057600080fd5b610234600160a060020a036004351661088a565b60405190815260200160405180910390f35b341561025157600080fd5b610117600160a060020a03600435811690602435166fffffffffffffffffffffffffffffffff60443516610905565b005b341561028d57600080fd5b610295610a25565b604051901515815260200160405180910390f35b34156102b457600080fd5b610117600160a060020a0360043516610a2e565b005b34156102d557600080fd5b6102dd610aac565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b838110156103225780820151818401525b602001610309565b50505050905001838103825284818151815260200191508051906020019060200280838360005b838110156103625780820151818401525b602001610349565b5050505090500194505050505060405180910390f35b341561038357600080fd5b610117610acf565b005b341561039857600080fd5b6103a3600435610b55565b604051600160a060020a03909116815260200160405180910390f35b34156103ca57600080fd5b6103a3610b87565b604051600160a060020a03909116815260200160405180910390f35b34156103f957600080fd5b610234610b96565b60405190815260200160405180910390f35b341561041e57600080fd5b610426610c16565b60405163ffffffff909116815260200160405180910390f35b341561044a57600080fd5b610117610c1d565b005b341561045f57600080fd5b610117610cb4565b005b341561047457600080fd5b6103a3610d39565b604051600160a060020a03909116815260200160405180910390f35b34156104a357600080fd5b6103a3610d48565b604051600160a060020a03909116815260200160405180910390f35b34156104d257600080fd5b610117600160a060020a0360043516610d5c565b005b34156104f357600080fd5b6103a3600435610f57565b604051600160a060020a03909116815260200160405180910390f35b341561052557600080fd5b610295600160a060020a0360043581169060243516610f89565b604051901515815260200160405180910390f35b341561055e57600080fd5b61029561108a565b604051901515815260200160405180910390f35b61059033600035600160e060020a031916611104565b151561059857fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a46003805460ff191660011790555b5b50505b565b61061833600035600160e060020a031916611104565b151561062057fe5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a25b5b50565b80600160a060020a031633600160a060020a03161415156106a057600080fd5b60048054600181016106b2838261141b565b916000526020600020900160005b8154600160a060020a038086166101009390930a92830292021916179055507fefaffcc73be9b765a62202dfb131e611b0309bd65ee334807b295db1ecb1b54f81604051600160a060020a03909116815260200160405180910390a15b50565b610728611445565b610730611445565b6004546000906040518059106107435750595b908082528060200260200182016040525b509150600090505b60045463ffffffff821610156107d2576004805463ffffffff831690811061078057fe5b906000526020600020900160005b9054906101000a9004600160a060020a0316828263ffffffff16815181106107b257fe5b600160a060020a039092166020928302909101909101525b60010161075c565b8192505b505090565b6107f133600035600160e060020a031916611104565b15156107f957fe5b82600160a060020a031663a9059cbb838360006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526fffffffffffffffffffffffffffffffff166024820152604401602060405180830381600087803b151561086857600080fd5b6102c65a03f1151561087957600080fd5b505050604051805150505b5b505050565b600081600160a060020a03166370a082313060006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b15156108e357600080fd5b6102c65a03f115156108f457600080fd5b50505060405180519150505b919050565b61091b33600035600160e060020a031916611104565b151561092357fe5b73__AssetsLib_____________________________6352d28bc8600285858560405160e060020a63ffffffff87160281526004810194909452600160a060020a039283166024850152911660448301526fffffffffffffffffffffffffffffffff16606482015260840160006040518083038186803b15156109a457600080fd5b6102c65a03f415156109b557600080fd5b5050507fd0d5f5786af393a6465a13be9c33dbd3ee9c29b1c32872f8aa0cd405ee62f8fe838383604051600160a060020a0393841681529190921660208201526fffffffffffffffffffffffffffffffff9091166040808301919091526060909101905180910390a15b5b505050565b60035460ff1681565b610a4433600035600160e060020a031916611104565b1515610a4c57fe5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167f1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada460405160405180910390a25b5b50565b610ab4611445565b610abc611445565b610ac66002611267565b915091505b9091565b610ae533600035600160e060020a031916611104565b1515610aed57fe5b73__AssetsLib_____________________________6376c86cd7600260405160e060020a63ffffffff8416028152600481019190915260240160006040518083038186803b1515610b3d57600080fd5b6102c65a03f4151561088457600080fd5b5050505b5b565b6004805482908110610b6357fe5b906000526020600020900160005b915054906101000a9004600160a060020a031681565b600154600160a060020a031681565b600073__AssetsLib_____________________________63922ddcf26002836040516020015260405160e060020a63ffffffff8416028152600481019190915260240160206040518083038186803b1515610bf057600080fd5b6102c65a03f41515610c0157600080fd5b505050604051805163ffffffff169150505b90565b6004545b90565b600160a060020a0330163115610c3257600080fd5b610c3a61108a565b15610c4457600080fd5b6003546101009004600160a060020a031663a75fe8e13060405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401600060405180830381600087803b1515610c9957600080fd5b6102c65a03f11515610caa57600080fd5b506000915050ff5b565b610cca33600035600160e060020a031916611104565b1515610cd257fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a46003805460ff191690555b5b50505b565b600054600160a060020a031681565b6003546101009004600160a060020a031681565b600080805b60045463ffffffff83161015610dce576004805463ffffffff8416908110610d8557fe5b906000526020600020900160005b9054906101000a9004600160a060020a0316600160a060020a031684600160a060020a03161415610dc2578192505b5b816001019150610d61565b60045463ffffffff841610610de257600080fd5b6004805463ffffffff8516908110610df657fe5b906000526020600020900160005b9054906101000a9004600160a060020a0316600160a060020a031633600160a060020a0316141515610e3557600080fd5b506004546000190163ffffffff831681901115610e5157600080fd5b7fe5b46e8ce81c693c1c4aaac9f27e523da69aab82fb2459d7819867c01869d99160048463ffffffff16815481101515610e8757fe5b906000526020600020900160005b9054906101000a9004600160a060020a0316604051600160a060020a03909116815260200160405180910390a163ffffffff83168114610f43576004805482908110610edd57fe5b906000526020600020900160005b9054906101000a9004600160a060020a031660048463ffffffff16815481101515610f1257fe5b906000526020600020900160005b6101000a815481600160a060020a030219169083600160a060020a031602179055505b80610f4f60048261141b565b505b50505050565b6002805482908110610b6357fe5b906000526020600020900160005b915054906101000a9004600160a060020a031681565b6000610fa133600035600160e060020a031916611104565b1515610fa957fe5b73__AssetsLib_____________________________638a7fcb2d6002858560006040516020015260405160e060020a63ffffffff86160281526004810193909352600160a060020a03918216602484015216604482015260640160206040518083038186803b151561101a57600080fd5b6102c65a03f4151561102b57600080fd5b505050604051805190501561107f577f37803e2125c48ee96c38ddf04e826daf335b0e1603579040fd275aba6d06b6fc83604051600160a060020a03909116815260200160405180910390a1506001611083565b5060005b5b92915050565b600073__AssetsLib_____________________________632a6a878f6002836040516020015260405160e060020a63ffffffff8416028152600481019190915260240160206040518083038186803b15156110e457600080fd5b6102c65a03f415156110f557600080fd5b50505060405180519150505b90565b600030600160a060020a031683600160a060020a0316141561112857506001611083565b600154600160a060020a03848116911614801561114e5750600054600160a060020a0316155b1561115b57506001611083565b600054600160a060020a031615156111c3577feb91385a0d70e44b915093d9ddfe6c8b41f2c56729431b1405f40aee5d874be78383604051600160a060020a039092168252600160e060020a03191660208201526040908101905180910390a1506000611083565b60008054600160a060020a03169063b7009613908590309086906040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152600160e060020a03199091166044820152606401602060405180830381600087803b151561123c57600080fd5b6102c65a03f1151561124d57600080fd5b505050604051805190509050611083565b5b5b5b92915050565b61126f611445565b611277611445565b61127f611445565b611287611445565b84546000906040518059106112995750595b908082528060200260200182016040525b5086549093506040518059106112bd5750595b908082528060200260200182016040525b509150600090505b855463ffffffff8216101561140c57858163ffffffff168154811015156112f957fe5b906000526020600020900160005b9054906101000a9004600160a060020a0316838263ffffffff168151811061132b57fe5b600160a060020a039092166020928302909101909101528554869063ffffffff831690811061135657fe5b906000526020600020900160005b9054906101000a9004600160a060020a0316600160a060020a03166370a082313060006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b15156113cc57600080fd5b6102c65a03f115156113dd57600080fd5b50505060405180519050828263ffffffff16815181106113f957fe5b602090810290910101525b6001016112d6565b8282945094505b505050915091565b81548183558181151161088457600083815260209020610884918101908301611493565b5b505050565b60206040519081016040526000815290565b60206040519081016040526000815290565b81548183558181151161088457600083815260209020610884918101908301611493565b5b505050565b610c1391905b808211156114ad5760008155600101611499565b5090565b905600a165627a7a7230582049917899d958ccc10564329eb835d47a88ac78b023076d73e8c31984d3f5fae00029"`

// DeployWallet deploys a new Ethereum contract, binding an instance of Wallet to it.
func DeployWallet(auth *bind.TransactOpts, backend bind.ContractBackend, directory_ common.Address) (common.Address, *types.Transaction, *Wallet, error) {
	parsed, err := abi.JSON(strings.NewReader(WalletABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WalletBin), backend, directory_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Wallet{WalletCaller: WalletCaller{contract: contract}, WalletTransactor: WalletTransactor{contract: contract}}, nil
}

// Wallet is an auto generated Go binding around an Ethereum contract.
type Wallet struct {
	WalletCaller     // Read-only binding to the contract
	WalletTransactor // Write-only binding to the contract
}

// WalletCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletSession struct {
	Contract     *Wallet           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletCallerSession struct {
	Contract *WalletCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WalletTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletTransactorSession struct {
	Contract     *WalletTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletRaw struct {
	Contract *Wallet // Generic contract binding to access the raw methods on
}

// WalletCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletCallerRaw struct {
	Contract *WalletCaller // Generic read-only contract binding to access the raw methods on
}

// WalletTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletTransactorRaw struct {
	Contract *WalletTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWallet creates a new instance of Wallet, bound to a specific deployed contract.
func NewWallet(address common.Address, backend bind.ContractBackend) (*Wallet, error) {
	contract, err := bindWallet(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Wallet{WalletCaller: WalletCaller{contract: contract}, WalletTransactor: WalletTransactor{contract: contract}}, nil
}

// NewWalletCaller creates a new read-only instance of Wallet, bound to a specific deployed contract.
func NewWalletCaller(address common.Address, caller bind.ContractCaller) (*WalletCaller, error) {
	contract, err := bindWallet(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &WalletCaller{contract: contract}, nil
}

// NewWalletTransactor creates a new write-only instance of Wallet, bound to a specific deployed contract.
func NewWalletTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletTransactor, error) {
	contract, err := bindWallet(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &WalletTransactor{contract: contract}, nil
}

// bindWallet binds a generic wrapper to an already deployed contract.
func bindWallet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WalletABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wallet *WalletRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Wallet.Contract.WalletCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wallet *WalletRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.Contract.WalletTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wallet *WalletRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wallet.Contract.WalletTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wallet *WalletCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Wallet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wallet *WalletTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wallet *WalletTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wallet.Contract.contract.Transact(opts, method, params...)
}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets( uint256) constant returns(address)
func (_Wallet *WalletCaller) Assets(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "assets", arg0)
	return *ret0, err
}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets( uint256) constant returns(address)
func (_Wallet *WalletSession) Assets(arg0 *big.Int) (common.Address, error) {
	return _Wallet.Contract.Assets(&_Wallet.CallOpts, arg0)
}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets( uint256) constant returns(address)
func (_Wallet *WalletCallerSession) Assets(arg0 *big.Int) (common.Address, error) {
	return _Wallet.Contract.Assets(&_Wallet.CallOpts, arg0)
}

// AssetsLen is a free data retrieval call binding the contract method 0xa37685e5.
//
// Solidity: function assetsLen() constant returns(uint256)
func (_Wallet *WalletCaller) AssetsLen(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "assetsLen")
	return *ret0, err
}

// AssetsLen is a free data retrieval call binding the contract method 0xa37685e5.
//
// Solidity: function assetsLen() constant returns(uint256)
func (_Wallet *WalletSession) AssetsLen() (*big.Int, error) {
	return _Wallet.Contract.AssetsLen(&_Wallet.CallOpts)
}

// AssetsLen is a free data retrieval call binding the contract method 0xa37685e5.
//
// Solidity: function assetsLen() constant returns(uint256)
func (_Wallet *WalletCallerSession) AssetsLen() (*big.Int, error) {
	return _Wallet.Contract.AssetsLen(&_Wallet.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Wallet *WalletCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "authority")
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Wallet *WalletSession) Authority() (common.Address, error) {
	return _Wallet.Contract.Authority(&_Wallet.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Wallet *WalletCallerSession) Authority() (common.Address, error) {
	return _Wallet.Contract.Authority(&_Wallet.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(token address) constant returns(uint256)
func (_Wallet *WalletCaller) BalanceOf(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "balanceOf", token)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(token address) constant returns(uint256)
func (_Wallet *WalletSession) BalanceOf(token common.Address) (*big.Int, error) {
	return _Wallet.Contract.BalanceOf(&_Wallet.CallOpts, token)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(token address) constant returns(uint256)
func (_Wallet *WalletCallerSession) BalanceOf(token common.Address) (*big.Int, error) {
	return _Wallet.Contract.BalanceOf(&_Wallet.CallOpts, token)
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(address[], uint256[])
func (_Wallet *WalletCaller) Balances(opts *bind.CallOpts) ([]common.Address, []*big.Int, error) {
	var (
		ret0 = new([]common.Address)
		ret1 = new([]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Wallet.contract.Call(opts, out, "balances")
	return *ret0, *ret1, err
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(address[], uint256[])
func (_Wallet *WalletSession) Balances() ([]common.Address, []*big.Int, error) {
	return _Wallet.Contract.Balances(&_Wallet.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(address[], uint256[])
func (_Wallet *WalletCallerSession) Balances() ([]common.Address, []*big.Int, error) {
	return _Wallet.Contract.Balances(&_Wallet.CallOpts)
}

// Directory is a free data retrieval call binding the contract method 0xc41c2f24.
//
// Solidity: function directory() constant returns(address)
func (_Wallet *WalletCaller) Directory(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "directory")
	return *ret0, err
}

// Directory is a free data retrieval call binding the contract method 0xc41c2f24.
//
// Solidity: function directory() constant returns(address)
func (_Wallet *WalletSession) Directory() (common.Address, error) {
	return _Wallet.Contract.Directory(&_Wallet.CallOpts)
}

// Directory is a free data retrieval call binding the contract method 0xc41c2f24.
//
// Solidity: function directory() constant returns(address)
func (_Wallet *WalletCallerSession) Directory() (common.Address, error) {
	return _Wallet.Contract.Directory(&_Wallet.CallOpts)
}

// HasFunds is a free data retrieval call binding the contract method 0xf2cd59f6.
//
// Solidity: function hasFunds() constant returns(bool)
func (_Wallet *WalletCaller) HasFunds(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "hasFunds")
	return *ret0, err
}

// HasFunds is a free data retrieval call binding the contract method 0xf2cd59f6.
//
// Solidity: function hasFunds() constant returns(bool)
func (_Wallet *WalletSession) HasFunds() (bool, error) {
	return _Wallet.Contract.HasFunds(&_Wallet.CallOpts)
}

// HasFunds is a free data retrieval call binding the contract method 0xf2cd59f6.
//
// Solidity: function hasFunds() constant returns(bool)
func (_Wallet *WalletCallerSession) HasFunds() (bool, error) {
	return _Wallet.Contract.HasFunds(&_Wallet.CallOpts)
}

// ListVaults is a free data retrieval call binding the contract method 0x50cc258e.
//
// Solidity: function listVaults() constant returns(address[])
func (_Wallet *WalletCaller) ListVaults(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "listVaults")
	return *ret0, err
}

// ListVaults is a free data retrieval call binding the contract method 0x50cc258e.
//
// Solidity: function listVaults() constant returns(address[])
func (_Wallet *WalletSession) ListVaults() ([]common.Address, error) {
	return _Wallet.Contract.ListVaults(&_Wallet.CallOpts)
}

// ListVaults is a free data retrieval call binding the contract method 0x50cc258e.
//
// Solidity: function listVaults() constant returns(address[])
func (_Wallet *WalletCallerSession) ListVaults() ([]common.Address, error) {
	return _Wallet.Contract.ListVaults(&_Wallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Wallet *WalletCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Wallet *WalletSession) Owner() (common.Address, error) {
	return _Wallet.Contract.Owner(&_Wallet.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Wallet *WalletCallerSession) Owner() (common.Address, error) {
	return _Wallet.Contract.Owner(&_Wallet.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_Wallet *WalletCaller) Stopped(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "stopped")
	return *ret0, err
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_Wallet *WalletSession) Stopped() (bool, error) {
	return _Wallet.Contract.Stopped(&_Wallet.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_Wallet *WalletCallerSession) Stopped() (bool, error) {
	return _Wallet.Contract.Stopped(&_Wallet.CallOpts)
}

// VaultCount is a free data retrieval call binding the contract method 0xa7c6a100.
//
// Solidity: function vaultCount() constant returns(uint32)
func (_Wallet *WalletCaller) VaultCount(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "vaultCount")
	return *ret0, err
}

// VaultCount is a free data retrieval call binding the contract method 0xa7c6a100.
//
// Solidity: function vaultCount() constant returns(uint32)
func (_Wallet *WalletSession) VaultCount() (uint32, error) {
	return _Wallet.Contract.VaultCount(&_Wallet.CallOpts)
}

// VaultCount is a free data retrieval call binding the contract method 0xa7c6a100.
//
// Solidity: function vaultCount() constant returns(uint32)
func (_Wallet *WalletCallerSession) VaultCount() (uint32, error) {
	return _Wallet.Contract.VaultCount(&_Wallet.CallOpts)
}

// Vaults is a free data retrieval call binding the contract method 0x8c64ea4a.
//
// Solidity: function vaults( uint256) constant returns(address)
func (_Wallet *WalletCaller) Vaults(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Wallet.contract.Call(opts, out, "vaults", arg0)
	return *ret0, err
}

// Vaults is a free data retrieval call binding the contract method 0x8c64ea4a.
//
// Solidity: function vaults( uint256) constant returns(address)
func (_Wallet *WalletSession) Vaults(arg0 *big.Int) (common.Address, error) {
	return _Wallet.Contract.Vaults(&_Wallet.CallOpts, arg0)
}

// Vaults is a free data retrieval call binding the contract method 0x8c64ea4a.
//
// Solidity: function vaults( uint256) constant returns(address)
func (_Wallet *WalletCallerSession) Vaults(arg0 *big.Int) (common.Address, error) {
	return _Wallet.Contract.Vaults(&_Wallet.CallOpts, arg0)
}

// AddAsset is a paid mutator transaction binding the contract method 0x712e6279.
//
// Solidity: function addAsset(token address, src address, wad uint128) returns()
func (_Wallet *WalletTransactor) AddAsset(opts *bind.TransactOpts, token common.Address, src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "addAsset", token, src, wad)
}

// AddAsset is a paid mutator transaction binding the contract method 0x712e6279.
//
// Solidity: function addAsset(token address, src address, wad uint128) returns()
func (_Wallet *WalletSession) AddAsset(token common.Address, src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.AddAsset(&_Wallet.TransactOpts, token, src, wad)
}

// AddAsset is a paid mutator transaction binding the contract method 0x712e6279.
//
// Solidity: function addAsset(token address, src address, wad uint128) returns()
func (_Wallet *WalletTransactorSession) AddAsset(token common.Address, src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.AddAsset(&_Wallet.TransactOpts, token, src, wad)
}

// AddVault is a paid mutator transaction binding the contract method 0x256b5a02.
//
// Solidity: function addVault(vault_ address) returns()
func (_Wallet *WalletTransactor) AddVault(opts *bind.TransactOpts, vault_ common.Address) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "addVault", vault_)
}

// AddVault is a paid mutator transaction binding the contract method 0x256b5a02.
//
// Solidity: function addVault(vault_ address) returns()
func (_Wallet *WalletSession) AddVault(vault_ common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.AddVault(&_Wallet.TransactOpts, vault_)
}

// AddVault is a paid mutator transaction binding the contract method 0x256b5a02.
//
// Solidity: function addVault(vault_ address) returns()
func (_Wallet *WalletTransactorSession) AddVault(vault_ common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.AddVault(&_Wallet.TransactOpts, vault_)
}

// CleanStorage is a paid mutator transaction binding the contract method 0x7e84d36b.
//
// Solidity: function cleanStorage() returns()
func (_Wallet *WalletTransactor) CleanStorage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "cleanStorage")
}

// CleanStorage is a paid mutator transaction binding the contract method 0x7e84d36b.
//
// Solidity: function cleanStorage() returns()
func (_Wallet *WalletSession) CleanStorage() (*types.Transaction, error) {
	return _Wallet.Contract.CleanStorage(&_Wallet.TransactOpts)
}

// CleanStorage is a paid mutator transaction binding the contract method 0x7e84d36b.
//
// Solidity: function cleanStorage() returns()
func (_Wallet *WalletTransactorSession) CleanStorage() (*types.Transaction, error) {
	return _Wallet.Contract.CleanStorage(&_Wallet.TransactOpts)
}

// Remove is a paid mutator transaction binding the contract method 0xa7f43779.
//
// Solidity: function remove() returns()
func (_Wallet *WalletTransactor) Remove(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "remove")
}

// Remove is a paid mutator transaction binding the contract method 0xa7f43779.
//
// Solidity: function remove() returns()
func (_Wallet *WalletSession) Remove() (*types.Transaction, error) {
	return _Wallet.Contract.Remove(&_Wallet.TransactOpts)
}

// Remove is a paid mutator transaction binding the contract method 0xa7f43779.
//
// Solidity: function remove() returns()
func (_Wallet *WalletTransactorSession) Remove() (*types.Transaction, error) {
	return _Wallet.Contract.Remove(&_Wallet.TransactOpts)
}

// RemoveVault is a paid mutator transaction binding the contract method 0xceb68c23.
//
// Solidity: function removeVault(vault address) returns()
func (_Wallet *WalletTransactor) RemoveVault(opts *bind.TransactOpts, vault common.Address) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "removeVault", vault)
}

// RemoveVault is a paid mutator transaction binding the contract method 0xceb68c23.
//
// Solidity: function removeVault(vault address) returns()
func (_Wallet *WalletSession) RemoveVault(vault common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.RemoveVault(&_Wallet.TransactOpts, vault)
}

// RemoveVault is a paid mutator transaction binding the contract method 0xceb68c23.
//
// Solidity: function removeVault(vault address) returns()
func (_Wallet *WalletTransactorSession) RemoveVault(vault common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.RemoveVault(&_Wallet.TransactOpts, vault)
}

// RmAsset is a paid mutator transaction binding the contract method 0xd1da3e7b.
//
// Solidity: function rmAsset(token address, dst address) returns(bool)
func (_Wallet *WalletTransactor) RmAsset(opts *bind.TransactOpts, token common.Address, dst common.Address) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "rmAsset", token, dst)
}

// RmAsset is a paid mutator transaction binding the contract method 0xd1da3e7b.
//
// Solidity: function rmAsset(token address, dst address) returns(bool)
func (_Wallet *WalletSession) RmAsset(token common.Address, dst common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.RmAsset(&_Wallet.TransactOpts, token, dst)
}

// RmAsset is a paid mutator transaction binding the contract method 0xd1da3e7b.
//
// Solidity: function rmAsset(token address, dst address) returns(bool)
func (_Wallet *WalletTransactorSession) RmAsset(token common.Address, dst common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.RmAsset(&_Wallet.TransactOpts, token, dst)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_Wallet *WalletTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_Wallet *WalletSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.SetAuthority(&_Wallet.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_Wallet *WalletTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.SetAuthority(&_Wallet.TransactOpts, authority_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_Wallet *WalletTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_Wallet *WalletSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.SetOwner(&_Wallet.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_Wallet *WalletTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Wallet.Contract.SetOwner(&_Wallet.TransactOpts, owner_)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Wallet *WalletTransactor) Start(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "start")
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Wallet *WalletSession) Start() (*types.Transaction, error) {
	return _Wallet.Contract.Start(&_Wallet.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Wallet *WalletTransactorSession) Start() (*types.Transaction, error) {
	return _Wallet.Contract.Start(&_Wallet.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Wallet *WalletTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Wallet *WalletSession) Stop() (*types.Transaction, error) {
	return _Wallet.Contract.Stop(&_Wallet.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Wallet *WalletTransactorSession) Stop() (*types.Transaction, error) {
	return _Wallet.Contract.Stop(&_Wallet.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0x56e596a8.
//
// Solidity: function transfer(token address, dst address, wad uint128) returns()
func (_Wallet *WalletTransactor) Transfer(opts *bind.TransactOpts, token common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Wallet.contract.Transact(opts, "transfer", token, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0x56e596a8.
//
// Solidity: function transfer(token address, dst address, wad uint128) returns()
func (_Wallet *WalletSession) Transfer(token common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.Transfer(&_Wallet.TransactOpts, token, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0x56e596a8.
//
// Solidity: function transfer(token address, dst address, wad uint128) returns()
func (_Wallet *WalletTransactorSession) Transfer(token common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Wallet.Contract.Transfer(&_Wallet.TransactOpts, token, dst, wad)
}

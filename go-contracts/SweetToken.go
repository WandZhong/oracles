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

// SweetTokenABI is the input ABI used to generate the binding from.
const SweetTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guy\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"logic_\",\"type\":\"address\"}],\"name\":\"setLogic\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"pull\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name_\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"logic\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"_allowance\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"name_\",\"type\":\"string\"},{\"name\":\"symbol_\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"constructor\"},{\"payable\":true,\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"LogBurn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"who\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"LogMint\",\"type\":\"event\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"foo\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"bar\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fax\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"UnauthorizedAccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// SweetTokenBin is the compiled bytecode used for deploying new contracts.
const SweetTokenBin = `"0x6060604052601260045534156200001557600080fd5b604051620014d4380380620014d4833981016040528080518201919060200180519150505b81815b5b60018054600160a060020a03191633600160a060020a03169081179091557fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a25b60038280516200009e929160200190620000b0565b5060028190555b50505b50506200015a565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620000f357805160ff191683800117855562000123565b8280016001018555821562000123579182015b828111156200012357825182559160200191906001019062000106565b5b506200013292915062000136565b5090565b6200015791905b808211156200013257600081556001016200013d565b5090565b90565b61136a806200016a6000396000f3006060604052361561010c5763ffffffff60e060020a60003504166306fdde03811461011557806307da68f5146101a0578063095ea7b3146101b557806313af4035146101eb57806318160ddd1461020c57806323b872dd14610231578063313ce5671461026d57806369d3e20e1461029257806370a08231146102b357806371857000146102e457806375f12b21146103175780637a9e5e4b1461033e5780638402181f1461035f5780638da5cb5b1461039e57806390bc1693146103cd57806395d89b41146103ee578063a9059cbb14610413578063be9a655514610449578063bf7e214f1461045e578063c47f00271461048d578063d7dfa0dd146104e0578063dd62ed3e1461050f575b5b600080fd5b5b005b341561012057600080fd5b610128610546565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156101655780820151818401525b60200161014c565b50505050905090810190601f1680156101925780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156101ab57600080fd5b6101126105e4565b005b34156101c057600080fd5b6101d7600160a060020a0360043516602435610683565b604051901515815260200160405180910390f35b34156101f657600080fd5b610112600160a060020a03600435166107d7565b005b341561021757600080fd5b61021f610855565b60405190815260200160405180910390f35b341561023c57600080fd5b6101d7600160a060020a03600435811690602435166044356108bf565b604051901515815260200160405180910390f35b341561027857600080fd5b61021f610a0e565b60405190815260200160405180910390f35b341561029d57600080fd5b6101126001608060020a0360043516610a14565b005b34156102be57600080fd5b61021f600160a060020a0360043516610b50565b60405190815260200160405180910390f35b34156102ef57600080fd5b6101d7600160a060020a0360043516610bcd565b604051901515815260200160405180910390f35b341561032257600080fd5b6101d7610c8b565b604051901515815260200160405180910390f35b341561034957600080fd5b610112600160a060020a0360043516610c9b565b005b341561036a57600080fd5b6101d7600160a060020a03600435166001608060020a0360243516610d19565b604051901515815260200160405180910390f35b34156103a957600080fd5b6103b1610d38565b604051600160a060020a03909116815260200160405180910390f35b34156103d857600080fd5b6101126001608060020a0360043516610d47565b005b34156103f957600080fd5b61021f610e73565b60405190815260200160405180910390f35b341561041e57600080fd5b6101d7600160a060020a0360043516602435610e79565b604051901515815260200160405180910390f35b341561045457600080fd5b610112610fc7565b005b341561046957600080fd5b6103b1611060565b604051600160a060020a03909116815260200160405180910390f35b341561049857600080fd5b61011260046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061106f95505050505050565b005b34156104eb57600080fd5b6103b16110a6565b604051600160a060020a03909116815260200160405180910390f35b341561051a57600080fd5b61021f600160a060020a03600435811690602435166110b5565b60405190815260200160405180910390f35b60038054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105dc5780601f106105b1576101008083540402835291602001916105dc565b820191906000526020600020905b8154815290600101906020018083116105bf57829003601f168201915b505050505081565b6105fa33600035600160e060020a03191661113b565b151561060257fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a46001805474ff0000000000000000000000000000000000000000191660a060020a1790555b5b50505b565b600154600090819060a060020a900460ff161561069c57fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a4600554600160a060020a031663e1f21c6733888860006040516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b151561075e57600080fd5b6102c65a03f1151561076f57600080fd5b505050604051805193505082156107c85785600160a060020a031633600160a060020a03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258760405190815260200160405180910390a35b8293505b5b50505b5092915050565b6107ed33600035600160e060020a03191661113b565b15156107f557fe5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a25b5b50565b600554600090600160a060020a03166318160ddd82604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561089f57600080fd5b6102c65a03f115156108b057600080fd5b50505060405180519150505b90565b600154600090819060a060020a900460ff16156108d857fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a4600554600160a060020a03166323b872dd88888860006040516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b151561099a57600080fd5b6102c65a03f115156109ab57600080fd5b50505060405180519050925085600160a060020a031687600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8760405190815260200160405180910390a38293505b5b50505b509392505050565b60045481565b610a2a33600035600160e060020a03191661113b565b1515610a3257fe5b60015460a060020a900460ff1615610a4657fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a4600554600160a060020a03166369d3e20e8460405160e060020a63ffffffff84160281526001608060020a039091166004820152602401600060405180830381600087803b1515610aec57600080fd5b6102c65a03f11515610afd57600080fd5b50505033600160a060020a03167f9f494565851dbcb31fb5198ca217cda6833282fadb96ba9431bd19c82afc1dd3846040516001608060020a03909116815260200160405180910390a25b5b50505b5b50565b600554600090600160a060020a03166370a0823183836040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b1515610bab57600080fd5b6102c65a03f11515610bbc57600080fd5b50505060405180519150505b919050565b6000610be533600035600160e060020a03191661113b565b1515610bed57fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a4600160a060020a0384161515610c5557fe5b6005805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038616179055600192505b5b50505b919050565b60015460a060020a900460ff1681565b610cb133600035600160e060020a03191661113b565b1515610cb957fe5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167f1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada460405160405180910390a25b5b50565b6000610d2f8333846001608060020a03166108bf565b90505b92915050565b600154600160a060020a031681565b60015460a060020a900460ff1615610d5b57fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a4600554600160a060020a0316637261e469338560405160e060020a63ffffffff8516028152600160a060020a0390921660048301526001608060020a03166024820152604401600060405180830381600087803b1515610e1057600080fd5b6102c65a03f11515610e2157600080fd5b50505033600160a060020a03167f38d762ef507761291a578e921acfe29c1af31a7331ea03e391cf16cfc4d4f581846040516001608060020a03909116815260200160405180910390a25b5b50505b50565b60025481565b600154600090819060a060020a900460ff1615610e9257fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a4600554600160a060020a031663beabacc833888860006040516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b1515610f5457600080fd5b6102c65a03f11515610f6557600080fd5b50505060405180519050925085600160a060020a031633600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8760405190815260200160405180910390a38293505b5b50505b5092915050565b610fdd33600035600160e060020a03191661113b565b1515610fe557fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a46001805474ff0000000000000000000000000000000000000000191690555b5b50505b565b600054600160a060020a031681565b61108533600035600160e060020a03191661113b565b151561108d57fe5b600381805161067d92916020019061129e565b505b5b50565b600554600160a060020a031681565b600554600090600160a060020a031663dd62ed3e8484846040516020015260405160e060020a63ffffffff8516028152600160a060020a03928316600482015291166024820152604401602060405180830381600087803b151561111857600080fd5b6102c65a03f1151561112957600080fd5b50505060405180519150505b92915050565b600030600160a060020a031683600160a060020a0316141561115f57506001610d32565b600154600160a060020a0384811691161480156111855750600054600160a060020a0316155b1561119257506001610d32565b600054600160a060020a031615156111fa577feb91385a0d70e44b915093d9ddfe6c8b41f2c56729431b1405f40aee5d874be78383604051600160a060020a039092168252600160e060020a03191660208201526040908101905180910390a1506000610d32565b60008054600160a060020a03169063b7009613908590309086906040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152600160e060020a03199091166044820152606401602060405180830381600087803b151561127357600080fd5b6102c65a03f1151561128457600080fd5b505050604051805190509050610d32565b5b5b5b92915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106112df57805160ff191683800117855561130c565b8280016001018555821561130c579182015b8281111561130c5782518255916020019190600101906112f1565b5b5061131992915061131d565b5090565b6108bc91905b808211156113195760008155600101611323565b5090565b905600a165627a7a72305820db70e09285ab9ec4ae534b31230656febf8679dc70a68bc9c30241eb0eabaa6d0029"`

// DeploySweetToken deploys a new Ethereum contract, binding an instance of SweetToken to it.
func DeploySweetToken(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ [32]byte) (common.Address, *types.Transaction, *SweetToken, error) {
	parsed, err := abi.JSON(strings.NewReader(SweetTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SweetTokenBin), backend, name_, symbol_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SweetToken{SweetTokenCaller: SweetTokenCaller{contract: contract}, SweetTokenTransactor: SweetTokenTransactor{contract: contract}}, nil
}

// SweetToken is an auto generated Go binding around an Ethereum contract.
type SweetToken struct {
	SweetTokenCaller     // Read-only binding to the contract
	SweetTokenTransactor // Write-only binding to the contract
}

// SweetTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type SweetTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SweetTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SweetTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SweetTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SweetTokenSession struct {
	Contract     *SweetToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SweetTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SweetTokenCallerSession struct {
	Contract *SweetTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SweetTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SweetTokenTransactorSession struct {
	Contract     *SweetTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SweetTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type SweetTokenRaw struct {
	Contract *SweetToken // Generic contract binding to access the raw methods on
}

// SweetTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SweetTokenCallerRaw struct {
	Contract *SweetTokenCaller // Generic read-only contract binding to access the raw methods on
}

// SweetTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SweetTokenTransactorRaw struct {
	Contract *SweetTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSweetToken creates a new instance of SweetToken, bound to a specific deployed contract.
func NewSweetToken(address common.Address, backend bind.ContractBackend) (*SweetToken, error) {
	contract, err := bindSweetToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SweetToken{SweetTokenCaller: SweetTokenCaller{contract: contract}, SweetTokenTransactor: SweetTokenTransactor{contract: contract}}, nil
}

// NewSweetTokenCaller creates a new read-only instance of SweetToken, bound to a specific deployed contract.
func NewSweetTokenCaller(address common.Address, caller bind.ContractCaller) (*SweetTokenCaller, error) {
	contract, err := bindSweetToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &SweetTokenCaller{contract: contract}, nil
}

// NewSweetTokenTransactor creates a new write-only instance of SweetToken, bound to a specific deployed contract.
func NewSweetTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*SweetTokenTransactor, error) {
	contract, err := bindSweetToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &SweetTokenTransactor{contract: contract}, nil
}

// bindSweetToken binds a generic wrapper to an already deployed contract.
func bindSweetToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SweetTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SweetToken *SweetTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SweetToken.Contract.SweetTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SweetToken *SweetTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SweetToken.Contract.SweetTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SweetToken *SweetTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SweetToken.Contract.SweetTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SweetToken *SweetTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SweetToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SweetToken *SweetTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SweetToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SweetToken *SweetTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SweetToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(_allowance uint256)
func (_SweetToken *SweetTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SweetToken.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(_allowance uint256)
func (_SweetToken *SweetTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SweetToken.Contract.Allowance(&_SweetToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(_allowance uint256)
func (_SweetToken *SweetTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SweetToken.Contract.Allowance(&_SweetToken.CallOpts, owner, spender)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_SweetToken *SweetTokenCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SweetToken.contract.Call(opts, out, "authority")
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_SweetToken *SweetTokenSession) Authority() (common.Address, error) {
	return _SweetToken.Contract.Authority(&_SweetToken.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_SweetToken *SweetTokenCallerSession) Authority() (common.Address, error) {
	return _SweetToken.Contract.Authority(&_SweetToken.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(value uint256)
func (_SweetToken *SweetTokenCaller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SweetToken.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(value uint256)
func (_SweetToken *SweetTokenSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _SweetToken.Contract.BalanceOf(&_SweetToken.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(value uint256)
func (_SweetToken *SweetTokenCallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _SweetToken.Contract.BalanceOf(&_SweetToken.CallOpts, who)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_SweetToken *SweetTokenCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SweetToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_SweetToken *SweetTokenSession) Decimals() (*big.Int, error) {
	return _SweetToken.Contract.Decimals(&_SweetToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_SweetToken *SweetTokenCallerSession) Decimals() (*big.Int, error) {
	return _SweetToken.Contract.Decimals(&_SweetToken.CallOpts)
}

// Logic is a free data retrieval call binding the contract method 0xd7dfa0dd.
//
// Solidity: function logic() constant returns(address)
func (_SweetToken *SweetTokenCaller) Logic(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SweetToken.contract.Call(opts, out, "logic")
	return *ret0, err
}

// Logic is a free data retrieval call binding the contract method 0xd7dfa0dd.
//
// Solidity: function logic() constant returns(address)
func (_SweetToken *SweetTokenSession) Logic() (common.Address, error) {
	return _SweetToken.Contract.Logic(&_SweetToken.CallOpts)
}

// Logic is a free data retrieval call binding the contract method 0xd7dfa0dd.
//
// Solidity: function logic() constant returns(address)
func (_SweetToken *SweetTokenCallerSession) Logic() (common.Address, error) {
	return _SweetToken.Contract.Logic(&_SweetToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SweetToken *SweetTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _SweetToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SweetToken *SweetTokenSession) Name() (string, error) {
	return _SweetToken.Contract.Name(&_SweetToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_SweetToken *SweetTokenCallerSession) Name() (string, error) {
	return _SweetToken.Contract.Name(&_SweetToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SweetToken *SweetTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SweetToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SweetToken *SweetTokenSession) Owner() (common.Address, error) {
	return _SweetToken.Contract.Owner(&_SweetToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SweetToken *SweetTokenCallerSession) Owner() (common.Address, error) {
	return _SweetToken.Contract.Owner(&_SweetToken.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_SweetToken *SweetTokenCaller) Stopped(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SweetToken.contract.Call(opts, out, "stopped")
	return *ret0, err
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_SweetToken *SweetTokenSession) Stopped() (bool, error) {
	return _SweetToken.Contract.Stopped(&_SweetToken.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_SweetToken *SweetTokenCallerSession) Stopped() (bool, error) {
	return _SweetToken.Contract.Stopped(&_SweetToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(bytes32)
func (_SweetToken *SweetTokenCaller) Symbol(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _SweetToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(bytes32)
func (_SweetToken *SweetTokenSession) Symbol() ([32]byte, error) {
	return _SweetToken.Contract.Symbol(&_SweetToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(bytes32)
func (_SweetToken *SweetTokenCallerSession) Symbol() ([32]byte, error) {
	return _SweetToken.Contract.Symbol(&_SweetToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SweetToken *SweetTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SweetToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SweetToken *SweetTokenSession) TotalSupply() (*big.Int, error) {
	return _SweetToken.Contract.TotalSupply(&_SweetToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SweetToken *SweetTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _SweetToken.Contract.TotalSupply(&_SweetToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(guy address, wad uint256) returns(bool)
func (_SweetToken *SweetTokenTransactor) Approve(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetToken.contract.Transact(opts, "approve", guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(guy address, wad uint256) returns(bool)
func (_SweetToken *SweetTokenSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetToken.Contract.Approve(&_SweetToken.TransactOpts, guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(guy address, wad uint256) returns(bool)
func (_SweetToken *SweetTokenTransactorSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetToken.Contract.Approve(&_SweetToken.TransactOpts, guy, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x90bc1693.
//
// Solidity: function burn(wad uint128) returns()
func (_SweetToken *SweetTokenTransactor) Burn(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _SweetToken.contract.Transact(opts, "burn", wad)
}

// Burn is a paid mutator transaction binding the contract method 0x90bc1693.
//
// Solidity: function burn(wad uint128) returns()
func (_SweetToken *SweetTokenSession) Burn(wad *big.Int) (*types.Transaction, error) {
	return _SweetToken.Contract.Burn(&_SweetToken.TransactOpts, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x90bc1693.
//
// Solidity: function burn(wad uint128) returns()
func (_SweetToken *SweetTokenTransactorSession) Burn(wad *big.Int) (*types.Transaction, error) {
	return _SweetToken.Contract.Burn(&_SweetToken.TransactOpts, wad)
}

// Mint is a paid mutator transaction binding the contract method 0x69d3e20e.
//
// Solidity: function mint(wad uint128) returns()
func (_SweetToken *SweetTokenTransactor) Mint(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _SweetToken.contract.Transact(opts, "mint", wad)
}

// Mint is a paid mutator transaction binding the contract method 0x69d3e20e.
//
// Solidity: function mint(wad uint128) returns()
func (_SweetToken *SweetTokenSession) Mint(wad *big.Int) (*types.Transaction, error) {
	return _SweetToken.Contract.Mint(&_SweetToken.TransactOpts, wad)
}

// Mint is a paid mutator transaction binding the contract method 0x69d3e20e.
//
// Solidity: function mint(wad uint128) returns()
func (_SweetToken *SweetTokenTransactorSession) Mint(wad *big.Int) (*types.Transaction, error) {
	return _SweetToken.Contract.Mint(&_SweetToken.TransactOpts, wad)
}

// Pull is a paid mutator transaction binding the contract method 0x8402181f.
//
// Solidity: function pull(src address, wad uint128) returns(bool)
func (_SweetToken *SweetTokenTransactor) Pull(opts *bind.TransactOpts, src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetToken.contract.Transact(opts, "pull", src, wad)
}

// Pull is a paid mutator transaction binding the contract method 0x8402181f.
//
// Solidity: function pull(src address, wad uint128) returns(bool)
func (_SweetToken *SweetTokenSession) Pull(src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetToken.Contract.Pull(&_SweetToken.TransactOpts, src, wad)
}

// Pull is a paid mutator transaction binding the contract method 0x8402181f.
//
// Solidity: function pull(src address, wad uint128) returns(bool)
func (_SweetToken *SweetTokenTransactorSession) Pull(src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetToken.Contract.Pull(&_SweetToken.TransactOpts, src, wad)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_SweetToken *SweetTokenTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _SweetToken.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_SweetToken *SweetTokenSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _SweetToken.Contract.SetAuthority(&_SweetToken.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_SweetToken *SweetTokenTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _SweetToken.Contract.SetAuthority(&_SweetToken.TransactOpts, authority_)
}

// SetLogic is a paid mutator transaction binding the contract method 0x71857000.
//
// Solidity: function setLogic(logic_ address) returns(bool)
func (_SweetToken *SweetTokenTransactor) SetLogic(opts *bind.TransactOpts, logic_ common.Address) (*types.Transaction, error) {
	return _SweetToken.contract.Transact(opts, "setLogic", logic_)
}

// SetLogic is a paid mutator transaction binding the contract method 0x71857000.
//
// Solidity: function setLogic(logic_ address) returns(bool)
func (_SweetToken *SweetTokenSession) SetLogic(logic_ common.Address) (*types.Transaction, error) {
	return _SweetToken.Contract.SetLogic(&_SweetToken.TransactOpts, logic_)
}

// SetLogic is a paid mutator transaction binding the contract method 0x71857000.
//
// Solidity: function setLogic(logic_ address) returns(bool)
func (_SweetToken *SweetTokenTransactorSession) SetLogic(logic_ common.Address) (*types.Transaction, error) {
	return _SweetToken.Contract.SetLogic(&_SweetToken.TransactOpts, logic_)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(name_ string) returns()
func (_SweetToken *SweetTokenTransactor) SetName(opts *bind.TransactOpts, name_ string) (*types.Transaction, error) {
	return _SweetToken.contract.Transact(opts, "setName", name_)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(name_ string) returns()
func (_SweetToken *SweetTokenSession) SetName(name_ string) (*types.Transaction, error) {
	return _SweetToken.Contract.SetName(&_SweetToken.TransactOpts, name_)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(name_ string) returns()
func (_SweetToken *SweetTokenTransactorSession) SetName(name_ string) (*types.Transaction, error) {
	return _SweetToken.Contract.SetName(&_SweetToken.TransactOpts, name_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_SweetToken *SweetTokenTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _SweetToken.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_SweetToken *SweetTokenSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _SweetToken.Contract.SetOwner(&_SweetToken.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_SweetToken *SweetTokenTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _SweetToken.Contract.SetOwner(&_SweetToken.TransactOpts, owner_)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_SweetToken *SweetTokenTransactor) Start(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SweetToken.contract.Transact(opts, "start")
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_SweetToken *SweetTokenSession) Start() (*types.Transaction, error) {
	return _SweetToken.Contract.Start(&_SweetToken.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_SweetToken *SweetTokenTransactorSession) Start() (*types.Transaction, error) {
	return _SweetToken.Contract.Start(&_SweetToken.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_SweetToken *SweetTokenTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SweetToken.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_SweetToken *SweetTokenSession) Stop() (*types.Transaction, error) {
	return _SweetToken.Contract.Stop(&_SweetToken.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_SweetToken *SweetTokenTransactorSession) Stop() (*types.Transaction, error) {
	return _SweetToken.Contract.Stop(&_SweetToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(dst address, wad uint256) returns(bool)
func (_SweetToken *SweetTokenTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetToken.contract.Transact(opts, "transfer", dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(dst address, wad uint256) returns(bool)
func (_SweetToken *SweetTokenSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetToken.Contract.Transfer(&_SweetToken.TransactOpts, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(dst address, wad uint256) returns(bool)
func (_SweetToken *SweetTokenTransactorSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetToken.Contract.Transfer(&_SweetToken.TransactOpts, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_SweetToken *SweetTokenTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetToken.contract.Transact(opts, "transferFrom", src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_SweetToken *SweetTokenSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetToken.Contract.TransferFrom(&_SweetToken.TransactOpts, src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_SweetToken *SweetTokenTransactorSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetToken.Contract.TransferFrom(&_SweetToken.TransactOpts, src, dst, wad)
}

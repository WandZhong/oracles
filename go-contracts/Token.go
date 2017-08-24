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

// TokenABI is the input ABI used to generate the binding from.
const TokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guy\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"push\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"logic_\",\"type\":\"address\"}],\"name\":\"setLogic\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"pull\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name_\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"logic\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"_allowance\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"name_\",\"type\":\"string\"},{\"name\":\"symbol_\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"foo\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"bar\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fax\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"UnauthorizedAccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// TokenBin is the compiled bytecode used for deploying new contracts.
const TokenBin = `"0x6060604052601260045534156200001557600080fd5b604051620015c5380380620015c5833981016040528080518201919060200180519150505b5b60018054600160a060020a03191633600160a060020a03169081179091557fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a25b60038280516200009b929160200190620000aa565b5060028190555b505062000154565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620000ed57805160ff19168380011785556200011d565b828001600101855582156200011d579182015b828111156200011d57825182559160200191906001019062000100565b5b506200012c92915062000130565b5090565b6200015191905b808211156200012c576000815560010162000137565b5090565b90565b61146180620001646000396000f300606060405236156101175763ffffffff60e060020a60003504166306fdde03811461011c57806307da68f5146101a7578063095ea7b3146101bc57806313af4035146101f257806318160ddd1461021357806323b872dd14610238578063313ce567146102745780633452f51d1461029957806369d3e20e146102d857806370a08231146102f9578063718570001461032a57806375f12b211461035d5780637a9e5e4b146103845780638402181f146103a55780638da5cb5b146103e457806390bc16931461041357806395d89b4114610434578063a9059cbb14610459578063be9a65551461048f578063bf7e214f146104a4578063c47f0027146104d3578063d7dfa0dd14610526578063dd62ed3e14610555575b600080fd5b341561012757600080fd5b61012f61058c565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561016c5780820151818401525b602001610153565b50505050905090810190601f1680156101995780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156101b257600080fd5b6101ba61062a565b005b34156101c757600080fd5b6101de600160a060020a03600435166024356106c9565b604051901515815260200160405180910390f35b34156101fd57600080fd5b6101ba600160a060020a036004351661081d565b005b341561021e57600080fd5b61022661089b565b60405190815260200160405180910390f35b341561024357600080fd5b6101de600160a060020a0360043581169060243516604435610905565b604051901515815260200160405180910390f35b341561027f57600080fd5b610226610a54565b60405190815260200160405180910390f35b34156102a457600080fd5b6101de600160a060020a03600435166001608060020a0360243516610a5a565b604051901515815260200160405180910390f35b34156102e357600080fd5b6101ba6001608060020a0360043516610a78565b005b341561030457600080fd5b610226600160a060020a0360043516610c81565b60405190815260200160405180910390f35b341561033557600080fd5b6101de600160a060020a0360043516610cfe565b604051901515815260200160405180910390f35b341561036857600080fd5b6101de610daa565b604051901515815260200160405180910390f35b341561038f57600080fd5b6101ba600160a060020a0360043516610dba565b005b34156103b057600080fd5b6101de600160a060020a03600435166001608060020a0360243516610e38565b604051901515815260200160405180910390f35b34156103ef57600080fd5b6103f7610e57565b604051600160a060020a03909116815260200160405180910390f35b341561041e57600080fd5b6101ba6001608060020a0360043516610e66565b005b341561043f57600080fd5b610226610f6a565b60405190815260200160405180910390f35b341561046457600080fd5b6101de600160a060020a0360043516602435610f70565b604051901515815260200160405180910390f35b341561049a57600080fd5b6101ba6110be565b005b34156104af57600080fd5b6103f7611157565b604051600160a060020a03909116815260200160405180910390f35b34156104de57600080fd5b6101ba60046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061116695505050505050565b005b341561053157600080fd5b6103f761119d565b604051600160a060020a03909116815260200160405180910390f35b341561056057600080fd5b610226600160a060020a03600435811690602435166111ac565b60405190815260200160405180910390f35b60038054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106225780601f106105f757610100808354040283529160200191610622565b820191906000526020600020905b81548152906001019060200180831161060557829003601f168201915b505050505081565b61064033600035600160e060020a031916611232565b151561064857fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a46001805474ff0000000000000000000000000000000000000000191660a060020a1790555b5b50505b565b600154600090819060a060020a900460ff16156106e257fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a4600554600160a060020a031663e1f21c6733888860006040516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b15156107a457600080fd5b6102c65a03f115156107b557600080fd5b5050506040518051935050821561080e5785600160a060020a031633600160a060020a03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258760405190815260200160405180910390a35b8293505b5b50505b5092915050565b61083333600035600160e060020a031916611232565b151561083b57fe5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a25b5b50565b600554600090600160a060020a03166318160ddd82604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156108e557600080fd5b6102c65a03f115156108f657600080fd5b50505060405180519150505b90565b600154600090819060a060020a900460ff161561091e57fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a4600554600160a060020a03166323b872dd88888860006040516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b15156109e057600080fd5b6102c65a03f115156109f157600080fd5b50505060405180519050925085600160a060020a031687600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8760405190815260200160405180910390a38293505b5b50505b509392505050565b60045481565b6000610a6f83836001608060020a0316610f70565b90505b92915050565b610a8e33600035600160e060020a031916611232565b1515610a9657fe5b60015460a060020a900460ff1615610aaa57fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a4600554600160a060020a03166369d3e20e8460405160e060020a63ffffffff84160281526001608060020a039091166004820152602401600060405180830381600087803b1515610b5057600080fd5b6102c65a03f11515610b6157600080fd5b5050600554600160a060020a031690506373d4a13a6000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610bad57600080fd5b6102c65a03f11515610bbe57600080fd5b50505060405180519050600160a060020a0316638da5cb5b6000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610c0d57600080fd5b6102c65a03f11515610c1e57600080fd5b50505060405180519050600160a060020a031630600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef856040516001608060020a03909116815260200160405180910390a35b5b50505b5b50565b600554600090600160a060020a03166370a0823183836040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b1515610cdc57600080fd5b6102c65a03f11515610ced57600080fd5b50505060405180519150505b919050565b6000610d1633600035600160e060020a031916611232565b1515610d1e57fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a46005805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038616179055600192505b5b50505b919050565b60015460a060020a900460ff1681565b610dd033600035600160e060020a031916611232565b1515610dd857fe5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167f1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada460405160405180910390a25b5b50565b6000610a6f8333846001608060020a0316610905565b90505b92915050565b600154600160a060020a031681565b610e7c33600035600160e060020a031916611232565b1515610e8457fe5b60015460a060020a900460ff1615610e9857fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a4600554600160a060020a0316637261e469338560405160e060020a63ffffffff8516028152600160a060020a0390921660048301526001608060020a03166024820152604401600060405180830381600087803b1515610f4d57600080fd5b6102c65a03f11515610f5e57600080fd5b5050505b5b50505b5b50565b60025481565b600154600090819060a060020a900460ff1615610f8957fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a4600554600160a060020a031663beabacc833888860006040516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b151561104b57600080fd5b6102c65a03f1151561105c57600080fd5b50505060405180519050925085600160a060020a031633600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8760405190815260200160405180910390a38293505b5b50505b5092915050565b6110d433600035600160e060020a031916611232565b15156110dc57fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a46001805474ff0000000000000000000000000000000000000000191690555b5b50505b565b600054600160a060020a031681565b61117c33600035600160e060020a031916611232565b151561118457fe5b60038180516106c3929160200190611395565b505b5b50565b600554600160a060020a031681565b600554600090600160a060020a031663dd62ed3e8484846040516020015260405160e060020a63ffffffff8516028152600160a060020a03928316600482015291166024820152604401602060405180830381600087803b151561120f57600080fd5b6102c65a03f1151561122057600080fd5b50505060405180519150505b92915050565b600030600160a060020a031683600160a060020a0316141561125657506001610a72565b600154600160a060020a03848116911614801561127c5750600054600160a060020a0316155b1561128957506001610a72565b600054600160a060020a031615156112f1577feb91385a0d70e44b915093d9ddfe6c8b41f2c56729431b1405f40aee5d874be78383604051600160a060020a039092168252600160e060020a03191660208201526040908101905180910390a1506000610a72565b60008054600160a060020a03169063b7009613908590309086906040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152600160e060020a03199091166044820152606401602060405180830381600087803b151561136a57600080fd5b6102c65a03f1151561137b57600080fd5b505050604051805190509050610a72565b5b5b5b92915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106113d657805160ff1916838001178555611403565b82800160010185558215611403579182015b828111156114035782518255916020019190600101906113e8565b5b50611410929150611414565b5090565b61090291905b80821115611410576000815560010161141a565b5090565b905600a165627a7a723058208517c606fccdeebff0265c1d468698d7b53f64e7abafac4b335a262702dabdb60029"`

// DeployToken deploys a new Ethereum contract, binding an instance of Token to it.
func DeployToken(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ [32]byte) (common.Address, *types.Transaction, *Token, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenBin), backend, name_, symbol_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}}, nil
}

// Token is an auto generated Go binding around an Ethereum contract.
type Token struct {
	TokenCaller     // Read-only binding to the contract
	TokenTransactor // Write-only binding to the contract
}

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenSession struct {
	Contract     *Token            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenCallerSession struct {
	Contract *TokenCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// TokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenTransactorSession struct {
	Contract     *TokenTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenRaw struct {
	Contract *Token // Generic contract binding to access the raw methods on
}

// TokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenCallerRaw struct {
	Contract *TokenCaller // Generic read-only contract binding to access the raw methods on
}

// TokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenTransactorRaw struct {
	Contract *TokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewToken creates a new instance of Token, bound to a specific deployed contract.
func NewToken(address common.Address, backend bind.ContractBackend) (*Token, error) {
	contract, err := bindToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Token{TokenCaller: TokenCaller{contract: contract}, TokenTransactor: TokenTransactor{contract: contract}}, nil
}

// NewTokenCaller creates a new read-only instance of Token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{contract: contract}, nil
}

// NewTokenTransactor creates a new write-only instance of Token, bound to a specific deployed contract.
func NewTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenTransactor, error) {
	contract, err := bindToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &TokenTransactor{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.TokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.TokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Token *TokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Token.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Token *TokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Token *TokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Token.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(_allowance uint256)
func (_Token *TokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(_allowance uint256)
func (_Token *TokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(_allowance uint256)
func (_Token *TokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Token.Contract.Allowance(&_Token.CallOpts, owner, spender)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Token *TokenCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "authority")
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Token *TokenSession) Authority() (common.Address, error) {
	return _Token.Contract.Authority(&_Token.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Token *TokenCallerSession) Authority() (common.Address, error) {
	return _Token.Contract.Authority(&_Token.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(value uint256)
func (_Token *TokenCaller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(value uint256)
func (_Token *TokenSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(value uint256)
func (_Token *TokenCallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _Token.Contract.BalanceOf(&_Token.CallOpts, who)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_Token *TokenCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_Token *TokenSession) Decimals() (*big.Int, error) {
	return _Token.Contract.Decimals(&_Token.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_Token *TokenCallerSession) Decimals() (*big.Int, error) {
	return _Token.Contract.Decimals(&_Token.CallOpts)
}

// Logic is a free data retrieval call binding the contract method 0xd7dfa0dd.
//
// Solidity: function logic() constant returns(address)
func (_Token *TokenCaller) Logic(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "logic")
	return *ret0, err
}

// Logic is a free data retrieval call binding the contract method 0xd7dfa0dd.
//
// Solidity: function logic() constant returns(address)
func (_Token *TokenSession) Logic() (common.Address, error) {
	return _Token.Contract.Logic(&_Token.CallOpts)
}

// Logic is a free data retrieval call binding the contract method 0xd7dfa0dd.
//
// Solidity: function logic() constant returns(address)
func (_Token *TokenCallerSession) Logic() (common.Address, error) {
	return _Token.Contract.Logic(&_Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Token *TokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Token *TokenSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Token *TokenCallerSession) Name() (string, error) {
	return _Token.Contract.Name(&_Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Token *TokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Token *TokenSession) Owner() (common.Address, error) {
	return _Token.Contract.Owner(&_Token.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Token *TokenCallerSession) Owner() (common.Address, error) {
	return _Token.Contract.Owner(&_Token.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_Token *TokenCaller) Stopped(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "stopped")
	return *ret0, err
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_Token *TokenSession) Stopped() (bool, error) {
	return _Token.Contract.Stopped(&_Token.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_Token *TokenCallerSession) Stopped() (bool, error) {
	return _Token.Contract.Stopped(&_Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(bytes32)
func (_Token *TokenCaller) Symbol(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(bytes32)
func (_Token *TokenSession) Symbol() ([32]byte, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(bytes32)
func (_Token *TokenCallerSession) Symbol() ([32]byte, error) {
	return _Token.Contract.Symbol(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Token.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Token *TokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Token.Contract.TotalSupply(&_Token.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(guy address, wad uint256) returns(bool)
func (_Token *TokenTransactor) Approve(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "approve", guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(guy address, wad uint256) returns(bool)
func (_Token *TokenSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(guy address, wad uint256) returns(bool)
func (_Token *TokenTransactorSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Approve(&_Token.TransactOpts, guy, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x90bc1693.
//
// Solidity: function burn(wad uint128) returns()
func (_Token *TokenTransactor) Burn(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "burn", wad)
}

// Burn is a paid mutator transaction binding the contract method 0x90bc1693.
//
// Solidity: function burn(wad uint128) returns()
func (_Token *TokenSession) Burn(wad *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Burn(&_Token.TransactOpts, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x90bc1693.
//
// Solidity: function burn(wad uint128) returns()
func (_Token *TokenTransactorSession) Burn(wad *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Burn(&_Token.TransactOpts, wad)
}

// Mint is a paid mutator transaction binding the contract method 0x69d3e20e.
//
// Solidity: function mint(wad uint128) returns()
func (_Token *TokenTransactor) Mint(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "mint", wad)
}

// Mint is a paid mutator transaction binding the contract method 0x69d3e20e.
//
// Solidity: function mint(wad uint128) returns()
func (_Token *TokenSession) Mint(wad *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Mint(&_Token.TransactOpts, wad)
}

// Mint is a paid mutator transaction binding the contract method 0x69d3e20e.
//
// Solidity: function mint(wad uint128) returns()
func (_Token *TokenTransactorSession) Mint(wad *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Mint(&_Token.TransactOpts, wad)
}

// Pull is a paid mutator transaction binding the contract method 0x8402181f.
//
// Solidity: function pull(src address, wad uint128) returns(bool)
func (_Token *TokenTransactor) Pull(opts *bind.TransactOpts, src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "pull", src, wad)
}

// Pull is a paid mutator transaction binding the contract method 0x8402181f.
//
// Solidity: function pull(src address, wad uint128) returns(bool)
func (_Token *TokenSession) Pull(src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Pull(&_Token.TransactOpts, src, wad)
}

// Pull is a paid mutator transaction binding the contract method 0x8402181f.
//
// Solidity: function pull(src address, wad uint128) returns(bool)
func (_Token *TokenTransactorSession) Pull(src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Pull(&_Token.TransactOpts, src, wad)
}

// Push is a paid mutator transaction binding the contract method 0x3452f51d.
//
// Solidity: function push(dst address, wad uint128) returns(bool)
func (_Token *TokenTransactor) Push(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "push", dst, wad)
}

// Push is a paid mutator transaction binding the contract method 0x3452f51d.
//
// Solidity: function push(dst address, wad uint128) returns(bool)
func (_Token *TokenSession) Push(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Push(&_Token.TransactOpts, dst, wad)
}

// Push is a paid mutator transaction binding the contract method 0x3452f51d.
//
// Solidity: function push(dst address, wad uint128) returns(bool)
func (_Token *TokenTransactorSession) Push(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Push(&_Token.TransactOpts, dst, wad)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_Token *TokenTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_Token *TokenSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetAuthority(&_Token.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_Token *TokenTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetAuthority(&_Token.TransactOpts, authority_)
}

// SetLogic is a paid mutator transaction binding the contract method 0x71857000.
//
// Solidity: function setLogic(logic_ address) returns(bool)
func (_Token *TokenTransactor) SetLogic(opts *bind.TransactOpts, logic_ common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setLogic", logic_)
}

// SetLogic is a paid mutator transaction binding the contract method 0x71857000.
//
// Solidity: function setLogic(logic_ address) returns(bool)
func (_Token *TokenSession) SetLogic(logic_ common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetLogic(&_Token.TransactOpts, logic_)
}

// SetLogic is a paid mutator transaction binding the contract method 0x71857000.
//
// Solidity: function setLogic(logic_ address) returns(bool)
func (_Token *TokenTransactorSession) SetLogic(logic_ common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetLogic(&_Token.TransactOpts, logic_)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(name_ string) returns()
func (_Token *TokenTransactor) SetName(opts *bind.TransactOpts, name_ string) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setName", name_)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(name_ string) returns()
func (_Token *TokenSession) SetName(name_ string) (*types.Transaction, error) {
	return _Token.Contract.SetName(&_Token.TransactOpts, name_)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(name_ string) returns()
func (_Token *TokenTransactorSession) SetName(name_ string) (*types.Transaction, error) {
	return _Token.Contract.SetName(&_Token.TransactOpts, name_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_Token *TokenTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_Token *TokenSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetOwner(&_Token.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_Token *TokenTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Token.Contract.SetOwner(&_Token.TransactOpts, owner_)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Token *TokenTransactor) Start(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "start")
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Token *TokenSession) Start() (*types.Transaction, error) {
	return _Token.Contract.Start(&_Token.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Token *TokenTransactorSession) Start() (*types.Transaction, error) {
	return _Token.Contract.Start(&_Token.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Token *TokenTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Token *TokenSession) Stop() (*types.Transaction, error) {
	return _Token.Contract.Stop(&_Token.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Token *TokenTransactorSession) Stop() (*types.Transaction, error) {
	return _Token.Contract.Stop(&_Token.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(dst address, wad uint256) returns(bool)
func (_Token *TokenTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transfer", dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(dst address, wad uint256) returns(bool)
func (_Token *TokenSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(dst address, wad uint256) returns(bool)
func (_Token *TokenTransactorSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Token.Contract.Transfer(&_Token.TransactOpts, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_Token *TokenTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Token.contract.Transact(opts, "transferFrom", src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_Token *TokenSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_Token *TokenTransactorSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Token.Contract.TransferFrom(&_Token.TransactOpts, src, dst, wad)
}

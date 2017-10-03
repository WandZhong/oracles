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

// BridgeTokenABI is the input ABI used to generate the binding from.
const BridgeTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guy\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"brgAmount\",\"type\":\"uint128\"},{\"name\":\"vault\",\"type\":\"address\"},{\"name\":\"uouIndex\",\"type\":\"uint256\"}],\"name\":\"repayUou\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"mintFor\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"logic_\",\"type\":\"address\"}],\"name\":\"setLogic\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"pull\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name_\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"logic\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"_allowance\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"name_\",\"type\":\"string\"},{\"name\":\"symbol_\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"src\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"LogBurn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"src\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"LogMint\",\"type\":\"event\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"foo\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"bar\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fax\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"UnauthorizedAccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// BridgeTokenBin is the compiled bytecode used for deploying new contracts.
const BridgeTokenBin = `"0x6060604052601260045534156200001557600080fd5b6040516200160038038062001600833981016040528080518201919060200180519150505b81815b5b60018054600160a060020a03191633600160a060020a03169081179091557fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a25b60038280516200009e929160200190620000b0565b5060028190555b50505b50506200015a565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620000f357805160ff191683800117855562000123565b8280016001018555821562000123579182015b828111156200012357825182559160200191906001019062000106565b5b506200013292915062000136565b5090565b6200015791905b808211156200013257600081556001016200013d565b5090565b90565b611496806200016a6000396000f300606060405236156101175763ffffffff60e060020a60003504166306fdde03811461011c57806307da68f5146101a7578063095ea7b3146101bc5780630badbc55146101f257806313af40351461022257806318160ddd1461024357806323b872dd1461026857806327cfc219146102a4578063313ce567146102d157806370a08231146102f6578063718570001461032757806375f12b211461035a5780637a9e5e4b146103815780638402181f146103a25780638da5cb5b146103e157806390bc16931461041057806395d89b4114610431578063a9059cbb14610456578063be9a65551461048c578063bf7e214f146104a1578063c47f0027146104d0578063d7dfa0dd14610523578063dd62ed3e14610552575b600080fd5b341561012757600080fd5b61012f610589565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561016c5780820151818401525b602001610153565b50505050905090810190601f1680156101995780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156101b257600080fd5b6101ba610627565b005b34156101c757600080fd5b6101de600160a060020a03600435166024356106c6565b604051901515815260200160405180910390f35b34156101fd57600080fd5b6101ba6001608060020a0360043516600160a060020a036024351660443561081a565b005b341561022d57600080fd5b6101ba600160a060020a0360043516610980565b005b341561024e57600080fd5b6102566109fe565b60405190815260200160405180910390f35b341561027357600080fd5b6101de600160a060020a0360043581169060243516604435610a68565b604051901515815260200160405180910390f35b34156102af57600080fd5b6101ba600160a060020a03600435166001608060020a0360243516610bb7565b005b34156102dc57600080fd5b610256610c76565b60405190815260200160405180910390f35b341561030157600080fd5b610256600160a060020a0360043516610c7c565b60405190815260200160405180910390f35b341561033257600080fd5b6101de600160a060020a0360043516610cf9565b604051901515815260200160405180910390f35b341561036557600080fd5b6101de610db7565b604051901515815260200160405180910390f35b341561038c57600080fd5b6101ba600160a060020a0360043516610dc7565b005b34156103ad57600080fd5b6101de600160a060020a03600435166001608060020a0360243516610e45565b604051901515815260200160405180910390f35b34156103ec57600080fd5b6103f4610e64565b604051600160a060020a03909116815260200160405180910390f35b341561041b57600080fd5b6101ba6001608060020a0360043516610e73565b005b341561043c57600080fd5b610256610f9f565b60405190815260200160405180910390f35b341561046157600080fd5b6101de600160a060020a0360043516602435610fa5565b604051901515815260200160405180910390f35b341561049757600080fd5b6101ba6110f3565b005b34156104ac57600080fd5b6103f461118c565b604051600160a060020a03909116815260200160405180910390f35b34156104db57600080fd5b6101ba60046024813581810190830135806020601f8201819004810201604051908101604052818152929190602084018383808284375094965061119b95505050505050565b005b341561052e57600080fd5b6103f46111d2565b604051600160a060020a03909116815260200160405180910390f35b341561055d57600080fd5b610256600160a060020a03600435811690602435166111e1565b60405190815260200160405180910390f35b60038054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561061f5780601f106105f45761010080835404028352916020019161061f565b820191906000526020600020905b81548152906001019060200180831161060257829003601f168201915b505050505081565b61063d33600035600160e060020a031916611267565b151561064557fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a46001805474ff0000000000000000000000000000000000000000191660a060020a1790555b5b50505b565b600154600090819060a060020a900460ff16156106df57fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a4600554600160a060020a031663e1f21c6733888860006040516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b15156107a157600080fd5b6102c65a03f115156107b257600080fd5b5050506040518051935050821561080b5785600160a060020a031633600160a060020a03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258760405190815260200160405180910390a35b8293505b5b50505b5092915050565b6000836001608060020a031661082f33610c7c565b101561083a57600080fd5b82600160a060020a0316631a93ebf0858460006040516020015260405160e060020a63ffffffff85160281526001608060020a0390921660048301526024820152604401602060405180830381600087803b151561089757600080fd5b6102c65a03f115156108a857600080fd5b5050506040518051600554909250600160a060020a03169050637261e469338360405160e060020a63ffffffff8516028152600160a060020a0390921660048301526001608060020a03166024820152604401600060405180830381600087803b151561091457600080fd5b6102c65a03f1151561092557600080fd5b50505082600160a060020a031633600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516001608060020a03909116815260200160405180910390a35b50505050565b61099633600035600160e060020a031916611267565b151561099e57fe5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a25b5b50565b600554600090600160a060020a03166318160ddd82604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610a4857600080fd5b6102c65a03f11515610a5957600080fd5b50505060405180519150505b90565b600154600090819060a060020a900460ff1615610a8157fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a4600554600160a060020a03166323b872dd88888860006040516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b1515610b4357600080fd5b6102c65a03f11515610b5457600080fd5b50505060405180519050925085600160a060020a031687600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8760405190815260200160405180910390a38293505b5b50505b509392505050565b600554600160a060020a03166327cfc219838360405160e060020a63ffffffff8516028152600160a060020a0390921660048301526001608060020a03166024820152604401600060405180830381600087803b1515610c1657600080fd5b6102c65a03f11515610c2757600080fd5b50505081600160a060020a03167f9f494565851dbcb31fb5198ca217cda6833282fadb96ba9431bd19c82afc1dd3826040516001608060020a03909116815260200160405180910390a25b5050565b60045481565b600554600090600160a060020a03166370a0823183836040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b1515610cd757600080fd5b6102c65a03f11515610ce857600080fd5b50505060405180519150505b919050565b6000610d1133600035600160e060020a031916611267565b1515610d1957fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a4600160a060020a0384161515610d8157fe5b6005805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038616179055600192505b5b50505b919050565b60015460a060020a900460ff1681565b610ddd33600035600160e060020a031916611267565b1515610de557fe5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167f1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada460405160405180910390a25b5b50565b6000610e5b8333846001608060020a0316610a68565b90505b92915050565b600154600160a060020a031681565b60015460a060020a900460ff1615610e8757fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a4600554600160a060020a0316637261e469338560405160e060020a63ffffffff8516028152600160a060020a0390921660048301526001608060020a03166024820152604401600060405180830381600087803b1515610f3c57600080fd5b6102c65a03f11515610f4d57600080fd5b50505033600160a060020a03167f38d762ef507761291a578e921acfe29c1af31a7331ea03e391cf16cfc4d4f581846040516001608060020a03909116815260200160405180910390a25b5b50505b50565b60025481565b600154600090819060a060020a900460ff1615610fbe57fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a4600554600160a060020a031663beabacc833888860006040516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b151561108057600080fd5b6102c65a03f1151561109157600080fd5b50505060405180519050925085600160a060020a031633600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8760405190815260200160405180910390a38293505b5b50505b5092915050565b61110933600035600160e060020a031916611267565b151561111157fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a46001805474ff0000000000000000000000000000000000000000191690555b5b50505b565b600054600160a060020a031681565b6111b133600035600160e060020a031916611267565b15156111b957fe5b60038180516106c09291602001906113ca565b505b5b50565b600554600160a060020a031681565b600554600090600160a060020a031663dd62ed3e8484846040516020015260405160e060020a63ffffffff8516028152600160a060020a03928316600482015291166024820152604401602060405180830381600087803b151561124457600080fd5b6102c65a03f1151561125557600080fd5b50505060405180519150505b92915050565b600030600160a060020a031683600160a060020a0316141561128b57506001610e5e565b600154600160a060020a0384811691161480156112b15750600054600160a060020a0316155b156112be57506001610e5e565b600054600160a060020a03161515611326577feb91385a0d70e44b915093d9ddfe6c8b41f2c56729431b1405f40aee5d874be78383604051600160a060020a039092168252600160e060020a03191660208201526040908101905180910390a1506000610e5e565b60008054600160a060020a03169063b7009613908590309086906040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152600160e060020a03199091166044820152606401602060405180830381600087803b151561139f57600080fd5b6102c65a03f115156113b057600080fd5b505050604051805190509050610e5e565b5b5b5b92915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061140b57805160ff1916838001178555611438565b82800160010185558215611438579182015b8281111561143857825182559160200191906001019061141d565b5b50611445929150611449565b5090565b610a6591905b80821115611445576000815560010161144f565b5090565b905600a165627a7a723058200dfbe594fd3c677aae29cc0be9dd16a2b9fda419a0040c43ecf205c08f6bfde60029"`

// DeployBridgeToken deploys a new Ethereum contract, binding an instance of BridgeToken to it.
func DeployBridgeToken(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ [32]byte) (common.Address, *types.Transaction, *BridgeToken, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BridgeTokenBin), backend, name_, symbol_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeToken{BridgeTokenCaller: BridgeTokenCaller{contract: contract}, BridgeTokenTransactor: BridgeTokenTransactor{contract: contract}}, nil
}

// BridgeToken is an auto generated Go binding around an Ethereum contract.
type BridgeToken struct {
	BridgeTokenCaller     // Read-only binding to the contract
	BridgeTokenTransactor // Write-only binding to the contract
}

// BridgeTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeTokenSession struct {
	Contract     *BridgeToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeTokenCallerSession struct {
	Contract *BridgeTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BridgeTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTokenTransactorSession struct {
	Contract     *BridgeTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BridgeTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeTokenRaw struct {
	Contract *BridgeToken // Generic contract binding to access the raw methods on
}

// BridgeTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeTokenCallerRaw struct {
	Contract *BridgeTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTokenTransactorRaw struct {
	Contract *BridgeTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeToken creates a new instance of BridgeToken, bound to a specific deployed contract.
func NewBridgeToken(address common.Address, backend bind.ContractBackend) (*BridgeToken, error) {
	contract, err := bindBridgeToken(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeToken{BridgeTokenCaller: BridgeTokenCaller{contract: contract}, BridgeTokenTransactor: BridgeTokenTransactor{contract: contract}}, nil
}

// NewBridgeTokenCaller creates a new read-only instance of BridgeToken, bound to a specific deployed contract.
func NewBridgeTokenCaller(address common.Address, caller bind.ContractCaller) (*BridgeTokenCaller, error) {
	contract, err := bindBridgeToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTokenCaller{contract: contract}, nil
}

// NewBridgeTokenTransactor creates a new write-only instance of BridgeToken, bound to a specific deployed contract.
func NewBridgeTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTokenTransactor, error) {
	contract, err := bindBridgeToken(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &BridgeTokenTransactor{contract: contract}, nil
}

// bindBridgeToken binds a generic wrapper to an already deployed contract.
func bindBridgeToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeToken *BridgeTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BridgeToken.Contract.BridgeTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeToken *BridgeTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeToken.Contract.BridgeTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeToken *BridgeTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeToken.Contract.BridgeTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeToken *BridgeTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BridgeToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeToken *BridgeTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeToken *BridgeTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(_allowance uint256)
func (_BridgeToken *BridgeTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BridgeToken.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(_allowance uint256)
func (_BridgeToken *BridgeTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BridgeToken.Contract.Allowance(&_BridgeToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(_allowance uint256)
func (_BridgeToken *BridgeTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _BridgeToken.Contract.Allowance(&_BridgeToken.CallOpts, owner, spender)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_BridgeToken *BridgeTokenCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeToken.contract.Call(opts, out, "authority")
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_BridgeToken *BridgeTokenSession) Authority() (common.Address, error) {
	return _BridgeToken.Contract.Authority(&_BridgeToken.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_BridgeToken *BridgeTokenCallerSession) Authority() (common.Address, error) {
	return _BridgeToken.Contract.Authority(&_BridgeToken.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(value uint256)
func (_BridgeToken *BridgeTokenCaller) BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BridgeToken.contract.Call(opts, out, "balanceOf", who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(value uint256)
func (_BridgeToken *BridgeTokenSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _BridgeToken.Contract.BalanceOf(&_BridgeToken.CallOpts, who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(who address) constant returns(value uint256)
func (_BridgeToken *BridgeTokenCallerSession) BalanceOf(who common.Address) (*big.Int, error) {
	return _BridgeToken.Contract.BalanceOf(&_BridgeToken.CallOpts, who)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_BridgeToken *BridgeTokenCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BridgeToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_BridgeToken *BridgeTokenSession) Decimals() (*big.Int, error) {
	return _BridgeToken.Contract.Decimals(&_BridgeToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint256)
func (_BridgeToken *BridgeTokenCallerSession) Decimals() (*big.Int, error) {
	return _BridgeToken.Contract.Decimals(&_BridgeToken.CallOpts)
}

// Logic is a free data retrieval call binding the contract method 0xd7dfa0dd.
//
// Solidity: function logic() constant returns(address)
func (_BridgeToken *BridgeTokenCaller) Logic(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeToken.contract.Call(opts, out, "logic")
	return *ret0, err
}

// Logic is a free data retrieval call binding the contract method 0xd7dfa0dd.
//
// Solidity: function logic() constant returns(address)
func (_BridgeToken *BridgeTokenSession) Logic() (common.Address, error) {
	return _BridgeToken.Contract.Logic(&_BridgeToken.CallOpts)
}

// Logic is a free data retrieval call binding the contract method 0xd7dfa0dd.
//
// Solidity: function logic() constant returns(address)
func (_BridgeToken *BridgeTokenCallerSession) Logic() (common.Address, error) {
	return _BridgeToken.Contract.Logic(&_BridgeToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BridgeToken *BridgeTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BridgeToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BridgeToken *BridgeTokenSession) Name() (string, error) {
	return _BridgeToken.Contract.Name(&_BridgeToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_BridgeToken *BridgeTokenCallerSession) Name() (string, error) {
	return _BridgeToken.Contract.Name(&_BridgeToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BridgeToken *BridgeTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BridgeToken *BridgeTokenSession) Owner() (common.Address, error) {
	return _BridgeToken.Contract.Owner(&_BridgeToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BridgeToken *BridgeTokenCallerSession) Owner() (common.Address, error) {
	return _BridgeToken.Contract.Owner(&_BridgeToken.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_BridgeToken *BridgeTokenCaller) Stopped(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BridgeToken.contract.Call(opts, out, "stopped")
	return *ret0, err
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_BridgeToken *BridgeTokenSession) Stopped() (bool, error) {
	return _BridgeToken.Contract.Stopped(&_BridgeToken.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_BridgeToken *BridgeTokenCallerSession) Stopped() (bool, error) {
	return _BridgeToken.Contract.Stopped(&_BridgeToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(bytes32)
func (_BridgeToken *BridgeTokenCaller) Symbol(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _BridgeToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(bytes32)
func (_BridgeToken *BridgeTokenSession) Symbol() ([32]byte, error) {
	return _BridgeToken.Contract.Symbol(&_BridgeToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(bytes32)
func (_BridgeToken *BridgeTokenCallerSession) Symbol() ([32]byte, error) {
	return _BridgeToken.Contract.Symbol(&_BridgeToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BridgeToken *BridgeTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BridgeToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BridgeToken *BridgeTokenSession) TotalSupply() (*big.Int, error) {
	return _BridgeToken.Contract.TotalSupply(&_BridgeToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BridgeToken *BridgeTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BridgeToken.Contract.TotalSupply(&_BridgeToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(guy address, wad uint256) returns(bool)
func (_BridgeToken *BridgeTokenTransactor) Approve(opts *bind.TransactOpts, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeToken.contract.Transact(opts, "approve", guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(guy address, wad uint256) returns(bool)
func (_BridgeToken *BridgeTokenSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.Approve(&_BridgeToken.TransactOpts, guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(guy address, wad uint256) returns(bool)
func (_BridgeToken *BridgeTokenTransactorSession) Approve(guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.Approve(&_BridgeToken.TransactOpts, guy, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x90bc1693.
//
// Solidity: function burn(wad uint128) returns()
func (_BridgeToken *BridgeTokenTransactor) Burn(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _BridgeToken.contract.Transact(opts, "burn", wad)
}

// Burn is a paid mutator transaction binding the contract method 0x90bc1693.
//
// Solidity: function burn(wad uint128) returns()
func (_BridgeToken *BridgeTokenSession) Burn(wad *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.Burn(&_BridgeToken.TransactOpts, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x90bc1693.
//
// Solidity: function burn(wad uint128) returns()
func (_BridgeToken *BridgeTokenTransactorSession) Burn(wad *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.Burn(&_BridgeToken.TransactOpts, wad)
}

// MintFor is a paid mutator transaction binding the contract method 0x27cfc219.
//
// Solidity: function mintFor(recipient address, wad uint128) returns()
func (_BridgeToken *BridgeTokenTransactor) MintFor(opts *bind.TransactOpts, recipient common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeToken.contract.Transact(opts, "mintFor", recipient, wad)
}

// MintFor is a paid mutator transaction binding the contract method 0x27cfc219.
//
// Solidity: function mintFor(recipient address, wad uint128) returns()
func (_BridgeToken *BridgeTokenSession) MintFor(recipient common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.MintFor(&_BridgeToken.TransactOpts, recipient, wad)
}

// MintFor is a paid mutator transaction binding the contract method 0x27cfc219.
//
// Solidity: function mintFor(recipient address, wad uint128) returns()
func (_BridgeToken *BridgeTokenTransactorSession) MintFor(recipient common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.MintFor(&_BridgeToken.TransactOpts, recipient, wad)
}

// Pull is a paid mutator transaction binding the contract method 0x8402181f.
//
// Solidity: function pull(src address, wad uint128) returns(bool)
func (_BridgeToken *BridgeTokenTransactor) Pull(opts *bind.TransactOpts, src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeToken.contract.Transact(opts, "pull", src, wad)
}

// Pull is a paid mutator transaction binding the contract method 0x8402181f.
//
// Solidity: function pull(src address, wad uint128) returns(bool)
func (_BridgeToken *BridgeTokenSession) Pull(src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.Pull(&_BridgeToken.TransactOpts, src, wad)
}

// Pull is a paid mutator transaction binding the contract method 0x8402181f.
//
// Solidity: function pull(src address, wad uint128) returns(bool)
func (_BridgeToken *BridgeTokenTransactorSession) Pull(src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.Pull(&_BridgeToken.TransactOpts, src, wad)
}

// RepayUou is a paid mutator transaction binding the contract method 0x0badbc55.
//
// Solidity: function repayUou(brgAmount uint128, vault address, uouIndex uint256) returns()
func (_BridgeToken *BridgeTokenTransactor) RepayUou(opts *bind.TransactOpts, brgAmount *big.Int, vault common.Address, uouIndex *big.Int) (*types.Transaction, error) {
	return _BridgeToken.contract.Transact(opts, "repayUou", brgAmount, vault, uouIndex)
}

// RepayUou is a paid mutator transaction binding the contract method 0x0badbc55.
//
// Solidity: function repayUou(brgAmount uint128, vault address, uouIndex uint256) returns()
func (_BridgeToken *BridgeTokenSession) RepayUou(brgAmount *big.Int, vault common.Address, uouIndex *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.RepayUou(&_BridgeToken.TransactOpts, brgAmount, vault, uouIndex)
}

// RepayUou is a paid mutator transaction binding the contract method 0x0badbc55.
//
// Solidity: function repayUou(brgAmount uint128, vault address, uouIndex uint256) returns()
func (_BridgeToken *BridgeTokenTransactorSession) RepayUou(brgAmount *big.Int, vault common.Address, uouIndex *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.RepayUou(&_BridgeToken.TransactOpts, brgAmount, vault, uouIndex)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_BridgeToken *BridgeTokenTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _BridgeToken.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_BridgeToken *BridgeTokenSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _BridgeToken.Contract.SetAuthority(&_BridgeToken.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_BridgeToken *BridgeTokenTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _BridgeToken.Contract.SetAuthority(&_BridgeToken.TransactOpts, authority_)
}

// SetLogic is a paid mutator transaction binding the contract method 0x71857000.
//
// Solidity: function setLogic(logic_ address) returns(bool)
func (_BridgeToken *BridgeTokenTransactor) SetLogic(opts *bind.TransactOpts, logic_ common.Address) (*types.Transaction, error) {
	return _BridgeToken.contract.Transact(opts, "setLogic", logic_)
}

// SetLogic is a paid mutator transaction binding the contract method 0x71857000.
//
// Solidity: function setLogic(logic_ address) returns(bool)
func (_BridgeToken *BridgeTokenSession) SetLogic(logic_ common.Address) (*types.Transaction, error) {
	return _BridgeToken.Contract.SetLogic(&_BridgeToken.TransactOpts, logic_)
}

// SetLogic is a paid mutator transaction binding the contract method 0x71857000.
//
// Solidity: function setLogic(logic_ address) returns(bool)
func (_BridgeToken *BridgeTokenTransactorSession) SetLogic(logic_ common.Address) (*types.Transaction, error) {
	return _BridgeToken.Contract.SetLogic(&_BridgeToken.TransactOpts, logic_)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(name_ string) returns()
func (_BridgeToken *BridgeTokenTransactor) SetName(opts *bind.TransactOpts, name_ string) (*types.Transaction, error) {
	return _BridgeToken.contract.Transact(opts, "setName", name_)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(name_ string) returns()
func (_BridgeToken *BridgeTokenSession) SetName(name_ string) (*types.Transaction, error) {
	return _BridgeToken.Contract.SetName(&_BridgeToken.TransactOpts, name_)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(name_ string) returns()
func (_BridgeToken *BridgeTokenTransactorSession) SetName(name_ string) (*types.Transaction, error) {
	return _BridgeToken.Contract.SetName(&_BridgeToken.TransactOpts, name_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_BridgeToken *BridgeTokenTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _BridgeToken.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_BridgeToken *BridgeTokenSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _BridgeToken.Contract.SetOwner(&_BridgeToken.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_BridgeToken *BridgeTokenTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _BridgeToken.Contract.SetOwner(&_BridgeToken.TransactOpts, owner_)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_BridgeToken *BridgeTokenTransactor) Start(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeToken.contract.Transact(opts, "start")
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_BridgeToken *BridgeTokenSession) Start() (*types.Transaction, error) {
	return _BridgeToken.Contract.Start(&_BridgeToken.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_BridgeToken *BridgeTokenTransactorSession) Start() (*types.Transaction, error) {
	return _BridgeToken.Contract.Start(&_BridgeToken.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_BridgeToken *BridgeTokenTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeToken.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_BridgeToken *BridgeTokenSession) Stop() (*types.Transaction, error) {
	return _BridgeToken.Contract.Stop(&_BridgeToken.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_BridgeToken *BridgeTokenTransactorSession) Stop() (*types.Transaction, error) {
	return _BridgeToken.Contract.Stop(&_BridgeToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(dst address, wad uint256) returns(bool)
func (_BridgeToken *BridgeTokenTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeToken.contract.Transact(opts, "transfer", dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(dst address, wad uint256) returns(bool)
func (_BridgeToken *BridgeTokenSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.Transfer(&_BridgeToken.TransactOpts, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(dst address, wad uint256) returns(bool)
func (_BridgeToken *BridgeTokenTransactorSession) Transfer(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.Transfer(&_BridgeToken.TransactOpts, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_BridgeToken *BridgeTokenTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeToken.contract.Transact(opts, "transferFrom", src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_BridgeToken *BridgeTokenSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.TransferFrom(&_BridgeToken.TransactOpts, src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_BridgeToken *BridgeTokenTransactorSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeToken.Contract.TransferFrom(&_BridgeToken.TransactOpts, src, dst, wad)
}

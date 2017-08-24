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
const SweetTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"guy\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"push\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"uint128\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"logic_\",\"type\":\"address\"}],\"name\":\"setLogic\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"pull\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"name_\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"logic\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"_allowance\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"name_\",\"type\":\"string\"},{\"name\":\"symbol_\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"constructor\"},{\"payable\":true,\"type\":\"fallback\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"foo\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"bar\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fax\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"UnauthorizedAccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// SweetTokenBin is the compiled bytecode used for deploying new contracts.
const SweetTokenBin = `"0x6060604052601260045534156200001557600080fd5b6040516200140b3803806200140b833981016040528080518201919060200180519150505b81815b5b60018054600160a060020a03191633600160a060020a03169081179091557fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a25b60038280516200009e929160200190620000b0565b5060028190555b50505b50506200015a565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620000f357805160ff191683800117855562000123565b8280016001018555821562000123579182015b828111156200012357825182559160200191906001019062000106565b5b506200013292915062000136565b5090565b6200015791905b808211156200013257600081556001016200013d565b5090565b90565b6112a1806200016a6000396000f300606060405236156101175763ffffffff60e060020a60003504166306fdde03811461012057806307da68f5146101ab578063095ea7b3146101c057806313af4035146101f657806318160ddd1461021757806323b872dd1461023c578063313ce567146102785780633452f51d1461029d57806369d3e20e146102e557806370a082311461030f578063718570001461034057806375f12b21146103735780637a9e5e4b1461039a5780638402181f146103bb5780638da5cb5b1461040357806390bc16931461043257806395d89b411461045c578063a9059cbb14610481578063be9a6555146104b7578063bf7e214f146104cc578063c47f0027146104fb578063d7dfa0dd1461054e578063dd62ed3e1461057d575b5b600080fd5b5b005b341561012b57600080fd5b6101336105b4565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156101705780820151818401525b602001610157565b50505050905090810190601f16801561019d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156101b657600080fd5b61011d610652565b005b34156101cb57600080fd5b6101e2600160a060020a03600435166024356106f1565b604051901515815260200160405180910390f35b341561020157600080fd5b61011d600160a060020a0360043516610845565b005b341561022257600080fd5b61022a6108c3565b60405190815260200160405180910390f35b341561024757600080fd5b6101e2600160a060020a036004358116906024351660443561092d565b604051901515815260200160405180910390f35b341561028357600080fd5b61022a610a7c565b60405190815260200160405180910390f35b34156102a857600080fd5b6101e2600160a060020a03600435166fffffffffffffffffffffffffffffffff60243516610a82565b604051901515815260200160405180910390f35b34156102f057600080fd5b61011d6fffffffffffffffffffffffffffffffff60043516610aa9565b005b341561031a57600080fd5b61022a600160a060020a0360043516610aaf565b60405190815260200160405180910390f35b341561034b57600080fd5b6101e2600160a060020a0360043516610b2c565b604051901515815260200160405180910390f35b341561037e57600080fd5b6101e2610bd8565b604051901515815260200160405180910390f35b34156103a557600080fd5b61011d600160a060020a0360043516610be8565b005b34156103c657600080fd5b6101e2600160a060020a03600435166fffffffffffffffffffffffffffffffff60243516610c66565b604051901515815260200160405180910390f35b341561040e57600080fd5b610416610c8e565b604051600160a060020a03909116815260200160405180910390f35b341561043d57600080fd5b61011d6fffffffffffffffffffffffffffffffff60043516610c9d565b005b341561046757600080fd5b61022a610daa565b60405190815260200160405180910390f35b341561048c57600080fd5b6101e2600160a060020a0360043516602435610db0565b604051901515815260200160405180910390f35b34156104c257600080fd5b61011d610efe565b005b34156104d757600080fd5b610416610f97565b604051600160a060020a03909116815260200160405180910390f35b341561050657600080fd5b61011d60046024813581810190830135806020601f82018190048102016040519081016040528181529291906020840183838082843750949650610fa695505050505050565b005b341561055957600080fd5b610416610fdd565b604051600160a060020a03909116815260200160405180910390f35b341561058857600080fd5b61022a600160a060020a0360043581169060243516610fec565b60405190815260200160405180910390f35b60038054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561064a5780601f1061061f5761010080835404028352916020019161064a565b820191906000526020600020905b81548152906001019060200180831161062d57829003601f168201915b505050505081565b61066833600035600160e060020a031916611072565b151561067057fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a46001805474ff0000000000000000000000000000000000000000191660a060020a1790555b5b50505b565b600154600090819060a060020a900460ff161561070a57fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a4600554600160a060020a031663e1f21c6733888860006040516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b15156107cc57600080fd5b6102c65a03f115156107dd57600080fd5b505050604051805193505082156108365785600160a060020a031633600160a060020a03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258760405190815260200160405180910390a35b8293505b5b50505b5092915050565b61085b33600035600160e060020a031916611072565b151561086357fe5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a25b5b50565b600554600090600160a060020a03166318160ddd82604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561090d57600080fd5b6102c65a03f1151561091e57600080fd5b50505060405180519150505b90565b600154600090819060a060020a900460ff161561094657fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a4600554600160a060020a03166323b872dd88888860006040516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b1515610a0857600080fd5b6102c65a03f11515610a1957600080fd5b50505060405180519050925085600160a060020a031687600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8760405190815260200160405180910390a38293505b5b50505b509392505050565b60045481565b6000610aa083836fffffffffffffffffffffffffffffffff16610db0565b90505b92915050565bfe5b5b50565b600554600090600160a060020a03166370a0823183836040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b1515610b0a57600080fd5b6102c65a03f11515610b1b57600080fd5b50505060405180519150505b919050565b6000610b4433600035600160e060020a031916611072565b1515610b4c57fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a46005805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038616179055600192505b5b50505b919050565b60015460a060020a900460ff1681565b610bfe33600035600160e060020a031916611072565b1515610c0657fe5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167f1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada460405160405180910390a25b5b50565b6000610aa08333846fffffffffffffffffffffffffffffffff1661092d565b90505b92915050565b600154600160a060020a031681565b610cb333600035600160e060020a031916611072565b1515610cbb57fe5b60015460a060020a900460ff1615610ccf57fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a4600554600160a060020a0316637261e469338560405160e060020a63ffffffff8516028152600160a060020a0390921660048301526fffffffffffffffffffffffffffffffff166024820152604401600060405180830381600087803b1515610d8d57600080fd5b6102c65a03f11515610d9e57600080fd5b5050505b5b50505b5b50565b60025481565b600154600090819060a060020a900460ff1615610dc957fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a4600554600160a060020a031663beabacc833888860006040516020015260405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401602060405180830381600087803b1515610e8b57600080fd5b6102c65a03f11515610e9c57600080fd5b50505060405180519050925085600160a060020a031633600160a060020a03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8760405190815260200160405180910390a38293505b5b50505b5092915050565b610f1433600035600160e060020a031916611072565b1515610f1c57fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a46001805474ff0000000000000000000000000000000000000000191690555b5b50505b565b600054600160a060020a031681565b610fbc33600035600160e060020a031916611072565b1515610fc457fe5b60038180516106eb9291602001906111d5565b505b5b50565b600554600160a060020a031681565b600554600090600160a060020a031663dd62ed3e8484846040516020015260405160e060020a63ffffffff8516028152600160a060020a03928316600482015291166024820152604401602060405180830381600087803b151561104f57600080fd5b6102c65a03f1151561106057600080fd5b50505060405180519150505b92915050565b600030600160a060020a031683600160a060020a0316141561109657506001610aa3565b600154600160a060020a0384811691161480156110bc5750600054600160a060020a0316155b156110c957506001610aa3565b600054600160a060020a03161515611131577feb91385a0d70e44b915093d9ddfe6c8b41f2c56729431b1405f40aee5d874be78383604051600160a060020a039092168252600160e060020a03191660208201526040908101905180910390a1506000610aa3565b60008054600160a060020a03169063b7009613908590309086906040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152600160e060020a03199091166044820152606401602060405180830381600087803b15156111aa57600080fd5b6102c65a03f115156111bb57600080fd5b505050604051805190509050610aa3565b5b5b5b92915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061121657805160ff1916838001178555611243565b82800160010185558215611243579182015b82811115611243578251825591602001919060010190611228565b5b50611250929150611254565b5090565b61092a91905b80821115611250576000815560010161125a565b5090565b905600a165627a7a723058200ce31c822994fb518a499a4837471190627a6b80b4e84597835c2bf1e81563e00029"`

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
// Solidity: function mint( uint128) returns()
func (_SweetToken *SweetTokenTransactor) Mint(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _SweetToken.contract.Transact(opts, "mint", arg0)
}

// Mint is a paid mutator transaction binding the contract method 0x69d3e20e.
//
// Solidity: function mint( uint128) returns()
func (_SweetToken *SweetTokenSession) Mint(arg0 *big.Int) (*types.Transaction, error) {
	return _SweetToken.Contract.Mint(&_SweetToken.TransactOpts, arg0)
}

// Mint is a paid mutator transaction binding the contract method 0x69d3e20e.
//
// Solidity: function mint( uint128) returns()
func (_SweetToken *SweetTokenTransactorSession) Mint(arg0 *big.Int) (*types.Transaction, error) {
	return _SweetToken.Contract.Mint(&_SweetToken.TransactOpts, arg0)
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

// Push is a paid mutator transaction binding the contract method 0x3452f51d.
//
// Solidity: function push(dst address, wad uint128) returns(bool)
func (_SweetToken *SweetTokenTransactor) Push(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetToken.contract.Transact(opts, "push", dst, wad)
}

// Push is a paid mutator transaction binding the contract method 0x3452f51d.
//
// Solidity: function push(dst address, wad uint128) returns(bool)
func (_SweetToken *SweetTokenSession) Push(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetToken.Contract.Push(&_SweetToken.TransactOpts, dst, wad)
}

// Push is a paid mutator transaction binding the contract method 0x3452f51d.
//
// Solidity: function push(dst address, wad uint128) returns(bool)
func (_SweetToken *SweetTokenTransactorSession) Push(dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetToken.Contract.Push(&_SweetToken.TransactOpts, dst, wad)
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

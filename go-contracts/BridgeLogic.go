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

// BridgeLogicABI is the input ABI used to generate the binding from.
const BridgeLogicABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token_\",\"type\":\"address\"}],\"name\":\"setToken\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"guy\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"token_\",\"type\":\"address\"},{\"name\":\"data_\",\"type\":\"address\"},{\"name\":\"supply_\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"foo\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"bar\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fax\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"UnauthorizedAccess\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// BridgeLogicBin is the compiled bytecode used for deploying new contracts.
const BridgeLogicBin = `"0x6060604052341561000f57600080fd5b604051606080611a328339810160405280805191906020018051919060200180519150505b5b60018054600160a060020a03191633600160a060020a03169081179091557fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a25b600160a060020a038316151561009557600080fd5b600160a060020a0382161515610110573081336100b0610150565b600160a060020a039384168152602081019290925290911660408083019190915260609091019051809103906000f08015156100eb57600080fd5b60028054600160a060020a031916600160a060020a039290921691909117905561012c565b60028054600160a060020a031916600160a060020a0384161790555b60038054600160a060020a031916600160a060020a0385161790555b505050610160565b6040516103b68061167c83390190565b61150d8061016f6000396000f300606060405236156100e05763ffffffff60e060020a60003504166307da68f581146100e557806313af4035146100fa578063144fa6d71461011b57806318160ddd1461013c57806323b872dd1461016157806369d3e20e1461019d57806370a08231146101c75780637261e469146101f857806373d4a13a1461022e57806375f12b211461025d5780637a9e5e4b146102845780638da5cb5b146102a5578063be9a6555146102d4578063beabacc8146102e9578063bf7e214f14610325578063dd62ed3e14610354578063e1f21c671461038b578063fc0c546a146103c7575b600080fd5b34156100f057600080fd5b6100f86103f6565b005b341561010557600080fd5b6100f8600160a060020a03600435166104a6565b005b341561012657600080fd5b6100f8600160a060020a0360043516610524565b005b341561014757600080fd5b61014f61056f565b60405190815260200160405180910390f35b341561016c57600080fd5b610189600160a060020a03600435811690602435166044356105d9565b604051901515815260200160405180910390f35b34156101a857600080fd5b6100f86fffffffffffffffffffffffffffffffff600435166109b7565b005b34156101d257600080fd5b61014f600160a060020a0360043516610c4b565b60405190815260200160405180910390f35b341561020357600080fd5b6100f8600160a060020a03600435166fffffffffffffffffffffffffffffffff60243516610cc8565b005b341561023957600080fd5b610241610ec9565b604051600160a060020a03909116815260200160405180910390f35b341561026857600080fd5b610189610ed8565b604051901515815260200160405180910390f35b341561028f57600080fd5b6100f8600160a060020a0360043516610ef9565b005b34156102b057600080fd5b610241610f77565b604051600160a060020a03909116815260200160405180910390f35b34156102df57600080fd5b6100f8610f86565b005b34156102f457600080fd5b610189600160a060020a036004358116906024351660443561101f565b604051901515815260200160405180910390f35b341561033057600080fd5b610241611213565b604051600160a060020a03909116815260200160405180910390f35b341561035f57600080fd5b61014f600160a060020a0360043581169060243516611222565b60405190815260200160405180910390f35b341561039657600080fd5b610189600160a060020a03600435811690602435166044356112a8565b604051901515815260200160405180910390f35b34156103d257600080fd5b610241611347565b604051600160a060020a03909116815260200160405180910390f35b61040c33600035600160e060020a031916611356565b151561041457fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a46001805474ff00000000000000000000000000000000000000001916740100000000000000000000000000000000000000001790555b5b50505b565b6104bc33600035600160e060020a031916611356565b15156104c457fe5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a25b5b50565b61053a33600035600160e060020a031916611356565b151561054257fe5b6003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b5b50565b600254600090600160a060020a031663047fc9aa82604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156105b957600080fd5b6102c65a03f115156105ca57600080fd5b50505060405180519150505b90565b60035460009033600160a060020a039081169116146105f457fe5b6002548290600160a060020a03166327e235e38660006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b151561064f57600080fd5b6102c65a03f1151561066057600080fd5b505050604051805190501015151561067757600080fd5b6002548290600160a060020a031663a32ce11e868660006040516020015260405160e060020a63ffffffff8516028152600160a060020a03928316600482015291166024820152604401602060405180830381600087803b15156106da57600080fd5b6102c65a03f115156106eb57600080fd5b505050604051805190501015151561070257600080fd5b600254600160a060020a031663e32d5cf8858561078f8463a32ce11e848460006040516020015260405160e060020a63ffffffff8516028152600160a060020a03928316600482015291166024820152604401602060405180830381600087803b151561076e57600080fd5b6102c65a03f1151561077f57600080fd5b50505060405180519050876114b9565b60405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401600060405180830381600087803b15156107de57600080fd5b6102c65a03f115156107ef57600080fd5b5050600254600160a060020a031690506328e69b1685610877836327e235e38360006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b151561085657600080fd5b6102c65a03f1151561086757600080fd5b50505060405180519050866114b9565b60405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401600060405180830381600087803b15156108ba57600080fd5b6102c65a03f115156108cb57600080fd5b5050600254600160a060020a031690506328e69b1684610953836327e235e38360006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b151561093257600080fd5b6102c65a03f1151561094357600080fd5b50505060405180519050866114cd565b60405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401600060405180830381600087803b151561099657600080fd5b6102c65a03f115156109a757600080fd5b505050600190505b5b9392505050565b60035433600160a060020a039081169116146109cf57fe5b600254600160a060020a03166328e69b1681638da5cb5b6000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610a1d57600080fd5b6102c65a03f11515610a2e57600080fd5b5050506040518051600254909150610b2090600160a060020a03166327e235e381638da5cb5b6000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610a8b57600080fd5b6102c65a03f11515610a9c57600080fd5b5050506040518051905060006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b1515610aed57600080fd5b6102c65a03f11515610afe57600080fd5b50505060405180519050856fffffffffffffffffffffffffffffffff166114cd565b60405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401600060405180830381600087803b1515610b6357600080fd5b6102c65a03f11515610b7457600080fd5b5050600254600160a060020a03169050633b4c4b25610bfc8263047fc9aa6000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610bc957600080fd5b6102c65a03f11515610bda57600080fd5b50505060405180519050846fffffffffffffffffffffffffffffffff166114cd565b60405160e060020a63ffffffff84160281526004810191909152602401600060405180830381600087803b1515610c3257600080fd5b6102c65a03f11515610c4357600080fd5b5050505b5b50565b600254600090600160a060020a03166327e235e383836040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b1515610ca657600080fd5b6102c65a03f11515610cb757600080fd5b50505060405180519150505b919050565b60035433600160a060020a03908116911614610ce057fe5b806fffffffffffffffffffffffffffffffff16610cfc83610c4b565b1015610d0757600080fd5b600254600160a060020a03166328e69b1683610d9d836327e235e38360006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b1515610d6a57600080fd5b6102c65a03f11515610d7b57600080fd5b50505060405180519050856fffffffffffffffffffffffffffffffff166114b9565b60405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401600060405180830381600087803b1515610de057600080fd5b6102c65a03f11515610df157600080fd5b5050600254600160a060020a03169050633b4c4b25610e798263047fc9aa6000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610e4657600080fd5b6102c65a03f11515610e5757600080fd5b50505060405180519050846fffffffffffffffffffffffffffffffff166114b9565b60405160e060020a63ffffffff84160281526004810191909152602401600060405180830381600087803b1515610eaf57600080fd5b6102c65a03f11515610ec057600080fd5b5050505b5b5050565b600254600160a060020a031681565b60015474010000000000000000000000000000000000000000900460ff1681565b610f0f33600035600160e060020a031916611356565b1515610f1757fe5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167f1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada460405160405180910390a25b5b50565b600154600160a060020a031681565b610f9c33600035600160e060020a031916611356565b1515610fa457fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a46001805474ff0000000000000000000000000000000000000000191690555b5b50505b565b60035460009033600160a060020a0390811691161461103a57fe5b8161104485610c4b565b101561104f57600080fd5b600254600160a060020a03166328e69b1685610877836327e235e38360006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b151561085657600080fd5b6102c65a03f1151561086757600080fd5b50505060405180519050866114b9565b60405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401600060405180830381600087803b15156108ba57600080fd5b6102c65a03f115156108cb57600080fd5b5050600254600160a060020a031690506328e69b1684610953836327e235e38360006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b151561093257600080fd5b6102c65a03f1151561094357600080fd5b50505060405180519050866114cd565b60405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401600060405180830381600087803b151561099657600080fd5b6102c65a03f115156109a757600080fd5b505050600190505b5b9392505050565b600054600160a060020a031681565b600254600090600160a060020a031663a32ce11e8484846040516020015260405160e060020a63ffffffff8516028152600160a060020a03928316600482015291166024820152604401602060405180830381600087803b151561128557600080fd5b6102c65a03f1151561129657600080fd5b50505060405180519150505b92915050565b60035460009033600160a060020a039081169116146112c357fe5b600254600160a060020a031663e32d5cf885858560405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401600060405180830381600087803b151561099657600080fd5b6102c65a03f115156109a757600080fd5b505050600190505b5b9392505050565b600354600160a060020a031681565b600030600160a060020a031683600160a060020a0316141561137a575060016112a2565b600154600160a060020a0384811691161480156113a05750600054600160a060020a0316155b156113ad575060016112a2565b600054600160a060020a03161515611415577feb91385a0d70e44b915093d9ddfe6c8b41f2c56729431b1405f40aee5d874be78383604051600160a060020a039092168252600160e060020a03191660208201526040908101905180910390a15060006112a2565b60008054600160a060020a03169063b7009613908590309086906040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152600160e060020a03199091166044820152606401602060405180830381600087803b151561148e57600080fd5b6102c65a03f1151561149f57600080fd5b5050506040518051905090506112a2565b5b5b5b92915050565b808203828111156112a257fe5b5b92915050565b808201828110156112a257fe5b5b929150505600a165627a7a723058202d67bc145286c3a013441b5a00e9e0a4614ac6bc19a3e11df91825810c808a1a00296060604052341561000f57600080fd5b6040516060806103b68339810160405280805191906020018051919060200180519150505b60038054600160a060020a03808616600160a060020a03199283161790925560008481556004805485851693169290921791829055911681526001602052604090208290555b5050505b6103298061008d6000396000f3006060604052361561008b5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663047fc9aa811461009057806327e235e3146100b557806328e69b16146100e65780633a3d523f1461010a5780633b4c4b251461012b5780638da5cb5b14610143578063a32ce11e14610172578063e32d5cf8146101a9575b600080fd5b341561009b57600080fd5b6100a36101d3565b60405190815260200160405180910390f35b34156100c057600080fd5b6100a3600160a060020a03600435166101d9565b60405190815260200160405180910390f35b34156100f157600080fd5b610108600160a060020a03600435166024356101eb565b005b341561011557600080fd5b610108600160a060020a0360043516610224565b005b341561013657600080fd5b610108600435610268565b005b341561014e57600080fd5b61015661028a565b604051600160a060020a03909116815260200160405180910390f35b341561017d57600080fd5b6100a3600160a060020a0360043581169060243516610299565b60405190815260200160405180910390f35b34156101b457600080fd5b610108600160a060020a03600435811690602435166044356102b6565b005b60005481565b60016020526000908152604090205481565b60035433600160a060020a0390811691161461020357fe5b600160a060020a03821660009081526001602052604090208190555b5b5050565b60045433600160a060020a0390811691161461023c57fe5b6003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b50565b60035433600160a060020a0390811691161461028057fe5b60008190555b5b50565b600454600160a060020a031681565b600260209081526000928352604080842090915290825290205481565b60035433600160a060020a039081169116146102ce57fe5b600160a060020a0380841660009081526002602090815260408083209386168352929052208190555b5b5050505600a165627a7a72305820b555670b782a6fbd07366fb67f18ffc34a831f86e97df3126bb5cf286ddfb1740029"`

// DeployBridgeLogic deploys a new Ethereum contract, binding an instance of BridgeLogic to it.
func DeployBridgeLogic(auth *bind.TransactOpts, backend bind.ContractBackend, token_ common.Address, data_ common.Address, supply_ *big.Int) (common.Address, *types.Transaction, *BridgeLogic, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeLogicABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BridgeLogicBin), backend, token_, data_, supply_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BridgeLogic{BridgeLogicCaller: BridgeLogicCaller{contract: contract}, BridgeLogicTransactor: BridgeLogicTransactor{contract: contract}}, nil
}

// BridgeLogic is an auto generated Go binding around an Ethereum contract.
type BridgeLogic struct {
	BridgeLogicCaller     // Read-only binding to the contract
	BridgeLogicTransactor // Write-only binding to the contract
}

// BridgeLogicCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeLogicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeLogicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeLogicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeLogicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeLogicSession struct {
	Contract     *BridgeLogic      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeLogicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeLogicCallerSession struct {
	Contract *BridgeLogicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BridgeLogicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeLogicTransactorSession struct {
	Contract     *BridgeLogicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BridgeLogicRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeLogicRaw struct {
	Contract *BridgeLogic // Generic contract binding to access the raw methods on
}

// BridgeLogicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeLogicCallerRaw struct {
	Contract *BridgeLogicCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeLogicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeLogicTransactorRaw struct {
	Contract *BridgeLogicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridgeLogic creates a new instance of BridgeLogic, bound to a specific deployed contract.
func NewBridgeLogic(address common.Address, backend bind.ContractBackend) (*BridgeLogic, error) {
	contract, err := bindBridgeLogic(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BridgeLogic{BridgeLogicCaller: BridgeLogicCaller{contract: contract}, BridgeLogicTransactor: BridgeLogicTransactor{contract: contract}}, nil
}

// NewBridgeLogicCaller creates a new read-only instance of BridgeLogic, bound to a specific deployed contract.
func NewBridgeLogicCaller(address common.Address, caller bind.ContractCaller) (*BridgeLogicCaller, error) {
	contract, err := bindBridgeLogic(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeLogicCaller{contract: contract}, nil
}

// NewBridgeLogicTransactor creates a new write-only instance of BridgeLogic, bound to a specific deployed contract.
func NewBridgeLogicTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeLogicTransactor, error) {
	contract, err := bindBridgeLogic(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &BridgeLogicTransactor{contract: contract}, nil
}

// bindBridgeLogic binds a generic wrapper to an already deployed contract.
func bindBridgeLogic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeLogicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeLogic *BridgeLogicRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BridgeLogic.Contract.BridgeLogicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeLogic *BridgeLogicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeLogic.Contract.BridgeLogicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeLogic *BridgeLogicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeLogic.Contract.BridgeLogicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BridgeLogic *BridgeLogicCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BridgeLogic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BridgeLogic *BridgeLogicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeLogic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BridgeLogic *BridgeLogicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BridgeLogic.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(src address, guy address) constant returns(uint256)
func (_BridgeLogic *BridgeLogicCaller) Allowance(opts *bind.CallOpts, src common.Address, guy common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BridgeLogic.contract.Call(opts, out, "allowance", src, guy)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(src address, guy address) constant returns(uint256)
func (_BridgeLogic *BridgeLogicSession) Allowance(src common.Address, guy common.Address) (*big.Int, error) {
	return _BridgeLogic.Contract.Allowance(&_BridgeLogic.CallOpts, src, guy)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(src address, guy address) constant returns(uint256)
func (_BridgeLogic *BridgeLogicCallerSession) Allowance(src common.Address, guy common.Address) (*big.Int, error) {
	return _BridgeLogic.Contract.Allowance(&_BridgeLogic.CallOpts, src, guy)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_BridgeLogic *BridgeLogicCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeLogic.contract.Call(opts, out, "authority")
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_BridgeLogic *BridgeLogicSession) Authority() (common.Address, error) {
	return _BridgeLogic.Contract.Authority(&_BridgeLogic.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_BridgeLogic *BridgeLogicCallerSession) Authority() (common.Address, error) {
	return _BridgeLogic.Contract.Authority(&_BridgeLogic.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(src address) constant returns(uint256)
func (_BridgeLogic *BridgeLogicCaller) BalanceOf(opts *bind.CallOpts, src common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BridgeLogic.contract.Call(opts, out, "balanceOf", src)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(src address) constant returns(uint256)
func (_BridgeLogic *BridgeLogicSession) BalanceOf(src common.Address) (*big.Int, error) {
	return _BridgeLogic.Contract.BalanceOf(&_BridgeLogic.CallOpts, src)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(src address) constant returns(uint256)
func (_BridgeLogic *BridgeLogicCallerSession) BalanceOf(src common.Address) (*big.Int, error) {
	return _BridgeLogic.Contract.BalanceOf(&_BridgeLogic.CallOpts, src)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(address)
func (_BridgeLogic *BridgeLogicCaller) Data(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeLogic.contract.Call(opts, out, "data")
	return *ret0, err
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(address)
func (_BridgeLogic *BridgeLogicSession) Data() (common.Address, error) {
	return _BridgeLogic.Contract.Data(&_BridgeLogic.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(address)
func (_BridgeLogic *BridgeLogicCallerSession) Data() (common.Address, error) {
	return _BridgeLogic.Contract.Data(&_BridgeLogic.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BridgeLogic *BridgeLogicCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeLogic.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BridgeLogic *BridgeLogicSession) Owner() (common.Address, error) {
	return _BridgeLogic.Contract.Owner(&_BridgeLogic.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_BridgeLogic *BridgeLogicCallerSession) Owner() (common.Address, error) {
	return _BridgeLogic.Contract.Owner(&_BridgeLogic.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_BridgeLogic *BridgeLogicCaller) Stopped(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BridgeLogic.contract.Call(opts, out, "stopped")
	return *ret0, err
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_BridgeLogic *BridgeLogicSession) Stopped() (bool, error) {
	return _BridgeLogic.Contract.Stopped(&_BridgeLogic.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_BridgeLogic *BridgeLogicCallerSession) Stopped() (bool, error) {
	return _BridgeLogic.Contract.Stopped(&_BridgeLogic.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_BridgeLogic *BridgeLogicCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BridgeLogic.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_BridgeLogic *BridgeLogicSession) Token() (common.Address, error) {
	return _BridgeLogic.Contract.Token(&_BridgeLogic.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_BridgeLogic *BridgeLogicCallerSession) Token() (common.Address, error) {
	return _BridgeLogic.Contract.Token(&_BridgeLogic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BridgeLogic *BridgeLogicCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BridgeLogic.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BridgeLogic *BridgeLogicSession) TotalSupply() (*big.Int, error) {
	return _BridgeLogic.Contract.TotalSupply(&_BridgeLogic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_BridgeLogic *BridgeLogicCallerSession) TotalSupply() (*big.Int, error) {
	return _BridgeLogic.Contract.TotalSupply(&_BridgeLogic.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(src address, guy address, wad uint256) returns(bool)
func (_BridgeLogic *BridgeLogicTransactor) Approve(opts *bind.TransactOpts, src common.Address, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeLogic.contract.Transact(opts, "approve", src, guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(src address, guy address, wad uint256) returns(bool)
func (_BridgeLogic *BridgeLogicSession) Approve(src common.Address, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeLogic.Contract.Approve(&_BridgeLogic.TransactOpts, src, guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(src address, guy address, wad uint256) returns(bool)
func (_BridgeLogic *BridgeLogicTransactorSession) Approve(src common.Address, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeLogic.Contract.Approve(&_BridgeLogic.TransactOpts, src, guy, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x7261e469.
//
// Solidity: function burn(src address, wad uint128) returns()
func (_BridgeLogic *BridgeLogicTransactor) Burn(opts *bind.TransactOpts, src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeLogic.contract.Transact(opts, "burn", src, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x7261e469.
//
// Solidity: function burn(src address, wad uint128) returns()
func (_BridgeLogic *BridgeLogicSession) Burn(src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeLogic.Contract.Burn(&_BridgeLogic.TransactOpts, src, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x7261e469.
//
// Solidity: function burn(src address, wad uint128) returns()
func (_BridgeLogic *BridgeLogicTransactorSession) Burn(src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeLogic.Contract.Burn(&_BridgeLogic.TransactOpts, src, wad)
}

// Mint is a paid mutator transaction binding the contract method 0x69d3e20e.
//
// Solidity: function mint(wad uint128) returns()
func (_BridgeLogic *BridgeLogicTransactor) Mint(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _BridgeLogic.contract.Transact(opts, "mint", wad)
}

// Mint is a paid mutator transaction binding the contract method 0x69d3e20e.
//
// Solidity: function mint(wad uint128) returns()
func (_BridgeLogic *BridgeLogicSession) Mint(wad *big.Int) (*types.Transaction, error) {
	return _BridgeLogic.Contract.Mint(&_BridgeLogic.TransactOpts, wad)
}

// Mint is a paid mutator transaction binding the contract method 0x69d3e20e.
//
// Solidity: function mint(wad uint128) returns()
func (_BridgeLogic *BridgeLogicTransactorSession) Mint(wad *big.Int) (*types.Transaction, error) {
	return _BridgeLogic.Contract.Mint(&_BridgeLogic.TransactOpts, wad)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_BridgeLogic *BridgeLogicTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _BridgeLogic.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_BridgeLogic *BridgeLogicSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _BridgeLogic.Contract.SetAuthority(&_BridgeLogic.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_BridgeLogic *BridgeLogicTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _BridgeLogic.Contract.SetAuthority(&_BridgeLogic.TransactOpts, authority_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_BridgeLogic *BridgeLogicTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _BridgeLogic.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_BridgeLogic *BridgeLogicSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _BridgeLogic.Contract.SetOwner(&_BridgeLogic.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_BridgeLogic *BridgeLogicTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _BridgeLogic.Contract.SetOwner(&_BridgeLogic.TransactOpts, owner_)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(token_ address) returns()
func (_BridgeLogic *BridgeLogicTransactor) SetToken(opts *bind.TransactOpts, token_ common.Address) (*types.Transaction, error) {
	return _BridgeLogic.contract.Transact(opts, "setToken", token_)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(token_ address) returns()
func (_BridgeLogic *BridgeLogicSession) SetToken(token_ common.Address) (*types.Transaction, error) {
	return _BridgeLogic.Contract.SetToken(&_BridgeLogic.TransactOpts, token_)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(token_ address) returns()
func (_BridgeLogic *BridgeLogicTransactorSession) SetToken(token_ common.Address) (*types.Transaction, error) {
	return _BridgeLogic.Contract.SetToken(&_BridgeLogic.TransactOpts, token_)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_BridgeLogic *BridgeLogicTransactor) Start(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeLogic.contract.Transact(opts, "start")
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_BridgeLogic *BridgeLogicSession) Start() (*types.Transaction, error) {
	return _BridgeLogic.Contract.Start(&_BridgeLogic.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_BridgeLogic *BridgeLogicTransactorSession) Start() (*types.Transaction, error) {
	return _BridgeLogic.Contract.Start(&_BridgeLogic.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_BridgeLogic *BridgeLogicTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BridgeLogic.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_BridgeLogic *BridgeLogicSession) Stop() (*types.Transaction, error) {
	return _BridgeLogic.Contract.Stop(&_BridgeLogic.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_BridgeLogic *BridgeLogicTransactorSession) Stop() (*types.Transaction, error) {
	return _BridgeLogic.Contract.Stop(&_BridgeLogic.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(src address, dst address, wad uint256) returns(bool)
func (_BridgeLogic *BridgeLogicTransactor) Transfer(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeLogic.contract.Transact(opts, "transfer", src, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(src address, dst address, wad uint256) returns(bool)
func (_BridgeLogic *BridgeLogicSession) Transfer(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeLogic.Contract.Transfer(&_BridgeLogic.TransactOpts, src, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(src address, dst address, wad uint256) returns(bool)
func (_BridgeLogic *BridgeLogicTransactorSession) Transfer(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeLogic.Contract.Transfer(&_BridgeLogic.TransactOpts, src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_BridgeLogic *BridgeLogicTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeLogic.contract.Transact(opts, "transferFrom", src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_BridgeLogic *BridgeLogicSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeLogic.Contract.TransferFrom(&_BridgeLogic.TransactOpts, src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_BridgeLogic *BridgeLogicTransactorSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _BridgeLogic.Contract.TransferFrom(&_BridgeLogic.TransactOpts, src, dst, wad)
}

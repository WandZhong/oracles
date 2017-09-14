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

// SweetTokenLogicABI is the input ABI used to generate the binding from.
const SweetTokenLogicABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token_\",\"type\":\"address\"}],\"name\":\"setToken\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"freeTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"isFree\",\"type\":\"bool\"}],\"name\":\"setFreeTransfer\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"listName\",\"type\":\"bytes32\"}],\"name\":\"addWhiteList\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"listName\",\"type\":\"bytes32\"},{\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"removeFromWhiteList\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"whiteLists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"listNames\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"listName\",\"type\":\"bytes32\"},{\"name\":\"guy\",\"type\":\"address\"}],\"name\":\"addToWhiteList\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"uint128\"}],\"name\":\"mint\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint128\"}],\"name\":\"burn\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"data\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"guy\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"token_\",\"type\":\"address\"},{\"name\":\"data_\",\"type\":\"address\"},{\"name\":\"supply_\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"UnauthorizedAccess\",\"type\":\"event\"}]"

// SweetTokenLogicBin is the compiled bytecode used for deploying new contracts.
const SweetTokenLogicBin = `"0x60606040526006805460ff19166001179055341561001c57600080fd5b6040516060806117ed8339810160405280805191906020018051919060200180519150505b8282825b5b60018054600160a060020a03191633600160a060020a03169081179091557fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a25b600160a060020a03831615156100a657600080fd5b600160a060020a0382161515610121573081336100c1610165565b600160a060020a039384168152602081019290925290911660408083019190915260609091019051809103906000f08015156100fc57600080fd5b60028054600160a060020a031916600160a060020a039290921691909117905561013d565b60028054600160a060020a031916600160a060020a0384161790555b60038054600160a060020a031916600160a060020a0385161790555b5050505b505050610175565b6040516103b68061143783390190565b6112b3806101846000396000f3006060604052361561010c5763ffffffff60e060020a60003504166313af40358114610111578063144fa6d71461013257806318160ddd14610153578063211ed6c11461017857806323b872dd1461019f5780632abfaf1f146101db5780633bd04d69146101f55780633f32aa701461020d5780634b4d0d45146102315780636627262314610267578063693382a91461028f57806369d3e20e146102b357806370a08231146102dd5780637261e4691461030e57806373d4a13a146103445780637a9e5e4b146103735780638da5cb5b14610394578063beabacc8146103c3578063bf7e214f146103ff578063dd62ed3e1461042e578063e1f21c6714610465578063fc0c546a146104a1575b600080fd5b341561011c57600080fd5b610130600160a060020a03600435166104d0565b005b341561013d57600080fd5b610130600160a060020a036004351661054e565b005b341561015e57600080fd5b610166610599565b60405190815260200160405180910390f35b341561018357600080fd5b61018b610603565b604051901515815260200160405180910390f35b34156101aa57600080fd5b61018b600160a060020a036004358116906024351660443561060c565b604051901515815260200160405180910390f35b34156101e657600080fd5b6101306004351515610935565b005b341561020057600080fd5b610130600435610966565b005b341561021857600080fd5b610130600435600160a060020a03602435166109c1565b005b341561023c57600080fd5b61018b600160a060020a0360043516602435610a24565b604051901515815260200160405180910390f35b341561027257600080fd5b610166600435610a44565b60405190815260200160405180910390f35b341561029a57600080fd5b610130600435600160a060020a0360243516610a67565b005b34156102be57600080fd5b6101306fffffffffffffffffffffffffffffffff60043516610acd565b005b34156102e857600080fd5b610166600160a060020a0360043516610aec565b60405190815260200160405180910390f35b341561031957600080fd5b610130600160a060020a03600435166fffffffffffffffffffffffffffffffff60243516610acd565b005b341561034f57600080fd5b610357610b89565b604051600160a060020a03909116815260200160405180910390f35b341561037e57600080fd5b610130600160a060020a0360043516610b98565b005b341561039f57600080fd5b610357610c16565b604051600160a060020a03909116815260200160405180910390f35b34156103ce57600080fd5b61018b600160a060020a0360043581169060243516604435610c25565b604051901515815260200160405180910390f35b341561040a57600080fd5b610357610e5d565b604051600160a060020a03909116815260200160405180910390f35b341561043957600080fd5b610166600160a060020a0360043581169060243516610e6c565b60405190815260200160405180910390f35b341561047057600080fd5b61018b600160a060020a0360043581169060243516604435610ef2565b604051901515815260200160405180910390f35b34156104ac57600080fd5b610357610f91565b604051600160a060020a03909116815260200160405180910390f35b6104e633600035600160e060020a031916610fa0565b15156104ee57fe5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a25b5b50565b61056433600035600160e060020a031916610fa0565b151561056c57fe5b6003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b5b50565b600254600090600160a060020a031663047fc9aa82604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156105e357600080fd5b6102c65a03f115156105f457600080fd5b50505060405180519150505b90565b60065460ff1681565b60035460009033600160a060020a0390811691161461062757fe5b6006548490849060ff16806106495750600154600160a060020a038381169116145b806106615750600154600160a060020a038281169116145b8061067157506106718282611103565b5b151561067d57600080fd5b600254600160a060020a031663e32d5cf8878761070a8463a32ce11e848460006040516020015260405160e060020a63ffffffff8516028152600160a060020a03928316600482015291166024820152604401602060405180830381600087803b15156106e957600080fd5b6102c65a03f115156106fa57600080fd5b50505060405180519050896111b6565b60405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401600060405180830381600087803b151561075957600080fd5b6102c65a03f1151561076a57600080fd5b5050600254600160a060020a031690506328e69b16876107f2836327e235e38360006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b15156107d157600080fd5b6102c65a03f115156107e257600080fd5b50505060405180519050886111b6565b60405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401600060405180830381600087803b151561083557600080fd5b6102c65a03f1151561084657600080fd5b5050600254600160a060020a031690506328e69b16866108ce836327e235e38360006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b15156108ad57600080fd5b6102c65a03f115156108be57600080fd5b50505060405180519050886111cd565b60405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401600060405180830381600087803b151561091157600080fd5b6102c65a03f1151561092257600080fd5b505050600192505b5b50505b9392505050565b61094b33600035600160e060020a031916610fa0565b151561095357fe5b6006805460ff19168215151790555b5b50565b61097c33600035600160e060020a031916610fa0565b151561098457fe5b61098d816111e4565b1561099757600080fd5b60048054600181016109a9838261123c565b916000526020600020900160005b50829055505b5b50565b6109d733600035600160e060020a031916610fa0565b15156109df57fe5b6109e8826111e4565b15156109f357600080fd5b600160a060020a03811660009081526005602090815260408083208584529091529020805460ff191690555b5b5050565b600560209081526000928352604080842090915290825290205460ff1681565b6004805482908110610a5257fe5b906000526020600020900160005b5054905081565b610a7d33600035600160e060020a031916610fa0565b1515610a8557fe5b610a8e826111e4565b1515610a9957600080fd5b600160a060020a03811660009081526005602090815260408083208584529091529020805460ff191660011790555b5b5050565b60035433600160a060020a03908116911614610ae557fe5bfe5b5b5b50565b600254600090600160a060020a03166327e235e383836040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b1515610b4757600080fd5b6102c65a03f11515610b5857600080fd5b50505060405180519150505b919050565b60035433600160a060020a03908116911614610ae557fe5bfe5b5b5b5050565b600254600160a060020a031681565b610bae33600035600160e060020a031916610fa0565b1515610bb657fe5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167f1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada460405160405180910390a25b5b50565b600154600160a060020a031681565b60035460009033600160a060020a03908116911614610c4057fe5b6006548490849060ff1680610c625750600154600160a060020a038381169116145b80610c7a5750600154600160a060020a038281169116145b80610c8a5750610c8a8282611103565b5b1515610c9657600080fd5b600254600160a060020a03166328e69b16876107f2836327e235e38360006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b15156107d157600080fd5b6102c65a03f115156107e257600080fd5b50505060405180519050886111b6565b60405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401600060405180830381600087803b151561083557600080fd5b6102c65a03f1151561084657600080fd5b5050600254600160a060020a031690506328e69b16866108ce836327e235e38360006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b15156108ad57600080fd5b6102c65a03f115156108be57600080fd5b50505060405180519050886111cd565b60405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401600060405180830381600087803b151561091157600080fd5b6102c65a03f1151561092257600080fd5b505050600192505b5b50505b9392505050565b600054600160a060020a031681565b600254600090600160a060020a031663a32ce11e8484846040516020015260405160e060020a63ffffffff8516028152600160a060020a03928316600482015291166024820152604401602060405180830381600087803b1515610ecf57600080fd5b6102c65a03f11515610ee057600080fd5b50505060405180519150505b92915050565b60035460009033600160a060020a03908116911614610f0d57fe5b600254600160a060020a031663e32d5cf885858560405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401600060405180830381600087803b1515610f7057600080fd5b6102c65a03f11515610f8157600080fd5b505050600190505b5b9392505050565b600354600160a060020a031681565b600030600160a060020a031683600160a060020a03161415610fc457506001610eec565b600154600160a060020a038481169116148015610fea5750600054600160a060020a0316155b15610ff757506001610eec565b600054600160a060020a0316151561105f577feb91385a0d70e44b915093d9ddfe6c8b41f2c56729431b1405f40aee5d874be78383604051600160a060020a039092168252600160e060020a03191660208201526040908101905180910390a1506000610eec565b60008054600160a060020a03169063b7009613908590309086906040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152600160e060020a03199091166044820152606401602060405180830381600087803b15156110d857600080fd5b6102c65a03f115156110e957600080fd5b505050604051805190509050610eec565b5b5b5b92915050565b600080805b60045460ff831610156111a9576004805460ff841690811061112657fe5b906000526020600020900160005b5054600160a060020a038616600090815260056020908152604080832084845290915290205490915060ff16801561118f5750600160a060020a038416600090815260056020908152604080832084845290915290205460ff165b1561119d57600192506111ae565b5b600190910190611108565b600092505b505092915050565b80820382811115610eec57600080fd5b5b92915050565b80820182811015610eec57600080fd5b5b92915050565b6000805b60045460ff82161015611231576004805484919060ff841690811061120957fe5b906000526020600020900160005b505414156112285760019150611236565b5b6001016111e8565b600091505b50919050565b81548183558181151161126057600083815260209020611260918101908301611266565b5b505050565b61060091905b80821115611280576000815560010161126c565b5090565b905600a165627a7a723058207ed7bef76d024d6306568d4d72a82273b86607f822c5ea818e0f2cf7707580f200296060604052341561000f57600080fd5b6040516060806103b68339810160405280805191906020018051919060200180519150505b60038054600160a060020a03808616600160a060020a03199283161790925560008481556004805485851693169290921791829055911681526001602052604090208290555b5050505b6103298061008d6000396000f3006060604052361561008b5763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663047fc9aa811461009057806327e235e3146100b557806328e69b16146100e65780633a3d523f1461010a5780633b4c4b251461012b5780638da5cb5b14610143578063a32ce11e14610172578063e32d5cf8146101a9575b600080fd5b341561009b57600080fd5b6100a36101d3565b60405190815260200160405180910390f35b34156100c057600080fd5b6100a3600160a060020a03600435166101d9565b60405190815260200160405180910390f35b34156100f157600080fd5b610108600160a060020a03600435166024356101eb565b005b341561011557600080fd5b610108600160a060020a0360043516610224565b005b341561013657600080fd5b610108600435610268565b005b341561014e57600080fd5b61015661028a565b604051600160a060020a03909116815260200160405180910390f35b341561017d57600080fd5b6100a3600160a060020a0360043581169060243516610299565b60405190815260200160405180910390f35b34156101b457600080fd5b610108600160a060020a03600435811690602435166044356102b6565b005b60005481565b60016020526000908152604090205481565b60035433600160a060020a0390811691161461020357fe5b600160a060020a03821660009081526001602052604090208190555b5b5050565b60045433600160a060020a0390811691161461023c57fe5b6003805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b50565b60035433600160a060020a0390811691161461028057fe5b60008190555b5b50565b600454600160a060020a031681565b600260209081526000928352604080842090915290825290205481565b60035433600160a060020a039081169116146102ce57fe5b600160a060020a0380841660009081526002602090815260408083209386168352929052208190555b5b5050505600a165627a7a72305820c5749d991a4fd582a2f94abde69da9ff2275f11721179661d6cf288b5111c0220029"`

// DeploySweetTokenLogic deploys a new Ethereum contract, binding an instance of SweetTokenLogic to it.
func DeploySweetTokenLogic(auth *bind.TransactOpts, backend bind.ContractBackend, token_ common.Address, data_ common.Address, supply_ *big.Int) (common.Address, *types.Transaction, *SweetTokenLogic, error) {
	parsed, err := abi.JSON(strings.NewReader(SweetTokenLogicABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SweetTokenLogicBin), backend, token_, data_, supply_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SweetTokenLogic{SweetTokenLogicCaller: SweetTokenLogicCaller{contract: contract}, SweetTokenLogicTransactor: SweetTokenLogicTransactor{contract: contract}}, nil
}

// SweetTokenLogic is an auto generated Go binding around an Ethereum contract.
type SweetTokenLogic struct {
	SweetTokenLogicCaller     // Read-only binding to the contract
	SweetTokenLogicTransactor // Write-only binding to the contract
}

// SweetTokenLogicCaller is an auto generated read-only Go binding around an Ethereum contract.
type SweetTokenLogicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SweetTokenLogicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SweetTokenLogicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SweetTokenLogicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SweetTokenLogicSession struct {
	Contract     *SweetTokenLogic  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SweetTokenLogicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SweetTokenLogicCallerSession struct {
	Contract *SweetTokenLogicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// SweetTokenLogicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SweetTokenLogicTransactorSession struct {
	Contract     *SweetTokenLogicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// SweetTokenLogicRaw is an auto generated low-level Go binding around an Ethereum contract.
type SweetTokenLogicRaw struct {
	Contract *SweetTokenLogic // Generic contract binding to access the raw methods on
}

// SweetTokenLogicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SweetTokenLogicCallerRaw struct {
	Contract *SweetTokenLogicCaller // Generic read-only contract binding to access the raw methods on
}

// SweetTokenLogicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SweetTokenLogicTransactorRaw struct {
	Contract *SweetTokenLogicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSweetTokenLogic creates a new instance of SweetTokenLogic, bound to a specific deployed contract.
func NewSweetTokenLogic(address common.Address, backend bind.ContractBackend) (*SweetTokenLogic, error) {
	contract, err := bindSweetTokenLogic(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SweetTokenLogic{SweetTokenLogicCaller: SweetTokenLogicCaller{contract: contract}, SweetTokenLogicTransactor: SweetTokenLogicTransactor{contract: contract}}, nil
}

// NewSweetTokenLogicCaller creates a new read-only instance of SweetTokenLogic, bound to a specific deployed contract.
func NewSweetTokenLogicCaller(address common.Address, caller bind.ContractCaller) (*SweetTokenLogicCaller, error) {
	contract, err := bindSweetTokenLogic(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &SweetTokenLogicCaller{contract: contract}, nil
}

// NewSweetTokenLogicTransactor creates a new write-only instance of SweetTokenLogic, bound to a specific deployed contract.
func NewSweetTokenLogicTransactor(address common.Address, transactor bind.ContractTransactor) (*SweetTokenLogicTransactor, error) {
	contract, err := bindSweetTokenLogic(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &SweetTokenLogicTransactor{contract: contract}, nil
}

// bindSweetTokenLogic binds a generic wrapper to an already deployed contract.
func bindSweetTokenLogic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SweetTokenLogicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SweetTokenLogic *SweetTokenLogicRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SweetTokenLogic.Contract.SweetTokenLogicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SweetTokenLogic *SweetTokenLogicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.SweetTokenLogicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SweetTokenLogic *SweetTokenLogicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.SweetTokenLogicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SweetTokenLogic *SweetTokenLogicCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SweetTokenLogic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SweetTokenLogic *SweetTokenLogicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SweetTokenLogic *SweetTokenLogicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(src address, spender address) constant returns(uint256)
func (_SweetTokenLogic *SweetTokenLogicCaller) Allowance(opts *bind.CallOpts, src common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SweetTokenLogic.contract.Call(opts, out, "allowance", src, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(src address, spender address) constant returns(uint256)
func (_SweetTokenLogic *SweetTokenLogicSession) Allowance(src common.Address, spender common.Address) (*big.Int, error) {
	return _SweetTokenLogic.Contract.Allowance(&_SweetTokenLogic.CallOpts, src, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(src address, spender address) constant returns(uint256)
func (_SweetTokenLogic *SweetTokenLogicCallerSession) Allowance(src common.Address, spender common.Address) (*big.Int, error) {
	return _SweetTokenLogic.Contract.Allowance(&_SweetTokenLogic.CallOpts, src, spender)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_SweetTokenLogic *SweetTokenLogicCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SweetTokenLogic.contract.Call(opts, out, "authority")
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_SweetTokenLogic *SweetTokenLogicSession) Authority() (common.Address, error) {
	return _SweetTokenLogic.Contract.Authority(&_SweetTokenLogic.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_SweetTokenLogic *SweetTokenLogicCallerSession) Authority() (common.Address, error) {
	return _SweetTokenLogic.Contract.Authority(&_SweetTokenLogic.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(src address) constant returns(uint256)
func (_SweetTokenLogic *SweetTokenLogicCaller) BalanceOf(opts *bind.CallOpts, src common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SweetTokenLogic.contract.Call(opts, out, "balanceOf", src)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(src address) constant returns(uint256)
func (_SweetTokenLogic *SweetTokenLogicSession) BalanceOf(src common.Address) (*big.Int, error) {
	return _SweetTokenLogic.Contract.BalanceOf(&_SweetTokenLogic.CallOpts, src)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(src address) constant returns(uint256)
func (_SweetTokenLogic *SweetTokenLogicCallerSession) BalanceOf(src common.Address) (*big.Int, error) {
	return _SweetTokenLogic.Contract.BalanceOf(&_SweetTokenLogic.CallOpts, src)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(address)
func (_SweetTokenLogic *SweetTokenLogicCaller) Data(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SweetTokenLogic.contract.Call(opts, out, "data")
	return *ret0, err
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(address)
func (_SweetTokenLogic *SweetTokenLogicSession) Data() (common.Address, error) {
	return _SweetTokenLogic.Contract.Data(&_SweetTokenLogic.CallOpts)
}

// Data is a free data retrieval call binding the contract method 0x73d4a13a.
//
// Solidity: function data() constant returns(address)
func (_SweetTokenLogic *SweetTokenLogicCallerSession) Data() (common.Address, error) {
	return _SweetTokenLogic.Contract.Data(&_SweetTokenLogic.CallOpts)
}

// FreeTransfer is a free data retrieval call binding the contract method 0x211ed6c1.
//
// Solidity: function freeTransfer() constant returns(bool)
func (_SweetTokenLogic *SweetTokenLogicCaller) FreeTransfer(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SweetTokenLogic.contract.Call(opts, out, "freeTransfer")
	return *ret0, err
}

// FreeTransfer is a free data retrieval call binding the contract method 0x211ed6c1.
//
// Solidity: function freeTransfer() constant returns(bool)
func (_SweetTokenLogic *SweetTokenLogicSession) FreeTransfer() (bool, error) {
	return _SweetTokenLogic.Contract.FreeTransfer(&_SweetTokenLogic.CallOpts)
}

// FreeTransfer is a free data retrieval call binding the contract method 0x211ed6c1.
//
// Solidity: function freeTransfer() constant returns(bool)
func (_SweetTokenLogic *SweetTokenLogicCallerSession) FreeTransfer() (bool, error) {
	return _SweetTokenLogic.Contract.FreeTransfer(&_SweetTokenLogic.CallOpts)
}

// ListNames is a free data retrieval call binding the contract method 0x66272623.
//
// Solidity: function listNames( uint256) constant returns(bytes32)
func (_SweetTokenLogic *SweetTokenLogicCaller) ListNames(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _SweetTokenLogic.contract.Call(opts, out, "listNames", arg0)
	return *ret0, err
}

// ListNames is a free data retrieval call binding the contract method 0x66272623.
//
// Solidity: function listNames( uint256) constant returns(bytes32)
func (_SweetTokenLogic *SweetTokenLogicSession) ListNames(arg0 *big.Int) ([32]byte, error) {
	return _SweetTokenLogic.Contract.ListNames(&_SweetTokenLogic.CallOpts, arg0)
}

// ListNames is a free data retrieval call binding the contract method 0x66272623.
//
// Solidity: function listNames( uint256) constant returns(bytes32)
func (_SweetTokenLogic *SweetTokenLogicCallerSession) ListNames(arg0 *big.Int) ([32]byte, error) {
	return _SweetTokenLogic.Contract.ListNames(&_SweetTokenLogic.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SweetTokenLogic *SweetTokenLogicCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SweetTokenLogic.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SweetTokenLogic *SweetTokenLogicSession) Owner() (common.Address, error) {
	return _SweetTokenLogic.Contract.Owner(&_SweetTokenLogic.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_SweetTokenLogic *SweetTokenLogicCallerSession) Owner() (common.Address, error) {
	return _SweetTokenLogic.Contract.Owner(&_SweetTokenLogic.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_SweetTokenLogic *SweetTokenLogicCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _SweetTokenLogic.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_SweetTokenLogic *SweetTokenLogicSession) Token() (common.Address, error) {
	return _SweetTokenLogic.Contract.Token(&_SweetTokenLogic.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() constant returns(address)
func (_SweetTokenLogic *SweetTokenLogicCallerSession) Token() (common.Address, error) {
	return _SweetTokenLogic.Contract.Token(&_SweetTokenLogic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SweetTokenLogic *SweetTokenLogicCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _SweetTokenLogic.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SweetTokenLogic *SweetTokenLogicSession) TotalSupply() (*big.Int, error) {
	return _SweetTokenLogic.Contract.TotalSupply(&_SweetTokenLogic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_SweetTokenLogic *SweetTokenLogicCallerSession) TotalSupply() (*big.Int, error) {
	return _SweetTokenLogic.Contract.TotalSupply(&_SweetTokenLogic.CallOpts)
}

// WhiteLists is a free data retrieval call binding the contract method 0x4b4d0d45.
//
// Solidity: function whiteLists( address,  bytes32) constant returns(bool)
func (_SweetTokenLogic *SweetTokenLogicCaller) WhiteLists(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _SweetTokenLogic.contract.Call(opts, out, "whiteLists", arg0, arg1)
	return *ret0, err
}

// WhiteLists is a free data retrieval call binding the contract method 0x4b4d0d45.
//
// Solidity: function whiteLists( address,  bytes32) constant returns(bool)
func (_SweetTokenLogic *SweetTokenLogicSession) WhiteLists(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _SweetTokenLogic.Contract.WhiteLists(&_SweetTokenLogic.CallOpts, arg0, arg1)
}

// WhiteLists is a free data retrieval call binding the contract method 0x4b4d0d45.
//
// Solidity: function whiteLists( address,  bytes32) constant returns(bool)
func (_SweetTokenLogic *SweetTokenLogicCallerSession) WhiteLists(arg0 common.Address, arg1 [32]byte) (bool, error) {
	return _SweetTokenLogic.Contract.WhiteLists(&_SweetTokenLogic.CallOpts, arg0, arg1)
}

// AddToWhiteList is a paid mutator transaction binding the contract method 0x693382a9.
//
// Solidity: function addToWhiteList(listName bytes32, guy address) returns()
func (_SweetTokenLogic *SweetTokenLogicTransactor) AddToWhiteList(opts *bind.TransactOpts, listName [32]byte, guy common.Address) (*types.Transaction, error) {
	return _SweetTokenLogic.contract.Transact(opts, "addToWhiteList", listName, guy)
}

// AddToWhiteList is a paid mutator transaction binding the contract method 0x693382a9.
//
// Solidity: function addToWhiteList(listName bytes32, guy address) returns()
func (_SweetTokenLogic *SweetTokenLogicSession) AddToWhiteList(listName [32]byte, guy common.Address) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.AddToWhiteList(&_SweetTokenLogic.TransactOpts, listName, guy)
}

// AddToWhiteList is a paid mutator transaction binding the contract method 0x693382a9.
//
// Solidity: function addToWhiteList(listName bytes32, guy address) returns()
func (_SweetTokenLogic *SweetTokenLogicTransactorSession) AddToWhiteList(listName [32]byte, guy common.Address) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.AddToWhiteList(&_SweetTokenLogic.TransactOpts, listName, guy)
}

// AddWhiteList is a paid mutator transaction binding the contract method 0x3bd04d69.
//
// Solidity: function addWhiteList(listName bytes32) returns()
func (_SweetTokenLogic *SweetTokenLogicTransactor) AddWhiteList(opts *bind.TransactOpts, listName [32]byte) (*types.Transaction, error) {
	return _SweetTokenLogic.contract.Transact(opts, "addWhiteList", listName)
}

// AddWhiteList is a paid mutator transaction binding the contract method 0x3bd04d69.
//
// Solidity: function addWhiteList(listName bytes32) returns()
func (_SweetTokenLogic *SweetTokenLogicSession) AddWhiteList(listName [32]byte) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.AddWhiteList(&_SweetTokenLogic.TransactOpts, listName)
}

// AddWhiteList is a paid mutator transaction binding the contract method 0x3bd04d69.
//
// Solidity: function addWhiteList(listName bytes32) returns()
func (_SweetTokenLogic *SweetTokenLogicTransactorSession) AddWhiteList(listName [32]byte) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.AddWhiteList(&_SweetTokenLogic.TransactOpts, listName)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(src address, guy address, wad uint256) returns(bool)
func (_SweetTokenLogic *SweetTokenLogicTransactor) Approve(opts *bind.TransactOpts, src common.Address, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetTokenLogic.contract.Transact(opts, "approve", src, guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(src address, guy address, wad uint256) returns(bool)
func (_SweetTokenLogic *SweetTokenLogicSession) Approve(src common.Address, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.Approve(&_SweetTokenLogic.TransactOpts, src, guy, wad)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(src address, guy address, wad uint256) returns(bool)
func (_SweetTokenLogic *SweetTokenLogicTransactorSession) Approve(src common.Address, guy common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.Approve(&_SweetTokenLogic.TransactOpts, src, guy, wad)
}

// Burn is a paid mutator transaction binding the contract method 0x7261e469.
//
// Solidity: function burn( address,  uint128) returns()
func (_SweetTokenLogic *SweetTokenLogicTransactor) Burn(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _SweetTokenLogic.contract.Transact(opts, "burn", arg0, arg1)
}

// Burn is a paid mutator transaction binding the contract method 0x7261e469.
//
// Solidity: function burn( address,  uint128) returns()
func (_SweetTokenLogic *SweetTokenLogicSession) Burn(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.Burn(&_SweetTokenLogic.TransactOpts, arg0, arg1)
}

// Burn is a paid mutator transaction binding the contract method 0x7261e469.
//
// Solidity: function burn( address,  uint128) returns()
func (_SweetTokenLogic *SweetTokenLogicTransactorSession) Burn(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.Burn(&_SweetTokenLogic.TransactOpts, arg0, arg1)
}

// Mint is a paid mutator transaction binding the contract method 0x69d3e20e.
//
// Solidity: function mint( uint128) returns()
func (_SweetTokenLogic *SweetTokenLogicTransactor) Mint(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _SweetTokenLogic.contract.Transact(opts, "mint", arg0)
}

// Mint is a paid mutator transaction binding the contract method 0x69d3e20e.
//
// Solidity: function mint( uint128) returns()
func (_SweetTokenLogic *SweetTokenLogicSession) Mint(arg0 *big.Int) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.Mint(&_SweetTokenLogic.TransactOpts, arg0)
}

// Mint is a paid mutator transaction binding the contract method 0x69d3e20e.
//
// Solidity: function mint( uint128) returns()
func (_SweetTokenLogic *SweetTokenLogicTransactorSession) Mint(arg0 *big.Int) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.Mint(&_SweetTokenLogic.TransactOpts, arg0)
}

// RemoveFromWhiteList is a paid mutator transaction binding the contract method 0x3f32aa70.
//
// Solidity: function removeFromWhiteList(listName bytes32, guy address) returns()
func (_SweetTokenLogic *SweetTokenLogicTransactor) RemoveFromWhiteList(opts *bind.TransactOpts, listName [32]byte, guy common.Address) (*types.Transaction, error) {
	return _SweetTokenLogic.contract.Transact(opts, "removeFromWhiteList", listName, guy)
}

// RemoveFromWhiteList is a paid mutator transaction binding the contract method 0x3f32aa70.
//
// Solidity: function removeFromWhiteList(listName bytes32, guy address) returns()
func (_SweetTokenLogic *SweetTokenLogicSession) RemoveFromWhiteList(listName [32]byte, guy common.Address) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.RemoveFromWhiteList(&_SweetTokenLogic.TransactOpts, listName, guy)
}

// RemoveFromWhiteList is a paid mutator transaction binding the contract method 0x3f32aa70.
//
// Solidity: function removeFromWhiteList(listName bytes32, guy address) returns()
func (_SweetTokenLogic *SweetTokenLogicTransactorSession) RemoveFromWhiteList(listName [32]byte, guy common.Address) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.RemoveFromWhiteList(&_SweetTokenLogic.TransactOpts, listName, guy)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_SweetTokenLogic *SweetTokenLogicTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _SweetTokenLogic.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_SweetTokenLogic *SweetTokenLogicSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.SetAuthority(&_SweetTokenLogic.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_SweetTokenLogic *SweetTokenLogicTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.SetAuthority(&_SweetTokenLogic.TransactOpts, authority_)
}

// SetFreeTransfer is a paid mutator transaction binding the contract method 0x2abfaf1f.
//
// Solidity: function setFreeTransfer(isFree bool) returns()
func (_SweetTokenLogic *SweetTokenLogicTransactor) SetFreeTransfer(opts *bind.TransactOpts, isFree bool) (*types.Transaction, error) {
	return _SweetTokenLogic.contract.Transact(opts, "setFreeTransfer", isFree)
}

// SetFreeTransfer is a paid mutator transaction binding the contract method 0x2abfaf1f.
//
// Solidity: function setFreeTransfer(isFree bool) returns()
func (_SweetTokenLogic *SweetTokenLogicSession) SetFreeTransfer(isFree bool) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.SetFreeTransfer(&_SweetTokenLogic.TransactOpts, isFree)
}

// SetFreeTransfer is a paid mutator transaction binding the contract method 0x2abfaf1f.
//
// Solidity: function setFreeTransfer(isFree bool) returns()
func (_SweetTokenLogic *SweetTokenLogicTransactorSession) SetFreeTransfer(isFree bool) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.SetFreeTransfer(&_SweetTokenLogic.TransactOpts, isFree)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_SweetTokenLogic *SweetTokenLogicTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _SweetTokenLogic.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_SweetTokenLogic *SweetTokenLogicSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.SetOwner(&_SweetTokenLogic.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_SweetTokenLogic *SweetTokenLogicTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.SetOwner(&_SweetTokenLogic.TransactOpts, owner_)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(token_ address) returns()
func (_SweetTokenLogic *SweetTokenLogicTransactor) SetToken(opts *bind.TransactOpts, token_ common.Address) (*types.Transaction, error) {
	return _SweetTokenLogic.contract.Transact(opts, "setToken", token_)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(token_ address) returns()
func (_SweetTokenLogic *SweetTokenLogicSession) SetToken(token_ common.Address) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.SetToken(&_SweetTokenLogic.TransactOpts, token_)
}

// SetToken is a paid mutator transaction binding the contract method 0x144fa6d7.
//
// Solidity: function setToken(token_ address) returns()
func (_SweetTokenLogic *SweetTokenLogicTransactorSession) SetToken(token_ common.Address) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.SetToken(&_SweetTokenLogic.TransactOpts, token_)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(src address, dst address, wad uint256) returns(bool)
func (_SweetTokenLogic *SweetTokenLogicTransactor) Transfer(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetTokenLogic.contract.Transact(opts, "transfer", src, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(src address, dst address, wad uint256) returns(bool)
func (_SweetTokenLogic *SweetTokenLogicSession) Transfer(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.Transfer(&_SweetTokenLogic.TransactOpts, src, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(src address, dst address, wad uint256) returns(bool)
func (_SweetTokenLogic *SweetTokenLogicTransactorSession) Transfer(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.Transfer(&_SweetTokenLogic.TransactOpts, src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_SweetTokenLogic *SweetTokenLogicTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetTokenLogic.contract.Transact(opts, "transferFrom", src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_SweetTokenLogic *SweetTokenLogicSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.TransferFrom(&_SweetTokenLogic.TransactOpts, src, dst, wad)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(src address, dst address, wad uint256) returns(bool)
func (_SweetTokenLogic *SweetTokenLogicTransactorSession) TransferFrom(src common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _SweetTokenLogic.Contract.TransferFrom(&_SweetTokenLogic.TransactOpts, src, dst, wad)
}

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

// VaultABI is the input ABI used to generate the binding from.
const VaultABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"brgAmount\",\"type\":\"uint128\"},{\"name\":\"uouIndex\",\"type\":\"uint256\"}],\"name\":\"repayUou\",\"outputs\":[{\"name\":\"\",\"type\":\"uint128\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uouIndex\",\"type\":\"uint256\"}],\"name\":\"rejectUouRequest\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"brgBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"amountDue\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"uouCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isEmpty\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"addAsset\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"cleanStorage\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uous\",\"outputs\":[{\"name\":\"initialAmount\",\"type\":\"uint128\"},{\"name\":\"repaidAmount\",\"type\":\"uint128\"},{\"name\":\"fee\",\"type\":\"uint128\"},{\"name\":\"time\",\"type\":\"uint256\"},{\"name\":\"decision\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"brgAmount\",\"type\":\"uint128\"}],\"name\":\"requestUou\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assets\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"}],\"name\":\"rmAsset\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uouIndex\",\"type\":\"uint256\"}],\"name\":\"acceptUouRequest\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"wallet_\",\"type\":\"address\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"brgAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"uouIndex\",\"type\":\"uint256\"}],\"name\":\"UouRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"brgAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"uouIndex\",\"type\":\"uint256\"}],\"name\":\"UouRequestApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"brgAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"uouIndex\",\"type\":\"uint256\"}],\"name\":\"UouRequestDeclined\",\"type\":\"event\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"foo\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"bar\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fax\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"UnauthorizedAccess\",\"type\":\"event\"}]"

// VaultBin is the compiled bytecode used for deploying new contracts.
const VaultBin = `"0x6060604052341561000f57600080fd5b604051602080611870833981016040528080519150505b5b60018054600160a060020a03191633600160a060020a03169081179091557fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a25b60048054600160a060020a031916600160a060020a03838116919091179182905516638da5cb5b6000604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15156100e957600080fd5b6102c65a03f115156100fa57600080fd5b505050604051805160018054600160a060020a031916600160a060020a0392909216919091179055505b505b61173b806101356000396000f300606060405236156101225763ffffffff60e060020a60003504166307da68f5811461012757806313af40351461013c5780631a93ebf01461015d57806322d8fe401461019b57806324101d40146101b35780632f195680146101d8578063521eb273146101fd57806356e596a81461022c5780635e7624241461025f578063681fe70c1461028457806370a08231146102ab578063712e6279146102dc57806375f12b211461030f5780637a9e5e4b146103365780637bb98a68146103575780637e84d36b146103be5780638da5cb5b146103d35780639913314114610402578063a52d7ffb14610466578063be9a655514610487578063bf7e214f1461049c578063cf35bdd0146104cb578063d1da3e7b146104fd578063dc2628fa14610524575b600080fd5b341561013257600080fd5b61013a61053c565b005b341561014757600080fd5b61013a600160a060020a03600435166105ec565b005b341561016857600080fd5b61017f6001608060020a036004351660243561066a565b6040516001608060020a03909116815260200160405180910390f35b34156101a657600080fd5b61013a600435610813565b005b34156101be57600080fd5b6101c6610918565b60405190815260200160405180910390f35b34156101e357600080fd5b6101c6610a53565b60405190815260200160405180910390f35b341561020857600080fd5b610210610a59565b604051600160a060020a03909116815260200160405180910390f35b341561023757600080fd5b61013a600160a060020a03600435811690602435166001608060020a0360443516610a68565b005b341561026a57600080fd5b6101c6610b0e565b60405190815260200160405180910390f35b341561028f57600080fd5b610297610b15565b604051901515815260200160405180910390f35b34156102b657600080fd5b6101c6600160a060020a0360043516610b21565b60405190815260200160405180910390f35b34156102e757600080fd5b61013a600160a060020a03600435811690602435166001608060020a0360443516610b9c565b005b341561031a57600080fd5b610297610c4d565b604051901515815260200160405180910390f35b341561034157600080fd5b61013a600160a060020a0360043516610c6e565b005b341561036257600080fd5b61036a610cec565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156103aa5780820151818401525b602001610391565b505050509050019250505060405180910390f35b34156103c957600080fd5b61013a610d04565b005b34156103de57600080fd5b610210610d8a565b604051600160a060020a03909116815260200160405180910390f35b341561040d57600080fd5b610418600435610d99565b6040516001608060020a038087168252858116602083015284166040820152606081018390526080810182600281111561044e57fe5b60ff1681526020019550505050505060405180910390f35b341561047157600080fd5b61013a6001608060020a0360043516610dee565b005b341561049257600080fd5b61013a610f55565b005b34156104a757600080fd5b610210610fee565b604051600160a060020a03909116815260200160405180910390f35b34156104d657600080fd5b610210600435610ffd565b604051600160a060020a03909116815260200160405180910390f35b341561050857600080fd5b61013a600160a060020a036004358116906024351661102f565b005b341561052f57600080fd5b61013a6004356110cf565b005b61055233600035600160e060020a03191661137b565b151561055a57fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a46001805474ff00000000000000000000000000000000000000001916740100000000000000000000000000000000000000001790555b5b50505b565b61060233600035600160e060020a03191661137b565b151561060a57fe5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a25b5b50565b60045460009081908190600160a060020a031663c41c2f2482604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156106b857600080fd5b6102c65a03f115156106c957600080fd5b50505060405180519050600160a060020a0316634f9c8fe86000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561071857600080fd5b6102c65a03f1151561072957600080fd5b50505060405180519050600160a060020a031633600160a060020a031614151561075257600080fd5b600380548590811061076057fe5b906000526020600020906004020160005b50915060025b600383015460ff16600281111561078a57fe5b1461079457600080fd5b81546107b3906001608060020a0380821691608060020a9004166114d9565b9050806001608060020a0316856001608060020a031611156107d3578094505b81546107ef90608060020a90046001608060020a0316866114f9565b82546001608060020a03918216608060020a0291161782558492505b505092915050565b60005b600380548390811061082457fe5b906000526020600020906004020160005b506003015460ff16600281111561084857fe5b1461085257600080fd5b600160038281548110151561086357fe5b906000526020600020906004020160005b50600301805460ff1916600183600281111561088c57fe5b02179055507f3ed382a18ac4f16fc64863bd31865023dddcfb68ad39027d444747a2d9a3d40b306003838154811015156108c257fe5b906000526020600020906004020160005b50546001608060020a031683604051600160a060020a0390931683526001608060020a0390911660208301526040808301919091526060909101905180910390a15b50565b600454600090600160a060020a031663c41c2f2482604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561096257600080fd5b6102c65a03f1151561097357600080fd5b50505060405180519050600160a060020a0316634f9c8fe86000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156109c257600080fd5b6102c65a03f115156109d357600080fd5b50505060405180519050600160a060020a03166370a082313060006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b1515610a3357600080fd5b6102c65a03f11515610a4457600080fd5b50505060405180519150505b90565b60055481565b600454600160a060020a031681565b610a7e33600035600160e060020a03191661137b565b1515610a8657fe5b82600160a060020a031663a9059cbb838360006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526001608060020a03166024820152604401602060405180830381600087803b1515610aec57600080fd5b6102c65a03f11515610afd57600080fd5b505050604051805150505b5b505050565b6003545b90565b60035460009011155b90565b600081600160a060020a03166370a082313060006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b1515610b7a57600080fd5b6102c65a03f11515610b8b57600080fd5b50505060405180519150505b919050565b610bb233600035600160e060020a03191661137b565b1515610bba57fe5b73__AssetsLib_____________________________6352d28bc8600285858560405160e060020a63ffffffff87160281526004810194909452600160a060020a039283166024850152911660448301526001608060020a0316606482015260840160006040518083038186803b1515610c3257600080fd5b6102c65a03f41515610c4357600080fd5b5050505b5b505050565b60015474010000000000000000000000000000000000000000900460ff1681565b610c8433600035600160e060020a03191661137b565b1515610c8c57fe5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167f1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada460405160405180910390a25b5b50565b610cf461164c565b610cfe6002611519565b90505b90565b610d1a33600035600160e060020a03191661137b565b1515610d2257fe5b73__AssetsLib_____________________________6376c86cd7600260405160e060020a63ffffffff8416028152600481019190915260240160006040518083038186803b1515610d7257600080fd5b6102c65a03f41515610b0857600080fd5b5050505b5b565b600154600160a060020a031681565b6003805482908110610da757fe5b906000526020600020906004020160005b508054600182015460028301546003909301546001608060020a038084169550608060020a909304831693919092169160ff1685565b610df661165e565b610e0c33600035600160e060020a03191661137b565b1515610e1457fe5b6001608060020a03821681524260608201526003805460018101610e38838261168f565b916000526020600020906004020160005b508290815181546fffffffffffffffffffffffffffffffff19166001608060020a0391909116178155602082015181546001608060020a03918216608060020a02911617815560408201516001820180546fffffffffffffffffffffffffffffffff19166001608060020a039290921691909117905560608201518160020155608082015160038201805460ff19166001836002811115610ee657fe5b02179055505050507f67b3996492d0393b3c7c2e35a1d67de008977701daa8880bf9c0f91aadf5529a3083600160038054905003604051600160a060020a0390931683526001608060020a0390911660208301526040808301919091526060909101905180910390a15b5b5050565b610f6b33600035600160e060020a03191661137b565b1515610f7357fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a46001805474ff0000000000000000000000000000000000000000191690555b5b50505b565b600054600160a060020a031681565b600280548290811061100b57fe5b906000526020600020900160005b915054906101000a9004600160a060020a031681565b61104533600035600160e060020a03191661137b565b151561104d57fe5b73__AssetsLib_____________________________638a7fcb2d6002848460405160e060020a63ffffffff86160281526004810193909352600160a060020a03918216602484015216604482015260640160006040518083038186803b15156110b557600080fd5b6102c65a03f415156110c657600080fd5b5050505b5b5050565b60005b60038054839081106110e057fe5b906000526020600020906004020160005b506003015460ff16600281111561110457fe5b1461110e57600080fd5b600260038281548110151561111f57fe5b906000526020600020906004020160005b50600301805460ff1916600183600281111561114857fe5b021790555061118660038281548110151561115f57fe5b906000526020600020906004020160005b50546005546001608060020a0390911690611638565b600555600454600160a060020a031663c41c2f246000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156111d157600080fd5b6102c65a03f115156111e257600080fd5b50505060405180519050600160a060020a0316634f9c8fe86000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561123157600080fd5b6102c65a03f1151561124257600080fd5b50505060405180519050600160a060020a031663591a91c560038381548110151561126957fe5b906000526020600020906004020160005b50546001546001608060020a0390911690600160a060020a031660405160e060020a63ffffffff85160281526001608060020a039092166004830152600160a060020a03166024820152604401600060405180830381600087803b15156112e057600080fd5b6102c65a03f115156112f157600080fd5b5050507f65025c7799d1ede98b1ef493282c86b38a8552da89e43f27da441fb2019096f3306003838154811015156108c257fe5b906000526020600020906004020160005b50546001608060020a031683604051600160a060020a0390931683526001608060020a0390911660208301526040808301919091526060909101905180910390a15b50565b600030600160a060020a031683600160a060020a0316141561139f575060016114d0565b600154600160a060020a0384811691161480156113c55750600054600160a060020a0316155b156113d2575060016114d0565b600054600160a060020a0316151561143a577feb91385a0d70e44b915093d9ddfe6c8b41f2c56729431b1405f40aee5d874be78383604051600160a060020a039092168252600160e060020a03191660208201526040908101905180910390a15060006114d0565b60008054600160a060020a03169063b7009613908590309086906040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152600160e060020a03199091166044820152606401602060405180830381600087803b15156114b357600080fd5b6102c65a03f115156114c457600080fd5b50505060405180519150505b5b5b5b92915050565b8082036001608060020a0380841690821611156114d057fe5b5b92915050565b8082016001608060020a0380841690821610156114d057fe5b5b92915050565b61152161164c565b61152961164c565b825460009060405180591061153b5750595b908082528060200260200182016040525b509150600090505b835463ffffffff8216101561162d57838163ffffffff1681548110151561157757fe5b906000526020600020900160005b9054906101000a9004600160a060020a0316600160a060020a03166370a082313060006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b15156115ed57600080fd5b6102c65a03f115156115fe57600080fd5b50505060405180519050828263ffffffff168151811061161a57fe5b602090810290910101525b600101611554565b8192505b5050919050565b808201828110156114d057fe5b5b92915050565b60206040519081016040526000815290565b60a060405190810160409081526000808352602083018190529082018190526060820181905260808201905b905290565b815481835581811511610b0857600402816004028360005260206000209182019101610b0891906116c1565b5b505050565b610a5091905b808211156117085760008082556001820180546fffffffffffffffffffffffffffffffff19169055600282015560038101805460ff191690556004016116c7565b5090565b905600a165627a7a72305820b2e43e3d5721687cf2be25e791aecceda11992c28ba88399ee1d73c63f0d617f0029"`

// DeployVault deploys a new Ethereum contract, binding an instance of Vault to it.
func DeployVault(auth *bind.TransactOpts, backend bind.ContractBackend, wallet_ common.Address) (common.Address, *types.Transaction, *Vault, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(VaultBin), backend, wallet_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Vault{VaultCaller: VaultCaller{contract: contract}, VaultTransactor: VaultTransactor{contract: contract}}, nil
}

// Vault is an auto generated Go binding around an Ethereum contract.
type Vault struct {
	VaultCaller     // Read-only binding to the contract
	VaultTransactor // Write-only binding to the contract
}

// VaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultSession struct {
	Contract     *Vault            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultCallerSession struct {
	Contract *VaultCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultTransactorSession struct {
	Contract     *VaultTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultRaw struct {
	Contract *Vault // Generic contract binding to access the raw methods on
}

// VaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultCallerRaw struct {
	Contract *VaultCaller // Generic read-only contract binding to access the raw methods on
}

// VaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultTransactorRaw struct {
	Contract *VaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVault creates a new instance of Vault, bound to a specific deployed contract.
func NewVault(address common.Address, backend bind.ContractBackend) (*Vault, error) {
	contract, err := bindVault(address, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vault{VaultCaller: VaultCaller{contract: contract}, VaultTransactor: VaultTransactor{contract: contract}}, nil
}

// NewVaultCaller creates a new read-only instance of Vault, bound to a specific deployed contract.
func NewVaultCaller(address common.Address, caller bind.ContractCaller) (*VaultCaller, error) {
	contract, err := bindVault(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &VaultCaller{contract: contract}, nil
}

// NewVaultTransactor creates a new write-only instance of Vault, bound to a specific deployed contract.
func NewVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultTransactor, error) {
	contract, err := bindVault(address, nil, transactor)
	if err != nil {
		return nil, err
	}
	return &VaultTransactor{contract: contract}, nil
}

// bindVault binds a generic wrapper to an already deployed contract.
func bindVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VaultABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vault *VaultRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Vault.Contract.VaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vault *VaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.Contract.VaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vault *VaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vault.Contract.VaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vault *VaultCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Vault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vault *VaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vault *VaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vault.Contract.contract.Transact(opts, method, params...)
}

// AmountDue is a free data retrieval call binding the contract method 0x2f195680.
//
// Solidity: function amountDue() constant returns(uint256)
func (_Vault *VaultCaller) AmountDue(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "amountDue")
	return *ret0, err
}

// AmountDue is a free data retrieval call binding the contract method 0x2f195680.
//
// Solidity: function amountDue() constant returns(uint256)
func (_Vault *VaultSession) AmountDue() (*big.Int, error) {
	return _Vault.Contract.AmountDue(&_Vault.CallOpts)
}

// AmountDue is a free data retrieval call binding the contract method 0x2f195680.
//
// Solidity: function amountDue() constant returns(uint256)
func (_Vault *VaultCallerSession) AmountDue() (*big.Int, error) {
	return _Vault.Contract.AmountDue(&_Vault.CallOpts)
}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets( uint256) constant returns(address)
func (_Vault *VaultCaller) Assets(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "assets", arg0)
	return *ret0, err
}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets( uint256) constant returns(address)
func (_Vault *VaultSession) Assets(arg0 *big.Int) (common.Address, error) {
	return _Vault.Contract.Assets(&_Vault.CallOpts, arg0)
}

// Assets is a free data retrieval call binding the contract method 0xcf35bdd0.
//
// Solidity: function assets( uint256) constant returns(address)
func (_Vault *VaultCallerSession) Assets(arg0 *big.Int) (common.Address, error) {
	return _Vault.Contract.Assets(&_Vault.CallOpts, arg0)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Vault *VaultCaller) Authority(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "authority")
	return *ret0, err
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Vault *VaultSession) Authority() (common.Address, error) {
	return _Vault.Contract.Authority(&_Vault.CallOpts)
}

// Authority is a free data retrieval call binding the contract method 0xbf7e214f.
//
// Solidity: function authority() constant returns(address)
func (_Vault *VaultCallerSession) Authority() (common.Address, error) {
	return _Vault.Contract.Authority(&_Vault.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(token address) constant returns(uint256)
func (_Vault *VaultCaller) BalanceOf(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "balanceOf", token)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(token address) constant returns(uint256)
func (_Vault *VaultSession) BalanceOf(token common.Address) (*big.Int, error) {
	return _Vault.Contract.BalanceOf(&_Vault.CallOpts, token)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(token address) constant returns(uint256)
func (_Vault *VaultCallerSession) BalanceOf(token common.Address) (*big.Int, error) {
	return _Vault.Contract.BalanceOf(&_Vault.CallOpts, token)
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(uint256[])
func (_Vault *VaultCaller) Balances(opts *bind.CallOpts) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "balances")
	return *ret0, err
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(uint256[])
func (_Vault *VaultSession) Balances() ([]*big.Int, error) {
	return _Vault.Contract.Balances(&_Vault.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(uint256[])
func (_Vault *VaultCallerSession) Balances() ([]*big.Int, error) {
	return _Vault.Contract.Balances(&_Vault.CallOpts)
}

// BrgBalance is a free data retrieval call binding the contract method 0x24101d40.
//
// Solidity: function brgBalance() constant returns(uint256)
func (_Vault *VaultCaller) BrgBalance(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "brgBalance")
	return *ret0, err
}

// BrgBalance is a free data retrieval call binding the contract method 0x24101d40.
//
// Solidity: function brgBalance() constant returns(uint256)
func (_Vault *VaultSession) BrgBalance() (*big.Int, error) {
	return _Vault.Contract.BrgBalance(&_Vault.CallOpts)
}

// BrgBalance is a free data retrieval call binding the contract method 0x24101d40.
//
// Solidity: function brgBalance() constant returns(uint256)
func (_Vault *VaultCallerSession) BrgBalance() (*big.Int, error) {
	return _Vault.Contract.BrgBalance(&_Vault.CallOpts)
}

// IsEmpty is a free data retrieval call binding the contract method 0x681fe70c.
//
// Solidity: function isEmpty() constant returns(bool)
func (_Vault *VaultCaller) IsEmpty(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "isEmpty")
	return *ret0, err
}

// IsEmpty is a free data retrieval call binding the contract method 0x681fe70c.
//
// Solidity: function isEmpty() constant returns(bool)
func (_Vault *VaultSession) IsEmpty() (bool, error) {
	return _Vault.Contract.IsEmpty(&_Vault.CallOpts)
}

// IsEmpty is a free data retrieval call binding the contract method 0x681fe70c.
//
// Solidity: function isEmpty() constant returns(bool)
func (_Vault *VaultCallerSession) IsEmpty() (bool, error) {
	return _Vault.Contract.IsEmpty(&_Vault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Vault *VaultCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Vault *VaultSession) Owner() (common.Address, error) {
	return _Vault.Contract.Owner(&_Vault.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Vault *VaultCallerSession) Owner() (common.Address, error) {
	return _Vault.Contract.Owner(&_Vault.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_Vault *VaultCaller) Stopped(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "stopped")
	return *ret0, err
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_Vault *VaultSession) Stopped() (bool, error) {
	return _Vault.Contract.Stopped(&_Vault.CallOpts)
}

// Stopped is a free data retrieval call binding the contract method 0x75f12b21.
//
// Solidity: function stopped() constant returns(bool)
func (_Vault *VaultCallerSession) Stopped() (bool, error) {
	return _Vault.Contract.Stopped(&_Vault.CallOpts)
}

// UouCount is a free data retrieval call binding the contract method 0x5e762424.
//
// Solidity: function uouCount() constant returns(uint256)
func (_Vault *VaultCaller) UouCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "uouCount")
	return *ret0, err
}

// UouCount is a free data retrieval call binding the contract method 0x5e762424.
//
// Solidity: function uouCount() constant returns(uint256)
func (_Vault *VaultSession) UouCount() (*big.Int, error) {
	return _Vault.Contract.UouCount(&_Vault.CallOpts)
}

// UouCount is a free data retrieval call binding the contract method 0x5e762424.
//
// Solidity: function uouCount() constant returns(uint256)
func (_Vault *VaultCallerSession) UouCount() (*big.Int, error) {
	return _Vault.Contract.UouCount(&_Vault.CallOpts)
}

// Uous is a free data retrieval call binding the contract method 0x99133141.
//
// Solidity: function uous( uint256) constant returns(initialAmount uint128, repaidAmount uint128, fee uint128, time uint256, decision uint8)
func (_Vault *VaultCaller) Uous(opts *bind.CallOpts, arg0 *big.Int) (struct {
	InitialAmount *big.Int
	RepaidAmount  *big.Int
	Fee           *big.Int
	Time          *big.Int
	Decision      uint8
}, error) {
	ret := new(struct {
		InitialAmount *big.Int
		RepaidAmount  *big.Int
		Fee           *big.Int
		Time          *big.Int
		Decision      uint8
	})
	out := ret
	err := _Vault.contract.Call(opts, out, "uous", arg0)
	return *ret, err
}

// Uous is a free data retrieval call binding the contract method 0x99133141.
//
// Solidity: function uous( uint256) constant returns(initialAmount uint128, repaidAmount uint128, fee uint128, time uint256, decision uint8)
func (_Vault *VaultSession) Uous(arg0 *big.Int) (struct {
	InitialAmount *big.Int
	RepaidAmount  *big.Int
	Fee           *big.Int
	Time          *big.Int
	Decision      uint8
}, error) {
	return _Vault.Contract.Uous(&_Vault.CallOpts, arg0)
}

// Uous is a free data retrieval call binding the contract method 0x99133141.
//
// Solidity: function uous( uint256) constant returns(initialAmount uint128, repaidAmount uint128, fee uint128, time uint256, decision uint8)
func (_Vault *VaultCallerSession) Uous(arg0 *big.Int) (struct {
	InitialAmount *big.Int
	RepaidAmount  *big.Int
	Fee           *big.Int
	Time          *big.Int
	Decision      uint8
}, error) {
	return _Vault.Contract.Uous(&_Vault.CallOpts, arg0)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_Vault *VaultCaller) Wallet(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "wallet")
	return *ret0, err
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_Vault *VaultSession) Wallet() (common.Address, error) {
	return _Vault.Contract.Wallet(&_Vault.CallOpts)
}

// Wallet is a free data retrieval call binding the contract method 0x521eb273.
//
// Solidity: function wallet() constant returns(address)
func (_Vault *VaultCallerSession) Wallet() (common.Address, error) {
	return _Vault.Contract.Wallet(&_Vault.CallOpts)
}

// AcceptUouRequest is a paid mutator transaction binding the contract method 0xdc2628fa.
//
// Solidity: function acceptUouRequest(uouIndex uint256) returns()
func (_Vault *VaultTransactor) AcceptUouRequest(opts *bind.TransactOpts, uouIndex *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "acceptUouRequest", uouIndex)
}

// AcceptUouRequest is a paid mutator transaction binding the contract method 0xdc2628fa.
//
// Solidity: function acceptUouRequest(uouIndex uint256) returns()
func (_Vault *VaultSession) AcceptUouRequest(uouIndex *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.AcceptUouRequest(&_Vault.TransactOpts, uouIndex)
}

// AcceptUouRequest is a paid mutator transaction binding the contract method 0xdc2628fa.
//
// Solidity: function acceptUouRequest(uouIndex uint256) returns()
func (_Vault *VaultTransactorSession) AcceptUouRequest(uouIndex *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.AcceptUouRequest(&_Vault.TransactOpts, uouIndex)
}

// AddAsset is a paid mutator transaction binding the contract method 0x712e6279.
//
// Solidity: function addAsset(token address, src address, wad uint128) returns()
func (_Vault *VaultTransactor) AddAsset(opts *bind.TransactOpts, token common.Address, src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "addAsset", token, src, wad)
}

// AddAsset is a paid mutator transaction binding the contract method 0x712e6279.
//
// Solidity: function addAsset(token address, src address, wad uint128) returns()
func (_Vault *VaultSession) AddAsset(token common.Address, src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.AddAsset(&_Vault.TransactOpts, token, src, wad)
}

// AddAsset is a paid mutator transaction binding the contract method 0x712e6279.
//
// Solidity: function addAsset(token address, src address, wad uint128) returns()
func (_Vault *VaultTransactorSession) AddAsset(token common.Address, src common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.AddAsset(&_Vault.TransactOpts, token, src, wad)
}

// CleanStorage is a paid mutator transaction binding the contract method 0x7e84d36b.
//
// Solidity: function cleanStorage() returns()
func (_Vault *VaultTransactor) CleanStorage(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "cleanStorage")
}

// CleanStorage is a paid mutator transaction binding the contract method 0x7e84d36b.
//
// Solidity: function cleanStorage() returns()
func (_Vault *VaultSession) CleanStorage() (*types.Transaction, error) {
	return _Vault.Contract.CleanStorage(&_Vault.TransactOpts)
}

// CleanStorage is a paid mutator transaction binding the contract method 0x7e84d36b.
//
// Solidity: function cleanStorage() returns()
func (_Vault *VaultTransactorSession) CleanStorage() (*types.Transaction, error) {
	return _Vault.Contract.CleanStorage(&_Vault.TransactOpts)
}

// RejectUouRequest is a paid mutator transaction binding the contract method 0x22d8fe40.
//
// Solidity: function rejectUouRequest(uouIndex uint256) returns()
func (_Vault *VaultTransactor) RejectUouRequest(opts *bind.TransactOpts, uouIndex *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "rejectUouRequest", uouIndex)
}

// RejectUouRequest is a paid mutator transaction binding the contract method 0x22d8fe40.
//
// Solidity: function rejectUouRequest(uouIndex uint256) returns()
func (_Vault *VaultSession) RejectUouRequest(uouIndex *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.RejectUouRequest(&_Vault.TransactOpts, uouIndex)
}

// RejectUouRequest is a paid mutator transaction binding the contract method 0x22d8fe40.
//
// Solidity: function rejectUouRequest(uouIndex uint256) returns()
func (_Vault *VaultTransactorSession) RejectUouRequest(uouIndex *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.RejectUouRequest(&_Vault.TransactOpts, uouIndex)
}

// RepayUou is a paid mutator transaction binding the contract method 0x1a93ebf0.
//
// Solidity: function repayUou(brgAmount uint128, uouIndex uint256) returns(uint128)
func (_Vault *VaultTransactor) RepayUou(opts *bind.TransactOpts, brgAmount *big.Int, uouIndex *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "repayUou", brgAmount, uouIndex)
}

// RepayUou is a paid mutator transaction binding the contract method 0x1a93ebf0.
//
// Solidity: function repayUou(brgAmount uint128, uouIndex uint256) returns(uint128)
func (_Vault *VaultSession) RepayUou(brgAmount *big.Int, uouIndex *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.RepayUou(&_Vault.TransactOpts, brgAmount, uouIndex)
}

// RepayUou is a paid mutator transaction binding the contract method 0x1a93ebf0.
//
// Solidity: function repayUou(brgAmount uint128, uouIndex uint256) returns(uint128)
func (_Vault *VaultTransactorSession) RepayUou(brgAmount *big.Int, uouIndex *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.RepayUou(&_Vault.TransactOpts, brgAmount, uouIndex)
}

// RequestUou is a paid mutator transaction binding the contract method 0xa52d7ffb.
//
// Solidity: function requestUou(brgAmount uint128) returns()
func (_Vault *VaultTransactor) RequestUou(opts *bind.TransactOpts, brgAmount *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "requestUou", brgAmount)
}

// RequestUou is a paid mutator transaction binding the contract method 0xa52d7ffb.
//
// Solidity: function requestUou(brgAmount uint128) returns()
func (_Vault *VaultSession) RequestUou(brgAmount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.RequestUou(&_Vault.TransactOpts, brgAmount)
}

// RequestUou is a paid mutator transaction binding the contract method 0xa52d7ffb.
//
// Solidity: function requestUou(brgAmount uint128) returns()
func (_Vault *VaultTransactorSession) RequestUou(brgAmount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.RequestUou(&_Vault.TransactOpts, brgAmount)
}

// RmAsset is a paid mutator transaction binding the contract method 0xd1da3e7b.
//
// Solidity: function rmAsset(token address, dst address) returns()
func (_Vault *VaultTransactor) RmAsset(opts *bind.TransactOpts, token common.Address, dst common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "rmAsset", token, dst)
}

// RmAsset is a paid mutator transaction binding the contract method 0xd1da3e7b.
//
// Solidity: function rmAsset(token address, dst address) returns()
func (_Vault *VaultSession) RmAsset(token common.Address, dst common.Address) (*types.Transaction, error) {
	return _Vault.Contract.RmAsset(&_Vault.TransactOpts, token, dst)
}

// RmAsset is a paid mutator transaction binding the contract method 0xd1da3e7b.
//
// Solidity: function rmAsset(token address, dst address) returns()
func (_Vault *VaultTransactorSession) RmAsset(token common.Address, dst common.Address) (*types.Transaction, error) {
	return _Vault.Contract.RmAsset(&_Vault.TransactOpts, token, dst)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_Vault *VaultTransactor) SetAuthority(opts *bind.TransactOpts, authority_ common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setAuthority", authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_Vault *VaultSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetAuthority(&_Vault.TransactOpts, authority_)
}

// SetAuthority is a paid mutator transaction binding the contract method 0x7a9e5e4b.
//
// Solidity: function setAuthority(authority_ address) returns()
func (_Vault *VaultTransactorSession) SetAuthority(authority_ common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetAuthority(&_Vault.TransactOpts, authority_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_Vault *VaultTransactor) SetOwner(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "setOwner", owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_Vault *VaultSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetOwner(&_Vault.TransactOpts, owner_)
}

// SetOwner is a paid mutator transaction binding the contract method 0x13af4035.
//
// Solidity: function setOwner(owner_ address) returns()
func (_Vault *VaultTransactorSession) SetOwner(owner_ common.Address) (*types.Transaction, error) {
	return _Vault.Contract.SetOwner(&_Vault.TransactOpts, owner_)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Vault *VaultTransactor) Start(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "start")
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Vault *VaultSession) Start() (*types.Transaction, error) {
	return _Vault.Contract.Start(&_Vault.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0xbe9a6555.
//
// Solidity: function start() returns()
func (_Vault *VaultTransactorSession) Start() (*types.Transaction, error) {
	return _Vault.Contract.Start(&_Vault.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Vault *VaultTransactor) Stop(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "stop")
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Vault *VaultSession) Stop() (*types.Transaction, error) {
	return _Vault.Contract.Stop(&_Vault.TransactOpts)
}

// Stop is a paid mutator transaction binding the contract method 0x07da68f5.
//
// Solidity: function stop() returns()
func (_Vault *VaultTransactorSession) Stop() (*types.Transaction, error) {
	return _Vault.Contract.Stop(&_Vault.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0x56e596a8.
//
// Solidity: function transfer(token address, dst address, wad uint128) returns()
func (_Vault *VaultTransactor) Transfer(opts *bind.TransactOpts, token common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "transfer", token, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0x56e596a8.
//
// Solidity: function transfer(token address, dst address, wad uint128) returns()
func (_Vault *VaultSession) Transfer(token common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Transfer(&_Vault.TransactOpts, token, dst, wad)
}

// Transfer is a paid mutator transaction binding the contract method 0x56e596a8.
//
// Solidity: function transfer(token address, dst address, wad uint128) returns()
func (_Vault *VaultTransactorSession) Transfer(token common.Address, dst common.Address, wad *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Transfer(&_Vault.TransactOpts, token, dst, wad)
}

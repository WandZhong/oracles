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

// VaultABI is the input ABI used to generate the binding from.
const VaultABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"stop\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\"}],\"name\":\"setOwner\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"brgAmount\",\"type\":\"uint128\"},{\"name\":\"uouIndex\",\"type\":\"uint256\"}],\"name\":\"repayUou\",\"outputs\":[{\"name\":\"\",\"type\":\"uint128\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uouIndex\",\"type\":\"uint256\"}],\"name\":\"rejectUouRequest\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"brgBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"amountDue\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wallet\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"uouCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isEmpty\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"src\",\"type\":\"address\"},{\"name\":\"wad\",\"type\":\"uint128\"}],\"name\":\"addAsset\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stopped\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"authority_\",\"type\":\"address\"}],\"name\":\"setAuthority\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"balances\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"},{\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"cleanStorage\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uous\",\"outputs\":[{\"name\":\"initialAmount\",\"type\":\"uint128\"},{\"name\":\"repaidAmount\",\"type\":\"uint128\"},{\"name\":\"fee\",\"type\":\"uint128\"},{\"name\":\"time\",\"type\":\"uint256\"},{\"name\":\"decision\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"assetsLen\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"brgAmount\",\"type\":\"uint128\"}],\"name\":\"requestUou\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"remove\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"start\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"authority\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"assets\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"token\",\"type\":\"address\"},{\"name\":\"dst\",\"type\":\"address\"}],\"name\":\"rmAsset\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"uouIndex\",\"type\":\"uint256\"}],\"name\":\"acceptUouRequest\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hasFunds\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"wallet_\",\"type\":\"address\"}],\"payable\":false,\"type\":\"constructor\"},{\"payable\":true,\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"brgAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"uouIndex\",\"type\":\"uint256\"}],\"name\":\"UouRequested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"brgAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"uouIndex\",\"type\":\"uint256\"}],\"name\":\"UouRequestApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"brgAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"uouIndex\",\"type\":\"uint256\"}],\"name\":\"UouRequestDeclined\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"AssetAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"token\",\"type\":\"address\"}],\"name\":\"AssetRemoved\",\"type\":\"event\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"name\":\"guy\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"foo\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"bar\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"wad\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fax\",\"type\":\"bytes\"}],\"name\":\"LogNote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"authority\",\"type\":\"address\"}],\"name\":\"LogSetAuthority\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LogSetOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"sig\",\"type\":\"bytes4\"}],\"name\":\"UnauthorizedAccess\",\"type\":\"event\"}]"

// VaultBin is the compiled bytecode used for deploying new contracts.
const VaultBin = `"0x6060604052341561000f57600080fd5b604051602080611e03833981016040528080519150505b5b60018054600160a060020a03191633600160a060020a03169081179091557fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a25b60048054600160a060020a031916600160a060020a03838116919091179182905516638da5cb5b6000604051602001526040518163ffffffff167c0100000000000000000000000000000000000000000000000000000000028152600401602060405180830381600087803b15156100e957600080fd5b6102c65a03f115156100fa57600080fd5b505050604051805160018054600160a060020a031916600160a060020a0392831617905560045416905063256b5a02306040517c010000000000000000000000000000000000000000000000000000000063ffffffff8416028152600160a060020a039091166004820152602401600060405180830381600087803b151561018157600080fd5b6102c65a03f1151561019257600080fd5b5050505b505b611c5c806101a76000396000f300606060405236156101435763ffffffff60e060020a60003504166307da68f5811461014757806313af40351461015c5780631a93ebf01461017d57806322d8fe40146101bb57806324101d40146101d35780632f195680146101f8578063521eb2731461021d57806356e596a81461024c5780635e7624241461027f578063681fe70c146102a457806370a08231146102cb578063712e6279146102fc57806375f12b211461032f5780637a9e5e4b146103565780637bb98a68146103775780637e84d36b146104255780638da5cb5b1461043a5780639913314114610469578063a37685e5146104cd578063a52d7ffb146104f2578063a7f4377914610513578063be9a655514610528578063bf7e214f1461053d578063cf35bdd01461056c578063d1da3e7b1461059e578063dc2628fa146105d7578063f2cd59f6146105ef575b5b5b005b341561015257600080fd5b610143610616565b005b341561016757600080fd5b610143600160a060020a03600435166106c6565b005b341561018857600080fd5b61019f6001608060020a0360043516602435610744565b6040516001608060020a03909116815260200160405180910390f35b34156101c657600080fd5b6101436004356108ed565b005b34156101de57600080fd5b6101e66109f2565b60405190815260200160405180910390f35b341561020357600080fd5b6101e6610b2d565b60405190815260200160405180910390f35b341561022857600080fd5b610230610b33565b604051600160a060020a03909116815260200160405180910390f35b341561025757600080fd5b610143600160a060020a03600435811690602435166001608060020a0360443516610b42565b005b341561028a57600080fd5b6101e6610be8565b60405190815260200160405180910390f35b34156102af57600080fd5b6102b7610bef565b604051901515815260200160405180910390f35b34156102d657600080fd5b6101e6600160a060020a0360043516610bfb565b60405190815260200160405180910390f35b341561030757600080fd5b610143600160a060020a03600435811690602435166001608060020a0360443516610c76565b005b341561033a57600080fd5b6102b7610d84565b604051901515815260200160405180910390f35b341561036157600080fd5b610143600160a060020a0360043516610da5565b005b341561038257600080fd5b61038a610e23565b604051808060200180602001838103835285818151815260200191508051906020019060200280838360005b838110156103cf5780820151818401525b6020016103b6565b50505050905001838103825284818151815260200191508051906020019060200280838360005b8381101561040f5780820151818401525b6020016103f6565b5050505090500194505050505060405180910390f35b341561043057600080fd5b610143610e46565b005b341561044557600080fd5b610230610ecc565b604051600160a060020a03909116815260200160405180910390f35b341561047457600080fd5b61047f600435610edb565b6040516001608060020a03808716825285811660208301528416604082015260608101839052608081018260028111156104b557fe5b60ff1681526020019550505050505060405180910390f35b34156104d857600080fd5b6101e6610f30565b60405190815260200160405180910390f35b34156104fd57600080fd5b6101436001608060020a0360043516610fb0565b005b341561051e57600080fd5b6101436110f8565b005b341561053357600080fd5b6101436112ec565b005b341561054857600080fd5b610230611385565b604051600160a060020a03909116815260200160405180910390f35b341561057757600080fd5b610230600435611394565b604051600160a060020a03909116815260200160405180910390f35b34156105a957600080fd5b6102b7600160a060020a03600435811690602435166113c6565b604051901515815260200160405180910390f35b34156105e257600080fd5b6101436004356114c7565b005b34156105fa57600080fd5b6102b761176d565b604051901515815260200160405180910390f35b61062c33600035600160e060020a0319166117e7565b151561063457fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a46001805474ff00000000000000000000000000000000000000001916740100000000000000000000000000000000000000001790555b5b50505b565b6106dc33600035600160e060020a0319166117e7565b15156106e457fe5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167fce241d7ca1f669fee44b6fc00b8eba2df3bb514eed0f6f668f8f89096e81ed9460405160405180910390a25b5b50565b60045460009081908190600160a060020a031663c41c2f2482604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561079257600080fd5b6102c65a03f115156107a357600080fd5b50505060405180519050600160a060020a0316634f9c8fe86000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156107f257600080fd5b6102c65a03f1151561080357600080fd5b50505060405180519050600160a060020a031633600160a060020a031614151561082c57600080fd5b600380548590811061083a57fe5b906000526020600020906004020160005b50915060025b600383015460ff16600281111561086457fe5b1461086e57600080fd5b815461088d906001608060020a0380821691608060020a90041661194a565b9050806001608060020a0316856001608060020a031611156108ad578094505b81546108c990608060020a90046001608060020a03168661196d565b82546001608060020a03918216608060020a0291161782558492505b505092915050565b60005b60038054839081106108fe57fe5b906000526020600020906004020160005b506003015460ff16600281111561092257fe5b1461092c57600080fd5b600160038281548110151561093d57fe5b906000526020600020906004020160005b50600301805460ff1916600183600281111561096657fe5b02179055507f3ed382a18ac4f16fc64863bd31865023dddcfb68ad39027d444747a2d9a3d40b3060038381548110151561099c57fe5b906000526020600020906004020160005b50546001608060020a031683604051600160a060020a0390931683526001608060020a0390911660208301526040808301919091526060909101905180910390a15b50565b600454600090600160a060020a031663c41c2f2482604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610a3c57600080fd5b6102c65a03f11515610a4d57600080fd5b50505060405180519050600160a060020a0316634f9c8fe86000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610a9c57600080fd5b6102c65a03f11515610aad57600080fd5b50505060405180519050600160a060020a03166370a082313060006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b1515610b0d57600080fd5b6102c65a03f11515610b1e57600080fd5b50505060405180519150505b90565b60055481565b600454600160a060020a031681565b610b5833600035600160e060020a0319166117e7565b1515610b6057fe5b82600160a060020a031663a9059cbb838360006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526001608060020a03166024820152604401602060405180830381600087803b1515610bc657600080fd5b6102c65a03f11515610bd757600080fd5b505050604051805150505b5b505050565b6003545b90565b60035460009011155b90565b600081600160a060020a03166370a082313060006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b1515610c5457600080fd5b6102c65a03f11515610c6557600080fd5b50505060405180519150505b919050565b610c8c33600035600160e060020a0319166117e7565b1515610c9457fe5b73__AssetsLib_____________________________6352d28bc8600285858560405160e060020a63ffffffff87160281526004810194909452600160a060020a039283166024850152911660448301526001608060020a0316606482015260840160006040518083038186803b1515610d0c57600080fd5b6102c65a03f41515610d1d57600080fd5b5050507fd0d5f5786af393a6465a13be9c33dbd3ee9c29b1c32872f8aa0cd405ee62f8fe838383604051600160a060020a0393841681529190921660208201526001608060020a039091166040808301919091526060909101905180910390a15b5b505050565b60015474010000000000000000000000000000000000000000900460ff1681565b610dbb33600035600160e060020a0319166117e7565b1515610dc357fe5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a038381169190911791829055167f1abebea81bfa2637f28358c371278fb15ede7ea8dd28d2e03b112ff6d936ada460405160405180910390a25b5b50565b610e2b611b5b565b610e33611b5b565b610e3d6002611990565b915091505b9091565b610e5c33600035600160e060020a0319166117e7565b1515610e6457fe5b73__AssetsLib_____________________________6376c86cd7600260405160e060020a63ffffffff8416028152600481019190915260240160006040518083038186803b1515610eb457600080fd5b6102c65a03f41515610be257600080fd5b5050505b5b565b600154600160a060020a031681565b6003805482908110610ee957fe5b906000526020600020906004020160005b508054600182015460028301546003909301546001608060020a038084169550608060020a909304831693919092169160ff1685565b600073__AssetsLib_____________________________63922ddcf26002836040516020015260405160e060020a63ffffffff8416028152600481019190915260240160206040518083038186803b1515610f8a57600080fd5b6102c65a03f41515610f9b57600080fd5b505050604051805163ffffffff169150505b90565b610fb8611b7f565b6001608060020a03821681524260608201526003805460018101610fdc8382611bb0565b916000526020600020906004020160005b508290815181546fffffffffffffffffffffffffffffffff19166001608060020a0391909116178155602082015181546001608060020a03918216608060020a02911617815560408201516001820180546fffffffffffffffffffffffffffffffff19166001608060020a039290921691909117905560608201518160020155608082015160038201805460ff1916600183600281111561108a57fe5b02179055505050507f67b3996492d0393b3c7c2e35a1d67de008977701daa8880bf9c0f91aadf5529a3083600160038054905003604051600160a060020a0390931683526001608060020a0390911660208301526040808301919091526060909101905180910390a15b5050565b6003546000901561110857600080fd5b5060005b60025463ffffffff82161015611276576002805463ffffffff831690811061113057fe5b906000526020600020900160005b905460045460028054600160a060020a036101009590950a90930484169363a9059cbb9392169163ffffffff861690811061117557fe5b906000526020600020900160005b9054906101000a9004600160a060020a0316600160a060020a03166370a082313060006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b15156111eb57600080fd5b6102c65a03f115156111fc57600080fd5b5050506040518051905060006040516020015260405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401602060405180830381600087803b151561125257600080fd5b6102c65a03f1151561126357600080fd5b505050604051805150505b60010161110c565b600454600160a060020a031663ceb68c233060405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401600060405180830381600087803b15156112c657600080fd5b6102c65a03f115156112d757600080fd5b5050600454600160a060020a03169050ff5b50565b61130233600035600160e060020a0319166117e7565b151561130a57fe5b600435602435808233600160a060020a031660008035600160e060020a0319169034903660405183815260406020820181815290820183905260608201848480828437820191505094505050505060405180910390a46001805474ff0000000000000000000000000000000000000000191690555b5b50505b565b600054600160a060020a031681565b60028054829081106113a257fe5b906000526020600020900160005b915054906101000a9004600160a060020a031681565b60006113de33600035600160e060020a0319166117e7565b15156113e657fe5b73__AssetsLib_____________________________638a7fcb2d6002858560006040516020015260405160e060020a63ffffffff86160281526004810193909352600160a060020a03918216602484015216604482015260640160206040518083038186803b151561145757600080fd5b6102c65a03f4151561146857600080fd5b50505060405180519050156114bc577f37803e2125c48ee96c38ddf04e826daf335b0e1603579040fd275aba6d06b6fc83604051600160a060020a03909116815260200160405180910390a15060016114c0565b5060005b5b92915050565b60005b60038054839081106114d857fe5b906000526020600020906004020160005b506003015460ff1660028111156114fc57fe5b1461150657600080fd5b600260038281548110151561151757fe5b906000526020600020906004020160005b50600301805460ff1916600183600281111561154057fe5b021790555061157e60038281548110151561155757fe5b906000526020600020906004020160005b50546005546001608060020a0390911690611b44565b600555600454600160a060020a031663c41c2f246000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156115c957600080fd5b6102c65a03f115156115da57600080fd5b50505060405180519050600160a060020a0316634f9c8fe86000604051602001526040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561162957600080fd5b6102c65a03f1151561163a57600080fd5b505050604051805160045460038054600160a060020a0393841694506327cfc21993909216918590811061166a57fe5b906000526020600020906004020160005b50546001608060020a031660405160e060020a63ffffffff8516028152600160a060020a0390921660048301526001608060020a03166024820152604401600060405180830381600087803b15156116d257600080fd5b6102c65a03f115156116e357600080fd5b5050507f65025c7799d1ede98b1ef493282c86b38a8552da89e43f27da441fb2019096f33060038381548110151561099c57fe5b906000526020600020906004020160005b50546001608060020a031683604051600160a060020a0390931683526001608060020a0390911660208301526040808301919091526060909101905180910390a15b50565b600073__AssetsLib_____________________________632a6a878f6002836040516020015260405160e060020a63ffffffff8416028152600481019190915260240160206040518083038186803b15156117c757600080fd5b6102c65a03f41515610b1e57600080fd5b50505060405180519150505b90565b600030600160a060020a031683600160a060020a0316141561180b575060016114c0565b600154600160a060020a0384811691161480156118315750600054600160a060020a0316155b1561183e575060016114c0565b600054600160a060020a031615156118a6577feb91385a0d70e44b915093d9ddfe6c8b41f2c56729431b1405f40aee5d874be78383604051600160a060020a039092168252600160e060020a03191660208201526040908101905180910390a15060006114c0565b60008054600160a060020a03169063b7009613908590309086906040516020015260405160e060020a63ffffffff8616028152600160a060020a039384166004820152919092166024820152600160e060020a03199091166044820152606401602060405180830381600087803b151561191f57600080fd5b6102c65a03f1151561193057600080fd5b5050506040518051905090506114c0565b5b5b5b92915050565b8082036001608060020a0380841690821611156114c057600080fd5b5b92915050565b8082016001608060020a0380841690821610156114c057600080fd5b5b92915050565b611998611b5b565b6119a0611b5b565b6119a8611b5b565b6119b0611b5b565b84546000906040518059106119c25750595b908082528060200260200182016040525b5086549093506040518059106119e65750595b908082528060200260200182016040525b509150600090505b855463ffffffff82161015611b3557858163ffffffff16815481101515611a2257fe5b906000526020600020900160005b9054906101000a9004600160a060020a0316838263ffffffff1681518110611a5457fe5b600160a060020a039092166020928302909101909101528554869063ffffffff8316908110611a7f57fe5b906000526020600020900160005b9054906101000a9004600160a060020a0316600160a060020a03166370a082313060006040516020015260405160e060020a63ffffffff8416028152600160a060020a039091166004820152602401602060405180830381600087803b1515611af557600080fd5b6102c65a03f11515611b0657600080fd5b50505060405180519050828263ffffffff1681518110611b2257fe5b602090810290910101525b6001016119ff565b8282945094505b505050915091565b808201828110156114c057600080fd5b5b92915050565b60206040519081016040526000815290565b60206040519081016040526000815290565b60a060405190810160409081526000808352602083018190529082018190526060820181905260808201905b905290565b815481835581811511610be257600402816004028360005260206000209182019101610be29190611be2565b5b505050565b610b2a91905b80821115611c295760008082556001820180546fffffffffffffffffffffffffffffffff19169055600282015560038101805460ff19169055600401611be8565b5090565b905600a165627a7a7230582015bf31556c7b94a9dcb73c575718118cd6967758279ba7bb328787a17d834e1b0029"`

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

// AssetsLen is a free data retrieval call binding the contract method 0xa37685e5.
//
// Solidity: function assetsLen() constant returns(uint256)
func (_Vault *VaultCaller) AssetsLen(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "assetsLen")
	return *ret0, err
}

// AssetsLen is a free data retrieval call binding the contract method 0xa37685e5.
//
// Solidity: function assetsLen() constant returns(uint256)
func (_Vault *VaultSession) AssetsLen() (*big.Int, error) {
	return _Vault.Contract.AssetsLen(&_Vault.CallOpts)
}

// AssetsLen is a free data retrieval call binding the contract method 0xa37685e5.
//
// Solidity: function assetsLen() constant returns(uint256)
func (_Vault *VaultCallerSession) AssetsLen() (*big.Int, error) {
	return _Vault.Contract.AssetsLen(&_Vault.CallOpts)
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
// Solidity: function balances() constant returns(address[], uint256[])
func (_Vault *VaultCaller) Balances(opts *bind.CallOpts) ([]common.Address, []*big.Int, error) {
	var (
		ret0 = new([]common.Address)
		ret1 = new([]*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
	}
	err := _Vault.contract.Call(opts, out, "balances")
	return *ret0, *ret1, err
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(address[], uint256[])
func (_Vault *VaultSession) Balances() ([]common.Address, []*big.Int, error) {
	return _Vault.Contract.Balances(&_Vault.CallOpts)
}

// Balances is a free data retrieval call binding the contract method 0x7bb98a68.
//
// Solidity: function balances() constant returns(address[], uint256[])
func (_Vault *VaultCallerSession) Balances() ([]common.Address, []*big.Int, error) {
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

// HasFunds is a free data retrieval call binding the contract method 0xf2cd59f6.
//
// Solidity: function hasFunds() constant returns(bool)
func (_Vault *VaultCaller) HasFunds(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Vault.contract.Call(opts, out, "hasFunds")
	return *ret0, err
}

// HasFunds is a free data retrieval call binding the contract method 0xf2cd59f6.
//
// Solidity: function hasFunds() constant returns(bool)
func (_Vault *VaultSession) HasFunds() (bool, error) {
	return _Vault.Contract.HasFunds(&_Vault.CallOpts)
}

// HasFunds is a free data retrieval call binding the contract method 0xf2cd59f6.
//
// Solidity: function hasFunds() constant returns(bool)
func (_Vault *VaultCallerSession) HasFunds() (bool, error) {
	return _Vault.Contract.HasFunds(&_Vault.CallOpts)
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

// Remove is a paid mutator transaction binding the contract method 0xa7f43779.
//
// Solidity: function remove() returns()
func (_Vault *VaultTransactor) Remove(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "remove")
}

// Remove is a paid mutator transaction binding the contract method 0xa7f43779.
//
// Solidity: function remove() returns()
func (_Vault *VaultSession) Remove() (*types.Transaction, error) {
	return _Vault.Contract.Remove(&_Vault.TransactOpts)
}

// Remove is a paid mutator transaction binding the contract method 0xa7f43779.
//
// Solidity: function remove() returns()
func (_Vault *VaultTransactorSession) Remove() (*types.Transaction, error) {
	return _Vault.Contract.Remove(&_Vault.TransactOpts)
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
// Solidity: function rmAsset(token address, dst address) returns(bool)
func (_Vault *VaultTransactor) RmAsset(opts *bind.TransactOpts, token common.Address, dst common.Address) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "rmAsset", token, dst)
}

// RmAsset is a paid mutator transaction binding the contract method 0xd1da3e7b.
//
// Solidity: function rmAsset(token address, dst address) returns(bool)
func (_Vault *VaultSession) RmAsset(token common.Address, dst common.Address) (*types.Transaction, error) {
	return _Vault.Contract.RmAsset(&_Vault.TransactOpts, token, dst)
}

// RmAsset is a paid mutator transaction binding the contract method 0xd1da3e7b.
//
// Solidity: function rmAsset(token address, dst address) returns(bool)
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

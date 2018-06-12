// Copyright (c) 2017 Sweetbridge Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package setup

import (
	"flag"
	"math/big"

	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum/wad"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/robert-zaremba/errstack"
	"github.com/robert-zaremba/go-bat"
)

// EthFlags represents common Ethereum client flags
type EthFlags struct {
	PkHex         *string
	PkFile        *string
	PkPwd         *string
	ContractsPath *string
	Network       *string
	NetworkConfig *string
}

// NewEthFlags associate ethereum client flags with the structure fields.
// This should be called before flag.Parse or Flag function.
func NewEthFlags() EthFlags {
	return EthFlags{
		PkHex:         flag.String("pk-hex", "", "private key - takes precedence over `pk-file` + `pk-pwd` [required]"),
		PkFile:        flag.String("pk-file", "", "path to the private key json file [required if `pk-hex` not specified]"),
		PkPwd:         flag.String("pk-pwd", "", "key file password [required if `pk-hex` not specified]"),
		ContractsPath: flag.String("contracts", "", "path to contract schemas directory [required]"),
		Network:       flag.String("eth-network", "", "blockchain network name, must be specified in eth-network-file file [required]"),
		NetworkConfig: flag.String("eth-network-file", "", "path to the file with network configuration [required]"),
	}
}

// Check validates the flags.
func (ef EthFlags) Check() error {
	if *ef.Network == "" || *ef.NetworkConfig == "" {
		if *ef.Network != "" {
			return errstack.NewReq("eth-network-file must be specified")
		}
		if *ef.NetworkConfig != "" {
			return errstack.NewReq("eth-network must be specified")
		}
		return errstack.NewReq("both eth-network and eth-network-file must be specified")
	}
	if *ef.PkHex != "" {
		if _, err := crypto.HexToECDSA(*ef.PkHex); err != nil {
			return errstack.WrapAsReq(err,
				"Can't parse ECDSA key. Expected valid hex string.")
		}
	} else {
		if err := bat.IsFile(*ef.PkFile); err != nil {
			return errstack.WrapAsReq(err, "Wrong pkFile")
		}
	}
	if *ef.ContractsPath == "" {
		return errstack.NewReq("All equired arguments must be specified")
	}
	return nil
}

// mustNewTxrFactory creates TxrFactory based on command flags.
// It panics in case of error.
func (ef EthFlags) mustNewTxrFactory() ethereum.TxrFactory {
	if *ef.PkHex != "" {
		p, err := ethereum.NewPrivKeyTxrFactory(*ef.PkHex)
		if err != nil {
			logger.Fatal("Malformed plain (hex) private key")
		}
		return p
	}
	p, err := ethereum.NewJSONTxrFactory(*ef.PkFile, *ef.PkPwd)
	if err != nil {
		logger.Fatal("Can't create TxrFactory based on JSON file", "filename", *ef.PkFile, err)
	}
	return p
}

// MustEthFactory creates ethclient and contract factory based on flag options
func (ef EthFlags) MustEthFactory() (*ethclient.Client, ethereum.ContractFactory) {
	var netmap = map[string]network{}
	if err := bat.DecodeJSONFile(*ef.NetworkConfig, &netmap, logger); err != nil {
		logger.Fatal("Can't parse eth network config file", "path", *ef.NetworkConfig, err)
	}
	n, ok := netmap[*ef.Network]
	if !ok {
		logger.Fatal("Unknown network name", "network", *ef.Network)
	}
	txrF := ef.mustNewTxrFactory()
	return MustEthFactory(n, *ef.ContractsPath, txrF)
}

// GasPriceFlags defines flag to customize gas price.
type GasPriceFlags struct {
	gasPriceStr *string
	GasPrice    *big.Int
}

// NewGasPriceFlags instantiates GasPriceFlags.
// This should be called before flag.Parse or Flag function.
func NewGasPriceFlags() *GasPriceFlags {
	return &GasPriceFlags{
		gasPriceStr: flag.String("gas-price", "", "gas price in Gwei to be used for single transaction. Must be a positive integer or 'default' (string). 'default' = use gas price oracle. [required]")}
}

// Check validates the flags.
func (gp *GasPriceFlags) Check() error {
	if *gp.gasPriceStr == "default" {
		gp.GasPrice = nil
		return nil
	}
	n, err := bat.Atoui64(*gp.gasPriceStr)
	if err != nil {
		return errstack.WrapAsReq(err, "Wrong gas-price value. Must be positive integer or 'default'")
	} else if n == 0 {
		return errstack.NewReq("Wrong gas-price value. Must be positive integer or 'default' (0 was given).")
	}
	gp.GasPrice = wad.GweiToWei(n)
	return nil
}

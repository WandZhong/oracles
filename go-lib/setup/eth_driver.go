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
	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"bitbucket.org/sweetbridge/oracles/go-lib/utils"
	"github.com/ethereum/go-ethereum/ethclient"
)

type network struct {
	ID   int
	Addr string
}

const testrpcID = 9

// MustEthFactory creates new eth client and ContractFactory based on the network name.
func MustEthFactory(n network, contractsPath string, txrF ethereum.TxrFactory) (
	*ethclient.Client, ethereum.ContractFactory) {

	logger.Debug("Creating Eth Client", "address", n.Addr, "id", n.ID)
	client, err := ethclient.Dial(n.Addr)
	utils.Assert(err, "Can't connect to the node "+n.Addr)
	sf, err := ethereum.NewSchemaFactory(contractsPath, n.ID)
	utils.Assert(err, "wrong contractsPath")
	cf := ethereum.NewContractFactory(client, sf, txrF, n.ID == testrpcID)
	return client, cf
}

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

package ethcrawl

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/robert-zaremba/errstack"
)

// chainHandle is a context for exploring ethereum
type chainHandle struct {
	chainLabel   string
	networkID    int64
	chainAddress string
	client       *ethclient.Client
	ctx          context.Context
}

// newChainHandle return an initialized chainHandle
func newChainHandle(ctx context.Context, chainAddress string, chainLabel string) (*chainHandle, errstack.E) {
	client, err := ethclient.Dial(chainAddress)
	if err != nil {
		return nil, errstack.WrapAsInf(err, "can't connect to Ethereum node")
	}

	netID, err := client.NetworkID(ctx)
	if err != nil {
		return nil, errstack.WrapAsInf(err, "can't get Ethereum network ID")
	}

	// define a default chainLabel if none is provided
	if chainLabel == "" {
		chainLabel = fmt.Sprintf("ETH-%s", netID.String())
	}

	return &chainHandle{
		chainAddress: chainAddress,
		chainLabel:   chainLabel,
		networkID:    netID.Int64(),
		client:       client,
		ctx:          ctx,
	}, nil
}

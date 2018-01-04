// Copyright (c) 2017 Sweetbridge Stiftung (Sweetbridge Foundation)
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

package trancheq

import (
	"math/big"

	contracts "bitbucket.org/sweetbridge/oracles/go-contracts"
	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"bitbucket.org/sweetbridge/oracles/go-lib/liquidity"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/robert-zaremba/errstack"
)

// SWCqueue globals
var (
	SWCqueueABI = ethereum.MustParseABI("SWCqueue", contracts.SWCqueueABI)
)

const (
	logSWCqueueDirectPledge = "LogSWCqueueDirectPledge"
)

func init() {
	ethereum.MustHaveEvents(SWCqueueABI, logSWCqueueDirectPledge)
}

// LogSWCqueueDirectPledge returns event
func LogSWCqueueDirectPledge() abi.Event { return SWCqueueABI.Events[logSWCqueueDirectPledge] }

// EventDirectPledge represents LogSWCqueueDirectPledge event payload
type EventDirectPledge struct {
	Who      common.Address
	Wad      *big.Int
	Currency liquidity.Currency
}

// Unmarshal blockchain log into the event structure
func (e *EventDirectPledge) Unmarshal(log *types.Log) errstack.E {
	curr := [3]byte{}
	dest := &[3]interface{}{&e.Who, &e.Wad, &curr}
	err := ethereum.UnmarshalEvent(dest, log.Data, LogSWCqueueDirectPledge())
	e.Currency = liquidity.Currency(curr[:])
	return err
}

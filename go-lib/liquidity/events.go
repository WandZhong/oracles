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

package liquidity

import (
	"math/big"

	"bitbucket.org/sweetbridge/oracles/go-contracts"
	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/robert-zaremba/errstack"
)

// globals
var (
	BRGusdABI = ethereum.MustParseABI("BRGusd", contracts.BridgeTokenABI)
)

const (
	logTokenTransfer = "Transfer"
)

func init() {
	ethereum.MustHaveEvents(BRGusdABI, logTokenTransfer)
}

// LogBRGusdTransfer returns event
func LogBRGusdTransfer() abi.Event { return BRGusdABI.Events[logTokenTransfer] }

// EventTokenTransfer represents Transfer event payload
type EventTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
}

// Unmarshal blockchain log into the event structure
func (e *EventTokenTransfer) Unmarshal(log *types.Log) errstack.E {
	err := ethereum.UnmarshalEvent(e, log.Data, LogBRGusdTransfer())
	if err != nil {
		return err
	}
	if len(log.Topics) < 3 {
		return errstack.NewInfF("Can't unmarshal EventTokenTransfer event. Not enough elements in Topic. Want at least: 3, got: %d", len(log.Topics))
	}
	if e.From, err = ethereum.HashToAddress(log.Topics[1]); err != nil {
		return err
	}
	e.To, err = ethereum.HashToAddress(log.Topics[2])
	return err
}

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

package payfwd

import (
	"bitbucket.org/sweetbridge/oracles/go-contracts"
	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/robert-zaremba/errstack"
)

// global events
var (
	ForwarderFactoryABI = ethereum.MustParseABI("ForwarderFactory",
		contracts.ForwarderFactoryABI)
)

const (
	logForwarderCreated = "LogForwarderCreated"
)

func init() {
	ethereum.MustHaveEvents(ForwarderFactoryABI, logForwarderCreated)
}

// LogForwarderCreated returns event
func LogForwarderCreated() abi.Event {
	return ForwarderFactoryABI.Events[logForwarderCreated]
}

// EventForwarderCreated represents LogForwarderCreated event payload
type EventForwarderCreated struct {
	Forwarder common.Address
}

// Unmarshal blockchain log into the event structure
func (e *EventForwarderCreated) Unmarshal(log types.Log) errstack.E {
	return ethereum.UnmarshalEvent(&e.Forwarder, log.Data, LogForwarderCreated())
}

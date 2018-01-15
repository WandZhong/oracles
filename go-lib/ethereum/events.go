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

package ethereum

import (
	"context"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/robert-zaremba/errstack"
)

// SubscribeSimple is a simple utility function to create events subscription
func SubscribeSimple(ctx context.Context,
	client *ethclient.Client,
	topics [][]common.Hash, addresses []common.Address) (<-chan types.Log, ethereum.Subscription, errstack.E) {
	query := ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   nil,
		Topics:    topics,
		Addresses: addresses}
	var events = make(chan types.Log, 5) // to quickler consume batch events
	s, err := client.SubscribeFilterLogs(ctx, query, events)
	return events, s, errstack.WrapAsDomain(err, "Can't create Ethereum Subscription")
}

// UnmarshalEvent blockchain log into the event structure
// `dest` must be a pointer to initialized structure
func UnmarshalEvent(dest interface{}, data []byte, e abi.Event) errstack.E {
	a := abi.ABI{Events: map[string]abi.Event{"e": e}}
	err := a.Unpack(dest, "e", data)
	return errstack.WrapAsInf(err,
		"Probably the ABI doesn't match with the contract version")
}

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

package trancheq

import (
	"time"

	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"bitbucket.org/sweetbridge/oracles/go-lib/liquidity"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/robert-zaremba/errstack"
	pgt "github.com/robert-zaremba/go-pgt"
)

// Pledge represents pledge_queue entry
type Pledge struct {
	tableName struct{} `sql:"pledge_queue"`

	Tx         string              `sql:"pledge_queue_id,pk"` // ID is a Tx hash
	WalletAddr ethereum.PgtAddress `sql:"wallet_addr"`
	CtrAddr    ethereum.PgtAddress `sql:"ctr_addr,notnull"`
	Wad        pgt.BigInt          `sql:",notnull"`
	CreatedAt  time.Time           `sql:"created_at"`
	Currency   liquidity.Currency  `sql:",notnull"`
	Direct     bool                `sql:",notnull"`
}

// LogPledger is an interface for an object which is able to construct
// new pledges from a Log
type LogPledger interface {
	NewPledge(log *types.Log) (Pledge, errstack.E)
}

// NewPledgeFromDirectPledge constructs new Pledge from the event
func NewPledgeFromDirectPledge(log *types.Log) (Pledge, errstack.E) {
	var e EventDirectPledge
	if err := e.Unmarshal(log); err != nil {
		return Pledge{}, err
	}
	return Pledge{
		Tx:         log.TxHash.Hex(),
		WalletAddr: ethereum.PgtAddress{Address: e.Who},
		CtrAddr:    ethereum.PgtAddress{Address: log.Address},
		Wad:        pgt.BigInt{Int: e.Wad},
		Currency:   e.Currency,
		CreatedAt:  time.Now(),
		Direct:     true,
	}, nil
}

type transferPledger struct {
	brgs map[common.Address]liquidity.Currency
}

// NewLogTransferPledger constructs new Pledge from a Transfer event. It uses
// mapping for various BRG* flavours handling different currencies
func NewLogTransferPledger(brgs map[common.Address]liquidity.Currency) LogPledger {
	return transferPledger{brgs}
}

func (tp transferPledger) NewPledge(log *types.Log) (Pledge, errstack.E) {
	var e liquidity.EventTokenTransfer
	if err := e.Unmarshal(log); err != nil {
		return Pledge{}, err
	}
	curr, ok := tp.brgs[log.Address]
	if !ok {
		return Pledge{}, errstack.NewDomain(
			"Receiving transfer from unknown contract: " + log.Address.Hex())
	}
	return Pledge{
		Tx:         log.TxHash.Hex(),
		WalletAddr: ethereum.PgtAddress{Address: e.From},
		CtrAddr:    ethereum.PgtAddress{Address: e.To},
		Wad:        pgt.BigInt{Int: e.Value},
		Currency:   curr,
		CreatedAt:  time.Now(),
		Direct:     false,
	}, nil
}

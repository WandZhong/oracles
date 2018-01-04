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

package crawlers

import (
	"math/big"
	"time"
)

// BCBlock is a chain block Data structure
type BCBlock struct {
	*BCBlockHeader
	Transactions []*BCTrans
	Logs         []*BCLog
}

// BCBlockHeader contains Block's header details
type BCBlockHeader struct {
	BlockNum     int64     `json:"blockNum"`
	BlockHash    string    `json:"blockHash"`
	Timestamp    time.Time `json:"timestamp"`
	NetworkID    int64     `json:"networkId"`
	NetworkLabel string    `json:"networkLabel"`
}

// BCLog contains Chain log event data
type BCLog struct {
	ID        string   `json:"uid"`
	TxIndex   int64    `json:"txIndex"`
	TxHash    string   `json:"txHash"`
	LogIndex  uint     `json:"logIndex"`
	Address   string   `json:"address"`
	EventName string   `json:"eventName"`
	EventAgs  []string `json:"eventArgs"`
	Data      string   `json:"data"`
}

// BCTrans contains Chain transaction data
type BCTrans struct {
	ID                     string     `json:"uid"`
	TxIndex                int64      `json:"txIndex"`
	TxHash                 string     `json:"txHash"`
	GasPrice               uint64     `json:"gasPrice"`
	Gas                    uint64     `json:"gas"`
	CumulativeGasUsed      uint64     `json:"cumulativeGasUsed"`
	GasUsed                uint64     `json:"gasUsed"`
	From                   string     `json:"from"`
	To                     string     `json:"to,omitempty"`
	ContractCreated        bool       `json:"contractCreation,omitempty"`
	CreatedContractAddress string     `json:"createdContractAddress,omitempty"`
	Value                  *big.Float `json:"value"`
	RawValue               *big.Int   `json:"rawValue"`
	Data                   string     `json:"data"`
}

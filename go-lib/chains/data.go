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

package chains

import "time"

// BlockData is a chain block Data structure
type BlockData struct {
	Header       *HeaderData
	Transactions []*TransData
	Logs         []*LogData
}

// HeaderData contains Block's header details
type HeaderData struct {
	BlockNum     int64     `json:"blockNum"`
	BlockHash    string    `json:"blockHash"`
	Timestamp    time.Time `json:"timestamp"`
	NetworkID    int64     `json:"networkId"`
	NetworkLabel string    `json:"networkLabel"`
}

// LogData contains Chain log event data
type LogData struct {
	ID       string `json:"uid"`
	TxIndex  int64  `json:"txIndex"`
	TxHash   string `json:"txHash"`
	LogIndex uint   `json:"logIndex"`
	Address  string `json:"address"`
	Data     string `json:"data"`
}

// TransData contains Chain transaction data
type TransData struct {
	ID                     string `json:"uid"`
	TxIndex                int64  `json:"txIndex"`
	TxHash                 string `json:"txHash"`
	GasPrice               int64  `json:"gasPrice"`
	Gas                    int64  `json:"gas"`
	CumulativeGasUsed      int64  `json:"cumulativeGasUsed"`
	GasUsed                int64  `json:"gasUsed"`
	From                   string `json:"from"`
	To                     string `json:"to,omitempty"`
	ContractCreated        bool   `json:"contractCreation,omitempty"`
	CreatedContractAddress string `json:"createdContractAddress,omitempty"`
	Value                  int64  `json:"value"`
	Data                   string `json:"data"`
}

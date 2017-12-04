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

package ethcrawl

import (
	"bitbucket.org/sweetbridge/oracles/go-lib/chains"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/robert-zaremba/errstack"
)

// A log reader
type logReader struct {
	handle *ChainHandle
}

func newLogReader(handle *ChainHandle) *logReader {
	return &logReader{handle}
}

// reads the list of logs from an ethereum block.
func (r *logReader) readList(block *types.Block, filter chains.AddressFilter) ([]*chains.LogData, errstack.E) {
	// Filter on all the logs of the provided block
	logFilter := ethereum.FilterQuery{
		FromBlock: block.Number(),
		ToBlock:   block.Number(),
		Addresses: nil,
		Topics:    nil,
	}

	logs, err := r.handle.client.FilterLogs(r.handle.ctx, logFilter)
	if err != nil {
		return nil, errstack.NewDomain(err.Error())
	}
	var lst []*chains.LogData
	for _, log := range logs {
		data, err := r.read(&log, filter)
		if err != nil {
			return nil, err
		}
		lst = append(lst, data)
	}
	return lst, nil
}

// Returns a LogData object if it matches the provided filter
func (r *logReader) read(log *types.Log, filter chains.AddressFilter) (*chains.LogData, errstack.E) {
	addr := log.Address.String()
	if filter != nil && filter.MatchesNone(addr) {
		return nil, nil
	}

	data := &chains.LogData{
		LogIndex: log.Index,
		TxHash:   log.TxHash.String(),
		TxIndex:  int64(log.TxIndex),
		Data:     hex.EncodeToString(log.Data),
		Address:  addr,
	}
	data.ID = fmt.Sprintf("%s-%d", data.TxHash, data.LogIndex)
	return data, nil
}

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
	"time"

	"bitbucket.org/sweetbridge/oracles/go-lib/chains"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/robert-zaremba/errstack"
)

// A block reader  extracts the data from ethcrawl objects into a BlockData structure
type blockReader struct {
	chainCtx    *ChainHandle
	transReader *transReader
	logReader   *logReader
}

// Factory for a BlockReader
func newBlockReader(chainCtx *ChainHandle) *blockReader {
	return &blockReader{chainCtx, newTransReader(chainCtx), newLogReader(chainCtx)}
}

// Reads a block and returns the BlockData structure
func (r *blockReader) read(block *types.Block, filter chains.AddressFilter) (*chains.BlockData, errstack.E) {
	txData, err := r.transReader.readList(block, filter)
	if err != nil {
		return nil, err
	}

	logData, err := r.logReader.readList(block, filter)
	if err != nil {
		return nil, err
	}

	// If no filtered transaction and no filtered log, nothing to process
	if len(txData)+len(logData) == 0 {
		return nil, nil
	}

	return &chains.BlockData{
		Header: &chains.HeaderData{
			NetworkID:    r.chainCtx.networkID,
			NetworkLabel: r.chainCtx.chainLabel,
			BlockHash:    block.Hash().String(),
			BlockNum:     block.Number().Int64(),
			Timestamp:    time.Unix(block.Time().Int64(), 0).UTC(),
		},
		Transactions: txData,
		Logs:         logData,
	}, nil
}

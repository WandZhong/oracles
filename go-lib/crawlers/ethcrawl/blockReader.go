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

	"bitbucket.org/sweetbridge/oracles/go-lib/crawlers"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/robert-zaremba/errstack"
)

// A block reader  extracts the data from ecrawl objects into a BCBlock structure
type blockReader struct {
	handle      *chainHandle
	transReader *transReader
	logReader   *logReader
}

// BlockReadingMode represents the different modes of reading blocks.
type BlockReadingMode int

const (
	// ReadTransDetail indicates that the tx details must be read
	ReadTransDetail BlockReadingMode = 1
	// ReadTransReceipt indicates that the transaction receipt must be read
	ReadTransReceipt BlockReadingMode = 2
	// ReadTransLog indicates that the transaction logs must be read
	ReadTransLog              BlockReadingMode = 4
	ReadTransDetailAndReceipt                  = ReadTransDetail + ReadTransReceipt
	ReadAll                                    = ReadTransDetailAndReceipt + ReadTransLog
)

// Returns a blockReader
func newBlockReader(handle *chainHandle, filter crawlers.AddressFilter, readingMode BlockReadingMode) *blockReader {
	var txReader *transReader
	var logReader *logReader

	if readingMode&ReadTransDetail == ReadTransDetail {
		txReader = newTransReader(handle, filter, readingMode&ReadTransReceipt == ReadTransReceipt)
	}
	if readingMode&ReadTransLog == ReadTransLog {
		logReader = newLogReader(handle, filter)
	}
	return &blockReader{
		handle:      handle,
		transReader: txReader,
		logReader:   logReader,
	}
}

func (r *blockReader) read(block *types.Block) (*crawlers.BCBlock, errstack.E) {
	// Reading transactions if a transReader is setup
	var txData []*crawlers.BCTrans
	var err errstack.E
	if r.transReader != nil {
		txData, err = r.transReader.readList(block)
		if err != nil {
			return nil, err
		}
	}

	// Reading logs if a logReader is setup
	var logData []*crawlers.BCLog
	if r.logReader != nil {
		logData, err = r.logReader.readList(block)
		if err != nil {
			return nil, err
		}
	}

	// If no transaction and no log are included, then empty data
	if len(txData) == 0 && len(logData) == 0 {
		return nil, nil
	}

	return &crawlers.BCBlock{
		BCBlockHeader: &crawlers.BCBlockHeader{
			NetworkID:    r.handle.networkID,
			NetworkLabel: r.handle.chainLabel,
			BlockHash:    block.Hash().String(),
			BlockNum:     block.Number().Int64(),
			Timestamp:    time.Unix(block.Time().Int64(), 0).UTC(),
		},
		Transactions: txData,
		Logs:         logData,
	}, nil
}

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
	"bitbucket.org/sweetbridge/oracles/go-lib/log"
	"github.com/robert-zaremba/errstack"
)

var logger = log.Root()

// implementation of the a crawler for Ethereum
type ethCrawler struct {
	chainCtx *ChainHandle
	iter     *iterator
	proc     chains.DataProcessor
	reader   *blockReader
	filter   chains.AddressFilter
}

// NewCrawler returns a ethCrawler
// - handle: the chain context handle
// - offset: indication of the block position to start the processing from. Values are:
//    > 0: starts at the offset from the block 0
//    < 0: starts at the offset from the last block
// - filter: an address filter
// - dp: the data processor
// Returns a ethCrawler reference of an error
func NewCrawler(handle *ChainHandle, offset int64, filter chains.AddressFilter, dp chains.DataProcessor) (chains.ChainCrawler, errstack.E) {
	// DataProcessor mandatory
	if dp == nil {
		return nil, errstack.NewDomain("No DataProcessor provided")
	}

	// Initiates a block iterator
	iter, err := newIterator(handle, offset)
	if err != nil {
		return nil, err
	}

	return &ethCrawler{
		chainCtx: handle,
		iter:     iter,
		proc:     dp,
		reader:   newBlockReader(handle),
		filter:   filter,
	}, nil
}

// Process ...
func (r *ethCrawler) Process() errstack.E {
	for true {
		block, err := r.iter.next()
		if err != nil {
			return err
		}

		// no more block to process on the chain
		if block == nil {
			break
		}

		blockData, err := r.reader.read(block, r.filter)
		if err != nil {
			return err
		}

		if blockData == nil {
			continue
		}

		logger.Debug("Processing eth block", "block", blockData.Header.BlockNum, "count", len(blockData.Transactions))
		if err = r.proc.Process(blockData); err != nil {
			return err
		}
	}
	return nil
}

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

package ethcrawl

import (
	"context"
	"fmt"
	"time"

	"bitbucket.org/sweetbridge/oracles/go-lib/crawlers"
	"bitbucket.org/sweetbridge/oracles/go-lib/log"
	"github.com/robert-zaremba/errstack"
)

var logger = log.Root()

// implementation of a crawler for Ethereum
type crawler struct {
	handle *chainHandle
	delay  int64
	iter   *iterator
	proc   crawlers.BCBlockProcessor
	reader *blockReader
	filter crawlers.AddressFilter
}

// NewCrawler returns a fully setup crawler. Parameters are:
// -ctx: An execution context
// -flags: the execution options
// -filter: an address filter
// -processor: a BCBlockProcessor
// -readingMode: Indicates what to extract. Values are:
//    ReadTransDetail  : Transaction details
//    ReadTransReceipt : Include Transaction Receipts related info
//    ReadTransLog     : Include the transactions logs
func NewCrawler(ctx context.Context, flags *crawlers.CrawlerOpts, filter crawlers.AddressFilter, processor crawlers.BCBlockProcessor, readingMode BlockReadingMode) (crawlers.ChainCrawler, errstack.E) {
	// preconditions
	if processor == nil {
		return nil, errstack.NewDomain("processor is a required parameter")
	}

	handle, err := newChainHandle(ctx, flags.ChainAddress, "")
	if err != nil {
		return nil, errstack.WrapAsDomain(err, "Error initializing a chain handle")
	}

	iter, err := newIterator(handle, flags.BlockOffset)
	if err != nil {
		return nil, errstack.WrapAsDomain(err, "Iterator init failure ")
	}

	return &crawler{
		handle: handle,
		delay:  flags.Delay,
		iter:   iter,
		proc:   processor,
		reader: newBlockReader(handle, filter, readingMode),
		filter: filter,
	}, nil
}

// Process is the main function of the crawler. It is a loop thar processed up to the last block at each iteration
// the number of iteration is defined by the delay field (initialised by config)
// - delay ==  0: only one iteration. When the last block is reached it ends.
// - delay > 0: there is a pause between each iteration. Delay's unit is minute.
// This function can behave in two ways according to the value "delay" initialised by config.
// - if delay = 0:
func (r *crawler) Process() errstack.E {
	if r.delay == 0 {
		return r.process()
	}

	d := time.Duration(int64(time.Minute) * r.delay)
	for {
		if err := r.process(); err != nil {
			return err
		}
		time.Sleep(d)
	}
}

// process process all the block from the current position up to the last block of the chain and returns
func (r *crawler) process() errstack.E {
	for {
		block, err := r.iter.next()
		// error or no more block to process on the chain
		if err != nil || block == nil {
			return err
		}

		logger.Debug(fmt.Sprintf("Block :%d, Timetamp:%s (TxCount: %d)\n", block.Number(), time.Unix(block.Time().Int64(), 0).UTC().String(), block.Transactions().Len()))
		bd, err := r.reader.read(block)
		if err != nil {
			return err
		}

		// Nothing of interest has been found
		if bd == nil {
			continue
		}

		if err = r.proc.Process(bd); err != nil {
			return err
		}
	}
}

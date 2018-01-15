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

package main

import (
	"bitbucket.org/sweetbridge/oracles/go-lib/analytics/elastic"
	"bitbucket.org/sweetbridge/oracles/go-lib/crawlers"
	"github.com/robert-zaremba/errstack"
)

// elasticFeeder is an implementation of BCBlockProcessor to send records as messages to ElasticSearch
type elasticFeeder struct {
	feed *elastic.Feed
}

const (
	esTransIndex   = "chain-trans"
	esTransDocType = "trans"
	esTransMsgType = "Ethereum.Transactions"
	esLogsIndex    = "chain-log"
	esLogsDocType  = "log"
	esLogsMsgType  = "Ethereum.Logs"
	esSourceName   = "EthereumFeed"
)

func newElasticProcessor(host string) crawlers.BCBlockProcessor {
	return &elasticFeeder{feed: elastic.NewFeed(host, esSourceName)}
}

// Process is the implementation of the BCBlock processor
func (r *elasticFeeder) Process(block *crawlers.BCBlock) errstack.E {

	if err := r.processTransactions(block.BCBlockHeader, block.Transactions); err != nil {
		return err
	}
	return r.processLogs(block.BCBlockHeader, block.Logs)
}

type msgTransBody struct {
	*crawlers.BCBlockHeader
	*crawlers.BCTrans
}

// Send a message per transaction
func (r *elasticFeeder) processTransactions(blockHeader *crawlers.BCBlockHeader, transactions []*crawlers.BCTrans) errstack.E {
	for _, t := range transactions {
		ti := msgTransBody{blockHeader, t}
		if err := r.feed.Send(esTransIndex, esTransDocType, esTransMsgType, t.ID, &ti); err != nil {
			return err
		}
	}
	return nil
}

type msgLogBody struct {
	*crawlers.BCBlockHeader
	*crawlers.BCLog
}

// Sends a message per log
func (r *elasticFeeder) processLogs(blockHeader *crawlers.BCBlockHeader, logs []*crawlers.BCLog) errstack.E {
	for _, l := range logs {
		li := msgLogBody{blockHeader, l}
		if err := r.feed.Send(esLogsIndex, esLogsDocType, esLogsMsgType, l.ID, &li); err != nil {
			return err
		}
	}
	return nil
}

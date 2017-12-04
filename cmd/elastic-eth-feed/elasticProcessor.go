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

package main

import (
	"bitbucket.org/sweetbridge/oracles/go-lib/analytics/elastic"
	"bitbucket.org/sweetbridge/oracles/go-lib/chains"
	"github.com/robert-zaremba/errstack"
)

// elasticProcessor is an implementation of DataProcessor to send messages to ElasticSearch
type elasticProcessor struct {
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

func newElasticProcessor(host string) chains.DataProcessor {
	return &elasticProcessor{feed: elastic.NewFeed(host, esSourceName)}
}

func (r *elasticProcessor) Process(block *chains.BlockData) errstack.E {

	if err := r.processTransactions(block.Header, block.Transactions); err != nil {
		return err
	}
	return r.processLogs(block.Header, block.Logs)
}

type msgTransBody struct {
	*chains.HeaderData
	*chains.TransData
}

// Send a message per transaction
func (r *elasticProcessor) processTransactions(blockHeader *chains.HeaderData, transactions []*chains.TransData) errstack.E {
	for _, t := range transactions {
		ti := msgTransBody{blockHeader, t}
		if err := r.feed.Send(esTransIndex, esTransDocType, esTransMsgType, t.ID, &ti); err != nil {
			return err
		}
	}
	return nil
}

type msgLogBody struct {
	*chains.HeaderData
	*chains.LogData
}

// Sends a message per log
func (r *elasticProcessor) processLogs(blockHeader *chains.HeaderData, logs []*chains.LogData) errstack.E {
	for _, l := range logs {
		li := msgLogBody{blockHeader, l}
		if err := r.feed.Send(esLogsIndex, esLogsDocType, esLogsMsgType, l.ID, &li); err != nil {
			return err
		}
	}
	return nil
}

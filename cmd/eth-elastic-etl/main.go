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
	"bitbucket.org/sweetbridge/oracles/go-lib/crawlers/ethcrawl"
	"bitbucket.org/sweetbridge/oracles/go-lib/log"
	"bitbucket.org/sweetbridge/oracles/go-lib/setup"
	"context"
	"github.com/robert-zaremba/log15/rollbar"
)

var (
	logger = log.Root()
)

func main() {
	opts := readMainOpts()
	setup.MustLogger("eth-elastic-etl", *opts.Rollbar)
	defer rollbar.WaitForRollbar(logger)

	elastic := newElasticProcessor(opts.elasticOpts.Endpoint)

	crawler, err := ethcrawl.NewCrawler(context.Background(), opts.CrawlerOpts, nil, elastic, ethcrawl.ReadAll)
	if err != nil {
		logger.Error("Error during ChainCrawler initialization", err)
		return
	}
	if err = crawler.Process(); err != nil {
		logger.Error("Processing failed", err)
		return
	}
}

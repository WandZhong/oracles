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
	"context"
	"errors"
	"flag"
	"time"

	"bitbucket.org/sweetbridge/oracles/go-lib/chains"
	"bitbucket.org/sweetbridge/oracles/go-lib/chains/ethcrawl"
	"bitbucket.org/sweetbridge/oracles/go-lib/log"
)

type flags struct {
	*chains.Flags
	Elastic string
	Delay   int64
}

func readFlags() (*flags, error) {
	elastic := flag.String("elastic", "http://localhost:9200", "The url of the instance of ElasticSearch")
	delay := flag.Int64("delay", 1, "Delay (in minutes) before resuming processing loop")
	ethFlags := chains.NewFlags()

	if *elastic == "" {
		return nil, errors.New("Param: 'elastic'. Cannot hold an empty value")
	}
	if err := ethFlags.Check(); err != nil {
		return nil, err
	}

	return &flags{Flags: ethFlags, Elastic: *elastic, Delay: *delay}, nil
}

func main() {
	logger := log.Root()
	logger.Info("Started ...")
	flags, err := readFlags()
	if err != nil {
		logger.Fatal("Command line parameters", err)
	}

	handle, err := ethcrawl.NewChainHandle(context.Background(), flags.ChainAddress, "")
	if err != nil {
		logger.Fatal("ChainHandle Init failure", err)
	}

	chainProc, err := ethcrawl.NewCrawler(handle, flags.BlockOffset, nil, newElasticProcessor(flags.Elastic))
	if err != nil {
		logger.Fatal("ethProcessor Init failure", err)
	}

	delay := time.Duration(int64(time.Minute) * flags.Delay)
	for {
		if err := chainProc.Process(); err != nil {
			logger.Fatal("Processing failure", err)
		}
		time.Sleep(delay)
	}
}

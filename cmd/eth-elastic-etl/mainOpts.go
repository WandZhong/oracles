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
	"bitbucket.org/sweetbridge/oracles/go-lib/crawlers"
	"bitbucket.org/sweetbridge/oracles/go-lib/setup"
	"github.com/robert-zaremba/errstack"
)

// the program's options
type mainOpts struct {
	*crawlers.CrawlerOpts `group:"eth-crawlers" namespace:"crawlers"`
	*elasticOpts          `group:"elastic"  namespace:"elastic"`
	setup.RollbarFlags    `group:"rollbar" namespace:"rollbar"`
}

// Check applies validation rules
func (r *mainOpts) Check() errstack.E {
	// Checks crawler options
	if err := r.CrawlerOpts.Check(); err != nil {
		return errstack.WrapAsReq(err, "Ethereum crawler options error")
	}

	return nil
}

// Options for elasticsearch
type elasticOpts struct {
	Endpoint string `long:"endpoint" required:"yes" description:"ElasticSearch endpoint"`
}

func readMainOpts() mainOpts {
	var opts mainOpts
	if err := setup.ReadOpts(opts); err != nil {
		logger.Fatal("Initialization error", err)
	}
	if err := opts.Check(); err != nil {
		logger.Fatal("Options validation error", err)
	}
	return opts
}

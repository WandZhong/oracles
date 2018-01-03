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
	"bitbucket.org/sweetbridge/oracles/go-lib/crawlers"
	"bitbucket.org/sweetbridge/oracles/go-lib/setup"
	"github.com/robert-zaremba/errstack"
)

// the program's options
type mainOpts struct {
	*setup.PgFlags        `group:"postgres" namespace:"pg"`
	*contribDB            `group:"db-loaders" namespace:"ldr"`
	*crawlers.CrawlerOpts `group:"eth-crawlers" namespace:"crawl"`
	setup.RollbarFlags    `group:"rollbar" namespace:"rollbar"`
}

// DB specific options for the contributions
type contribDB struct {
	StoredFunc string `long:"dbproc" description:"stored procedure name" default:"add_trans_with_order"`
}

// Check applies validation rules
func (r *mainOpts) Check() errstack.E {
	//TODO: @RobertZ: It would be good to have a way to append errors without a key like Append(string)
	errb := errstack.NewBuilder()

	if err := r.PgFlags.Check(); err != nil {
		errb.Put("PgFlags", "Options errors")
	}
	if err := r.CrawlerOpts.Check(); err != nil {
		errb.Put("crawlerOpts", "Invalid options")
	}

	if errb.NotNil() {
		return errb.ToReqErr()
	}
	return nil
}

func readMainOptions() (*mainOpts, errstack.E) {
	var opts mainOpts
	if err := setup.ReadOpts(&opts); err != nil {
		logger.Error("Command line params error")
		return nil, err
	}
	if err := opts.Check(); err != nil {
		return nil, err
	}
	return &opts, nil
}

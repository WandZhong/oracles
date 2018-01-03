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

package crawlers

import (
	"github.com/jessevdk/go-flags"
	"github.com/robert-zaremba/errstack"
	. "gopkg.in/check.v1"
)

type CrawlerOptsSuite struct{}

func (r CrawlerOptsSuite) TestOpts(c *C) {
	scenarii := []struct {
		desc        string
		args        []string
		expectedErr bool
	}{
		{
			"no command line option",
			[]string{},
			true,
		}, {
			"Missing chain option",
			[]string{"--chain"},
			true,
		}, {
			"offset != 0",
			[]string{"--chain=nonull", "--offset=10"},
			false,
		}, {
			"Offset = 0",
			[]string{"--chain=nonull", "--offset=0"},
			true,
		}, {
			"Delay >  0",
			[]string{"--chain=nonull", "--offset=10", "--delay=1"},
			false,
		}, {
			"Delay ==  0",
			[]string{"--chain=nonull", "--offset=10", "--delay=0"},
			false,
		}, {
			"Delay < 0",
			[]string{"--chain=nonull", "--offset=10", "--delay=-1"},
			true,
		},
	}

	for _, sc := range scenarii {
		err := checkCrawlerOps(sc.args)
		if sc.expectedErr {
			c.Check(err, NotNil)
		} else {
			c.Check(err, IsNil)
		}
	}
}

func checkCrawlerOps(args []string) errstack.E {
	opts := CrawlerOpts{}
	_, err := flags.ParseArgs(&opts, args)
	if err != nil {
		return errstack.WrapAsReq(err, "")
	}
	return opts.Check()
}

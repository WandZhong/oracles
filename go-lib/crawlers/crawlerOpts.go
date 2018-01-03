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
	"github.com/robert-zaremba/errstack"
)

// CrawlerOpts contains the options for crawling a chain
type CrawlerOpts struct {
	ChainAddress string   `long:"chain" required:"yes" description:"The ethereum blockchain endpoint"`
	BlockOffset  int64    `long:"offset" default:"-240" description:"Positive: starts at the block number matching the offset value.  Negative: starts at the block's number that is offset from the last block"`
	Delay        int64    `long:"delay" default:"0" description:"The delay to wait for when the last block is reached, before resuming the processing. If O, the process ends when it reaches last block"`
	Addresses    []string `long:"filter" required:"yes" description:"Addresses used to filter the transactions or logs of interest"`
}

// Check applies validation rules
func (r *CrawlerOpts) Check() errstack.E {
	errb := errstack.NewBuilder()
	if r.BlockOffset == 0 {
		errb.Put("offset", "the value must be strictly positive or strictly negative")
	}
	if r.Delay < 0 {
		errb.Put("delay", "Must be superior or equal to 0")
	}
	if errb.NotNil() {
		return errb.ToReqErr()
	}
	return nil
}

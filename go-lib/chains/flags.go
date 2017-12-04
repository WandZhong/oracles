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

package chains

import (
	"flag"

	"github.com/robert-zaremba/errstack"
)

// Flags contains the flags for the crawlers packages
type Flags struct {
	ChainAddress string
	BlockOffset  int64
}

// NewFlags returns parsed flags
func NewFlags() *Flags {
	chainAddress := flag.String("chain", "", "the address of the blockchain")
	blockOffset := flag.Int64("offset", -10, "The starting point of the crawling.\n \t-Negative: offset from last block of the chain. \n\t-Positive: block number to start from")
	flag.Parse()
	return &Flags{*chainAddress, *blockOffset}
}

// Check applies validation on provided flags
func (r *Flags) Check() errstack.E {
	errb := errstack.NewBuilder()

	if r.ChainAddress == "" {
		errb.Put("network", "Mandatory field")
	}
	if r.BlockOffset == 0 {
		errb.Put("offset", "the value must be strictly positive or strictly negative")
	}
	return errb.ToReqErr()
}

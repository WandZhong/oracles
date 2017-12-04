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

package ethcrawl

import (
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/robert-zaremba/errstack"
)

type iterator struct {
	nextBlockNumber *big.Int
	increment       *big.Int
	handle          *ChainHandle
}

// Factory for a iterator
// - handle: the chainContext
// - offset is used to determine the starting point of the iteration. Valid values are:
//   > 0 : is the offeset of the block starting from block zero
//   < 0 : is the negative offset starting at the last block
func newIterator(handle *ChainHandle, offset int64) (*iterator, errstack.E) {

	var current *big.Int

	switch {
	case offset == 0:
		return nil, errstack.NewDomain("Invalid offset value")
	case offset > 0:
		current = big.NewInt(offset)
	case offset < 0:
		block, err := handle.client.BlockByNumber(handle.ctx, nil)
		if err != nil {
			return nil, errstack.NewDomain(err.Error())
		}
		current = block.Number().Add(block.Number(), big.NewInt(offset))
	}
	return &iterator{nextBlockNumber: current, increment: big.NewInt(1), handle: handle}, nil
}

// Returns the block retrieved or an error if any. Block value can be:
//  - block != nil : a valid block is found and returned
//  - block == nil and err == NotFound : no block to process
// nextBlockNumber is used to read the next block and incremented only when the block is successfully read.
func (r *iterator) next() (*types.Block, errstack.E) {
	block, err := r.handle.client.BlockByNumber(r.handle.ctx, r.nextBlockNumber)
	if err == ethereum.NotFound {
		return nil, nil
	}
	if err != nil {
		return nil, errstack.NewDomain(err.Error())
	}
	r.nextBlockNumber.Add(r.nextBlockNumber, r.increment)
	return block, nil
}

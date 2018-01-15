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

package ethereum

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

var one = big.NewInt(1)

// IncNonce increments nonce by one and returns updated nonce
func IncNonce(nonce *big.Int) *big.Int {
	return nonce.Add(nonce, one)
}

// IncTxoNonce increments transaction options nonce
func IncTxoNonce(txo *bind.TransactOpts, tx *types.Transaction) {
	if txo.Nonce == nil {
		txo.Nonce = big.NewInt(int64(tx.Nonce()))
	}
	IncNonce(txo.Nonce)
}

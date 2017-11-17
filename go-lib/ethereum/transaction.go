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

package ethereum

import (
	"fmt"
	"io"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/robert-zaremba/errstack"
)

// LogTx is a handy function to create debug log for successful transaction
func LogTx(msg string, tx *types.Transaction) {
	if tx != nil {
		logger.Debug(msg, "tx_hash", tx.Hash().Hex(), "nonce", tx.Nonce(),
			"gas", tx.Gas(), "gas_price", tx.GasPrice())
	} else {
		logger.Debug("Invalid transaction")
	}
}

// FlogTx logs transaction into a Writer
func FlogTx(w io.Writer, msg string, tx *types.Transaction) {
	if tx != nil {
		fmt.Fprintf(w, "%s\n\thash=%s, gas=%v, gas_price=%v\n",
			msg, tx.Hash().Hex(), tx.Gas(), tx.GasPrice())
	} else {
		_, err := w.Write([]byte(msg + ": invalid transaction\n"))
		errstack.Log(logger, err)
	}
}

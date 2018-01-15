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

package ethcrawl

import (
	"math/big"

	"bitbucket.org/sweetbridge/oracles/go-lib/crawlers"
	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"

	"math"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/robert-zaremba/errstack"
)

var (
	weiToEth = big.NewFloat(math.Pow10(18))
)

type transReader struct {
	handle         *chainHandle
	filter         crawlers.AddressFilter
	addReceiptInfo bool
}

func newTransReader(handle *chainHandle, filter crawlers.AddressFilter, addReceiptInfo bool) *transReader {
	return &transReader{
		handle:         handle,
		filter:         filter,
		addReceiptInfo: addReceiptInfo,
	}
}

// reads the list of transactions from an ethereum block.
func (r *transReader) readList(block *types.Block) ([]*crawlers.BCTrans, errstack.E) {
	txs := block.Transactions()
	var lst []*crawlers.BCTrans
	for i, tx := range txs {
		data, err := r.read(tx, i)
		if err != nil {
			return nil, err
		}
		if data != nil {
			lst = append(lst, data)
		}
	}

	return lst, nil
}

// Returns a TransData object if it matches the provided filter.
func (r *transReader) read(tx *types.Transaction, index int) (*crawlers.BCTrans, errstack.E) {
	var toAddr = tx.To()
	var toAddrStr, fromAddrStr string
	if toAddr != nil {
		toAddrStr = toAddr.String()
	}
	fromAddrStr = sender(tx).String()

	// Applying the filters
	if r.filter != nil {
		noMatch, err := r.filter.MatchesNone(toAddrStr, fromAddrStr)
		if err != nil {
			return nil, err
		}
		if noMatch {
			return nil, nil
		}
	}

	data := &crawlers.BCTrans{
		TxIndex:  int64(index),
		Data:     string(tx.Data()),
		TxHash:   tx.Hash().String(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice().Uint64(),
		From:     fromAddrStr,
		To:       toAddrStr,
		RawValue: tx.Value(),
	}
	data.ID = data.TxHash
	fl := new(big.Float)
	fl.SetInt(data.RawValue)
	data.Value = fl.Quo(fl, weiToEth)

	// Executed if TransactionReceipt are requested
	if r.addReceiptInfo {
		tr, err := r.handle.client.TransactionReceipt(r.handle.ctx, tx.Hash())
		if err != nil {
			if err != nil {
				return nil, errstack.WrapAsReq(err, "Can't get transaction receipt "+tx.Hash().Hex())
			}
		}

		data.GasUsed = tr.GasUsed
		data.CumulativeGasUsed = tr.CumulativeGasUsed

		contractAddrStr := ""
		contractAddr := tr.ContractAddress
		contractCreation := !ethereum.IsZeroAddr(contractAddr)
		if contractCreation {
			contractAddrStr = contractAddr.String()
		}
		data.CreatedContractAddress = contractAddrStr
		data.ContractCreated = contractCreation

	}
	return data, nil
}

func sender(tx *types.Transaction) *common.Address {
	var fromAddr common.Address
	v, _, _ := tx.RawSignatureValues()
	if v != nil {
		// make a best guess about the signer and use that to derive
		// the sender.
		signer := deriveSigner(tx, v)
		if f, err := types.Sender(signer, tx); err != nil { // derive but don't cache
		} else {
			fromAddr = common.BytesToAddress(f[:])
		}
	}
	return &fromAddr
}

func deriveSigner(tx *types.Transaction, v *big.Int) types.Signer {
	if v.Sign() != 0 && tx.Protected() {
		return types.NewEIP155Signer(tx.ChainId())
	}
	return types.HomesteadSigner{}
}

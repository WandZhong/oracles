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

	"bitbucket.org/sweetbridge/oracles/go-lib/chains"
	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/robert-zaremba/errstack"
)

type transReader struct {
	handle *ChainHandle
}

func newTransReader(handle *ChainHandle) *transReader {
	return &transReader{
		handle,
	}
}

// reads the list of transactions from an ethereum block.
func (r *transReader) readList(block *types.Block, filter chains.AddressFilter) ([]*chains.TransData, errstack.E) {
	txs := block.Transactions()
	var lst []*chains.TransData
	for i, tx := range txs {
		data, err := r.read(tx, i, filter)
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
func (r *transReader) read(tx *types.Transaction, index int, filter chains.AddressFilter) (*chains.TransData, errstack.E) {
	tr, err := r.handle.client.TransactionReceipt(r.handle.ctx, tx.Hash())
	if err != nil {
		return nil, errstack.NewDomain(err.Error())
	}

	var toAddr = tx.To()
	var toAddrStr, fromAddrStr string
	if toAddr != nil {
		toAddrStr = toAddr.String()
	}
	fromAddrStr = sender(tx).String()

	// Applying the filters
	if filter != nil && filter.MatchesNone(toAddrStr, fromAddrStr) {
		return nil, nil
	}

	contractAddrStr := ""
	contractAddr := tr.ContractAddress
	contractCreation := !ethereum.IsZeroAddr(contractAddr)
	if contractCreation {
		contractAddrStr = contractAddr.String()
	}

	data := &chains.TransData{
		TxIndex: int64(index),
		Data:    string(tx.Data()),
		TxHash:  tx.Hash().String(),

		CumulativeGasUsed: tr.CumulativeGasUsed.Int64(),
		Gas:               tx.Gas().Int64(),
		GasPrice:          tx.GasPrice().Int64(),
		GasUsed:           tr.GasUsed.Int64(),
		From:              fromAddrStr,
		To:                toAddrStr,
		Value:             tx.Value().Int64(),
		CreatedContractAddress: contractAddrStr,
		ContractCreated:        contractCreation,
	}
	data.ID = data.TxHash

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

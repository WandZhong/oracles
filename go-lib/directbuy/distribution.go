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

package directbuy

import (
	"fmt"
	"math/big"

	contracts "bitbucket.org/sweetbridge/oracles/go-contracts"
	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum/wad"
	"bitbucket.org/sweetbridge/oracles/go-lib/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/go-pg/pg"
	"github.com/robert-zaremba/errstack"
	"github.com/robert-zaremba/log15"
)

// Token is the minimum interface to execute distribution
type Token interface {
	BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error)
	Transfer(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error)
}

// CheckRequiredBalance verifies if the `cf` transactor address has enough `token` balance
// to cover all distribution.
func CheckRequiredBalance(summaries []Summary, tokenname string, token *contracts.SweetToken, cf ethereum.ContractFactory) errstack.E {
	var total = new(big.Int)
	for _, r := range summaries {
		total.Add(total, r.Amount)
	}
	addr := cf.Addr()
	logger.Debug(tokenname+" distribution account holder", "address", addr.Hex())
	balance, err := token.BalanceOf(nil, addr)
	if err != nil {
		return errstack.WrapAsInf(err, fmt.Sprintf("Can't check %s balance", tokenname))
	}
	if balance.Cmp(total) < 0 {
		return errstack.NewReqF("Not enough funds in the source account = %v, min_expected=%v",
			wad.WeiToString(balance), wad.WeiToInt(total))
	}
	logger.Debug("Distribution account balance", tokenname+".wei", balance.String(),
		"required", total.String())
	return nil
}

// Distribute performs on-chain token distirubtion based on aggregated summaries.
func Distribute(summaries []Summary, token Token, cf ethereum.ContractFactory, db *pg.DB) errstack.E {
	txo := cf.Txo()
	for _, s := range summaries {
		logger.Debug("Transfering", "amount", wad.WeiToString(s.Amount),
			"dest", s.Address.Hex(), "nonce", txo.Nonce)
		tx, err := token.Transfer(txo, s.Address, s.Amount)
		if err != nil {
			logger.Error("Can't transfer TOKEN", err)
			break
		} else {
			ethereum.LogTx("Transferred", tx)
			ethereum.IncTxoNonce(txo, tx)
			// logger.Debug(">>>> nonce after", "txo", txo.Nonce, "tx", tx.Nonce())
		}
		if db != nil {
			if err := s.SetStatusSent(db); err != nil {
				logger.Crit("Can't set status DONE", log15.Spew(s))
				return err
			}
		}
	}
	if db != nil {
		logger.Info("DirctBuy DB statuses haven't been updated")
	}
	return nil
}

// Summary represents the aggregated (per user) view of DirectBuy distribution
type Summary struct {
	Idx     int // report row index
	List    string
	Address common.Address
	Amount  *big.Int
	DBs     []DirectBuy
}

// String implements stringer interface
func (s Summary) String() string {
	return fmt.Sprint("{", s.Idx, " ", s.List, " ", s.Address.Hex(), " ", s.Amount, "}")
}

// SetStatusSent update the status for each DirectBuy record in summary to Done
func (s Summary) SetStatusSent(db *pg.DB) errstack.E {
	var ids = make([]int64, len(s.DBs))
	for i := range s.DBs {
		ids[i] = s.DBs[i].ID
		s.DBs[i].Status = StatusSent
	}
	res, err := db.Model(&DirectBuy{}).Set("status = ?", StatusSent).
		Where("direct_buy_id IN ?", pg.In(ids)).Update()
	return model.CheckRowsAffected(fmt.Sprintf("Update direct_buy[%d] status", ids),
		1, res, err)

}

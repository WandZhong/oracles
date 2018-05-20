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

	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum/wad"
	"bitbucket.org/sweetbridge/oracles/go-lib/model"
	"bitbucket.org/sweetbridge/oracles/go-lib/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/go-pg/pg"
	"github.com/robert-zaremba/errstack"
	"github.com/robert-zaremba/log15"
)

var zero = big.NewInt(0)

// Token is the minimum interface to execute distribution
type Token interface {
	BalanceOf(opts *bind.CallOpts, who common.Address) (*big.Int, error)
	Transfer(opts *bind.TransactOpts, dst common.Address, wad *big.Int) (*types.Transaction, error)
}

// CheckRequiredBalance verifies if the `cf` transactor address has enough `token` balance
// to cover all distribution.
func CheckRequiredBalance(summaries []Summary, tokenname string, token Token, cf ethereum.ContractFactory) errstack.E {
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
	logger.Debug("Distribution account balance",
		tokenname+".wei", balance.String(), tokenname, wad.WeiToString(balance),
		"required.wei", total.String(), "required", wad.WeiToString(total))
	return nil
}

// Distribute performs on-chain token distirubtion based of aggregated summaries.
func Distribute(dryRun bool, summaries []Summary, token Token, cf ethereum.ContractFactory, db *pg.DB) errstack.E {
	if len(summaries) == 0 {
		logger.Info("Nothing to distribute - no pending records in given tranche")
		return nil
	}
	if err := CheckRequiredBalance(summaries, "SWC", token, cf); err != nil {
		return err
	}
	if dryRun {
		logger.Debug("Dry run. Stopping execution.")
		return nil
	}
	logger.Info("Starting token distribution")
	var totalGas uint64
	txo := cf.Txo()
	for _, s := range summaries {
		logger.Debug("Transfering", "amount", wad.WeiToString(s.Amount),
			"dest", s.Address.Hex(), "nonce", txo.Nonce)
		if s.Amount.Cmp(zero) == 0 {
			logger.Warn("Ignoring, the amount_out is zero", "row", s.Idx)
			continue
		}
		tx, err := token.Transfer(txo, s.Address, s.Amount)
		if err != nil {
			logger.Error("Can't transfer TOKEN", err)
			break
		} else {
			ethereum.LogTx("Transferred", tx)
			totalGas += tx.Gas()
			ethereum.IncTxoNonce(txo, tx)
			// logger.Debug(">>>> nonce after", "txo", txo.Nonce, "tx", tx.Nonce())
		}
		if db != nil {
			if err := s.SetStatusSent(db); err != nil {
				logger.Info("Stopping distribution", "total_gas_spent", totalGas)
				logger.Crit("Can't set status DONE", log15.Spew(s))
				return err
			}
		}
	}
	logger.Info("Distribution finished", "total_gas_spent", totalGas)
	if db == nil {
		logger.Info("DirctBuy DB statuses haven't been updated")
	}
	return nil
}

// DistributeSWC performs on-chain SWC token distirubtion based of aggregated summaries.
func DistributeSWC(dryRun bool, summaries []Summary, cf ethereum.ContractFactory, db *pg.DB) errstack.E {
	swcC, addr, err := cf.GetSWC()
	utils.Assert(err, "Can't instantiate SWT contract")
	logger.Debug("Contract address", "swc", addr.Hex())

	return Distribute(dryRun, summaries, swcC, cf, db)
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
	if len(s.DBs) == 0 {
		return nil
	}
	var ids = make([]int64, len(s.DBs))
	for i := range s.DBs {
		ids[i] = s.DBs[i].ID
		s.DBs[i].Status = StatusSent
	}
	logger.Debug("updating status", "dbs", ids)
	res, err := db.Model(&DirectBuy{}).Set("status = ?", StatusSent).
		Where("direct_buy_id IN (?)", pg.In(ids)).Update()
	return model.CheckRowsAffected(fmt.Sprintf("Update direct_buy %v status", ids),
		len(ids), res, err)

}

/* Not finished. Tests are still failing
func checkWhitelist(summaries []directbuy.Summary, cf ethereum.ContractFactory) errstack.E {
	swcLogic, _, err := cf.GetSWClogic()
	utils.Assert(err, "Can't instantiate SWTlogic contract")

	var listName = [32]byte{}
	copy(listName[:], []byte("whitelist"))
	logger.Debug("Constructing whitelist name", "bytes", listName)
	txo := cf.Txo()
	if ok, err := swcLogic.ListExists(nil, listName); err != nil {
		return errstack.WrapAsInf(err, "Can't check if whitelist exists")
	} else if !ok {
		tx, err := swcLogic.AddWhiteList(txo, listName)
		if err != nil {
			return errstack.WrapAsInf(err, "Can't create SWC whitelist")
		}
		ethereum.LogTx("SWC whitelist created", tx)
	}

	for _, r := range summaries {
		if ok, err := swcLogic.WhiteLists(nil, r.Address, listName); err != nil {
			return errstack.WrapAsInfF(err, "Can't check if user %s is in the whitelist",
				r.Address)
		} else if ok {
			continue
		}
		tx, err := swcLogic.AddToWhiteList(txo, listName, r.Address)
		if err != nil {
			return errstack.WrapAsInfF(err, "Can't add user %s to the whitelist", r.Address)
		}
		ethereum.LogTx(fmt.Sprintf("User %s added to the whitelist", r.Address), tx)
	}
	return nil

}
*/

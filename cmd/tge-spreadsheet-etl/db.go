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

package main

import (
	"time"

	"bitbucket.org/sweetbridge/oracles/go-lib/directbuy"
	"bitbucket.org/sweetbridge/oracles/go-lib/model"
	"github.com/robert-zaremba/errstack"
	"github.com/robert-zaremba/go-pgt"
)

// missing user is not reported as an error!
func findUserByEmail(r Record) (user pgt.UUID, errS errstack.E) {
	_, err := db.QueryOne(&user, "SELECT id FROM individual WHERE lower(email_address) = ?", r.Email)
	if err != nil {
		errS = model.CheckPgNoRows("individual", err)
	}
	return user, errS
}

func createDirectBuys(records []Record) ([]directbuy.DirectBuy, errstack.E) {
	now := time.Now()
	var dbs []directbuy.DirectBuy
	var err errstack.E
	for _, r := range records {
		d := r.DirectBuy
		if d.UserID, err = findUserByEmail(r); err != nil {
			if !err.IsReq() {
				return dbs, err
			}
			logger.Warn("Can't find user", "email", r.Email, "name", r.FullName)
		}
		d.CreatedAt = now
		d.UpdatedAt = now
		if err = d.MkHash(r.TxHash); err != nil {
			return dbs, err
		}
		dbs = append(dbs, d)
	}
	return dbs, nil
}

func insert(records []Record) errstack.E {
	ds, err := createDirectBuys(records)
	if err != nil {
		return err
	}

	logger.Info("All direct_buys created. Inserting into DB...")
	result, errStd := db.Model(&ds).
		OnConflict("(direct_buy_id) DO UPDATE").
		Set(`hash = EXCLUDED.hash, individual_id = EXCLUDED.individual_id, email = EXCLUDED.email, tranche_id = EXCLUDED.tranche_id, amount_out = EXCLUDED.amount_out, amount_in = EXCLUDED.amount_in, currency_id = EXCLUDED.currency_id, usd_rate = EXCLUDED.usd_rate, sender_id = EXCLUDED.sender_id, updated_at = EXCLUDED.updated_at,
hash_out = CASE WHEN direct_buy.status=? THEN direct_buy.hash_out ELSE EXCLUDED.hash_out END,
nonce = CASE WHEN direct_buy.status=? THEN direct_buy.nonce ELSE EXCLUDED.nonce END,
status = CASE WHEN direct_buy.status=? THEN direct_buy.status ELSE EXCLUDED.status END`,
			directbuy.StatusSent, directbuy.StatusSent, directbuy.StatusSent).
		Insert()
	if errStd == nil {
		logger.Info("direct_buys insert finished", "rows_affected", result.RowsAffected())
	}
	return errstack.WrapAsInf(errStd, "DB direct_buy insert")
	// for i := range ds {
	// 	err = errstack.WrapAsInf(db.Insert(&ds[i]), "DB direct_buy insert")
	// 	if err != nil {
	// 		fmt.Println(ds[i])
	// 		return err
	// 	}
	// }
	// return nil
}

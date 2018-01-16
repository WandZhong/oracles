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
	"strings"
	"time"

	"bitbucket.org/sweetbridge/oracles/go-lib/liquidity"
	"github.com/robert-zaremba/errstack"
	"github.com/robert-zaremba/go-pgt"
)

// Transaction represents transactions_tge DB record
type Transaction struct {
	tableName struct{} `sql:"transactions_tge"`

	ID       pgt.UUID           `sql:"transaction_id,pk"`
	Amount   float64            `sql:"amount,notnull"`
	Currency liquidity.Currency `sql:"currency_id"`
	UserID   pgt.UUID           `sql:"individual_id"`
	Email    string             `sql:"email"`
	SenderID string             `sql:"sender_id"`

	TransactionDate time.Time `sql:"transaction_date,notnull"`
	CreatedAt       time.Time `sql:"created_at,notnull"`
	UpdatedAt       time.Time `sql:"updated_at,notnull"`
}

// NewTransactionFromRecord
func NewTransactionFromRecord(r Record, now time.Time) (Transaction, errstack.E) {
	var user pgt.UUID
	_, err := db.QueryOne(&user, "SELECT id FROM individual WHERE email_address = ?", r.Email)
	if err != nil {
		if strings.Index(err.Error(), "no rows in") < 0 {
			return Transaction{}, errstack.WrapAsInf(err, "can't query DB")
		}
		logger.Error("Can't find user", "email", r.Email, "name", r.FullName)
	}
	return Transaction{
		struct{}{}, pgt.RandomUUID(), r.Amount, r.Currency, user, r.Email, r.SenderID,
		r.Timestamp, now, now,
	}, nil
}

func createTransactions(records []Record) ([]Transaction, errstack.E) {
	now := time.Now()
	var ts []Transaction
	for _, r := range records {
		t, err := NewTransactionFromRecord(r, now)
		if err != nil {
			if !err.IsReq() {
				return ts, err
			}
			logger.Error("can't create transaction: " + err.Error())
		} else {
			ts = append(ts, t)
		}
	}
	return ts, nil
}

func insertTransactions(records []Record) errstack.E {
	ts, err := createTransactions(records)
	if err != nil {
		return err
	}
	logger.Info("All transactions properly created. Inserting into DB...")
	return errstack.WrapAsInf(db.Insert(&ts), "DB insert")
}

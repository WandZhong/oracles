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
	"hash/fnv"
	"strings"
	"time"

	"bitbucket.org/sweetbridge/oracles/go-lib/liquidity"
	"github.com/go-pg/pg"
	"github.com/robert-zaremba/errstack"
	pgt "github.com/robert-zaremba/go-pgt"
)

// DirectBuy represents the swc direct buy for the next distribution.
type DirectBuy struct {
	ID        int64              `sql:"direct_buy_id,pk"`
	Hash      []byte             `sql:"hash"`
	UserID    pgt.UUID           `sql:"individual_id"`
	Email     string             `sql:"email"`
	TrancheID uint64             `sql:"tranche_id,notnull"`
	AmountOut float64            `sql:"amount_out,notnull"`
	AmountIn  float64            `sql:"amount_in,notnull"`
	Currency  liquidity.Currency `sql:"currency_id"`
	UsdRate   float64            `sql:"usd_rate,notnull"`
	SenderID  string             `sql:"sender_id"`
	Status    Status             `sql:"status,notnull"`

	TransactionDate time.Time `sql:"transaction_date,notnull"`
	CreatedAt       time.Time `sql:"created_at,notnull"`
	UpdatedAt       time.Time `sql:"updated_at,notnull"`
}

// Status a type for DirectBuy.Status field
type Status uint

// ParseStatusErrp converts string status verose name to uint and sets error into errp
func ParseStatusErrp(s string, errp errstack.Putter) Status {
	switch strings.TrimSpace(strings.ToLower(s)) {
	case "on hold":
		return StatusOnHold
	case "pending":
		return StatusPending
	case "sent":
		return StatusSent
	}
	errp.Put("Unknown status " + s)
	return StatusOnHold
}

// DirectBuy statuses
const (
	StatusOnHold Status = iota
	StatusPending
	StatusSent
)

// MkHash computes and assigns hash to the DirectBuy record
func (d *DirectBuy) MkHash(txHash string) errstack.E {
	h := fnv.New128()
	for _, s := range []string{d.TransactionDate.String(),
		d.Email,
		string(d.Currency),
		txHash} {
		if _, err := h.Write([]byte(s)); err != nil {
			return errstack.WrapAsInf(err, "can't sum the string: "+s)
		}
	}
	d.Hash = h.Sum(nil)
	return nil
}

// GetPendingDirectBuys retrieves pending direct buys from given trancheID.
func GetPendingDirectBuys(trancheID string, db *pg.DB) ([]DirectBuy, errstack.E) {
	var ds = []DirectBuy{}
	err := db.Model(&ds).
		Where("tranche_id = ? AND status = ?", trancheID, StatusPending).
		Select()
	return ds, errstack.WrapAsInf(err, "Can't get DirectBuy records")
}

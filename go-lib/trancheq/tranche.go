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

package trancheq

import (
	"time"

	"bitbucket.org/sweetbridge/oracles/go-lib/liquidity"
	"bitbucket.org/sweetbridge/oracles/go-lib/model"
	"github.com/go-pg/pg"
	"github.com/robert-zaremba/errstack"
	pgt "github.com/robert-zaremba/go-pgt"
)

type price struct {
	tableName struct{} `sql:"tranche_prices"`

	TrancheID int64              `sql:"tranche_id,pk"`
	Currency  liquidity.Currency `sql:"currency_id"`
	Price     float64            `sql:"price"`
}

// TrancheDB represents tranches DB table entry
type TrancheDB struct {
	tableName struct{} `sql:"tranches"`

	ID         int64      `sql:"tranche_id,pk" json:"id"`
	TokenID    string     `sql:",notnull" json:"tokenID"`
	CreatedAt  time.Time  `sql:"created_at,notnull" json:"createdAt"`
	StartsAt   time.Time  `sql:"starts_at,notnull" json:"startsAt"`
	ExecutesAt time.Time  `sql:"executes_at,notnull" json:"executesAt"`
	EndsAt     *time.Time `sql:"ends_at" json:"endsAt"`
	Supply     pgt.BigInt `sql:"supply,notnull" json:"supply"`          // in wad
	MaxContrib pgt.BigInt `sql:"max_contrib,notnull" json:"maxContrib"` // in wad, 0=no limit
}

// Tranche represents tranche model
type Tranche struct {
	TrancheDB
	Prices map[liquidity.Currency]float64 `json:"prices"`
}

func newTrache(t TrancheDB, ps []price) Tranche {
	var out = Tranche{t, map[liquidity.Currency]float64{}}
	for _, p := range ps {
		out.Prices[p.Currency] = p.Price
	}
	return out
}

// Validate checks if all fields are sematically correct.
// It also resets the automatic fields (ID, CreatedAt)
func (t *Tranche) Validate() errstack.Builder {
	var errb = errstack.NewBuilder()
	t.ID = 0
	t.CreatedAt = time.Now().UTC()
	if t.ExecutesAt.Before(t.CreatedAt.Add(time.Hour)) {
		errb.Put("executesOn", "must be after 'now'+1hour")
	}
	if !t.ExecutesAt.After(t.StartsAt) {
		errb.Put("executesOn", "must be after `startsOn`")
	}
	if t.EndsAt != nil && !t.EndsAt.After(t.StartsAt) {
		errb.Put("endsOn", "must be after `startsOn`")
	}
	if t.Supply.Int == nil || t.Supply.Cmp(zero) <= 0 {
		errb.Put("supply", "must be bigger than zero")
	}
	if t.MaxContrib.Int != nil && t.MaxContrib.Cmp(zero) < 0 {
		errb.Put("maxContrib", "can't be negative")
	}
	return errb
}

// Save stores the object into DB
func (t *Tranche) Save(db *pg.DB) errstack.E {
	if err := db.Insert(&t.TrancheDB); err != nil {
		return errstack.WrapAsInf(err, "Can't insert Tranche")
	}
	if len(t.Prices) == 0 {
		return nil
	}
	var prices = make([]price, len(t.Prices))
	i := 0
	for curr, val := range t.Prices {
		prices[i].Currency = curr
		prices[i].Price = val
		prices[i].TrancheID = t.ID
		i++
	}
	return errstack.WrapAsInf(db.Insert(&prices), "Can't insert prices")
}

// GetTranche returns tranche from DB
func GetTranche(id int64, db *pg.DB) (Tranche, errstack.E) {
	var tdb = TrancheDB{ID: id}
	if err := model.CheckPgNoRows("token", db.Select(&tdb)); err != nil {
		return Tranche{}, err
	}
	var ps = []price{}
	err := model.CheckPgNoRows("tranche_prices",
		db.Model(&ps).Where("tranche_id = ?", id).Select())
	if err != nil {
		return Tranche{}, err
	}
	return newTrache(tdb, ps), nil
}

// GetTranches returns Tranches combined with prices from DB.
// TODO: in the future we need to return only active tranches.
func GetTranches(db *pg.DB) ([]Tranche, errstack.E) {
	var tranches = []TrancheDB{}
	err := db.Model(&tranches).Select()
	if err := model.ErrNotNoRows("tranche", err); err != nil {
		return nil, err
	}
	var resp = make([]Tranche, 0, len(tranches))
	if len(tranches) == 0 {
		return resp, nil
	}
	var prices = []price{}
	err = db.Model(&prices).Select()
	if err := model.ErrNotNoRows("prices", err); err != nil {
		return nil, err
	}

	var priceMap = map[int64][]price{}
	for _, p := range prices {
		priceMap[p.TrancheID] = append(priceMap[p.TrancheID], p)
	}
	for i := range tranches {
		ps := priceMap[tranches[i].ID]
		resp = append(resp, newTrache(tranches[i], ps))
	}
	return resp, nil
}

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

package trancheq

import (
	"time"

	"bitbucket.org/sweetbridge/oracles/go-lib/model"
	"github.com/go-pg/pg"
	"github.com/robert-zaremba/errstack"
	pgt "github.com/robert-zaremba/go-pgt"
)

// Tranche represents tranches DB table entry
type Tranche struct {
	ID         int64      `sql:"tranche_id,pk" json:"id"`
	TokenID    string     `sql:",notnull" json:"tokenID"`
	CreatedOn  time.Time  `sql:"created_on,notnull" json:"createdOn"`
	StartsOn   time.Time  `sql:"starts_on,notnull" json:"startsOn"`
	ExecutesOn time.Time  `sql:"executes_on,notnull" json:"executesOn"`
	Supply     pgt.BigInt `sql:"supply,notnull" json:"supply"`          // in wad
	MaxContrib pgt.BigInt `sql:"max_contrib,notnull" json:"maxContrib"` // in wad, 0=no limit

	// TODO: how to handle different flavors (brg*)
	PriceBRGusd float64 `sql:"price_brg_usd,notnull" json:"priceBRGusd"`
}

// Validate checks if all fields are sematically correct.
// It also resets the automatic fields (ID, CreatedOn)
func (t *Tranche) Validate() errstack.Builder {
	var errb = errstack.NewBuilder()
	t.ID = 0
	t.CreatedOn = time.Now().UTC()
	if t.ExecutesOn.Before(t.CreatedOn.Add(time.Hour)) {
		errb.Put("executesOn", "must be after 'now'+1hour")
	}
	if t.ExecutesOn.Before(t.StartsOn) {
		errb.Put("executesOn", "must be after `startsOn`")
	}
	if t.Supply.Int == nil || t.Supply.Cmp(zero) <= 0 {
		errb.Put("supply", "must be bigger than zero")
	}
	if t.MaxContrib.Int != nil && t.MaxContrib.Cmp(zero) < 0 {
		errb.Put("maxContrib", "can't be negative")
	}
	if t.PriceBRGusd <= 0 {
		errb.Put("priceBRGusd", "must be bigger than zero")
	}
	return errb
}

// GetTranche returns tranche from DB
func GetTranche(id int64, db *pg.DB) (Tranche, errstack.E) {
	var t = Tranche{ID: id}
	return t, model.CheckPgNoRows("token", db.Select(&t))
}

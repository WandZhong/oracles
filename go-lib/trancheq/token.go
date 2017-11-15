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
	"math/big"
	"time"

	"bitbucket.org/sweetbridge/oracles/go-lib/model"
	"github.com/go-pg/pg"
	"github.com/robert-zaremba/errstack"
	pgt "github.com/robert-zaremba/go-pgt"
)

var zero = big.NewInt(0)

// Token represents tokens DB table entry
type Token struct {
	ID        string    `sql:"token_id,pk"` // token name
	CreatedOn time.Time `sql:"created_on,notnull"`
	Comment   string    `sql:"comment,notnull"`

	MaxTotalContrib pgt.BigInt `sql:"max_total_contrib" json:"maxTotalContrib"` // in wad, 0 = no limit
}

// Validate checks if all fields are sematically correct.
// It also sets the automatic fields (ID, CreatedOn)
func (t *Token) Validate() errstack.Builder {
	var errb = errstack.NewBuilder()
	t.CreatedOn = time.Now().UTC()
	if len(t.ID) < 3 {
		errb.Put("id", "must be at least 3-characters long")
	}
	if t.MaxTotalContrib.Int != nil && t.MaxTotalContrib.Cmp(zero) < 0 {
		errb.Put("maxTotalContrib", "can't be negative")
	}
	return errb
}

// GetToken returns token from DB
func GetToken(id string, db *pg.DB) (Token, errstack.E) {
	var t = Token{ID: id}
	return t, model.CheckPgNoRows("token", db.Select(&t))
}

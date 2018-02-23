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
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/go-pg/pg"
	"github.com/robert-zaremba/errstack"
	pgt "github.com/robert-zaremba/go-pgt"
)

// CreateDistributionAccount creates a SWC distribution account record into DB
func CreateDistributionAccount(userID pgt.UUID, addr common.Address, db *pg.DB) errstack.E {
	now := time.Now().UTC()
	a := Account{
		struct{}{}, pgt.RandomUUID(), userID, addr.Hex(),
		distributionAccountName, "primary", now, now,
	}
	return errstack.WrapAsInf(db.Insert(&a), "Can't insert SWC distribution member_account")
}

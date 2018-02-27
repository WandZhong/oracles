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
	"strings"
	"time"

	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"bitbucket.org/sweetbridge/oracles/go-lib/model"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-pg/pg"
	"github.com/robert-zaremba/errstack"
	pgt "github.com/robert-zaremba/go-pgt"
)

const distributionAccountName = "Sweetcoin Distribution"

// Account represents DB members account record
type Account struct {
	tableName    struct{} `sql:"member_account"`
	ID           pgt.UUID `sql:"id,pk"`
	IndividualID pgt.UUID `sql:"individual_id"`
	Number       string   `sql:"account_number"`
	Name         string   `sql:"account_name"`
	Reference    string   `sql:"reference"`

	CreatedAt time.Time `sql:"created_at,notnull"`
	UpdatedAt time.Time `sql:"updated_at,notnull"`
}

// FindDistributionAccount selects user etherum account used to distribute SWC tokens.
func FindDistributionAccount(userID pgt.UUID, db *pg.DB) (common.Address, errstack.E) {
	if userID.Empty() {
		return ethereum.ZeroAddress, errstack.NewReq("No (empty) user ID")
	}

	var accountStr string
	var query = `SELECT account_number
	FROM member_account
	WHERE individual_id = ? AND account_number IS NOT NULL
	  AND reference = 'primary' AND account_name = ?
	LIMIT 1`
	_, err := db.QueryOne(&accountStr, query, userID, distributionAccountName)
	if err := model.CheckPgNoRows("account_number", err); err != nil {
		return ethereum.ZeroAddress, err
	}

	// blank address means that user doesn't want to get SWC (yet)
	if accountStr == "" {
		return ethereum.ZeroAddress,
			errstack.NewReq("Blank ethereum account (user want's SB to hold his SWC)")
	}
	if !strings.HasPrefix(accountStr, "0x") {
		logger.Warn("Inconsistent Ethereum address: no '0x' prefix",
			"individual_id", userID, "address", accountStr)
		accountStr = "0x" + accountStr
	}

	return ethereum.ParseAddress(accountStr)
}

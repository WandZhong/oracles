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

package main

import (
	"bitbucket.org/sweetbridge/oracles/go-lib/crawlers"
	"fmt"
	"github.com/go-pg/pg"
	"github.com/robert-zaremba/errstack"
	"math/big"
	"strings"
	"time"
)

const (
	// The statement to add a contribution
	// the name of the function comes from the configuration and is quoted.
	uploadStmt = `SELECT "___SF_NAME___"(
		$1::text,
		$2::timestamp,
		$3::numeric,
		$4::text,
		$5::text,
		$6::text,
		$7::text,
		$8::text,
		$9::text,
		$10::text,
		$11::text,
		$12::text,
		$13::text,
		$14::text,
		$15::text
	)`
)

var (
	// Used in comparing a big.Int to 0
	bigIntZero = big.NewInt(0)
)

// newContribDBLoader returns an initalized contribDbLoader, using provided options
func newContribDBLoader(opts mainOpts) (crawlers.BCBlockProcessor, errstack.E) {
	db := opts.PgFlags.MustConnect()
	var stmtSQL = strings.Replace(uploadStmt, "___SF_NAME___", opts.contribDB.StoredFunc, 1)
	stmt, err := db.Prepare(stmtSQL)
	if err != nil {
		return nil, errstack.WrapAsDomain(err, "Stmt preparation failed")
	}
	return &contribDbLoader{stmt}, nil
}

// A contriution record
type contribDetail struct {
	Sources       string
	Timestamp     time.Time
	Amount        float64
	AccountType   string
	AccountSource string
	FromAddress   string
	ToAddress     string
	TransRef      string
	Curr          string
}

// A wrapper around the DB Stored function for uploading transactions in DB
type contribDbLoader struct {
	stmt *pg.Stmt
}

// Process is the implementation of BCBlockProcessor.Process cf BCBlockProcessor
// Only the transactions are processed. Logs are ignored.
// The conditions to include a transaction are:
// - both To and From address must be valid non nil addresses
// - value of the transaction must be different from O
func (r *contribDbLoader) Process(block *crawlers.BCBlock) errstack.E {
	for _, t := range block.Transactions {
		if t.To == "" && t.From == "" {
			continue
		}
		if t.RawValue.Cmp(bigIntZero) == 0 {
			continue
		}

		amount, _ := t.Value.Float64()
		contrib := &contribDetail{
			Sources:       fmt.Sprintf("ethereum://%d/%s", block.NetworkID, t.TxHash),
			Timestamp:     block.Timestamp,
			Amount:        amount,
			AccountType:   "Ether",
			AccountSource: fmt.Sprintf("NetworkId-%d", block.NetworkID),
			FromAddress:   t.From,
			ToAddress:     t.To,
			TransRef:      t.TxHash,
			Curr:          "ETH",
		}

		if err := r.sendData(contrib); err != nil {
			return err
		}

	}
	return nil
}

// actual stored procedure invocation
func (r *contribDbLoader) sendData(c *contribDetail) errstack.E {
	logger.Debug(fmt.Sprintf("Ref:%q, TimeStamp:%s, Value:%f, From:%q, To:%q\n", c.TransRef, c.Timestamp.UTC().String(), c.Amount, c.FromAddress, c.ToAddress))
	_, err := r.stmt.ExecOne(
		c.Sources,
		c.Timestamp,
		c.Amount,
		c.AccountType,
		c.AccountSource,
		c.FromAddress,
		"N/A",
		nil,
		c.AccountType,
		c.AccountSource,
		c.ToAddress,
		"N/A",
		nil,
		c.TransRef,
		c.Curr)
	return errstack.WrapAsDomain(err, "Database writing failed")
}

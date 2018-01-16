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
	"fmt"
	"io"
	"strings"
	"time"

	"bitbucket.org/sweetbridge/oracles/go-lib/encoding"
	"bitbucket.org/sweetbridge/oracles/go-lib/liquidity"
	"github.com/robert-zaremba/errstack"
	bat "github.com/robert-zaremba/go-bat"
)

// Record represent SWC distribution record
type Record struct {
	RowNum    int
	Timestamp time.Time
	Email     string
	FullName  string
	SenderID  string
	Amount    float64
	Currency  liquidity.Currency
}

// String implements stringer interface
func (r Record) String() string {
	return fmt.Sprint("{", r.FullName, " <", r.Email, "> ", r.Amount, " ", r.Currency, " ", r.Timestamp, "}")
}

func readRecords(fname string) ([]Record, errstack.E) {
	reader, fclose, errS := encoding.NewCSVFileReader(fname, 2)
	if errS != nil {
		return nil, errS
	}
	defer errstack.CallAndLog(logger, fclose)

	var records []Record
	var errb = errstack.NewBuilder()
	for i := 3; ; i++ {
		errbRow := errb.ForkIdx(i)
		row, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			errbRow.Put("row", "Can't read the row. "+err.Error())
			continue
		}
		if row[0] == "" && row[1] == "" && row[2] == "" {
			logger.Info("Columns 1,2,3 empty. Skipping further reading", "current_row", i)
			break
		}
		var r = Record{RowNum: i, Email: row[3], FullName: row[4], SenderID: row[17]}
		r.Currency = liquidity.ParseCurrencyErrp(row[7], errbRow.Putter("currency"))
		r.Amount = readAmount(row[8], errbRow.Putter("amount"))
		r.Timestamp, err = time.Parse("1/2/2006", row[0])
		if err != nil {
			errbRow.Put("payment_date", err.Error())
		}
		r.Timestamp = r.Timestamp.UTC()
		records = append(records, r)
	}

	return records, errb.ToReqErr()
}

func readAmount(src string, errp errstack.Putter) float64 {
	// Firstly let's fix numbers with thousend's separator, eg: "41'000"
	src = strings.Replace(src, "'", "", -1)
	src = strings.Replace(src, ",", "", -1)
	v, err := bat.Atof64(src)
	if err != nil {
		errp.Put(err)
	} else if v <= 0 {
		errp.Put("must be positive: " + src)
	}
	return v
}

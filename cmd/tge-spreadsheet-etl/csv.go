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

	"bitbucket.org/sweetbridge/oracles/go-lib/directbuy"
	"bitbucket.org/sweetbridge/oracles/go-lib/encoding"
	"bitbucket.org/sweetbridge/oracles/go-lib/liquidity"
	"github.com/robert-zaremba/errstack"
	bat "github.com/robert-zaremba/go-bat"
)

const (
	rowDate = iota
	_       // rowEmailConf
	rowTranche
	_ // rowEscrow
	rowEmail
	rowFullName
	_ // rowLastName
	_ // rowFristName
	rowCurrency
	rowAmount
	rowFXRate
	_ // rowAmountUSD
	_ // rowSWCPrice
	rowSWCAmount
	_ // rowMarketValue
	_ // rowCryptoExRate
	_ // rowSWCStatus
	_ // rowMatchedAddr
	_ // rowEmpty1
	rowSenderID
	rowTxHash
	_ // rowDate2
	rowID
)

// Record represent TGE spreadsheet record
type Record struct {
	directbuy.DirectBuy
	TxHash   string
	FullName string
}

func read(fname string) ([]Record, errstack.E) {
	reader, fclose, errS := encoding.NewCSVFileReader(fname, 2)
	if errS != nil {
		return nil, errS
	}
	defer errstack.CallAndLog(logger, fclose)

	if rowID != 22 || rowCurrency != 8 {
		logger.Fatal("CSV file columns has changed")
	}

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
		if len(row) < 3 || row[0] == "" && row[1] == "" && row[2] == "" {
			logger.Info("Columns 1,2,3 empty. Skipping further reading", "current_row", i)
			break
		}
		if len(row) < rowID-1 {
			errbRow.Put("row", fmt.Sprintf("Expecting %d columns, got %d", rowID, len(row)))
			continue
		}

		if row[0] == "???" {
			continue
		}
		var r = Record{
			DirectBuy: directbuy.DirectBuy{
				Email:    row[rowEmail],
				SenderID: row[rowSenderID],
			},
			FullName: row[rowFullName],
			TxHash:   row[rowTxHash]}
		r.ID = bat.Atoi64Errp(row[rowID], errbRow.Putter("id"))
		r.Currency = liquidity.ParseCurrencyErrp(row[rowCurrency], errbRow.Putter("currency"))
		r.AmountIn = readAmount(row[rowAmount], errbRow.Putter("amount_in"))
		r.AmountOut = readAmount(row[rowSWCAmount], errbRow.Putter("amount_swc"))
		r.UsdRate = readAmount(row[rowFXRate], errbRow.Putter("fx_rate"))
		errp := errbRow.Putter("swc_price")
		r.TrancheID = uint64(bat.Atoi64Errp(row[rowTranche], errp))
		if r.TrancheID < 1 {
			errp.Put("tranche must be >= 1")
		}
		r.TransactionDate, err = time.Parse("1/2/2006", row[rowDate])
		if err != nil {
			errbRow.Put("payment_date", err.Error())
		}
		r.TransactionDate = r.TransactionDate.UTC()
		records = append(records, r)
	}

	return records, errb.ToReqErr()
}

func readAmount(src string, errp errstack.Putter) float64 {
	// Firstly let's fix numbers with thousend's separator, eg: "41'000"
	src = strings.Replace(src, "'", "", -1)
	src = strings.Replace(src, ",", "", -1)
	src = strings.TrimSpace(src)
	v, err := bat.Atof64(src)
	if err != nil {
		errp.Put(err)
	} else if v <= 0 {
		errp.Put("must be positive: " + src)
	}
	return v
}

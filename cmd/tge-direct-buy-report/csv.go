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
	"encoding/csv"
	"fmt"
	"math/big"
	"os"

	"bitbucket.org/sweetbridge/oracles/go-lib/directbuy"
	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum/wad"
	"github.com/ethereum/go-ethereum/common"
	"github.com/robert-zaremba/errstack"
)

func getDirectBuys(trancheID string) ([]directbuy.DirectBuy, errstack.E) {
	var ds = []directbuy.DirectBuy{}
	err := db.Model(&ds).
		Where("tranche_id = ?", trancheID).
		Select()
	return ds, errstack.WrapAsInf(err, "Can't get DirectBuy records")
}

func createDirectBuyReport(trancheID, filename string) errstack.E {
	ds, errS := getDirectBuys(trancheID)
	if errS != nil {
		return errS
	}
	logger.Debug("Fetched direct buys", "tranche_id", trancheID, "num_records", len(ds))

	fout, err := os.Create(filename)
	if err != nil {
		return errstack.WrapAsInf(err, "Can't create file for direct buy report")
	}
	defer errstack.CallAndLog(logger, fout.Close)
	w := csv.NewWriter(fout)

	// firstly let's aggregate records
	var totals = map[string]*SWCRecord{}
	for _, d := range ds {
		// if d.UserID.Empty() {
		// 	logger.Info("swc-distribution report. Ignoring DirectBuy - no usesr ID", "id", d.ID)
		// 	continue
		// }
		fmt.Println("adding", d.Email)
		if sr, ok := totals[d.UserID.String()]; !ok {
			totals[d.UserID.String()] = &SWCRecord{
				common.Address{}, // TODO
				wad.FToWei(d.AmountOut)}
		} else {
			sr.Amount.Add(sr.Amount, wad.FToWei(d.AmountOut))
		}
	}

	logger.Info("Writing direct buy report", "filename", filename)
	// now we create the report
	for _, sr := range totals {
		if err = w.Write([]string{"TGE1", sr.Address.Hex(), sr.Amount.String(), "", "false"}); err != nil {
			return errstack.WrapAsInf(err, "Can't write row into direct buy report")
		}
	}
	w.Flush()
	return nil
}

// SWCRecord represent SWC distribution record
// aligned to swc-distributor
type SWCRecord struct {
	Address common.Address
	Amount  *big.Int
}

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
	"os"
	"strings"

	"bitbucket.org/sweetbridge/oracles/go-lib/directbuy"
	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"bitbucket.org/sweetbridge/oracles/go-lib/model"
	"github.com/ethereum/go-ethereum/common"
	"github.com/robert-zaremba/errstack"
	"github.com/robert-zaremba/go-bat"
	"github.com/robert-zaremba/go-pgt"
)

func getDirectBuys(trancheID string) ([]directbuy.DirectBuy, errstack.E) {
	var ds = []directbuy.DirectBuy{}
	err := db.Model(&ds).
		Where("tranche_id = ? AND status < ?", trancheID, directbuy.StatusDone).
		Select()
	return ds, errstack.WrapAsInf(err, "Can't get DirectBuy records")
}

func queryDistributionAccount(userID pgt.UUID) (string, error) {
	var account string
	var query = `SELECT account_number
	FROM member_account
	WHERE individual_id = ? AND account_number IS NOT NULL
	  AND reference = 'primary' AND account_name = 'Sweetcoin Distribution'
	LIMIT 1`
	_, err := db.QueryOne(&account, query, userID, directbuy.StatusDone)
	return account, err
}

func findAddress(d *directbuy.DirectBuy) (common.Address, errstack.E) {
	if d.UserID.Empty() {
		logger.Warn("DirectBuy with unmatched individual ID",
			"direct_buy_id", d.ID, "user_email", d.Email)
		return ethereum.ZeroAddress, nil
	}
	addressStr, err := queryDistributionAccount(d.UserID)
	if err != nil {
		if err := model.ErrNotNoRows("account_number", err); err != nil {
			return ethereum.ZeroAddress, err
		}
		logger.Warn("Can't find user ethereum account",
			"email", d.Email, "individual_id", d.UserID)
		return ethereum.ZeroAddress, nil
	}
	if addressStr == "" {
		// blank address means that user doesn't want to get SWC (yet)
		logger.Warn("Blank ethereum account (user want's SB to hold his SWC)",
			"email", d.Email, "individual_id", d.UserID)
		return ethereum.ZeroAddress, nil
	}
	if !strings.HasPrefix(addressStr, "0x") {
		logger.Warn("Inconsistent Ethereum address: no '0x' prefix",
			"email", d.Email, "individual_id", d.UserID, "address", addressStr)
		addressStr = "0x" + addressStr
	}
	address, err := ethereum.ParseAddress(addressStr)
	if err != nil {
		logger.Error("Invalid ethereum address", err,
			"email", d.Email, "individual_id", d.UserID, "address", addressStr)
	}
	return address, nil
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

	logger.Debug("Aggregating directe buy", "tranche_id", trancheID)
	if *flags.setPendingStatus {
		logger.Info("Setting 'pending' status for matched direct buy", "tranche", trancheID)
	} else {
		logger.Info("NOT updating statuses for matched direct buys", "tranche", trancheID)
	}
	var totals = map[string]*SWCRecord{}
	for _, d := range ds {
		if d.Status == directbuy.StatusDone {
			continue
		}
		address, err := findAddress(&d)
		if err != nil {
			return err
		}
		if ethereum.IsZeroAddr(address) {
			continue
		}
		if sr, ok := totals[d.UserID.String()]; !ok {
			totals[d.UserID.String()] = &SWCRecord{
				address,
				d.AmountOut,
				fmt.Sprintf("email: %s; id: %v", d.Email, d.UserID)}
		} else {
			sr.Amount += d.AmountOut
		}
		if *flags.setPendingStatus {
			d.Status = directbuy.StatusPending
			if err := directbuy.UpdateStatus(d.ID, d.Status, db); err != nil {
				fmt.Println(">>>", err)
				return err
			}
		}
	}

	logger.Debug("Writing direct buy report", "filename", filename)
	if err = w.Write([]string{"list", "address", "amount", "comment", "done"}); err != nil {
		return errstack.WrapAsInf(err, "Can't write header into direct buy report")
	}
	listName := "TGE_" + trancheID
	for _, sr := range totals {
		if err = w.Write([]string{listName, sr.Address.Hex(), bat.F64toa(sr.Amount), sr.Comment, "false"}); err != nil {
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
	Amount  float64
	Comment string
}

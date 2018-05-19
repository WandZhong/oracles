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
	"flag"
	"os"

	"bitbucket.org/sweetbridge/oracles/go-lib/directbuy"
	"bitbucket.org/sweetbridge/oracles/go-lib/log"
	"bitbucket.org/sweetbridge/oracles/go-lib/setup"
	"github.com/go-pg/pg"
	"github.com/robert-zaremba/errstack"
	bat "github.com/robert-zaremba/go-bat"
)

var (
	db     *pg.DB
	logger = log.Root()
	flags  = mainFlags{PgFlags: setup.NewPgFlags(),
		tranche: new(uint64)}
	// setPendingStatus: flag.Bool("set-pending-status", false,
	// 	"set matched not completed directbuy statuses to 'pending'")}
)

type mainFlags struct {
	PgFlags setup.PgFlags
	tranche *uint64
}

func (f mainFlags) Check() error {
	var err error
	*f.tranche, err = bat.Atoui64(flag.Arg(0))
	if err != nil || *f.tranche == 0 {
		return errstack.NewReq("Tranche ID must be specified (not 0)")
	}
	return f.PgFlags.Check()
}

func init() {
	setup.FlagSimpleInit("tge-spreadsheet-etl", "tranche_id report_out_filename.csv",
		nil, flags)
	db = flags.PgFlags.MustConnect()
}

func main() {
	if err := createDirectBuyReport(*flags.tranche, flag.Arg(1)); err != nil {
		logger.Error("Can't build report", err)
	}
}

func createDirectBuyReport(trancheID uint64, filename string) errstack.E {
	logger.Debug("Writing direct buy report", "filename", filename)
	fout, errStd := os.Create(filename)
	if errStd != nil {
		return errstack.WrapAsInf(errStd, "Can't create file for direct buy report")
	}
	defer errstack.CallAndLog(logger, fout.Close)
	w := csv.NewWriter(fout)
	if err := w.Write([]string{"list", "address", "amount", "comment", "done"}); err != nil {
		return errstack.WrapAsInf(err, "Can't write header into direct buy report")
	}

	logger.Info("Fetching pending direct buys", "tranche", trancheID)
	records, err := directbuy.CreateDirectBuyReport("TGE", trancheID, db)
	if err != nil {
		return err
	}

	for _, sr := range records {
		err := w.Write([]string{sr.List, sr.Address.Hex(), bat.F64toa(sr.AmountOut), sr.Comment, "false"})
		if err != nil {
			return errstack.WrapAsInf(err, "Can't write row into direct buy report")
		}
	}
	w.Flush()
	return nil
}

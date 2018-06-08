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
	"crypto/md5"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"

	"bitbucket.org/sweetbridge/oracles/go-lib/directbuy"
	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum/wad"
	"github.com/robert-zaremba/errstack"
	bat "github.com/robert-zaremba/go-bat"
)

func readRecords(fname string) ([]directbuy.Summary, errstack.E) {
	f, err := os.Open(fname)
	if err != nil {
		return nil, errstack.WrapAsReq(err, "Can't open csv file: "+fname)
	}
	defer errstack.CallAndLog(logger, f.Close)
	var hasher = md5.New()
	var reader = csv.NewReader(io.TeeReader(f, hasher))
	reader.Comment = '#'

	header, err := reader.Read()
	if err != nil {
		return nil, errstack.NewReq("Can't read CSV header. " + err.Error())
	}
	expectedHeader := []string{"list", "address", "amount", "comment", "done"}
	if !bat.StrsEq(header, expectedHeader) {
		return nil, errstack.NewReqF("CSV header don't match. Expecting: %v",
			expectedHeader)
	}

	var records []directbuy.Summary
	var notDone = []string{"", "0", "no", "false"}
	var maxSWC = wad.ToWei(*flags.maxSWC)
	var errb = errstack.NewBuilder()
	for i := 2; ; i++ {
		errbRow := errb.ForkIdx(i)
		row, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			errbRow.Put("row", "Can't read the row. "+err.Error())
			continue
		}
		if bat.StrSliceIdx(notDone, strings.ToLower(row[4])) < 0 {
			logger.Info("Ignoring done row", "row", i)
			continue
		}
		var r = directbuy.Summary{List: row[0], Idx: i}
		r.Address = ethereum.ParseAddressErrp(row[1], errbRow.Putter("address"))
		errpAmount := errbRow.Putter("amount")
		r.Amount = wad.AfToNotNegWei(row[2], errpAmount)
		if r.Amount != nil && r.Amount.Cmp(maxSWC) > 0 {
			errpAmount.Put(fmt.Sprint("must be smaller or equal then ", *flags.maxSWC))
		}
		records = append(records, r)
	}

	checkControlSum(hex.EncodeToString(hasher.Sum(nil)), errb)
	return records, validate(records, errb)
}

func checkControlSum(controlSum string, errb errstack.Builder) {
	if *flags.expectedMd5 == "" {
		logger.Debug("Control sum not specified. Input file md5sum", "hash", controlSum)
	} else if *flags.expectedMd5 != controlSum {
		errb.Put("hash", "Input file doesn't match the control sum. Computed md5="+controlSum)
	}
}

func validate(rs []directbuy.Summary, errb errstack.Builder) errstack.E {
	// var dup = 0
	// var dupMap = map[common.Address]int{}
	for _, r := range rs {
		errbRow := errb.ForkIdx(r.Idx)
		if len(r.List) < 1 {
			errbRow.Put("list", "List name can't be empty")
		}
		// if dup = dupMap[r.Address]; dup > 0 {
		// 	errbRow.Put("duplication", "address already used in row "+strconv.Itoa(dup))
		// }
		// dupMap[r.Address] = r.Idx
	}
	return errb.ToReqErr()
}

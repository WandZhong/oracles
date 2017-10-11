package main

import (
	"crypto/md5"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"io"
	"math/big"
	"os"
	"strconv"
	"strings"

	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum/wad"
	"github.com/ethereum/go-ethereum/common"
	"github.com/robert-zaremba/errstack"
	bat "github.com/robert-zaremba/go-bat"
)

// Record represent SWC distribution record
type Record struct {
	List    string
	Address common.Address
	Amount  *big.Int
	Idx     int
}

// String implements stringer interface
func (r Record) String() string {
	return fmt.Sprint("{", r.List, " ", r.Address.Hex(), " ", r.Amount, "}")
}

func readRecords(fname string) ([]Record, errstack.E) {
	f, err := os.Open(fname)
	defer errstack.CallAndLog(logger, f.Close)
	if err != nil {
		return nil, errstack.WrapAsReq(err, "Can't open csv file: "+fname)
	}
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

	var records []Record
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
		var r = Record{List: row[0], Idx: i}
		r.Address = ethereum.ToAddressErrp(row[1], errbRow.Putter("address"))
		errpAmount := errbRow.Putter("amount")
		r.Amount = wad.AfToPosWei(row[2], errpAmount)
		if r.Amount.Cmp(maxSWC) > 0 {
			errpAmount.Put(fmt.Sprint("must be bigger or equal then ", *flags.maxSWC))
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

func validate(rs []Record, errb errstack.Builder) errstack.E {
	var dup = 0
	var dupMap = map[common.Address]int{}
	for _, r := range rs {
		errbRow := errb.ForkIdx(r.Idx)
		if len(r.List) < 3 {
			errbRow.Put("list", "List name should be at least 3 character long")
		}
		if dup = dupMap[r.Address]; dup > 0 {
			errbRow.Put("duplication", "address already used in row "+strconv.Itoa(dup))
		}
		dupMap[r.Address] = r.Idx
	}
	return errb.ToReqErr()
}

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
		logger.Fatal("Can't open csv file", "filename", fname, err)
	}
	var hasher = md5.New()
	var reader = csv.NewReader(io.TeeReader(f, hasher))
	reader.Comment = '#'

	var records []Record
	var notDone = []string{"", "0", "no", "false"}
	header, err := reader.Read() // skip the header
	if err != nil {
		return records, errstack.NewReq("Can't read CSV header. " + err.Error())
	}
	expectedHeader := []string{"list", "address", "amount", "comment", "done"}
	if !bat.StrsEq(header, expectedHeader) {
		return records, errstack.NewReqF("CSV header don't match. Expecting: %v",
			expectedHeader)
	}

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

	md5sum := hex.EncodeToString(hasher.Sum(nil))
	if *flags.expectedMd5 == "" {
		logger.Debug("Control sum not specified. Input file md5sum", "hash", md5sum)
	} else if *flags.expectedMd5 != md5sum {
		errb.Put("hash", "Input file doesn't match the control sum. Computed md5="+md5sum)
	}
	return records, errb.ToReqErr()
}

func validate(rs []Record) errstack.E {
	var dup = 0
	var dupMap = map[common.Address]int{}
	var errb = errstack.NewBuilder()
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

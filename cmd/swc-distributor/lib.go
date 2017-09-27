package main

import (
	"crypto/md5"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"

	"math/big"

	contracts "bitbucket.org/sweetbridge/oracles/go-contracts"
	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"bitbucket.org/sweetbridge/oracles/go-lib/setup"
	"bitbucket.org/sweetbridge/oracles/go-lib/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/robert-zaremba/errstack"
	bat "github.com/robert-zaremba/go-bat"
)

// Record represent SWC distribution record
type Record struct {
	List    string
	Address common.Address
	Amount  int64
	Idx     int
}

// String implements stringer interface
func (r Record) String() string {
	return fmt.Sprint("{", r.List, " ", r.Address.Hex(), " ", r.Amount, "}")
}

func readRecords(fname string) ([]Record, bool) {
	f, err := os.Open(fname)
	defer errstack.CallAndLog(logger, f.Close)
	if err != nil {
		logger.Fatal("Can't open csv file", "filename", fname, err)
	}
	var hasher = md5.New()
	var reader = csv.NewReader(io.TeeReader(f, hasher))
	reader.Comment = '#'

	var records []Record
	var row []string
	var ok = true
	var notDone = []string{"", "0", "no", "false"}
	if _, err = reader.Read(); err != nil { // skip the header
		logger.Fatal("Can't read CSV header", err)
	}
	for i := 2; ; i++ {
		if row, err = reader.Read(); err == io.EOF {
			break
		} else if err != nil {
			logger.Fatal("Can't read the row", "row", i, err)
		} else if len(row) < 5 {
			logger.Error("Not enough columns in the row. Required >= 5", "cols", len(row))
			ok = false
		}
		var r = Record{List: row[0], Idx: i}
		if r.Address, err = ethereum.ToAddress(row[1]); err != nil {
			logger.Error("Malformed address", "row", i, "value", row[1])
			ok = false
		}
		if r.Amount, err = bat.Atoi64(row[2]); err != nil {
			logger.Error("Wrong amount", "row", i, "value", row[1], err)
			ok = false
		}
		if bat.StrSliceIdx(notDone, strings.ToLower(row[4])) < 0 {
			logger.Info("Ignoring done row", "row", i)
			continue
		}
		records = append(records, r)
	}

	md5sum := hex.EncodeToString(hasher.Sum(nil))
	if *flags.expectedMd5 == "" {
		logger.Debug("Control sum not specified. Input file md5sum", "hash", md5sum)
	} else if *flags.expectedMd5 != md5sum {
		logger.Error("Input file doesn't match the control sum", "computed_md5", md5sum)
		ok = false
	}
	return records, ok
}

func validate(rs []Record) bool {
	var ok = true
	var dup = 0
	var dupMap = map[common.Address]int{}
	for _, r := range rs {
		if len(r.List) < 3 {
			logger.Error("List name should be at least 3 character long", "row", r.Idx)
			ok = false
		}
		if r.Amount < 0 {
			logger.Error("Amount can't be negative", "row", r.Idx)
			ok = false
		}
		if dup = dupMap[r.Address]; dup > 0 {
			logger.Error("Detected duplicated rows (addresses)", "row1", dup, "row2", r.Idx)
			ok = false
		}
		dupMap[r.Address] = r.Idx
	}
	return ok
}

func transferSWC(records []Record) {
	client := setup.EthClient(*flags.Network)
	cf := ethereum.MustNewContractFactorySF(client, *flags.ContractsPath, *flags.Network)
	swcC, addr, err := cf.GetSWC()
	utils.Assert(err, "Can't instantiate SWT contract")
	logger.Debug("Contract address", "swc", addr.Hex())
	checkSWCbalance(records, swcC)
	txo := flags.MustNewTxrFactory().Txo()
	for _, r := range records {
		logger.Info(fmt.Sprint("Transferring: ", r))
		tx, err := swcC.Transfer(txo, r.Address, big.NewInt(r.Amount))
		if err != nil {
			logger.Error("Can't transfer TOKEN", err)
		} else {
			ethereum.LogTx("Transferred", tx)
			ethereum.IncTxoNonce(txo, tx)
		}
	}
}

func checkSWCbalance(records []Record, token *contracts.SweetToken) {
	var total int64
	for _, r := range records {
		total += r.Amount
	}
	k := ethereum.MustReadKeySimple(*flags.PkFile)
	logger.Debug("Coinbase", "address", k.Address)
	b, err := token.BalanceOf(nil, k.Address)
	if err != nil {
		logger.Fatal("Can't check SWC balance", err)
	}
	totalWei := ethereum.ToWei(total)
	if b.Cmp(totalWei) < 0 {
		bInt := ethereum.WeiToInt(b)
		logger.Fatal("Not enough funds in the source account", "min_expected", total, "has", bInt)
	}
	logger.Debug("Distribution account balance", "wei", totalWei.String())
}

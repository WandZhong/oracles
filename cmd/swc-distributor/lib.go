package main

import (
	"crypto/md5"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strconv"
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

func readRecords(fname string) ([]Record, errstack.Builder) {
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
	var notDone = []string{"", "0", "no", "false"}
	var errb = errstack.NewBuilder()
	header, err := reader.Read() // skip the header
	if err != nil {
		errb.Put("header", "Can't read CSV header. "+err.Error())
		return records, errb
	}
	fmt.Println("header ", header)
	for i := 2; ; i++ {
		errbRow := errb.ForkIdx(i)
		if row, err = reader.Read(); err == io.EOF {
			break
		} else if err != nil {
			logger.Error("Can't read the row", "row", i, err)
			os.Exit(1)
		} else if len(row) < 5 {
			errbRow.Put("row",
				"Not enough columns. Required >= 5, got: "+strconv.Itoa(len(row)))
		}
		var r = Record{List: row[0], Idx: i}
		r.Address = ethereum.ToAddressErrp(row[1], errbRow.Putter("address"))
		r.Amount = bat.Atoi64Errp(row[2], errbRow.Putter("amount"))
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
		errb.Put("hash", "Input file doesn't match the control sum. Computed md5="+md5sum)
	}
	return records, errb
}

func validate(rs []Record) errstack.Builder {
	var dup = 0
	var dupMap = map[common.Address]int{}
	var errb = errstack.NewBuilder()
	for _, r := range rs {
		errbRow := errb.ForkIdx(r.Idx)
		if len(r.List) < 3 {
			errbRow.Put("list", "List name should be at least 3 character long")
		}
		if r.Amount < 0 {
			errbRow.Put("amount", "Can't be negative")
		}
		if dup = dupMap[r.Address]; dup > 0 {
			errbRow.Put("duplication", "address already used in row "+strconv.Itoa(dup))
		}
		dupMap[r.Address] = r.Idx
	}
	return errb
}

func transferSWC(records []Record) {
	_, cf := setup.MustEthClient(*flags.Network, *flags.ContractsPath)
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

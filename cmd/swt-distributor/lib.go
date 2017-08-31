package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"math/big"

	contracts "bitbucket.org/sweetbridge/oracles/go-contracts"
	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"bitbucket.org/sweetbridge/oracles/go-lib/setup"
	"github.com/ethereum/go-ethereum/common"
	bat "github.com/robert-zaremba/go-bat"
)

// Record represent SWT distribution record
type Record struct {
	List    string
	Address common.Address
	Amount  int64
}

// String implements stringer interface
func (r Record) String() string {
	return fmt.Sprint("{", r.List, " ", r.Address.Hex(), " ", r.Amount, "}")
}

func readRecords(fname string) ([]Record, bool) {
	f, err := os.Open(fname)
	if err != nil {
		logger.Fatal("Can't open csv file", "filename", fname, err)
	}
	reader := csv.NewReader(f)
	var records []Record
	var row []string
	var ok = true
	if _, err = reader.Read(); err != nil { // skip the header
		logger.Fatal("Can't read CSV header", err)
	}
	for i := 2; ; i++ {
		if row, err = reader.Read(); err == io.EOF {
			break
		} else if err != nil {
			logger.Fatal("Can't read the row", "row", i, err)
		} else if len(row) < 3 {
			logger.Error("Not enough columns in the row. Required >= 3", "cols", len(row))
			ok = false
		}
		var r = Record{List: row[0]}
		if r.Address, err = ethereum.ToAddress(row[1]); err != nil {
			logger.Error("Malformed address", "row", i, "value", row[1], err)
			ok = false
		}
		if r.Amount, err = bat.Atoi64(row[2]); err != nil {
			logger.Error("Wrong amount", "row", i, "value", row[1], err)
			ok = false
		}
		records = append(records, r)
	}
	return records, ok
}

func validate(rs []Record) bool {
	var ok = true
	var i = 1
	for _, r := range rs {
		i++
		if len(r.List) < 3 {
			logger.Error("List name should be at least 3 character long", "row", i)
			ok = false
		}
		if r.Amount < 0 {
			logger.Error("Amount can't be negative", "row", i)
			ok = false
		}
	}
	return ok
}

func transferSWT(records []Record, swtAddr common.Address) {
	client := setup.EthClient(*ethHost)
	token, err := contracts.NewSweetToken(swtAddr, client)
	if err != nil {
		logger.Fatal("Can't obtain TOKEN contract", err)
	}
	txo := ethereum.MustFileTxr(*pkFile, *pkPwd)
	for _, r := range records {
		logger.Info(fmt.Sprint("Transferring: ", r))
		tx, err := token.Transfer(txo, r.Address, big.NewInt(r.Amount))
		if err != nil {
			logger.Error("Can't transfer TOKEN", err)
		} else {
			logger.Debug("Transferred", "gas", tx.Gas(), "gas_price", tx.GasPrice())
			fmt.Println(txo.Nonce, tx.Nonce())
			if txo.Nonce == nil {
				txo.Nonce = big.NewInt(int64(tx.Nonce()))
			}
			ethereum.IncNonce(txo.Nonce)
		}
	}
}

package ethereum

import (
	"fmt"
	"io"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/robert-zaremba/errstack"
)

// LogTx is a handy function to create debug log for successful transaction
func LogTx(msg string, tx *types.Transaction) {
	if tx != nil {
		logger.Debug(msg, "tx_hash", tx.Hash().Hex(), "nonce", tx.Nonce(),
			"gas", tx.Gas(), "gas_price", tx.GasPrice())
	} else {
		logger.Debug("Invalid transaction")
	}
}

// FlogTx logs transaction into a Writer
func FlogTx(w io.Writer, msg string, tx *types.Transaction) {
	if tx != nil {
		fmt.Fprintf(w, "%s\n\thash=%s, gas=%v, gas_price=%v\n",
			msg, tx.Hash().Hex(), tx.Gas(), tx.GasPrice())
	} else {
		_, err := w.Write([]byte(msg + ": invalid transaction\n"))
		errstack.Log(logger, err)
	}
}

package ethereum

import "github.com/ethereum/go-ethereum/core/types"

// LogTx is a handy function to create debug log for successful transaction
func LogTx(msg string, tx *types.Transaction) {
	if tx != nil {
		logger.Debug(msg, "hash", tx.Hash().Hex(), "gas", tx.Gas(), "gas_price", tx.GasPrice())
	} else {
		logger.Debug("Invalid transaction")
	}
}

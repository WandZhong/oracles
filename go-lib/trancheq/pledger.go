package trancheq

import (
	"math/big"

	contracts "bitbucket.org/sweetbridge/oracles/go-contracts"
	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"bitbucket.org/sweetbridge/oracles/go-lib/liquidity"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/robert-zaremba/errstack"
)

// Pledger provides functionality for creating pledges
type Pledger struct {
	BRG      *contracts.BridgeToken
	SWCQ     *contracts.SWCqueue
	SWCQaddr common.Address
	cf       ethereum.ContractFactory
}

// NewPledger constructs Pledger
func NewPledger(cf ethereum.ContractFactory) (Pledger, errstack.E) {
	brgC, brgAddr, err := cf.GetBRG()
	if err != nil {
		return Pledger{}, err
	}
	swcqC, swcqAddr, err := cf.GetSWCqueue()
	if err != nil {
		return Pledger{}, err
	}
	logger.Debug("Pledger contract addresses:", "brg", brgAddr.Hex(), "swcq", swcqAddr.Hex())
	return Pledger{brgC, swcqC, swcqAddr, cf}, nil
}

// Post mints new BRG for the SWCqueue and posts the pledge request to the blockchain
// It reteurns transaction objects for Mint and for DirectPledge respectively.
func (plr Pledger) Post(dst common.Address, wei *big.Int, curr liquidity.Currency) (
	*types.Transaction, *types.Transaction, error) {

	txo := plr.cf.Txo()
	txMint, err := plr.BRG.MintFor(txo, plr.SWCQaddr, wei)
	if err != nil {
		logger.Error("Can't mint BRG", err)
		return nil, nil, err
	}
	ethereum.LogTx("BRG minted for SWCq", txMint)
	ethereum.IncTxoNonce(txo, txMint)
	txPledge, err := plr.SWCQ.DirectPledge(txo, dst, wei, curr)
	if err != nil {
		logger.Error("Can't pledge", err)
		return nil, nil, err
	}
	ethereum.LogTx("SWCq direct pledge log recorded", txPledge)
	return txMint, txPledge, nil
}

package trancheq

import (
	"time"

	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"bitbucket.org/sweetbridge/oracles/go-lib/liquidity"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/robert-zaremba/errstack"
	pgt "github.com/robert-zaremba/go-pgt"
)

// Pledge represents swc_queue entry
type Pledge struct {
	tableName struct{} `sql:"pledge_queue"`

	Tx         string              `sql:"pledge_queue_id,pk"` // ID is a Tx hash
	WalletAddr ethereum.PgtAddress `sql:"wallet_addr"`
	CtrAddr    ethereum.PgtAddress `sql:"ctr_addr,notnull"`
	Wad        pgt.BigInt          `sql:",notnull"`
	CreatedOn  time.Time           `sql:"created_on"`
	Currency   liquidity.Currency  `sql:",notnull"`
	Direct     bool                `sql:",notnull"`
}

// NewPledgeFromDirectPledge constructs new Pledge from the event
func NewPledgeFromDirectPledge(log *types.Log) (Pledge, errstack.E) {
	var e EventDirectPledge
	if err := e.Unmarshal(log); err != nil {
		return Pledge{}, err
	}
	return Pledge{
		Tx:         log.TxHash.Hex(),
		WalletAddr: ethereum.PgtAddress{Address: e.Who},
		CtrAddr:    ethereum.PgtAddress{Address: log.Address},
		Wad:        pgt.BigInt{Int: e.Wad},
		Currency:   e.Currency,
		CreatedOn:  time.Now(),
		Direct:     true,
	}, nil
}

// NewPledgeFromTransfer constructs new Pledge from the event
func NewPledgeFromTransfer(log *types.Log) (Pledge, errstack.E) {
	var e liquidity.EventTokenTransfer
	if err := e.Unmarshal(log); err != nil {
		return Pledge{}, err
	}
	return Pledge{
		Tx:         log.TxHash.Hex(),
		WalletAddr: ethereum.PgtAddress{Address: e.From},
		CtrAddr:    ethereum.PgtAddress{Address: e.To},
		Wad:        pgt.BigInt{Int: e.Value},
		// TODO: find the right currency based on log.Address
		Currency:  liquidity.CurrUSD,
		CreatedOn: time.Now(),
		Direct:    false,
	}, nil
}

package swcq

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
	tableName struct{} `sql:"swc_queue"`

	Tx        string              `sql:"swc_queue_id,pk"` // ID is a Tx hash
	UserAddr  ethereum.PgtAddress `sql:"user_addr"`
	CtrAddr   ethereum.PgtAddress `sql:"ctr_addr,notnull"`
	Wad       pgt.BigInt          `sql:",notnull"`
	CreatedOn time.Time           `sql:"created_on"`
	Currency  liquidity.Currency  `sql:",notnull"`
	Direct    bool                `sql:",notnull"`
}

// NewPledgeFromDirectPledgeLog constructs new Pledge from the event
func NewPledgeFromDirectPledgeLog(log *types.Log) (Pledge, errstack.E) {
	var e EventDirectPledge
	if err := e.Unmarshal(log); err != nil {
		return Pledge{}, err
	}
	return Pledge{
		Tx:        log.TxHash.Hex(),
		UserAddr:  ethereum.PgtAddress{Address: e.Who},
		CtrAddr:   ethereum.PgtAddress{Address: log.Address},
		Wad:       pgt.BigInt{Int: e.Wad},
		Currency:  e.Currency,
		CreatedOn: time.Now(),
		Direct:    true,
	}, nil
}

// NewPledgeFromTransfer constructs new Pledge from the event
func NewPledgeFromTransfer(log *types.Log) (Pledge, errstack.E) {
	var e liquidity.EventTokenTransfer
	if err := e.Unmarshal(log); err != nil {
		return Pledge{}, err
	}
	return Pledge{
		Tx:       log.TxHash.Hex(),
		UserAddr: ethereum.PgtAddress{Address: e.From},
		CtrAddr:  ethereum.PgtAddress{Address: e.To},
		Wad:      pgt.BigInt{Int: e.Value},
		// TODO: find the right currency based on log.Address
		Currency:  liquidity.CurrUSD,
		CreatedOn: time.Now(),
		Direct:    false,
	}, nil
}

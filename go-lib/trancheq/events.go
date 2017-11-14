package trancheq

import (
	"math/big"

	contracts "bitbucket.org/sweetbridge/oracles/go-contracts"
	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"bitbucket.org/sweetbridge/oracles/go-lib/liquidity"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/robert-zaremba/errstack"
)

// SWCqueue globals
var (
	SWCqueueABI = ethereum.MustParseABI("SWCqueue", contracts.SWCqueueABI)
)

const (
	logSWCqueueDirectPledge = "LogSWCqueueDirectPledge"
)

func init() {
	ethereum.MustHaveEvents(SWCqueueABI, logSWCqueueDirectPledge)
}

// LogSWCqueueDirectPledge returns event
func LogSWCqueueDirectPledge() abi.Event { return SWCqueueABI.Events[logSWCqueueDirectPledge] }

// EventDirectPledge represents LogSWCqueueDirectPledge event payload
type EventDirectPledge struct {
	Who      common.Address
	Wad      *big.Int
	Currency liquidity.Currency
}

// Unmarshal blockchain log into the event structure
func (e *EventDirectPledge) Unmarshal(log *types.Log) errstack.E {
	return ethereum.UnmarshalEvent(e, log.Data, LogSWCqueueDirectPledge())
}

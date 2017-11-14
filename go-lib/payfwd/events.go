package payfwd

import (
	"bitbucket.org/sweetbridge/oracles/go-contracts"
	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"bitbucket.org/sweetbridge/oracles/go-lib/log"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/robert-zaremba/errstack"
)

var (
	ForwarderFactoryABI abi.ABI
	logger              = log.Root()
)

const (
	logForwarderCreated = "LogForwarderCreated"
)

func init() {
	ForwarderFactoryABI = ethereum.MustParseABI("ForwarderFactory", contracts.ForwarderFactoryABI)
}

func LogForwarderCreated() abi.Event {
	return ForwarderFactoryABI.Events[logForwarderCreated]
}

type EventForwarderCreated struct {
	Forwarder common.Address
}

// Unmarshal blockchain log into the event structure
func (e *EventForwarderCreated) Unmarshal(log types.Log) errstack.E {
	return ethereum.UnmarshalEvent(&e.Forwarder, log.Data, LogForwarderCreated())
}

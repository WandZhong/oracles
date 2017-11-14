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

// global events
var (
	ForwarderFactoryABI = ethereum.MustParseABI("ForwarderFactory", contracts.ForwarderFactoryABI)
	logger              = log.Root()
)

const (
	logForwarderCreated = "LogForwarderCreated"
)

func init() {
	ethereum.MustHaveEvents(ForwarderFactoryABI, logForwarderCreated)
}

// LogForwarderCreated returns event
func LogForwarderCreated() abi.Event {
	return ForwarderFactoryABI.Events[logForwarderCreated]
}

// EventForwarderCreated represents LogForwarderCreated event payload
type EventForwarderCreated struct {
	Forwarder common.Address
}

// Unmarshal blockchain log into the event structure
func (e *EventForwarderCreated) Unmarshal(log types.Log) errstack.E {
	return ethereum.UnmarshalEvent(&e.Forwarder, log.Data, LogForwarderCreated())
}

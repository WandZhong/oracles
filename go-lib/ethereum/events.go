package ethereum

import (
	"context"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/robert-zaremba/errstack"
)

// SubscribeSimple is a simple utility function to create events subscription
func SubscribeSimple(ctx context.Context,
	client *ethclient.Client,
	topics [][]common.Hash, addresses []common.Address) (<-chan types.Log, ethereum.Subscription, errstack.E) {
	query := ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   nil,
		Topics:    topics,
		Addresses: addresses}
	var events = make(chan types.Log, 5) // to quickler consume batch events
	s, err := client.SubscribeFilterLogs(ctx, query, events)
	return events, s, errstack.WrapAsDomain(err, "Can't create Ethereum Subscription")
}

// UnmarshalEvent blockchain log into the event structure
// `dest` must be a pointer to initialized structure
func UnmarshalEvent(dest interface{}, data []byte, e abi.Event) errstack.E {
	a := abi.ABI{Events: map[string]abi.Event{"e": e}}
	err := a.Unpack(dest, "e", data)
	return errstack.WrapAsInf(err,
		"Probably the ABI doesn't match with the contract version")
}

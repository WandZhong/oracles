package main

import (
	"context"

	"bitbucket.org/sweetbridge/oracles/go-lib/utils"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/robert-zaremba/log15"
)

// TODO filter topics ...
func listenPledge(stopChan <-chan struct{}) {
	query := ethereum.FilterQuery{
		FromBlock: nil,
		ToBlock:   nil,
		// TODO:  Topics [][]common.Hash
		Addresses: []common.Address{addrSWCq, addrBrg}}
	var logs = make(chan types.Log, 2) // TODO verify if we need buffered chan
	s, err := client.SubscribeFilterLogs(context.TODO(), query, logs)
	utils.Assert(err, "Can't subscribe for SWC direct pledge events")
	errChan := s.Err()
	for {
		select {
		case err := <-errChan:
			logger.Error("Logs subscription error", err)
			break
		case l := <-logs:
			logger.Info("new log", log15.Spew(l))
		case <-stopChan:
			logger.Info("Closing log consumption")
			break
		}
	}
}

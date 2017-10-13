package main

import (
	"context"

	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"bitbucket.org/sweetbridge/oracles/go-lib/swcq"
	"bitbucket.org/sweetbridge/oracles/go-lib/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/robert-zaremba/log15"
)

func listenPledge() {
	events, s, err := ethereum.SubscribeSimple(context.TODO(), client,
		[][]common.Hash{{swcq.TitleLogSWCqueueDirectPledge}},
		[]common.Address{addrSWCq})
	utils.Assert(err, "Can't subscribe for SWC direct pledge events")
	errChan := s.Err()
loop:
	for {
		select {
		case err := <-errChan:
			s.Unsubscribe()
			logger.Error("Events subscription error", err)
			break loop
		case e := <-events:
			logger.Info("new log", log15.Alone("event", e.String()))
		}
	}
}

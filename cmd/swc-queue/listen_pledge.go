package main

import (
	"context"

	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"bitbucket.org/sweetbridge/oracles/go-lib/swcq"
	"bitbucket.org/sweetbridge/oracles/go-lib/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/robert-zaremba/log15"
)

func listenPledge() {
	logger.Debug("Started listening on pledge event")
	logs, s, err := ethereum.SubscribeSimple(context.TODO(), client,
		[][]common.Hash{{swcq.LogSWCqueueDirectPledge().Id()}},
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
		case l := <-logs:
			logger.Debug("new log", log15.Alone("event", l.String()))
			createPledge(l)
		}
	}
}

func createPledge(log types.Log) {
	p, err := swcq.NewPledgeFromDirectPledgeLog(log)
	if err != nil {
		logger.Error("Can't construct pledge", err)
		return
	}
	if err := db.Insert(&p); err != nil {
		logger.Error("Can't insert pledge", err)
	} else {
		logger.Info("pledge inserted")
	}
}

func find() {
	var p swcq.Pledge
	_, err := db.QueryOne(&p, `SELECT * FROM swc_queue WHERE swc_queue_id = ?`,
		"0x0000000000000000000000000000000000000000000000000000000000000000")
	if err != nil {
		logger.Error("Can't find a pledge", err)
	} else {
		logger.Info("pledge found", log15.Spew(p))
	}
}

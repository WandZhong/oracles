package main

import (
	"context"
	"math/big"
	"time"

	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"bitbucket.org/sweetbridge/oracles/go-lib/swcq"
	"bitbucket.org/sweetbridge/oracles/go-lib/utils"
	"github.com/ethereum/go-ethereum/common"
	pgt "github.com/robert-zaremba/go-pgt"
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

func createPledge() {
	p := swcq.Pledge{
		ID:        pgt.RandomUUID(),
		UserID:    pgt.RandomUUID(),
		Wad:       pgt.BigInt{Int: big.NewInt(1000)},
		Currency:  "USD",
		CreatedOn: time.Now(),
		Direct:    false,
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
		"bc401de4-047c-4377-886d-dcb1709c130f")
	if err != nil {
		logger.Error("Can't find a pledge", err)
	} else {
		logger.Info("pledge found", log15.Spew(p))
	}
}

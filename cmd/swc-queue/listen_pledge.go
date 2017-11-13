package main

import (
	"context"

	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"bitbucket.org/sweetbridge/oracles/go-lib/liquidity"
	"bitbucket.org/sweetbridge/oracles/go-lib/swcq"
	"bitbucket.org/sweetbridge/oracles/go-lib/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/robert-zaremba/errstack"
	"github.com/robert-zaremba/log15"
)

type listenCallback func(swcq.Pledge) error

func listen(ctx context.Context, title string, topic common.Hash, address common.Address,
	parser func(*types.Log) (swcq.Pledge, errstack.E),
	callback listenCallback) {

	logger.Debug("Started listening on " + title)
	logs, s, err := ethereum.SubscribeSimple(ctx, client,
		[][]common.Hash{{topic}},
		[]common.Address{address})
	utils.Assert(err, "Can't subscribe for "+title)
	errChan := s.Err()

	for {
		select {
		case err := <-errChan:
			s.Unsubscribe()
			logger.Error("Events subscription error", "stream", title, err)
			return
		case <-ctx.Done():
			s.Unsubscribe()
			logger.Info("Context cancellation called. Closing event listener")
			return
		case l := <-logs:
			logger.Debug("new log", "stream", title, log15.Alone("event", l.String()))
			p, err := parser(&l)
			if err != nil {
				logger.Error("Can't construct pledge", err)
			} else {
				callback(p)
			}
		}
	}
}

func listenPledge(ctx context.Context, callback listenCallback) {
	listen(ctx, "SWCq DirectPledge", swcq.LogSWCqueueDirectPledge().Id(), addrSWCq,
		swcq.NewPledgeFromDirectPledgeLog, callback)
}

func listenTransfer(ctx context.Context, callback listenCallback) {
	listen(ctx, "SWCq BRG transfer", liquidity.LogBRGusdTransfer().Id(), addrBrg,
		swcq.NewPledgeFromTransfer, callback)
}

func createPledge(p swcq.Pledge) error {
	if err := validatePledge(&p); err != nil {
		return err
	}
	if err := db.Insert(&p); err != nil {
		logger.Error("Can't insert pledge", err)
		return err
	}
	logger.Info("pledge inserted", log15.Alone("pledge", &p))
	return nil
}

func validatePledge(p *swcq.Pledge) errstack.E {
	if !ethereum.AddrEqual(p.CtrAddr.Address, addrSWCq) {
		addr := p.CtrAddr.Hex()
		logger.Debug("Receiving pledge for unsupported contract", "ctr_addr", addr)
		return errstack.NewDomain("Unsupported queue contract: " + addr)
	}
	return nil
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

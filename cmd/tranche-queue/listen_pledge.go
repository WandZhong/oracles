package main

import (
	"context"

	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"bitbucket.org/sweetbridge/oracles/go-lib/liquidity"
	"bitbucket.org/sweetbridge/oracles/go-lib/trancheq"
	"bitbucket.org/sweetbridge/oracles/go-lib/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/robert-zaremba/errstack"
	"github.com/robert-zaremba/log15"
)

type listenCallback func(trancheq.Pledge) error

func listen(ctx context.Context, title string, topic common.Hash, address common.Address,
	parser func(*types.Log) (trancheq.Pledge, errstack.E),
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
				if err := callback(p); err != nil {
					logger.Error("can't process the pledge", err)
				}
			}
		}
	}
}

func listenDirectPledge(ctx context.Context, callback listenCallback) {
	listen(ctx, "SWCq DirectPledge", trancheq.LogSWCqueueDirectPledge().Id(), addrSWCq,
		trancheq.NewPledgeFromDirectPledge, callback)
}

func listenTransfer(ctx context.Context, callback listenCallback) {
	tp := trancheq.NewLogTransferPledger(map[common.Address]liquidity.Currency{
		addrBrg: liquidity.CurrUSD})
	listen(ctx, "SWCq BRG transfer", liquidity.LogBRGusdTransfer().Id(), addrBrg,
		tp.NewPledge, callback)
}

func createPledge(p trancheq.Pledge) error {
	if err := validatePledge(&p); err != nil {
		return err
	}
	if err := db.Insert(&p); err != nil {
		return errstack.WrapAsReq(err, "Can't insert pledge")
	}
	logger.Info("pledge inserted", log15.Alone("pledge", &p))
	return nil
}

func validatePledge(p *trancheq.Pledge) errstack.E {
	if p.CtrAddr.Address != addrSWCq {
		addr := p.CtrAddr.Hex()
		logger.Debug("Receiving pledge from unsupported contract", "ctr_addr", addr)
		return errstack.NewDomain("Unsupported queue contract: " + addr)
	}
	return nil
}

func find() {
	var p trancheq.Pledge
	_, err := db.QueryOne(&p, `SELECT * FROM swc_queue WHERE swc_queue_id = ?`,
		"0x0000000000000000000000000000000000000000000000000000000000000000")
	if err != nil {
		logger.Error("Can't find a pledge", err)
	} else {
		logger.Info("pledge found", log15.Spew(p))
	}
}

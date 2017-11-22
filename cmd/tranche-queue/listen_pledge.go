// Copyright (c) 2017 Sweetbridge Stiftung (Sweetbridge Foundation)
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
		return errstack.WrapAsInf(err, "Can't insert pledge")
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

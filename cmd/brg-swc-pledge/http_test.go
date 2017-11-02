package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"net/url"
	"time"

	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"bitbucket.org/sweetbridge/oracles/go-lib/swcq"
	"bitbucket.org/sweetbridge/oracles/go-lib/test/itest"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/robert-zaremba/log15"
	. "gopkg.in/check.v1"
)

type PledgeS struct{}

func (suite *PledgeS) SetUpSuite(c *C) {
	if !*flagIntegration {
		c.Skip("-integration not provided")
	}
	setupContracts()
}

func (suite *PledgeS) subscribe(c *C) (<-chan types.Log, <-chan struct{}, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 40*time.Second)
	events, s, err := ethereum.SubscribeSimple(ctx, client,
		[][]common.Hash{{swcq.LogSWCqueueDirectPledge().Id()}},
		[]common.Address{pledger.SWCQaddr})
	c.Assert(err, IsNil, Commentf("Can't subscribe for SWC direct pledge events"))
	var out = make(chan types.Log)

	go func() {
		logger.Debug("Start event listening")
		select {
		case err := <-s.Err():
			s.Unsubscribe()
			c.Check(err, IsNil, "Events subscriptoin error")
			cancel()
		case e := <-events:
			s.Unsubscribe()
			s, _ := e.MarshalJSON()
			logger.Info("new log", log15.Alone("event", string(s)))
			out <- e
		case <-ctx.Done():
			c.Error("Context cancelled externally. ", ctx.Err())
			s.Unsubscribe()
		}
	}()

	return out, ctx.Done(), cancel
}

func (suite *PledgeS) TestDirectPledge(c *C) {
	data := url.Values{
		"address":  {"0xce0d46d924cc8437c806721496599fc3ffa268"},
		"brg":      {"10"},
		"currency": {"usd"}}
	rc := itest.NewRoutingPostCtx(data)
	events, done, cancel := suite.subscribe(c)
	if err := httpPostPledge(rc); err != nil {
		cancel()
		c.Assert(err, IsNil)
	}
	resp := rc.Response.(*httptest.ResponseRecorder)
	c.Assert(resp.Code, Equals, http.StatusOK)

	select {
	case <-done:
	case e := <-events:
		c.Assert(e.Address, Equals, pledger.SWCQaddr)
		c.Assert(e.Data, Equals, []byte("00000000000000000000000000ce0d46d924cc8437c806721496599fc3ffa2680000000000000000000000000000000000000000000000008ac7230489e800007573640000000000000000000000000000000000000000000000000000000000"))
	}
}

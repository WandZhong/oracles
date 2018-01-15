// Copyright (c) 2017 Sweetbridge Inc.
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
	"encoding/json"
	"time"

	"bitbucket.org/sweetbridge/oracles/go-lib/liquidity"
	"bitbucket.org/sweetbridge/oracles/go-lib/model"
	"bitbucket.org/sweetbridge/oracles/go-lib/test/itest"
	"bitbucket.org/sweetbridge/oracles/go-lib/trancheq"
	bat "github.com/robert-zaremba/go-bat"
	pgt "github.com/robert-zaremba/go-pgt"
	. "github.com/scale-it/checkers"
	. "gopkg.in/check.v1"
)

type TrancheS struct {
	token   trancheq.Token
	tranche trancheq.Tranche
	now     time.Time
}

func (suite *TrancheS) SetUpSuite(c *C) {
	suite.now = time.Now().UTC()
	suite.token = trancheq.Token{ID: "SWC12"}
	_, err := db.Exec(`DELETE FROM tranches WHERE token_id = ?`, suite.token.ID)
	c.Assert(err, IsNil)
	_, err = db.Exec(`DELETE FROM tokens WHERE token_id = ?`, suite.token.ID)
	c.Assert(err, IsNil)

	c.Assert(model.ValidateSave(&suite.token, db), IsNil)

	endsOn := suite.now.Add(time.Hour * 36)
	suite.tranche = trancheq.Tranche{
		TrancheDB: trancheq.TrancheDB{
			ID:         0, // will be autoasigned
			TokenID:    suite.token.ID,
			CreatedAt:  suite.now,
			StartsAt:   suite.now.Add(time.Hour * 24),
			EndsAt:     &endsOn,
			ExecutesAt: suite.now.Add(time.Hour * 48),
			Supply:     pgt.NewBigInt(10000)},
		Prices: map[liquidity.Currency]float64{
			liquidity.CurrUSD:  1.92,
			liquidity.CurrcETH: 0.001,
		}}
}

func (suite *TrancheS) TearDownSuite(c *C) {
	_, err := db.Exec(`DELETE FROM tranches WHERE token_id = ?`, suite.token.ID)
	c.Assert(err, IsNil)
	c.Assert(db.Delete(&suite.token), IsNil)
}

func (suite *TrancheS) create(expected trancheq.Tranche, c *C) {
	data, err := json.Marshal(expected)
	c.Assert(err, IsNil)

	rc := itest.NewPostJSON(data, c)
	resp := itest.AssertHandlerOK(rc, postTranche, c)
	expected.ID, err = bat.Atoi64(resp.Body.String())
	c.Assert(err, IsNil)
	c.Assert(expected.ID > 0, IsTrue, Comment("obtained id:", expected.ID))

	t, err := trancheq.GetTranche(expected.ID, db)
	c.Assert(err, IsNil)
	c.Check(t.CreatedAt, WithinDuration, suite.now, time.Minute)
	c.Check(t.CreatedAt, Not(Equals), suite.now)
	c.Check(t.StartsAt, TimeEquals, expected.StartsAt)
	c.Check(t.ExecutesAt, TimeEquals, expected.ExecutesAt)
	c.Check(t.EndsAt, TimeEquals, expected.EndsAt)
	c.Check(t.Supply, DeepEquals, expected.Supply)
	c.Check(t.MaxContrib, DeepEquals, expected.MaxContrib)
	c.Assert(db.Delete(&t), IsNil)
}

func (suite *TrancheS) TestPostTranche(c *C) {
	expected := suite.tranche
	suite.create(expected, c)

	expected.MaxContrib = pgt.NewBigInt(601)
	expected.EndsAt = nil
	suite.create(expected, c)
}

func (suite *TrancheS) TestGetTranches(c *C) {
	// firstly let's add a tranche. Token is already there -> SetUpSuite
	c.Assert(suite.tranche.Save(db), IsNil)

	rc := itest.NewRoutingGetCtx("/")
	resp := itest.AssertHandlerOK(rc, getTranches, c)
	var trs TranchesResp
	err := bat.DecodeJSON(resp.Body, &trs)
	c.Assert(err, IsNil)

	var found bool
	for _, t := range trs.Tokens {
		if t.ID == suite.token.ID {
			found = true
			break
		}
	}
	c.Check(found, IsTrue, Comment("Suite token must be included"))

	found = false
	for _, t := range trs.Tranches {
		if t.ID == suite.tranche.ID {
			found = true
			c.Check(t.Prices, HasLen, len(suite.tranche.Prices))
			break
		}
	}
	c.Check(found, IsTrue, Comment("Suite tranche must be included"))
}

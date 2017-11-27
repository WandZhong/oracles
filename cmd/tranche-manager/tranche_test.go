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
	"encoding/json"
	"time"

	"bitbucket.org/sweetbridge/oracles/go-lib/test/itest"
	"bitbucket.org/sweetbridge/oracles/go-lib/trancheq"
	"github.com/robert-zaremba/go-bat"
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
	suite.token.Validate()
	c.Assert(db.Insert(&suite.token), IsNil)

	endsOn := suite.now.Add(time.Hour * 36)
	suite.tranche = trancheq.Tranche{
		ID:          0,
		TokenID:     suite.token.ID,
		CreatedOn:   suite.now,
		StartsOn:    suite.now.Add(time.Hour * 24),
		EndsOn:      &endsOn,
		ExecutesOn:  suite.now.Add(time.Hour * 48),
		Supply:      pgt.NewBigInt(10000),
		PriceBRGusd: 1.92,
	}
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
	c.Check(t.CreatedOn, WithinDuration, suite.now, time.Minute)
	c.Check(t.CreatedOn, Not(Equals), suite.now)
	c.Check(t.StartsOn, TimeEquals, expected.StartsOn)
	c.Check(t.ExecutesOn, TimeEquals, expected.ExecutesOn)
	c.Check(t.EndsOn, TimeEquals, expected.EndsOn)
	c.Check(t.Supply, DeepEquals, expected.Supply)
	c.Check(t.MaxContrib, DeepEquals, expected.MaxContrib)
	c.Assert(db.Delete(&t), IsNil)
}

func (suite *TrancheS) TestPostTranche(c *C) {
	expected := suite.tranche
	suite.create(expected, c)

	expected.MaxContrib = pgt.NewBigInt(601)
	expected.EndsOn = nil
	suite.create(expected, c)
}

func (suite *TrancheS) TestGetTranches(c *C) {
	// firstly let's add a tranche. Token is already there -> SetUpSuite
	c.Assert(db.Insert(&suite.tranche), IsNil)

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
			break
		}
	}
	c.Check(found, IsTrue, Comment("Suite tranche must be included"))
}

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
	"time"

	"bitbucket.org/sweetbridge/oracles/go-lib/test/itest"
	"bitbucket.org/sweetbridge/oracles/go-lib/trancheq"
	pgt "github.com/robert-zaremba/go-pgt"
	. "github.com/scale-it/checkers"
	. "gopkg.in/check.v1"
)

type TokenS struct{}

func (suite *TokenS) create(data string, expected trancheq.Token, c *C) {
	rc := itest.NewPostJSON([]byte(data), c)
	itest.AssertHandlerOK(rc, postToken, c)

	t, err := trancheq.GetToken(expected.ID, db)
	c.Assert(err, IsNil)
	c.Check(t.CreatedAt, WithinDuration, time.Now(), time.Minute)
	c.Check(t.ID, Equals, expected.ID)
	c.Check(t.MaxTotalContrib.Int, DeepEquals, expected.MaxTotalContrib.Int)
	c.Assert(db.Delete(&expected), IsNil)
}

func (suite *TokenS) TestPostToken(c *C) {
	data := `{"id": "SWC1"}`
	t := trancheq.Token{ID: "SWC1"}
	suite.create(data, t, c)

	data = `{"id": "SWC2", "maxTotalContrib": null}`
	t.ID = "SWC2"
	suite.create(data, t, c)

	data = `{"id": "SWC3", "maxTotalContrib": 2141231}`
	t = trancheq.Token{ID: "SWC3", MaxTotalContrib: pgt.NewBigInt(2141231)}
	suite.create(data, t, c)
}

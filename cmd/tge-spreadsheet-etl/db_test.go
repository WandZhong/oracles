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
	"strconv"
	"time"

	"bitbucket.org/sweetbridge/oracles/go-lib/directbuy"
	"bitbucket.org/sweetbridge/oracles/go-lib/members"
	"bitbucket.org/sweetbridge/oracles/go-lib/test/ethereumt"
	"github.com/robert-zaremba/errstack"
	pgt "github.com/robert-zaremba/go-pgt"
)

func createUserAndAccount(email string, ag *ethereumt.AddrGen) errstack.E {
	id := strconv.Itoa(ag.Counter)
	now := time.Now().UTC()
	user := members.Individual{ID: pgt.RandomUUID(), AuthID: "auth_" + id, Email: email,
		EmailVerified: false, Country: "PL", Language: "PL", CreatedAt: now, UpdatedAt: now}
	logger.Debug("inserting new user", "id", user.ID, "email", user.Email)
	if err := db.Insert(&user); err != nil {
		return errstack.WrapAsInfF(err, "can't insert user [%v]", user.ID)
	}
	return directbuy.CreateDistributionAccount(user.ID, ag.Next(), db)
}

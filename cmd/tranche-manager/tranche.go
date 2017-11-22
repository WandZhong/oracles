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

	"bitbucket.org/sweetbridge/oracles/go-lib/model"
	"bitbucket.org/sweetbridge/oracles/go-lib/trancheq"
	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/robert-zaremba/errstack"
)

func postTranche(c *routing.Context) (err error) {
	var obj trancheq.Tranche
	if err := model.DecodeAndSave(c.Request.Body, &obj, db); err != nil {
		return err
	}
	return c.Write(obj.ID)
}

// TranchesResp is a structure for getTranches response
type TranchesResp struct {
	Tokens   []trancheq.Token   `json:"tokens"`
	Tranches []trancheq.Tranche `json:"tranches"`
}

func getTranches(c *routing.Context) (err error) {
	var resp = TranchesResp{[]trancheq.Token{}, []trancheq.Tranche{}}
	if err = db.Model(&resp.Tokens).Select(); err != nil {
		return errstack.WrapAsInf(err, "Can't fetch tokens")
	}
	if err = db.Model(&resp.Tranches).Select(); err != nil {
		return errstack.WrapAsInf(err, "Can't fetch tranches")
	}
	return json.NewEncoder(c.Response).Encode(resp)
}

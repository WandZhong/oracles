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

	"github.com/blockcypher/gobcy"
	"github.com/go-ozzo/ozzo-routing"
	"github.com/robert-zaremba/errstack"
)

type network struct {
	coin  string
	chain string
}

var (
	networks = map[string]network{
		"main": {"btc", "main"},
		"test": {"bcy", "test"},
	}
)

type bcyReturnData struct {
	Address string
	ID      string
}

func handleBtcCreate(ctx *routing.Context) error {
	callBack := ctx.Request.PostFormValue("callBack")

	forwardAddress := ctx.Request.PostFormValue("toAddress")
	if len(forwardAddress) == 0 {
		forwardAddress = *flags.btcPool
	}
	payfwd, err := bcyAPI.CreatePayFwd(gobcy.PayFwd{Destination: forwardAddress, CallbackURL: callBack})
	if err != nil {
		return errstack.WrapAsInf(err, "creating BTC forwarder")
	}
	js, err := json.Marshal(bcyReturnData{payfwd.InputAddr, payfwd.ID})
	if err != nil {
		return errstack.WrapAsInf(err, "unmarshall response")
	}
	return ctx.Write(string(js))
}

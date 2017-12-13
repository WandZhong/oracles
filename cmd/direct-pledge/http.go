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
	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum/wad"
	"bitbucket.org/sweetbridge/oracles/go-lib/liquidity"
	"github.com/go-ozzo/ozzo-routing"
	"github.com/robert-zaremba/errstack"
)

func httpPostPledge(c *routing.Context) (err error) {
	errb := errstack.NewBuilder()
	addr := ethereum.ParseAddressErrp(c.Request.PostFormValue("address"), errb.Putter("address"))
	wei := wad.AfToWei(c.Request.PostFormValue("brg"), errb.Putter("brg"))
	currParam := c.Request.PostFormValue("currency")
	currency := liquidity.ParseCurrencyErrp(currParam, errb.Putter("currency"))
	if errb.NotNil() {
		return errb.ToReqErr()
	}

	logger.Info("BRG-SWC pledge request received", "brg-wei", wei,
		"currency", currParam, "currencyCode", currency,
		"address", addr.Hex())

	txMint, txPledge, err := pledger.Post(addr, wei, currency)
	if err != nil {
		return err
	}
	ethereum.FlogTx(c.Response, "BRG minted for SWCq", txMint)
	ethereum.FlogTx(c.Response, "SWCq direct pledge log recorded", txPledge)
	return nil
}

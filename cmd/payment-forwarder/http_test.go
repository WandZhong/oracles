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

	"net/http"
	"net/http/httptest"
	"net/url"

	"bitbucket.org/sweetbridge/oracles/go-lib/payfwd"
	"bitbucket.org/sweetbridge/oracles/go-lib/test/itest"
	. "gopkg.in/check.v1"
)

type PaymentForwarderS struct{}

func (suite *PaymentForwarderS) SetUpSuite(c *C) {
	if !*flagIntegration {
		c.Skip("-integration not provided")
	}
}

func (suite *PaymentForwarderS) TestCreateBtcForwarder(c *C) {
	data := url.Values{
		"callBack":  {"https://callback.sweetbridge.com/paymentreceived/bBKL63mYGZkw2sTKS"},
		"toAddress": {"CDYMSgTswBZNCg5pDVJEbXBRTyT8z3VxgF"},
	}
	rc := itest.NewRoutingPostCtx(data)
	err := handleBtcCreate(rc)
	c.Assert(err, IsNil)

	res := rc.Response.(*httptest.ResponseRecorder)
	c.Assert(res.Code, Equals, http.StatusOK)

	var returnData bcyReturnData
	err = json.Unmarshal(res.Body.Bytes(), &returnData)
	c.Assert(err, IsNil)
	c.Assert(returnData.Address, NotNil)
}

func (suite *PaymentForwarderS) TestCreateEthForwarder(c *C) {
	data := url.Values{
		"investorId": {"bBKL63mYGZkw2sTKS"},
		"toAddress":  {"0x3532727c1126ddad9a6e9f935b74e41e7b1d4025"},
	}

	rc := itest.NewRoutingPostCtx(data)
	res := itest.AssertHandlerOK(rc, handleEthCreate, c)

	var returnData payfwd.EventForwarderCreated
	err := json.Unmarshal(res.Body.Bytes(), &returnData)
	c.Assert(err, IsNil)
	c.Assert(returnData.Forwarder, NotNil)
}

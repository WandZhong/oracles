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

package itest

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"

	"github.com/go-ozzo/ozzo-routing"
	"gopkg.in/check.v1"
)

// NewRoutingPostCtx returns new POST request with parameters encoded using url-encoding
func NewRoutingPostCtx(values url.Values) *routing.Context {
	// req := httptest.NewRequest("POST", "/", bytes.NewBufferString(values.Encode()))
	var req = http.Request{
		Method:   http.MethodPost,
		PostForm: values}
	return routing.NewContext(httptest.NewRecorder(), &req)
}

// NewPostJSON creates a JSON post request
func NewPostJSON(body []byte, c *check.C) *routing.Context {
	readerBody := bytes.NewReader(body)
	var req, err = http.NewRequest(http.MethodPost, "/", readerBody)
	c.Assert(err, check.IsNil)
	return routing.NewContext(httptest.NewRecorder(), req)
}

// NewRoutingGetCtx creates a new GET request
func NewRoutingGetCtx(url string) *routing.Context {
	req := httptest.NewRequest("GET", "/", nil)
	return routing.NewContext(httptest.NewRecorder(), req)
}

// AssertHandlerOK verifies if a HTTP request is successfull.
func AssertHandlerOK(rc *routing.Context, h routing.Handler, c *check.C) *httptest.ResponseRecorder {
	resp := rc.Response.(*httptest.ResponseRecorder)
	c.Assert(h(rc), check.IsNil,
		check.Commentf("Handler returned error. Status: %v", resp.Code))
	c.Assert(resp.Code, check.Equals, http.StatusOK)
	return resp
}

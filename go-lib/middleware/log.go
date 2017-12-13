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

package middleware

import (
	"net/http"
	"time"

	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/robert-zaremba/errstack"
	goweb "github.com/scale-it/go-web"
	"github.com/scale-it/go-web/handlers"
)

func logRequest(s string) { logger.Debug(s) }

// LogTrace tracks the request HTTP info and time.
// It also logs the server errors.
func LogTrace(c *routing.Context) error {
	t := time.Now()
	w := handlers.WrapWriter(c.Response)
	c.Response = w
	// originalPath := r.URL.Path // this can be overwritten by a middleware
	err := c.Next()
	status := w.Status()
	if err, ok := err.(routing.HTTPError); ok {
		status = err.StatusCode()
	}
	if err != nil {
		errStack, ok := err.(errstack.E)
		// We want to log only server errors
		if (status == 0 || status >= 500) && (!ok || !errStack.IsReq()) {
			logger.Error("Server error", "status", status, err)
		} else {
			logger.Debug("Client Error", "status", status, err)
		}
	}
	// We overwrite status only locally for the Trace
	if status == 0 {
		if err != nil {
			status = http.StatusInternalServerError
		} else {
			status = 200
		}
	}
	goweb.LogRequest(logRequest, c.Request, t, status, w.BytesWritten())
	return err
}

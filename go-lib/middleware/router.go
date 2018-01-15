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

package middleware

import (
	"net/http"
	"strings"

	"github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/content"
	"github.com/go-ozzo/ozzo-routing/cors"
	"github.com/robert-zaremba/errstack"
)

// StdRouter defines standard, default router
// prefix is the routing group which will be use in the service. The leading `/` is added
// automatically if absent in the prefix.
func StdRouter(prefix string) (*routing.Router, *routing.RouteGroup) {
	router := routing.New()
	router.Use(
		ErrorToStatusError,
		LogTrace,
		cors.Handler(cors.AllowAll),
		content.TypeNegotiator(content.JSON))
	if !strings.HasPrefix(prefix, "/") {
		prefix = "/" + prefix
	}
	r := router.Group(prefix)
	r.Get("/health-check", func(ctx *routing.Context) error {
		return ctx.Write("OK")
	})
	return router, r
}

// ErrorToStatusError handles error and convert it's to routing.HTTPError if possible.
// The HTTPError structure holds the error status.
func ErrorToStatusError(c *routing.Context) error {
	err := c.Next()
	if errS, ok := err.(errstack.E); ok {
		if errS.IsReq() {
			return routing.NewHTTPError(http.StatusBadRequest, errS.Error())
		}
		return routing.NewHTTPError(http.StatusInternalServerError, errS.Error())
	}
	return err
}

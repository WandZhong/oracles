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
	"strings"

	"github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/content"
	"github.com/go-ozzo/ozzo-routing/cors"
)

// StdRouter defines standard, default router
// prefix is the routing group which will be use in the service. The leading `/` is added
// automatically if absent in the prefix.
func StdRouter(prefix string) (*routing.Router, *routing.RouteGroup) {
	router := routing.New()
	router.Use(
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

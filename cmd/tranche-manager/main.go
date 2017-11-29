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
	"flag"

	"bitbucket.org/sweetbridge/oracles/go-lib/log"
	"bitbucket.org/sweetbridge/oracles/go-lib/middleware"
	"bitbucket.org/sweetbridge/oracles/go-lib/setup"
	"github.com/go-pg/pg"
	"github.com/robert-zaremba/log15/rollbar"
)

type mainFlags struct {
	setup.PgFlags
	setup.RollbarFlags
	port *string
}

var flags = mainFlags{PgFlags: setup.NewPgFlags(),
	RollbarFlags: setup.NewRollbarFlags(),
	port:         flag.String("port", "8000", "The HTTP listening port")}
var (
	db     *pg.DB
	logger = log.Root()
)

func init() {
	setup.FlagSimpleInit("tranche-manager", *flags.Rollbar, flags.PgFlags, flags.RollbarFlags)
	db = flags.MustConnect()
}

func main() {
	defer rollbar.WaitForRollbar(logger)

	r := middleware.StdRouter()
	// r.Post("/tokens", postToken)
	// r.Post("/tranches", postTranche)
	r.Get("/tranches", getTranches)
	setup.HTTPServer("tranche-manager", *flags.port, r)
}

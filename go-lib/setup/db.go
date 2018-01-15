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

package setup

import (
	"flag"
	"fmt"
	stdlog "log"
	"net"

	"bitbucket.org/sweetbridge/oracles/go-lib/log"
	"github.com/go-pg/pg"
	"github.com/robert-zaremba/errstack"
	bat "github.com/robert-zaremba/go-bat"
)

// pgLogger wraps log.Logger to provide io.Writer interface
type pgLogger struct {
	logger log.Logger
}

func (l pgLogger) Write(p []byte) (int, error) {
	logger.Debug(bat.UnsafeByteArrayToStr(p))
	return len(p), nil
}

// PgFlags represents PostgreSQL flags
type PgFlags struct {
	User *string `long:"user" required:"yes" description:"Username"`
	Pwd  *string `long:"pwd" description:"Password"`
	DB   *string `long:"db" description:"Database name"`
	Addr *string `long:"addr" description:"Server address" default:"localhost:5432"`
}

// NewPgFlags associate PostgreSQL client flags with the structure fields.
// This should be called before flag.Parse or Flag function.
func NewPgFlags() PgFlags {
	return PgFlags{
		User: flag.String("pg-user", "", "PostgreSQL username [required]"),
		Pwd:  flag.String("pg-pwd", "", "PostgreSQL user password"),
		DB:   flag.String("pg-db", "sw", "PostgreSQL database name"),
		Addr: flag.String("pg-addr", "localhost:5432", "PostgreSQL address"),
	}
}

// Check validates the flags.
func (pf PgFlags) Check() error {
	if *pf.User == "" {
		return errstack.NewReq("pg-user must be specifed")
	}
	if *pf.Addr != "" {
		msg := fmt.Sprintf("Malformed pg-addr=%q. Should be host:port", *pf.Addr)
		if _, _, err := net.SplitHostPort(*pf.Addr); err != nil {
			return errstack.WrapAsReq(err, msg)
		}
		if *pf.Addr == ":" {
			return errstack.NewReq(msg)
		}
	}
	return nil
}

// MustConnect initiates Postgersql connection manager. It panics in case of error
func (pf PgFlags) MustConnect() *pg.DB {
	l := pgLogger{logger.New("PG:")}
	pg.SetLogger(stdlog.New(l, "", 0))
	db := pg.Connect(&pg.Options{
		User:     *pf.User,
		Password: *pf.Pwd,
		Database: *pf.DB,
		Addr:     *pf.Addr})

	// let's test the connection
	var x int
	_, err := db.QueryOne(&x, "SELECT 1")
	if err != nil || x != 1 {
		logger.Fatal("Can't connect to the Postgresql", "addr", *pf.Addr, err)
	}
	return db
}

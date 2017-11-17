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

package setup

import (
	stdlog "log"

	"bitbucket.org/sweetbridge/oracles/go-lib/log"
	"github.com/go-pg/pg"
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

// MustPsql initiates Postgersql connection. It panics in case of error
func MustPsql() *pg.DB {
	l := pgLogger{logger.New("PG:")}
	pg.SetLogger(stdlog.New(l, "", 0))
	db := pg.Connect(&pg.Options{
		User:     "swc-queue",
		Password: "password",
		Database: "sw",
		Addr:     "localhost:5432"})

	// let's test the connection
	var x int
	_, err := db.QueryOne(&x, "SELECT 1")
	if err != nil || x != 1 {
		logger.Fatal("Can't connect to the Postgresql", err)
	}
	return db
}

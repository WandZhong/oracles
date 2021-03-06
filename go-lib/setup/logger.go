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
	"strings"

	"bitbucket.org/sweetbridge/oracles/go-lib/log"
	"github.com/robert-zaremba/log15/rollbar"
)

var flagColorLog = flag.Bool("log-colored", true, "Use colored log (good for terminal output)")

// MustLogger setups logger
func MustLogger(appname string, rollbartoken string) {
	assert(checkAppName(appname))
	appname = envName + "-" + appname
	rc := rollbar.Config{
		Version: GitVersion,
		Env:     appname,
		Token:   rollbartoken}
	// it's OK. Logger is a pointer, so we don't need to overwrite the global object
	_, err := log.New(log.RootName,
		log.Config{Color: *flagColorLog, TimeFmt: "sec", Level: "DEBUG"}, rc)
	assert(err)
	if strings.HasPrefix(envName, "prod") && rollbartoken == "" {
		logger.Error("Rollbar token must be set in production environment")
	}
	logger.Debug("Logger initialized", "app_version", GitVersion)
}

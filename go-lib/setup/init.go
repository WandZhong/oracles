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

/* Package providings setup / initialization of common structures, drivers, ...
 * for other tools and applications
 */

package setup

import (
	"os"
	"strings"

	"bitbucket.org/sweetbridge/oracles/go-lib/log"
)

// GitVersion should be substituted during build time by the git version. This is done
// using go linker flags:
// -ldflags "-X bitbucket.org/sweetbridge/oracles/go-lib/setup.GitVersion=$(git describe)
var GitVersion string

// envName is the application stage environment name (eg production, dev, backstage, ...).
// It is set using `SB_ENV` environment variable.
var envName string

var logger = log.Root()

// init initializes packages
func init() {
	if envName = os.Getenv("SB_ENV"); envName == "" {
		envName = "dev"
	} else {
		envName = strings.ToLower(envName)
		assert(checkAppName(envName))
	}
	if envName == "prod" {
		envName = "production"
	}
}

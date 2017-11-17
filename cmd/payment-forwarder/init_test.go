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
	"path"
	"testing"

	"bitbucket.org/sweetbridge/oracles/go-lib/setup"
	. "gopkg.in/check.v1"
)

var flagIntegration = flag.Bool("integration", false, "Include integration tests")

func Test(t *testing.T) { TestingT(t) }

func init() {
	setupFlags()

	flags.bcyNet = ptr("test")
	flags.PkHex = ptr("fdb2886b1ff5a0e60f9a4684e385aa7b77f064730304143f08ba96ca1a17effa")
	timeout := 600
	flags.txTimeout = &timeout
	flags.ContractsPath = ptr(path.Join(setup.Root(),
		"submodules/contract-deployments/development"))
	flags.Network = ptr("backstage-dev")

	initBcyAPI()
	setupContracts()

	Suite(&PaymentForwarderS{})
}

func ptr(x string) *string {
	return &x
}

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

//-bcy-key=632dec063dad470da415dd386041e377 -bcy-net=test -contracts=x -pk-hex=21637e41ea5ec2c0a08c734ddcf4fd3fc01eaa7618b412e09b76f86808cb4dd8 -eth-network=development -btc-pool=CDYMSgTswBZNCg5pDVJEbXBRTyT8z3VxgF -eth-pool=0x55f0a1a581E669f6d352849A895f2c7bA69Da872
func init() {
	setupFlags()

	flags.PkHex = ptr("fdb2886b1ff5a0e60f9a4684e385aa7b77f064730304143f08ba96ca1a17effa")
	timeout := 600
	flags.txTimeout = &timeout
	flags.ContractsPath = ptr(path.Join(setup.Root(), "submodules/contract-deployments/development"))
	flags.NetworkConfig = ptr(path.Join(setup.Root(), "../config/eth-networks.json"))
	flags.Network = ptr("backstage-dev")

	setupContracts()

	Suite(&PaymentForwarderS{})
}

func ptr(x string) *string {
	return &x
}

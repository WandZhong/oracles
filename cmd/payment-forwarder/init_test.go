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

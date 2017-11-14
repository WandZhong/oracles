package main

import (
	"flag"
	"testing"

	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	. "gopkg.in/check.v1"
)

var flagIntegration = flag.Bool("integration", false, "Include integration tests")
var cf ethereum.ContractFactory

func Test(t *testing.T) { TestingT(t) }
func init() {
	setupFlags()
	cf = setupContracts()

	Suite(&PledgeS{})
}

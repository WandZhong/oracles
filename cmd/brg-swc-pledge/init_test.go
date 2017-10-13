package main

import (
	"flag"
	"testing"

	. "gopkg.in/check.v1"
)

var flagIntegration = flag.Bool("integration", false, "Include integration tests")

func Test(t *testing.T) { TestingT(t) }
func init() {
	setupFlags()

	Suite(&PledgeS{})
}

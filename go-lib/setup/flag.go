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
	"flag"
	"fmt"
	"os"
)

// FlagFail - exits the main process and displays usage information
func FlagFail(err error) {
	logger.Error("!! Wrong CMD parameters !! Run with `-h` parameter to output a usage information", err)

	os.Exit(1)
}

// Flag parses flag and setups usage function.
// `positionalArgs` is a string representing positional args, eg: "arg1 arg2 arg3"
// `commands` if provided is a positional argument command description.
func Flag(positionalArgs string, commands ...string) {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr,
			"\nusage: %s <parameters> %s:\n", os.Args[0], positionalArgs)
		if len(commands) != 0 {
			fmt.Fprint(os.Stderr, "COMMANDS:", commands[0], "\n\n")
		}
		fmt.Fprintln(os.Stderr, "PARAMETERS:")
		flag.PrintDefaults()
	}
	flag.Parse()
}

// FlagValidate runs flag checkers
func FlagValidate(checkers ...Checker) {
	for _, c := range checkers {
		if err := c.Check(); err != nil {
			FlagFail(err)
		}
	}
}

// FlagSimpleInit provides a common functionality to setup the command line flags.
// `positionalArgs` documents expected positional argumentes, eg `"arg1 arg2 arg3"`.
// `rollbarKey` is a pointer, because it can be a flag, which is gonig to be initialized
//    in this function.
func FlagSimpleInit(name, positionalArgs string, rollbarKey *string, flags ...Checker) {
	Flag(positionalArgs)
	FlagValidate(flags...)
	MustLogger(name, *rollbarKey)
}

// Checker is an interface for type which has a Check function
type Checker interface {
	Check() error
}

// RollbarFlags wraps flags for Rollbar client
type RollbarFlags struct {
	Rollbar *string `long:"t" description:"rollbar token" default:""`
}

// NewRollbarFlags setups flags for Rollbar client
func NewRollbarFlags() RollbarFlags {
	return RollbarFlags{flag.String("rollbar", "", "rollbar token [required in production env]")}
}

// Check is a dummy Check interface implementation
func (r RollbarFlags) Check() error {
	return nil
}

// BaseOracleFlags represents common oracle flags
type BaseOracleFlags struct {
	EthFlags
	RollbarFlags
}

// NewBaseOracleFlags setups common flags
func NewBaseOracleFlags() BaseOracleFlags {
	return BaseOracleFlags{
		NewEthFlags(), NewRollbarFlags()}
}

// Check validates the flags. It may panic!
func (b BaseOracleFlags) Check() error {
	return b.EthFlags.Check()
}

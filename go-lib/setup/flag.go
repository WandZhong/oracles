package setup

import (
	"flag"
	"fmt"
	"os"
)

// FlagFail - exits the main process and displays usage information
func FlagFail(err error) {
	logger.Error("Wrong CMD parameters", err)
	flag.Usage()
	os.Exit(1)
}

// Flag parses flag and setups usage function.
// `positionalArgs` is a string representing positional args
// `commands` if provided is a positional argument command description.
func Flag(positionalArgs string, commands ...string) {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr,
			"Usage of %s <parameters> %s:\n", os.Args[0], positionalArgs)
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

// Checker is an interface for type which has a Check function
type Checker interface {
	Check() error
}

// RollbarFlags wraps flags for Rollbar client
type RollbarFlags struct {
	Rollbar *string
}

// NewRollbarFlags setups flags for Rollbar client
func NewRollbarFlags() RollbarFlags {
	return RollbarFlags{flag.String("rollbar", "", "rollbar token [required in production env]")}
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

package setup

import (
	"flag"
	"fmt"
	"os"
)

// FlagFail - exits the main process and displays usage information
func FlagFail() {
	logger.Error("Wrong CMD parameters")
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

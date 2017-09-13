package utils

import (
	"flag"
	"os"
)

// FlagFail - exits the main process and displays usage information
func FlagFail() {
	logger.Error("Wrong CMD parameters")
	flag.Usage()
	os.Exit(1)
}

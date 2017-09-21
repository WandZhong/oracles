package setup

import (
	"flag"
	"fmt"
	"os"

	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"github.com/robert-zaremba/errstack"
	bat "github.com/robert-zaremba/go-bat"
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

// EthFlags represents common Ethereum client flags
type EthFlags struct {
	PkFile        *string
	PkPwd         *string
	ContractsPath *string
	Network       *string
	Host          *string
}

// NewEthFlags associsate ethereum client flags with the structure fields.
// This should be called before flag.Parse or Flag function.
func NewEthFlags() EthFlags {
	return EthFlags{
		PkFile:        flag.String("pk", "", "path to the private key json file [required]"),
		PkPwd:         flag.String("pwd", "", "key file password [required]"),
		ContractsPath: flag.String("contracts", "", "path to contract schemas directory [required]"),
		Network:       flag.String("network", "", "blockchain network name [required]"),
		//		Host:          flag.String("host", "", "ethereum node address. 'http' prefix added automatically. If not provided, ethNetworkAddrs is used to get the address")
	}
}

// Check validates the flags. It may panic!
func (ef EthFlags) Check() error {
	if err := bat.IsFile(*ef.PkFile); err != nil {
		return errstack.WrapAsReq(err, "Wrong pkFile")
	}
	if _, ok := ethNetworkAddrs[*ef.Network]; !ok {
		return errstack.NewReq("unknown network: " + *ef.Network)
	}
	if *ef.PkPwd == "" || *ef.ContractsPath == "" {
		return errstack.NewReq("All equired arguments must be specified")
	}
	return nil
}

// MustNewTxrFactory creates TxrFactory based on command flags.
// It panics in case of error.
func (ef EthFlags) MustNewTxrFactory() ethereum.TxrFactory {
	p, err := ethereum.NewJSONTxrFactory(*ef.PkFile, *ef.PkPwd)
	if err != nil {
		logger.Fatal("Can't create TxrFactory based on JSON file", "filename", *ef.PkFile)
	}
	return p
}

package setup

import (
	"flag"

	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/robert-zaremba/errstack"
	bat "github.com/robert-zaremba/go-bat"
)

// EthFlags represents common Ethereum client flags
type EthFlags struct {
	PkFile        *string
	PkPwd         *string
	ContractsPath *string
	Network       *string
	Host          *string
}

// NewEthFlags associate ethereum client flags with the structure fields.
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
	if _, ok := ethNetworkMap[*ef.Network]; !ok {
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

// MustEthFactory creates ethclient and contract factory based on flag options
func (ef EthFlags) MustEthFactory() (*ethclient.Client, ethereum.ContractFactory) {
	txrF := ef.MustNewTxrFactory()
	return MustEthFactory(*ef.Network, *ef.ContractsPath, txrF)
}

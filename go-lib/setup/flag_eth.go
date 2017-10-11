package setup

import (
	"flag"
	"fmt"

	"bitbucket.org/sweetbridge/oracles/go-lib/ethereum"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/robert-zaremba/errstack"
	bat "github.com/robert-zaremba/go-bat"
)

// EthFlags represents common Ethereum client flags
type EthFlags struct {
	PkHex         *string
	PkFile        *string
	PkPwd         *string
	ContractsPath *string
	Network       *string
}

// NewEthFlags associate ethereum client flags with the structure fields.
// This should be called before flag.Parse or Flag function.
func NewEthFlags() EthFlags {
	return EthFlags{
		PkHex:         flag.String("pk-hex", "", "private key - takes precedence over `pk-file` + `pk-pwd` [required]"),
		PkFile:        flag.String("pk-file", "", "path to the private key json file [required if `pk-hex` not specified]"),
		PkPwd:         flag.String("pk-pwd", "", "key file password [required if `pk-hex` not specified]"),
		ContractsPath: flag.String("contracts", "", "path to contract schemas directory [required]"),
		Network: flag.String("eth-network", "",
			fmt.Sprintf("blockchain network name [required]\n\tAvailable options: %s", ethNetworks)),
	}
}

// Check validates the flags. It may panic!
func (ef EthFlags) Check() error {
	if _, ok := ethNetworkMap[*ef.Network]; !ok {
		return errstack.NewReq("unknown network: " + *ef.Network)
	}
	if *ef.PkHex != "" {
		if _, err := crypto.HexToECDSA(*ef.PkHex); err != nil {
			return errstack.WrapAsReq(err,
				"Can't parse ECDSA key. Expected valid hex string.")
		}
	} else {
		if err := bat.IsFile(*ef.PkFile); err != nil {
			return errstack.WrapAsReq(err, "Wrong pkFile")
		}
	}
	if *ef.ContractsPath == "" {
		return errstack.NewReq("All equired arguments must be specified")
	}
	return nil
}

// mustNewTxrFactory creates TxrFactory based on command flags.
// It panics in case of error.
func (ef EthFlags) mustNewTxrFactory() ethereum.TxrFactory {
	if *ef.PkHex != "" {
		p, err := ethereum.NewPrivKeyTxrFactory(*ef.PkHex)
		if err != nil {
			logger.Fatal("Malformed plain (hex) private key")
		}
		return p
	}
	p, err := ethereum.NewJSONTxrFactory(*ef.PkFile, *ef.PkPwd)
	if err != nil {
		logger.Fatal("Can't create TxrFactory based on JSON file", "filename", *ef.PkFile, err)
	}
	return p
}

// MustEthFactory creates ethclient and contract factory based on flag options
func (ef EthFlags) MustEthFactory() (*ethclient.Client, ethereum.ContractFactory) {
	txrF := ef.mustNewTxrFactory()
	return MustEthFactory(*ef.Network, *ef.ContractsPath, txrF)
}

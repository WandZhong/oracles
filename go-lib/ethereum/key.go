package ethereum

import (
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/robert-zaremba/errstack"
)

// for private key: crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")

// NewKeyTxr creates new transactor using a hex string of a ECDSA key.
func NewKeyTxr(hexkey string) (*bind.TransactOpts, errstack.E) {
	key, err := crypto.HexToECDSA(hexkey)
	if err != nil {
		return nil, errstack.WrapAsReq(err,
			"Can't parse ECDSA key file. Expected valid hex string.")
	}
	return bind.NewKeyedTransactor(key), nil
}

// MustKeyTxr calls NewKeyTxr and panics if it get's an error.
func MustKeyTxr(hexkey string) *bind.TransactOpts {
	tx, err := NewKeyTxr(hexkey)
	if err != nil {
		logger.Fatal("Can't make transactor from hex string", "key", hexkey, err)
	}
	return tx
}

// NewFileTxr creates new transactor using the JSON key file
func NewFileTxr(JSONFile string, password string) (*bind.TransactOpts, errstack.E) {
	kr, err := os.Open(JSONFile)
	if err != nil {
		return nil, errstack.WrapAsReq(err, "Can't open JSON key file")
	}
	res, err := bind.NewTransactor(kr, password)
	if err != nil {
		return res, errstack.WrapAsReq(err, "Can't decode the file")
	}
	return res, nil
}

// MustFileTxr calls JKeyTxr and panics if it get's an error.
func MustFileTxr(JSONFile string, password string) *bind.TransactOpts {
	tx, err := NewFileTxr(JSONFile, password)
	if err != nil {
		logger.Fatal("Can't make transactor from JSON pk file", "file", JSONFile, err)
	}
	return tx
}

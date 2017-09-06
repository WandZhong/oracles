package ethereum

import (
	"encoding/hex"
	"encoding/json"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/robert-zaremba/errstack"
	bat "github.com/robert-zaremba/go-bat"
)

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

// KeySimple is a simple version of keystore.Key structure
type KeySimple struct {
	Address common.Address
	ID      string
	Version int
}

// UnmarshalJSON implement interface for json.Unmrashall
func (k *KeySimple) UnmarshalJSON(j []byte) (err error) {
	var tmp = struct {
		Address string `json:"address"`
		ID      string `json:"id"`
		Version int    `json:"version"`
	}{}
	err = json.Unmarshal(j, &tmp)
	if err != nil {
		return nil
	}
	addr, err := hex.DecodeString(tmp.Address)
	if err != nil {
		return err
	}
	k.ID = tmp.Address
	k.Version = tmp.Version
	k.Address = common.BytesToAddress(addr)
	return nil
}

// ReadKeySimple reads the JSON key into the key object
func ReadKeySimple(JSONFile string) (KeySimple, errstack.E) {
	var k KeySimple
	return k, bat.DecodeJSONFile(JSONFile, &k, logger)
}

// MustReadKeySimple calls ReadKeySimple and panics if it get's an error
func MustReadKeySimple(JSONFile string) KeySimple {
	k, err := ReadKeySimple(JSONFile)
	if err != nil {
		logger.Fatal("Can't read file", "path", JSONFile, err)
	}
	return k
}

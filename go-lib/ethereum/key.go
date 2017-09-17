package ethereum

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/robert-zaremba/errstack"
	bat "github.com/robert-zaremba/go-bat"
)

// TxrFactory wraps parameters to create bind.TransactOpts
type TxrFactory interface {
	Txo() *bind.TransactOpts
}

type txrFactory struct {
	privKey *ecdsa.PrivateKey
}

// NewJSONTxrFactory creates TxrFactory using on JSON account file and passphrase
func NewJSONTxrFactory(filename, passphrase string) (TxrFactory, errstack.E) {
	data, err := bat.ReadFile(filename, logger)
	if err != nil {
		return nil, err
	}
	key, errStd := keystore.DecryptKey(data, passphrase)
	if errStd != nil {
		return nil, errstack.WrapAsReq(err, "Wrong passphrase")
	}
	return txrFactory{key.PrivateKey}, nil
}

// NewPrivKeyTxrFactory creates new transactor using a hex string of a ECDSA key.
func NewPrivKeyTxrFactory(hexkey string) (TxrFactory, errstack.E) {
	key, err := crypto.HexToECDSA(hexkey)
	if err != nil {
		return nil, errstack.WrapAsReq(err,
			"Can't parse ECDSA key. Expected valid hex string.")
	}
	return txrFactory{key}, nil
}

// Txo implements TxrFactory interface
func (tp txrFactory) Txo() *bind.TransactOpts {
	return bind.NewKeyedTransactor(tp.privKey)
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

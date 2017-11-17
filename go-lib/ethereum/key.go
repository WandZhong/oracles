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
	Addr() common.Address
}

type txrFactory struct {
	privKey *ecdsa.PrivateKey
	addr    common.Address
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
	return txrFactory{key.PrivateKey, key.Address}, nil
}

// NewPrivKeyTxrFactory creates new transactor using a hex string of a ECDSA key.
func NewPrivKeyTxrFactory(hexkey string) (TxrFactory, errstack.E) {
	key, err := crypto.HexToECDSA(hexkey)
	if err != nil {
		return nil, errstack.WrapAsReq(err,
			"Can't parse ECDSA key. Expected valid hex string.")
	}
	addr := crypto.PubkeyToAddress(key.PublicKey)
	return txrFactory{key, addr}, nil
}

// Txo implements TxrFactory interface
func (tp txrFactory) Txo() *bind.TransactOpts {
	return bind.NewKeyedTransactor(tp.privKey)
}

// Txo implements TxrFactory interface
func (tp txrFactory) Addr() common.Address {
	return tp.addr
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

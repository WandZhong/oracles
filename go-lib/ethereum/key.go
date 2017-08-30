package ethereum

import (
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/robert-zaremba/errstack"
)

// JKeyTransactor creates new ethereum transactor using the JSON key file
func JKeyTransactor(JSONFile string, password string) (*bind.TransactOpts, errstack.E) {
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
